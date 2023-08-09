package RoleModel

import (
	"appbiomedica/Conexion"
	"github.com/jinzhu/gorm"
)

func (rm *RoleModule) AddRoleModule(tx *gorm.DB) error {
	return tx.Create(rm).Error
}
func ConsultNameModule(roleId []string) (error, *[]RoleModule) {
	data := make([]RoleModule, 0)
	db := Conexion.GetDB()
	err := db.Where("role_id IN (?)", roleId).Preload("Modulo").Find(&data).Error
	if err != nil {
		return err, nil
	}
	return nil, &data
}
