package ProveedoresControlador

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/ProveedoresModel"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

//********************************************************************************************************
/* ****************    Controlador para consultar uno o todos los contactos de los proveedores     ***************** */
func GetContactoProveedores(c *gin.Context) {
	id_proveedor := c.DefaultQuery("id_proveedor", "")
	id_contacto := c.DefaultQuery("id_contacto", "")
	nombre_contacto := c.DefaultQuery("nombre_contacto", "")
	err, resultado := ProveedoresModel.ConsultarContactoProveedor(id_contacto, id_proveedor, nombre_contacto)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}

//********************************************************************************************************
/* *********************    Controlador para Agregar un nuevo proveedor     ********************* */
func PostContactoProveedor(c *gin.Context) {
	var provedor ProveedoresModel.ContactoProveedor
	err := c.BindJSON(&provedor)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error al castear la estructura", err.Error(), c)
		return
	}
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al iniciar la transaccion", tx.Error, c)
		return
	}
	provedor.Id = bson.NewObjectId().Hex()
	provedor.Estado = true
	err = provedor.AgregarContactoProveedor(tx)
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
	Estructuras.Responder(http.StatusOK, "Datos agregados correctamente", provedor, c)
	return
}

//********************************************************************************************************
/* *********************    Controlador para actualizar un proveedor ya registrado     ********************* */
func UpdateContactoProveedor(c *gin.Context) {
	var provedor ProveedoresModel.ContactoProveedor
	id_provedor := c.DefaultQuery("id_contacto", "")
	err := c.BindJSON(&provedor)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error al castear la estructura", err.Error(), c)
		return
	}
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al conectar a la base de datos", tx.Error, c)
		return
	}
	if err = provedor.ActualizarContactoProveedor(id_provedor, tx); err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al actualizar los datos", err.Error(), c)
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al actualizar el proveedor", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos agregados correctamente", provedor, c)
	return
}

//********************************************************************************************************
/* *********************    Controlador para eliminar un proveedor     ********************* */
func DeleteContactoProveedor(c *gin.Context) {
	id_proveedor := c.DefaultQuery("id_contacto", "")
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al conectar a la base de datos", tx.Error, c)
		return
	}
	err, resultado := ProveedoresModel.EliminarContactoProveedor(id_proveedor, tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al eliminar los datos", err.Error(), c)
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al eliminar el contacto", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Contacto Eliminado Correctamente", resultado, c)
	return
}
