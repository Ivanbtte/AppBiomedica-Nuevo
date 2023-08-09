package FiltrosControlador

import (
	"appbiomedica/Modelos/FiltrosModel"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"net/http"
)

//*****************************************************************************************************************
//consulta todos los temas
func GetTemas(c *gin.Context) {
	err, resultado := FiltrosModel.ConsultaTemas()
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}