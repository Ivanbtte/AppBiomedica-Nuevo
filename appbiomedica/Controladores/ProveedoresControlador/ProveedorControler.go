package ProveedoresControlador

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/ProveedoresModel"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
	"strings"
)

//********************************************************************************************************
/* *********************    Controlador para consultar uno o todos los proveedores     ********************* */
func GetProveedores(c *gin.Context) {
	id_proveedor := c.DefaultQuery("id_proveedor", "")
	nombre_proveedor := c.DefaultQuery("nombre_proveedor", "")
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
	arrayFiltro := []string{}
	if nombre_proveedor != "" {
		arrayFiltro = strings.Split(nombre_proveedor, " ")
	}
	err, resultado, total := ProveedoresModel.ConsultarProveedor(id_proveedor, arrayFiltro, num_pag, num_reg)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", gin.H{"total": total, "registros": resultado}, c)
	return
}

//********************************************************************************************************
/* *********************    Controlador para consultar proveedores por nombres     ********************* */
func GetProveedoresNombre(c *gin.Context) {
	filtros := c.DefaultQuery("filtro", "")
	arrayFiltro := []string{}
	if filtros != "" {
		arrayFiltro = strings.Split(filtros, " ")
	}
	err, resultado := ProveedoresModel.ConsultarProveedorFiltro(arrayFiltro)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}

//********************************************************************************************************
/* *********************    Controlador para Agregar un nuevo proveedor     ********************* */
func PostProveedor(c *gin.Context) {
	var provedor ProveedoresModel.Proveedor
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
	provedor.Estado = false
	provedor.NombreEmpresa = Estructuras.FormatoTextoMayusculas(provedor.NombreEmpresa)
	provedor.RFC = Estructuras.FormatoTextoMayusculas(provedor.RFC)
	provedor.AliasEmpresa = Estructuras.FormatoTextoMayusculas(provedor.AliasEmpresa)
	duplicado := ProveedoresModel.VerificaProveedor(provedor.NProvImss)
	if duplicado {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Este Numero de proveedor ya esta registrado", err.Error(), c)
		return
	}
	err = provedor.AgregarProveedor(tx)
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
func UpdateProveedor(c *gin.Context) {
	var provedor ProveedoresModel.Proveedor
	id_provedor := c.DefaultQuery("id_proveedor", "")
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
	provedor.NombreEmpresa = Estructuras.FormatoTextoMayusculas(provedor.NombreEmpresa)
	provedor.RFC = Estructuras.FormatoTextoMayusculas(provedor.RFC)
	provedor.AliasEmpresa = Estructuras.FormatoTextoMayusculas(provedor.AliasEmpresa)
	if err = provedor.ActualizarProveedor(id_provedor, tx); err != nil {
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
	Estructuras.Responder(http.StatusOK, "Datos actualizados correctamente", provedor, c)
	return
}

//********************************************************************************************************
/* *********************    Controlador para eliminar un proveedor     ********************* */
func DeleteProveedor(c *gin.Context) {
	id_proveedor := c.DefaultQuery("id_proveedor", "")
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al conectar a la base de datos", tx.Error, c)
		return
	}
	err, resultado := ProveedoresModel.EliminarProveedor(id_proveedor, tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al eliminar los datos", err.Error(), c)
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al eliminar el proveedor", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Proveedor Eliminado Correctamente", resultado, c)
	return
}
