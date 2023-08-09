package PermisoModel

import (
	"appbiomedica/Conexion"
	"errors"
	"github.com/jinzhu/gorm"
)

//********************************************************************************************************
//-----------------------Consultar todos los permisos existentes

func ConsultarPermisos() (error, *[] Permiso) {
	var datos [] Permiso
	db := Conexion.GetDB()
	err := db.Find(&datos).Error
	if err != nil {
		return err, nil
	}
	return nil, &datos

}

//********************************************************************************************************
//----------------------------Agrega un permiso
func (p *Permiso) AgregarPermiso(tx *gorm.DB) error {
	return tx.Create(p).Error

}

//********************************************************************************************************
//-----------------------Actualizar un permiso

func (a Permiso) ActualizarPermiso(id string) error {
	db := Conexion.GetDB()
	var permisos Permiso
	tx := db.Begin()
	err := tx.Where("id = ?", id).First(&permisos).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro el id")
		}
		return err
	}
	err = tx.Model(&a).Where("id=?", id).Update(Permiso{Descripcion: a.Descripcion}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

//********************************************************************************************************
//-----------------------Eliminar un permiso

func (a Permiso) EliminarPermiso(id string) error {
	db := Conexion.GetDB()
	var permisos Permiso
	tx := db.Begin()
	err := tx.Where("id = ?", id).First(&permisos).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro el id")
		}
		return err
	}
	if !permisos.Estado {
		tx.Rollback()
		return errors.New("Este id ya ha sido inhabilitado")
	}
	err = tx.Model(&a).Where("id=?", id).Where("estado = ?", true).Update("estado", false).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
