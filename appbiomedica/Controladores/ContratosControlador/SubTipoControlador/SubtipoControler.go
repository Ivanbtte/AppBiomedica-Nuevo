package SubTipoControlador

import (
	"appbiomedica/Modelos/ContratoModel/PermisosContratoModelo"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*****  Controlador para consultar los subtipos de contratos  *****/
func GetSubTipo(c *gin.Context) {
	tipo_contrato := c.DefaultQuery("tipo_contrato", "")
	resultado, err := PermisosContratoModelo.ConsultarSubTipo(tipo_contrato)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}
