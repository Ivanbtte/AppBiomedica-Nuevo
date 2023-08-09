package AdministradorControlador

import (
	"appbiomedica/Modelos/ContratoModel/AdminAuxModelo"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"net/http"
)

//********************************************************************************************************
/* *********************    Controlador para consultar administradores o auxiliares     ********************* */
func GetAdministradores(c *gin.Context) {
	numero_contrato := c.DefaultQuery("numero_contrato", "")
	estado := false
	err, datos := AdminAuxModelo.ConsultarAdministrador(numero_contrato)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	if len(*datos) > 0 {
		estado = true
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", gin.H{"datos": datos,
		"estado": estado}, c)
	return
}