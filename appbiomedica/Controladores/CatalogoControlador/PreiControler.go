package CatalogoControlador

import (
	"appbiomedica/Modelos/CatalogosModel"
	"appbiomedica/Modulos/Estructuras"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

//*****************************************************************************************************************
//consulta todos los equipos prei
func GetPrei(c *gin.Context) {
	cuadro_basico := c.DefaultQuery("cuadro_id", "0")
	nombre := c.DefaultQuery("nombre", "")
	grupo := c.DefaultQuery("grupo", "")
	generico := c.DefaultQuery("generico", "")
	especifico := c.DefaultQuery("especifico", "")
	num_pagina := c.DefaultQuery("pagina", "1")
	num_registro := c.DefaultQuery("num_registros", "25")
	nombre = strings.Trim(nombre, " ")
	nombres := []string{}
	Id_grupo, err := strconv.Atoi(cuadro_basico)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Id del grupo incorrecto", err.Error(), c)
		return
	}
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
	if nombre != "" {
		nombres = strings.Split(nombre, " ")
	}
	err, resultado, total := CatalogosModel.ConsultaPrei(Id_grupo, nombres, grupo, generico, especifico, num_pag, num_reg)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", gin.H{"total": total, "registros": resultado}, c)
	return
}

//*****************************************************************************************************************
//consulta  una calve
func GetUnaClave(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	id_articulo := c.DefaultQuery("id_articulo", "")
	if id_articulo == "" && id == "" {
		Estructuras.Responder(http.StatusInternalServerError,
			"Debes ingresar un parametro de busqueda", errors.New("Falta un dato"), c)
		return
	}
	err, resultado := CatalogosModel.ConsultaUnaClave(id, id_articulo)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}

/* ********************* Controlador para regresar grupos, genericos, especifico PREI    ********************* */

func GetGruposFiltros(c *gin.Context) {
	grupoId := c.DefaultQuery("grupo", "")
	generico := c.DefaultQuery("generico", "")
	grupo := c.DefaultQuery("grupos", "false")
	grupos, err := strconv.ParseBool(grupo)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al parsear booleano", err.Error(), c)
		return
	}
	err, resultado := CatalogosModel.ConsultarGruposPrei(grupoId, generico, grupos)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return

}
