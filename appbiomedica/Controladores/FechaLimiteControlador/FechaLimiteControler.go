package FechaLimiteControlador

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/FechaLimiteModel"
	"appbiomedica/Modulos/Estructuras"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

//*****************************************************************************************************************
/* *********************    Controlador para regresar la fecha limite actual     ********************* */
func GetFecha(c *gin.Context) {
	err, resultado := FechaLimiteModel.ConsultaFechaActual()
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", resultado, c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}

//*****************************************************************************************************************
/* *********************    Controlador para agregar una fecha limite       ********************* */
func PostFecha(c *gin.Context) {
	var fecha FechaLimiteModel.FechaLimite
	err := c.BindJSON(&fecha)
	fmt.Print(fecha)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error al castear la estructura", err.Error(), c)
		return
	}

	fecha.Id = bson.NewObjectId().Hex()
	fecha.Estado = true
	err, resultado := fecha.AgregarFecha()
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al insertar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos agregados correctamente", resultado, c)
	return
}

//*****************************************************************************************************************
/* *********************    Controlador para actuzalizar una fecha limite     ********************* */
func UpdateFecha(c *gin.Context) {
	var fechas FechaLimiteModel.FechaLimite
	err := c.BindJSON(&fechas)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error en la estructura", err.Error(), c)
	}
	fmt.Println(fechas)
	//inicio transaccion en la base de datos
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al iniciar la transaccion", tx.Error, c)
		return
	}
	err = fechas.ActualizarFecha(fechas.Id, tx)
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
	Estructuras.Responder(http.StatusOK, "Datos actualizados correctamente", fechas, c)
	return
}
