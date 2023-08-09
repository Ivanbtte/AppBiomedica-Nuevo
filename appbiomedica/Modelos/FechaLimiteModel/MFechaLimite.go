package FechaLimiteModel

import (
	"appbiomedica/Conexion"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

//********************************************************************************************************
/* *********************    Funcion que agrega una fecha limite      ********************* */
func (f *FechaLimite) AgregarFecha() (error, *FechaLimite) {
	db := Conexion.GetDB()
	tx := db.Begin()
	err := tx.Create(f).Error
	if err != nil {
		tx.Rollback()
		return err, nil
	}
	return tx.Commit().Error, f
}

//********************************************************************************************************
/* *********************    Funcion que regresa la fecha limite actual     ********************* */
func ConsultaFechaActual() (error, *FechaLimite) {
	db := Conexion.GetDB()
	var datos FechaLimite
	err := db.Where("estado = ?", true).Find(&datos).Error
	if err != nil {
		return err, nil
	}
	return nil, &datos
}

//********************************************************************************************************
/* *********************    Funcion para actualizar una fecha limite     ********************* */
func (a *FechaLimite) ActualizarFecha(id string, tx *gorm.DB) error {
	var fechas FechaLimite
	fmt.Println("esto llega al actualizar weee", a)
	err := tx.Where("id = ?", id).First(&fechas).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro el id")
		}
		return err
	}
	return tx.Model(&a).Where("id=?", id).Update(FechaLimite{FechaInicio: a.FechaInicio, FechaLimite: a.FechaLimite}).Error
}
func VerificaFecha() (*FechaLimite, error) {
	var fechas FechaLimite
	db := Conexion.GetDB()
	err := db.Where("estado = ?", true).Find(&fechas).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("Aun no hay fechas establecidas")
		}
		return nil, err
	}
	fmt.Println("estas son las dehcas", &fechas)
	return &fechas, nil
}
