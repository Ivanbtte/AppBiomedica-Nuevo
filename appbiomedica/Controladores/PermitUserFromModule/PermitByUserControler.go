package PermitUserFromModule

import (
	"appbiomedica/Modelos/ModuloModel"
)

type ActionsPermit struct {
	Action ModuloModel.Action `json:"action,omitempty"`
	Permit bool               `json:"permit"`
}

//func GetPermitsByUser(ctx *gin.Context) {
//	moduleId := ctx.DefaultQuery("moduleId", "")
//	userId := ctx.DefaultQuery("userId", "")
//	/*****  Consult Actions by module  *****/
//	err, actions := ModuloModel.ConsultActionsByModule(moduleId)
//	if err != nil {
//		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los submodulos",
//			err.Error(), ctx)
//		return
//	}
//	actionPermit := make([]ActionsPermit, 0)
//	for _, item := range *actions {
//		_, permits := DepartamentosModel.ConsultPermitsByUser(item.Id, userId)
//		if permits != nil {
//			actionPermit = append(actionPermit, ActionsPermit{
//				Action: item,
//				Permit: permits.Permit,
//			})
//		} else {
//			actionPermit = append(actionPermit, ActionsPermit{
//				Action: item,
//				Permit: false,
//			})
//		}
//
//	}
//	Estructuras.Responder(http.StatusOK, "Submodulos consultados", actionPermit, ctx)
//}
