package DistribucionGeoControlador

import (
	"appbiomedica/Modelos/ModeloDistribucionGeografica"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"net/http"
)

//*****************************************************************************************************************
//consulta todos los roles
func GetMunicipios(c *gin.Context) {
	err, resultado := ModeloDistribucionGeografica.ConsultaMunicipios()
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}