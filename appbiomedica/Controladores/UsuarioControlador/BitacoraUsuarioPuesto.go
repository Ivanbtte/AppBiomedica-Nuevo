package UsuarioControlador

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/AuthModel"
	"appbiomedica/Modelos/DepartamentosModel"
	"appbiomedica/Modelos/UsuarioModel"
	"appbiomedica/Modulos/Estructuras"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func ValidarUsuarioPuesto(ctx *gin.Context) {
	bitacora := DepartamentosModel.BitacoraUsuarioPuesto{}
	var usuarioRole DepartamentosModel.UsuarioPuestoRole
	status := ctx.DefaultQuery("status", "")
	roleId := ctx.DefaultQuery("roleId", "")
	b, _ := strconv.ParseBool(status)
	err := ctx.BindJSON(&bitacora)
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
	res := DepartamentosModel.CheckPermitByRole("614ca03ed7be0a0b8e27fca4", bitacora.UsuarioModificadoId, false)
	if !res {
		usuarioRole.Id = bson.NewObjectId().Hex()
		usuarioRole.UsuarioPuestoId = bitacora.UsuarioModificadoId
		usuarioRole.RoleId = roleId
		usuarioRole.Status = true
		err = usuarioRole.AddUsuarioPuestoRoleModule(tx)
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "error al asignar el role al usuario", err.Error, ctx)
			return
		}
	} else {
		err = DepartamentosModel.ValidarUsuarioRole(bitacora.UsuarioModificadoId, "614ca03ed7be0a0b8e27fca4", true, tx)
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "error al validar usuario en el role", err, ctx)
			return
		}
	}
	err = DepartamentosModel.CambiarValidacionUsuarioPuesto(bitacora.UsuarioModificadoId, b, tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "error al actualizar el estado del usuario", err, ctx)
		return
	}

	bitacora.Id = bson.NewObjectId().Hex()
	bitacora.Fecha = time.Now()
	bitacora.Status = true
	err = bitacora.AddBitacoraUsuarioPuesto(tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "error al agregar la bitacora", err, ctx)
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al realizar el commit a la base de datos",
			err.Error(), ctx)
		return
	}
	Estructuras.Responder(http.StatusOK, "Usuario validado correctamente.", nil, ctx)
}

func GetTipoModificacion(ctx *gin.Context) {
	notId := ctx.DefaultQuery("notId", "")
	results, err := DepartamentosModel.ConsultTipoModificacion(notId)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "error al consultar los tipos de modificaciones", nil, ctx)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", results, ctx)
}

func UnsuscribeUsuarioPuesto(ctx *gin.Context) {
	unsuscribe := DepartamentosModel.BitacoraUsuarioPuesto{}
	idUser := ctx.DefaultQuery("idUser", "")
	err := ctx.BindJSON(&unsuscribe)
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
	fmt.Println(unsuscribe)
	if unsuscribe.TipoModificacionId == "61f1a790d7be0a2d4643b6ab" {
		//se eliminara del puesto, los roles, el usuario y los datos personales, credenciales
		fmt.Print("eliminar todos los registros del usuario")
		err = DepartamentosModel.CambiarValidacionUsuarioPuesto(unsuscribe.UsuarioModificadoId, false, tx)
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "error quitar la validacion del usuario en el puesto", err, ctx)
			return
		}
		err = DepartamentosModel.CambiarEstadoUsuarioPuesto(unsuscribe.UsuarioModificadoId, false, tx)
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "error al deshabilitar el usuario en el puesto", err, ctx)
			return
		}
		err = UsuarioModel.ValidarUsuario(idUser, false, tx)
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "error al deshabilitar usuario", err, ctx)
			return
		}
		err = AuthModel.ValidaCredencial(idUser, false, tx)
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "error al deshabilitar las credenciales", err, ctx)
			return
		}

	} else if unsuscribe.TipoModificacionId == "61f1a79ad7be0a2e43caa8c9" {
		//eliminar des puesto el usuario y roles
		fmt.Print("solo eliminar del puesto")
		err = DepartamentosModel.CambiarEstadoActionException(unsuscribe.UsuarioModificadoId, false, tx)
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "error al quitar los permisos especiales", err.Error(), ctx)
			return
		}
		err = DepartamentosModel.CambiarEstadoUsuarioRole(unsuscribe.UsuarioModificadoId,"614ca03ed7be0a0b8e27fca4", false, tx)
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "error al quitar los permisos del puesto", err, ctx)
			return
		}
		err = DepartamentosModel.CambiarValidacionUsuarioPuesto(unsuscribe.UsuarioModificadoId, false, tx)
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "error al actualizar el estado del usuario", err, ctx)
			return
		}
		err = DepartamentosModel.CambiarEstadoUsuarioPuesto(unsuscribe.UsuarioModificadoId, false, tx)
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "error al actualizar el estado del usuario", err, ctx)
			return
		}
	} else {
		//quitar validacion del usuario puesto
		err = DepartamentosModel.CambiarValidacionUsuarioPuesto(unsuscribe.UsuarioModificadoId, false, tx)
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "error al actualizar el estado del usuario", err, ctx)
			return
		}
		err = DepartamentosModel.CambiarEstadoUsuarioRole(unsuscribe.UsuarioModificadoId,"614ca03ed7be0a0b8e27fca4", false, tx)
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "error al quitar los permisos del puesto", err, ctx)
			return
		}
	}
	unsuscribe.Id = bson.NewObjectId().Hex()
	unsuscribe.Fecha = time.Now()
	unsuscribe.Status = true
	err = unsuscribe.AddBitacoraUsuarioPuesto(tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "error al agregar la bitacora", err, ctx)
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al realizar el commit a la base de datos",
			err.Error(), ctx)
		return
	}
	Estructuras.Responder(http.StatusOK, "Usuario validado correctamente.", nil, ctx)
}
