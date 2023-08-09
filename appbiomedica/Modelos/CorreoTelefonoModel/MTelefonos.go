package CorreoTelefonoModel

import (
	"appbiomedica/Conexion"
	"github.com/jinzhu/gorm"
)

//********************************************************************************************************
/* *********************    Funcion para agregar un correo       ********************* */
func (p *Telefonos) AgregaTelefono(tx *gorm.DB) error {
	return tx.Create(p).Error
}

//********************************************************************************************************
/* *********************    Funcion que busca un correo     ********************* */
func BuscarTelefono(numero string) (bool, *Telefonos) {
	var telefono Telefonos
	db := Conexion.GetDB()
	err := db.Where("telefono = ?", numero).First(&telefono).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		return false, nil
	}
	// repetido error
	return true, &telefono
}

//********************************************************************************************************
/* *********************    Funcion para actualizar un numero telefonico     ********************* */
func (tel *Telefonos) ActualizarTelefono(id string, tx *gorm.DB) error {
	return tx.Model(&tel).Where("id=?", id).Update(Telefonos{Telefono: tel.Telefono}).Error
}
