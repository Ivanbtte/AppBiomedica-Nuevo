package ContactoModelo

import (
	"appbiomedica/Conexion"
	"errors"
	"github.com/jinzhu/gorm"
)

//********************************************************************************************************
/* *********************    Funcion que agrega un Correo a un representante legal     ********************* */
func (p *CorreosContrato) AgregaCorreoContrato(tx *gorm.DB) error {
	return tx.Create(p).Error
}

//********************************************************************************************************
/* *********************    Funcion para consultar los correos del contrato     ********************* */
func ConsultarCorreos(n_contrato string) (*[]CorreosContrato, error) {
	var correos []CorreosContrato
	db := Conexion.GetDB()
	err := db.Where("contrato_numero_contrato = ?", n_contrato).Preload("Correos").Find(&correos).Error
	if err != nil {
		return nil, err
	}
	return &correos, nil
}

//********************************************************************************************************
/* *********************    Funcion para buscar un correo de un contrato     ********************* */
func ConsultarUnCorreo(id string) (*CorreosContrato, error) {
	var correo CorreosContrato
	db := Conexion.GetDB()
	err := db.Where("id = ?", id).Preload("Correos").First(&correo).Error
	if err != nil {
		return nil, err
	}
	return &correo, nil
}

//********************************************************************************************************
/* *****************    Funcion para actualizar el tipo de correo electronico de un contrato     ***************** */
func (cc *CorreosContrato) ActualizarTipoCorreo(id string, tx *gorm.DB) error {
	return tx.Model(&cc).Where("id=?", id).
		Update(CorreosContrato{TipoCorreo: cc.TipoCorreo}).Error
}

//********************************************************************************************************
/* *********************    Funcion para eliminar un correo de un contrato     ********************* */
func EliminarCorreoContrato(id string, tx *gorm.DB) (error, *CorreosContrato) {
	var item CorreosContrato
	err := tx.Where("id = ?", id).First(&item).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro el correo"), nil
		}
		return err, nil
	}
	return tx.Model(&CorreosContrato{}).Delete(&item).Error, &item
}
