package SubTipoControlador

import (
	"appbiomedica/Modelos/ContratoModel/ContratoModelo"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*****  Controlador para consultar que subtipo tiene asignado un contrato  *****/
func GetSubTipoContrato(c *gin.Context) {
	numero_contrato := c.DefaultQuery("numero_contrato", "")
	resultado, err := ContratoModelo.ConsultarSubTipoContrato(numero_contrato)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}
