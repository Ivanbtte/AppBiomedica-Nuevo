package TrasladoPacientesModel

import "github.com/jinzhu/gorm"

/*****  Funcion para agregar registro de activar o inhabilitar firma de la bitacora  *****/
func (rdf *RecordDisableFirm) AddUpdateFirm(tx *gorm.DB) error {
	return tx.Create(rdf).Error
}