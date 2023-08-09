package DepartamentoControlador

import (
	"appbiomedica/Modelos/DepartamentosModel"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDepartamento(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	err, resultado := DepartamentosModel.ConsultaDepartamento(id)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}
