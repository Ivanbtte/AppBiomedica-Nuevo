package UsuarioControlador

import (
	"appbiomedica/Modelos/DepartamentosModel"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUnidadUsuario(ctx *gin.Context) {
	userId := ctx.DefaultQuery("usuario_id", "")
	err, result := DepartamentosModel.ConsultUnitIdByUser(userId)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar unidad", err.Error(), ctx)
		return
	}
	Estructuras.Responder(http.StatusOK, "Unidad consultada", result, ctx)
}

type UserPermit struct {
	Permit bool                             `json:"permit"`
	User   DepartamentosModel.UsuarioPuesto `json:"user,omitempty"`
}

func GetUserByDepartment(ctx *gin.Context) {
	departmentId := ctx.DefaultQuery("departmentId", "")
	userId := ctx.DefaultQuery("userId", "")
	err, result := DepartamentosModel.ConsultUserByDepartment(departmentId, userId)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar usuarios del departamento",
			err.Error(),
			ctx)
		return
	}
	/***** Buscar si el usuario tiene algún permiso para el módulo de traslados *****/
	usersPermits := make([]UserPermit, 0)
	for _, user := range *result {
		state := DepartamentosModel.CheckPermitByRole("614ca03ed7be0a0b8e27fca4", user.Id,true)
		stateRol := DepartamentosModel.CheckPermitRoleByOrganigrama("614ca03ed7be0a0b8e27fca4", user.PuestoOrganigramaId)
		usersPermits = append(usersPermits, UserPermit{Permit: state || stateRol, User: user})
	}
	Estructuras.Responder(http.StatusOK, "Usuarios consultados", usersPermits, ctx)
}
