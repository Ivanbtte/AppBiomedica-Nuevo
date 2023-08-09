package PenasControlador

import (
	"appbiomedica/Modelos/ContratoModel/PenasModelo"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"net/http"
)

//********************************************************************************************************
/* *********************    Controlador para consultar penas deductivas     ********************* */
func GetPenasDeductivas(c *gin.Context) {
	numero_contrato := c.DefaultQuery("numero_contrato", "")
	err, datos, estado := PenasModelo.ConsultarPenaDeductiva(numero_contrato)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	if len(*datos) > 0 {
		estado = true
	} else {
		estado = false
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", gin.H{"datos": datos,
		"estado": estado}, c)
	return
}
