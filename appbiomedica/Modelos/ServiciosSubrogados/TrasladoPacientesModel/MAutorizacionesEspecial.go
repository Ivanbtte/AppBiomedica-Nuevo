package TrasladoPacientesModel

import "github.com/jinzhu/gorm"

func (ae *AutorizacionesEspeciales) AgregarAutorizacionEsp(tx *gorm.DB) error {
	return tx.Create(ae).Error
}
