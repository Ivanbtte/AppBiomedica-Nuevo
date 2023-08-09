package DepartamentosModel

import (
	"github.com/jinzhu/gorm"
)

func (ae *ActionExceptionByUser) AddExceptionAction(tx *gorm.DB) error {
	return tx.Create(ae).Error
}

func CambiarEstadoActionException(id string, status bool, tx *gorm.DB) error {
	var actionException ActionExceptionByUser
	err := tx.Where("usuario_puesto_id = ?", id).First(&actionException).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil
		}
		tx.Rollback()
		return err
	}
	return tx.Model(&ActionExceptionByUser{}).Where("usuario_puesto_id=?", id).Updates(map[string]interface{}{"status": status}).Error
}