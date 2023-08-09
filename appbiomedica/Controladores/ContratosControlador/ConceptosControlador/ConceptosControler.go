package ConceptosControlador

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/CatalogosModel"
	"appbiomedica/Modelos/ContratoModel/ConceptoContratoModelo"
	"appbiomedica/Modelos/ContratoModel/ContratoModelo"
	"appbiomedica/Modelos/ContratoModel/DistribucionModelo"
	"appbiomedica/Modulos/Estructuras"
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
	"strings"
)

type ConceptosNuevos struct {
	ConceptoContrato []ConceptoContratoModelo.ConceptoContrato `gorm:"-"`
}

//********************************************************************************************************
/* *********************    Controlador para Agregar un nuevo contrato     ********************* */
func PostConceptoContrato(c *gin.Context) {
	var concepto ConceptosNuevos
	id_contrato := c.DefaultQuery("id_contrato", "")
	err := c.BindJSON(&concepto)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error al castear la estructura", err.Error(), c)
		return
	}
	//********************************************************************************************************
	/* *********************    inicio mi transaccion de mi base de datos     ********************* */
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al iniciar la transaccion", tx.Error, c)
		return
	}
	var conceptos ConceptoContratoModelo.ConceptoContrato
	for _, conc := range concepto.ConceptoContrato {
		conceptos.Id = bson.NewObjectId().Hex()
		conceptos.PrecioUniSnIva = conc.PrecioUniSnIva
		conceptos.CantidadConcepto = conc.CantidadConcepto
		conceptos.CantidadAmpliada = conc.CantidadAmpliada
		conceptos.Marca = strings.ToUpper(conc.Marca)
		conceptos.Modelo = strings.ToUpper(conc.Modelo)
		conceptos.ObjetoContratacion = conc.ObjetoContratacion
		conceptos.FechaMaxEntrega = conc.FechaMaxEntrega
		conceptos.GarantiaBienes = conc.GarantiaBienes
		conceptos.PreiIdArticulo = conc.PreiIdArticulo
		conceptos.ContratoNumeroContrato = id_contrato
		duplicado, resultado := ConceptoContratoModelo.VerificaEquipo(conc.PreiIdArticulo, id_contrato)
		if duplicado == true {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "Este Equipo ya ha sido agregado", resultado, c)
			return
		}
		err := conceptos.AgregarConceptoContrato(tx)
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "Erros al inserta el concepto", err.Error(), c)
			return
		}
	}
	//********************************************************************************************************
	/* *********************   finalizar insercion y hacer el commit     ********************* */
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al realizar el commit", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos agregados correctamente", concepto, c)
	return
}

//********************************************************************************************************
/* *********************   Funcion para consultar todos los conceptos de un contrato     ********************* */
func GetConceptos(c *gin.Context) {
	id_contrato := c.DefaultQuery("id_contrato", "")
	if id_contrato == "" {
		Estructuras.Responder(http.StatusInternalServerError, "Debes ingrear el id contrato ", nil, c)
		return
	}
	err, resultado, suma := ConceptoContratoModelo.ConsultaConceptos(id_contrato)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente",
		gin.H{"total": suma, "registros": resultado}, c)
	return
}

//********************************************************************************************************
/* *********************   Funcion para actualizar un concepto de un contrato     ********************* */
func UpdateConcepto(ctx *gin.Context) {
	id_concepto := ctx.DefaultQuery("id_concepto", "")
	num_contrato := ctx.DefaultQuery("numero_contrato", "")
	var concepto ConceptoContratoModelo.ConceptoContrato
	var total_actualizar float64
	err := ctx.BindJSON(&concepto)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error en la estructura", err.Error(), ctx)
		return
	}
	fmt.Println(id_concepto, concepto)
	/*Buscar el monto total del contrato y el concepto*/
	_, id_prei := ConceptoContratoModelo.VerificaEquipo(concepto.PreiIdArticulo, num_contrato)
	if id_prei == nil {
		Estructuras.Responder(http.StatusBadRequest, "No se encontro el concepto en la base de datos",
			nil, ctx)
	}
	/*Calculando el nuevo total*/
	total_actualizar = float64(concepto.CantidadConcepto) * concepto.PrecioUniSnIva
	/****Verificando si hay diferencia en el monto total a actualizar*****/
	total_registrado := id_prei.PrecioUniSnIva * float64(id_prei.CantidadConcepto)
	if total_actualizar != total_registrado {
		/*Se validan que el precio o la cantida nuevas, no superen el monto total del contrato*/
		err, _, total := ConceptoContratoModelo.ConsultaConceptos(num_contrato)
		if err != nil {
			Estructuras.Responder(http.StatusBadRequest, "Error al verificar el monto registrado",
				err.Error(), ctx)
		}
		fmt.Println("Monto total del contrato: ", id_prei.Contrato.MontoTotal)
		total_restado := total - total_registrado
		fmt.Println("Restando el total del concepto", total, "-", total_registrado, "=", total_restado)
		monto_total_actualizar := total_restado + total_actualizar
		fmt.Println("total nuevo", total_restado, "+", total_actualizar, "=", monto_total_actualizar)
		if monto_total_actualizar > id_prei.Contrato.MontoTotal {
			Estructuras.Responder(http.StatusBadRequest,
				"Error al actualizar",
				errors.New("La cantidad o el precio por actualizar supera el monto total "+"registrado en el contrato").Error(),
				ctx)
			return
		}
	}
	concepto.Marca = strings.ToUpper(concepto.Marca)
	concepto.Modelo = strings.ToUpper(concepto.Modelo)
	fmt.Println(concepto)
	err = concepto.ActualizarConcepto(id_concepto)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al actualizar el concepto", err.Error(), ctx)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos actualizados correctamente", concepto, ctx)
	return
}

//********************************************************************************************************
/* *********************    funcion para leer un excel para verificar el contenido      ********************* */
func ValidarConceptosArchivo(c *gin.Context) {
	file, _ := c.FormFile("file")
	f, err := file.Open()
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al abrir el archivo", err, c)
		return
	}
	fmt.Println(f)
	defer f.Close()
	validado, errores, err := ValidarArchivo(f)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, err.Error(), errores, c)
		return
	}
	if validado == false {
		Estructuras.Responder(http.StatusInternalServerError, "Errores de Validacion", errores, c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Archivo Validado Correctamente", nil, c)
	return
}

//********************************************************************************************************
/* *********************    Controlador para agregar conceptos desde un archivo excel    ********************* */
func AgreegarConceptosArchivo(c *gin.Context) {
	var conceptos ConceptoContratoModelo.ConceptoContrato
	file, _ := c.FormFile("file")
	f, err := file.Open()
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al abrir el archivo", err, c)
		return
	}
	validado, errores, err := ValidarArchivo(f)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, err.Error(), errores, c)
		return
	}
	if validado == false {
		Estructuras.Responder(http.StatusInternalServerError, "Errores de Validacion", errores, c)
		return
	}
	//********************************************************************************************************
	/* *********************    inicio mi transaccion de mi base de datos     ********************* */
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al iniciar la transaccion", tx.Error, c)
		return
	}
	fi, err := file.Open()
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al abrir el archivo", err, c)
		return
	}
	ar, err := excelize.OpenReader(fi)
	if err != nil {
		println(err.Error(), "---error")
		Estructuras.Responder(http.StatusInternalServerError, "Error al abrir el archivo", err.Error(), c)
		return
	}
	rows, _ := ar.GetRows("Conceptos")
	//if err != nil {
	//	println(err.Error())
	//	Estructuras.Responder(http.StatusInternalServerError, "Error al abrir el archivo",
	//		errors.New("No se pudo leer el archivo o no tiene el formato correcto"), c)
	//	return
	//}
	fmt.Println(len(rows))
	for num, fila := range rows {
		fmt.Println("Entreeeeee", len(fila), num)
		if len(fila) == 8 && num > 0 {
			fmt.Println("Entre")
			precio, _ := strconv.ParseFloat(fila[0], 32)
			cantidad, _ := strconv.Atoi(fila[1])
			garantia, _ := strconv.Atoi(fila[5])
			_, clave := CatalogosModel.ConsultaUnaClave("", fila[6])
			conceptos.Id = bson.NewObjectId().Hex()
			conceptos.PrecioUniSnIva = float64(precio)
			conceptos.CantidadConcepto = cantidad
			conceptos.CantidadAmpliada = 0
			conceptos.CantidadDistribuida = 0
			conceptos.Marca = strings.ToUpper(fila[2])
			conceptos.Modelo = strings.ToUpper(fila[3])
			conceptos.ObjetoContratacion = clave.CuadroBasico.Grupo
			conceptos.FechaMaxEntrega = fila[4]
			conceptos.GarantiaBienes = garantia
			conceptos.PreiIdArticulo = clave.IdArticulo
			conceptos.ContratoNumeroContrato = fila[7]
			fmt.Println(conceptos)
			duplicado, resultado := ConceptoContratoModelo.VerificaEquipo(fila[6], fila[7])
			if duplicado == true {
				tx.Rollback()
				Estructuras.Responder(http.StatusInternalServerError, "Este Equipo ya ha sido agregado", resultado, c)
				return
			}
			err := conceptos.AgregarConceptoContrato(tx)
			if err != nil {
				tx.Rollback()
				Estructuras.Responder(http.StatusInternalServerError, "Erros al inserta el concepto", err.Error(), c)
				return
			}
		}
	}
	//********************************************************************************************************
	/* *********************   finalizar insercion y hacer el commit     ********************* */
	defer f.Close()
	defer fi.Close()
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al realizar el commit", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos agregados correctamente", nil, c)
	return
}

//********************************************************************************************************
/* *********************    Controlador para eliminar un concepto de un contrato     ********************* */
func DeleteConcepto(c *gin.Context) {
	id_concepto := c.DefaultQuery("id_concepto", "")
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al conectar a la base de datos", tx.Error, c)
		return
	}
	err, distribucion := DistribucionModelo.ConsultaDistribucion(id_concepto)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar distribucion", err.Error(), c)
		return
	}
	if len(*distribucion) > 0 {
		for _, distri := range *distribucion{
			/*Buscar cuantas distribuciones tienen la misma delegacion en el mismo contrato*/
			err, dis := DistribucionModelo.BuscarDistribucion(distri.ConceptoContrato.ContratoNumeroContrato,
				distri.UnidadMed.DelegacionesClvDele)
			if err != nil {
				tx.Rollback()
				Estructuras.Responder(http.StatusInternalServerError, "Error al buscar la distribucion", err.Error(), c)
				return
			}
			if len(*dis) <= 1 {
				err, _ = ContratoModelo.EliminarDeleContr(distri.ConceptoContrato.ContratoNumeroContrato,
					distri.UnidadMed.DelegacionesClvDele, tx)
				if err != nil {
					tx.Rollback()
					Estructuras.Responder(http.StatusInternalServerError, "Error al eliminar la delegacion del contrato",
						err.Error(), c)
					return
				}
			}
		}
	}
	err, resultado := ConceptoContratoModelo.EliminarConcepto(id_concepto, tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al eliminar el registro", err.Error(), c)
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al realizar el commit", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Concepto Eliminado Correctamente", resultado, c)
	return
}
