package PenasControlador

import (
	"appbiomedica/Modelos/ContratoModel/PenasModelo"
	"appbiomedica/Modulos/Estructuras"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//********************************************************************************************************
/* *********************    Controlador para consultar penas deductivas     ********************* */
func GetPenasConvencionales(c *gin.Context) {
	numero_contrato := c.DefaultQuery("numero_contrato", "")
	estado := false
	err, datos := PenasModelo.ConsultarPenaConvencional(numero_contrato)
	if datos != nil {
		estado = true
	}
	if err != nil {
		fmt.Println("Resultado convencional", datos)
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos",
			gin.H{"datos": err.Error(), "estado": estado}, c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", gin.H{"datos": datos,
		"estado": estado}, c)
	return
}
