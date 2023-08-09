package ProveedoresModel

import (
	"appbiomedica/Conexion"
	"errors"
	"github.com/jinzhu/gorm"
	"strings"
)

//********************************************************************************************************
/* *********************    Funcion para agregar un contacto de un proveedor      ********************* */
func (p *ContactoProveedor) AgregarContactoProveedor(tx *gorm.DB) error {
	return tx.Create(p).Error
}

//********************************************************************************************************
/* *********************    Funcion para actualizar un contacto de un proveedor     ********************* */
func (p *ContactoProveedor) ActualizarContactoProveedor(id string, tx *gorm.DB) error {
	var proveedor ContactoProveedor
	err := tx.Where("id = ?", id).First(&proveedor).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro el id")
		}
		return tx.Rollback().Error
	}
	return tx.Model(&p).Where("id=?", id).Update(ContactoProveedor{NombreCompleto: p.NombreCompleto,
		Departamento: p.Departamento, Cargo: p.Cargo, Telefono: p.Telefono, Extension: p.Extension, Celular: p.Celular,
		Correo: p.Correo, Comentarios: p.Comentarios, Asunto: p.Asunto}).Error
}

//********************************************************************************************************
/* *******************    Funcion para regresar la lista de proveedores o un solo proveedor     ******************** */
func ConsultarContactoProveedor(id_contacto, id_proveedor, nombre string) (error, *[]ContactoProveedor) {
	var resultado []ContactoProveedor
	db := Conexion.GetDB()
	if id_contacto != "" {
		db = db.Where("id = ?", id_contacto)
	}
	if nombre != "" {
		db = db.Where("LOWER (nombre_completo) LIKE ?", "%"+strings.ToLower(nombre)+"%")
	}
	err := db.Where("proveedor_n_prov_imss = ?", id_proveedor).Preload("Proveedor").Find(&resultado).Error
	if err != nil {
		return err, nil
	}
	return nil, &resultado
}

//********************************************************************************************************
/* *******************    Funcion para buscar un contacto del proveedor     ******************** */
func ConsultarContacto(nombre, nprov_imss string) *ContactoProveedor {
	var resultado ContactoProveedor
	db := Conexion.GetDB()
	err := db.Where("LOWER (nombre_completo) LIKE ? AND proveedor_n_prov_imss = ?", "%"+strings.ToLower(nombre)+"%", nprov_imss).
		Find(&resultado).Error
	if err != nil {
		return nil
	}
	return &resultado
}

//********************************************************************************************************
/* *********************    Funcion para eliminar un proveedor     ********************* */
func EliminarContactoProveedor(id string, tx *gorm.DB) (error, *ContactoProveedor) {
	var item ContactoProveedor
	err := tx.Where("id = ?", id).Find(&item).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro el id"), nil
		}
		return err, nil
	}
	return tx.Model(&ContactoProveedor{}).Where("id=?", id).Delete(&item).Error, &item
}
