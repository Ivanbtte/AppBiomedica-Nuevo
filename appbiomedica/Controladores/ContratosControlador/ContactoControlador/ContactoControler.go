package ContactoControlador

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/ContratoModel/ContactoModelo"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"net/http"
)

//********************************************************************************************************
/* *********************    Controlador para consultar los correos de un contrato     ********************* */
func GetCorreosContrato(c *gin.Context) {
	n_contrato := c.DefaultQuery("n_contrato", "")
	resultado, err := ContactoModelo.ConsultarCorreos(n_contrato)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}

//********************************************************************************************************
/* *********************    Controlador para consultar los correos de un contrato     ********************* */
func GetTelefonoContrato(c *gin.Context) {
	n_contrato := c.DefaultQuery("n_contrato", "")
	resultado, err := ContactoModelo.ConsultarTelefonos(n_contrato)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}

//********************************************************************************************************
/* *********************    Controlador para eliminar un correo de un contrato     ********************* */
func DeleteCorreoContrato(c *gin.Context) {
	id_correo := c.DefaultQuery("id_correo", "")
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al conectar a la base de datos", tx.Error, c)
		return
	}
	err, resultado := ContactoModelo.EliminarCorreoContrato(id_correo, tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al eliminar los datos", err.Error(), c)
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al eliminar el correo", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Correo eliminado correctamente", resultado, c)
	return
}
//********************************************************************************************************
/* *********************    Controlador para eliminar un telefono de un contrato     ********************* */
func DeleteTelefonoContrato(c *gin.Context) {
	id_telefono := c.DefaultQuery("id_telefono", "")
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al conectar a la base de datos", tx.Error, c)
		return
	}
	err, resultado := ContactoModelo.EliminarTelefonoContrato(id_telefono, tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al eliminar los datos", err.Error(), c)
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al eliminar el telefono", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Telefono eliminado correctamente", resultado, c)
	return
}