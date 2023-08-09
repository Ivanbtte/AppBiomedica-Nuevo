package ContactoModelo

import (
	"appbiomedica/Conexion"
	"errors"
	"github.com/jinzhu/gorm"
)

//********************************************************************************************************
/* *********************    Funcion que agrega un Telefono a un representante legal     ********************* */
func (p *TelefonoContrato) AgregaTelContrato(tx *gorm.DB) error {
	return tx.Create(p).Error
}

//********************************************************************************************************
/* *********************    Funcion para consultar los telefonos del contrato     ********************* */
func ConsultarTelefonos(n_contrato string) (*[]TelefonoContrato, error) {
	var telefono []TelefonoContrato
	db := Conexion.GetDB()
	err := db.Where("contrato_numero_contrato = ?", n_contrato).Preload("Telefonos").Find(&telefono).Error
	if err != nil {
		return nil, err
	}
	return &telefono, nil
}

//********************************************************************************************************
/* *********************    Funcion para buscar un telefono de un contrato     ********************* */
func ConsultarUnTelefono(id string) (*TelefonoContrato, error) {
	var telefono TelefonoContrato
	db := Conexion.GetDB()
	err := db.Where("id = ?", id).Preload("Telefonos").First(&telefono).Error
	if err != nil {
		return nil, err
	}
	return &telefono, nil
}

//********************************************************************************************************
/* ****************    Funcion para actualizar la extension o el tipo de telefono de un contrato     **************** */
func (tel *TelefonoContrato) ActualizarTelefonoContrato(id string, tx *gorm.DB) error {
	return tx.Model(&tel).Where("id=?", id).Update(TelefonoContrato{TipoTelefono: tel.TipoTelefono,
		Extension: tel.Extension}).Error
}

//********************************************************************************************************
/* *********************    Funcion para eliminar un telefono de un contrato     ********************* */
func EliminarTelefonoContrato(id string, tx *gorm.DB) (error, *TelefonoContrato) {
	var item TelefonoContrato
	err := tx.Where("id = ?", id).First(&item).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro el telefono"), nil
		}
		return err, nil
	}
	return tx.Model(&TelefonoContrato{}).Delete(&item).Error, &item
}
