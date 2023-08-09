package PermisoControlador

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/PermisoModel"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

//*****************************************************************************************************************
//consulta todos los permisos
func GetPermisos(c *gin.Context) {
	err, resultado := PermisoModel.ConsultarPermisos()
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}

//*****************************************************************************************************************
//agregar un nuevo permiso
func PostPermiso(c *gin.Context) {
	var permisos PermisoModel.Permiso
	err := c.BindJSON(&permisos)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error al castear la estructura", err.Error(), c)
		return
	}
	arrayErrores, err := Estructuras.ValidaEstructura(permisos)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al leer estructura", err.Error(), c)
		return
	}
	if len(arrayErrores) > 0 {
		Estructuras.Responder(http.StatusBadRequest, "Estructura invalida", arrayErrores, c)
		return
	}
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al iniciar la transaccion", tx.Error, c)
		return
	}
	permisos.Id = bson.NewObjectId().Hex()
	permisos.Estado = true
	err = permisos.AgregarPermiso(tx)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al insertar los datos", err.Error(), c)
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al insertar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos agregados correctamente", permisos, c)
	return
}

//*****************************************************************************************************************
//Actualiza un permiso
func UpdatePermiso(c *gin.Context) {
	id := c.Param("id")
	var permisos PermisoModel.Permiso
	err := c.BindJSON(&permisos)
	if id == "" {
		Estructuras.Responder(http.StatusBadRequest, "Debes ingresar el id", nil, c)
		return
	}
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error en la estructura", err.Error(), c)
	}
	arrayErrores, err := Estructuras.ValidaEstructura(permisos)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al leer estructura", err.Error(), c)
		return
	}
	if len(arrayErrores) > 0 {
		Estructuras.Responder(http.StatusBadRequest, "Estructura invalida", arrayErrores, c)
		return
	}

	err = permisos.ActualizarPermiso(id)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al actualizar datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos actualizados correctamente", permisos, c)
	return
}

//********************************************************************************************
//Elimina un permiso
func DeletePermiso(c *gin.Context) {
	id := c.Param("id")
	//status :=c.Param("status")
	var permisos PermisoModel.Permiso
	if id == "" {
		Estructuras.Responder(http.StatusBadRequest, "Debes ingresar el id", nil, c)
		return
	}
	err := permisos.EliminarPermiso(id)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "No se pudo eliminar el registro", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Registro eliminado correctamente", permisos.Id, c)
	return
}
