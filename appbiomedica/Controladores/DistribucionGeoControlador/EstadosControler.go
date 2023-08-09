package DistribucionGeoControlador

import (
	"appbiomedica/Modelos/ModeloDistribucionGeografica"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"net/http"
)

//*****************************************************************************************************************
//consulta todos los Estados
func GetEstados(c *gin.Context) {
	municipios := c.Query("municipios")
	if municipios == "true" {
		err, resultado := ModeloDistribucionGeografica.ConsultaEstadosMunicipios()

		if err != nil {
			Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
			return
		}
		Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
		return
	}
	err, resultado := ModeloDistribucionGeografica.ConsultaEstados()

	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return

}
