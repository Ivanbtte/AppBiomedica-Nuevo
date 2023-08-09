package DepartamentoControlador

import (
	"appbiomedica/Modelos/DepartamentosModel"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPositionOrganization(c *gin.Context) {
	departmentUnitId := c.DefaultQuery("departmentUnitId", "")
	departmentDelegationId := c.DefaultQuery("departmentDelegationId", "")
	err, result := DepartamentosModel.ConsultPositionOrganization(departmentUnitId, departmentDelegationId)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los puestos", err.Error(), c)
		return
	}
	if len(*result) != 0 {
		err, final := validPosition(*result)
		if err != nil {
			Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los puestos", err.Error(), c)
			return
		}
		Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", final, c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", result, c)
	return
}

func validPosition(available []DepartamentosModel.PuestoOrganigrama) (error, []DepartamentosModel.PuestoOrganigrama) {
	positions := make([]DepartamentosModel.PuestoOrganigrama, 0)
	for _, item := range available {
		err, total := DepartamentosModel.CountPositionsBusy(item.Id)
		if err != nil {
			return err, nil
		}
		if total == 0 || total < (2*item.NumeroPlazas) {
			positions = append(positions, item)
		}
	}
	return nil, positions
}
