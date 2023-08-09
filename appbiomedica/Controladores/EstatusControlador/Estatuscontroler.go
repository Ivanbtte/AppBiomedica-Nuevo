package EstatusControlador

import (
	"appbiomedica/Modelos/EstatusModel"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"net/http"
)

//*****************************************************************************************************************
/* *********************    Controlador para regresar la fecha limite actual     ********************* */
func GetEstatus(c *gin.Context) {
	err, resultado := EstatusModel.ConsultaEstatus()
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", resultado, c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}