package TrasladoPacientesModel

import "github.com/jinzhu/gorm"

/*****  Funcion para agregar un nuevo tipo de pasajero  *****/
func (tp *TipoPasajero) AgregarTipoPasajero(tx *gorm.DB) error {
	return tx.Create(tp).Error
}
