package UsuarioControlador

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/DepartamentosModel"
	"appbiomedica/Modelos/MenuModel"
	"appbiomedica/Modelos/RoleModel"
	"appbiomedica/Modelos/UsuarioModel"
	"appbiomedica/Modulos/Estructuras"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func GetUserByPosition(ctx *gin.Context) {
	userId := ctx.DefaultQuery("userId", "")
	user, err := DepartamentosModel.ConsultaUnUsuarioPuesto(userId)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar usuarios del departamento",
			err.Error(),
			ctx)
		return
	}
	Estructuras.Responder(http.StatusOK, "Usuarios consultados", user, ctx)
}
func GetUserValidation(ctx *gin.Context) {
	status := ctx.DefaultQuery("status", "")
	valid := ctx.DefaultQuery("valid", "")
	department := ctx.DefaultQuery("departmentId", "")
	unitMed := ctx.DefaultQuery("unitMedId", "0")
	clvDele := ctx.DefaultQuery("clvDele", "")
	unitMedId, _ := strconv.Atoi(unitMed)
	statusBool, err := strconv.ParseBool(status)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al parsear booleano", err.Error(), ctx)
		return
	}
	validBool, err := strconv.ParseBool(valid)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al parsear booleano", err.Error(), ctx)
		return
	}
	err, results := DepartamentosModel.ConsultUsersNotValid(statusBool, validBool, unitMedId, clvDele, department)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), ctx)
		return
	}
	Estructuras.Responder(http.StatusOK, "Usuarios consultados", results, ctx)
}

func AddNewUsuarioPuesto(ctx *gin.Context) {
	userType := ctx.DefaultQuery("userType", "")
	usuarioPuesto := DepartamentosModel.UsuarioPuesto{}
	update := ctx.DefaultQuery("update", "")
	usuario := UsuarioModel.Usuario{}
	err := ctx.BindJSON(&usuarioPuesto)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "error al analizar la estructura", err.Error(), ctx)
		return
	}
	state, _, err := DepartamentosModel.CheckNotRepeatUsuarioPuesto(usuarioPuesto.UsuarioId)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "error al buscar si el usuario tiene un puesto asignado", err.Error(), ctx)
		return
	}
	if update == "" {
		if state {
			Estructuras.Responder(http.StatusInternalServerError, "error este usuario ya tiene un puesto", nil, ctx)
			return
		}
	} else {

	}
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "error al iniciar la transacción de la base de datos",
			tx.Error, ctx)
		return
	}
	usuarioPuesto.Id = bson.NewObjectId().Hex()
	usuarioPuesto.FechaRegistro = time.Now()
	usuarioPuesto.Valid = false
	usuarioPuesto.Estado = true
	err = usuarioPuesto.AgregaUsuarioPuesto(tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "error al agregar el usuario al puesto",
			tx.Error, ctx)
		return
	}
	/***** asignar el permiso por default al usuario *****/
	var usuarioPuestoRole DepartamentosModel.UsuarioPuestoRole
	usuarioPuestoRole.Id = bson.NewObjectId().Hex()
	usuarioPuestoRole.UsuarioPuestoId = usuarioPuesto.Id
	usuarioPuestoRole.RoleId = "614c97bbd7be0a789bb6f02c" //Role Consulta por default
	usuarioPuestoRole.Status = true
	err = usuarioPuestoRole.AddUsuarioPuestoRoleModule(tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al asignar el permiso por default", err.Error(),
			ctx)
		return
	}
	usuario.UsuarioTipoId = userType
	err = usuario.UpdateUserType(usuarioPuesto.UsuarioId, tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "error al actualizar el tipo de usuario", err.Error(),
			ctx)
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al realizar el commit a la base de datos",
			err.Error(), ctx)
		return
	}
	err, rolesIds := ConsultRolesIdsByUser(usuarioPuesto.Id)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los roles", err.Error(), ctx)
		return
	}
	err, menuRole := MenuModel.ConsultMenuByRole(rolesIds)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "error al consultar permisos", err.Error(), ctx)
		return
	}
	err, permits := RoleModel.ConsultNameModule(rolesIds)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "error al consultar permisos", err.Error(), ctx)
		return
	}

	Estructuras.Responder(http.StatusOK, "Usuario registrado correctamente al puesto", gin.H{"data": usuarioPuesto,
		"menu": menuRole, "permits": permits}, ctx)
}

func AddNewRolePuesto(ctx *gin.Context) {
	var usuarioRole DepartamentosModel.UsuarioPuestoRole
	err := ctx.BindJSON(&usuarioRole)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "error al analizar la estructura", err.Error(), ctx)
		return
	}
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "error al iniciar la transacción de la base de datos",
			tx.Error, ctx)
		return
	}
	err = usuarioRole.AddUsuarioPuestoRoleModule(tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al asignar el rol al puesto", err.Error, ctx)
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al realizar el commit a la base de datos",
			err.Error(), ctx)
		return
	}
	Estructuras.Responder(http.StatusOK, "Rol asignado correctamente", usuarioRole, ctx)
}
