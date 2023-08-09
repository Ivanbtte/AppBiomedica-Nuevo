package FiltrosControlador

import (
	"appbiomedica/Modelos/FiltrosModel"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//*****************************************************************************************************************
//consulta todos los Grupos-temas
func GetGruposTemas(c *gin.Context) {
	todos := c.DefaultQuery("all", "true")
	tema := c.DefaultQuery("tema", "")
	temas := [] string{}
	if tema != "" {
		temas = strings.Split(tema, ",")
	}
	err, resultado := FiltrosModel.ConsultaGrupoTemas(todos, temas)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}
