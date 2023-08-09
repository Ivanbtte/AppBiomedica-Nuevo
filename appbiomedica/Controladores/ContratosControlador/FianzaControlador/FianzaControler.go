package FianzaControlador

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/ContratoModel/FianzaModelo"
	"appbiomedica/Modulos/Estructuras"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

//********************************************************************************************************
/* *********************    Controlador para consultar una fianza de un contrato     ********************* */
func GetFianza(c *gin.Context) {
	numero_contrato := c.DefaultQuery("numero_contrato", "")
	estado := false
	err, datos := FianzaModelo.ConsultarFianza(numero_contrato)
	if datos != nil {
		estado = true
	}
	if err != nil {
		fmt.Println("Resultado convencional", datos)
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos",
			gin.H{"datos": err.Error(), "estado": estado}, c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", gin.H{"datos": datos,
		"estado": estado}, c)
	return
}

//********************************************************************************************************
/* *********************    Controlador para agregar una fianza de un contrato     ********************* */
func PostFianza(c *gin.Context) {
	numero_contrato := c.DefaultQuery("numero_contrato", "")
	var fianza FianzaModelo.Fianza
	err := c.BindJSON(&fianza)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error al castear la estructura", err.Error(), c)
		return
	}
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al iniciar la transaccion", tx.Error, c)
		return
	}
	fianza.Id = bson.NewObjectId().Hex()
	fianza.Estado = true
	fianza.ContratoNumeroContrato = numero_contrato
	err = fianza.Agregarfianza(tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al insertar los datos", err.Error(), c)
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al realizar el commit", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos agregados correctamente", fianza, c)
	return
}
