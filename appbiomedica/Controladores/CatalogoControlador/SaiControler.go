package CatalogoControlador

import (
	"appbiomedica/Modelos/CatalogosModel"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

//*****************************************************************************************************************
/* *********************    Funcion para regresar genericos y especificos      ********************* */
func GetGenericosEspecificos(c *gin.Context) {
	grupoId := c.DefaultQuery("grupo", "")
	generico := c.DefaultQuery("generico", "")
	todosGrupos := c.DefaultQuery("todos_genericos", "false")
	todosEspecificos := c.DefaultQuery("todos_especificos", "false")
	Grupos, err := strconv.ParseBool(todosGrupos)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al parsear booleano", err.Error(), c)
		return
	}
	Especificos, err := strconv.ParseBool(todosEspecificos)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al parsear booleano", err.Error(), c)
		return
	}
	grupo := [] string{}
	if grupoId != "" {
		grupo = strings.Split(grupoId, ",")
	}
	err, resultado := CatalogosModel.ConsultarGenericosEspecifico(grupo, generico, Grupos, Especificos)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return

}

//*****************************************************************************************************************
//consulta todos los equipos sai
func GetSai(c *gin.Context) {
	grupoId := c.DefaultQuery("grupo", "0")
	generico := c.DefaultQuery("generico", "")
	especifico := c.DefaultQuery("especifico", "")
	temaId := c.DefaultQuery("tema", "")
	nombre := c.DefaultQuery("nombre", "")
	num_pagina := c.DefaultQuery("pagina", "1")
	num_registro := c.DefaultQuery("num_registros", "25")
	cuadro := c.DefaultQuery("cuadro", "")
	nombre = strings.Trim(nombre, " ")
	nombres := []string{}
	tema := [] string{}
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
	if temaId != "" {
		tema = strings.Split(temaId, ",")
	}
	if nombre != "" {
		nombres = strings.Split(nombre, " ")
	}
	err, resultado, total := CatalogosModel.ConsultaSai(grupoId, tema, nombres, num_pag, num_reg, cuadro, generico, especifico)

	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", gin.H{"total": total, "registros": resultado}, c)
	return

}

//*****************************************************************************************************************
//consulta  una calve sai
func GetClaveSai(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	clave := c.DefaultQuery("clave", "0")
	grupo := ""
	generico := ""
	especifico := ""
	diferenciador := ""
	variable := ""
	if clave != "0" {
		grupo = clave[0:3]
		generico = clave[3:6]
		especifico = clave[6:10]
		diferenciador = clave[10:12]
		variable = clave[12:14]
	}
	err, resultado := CatalogosModel.ConsultaClaveSai(id, grupo, generico, especifico, diferenciador, variable)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}
