package ProveedoresModel

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modulos/Query"
	"errors"
	"github.com/jinzhu/gorm"
)

//********************************************************************************************************
/* *********************    Funcion para agregar un proveedor nuevo     ********************* */
func (p *Proveedor) AgregarProveedor(tx *gorm.DB) error {
	return tx.Create(p).Error
}

//********************************************************************************************************
/* *********************    Funcion para actualizar un proveedor     ********************* */
func (p *Proveedor) ActualizarProveedor(id string, tx *gorm.DB) error {
	var proveedor Proveedor
	err := tx.Where("id = ?", id).First(&proveedor).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro el id")
		}
		return tx.Rollback().Error
	}
	return tx.Model(&p).Where("id=?", id).Update(Proveedor{AliasEmpresa: p.AliasEmpresa,
		NombreEmpresa: p.NombreEmpresa, Direccion: p.Direccion, Correo: p.Correo, RFC: p.RFC, NProvImss: p.NProvImss,
		Telefono: p.Telefono, EstadosId: p.EstadosId, Municipio: p.Municipio, CP: p.CP}).Error
}

//********************************************************************************************************
/* *******************    Funcion para regresar la lista de proveedores o un solo proveedor     ******************** */
func ConsultarProveedor(id string, nombre []string, numPagina, numRegistros int) (error, *[]Proveedor, int) {
	var resultado []Proveedor
	var total int
	db := Conexion.GetDB()
	if id != "" {
		db = db.Where("id = ?", id)
	}
	if len(nombre) != 0 {
		db = Query.ContruirQueryProveedor(db, nombre)
	}
	err := db.Find(&[]Proveedor{}).Count(&total).Error
	if err != nil {
		return err, nil, total
	}
	err = db.Offset((numPagina - 1) * numRegistros).Limit(numRegistros).Preload("Estados").Find(&resultado).Error
	if err != nil {
		return err, nil, total
	}
	return nil, &resultado, total
}

//********************************************************************************************************
/* *********************    Funcion para eliminar un proveedor     ********************* */
func EliminarProveedor(id string, tx *gorm.DB) (error, *Proveedor) {
	var item Proveedor
	err := tx.Where("id = ?", id).Find(&item).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro el id"), nil
		}
		return err, nil
	}
	return tx.Model(&Proveedor{}).Delete(&item).Error, &item
}

//********************************************************************************************************
/* *********************    Funcion para verificar que no exita un proveedor    ********************* */
func VerificaProveedor(n_p_imss string) bool {
	var provedor Proveedor
	db := Conexion.GetDB()
	err := db.Where("n_prov_imss = ?", n_p_imss).First(&provedor).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false
		}
		return false
	}
	// si esta
	return true
}

//********************************************************************************************************
/* *******************    Funcion para consultar proveedores por filtros     ******************** */
func ConsultarProveedorFiltro(filtro []string) (error, *[]Proveedor) {
	var resultado []Proveedor
	db := Conexion.GetDB()
	if len(filtro) != 0 {
		db = Query.ContruirQueryProveedor(db, filtro)
	}
	err := db.Order("nombre_empresa").Find(&resultado).Error
	if err != nil {
		return err, nil
	}
	return nil, &resultado
}
