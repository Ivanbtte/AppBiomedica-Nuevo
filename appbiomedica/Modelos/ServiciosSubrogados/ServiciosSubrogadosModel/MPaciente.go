package serviciossubrogadosmodel

import (
	"github.com/jinzhu/gorm"
)

func (p *Paciente) AddPaciente(tx *gorm.DB) error {
	return tx.Create(p).Error
}
