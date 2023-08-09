package CatalogoControlador

import (
	"appbiomedica/Modelos/CatalogosModel"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

/* *********************    controlador para consultar todos el catalog proforma     ********************* */
func GetProforma(c *gin.Context) {
	servicioId := c.DefaultQuery("servicio", "0")
	nombre := c.DefaultQuery("nombre", "")
	num_pagina := c.DefaultQuery("pagina", "1")
	num_registro := c.DefaultQuery("num_registros", "25")
	nombre = strings.Trim(nombre, " ")
	nombres := []string{}
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
	err, resultado, total := CatalogosModel.ConsultaProforma(servicioId, nombres, num_pag, num_reg)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente",
		gin.H{"total": total, "registros": resultado}, c)
	return
}

/* *********************    controlador para consultar una sola clave proforma    ********************* */
func GetClaveProforma(c *gin.Context) {
	id := c.Param("id")
	err, resultado := CatalogosModel.ConsultaClaveProforma(id)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}