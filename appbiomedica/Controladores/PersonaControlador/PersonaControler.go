package PersonaControlador

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/PersonaModel"
	"appbiomedica/Modulos/Estructuras"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
)

//*****************************************************************************************************************
//consulta todos las personas
func GetPersonas(c *gin.Context) {
	num_pagina := c.DefaultQuery("pagina", "1")
	num_registro := c.DefaultQuery("num_registros", "50")
	num_pag, err := strconv.Atoi(num_pagina)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Numero de pagina incorrecto", err.Error(), c)
		return
	}
	num_reg, err := strconv.Atoi(num_registro)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Numero de registros incorrecto", err.Error(), c)
		return
	}
	err, resultado, total := PersonaModel.ConsultaPersonas(num_pag, num_reg)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", gin.H{"total": total, "registros": resultado}, c)
	return
}

//*****************************************************************************************************************
//consulta  una persona
func GetUnaPersona(c *gin.Context) {
	id := c.Param("id")
	err, resultado := PersonaModel.ConsultaUnaPersona(id)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}

//*****************************************************************************************************************
//agregar una nueva persona
func PostPersona(c *gin.Context) {
	var personas PersonaModel.Persona
	err := c.BindJSON(&personas)
	fmt.Print(personas)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error al castear la estructura", err.Error(), c)
		return
	}
	/*	if personas.Estado == true {
		err = PersonaModel.VerificaCorreo(personas.Correo)
		if err != nil {
			Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
			return
		}
		personas.Estado = true

	}*/
	if PersonaModel.VerificaCorreoPersona(personas.Correo) {
		Estructuras.Responder(http.StatusInternalServerError, "Este correo ya ha sido registrado", nil, c)
		return
	}
	arrayErrores, err := Estructuras.ValidaEstructura(personas)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al leer estructura", err.Error(), c)
		return
	}
	if len(arrayErrores) > 0 {
		Estructuras.Responder(http.StatusBadRequest, "Estructura invalida", arrayErrores, c)
		return
	}
	if !govalidator.IsTime(personas.FechaNac, "2006-01-02") {
		Estructuras.Responder(http.StatusBadRequest, "Fecha de nacimiento, formato incorrecto tiene que ser "+
			"(yyyy-mm-dd)", nil, c)
		return
	}
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al iniciar la transaccion", tx.Error, c)
		return
	}
	personas.Id = bson.NewObjectId().Hex()
	personas.Estado = true
	err = personas.AgregaPersona(tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al insertar los datos", err.Error(), c)
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al insertar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos agregados correctamente", personas, c)
	return
}

//*****************************************************************************************************************
//Actualiza una persona

func UpdatePersona(c *gin.Context) {
	id := c.Param("id")
	var personas PersonaModel.Persona
	err := c.BindJSON(&personas)
	if id == "" {
		Estructuras.Responder(http.StatusBadRequest, "Debes ingresar el id", nil, c)
		return
	}
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error en la estructura", err.Error(), c)
		return
	}
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al iniciar la base de datos", tx.Error, c)
		return
	}
	err = personas.ActualizarPersona(id,tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al actualizar datos", err.Error(), c)
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al insertar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos actualizados correctamente", personas, c)
	return
}

/*//********************************************************************************************
//Elimina una persona
func DeletePersona(c *gin.Context) {
	id := c.Param("id")
	//status :=c.Param("status")
	var personas PersonaModel.Persona
	if id == "" {
		Estructuras.Responder(http.StatusBadRequest, "Debes ingresar el id", nil, c)
		return
	}
	err := personas.EliminarPersona(id)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "No se pudo eliminar el registro", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Registro eliminado correctamente", personas, c)
	return
}*/

func CheckAvailableEnrollment(ctx *gin.Context) {
	enrollment := ctx.DefaultQuery("enrollment", "")
	status, result := PersonaModel.CheckAvailablePerson(enrollment)
	if status {
		Estructuras.Responder(http.StatusOK, "no hay registro de esta matricula",
			gin.H{"status": true}, ctx)
		return
	}
	if !status {
		if result.Estado {
			Estructuras.Responder(http.StatusOK, "esta matricula se encuentra activa", gin.H{"status": false,
				"personStatus": true}, ctx)
			return
		} else {
			Estructuras.Responder(http.StatusOK, "esta matricula se encuentra registrada pero esta dada de baja",
				gin.H{"status": false, "personStatus": false, "data": result}, ctx)
			return
		}
	}
}
