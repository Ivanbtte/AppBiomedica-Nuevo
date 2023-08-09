package ModuloModel

import (
	"appbiomedica/Conexion"
	"github.com/jinzhu/gorm"
)

func (m *Modulo) AddModulo(tx *gorm.DB) error {
	return tx.Create(m).Error
}
func (a *Action) AddAction(tx *gorm.DB) error {
	return tx.Create(a).Error
}

func ConsultActionsByModule(moduleId string) (error, *[]Action) {
	var data []Action
	db := Conexion.GetDB()
	err := db.Where("modulo_id = ? ", moduleId).Preload("Modulo").Find(&data).Error
	if err != nil {
		return err, nil
	}
	return nil, &data
}
