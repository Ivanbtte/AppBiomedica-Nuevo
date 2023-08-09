package ConceptosServiciosControlador

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/ContratoModel/ConceptosServiciosModelo"
	"appbiomedica/Modelos/ContratoModel/ContratoModelo"
	"appbiomedica/Modulos/Estructuras"
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
)

//********************************************************************************************************
/* *********************    Controlador para Agregar un nuevo concepto o servicio     ********************* */
func PostConceptoServicioContrato(c *gin.Context) {
	var estructura ConceptosServiciosModelo.Concepto
	/*****  Query params  *****/
	tipo_contrato := c.DefaultQuery("tipo_contrato", "")
	sub_tipo_contrato := c.DefaultQuery("sub_tipo_contrato", "")
	file, _ := c.FormFile("file")
	fmt.Println(file.Header, "***", file.Filename)
	//---*
	f, err := file.Open()
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al abrir el archivo", err, c)
		return
	}
	fmt.Println(f)
	/*****  Decidir la estructura  *****/
	if sub_tipo_contrato != "" {
		estructura = Factory(sub_tipo_contrato)
		if estructura == nil {
			Estructuras.Responder(http.StatusInternalServerError, "Error al seleccionar la estructura", nil, c)
			return
		}
	} else {
		estructura = Factory(tipo_contrato)
		if estructura == nil {
			Estructuras.Responder(http.StatusInternalServerError, "Error al seleccionar la estructura", nil, c)
			return
		}
	}
	resultado, errores, err := estructura.ValidaConceptosArchivo(f)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, err.Error(), errores, c)
		return
	}
	if resultado == false {
		Estructuras.Responder(http.StatusInternalServerError, "Errores de Validacion", errores, c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Archivo Validado Correctamente", nil, c)
	return
}

/*****  Controlador para agregar conceptos desde un archivo validado  *****/
func AgregarConceptosArchivo(c *gin.Context) {
	var estructura ConceptosServiciosModelo.Concepto
	var concepto_contrato ConceptosServiciosModelo.Conceptos
	var sub_tipo_cont ContratoModelo.SubTipoContrato
	var id_concepto string
	/*****  Query params  *****/
	tipo_contrato := c.DefaultQuery("tipo_contrato", "")
	sub_tipo_contrato := c.DefaultQuery("sub_tipo_contrato", "")
	file, _ := c.FormFile("file")
	//---*
	f, err := file.Open()
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al abrir el archivo", err, c)
		return
	}
	/*****  Decidir la estructura  *****/
	if sub_tipo_contrato != "" {
		estructura = Factory(sub_tipo_contrato)
		if estructura == nil {
			Estructuras.Responder(http.StatusInternalServerError, "Error al seleccionar la estructura", nil, c)
			return
		}
	} else {
		estructura = Factory(tipo_contrato)
		if estructura == nil {
			Estructuras.Responder(http.StatusInternalServerError, "Error al seleccionar la estructura", nil, c)
			return
		}
	}

	/*****  Abrir el archivo de excel  *****/
	ar, err := excelize.OpenReader(f)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al abrir el archivo", err.Error(), c)
		return
	}
	rows, err := ar.GetRows("Conceptos")
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al abrir el archivo",
			errors.New("no se pudo leer el archivo o no tiene el formato correcto"), c)
		return
	}
	/* *********************    inicio mi transaccion de mi base de datos     ********************* */
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al iniciar la transaccion", tx.Error, c)
		return
	}
	fmt.Println(len(rows))
	for num, fila := range rows {
		fmt.Println("Entreeeeee", len(fila), num)
		if num > 0 {
			conceptoResul, estado := ConceptosServiciosModelo.BuscarConceptoContrato(Estructuras.ConvertLowercase(fila[0]))
			if !estado {
				/*****  Cambiar esto despues y ponerlo en un metodo separarado   *****/
				concepto_contrato.Id = bson.NewObjectId().Hex()
				concepto_contrato.ContratoNumeroContrato = Estructuras.FormatoTextoMayusculas(fila[0])
				err := concepto_contrato.AgregarConceptoContrato()
				if err != nil {
					tx.Rollback()
					Estructuras.Responder(http.StatusInternalServerError, "Error al insertar el concepto contrato",
						err.Error(), c)
					return
				}
				id_concepto = concepto_contrato.Id
			} else {
				id_concepto = conceptoResul.Id
			}
			fmt.Println("Agrego--------->>>>>>><", id_concepto)
			err := estructura.DefinirEstructura(fila, id_concepto).AgregarConcepto(tx)
			if err != nil {
				tx.Rollback()
				Estructuras.Responder(http.StatusInternalServerError, "Error al insertar el concepto", err.Error(), c)
				return
			}
		}
	}
	if sub_tipo_contrato != "" {
		resultado := ConceptosServiciosModelo.FiltraArray(rows, 0)
		for _, num_contrato := range resultado {
			sub_tipo_cont.ContratoNumeroContrato = Estructuras.FormatoTextoMayusculas(num_contrato)
			sub_tipo_cont.SubTipoId = sub_tipo_contrato
			sub_tipo_cont.Estado = true
			err := sub_tipo_cont.AgregarSubTipoContrato(tx)
			if err != nil {
				tx.Rollback()
				Estructuras.Responder(http.StatusInternalServerError, "Error al insertar el subtipo", err.Error(), c)
				return
			}
		}
	}
	//********************************************************************************************************
	/* *********************   finalizar insercion y hacer el commit     ********************* */
	defer f.Close()
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al realizar el commit", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos agregados correctamente", nil, c)
	return
}

/*****  Controlador para consultar conceptos de un contrato  *****/
func GetConceptos(c *gin.Context) {
	var estructura ConceptosServiciosModelo.Concepto

	/*****  Query params  *****/
	tipo_contrato := c.DefaultQuery("tipo_contrato", "")
	sub_tipo_contrato := c.DefaultQuery("sub_tipo_contrato", "")
	num_contrato := c.DefaultQuery("num_contrato", "")
	num_pagina := c.DefaultQuery("pagina", "1")
	num_registro := c.DefaultQuery("num_registros", "15")
	/*****    *****/
	num_pag, err := strconv.Atoi(num_pagina)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Numero de pagina incorrecto", err.Error(), c)
		return
	}
	num_reg, err := strconv.Atoi(num_registro)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Numero de registros incorrecto", err.Error(), c)
		return
	}
	/*****  Decidir la estructura  *****/
	if sub_tipo_contrato != "" {
		fmt.Println("soy un subtipo")
		estructura = Factory(sub_tipo_contrato)
		if estructura == nil {
			Estructuras.Responder(http.StatusInternalServerError, "Error al seleccionar la estructura", nil, c)
			return
		}
	} else {
		fmt.Println("Soy un tipo")
		estructura = Factory(tipo_contrato)
		if estructura == nil {
			Estructuras.Responder(http.StatusInternalServerError, "Error al seleccionar la estructura", nil, c)
			return
		}
	}
	/*****  consulto mis conceptos  *****/
	err, resultado, suma, total := estructura.ConsultarConceptos(num_contrato, num_pag, num_reg)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente",
		gin.H{"total": suma, "registros": resultado, "numReg": total}, c)
	return
}
