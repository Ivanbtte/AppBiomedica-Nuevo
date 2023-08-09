package PresupuestoControlador

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/PresupuestoModel"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

//*****************************************************************************************************************
/* *********************    Controlador para consultar un presupuesto     ********************* */
func GetPresupuesto(c *gin.Context) {
	err, resultado := PresupuestoModel.ConsultaPresupuesto()
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", resultado, c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}

//*****************************************************************************************************************
/* *********************    Controlador para agregar un presupuesto       ********************* */
func PostPresupuesto(c *gin.Context) {
	var presupuesto PresupuestoModel.Presupuesto
	err := c.BindJSON(&presupuesto)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error al castear la estructura", err.Error(), c)
		return
	}

	presupuesto.Id = bson.NewObjectId().Hex()
	presupuesto.Estado = true
	err, resultado := presupuesto.AgregarPresupuesto()
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al insertar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos agregados correctamente", resultado, c)
	return
}

//*****************************************************************************************************************
/* *********************    Controlador para actuzalizar un presupuesto     ********************* */
func UpdatePresupuesto(c *gin.Context) {
	var presupuesto PresupuestoModel.Presupuesto
	err := c.BindJSON(&presupuesto)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error en la estructura", err.Error(), c)
	}
	//inicio transaccion en la base de datos
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al iniciar la transaccion", tx.Error, c)
		return
	}
	err = presupuesto.ActualizarPresupuesto(presupuesto.Id, tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al actualizar datos", err.Error(), c)
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al actualizar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos actualizados correctamente", presupuesto, c)
	return
}
