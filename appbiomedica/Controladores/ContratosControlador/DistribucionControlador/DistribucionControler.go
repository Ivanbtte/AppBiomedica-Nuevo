package DistribucionControlador

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/CatalogosModel"
	"appbiomedica/Modelos/ContratoModel/ConceptoContratoModelo"
	"appbiomedica/Modelos/ContratoModel/ContratoModelo"
	"appbiomedica/Modelos/ContratoModel/DistribucionModelo"
	"appbiomedica/Modelos/ModeloUnidadM"
	"appbiomedica/Modulos/Estructuras"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
	"strings"
)

//********************************************************************************************************
/* *********************   Funcion para consultar toda la distribucion de un concepto     ********************* */
func GetDistribucion(c *gin.Context) {
	id_concepto := c.DefaultQuery("id_concepto", "")
	if id_concepto == "" {
		Estructuras.Responder(http.StatusInternalServerError, "Debes ingrear el id del concepto ", nil, c)
		return
	}
	err, resultado := DistribucionModelo.ConsultaDistribucion(id_concepto)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}

//********************************************************************************************************
/* *********************   Funcion para validar el archivo para agregar distribucion     ********************* */
func ValidarArchivoDistribucion(c *gin.Context) {
	distribucion := c.DefaultQuery("distribucion", "")
	file, _ := c.FormFile("file")
	f, err := file.Open()
	if err != nil {
		return
	}
	defer f.Close()
	if distribucion == "distribucion_sin_concepto" {
		respuesta, err, validado := ValidarDistribucionSinConceptos(f)
		if err != nil {
			Estructuras.Responder(http.StatusInternalServerError, err.Error(), respuesta, c)
			return
		}
		if validado == false {
			Estructuras.Responder(http.StatusInternalServerError, "Errores de Validacion", respuesta, c)
			return
		}
	}
	if distribucion == "distribucion_con_conceptos" {
		fmt.Println(file.Filename, "con conceptos")
		respuesta, err, validado := ValidarDistribucionConConceptos(f)
		if err != nil {
			Estructuras.Responder(http.StatusInternalServerError, err.Error(), respuesta, c)
			return
		}
		if validado == false {
			Estructuras.Responder(http.StatusInternalServerError, "Errores de Validacion", respuesta, c)
			return
		}
	}
	Estructuras.Responder(http.StatusOK, "Datos validados correctamente", nil, c)
	return
}

//********************************************************************************************************
/* ****************    Funcion para agregar Distribucion con conceptos registrados por archivo     **************** */
func AgregarArchivoDistribucion(c *gin.Context) {
	var ConcepD DistribucionModelo.DistribucionUnidadMed
	distribucion := c.DefaultQuery("distribucion", "")
	file, _ := c.FormFile("file")
	f, err := file.Open()
	if err != nil {
		return
	}
	defer f.Close()
	if distribucion == "distribucion_con_conceptos" {
		respuesta, err, validado := ValidarDistribucionConConceptos(f)
		if err != nil {
			Estructuras.Responder(http.StatusInternalServerError, err.Error(), respuesta, c)
			return
		}
		if validado == false {
			Estructuras.Responder(http.StatusInternalServerError, "Errores de Validacion", respuesta, c)
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
		rows, err := ar.GetRows("Distribucion")
		if err != nil {
			Estructuras.Responder(http.StatusInternalServerError,
				"Error al leer el archivo o no es el archivo original", err.Error(), c)
			return
		}
		for num, row := range rows {
			if num > 0 {
				cantidad, _ := strconv.Atoi(row[3])
				/*Validando que el id prei (equipo) ya se encuentre registrado en el mismo oontrato*/
				duplicado, equipo := ConceptoContratoModelo.VerificaEquipo(row[1], row[0])
				if duplicado == false {
					Estructuras.Responder(http.StatusInternalServerError,
						"Este Id Prei no ha sido registrado en el contrato", err.Error(), c)
					return
				}
				/*Validando que la unidad medica exista*/
				err, unidadMed := ModeloUnidadM.ConsultaUnidadMed(row[2])
				if err != nil {
					Estructuras.Responder(http.StatusInternalServerError,
						"No existe esta unidad medica", err.Error(), c)
					return
				}
				delegaciones := FiltraDelegaciones(rows, 2)
				fmt.Println("Delegaciones****", &delegaciones)
				/*Agregando delegaciones a contratos*/
				var delegacion ContratoModelo.Contrato_Delegaciones
				for _, deleg := range delegaciones {
					duplicado := ContratoModelo.RegresaDelegContrato(deleg.NumContrato, deleg.DelegacionId)
					if duplicado == nil {
						delegacion.ContratoNumeroContrato = deleg.NumContrato
						delegacion.DelegacionesClvDele = deleg.DelegacionId
						err = delegacion.AgregarDelegContrato(tx)
						if err != nil {
							tx.Rollback()
							Estructuras.Responder(http.StatusInternalServerError, "Error al agregar delegaciones contrato", err.Error(), c)
							return
						}
					}
				}
				ConcepD.Id = bson.NewObjectId().Hex()
				ConcepD.Cantidad = cantidad
				ConcepD.ConceptoContratoId = equipo.Id
				ConcepD.UnidadMedClvPresupuestal = unidadMed.ClvPresupuestal
				estado, _ := DistribucionModelo.ValidaRepetidos(ConcepD.UnidadMedClvPresupuestal, ConcepD.ConceptoContratoId)
				if estado {
					Estructuras.Responder(http.StatusInternalServerError,
						"Ya existe una distribucion para esta unidad y esta clave", err.Error(), c)
					return
				}
				err = ConcepD.AgregarDistribucion(tx)
				if err != nil {
					tx.Rollback()
					Estructuras.Responder(http.StatusInternalServerError, "Erros al registrar la distribucion", err.Error(), c)
					return
				}
				cantidadActu := equipo.CantidadDistribuida + cantidad
				cantidadConcepto := equipo.CantidadConcepto
				respuesta := ConceptoContratoModelo.ActualizarDistribucion(equipo.Id, cantidadActu, cantidadConcepto)
				if respuesta != nil {
					fmt.Println("Error-----", respuesta.Error())
					Estructuras.Responder(http.StatusInternalServerError, "Erros al registrar la distribucion", err.Error(), c)
					return
				}
			}
		}
		//********************************************************************************************************
		/* *********************   finalizar insercion y hacer el commit     ********************* */
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
	if distribucion == "distribucion_sin_concepto" {
		var conceptos ConceptoContratoModelo.ConceptoContrato
		respuesta, err, validado := ValidarDistribucionSinConceptos(f)
		if err != nil {
			Estructuras.Responder(http.StatusInternalServerError, err.Error(), respuesta, c)
			return
		}
		if validado == false {
			Estructuras.Responder(http.StatusInternalServerError, "Errores de Validacion", respuesta, c)
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
		defer fi.Close()
		ar, err := excelize.OpenReader(fi)
		if err != nil {
			Estructuras.Responder(http.StatusInternalServerError, "Error al abrir el archivo", err.Error(), c)
			return
		}
		rows, err := ar.GetRows("DistribucionConceptos")
		if err != nil {
			fmt.Println(err.Error())
			Estructuras.Responder(http.StatusInternalServerError,
				"Error al leer el archivo o no es el archivo original", err.Error(), c)
			return
		}
		resultado := FiltraArraypreiCantidad(rows, 1)
		fmt.Println("Conceptos****", &resultado)
		/*Agregando conceptos*/
		for _, concep := range resultado {
			duplicado, _ := ConceptoContratoModelo.VerificaEquipo(concep.IdPrei, concep.NumContrato)
			if duplicado == false {
				_, clave := CatalogosModel.ConsultaUnaClave("", concep.IdPrei)
				conceptos.Id = bson.NewObjectId().Hex()
				conceptos.PrecioUniSnIva = float64(concep.Precio)
				conceptos.CantidadConcepto = 0
				conceptos.CantidadAmpliada = 0
				conceptos.CantidadDistribuida = 0
				conceptos.Marca = strings.ToUpper(concep.Marca)
				conceptos.Modelo = strings.ToUpper(concep.Modelo)
				conceptos.ObjetoContratacion = clave.CuadroBasico.Grupo
				conceptos.FechaMaxEntrega = concep.Fecha
				conceptos.GarantiaBienes = concep.Garantia
				conceptos.PreiIdArticulo = clave.IdArticulo
				conceptos.ContratoNumeroContrato = concep.NumContrato
				fmt.Println(conceptos)
				err := conceptos.AgregarConceptoContrato(tx)
				if err != nil {
					tx.Rollback()
					Estructuras.Responder(http.StatusInternalServerError, "Erros al inserta el concepto", err.Error(), c)
					return
				}
			}
		}
		delegaciones := FiltraDelegaciones(rows, 7)
		fmt.Println("Delegaciones****", &delegaciones)
		/*Agregando delegaciones a contratos*/
		var delegacion ContratoModelo.Contrato_Delegaciones
		for _, deleg := range delegaciones {
			duplicado := ContratoModelo.RegresaDelegContrato(deleg.NumContrato, deleg.DelegacionId)
			if duplicado == nil {
				delegacion.ContratoNumeroContrato = deleg.NumContrato
				delegacion.DelegacionesClvDele = deleg.DelegacionId
				err = delegacion.AgregarDelegContrato(tx)
				if err != nil {
					tx.Rollback()
					Estructuras.Responder(http.StatusInternalServerError, "Error al agregar delegaciones contrato", err.Error(), c)
					return
				}
			}
		}
		err = tx.Commit().Error
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "Error al realizar el commit conceptos y delegaciones", err.Error(), c)
			return
		}
		txdis := Conexion.GetDB().Begin()
		if txdis.Error != nil {
			Estructuras.Responder(http.StatusInternalServerError, "Error al iniciar la transaccion", tx.Error, c)
			return
		}
		for num, row := range rows {
			if num > 0 {
				cantidad, _ := strconv.Atoi(row[8])
				/*Validando que el id prei (equipo) ya se encuentre registrado en el mismo oontrato*/
				duplicado, equipo := ConceptoContratoModelo.VerificaEquipo(row[1], row[0])
				if duplicado == false {
					Estructuras.Responder(http.StatusInternalServerError,
						"Este Id Prei no ha sido registrado en el contrato", err.Error(), c)
					return
				}
				/*Validando que la unidad medica exista*/
				err, unidadMed := ModeloUnidadM.ConsultaUnidadMed(row[7])
				if err != nil {
					Estructuras.Responder(http.StatusInternalServerError,
						"No existe esta unidad medica", err.Error(), c)
					return
				}
				ConcepD.Id = bson.NewObjectId().Hex()
				ConcepD.Cantidad = cantidad
				ConcepD.ConceptoContratoId = equipo.Id
				ConcepD.UnidadMedClvPresupuestal = unidadMed.ClvPresupuestal
				estado, _ := DistribucionModelo.ValidaRepetidos(ConcepD.UnidadMedClvPresupuestal, ConcepD.ConceptoContratoId)
				if estado {
					txdis.Rollback()
					Estructuras.Responder(http.StatusInternalServerError,
						"Ya existe una distribucion para esta unidad y esta clave", nil, c)
					return
				}
				err = ConcepD.AgregarDistribucion(txdis)
				if err != nil {
					fmt.Println("Error al agregar distribucion***", err.Error())
					txdis.Rollback()
					Estructuras.Responder(http.StatusInternalServerError, "Erros al registrar la distribucion", err.Error(), c)
					return
				}
				cantidadActu := equipo.CantidadDistribuida + cantidad
				cantidadConcepto := equipo.CantidadConcepto + cantidad
				fmt.Println(cantidadActu, "-*-*-*-*-*-*-*-*-*-*-*-*-*")
				respuesta := ConceptoContratoModelo.ActualizarDistribucion(equipo.Id, cantidadActu, cantidadConcepto)
				if respuesta != nil {
					fmt.Println("Error-----", respuesta.Error())
					Estructuras.Responder(http.StatusInternalServerError, "Erros al registrar la distribucion",
						respuesta.Error(), c)
					return
				}
			}
		}
		//********************************************************************************************************
		/* *********************   finalizar insercion y hacer el commit     ********************* */
		defer fi.Close()
		err = txdis.Commit().Error
		if err != nil {
			txdis.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "Error al realizar el commit de distribucion", err.Error(), c)
			return
		}
		Estructuras.Responder(http.StatusOK, "Datos agregados correctamente", nil, c)
		return
	}
}

//********************************************************************************************************
/* *********************    Agregar distribucion de manera manual desde el front-end     ********************* */
func AgregaDistribucion(c *gin.Context) {
	var distribucion DistribucionModelo.DistribucionUnidadMed
	var delegacion ContratoModelo.Contrato_Delegaciones
	n_contrato := c.DefaultQuery("NumContrato", "")
	clv_dele := c.DefaultQuery("ClvDele", "")
	cantidad := c.DefaultQuery("CantidadConcepto", "")
	err := c.BindJSON(&distribucion)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error al castear la estructura", err.Error(), c)
		return
	}
	cantidadCon, _ := strconv.Atoi(cantidad)
	/*Consulta la cantidad ya distribuida anteriormente*/
	_, distribuido := DistribucionModelo.VerificarCantidades(distribucion.ConceptoContratoId)
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al iniciar la transaccion", tx.Error, c)
		return
	}
	/*Verifica que no se repita la distribucion por concepto y unidad medica*/
	estado, _ := DistribucionModelo.ValidaRepetidos(distribucion.UnidadMedClvPresupuestal, distribucion.ConceptoContratoId)
	if estado {
		Estructuras.Responder(http.StatusInternalServerError,
			"Ya existe una distribucion para esta unidad y esta clave", err.Error(), c)
		return
	}
	/*Consulta que la delegacion y el numero de contrato, si ya existe no la agrega, si no agrega la delegacion*/
	duplicado := ContratoModelo.RegresaDelegContrato(n_contrato, clv_dele)
	if duplicado == nil {
		delegacion.ContratoNumeroContrato = n_contrato
		delegacion.DelegacionesClvDele = clv_dele
		err = delegacion.AgregarDelegContrato(tx)
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "Error al agregar delegaciones contrato", err.Error(), c)
			return
		}
	}
	/*Agrega la distribucion*/
	distribucion.Id = bson.NewObjectId().Hex()
	err = distribucion.AgregarDistribucion(tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Erros al registrar la distribucion", err.Error(), c)
		return
	}
	/*Suma la cantidad ya distribuida con la cantidad nueva por distribuir*/
	cantidadActu := distribuido + distribucion.Cantidad
	/*Actualiza el campo cantidad distribuida de la tabla concepto*/
	respuesta := ConceptoContratoModelo.ActualizarDistribucion(distribucion.ConceptoContratoId, cantidadActu, cantidadCon)
	if respuesta != nil {
		fmt.Println("Error-----", respuesta.Error())
		Estructuras.Responder(http.StatusInternalServerError, "Erros al registrar la distribucion", respuesta.Error(), c)
		return
	}
	/*Se realiza el commit a la base de datos*/
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al realizar el commit", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos agregados correctamente", distribucion, c)
	return
}

//********************************************************************************************************
/* *********************    Controlador para editar una distribucion     ********************* */
func EditarDistribucion(c *gin.Context) {
	var distribucion DistribucionModelo.DistribucionUnidadMed
	var delegacion ContratoModelo.Contrato_Delegaciones
	n_contrato := c.DefaultQuery("NumContrato", "")
	clv_dele := c.DefaultQuery("ClvDele", "")
	err := c.BindJSON(&distribucion)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error al castear la estructura", err.Error(), c)
		return
	}
	/*Consulta la cantidad ya distribuida anteriormente*/
	_, distribuido := DistribucionModelo.VerificarCantidades(distribucion.ConceptoContratoId)
	/*Consultar la distribucion*/
	err, resultado := DistribucionModelo.ConsultarDistribucion(distribucion.Id)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al buscar la distribucion", err.Error(), c)
		return
	}
	/*Inicio mi transaccion en mi  base de datos*/
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al iniciar la transaccion", tx.Error, c)
		return
	}
	if resultado.UnidadMed.DelegacionesClvDele != clv_dele {
		/*Buscar cuantas distribuciones tienen la misma delegacion en el mismo contrato*/
		err, dis := DistribucionModelo.BuscarDistribucion(n_contrato, resultado.UnidadMed.DelegacionesClvDele)
		if err != nil {
			Estructuras.Responder(http.StatusInternalServerError, "Error al buscar la distribucion", err.Error(), c)
			return
		}
		if len(*dis) <= 1 {
			err, _ = ContratoModelo.EliminarDeleContr(n_contrato, resultado.UnidadMed.DelegacionesClvDele, tx)
			if err != nil {
				Estructuras.Responder(http.StatusInternalServerError, "Error al eliminar la delegacion del contrato",
					err.Error(), c)
				return
			}
		}
	}
	/*Verifica que no se repita la distribucion por concepto y unidad medica*/
	estado, res := DistribucionModelo.ValidaRepetidos(distribucion.UnidadMedClvPresupuestal,
		distribucion.ConceptoContratoId)
	if estado && res.Id != distribucion.Id {
		Estructuras.Responder(http.StatusInternalServerError,
			"Ya existe una distribucion para esta unidad y esta clave", nil, c)
		return
	}
	/*Consulta que la delegacion y el numero de contrato, si ya existe no la agrega, si no agrega la delegacion*/
	duplicado := ContratoModelo.RegresaDelegContrato(n_contrato, clv_dele)
	if duplicado == nil {
		delegacion.ContratoNumeroContrato = n_contrato
		delegacion.DelegacionesClvDele = clv_dele
		err = delegacion.AgregarDelegContrato(tx)
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "Error al agregar delegaciones contrato", err.Error(), c)
			return
		}
	}
	/*Edita la distribucion*/
	err = distribucion.ActualizarDistribucion(distribucion.Id, tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al actualizar la distribucion", err.Error(), c)
		return
	}
	/*Suma la cantidad ya distribuida con la cantidad nueva por distribuir*/
	cantidadActu := distribuido + distribucion.Cantidad - resultado.Cantidad
	/*Actualiza el campo cantidad distribuida de la tabla concepto*/
	err = ConceptoContratoModelo.ActualizarDistribucion(distribucion.ConceptoContratoId, cantidadActu,
		resultado.ConceptoContrato.CantidadConcepto)
	if err != nil {
		fmt.Println("Error-----", err.Error())
		Estructuras.Responder(http.StatusInternalServerError, "Erros al registrar la distribucion", err.Error(), c)
		return
	}
	/*Se realiza el commit a la base de datos*/
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al realizar el commit", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos agregados correctamente", distribucion, c)
	return
}

//********************************************************************************************************
/* *********************    Controlador para eliminar una distribucion     ********************* */
func EliminaDistribucion(c *gin.Context) {
	id_distribucion := c.DefaultQuery("id", "")
	num_contrato := c.DefaultQuery("n_contrato", "")
	clv_dele := c.DefaultQuery("clv", "")
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al iniciar la transaccion", tx.Error, c)
		return
	}
	/*Buscar cuantas distribuciones tienen la misma delegacion en el mismo contrato*/
	err, dis := DistribucionModelo.BuscarDistribucion(num_contrato, clv_dele)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al buscar la distribucion", err.Error(), c)
		return
	}
	if len(*dis) <= 1 {
		err, _ = ContratoModelo.EliminarDeleContr(num_contrato, clv_dele, tx)
		if err != nil {
			Estructuras.Responder(http.StatusInternalServerError, "Error al eliminar la delegacion del contrato",
				err.Error(), c)
			return
		}
	}
	/*Consultar la distribucion*/
	estado, distribucion := DistribucionModelo.ConsultarDistribucion(id_distribucion)
	if estado != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al buscar la distribucion", estado.Error(), c)
		return
	}
	_, distribuido := DistribucionModelo.VerificarCantidades(distribucion.ConceptoContratoId)

	/*Actualiza el campo cantidad distribuida de la tabla concepto*/
	cantidadActual := distribuido - distribucion.Cantidad
	fmt.Println("Esta es la nueva cantidad distribuida del concepto...." , cantidadActual)
	respuesta := ConceptoContratoModelo.ActualizarDistribucion(distribucion.ConceptoContratoId, cantidadActual,
		distribucion.ConceptoContrato.CantidadConcepto)
	if respuesta != nil {
		fmt.Println("Error-----", respuesta.Error())
		Estructuras.Responder(http.StatusInternalServerError, "Erros al actualizar la cantidad", respuesta.Error(), c)
		return
	}
	err, _ = DistribucionModelo.EliminarDistribucion(id_distribucion, tx)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al eliminar la distribucion", err.Error(), c)
		return
	}
	/*Se realiza el commit a la base de datos*/
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al realizar el commit", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Distribucion eliminada correctamente", distribucion, c)
	return
}
