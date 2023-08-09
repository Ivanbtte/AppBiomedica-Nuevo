package ActionModuleControlador

import (
	"appbiomedica/Modelos/ModuloModel"
	"appbiomedica/Modelos/ServiciosSubrogados/TrasladoPacientesModel"
	"appbiomedica/Modulos/Estructuras"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetActionsByModule(ctx *gin.Context) {
	moduleId := ctx.DefaultQuery("moduleId", "")
	err, result := ModuloModel.ConsultActionsByModule(moduleId)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los modulos",
			err.Error(), ctx)
		return
	}
	Estructuras.Responder(http.StatusOK, "Modulos consultados", result, ctx)
}

//Funcion para consultar acciones de un modulo por unidad medica

func GetActionsByUnit(ctx *gin.Context) {
	unitId := ctx.DefaultQuery("unitId", "")
	moduleId := ctx.DefaultQuery("moduleId", "")
	unit, _ := strconv.Atoi(unitId)
	result, err := TrasladoPacientesModel.BuscarPermisosPorUnidad(unit, moduleId)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "error al consultar permisos", err, ctx)
		return
	}
	Estructuras.Responder(http.StatusOK, "datos consultados", result, ctx)
}
