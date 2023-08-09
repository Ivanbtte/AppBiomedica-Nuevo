package TrasladoPacientesModel

import "github.com/jinzhu/gorm"

/*****  Funcion para agregar registro de actualizacion de la bitacora  *****/
func (ruf *RecordUpdateFirm) AddUpdateFirm(tx *gorm.DB) error {
	return tx.Create(ruf).Error
}
