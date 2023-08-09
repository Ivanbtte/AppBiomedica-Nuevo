package PuestoRoleControlador

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/DepartamentosModel"
	"appbiomedica/Modulos/Estructuras"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func PostPuestoRole(ctx *gin.Context) {
	var puestoRole DepartamentosModel.PuestoOrganigramaRole
	err := ctx.BindJSON(&puestoRole)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error al castear la estructura", err.Error(), ctx)
		return
	}

	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al iniciar la transaccion", tx.Error, ctx)
		return
	}
	puestoRole.Id = bson.NewObjectId().Hex()
	puestoRole.Status = true
	err = puestoRole.AddPuestoOrganigramaRole(tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "error al guardar los datos", err, ctx)
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al realizar el commit", err.Error(), ctx)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos agregados correctamente", nil, ctx)
}
