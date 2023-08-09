package CatalogoControlador

import (
	"appbiomedica/Modelos/CatalogosModel"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"net/http"
)

/* *********************    controlador para consultar todos los servicios proforma     ********************* */
func GetServiciosProforma(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	if id != "" {
		err, resultado := CatalogosModel.ConsultarServicio(id)
		if err != nil {
			Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
			return
		}
		Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
		return
	} else {
		err, resultado := CatalogosModel.ConsultaServProforma()
		if err != nil {
			Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
			return
		}
		Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
		return
	}

}

/* *********************    controlador para consultar los nombres de los servicios de una clave    ********************* */
func GetClaveNombreServ(c *gin.Context) {
	id := c.Param("id")
	err, resultado := CatalogosModel.ConsultaNombresServicio(id)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}
