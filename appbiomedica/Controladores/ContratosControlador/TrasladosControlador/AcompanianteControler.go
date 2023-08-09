package TrasladosControlador

import (
	"appbiomedica/Modelos/ServiciosSubrogados/TrasladoPacientesModel"
	"appbiomedica/Modulos/Estructuras"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAcompaniantePorUnidad(ctx *gin.Context) {
	unidadMedId := ctx.DefaultQuery("unidadMedId", "")
	unidad, err := strconv.Atoi(unidadMedId)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "error en el numero de la unidad medica", nil, ctx)
		return
	}
	result, err := TrasladoPacientesModel.BuscarAcompaniantePorUnidad(unidad)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "error al buscar el ultimo registro del acompañante", nil, ctx)
		return
	}
	Estructuras.Responder(http.StatusAccepted, "Datos consultados", result, ctx)
}
