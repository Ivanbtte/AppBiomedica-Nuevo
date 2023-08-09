package ServiciosUnidadesControlador

import (
	"appbiomedica/Modelos/ServiciosUnidadesModel"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//*****************************************************************************************************************
/* *********************    Funcion que valida la peticion para regresar los servicios     ********************* */
func GetServiciosUnidades(c *gin.Context) {
	id:= c.DefaultQuery("id","")
	ids := [] string{}
	if id != "" {
		ids = strings.Split(id, ",")
	}
	err, resultado := ServiciosUnidadesModel.ConsultaServiciosUnidades(ids)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}