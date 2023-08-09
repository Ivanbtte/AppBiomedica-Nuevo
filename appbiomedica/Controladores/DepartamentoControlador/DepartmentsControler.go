package DepartamentoControlador

import (
	"appbiomedica/Modelos/DepartamentosModel"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetDepartmentHieracy(ctx *gin.Context) {
	unitMed := ctx.DefaultQuery("unitMedId", "")
	clvDele := ctx.DefaultQuery("clvDele", "")
	if unitMed != "0" {
		unitMedId, err := strconv.Atoi(unitMed)
		if err != nil {
			Estructuras.Responder(http.StatusBadRequest, "Numero de unidad medica incorrecto", err.Error(), ctx)
			return
		}
		err, parentUnit := DepartamentosModel.ConsultParentUnitId(unitMedId)
		if err != nil {
			Estructuras.Responder(http.StatusBadRequest, "No se han registrado departamentos para la unidad seleccionada",
				nil, ctx)
			return
		}
		err, result := DepartamentosModel.ConsultDepartmentUnit(parentUnit.Id)
		if err != nil {
			Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los departamentos", err.Error(), ctx)
			return
		}
		Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", result, ctx)
		return
	}
	if clvDele != "" {
		err, parentDele := DepartamentosModel.ConsultParentDelegationId(clvDele)
		if err != nil {
			Estructuras.Responder(http.StatusBadRequest, "No se han registrado departamentos para esta delegacion",
				nil, ctx)
			return
		}
		err, result := DepartamentosModel.ConsultDepartmentDelegation(parentDele.Id)
		if err != nil {
			Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los departamentos", err.Error(), ctx)
			return
		}
		Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", result, ctx)
		return
	}

}
