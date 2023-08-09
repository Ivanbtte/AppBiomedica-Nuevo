package ContratoDelegacionControlador

import (
	"appbiomedica/Modelos/ContratoModel/ContratoModelo"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"net/http"
)

//********************************************************************************************************
/* *********************   Funcion para consultar todas las delegaciones por numero de contrato
********************* */
func GetContratoDeleg(c *gin.Context) {
	num_contrato := c.DefaultQuery("num_contrato", "")
	if num_contrato == "" {
		Estructuras.Responder(http.StatusInternalServerError, "Debes ingrear el numero de contrato ", nil, c)
		return
	}
	err, resultado := ContratoModelo.ConsultaDelegContrato(num_contrato)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}
