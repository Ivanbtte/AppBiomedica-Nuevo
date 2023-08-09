package UnidadesMedicasControlador

import (
	"appbiomedica/Modelos/ModeloUnidadM"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
var TypeDele = 0
func GetDelegations(c *gin.Context) {
	typeDelegations := c.DefaultQuery("type", "")
	if typeDelegations != "" {
		TypeDele, _ = strconv.Atoi(typeDelegations)
	}
	err, resultado := ModeloUnidadM.ConsultsDelegations(TypeDele)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}
