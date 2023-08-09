package PresupuestoModel

import (
	"appbiomedica/Conexion"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

//********************************************************************************************************
/* *********************    Funcion que agrega un presupuesto      ********************* */
func (p *Presupuesto) AgregarPresupuesto() (error, *Presupuesto) {
	db := Conexion.GetDB()
	tx := db.Begin()
	err := tx.Create(p).Error
	if err != nil {
		tx.Rollback()
		return err, nil
	}
	return tx.Commit().Error, p
}

//********************************************************************************************************
/* *********************    Funcion que regresa el presupuesto actual     ********************* */
func ConsultaPresupuesto() (error, *Presupuesto) {
	db := Conexion.GetDB()
	var datos Presupuesto
	err := db.Where("estado = ?", true).Find(&datos).Error
	if err != nil {
		return err, nil
	}
	return nil, &datos
}

//********************************************************************************************************
/* *********************    Funcion para actualizar el presupuesto     ********************* */
func (a *Presupuesto) ActualizarPresupuesto(id string, tx *gorm.DB) error {
	var presupuesto Presupuesto
	fmt.Println("esto llega al actualizar weee", a)
	err := tx.Where("id = ?", id).First(&presupuesto).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro el id")
		}
		return err
	}
	return tx.Model(&a).Where("id=?", id).Update(Presupuesto{Presupuesto: a.Presupuesto}).Error
}
