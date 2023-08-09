package BitacorasControlador

import (
	"appbiomedica/Modelos/BitacoraModel"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"net/http"
)

//*****************************************************************************************************************
/* *********************    Controlador para regrsar el registro de una solicitud     ********************* */
func GetBitacoraSolicitud(c *gin.Context) {
	idSolicitud := c.DefaultQuery("id_solicitud", "0")
	err, resultado := BitacoraModel.ConsultaBitacoraSolicitud(idSolicitud)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}
