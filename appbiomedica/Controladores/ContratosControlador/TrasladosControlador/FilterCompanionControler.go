package TrasladosControlador

import (
	"appbiomedica/Modelos/ServiciosSubrogados/TrasladoPacientesModel"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetCompanionFilter(ctx *gin.Context) {
	unitId := ctx.DefaultQuery("unitId", "")
	state := ctx.DefaultQuery("state", "")
	destination := ctx.DefaultQuery("destination", "")
	origin := ctx.DefaultQuery("origin", "")
	numberContract := ctx.DefaultQuery("numberContract", "")
	dateStart := ctx.DefaultQuery("startDate", "")
	dateEnd := ctx.DefaultQuery("endDate", "")
	startDate := ""
	endDate := ""
	if dateStart != "" && dateEnd != "" {
		temporalDate, _ := time.Parse(time.RFC3339, dateStart)
		startDate = temporalDate.Truncate(24 * time.Hour).Format("2006-01-02 15:04:05")
		temporalDate2, _ := time.Parse(time.RFC3339, dateEnd)
		endDate = temporalDate2.Add(24 * time.Hour).Truncate(24 * time.Hour).Format("2006-01-02 15:04:05")
	}
	stateRequest, err := strconv.ParseBool(state)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error estado no valido", err.Error(), ctx)
		return
	}
	idUnit, err := strconv.Atoi(unitId)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Numero de Unidad Medica incorrecto", err.Error(), ctx)
		return
	}
	var contractsNumbers []string
	if numberContract != "" {
		contractsNumbers = strings.Split(numberContract, ",")
	}
	results, err := TrasladoPacientesModel.SearchCompanionFilter(idUnit, stateRequest, origin, destination,
		contractsNumbers, startDate, endDate)
	Estructuras.Responder(http.StatusOK, "Datos consultados", results, ctx)
	return
}
