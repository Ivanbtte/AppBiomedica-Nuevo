package MenuModel

import (
	"appbiomedica/Conexion"
	"github.com/jinzhu/gorm"
)

func (m *Menu) AddMenu(tx *gorm.DB) error {
	return tx.Create(m).Error
}

func ConsultModuleByUSer() (error, *[]Menu) {
	var data []Menu
	db := Conexion.GetDB()
	err := db.Where("menu_parent is null").Preload("Children").Preload("Modulo").Find(&data).Error
	if err != nil {
		return err, nil
	}
	return nil, &data
}

func ConsultModuleChildByParent(parentModule string) (error, *[]Menu) {
	var data []Menu
	db := Conexion.GetDB()
	err := db.Where("menu_parent = ? and not orden = ?", parentModule,
		1).Preload("Modulo").Order("orden asc").Find(&data).Error
	if err != nil {
		return err, nil
	} 
	return nil, &data                                                                          
}

func ConsultMenuByRole(roleId []string) (error, *[]Menu) {
	data := make([]Menu, 0)
	db := Conexion.GetDB()
	err := db.Joins("inner join modulo m on m.id = menu.modulo_id").
		Joins("inner join role_module rm on m.id = rm.modulo_id").
		Where("rm.role_id IN (?)", roleId).
		Where("menu_parent is null").
		Preload("Children", func(db *gorm.DB) *gorm.DB {
			return db.Joins("inner join modulo m on m.id = menu.modulo_id").
				Joins("inner join role_module rm on m.id = rm.modulo_id").
				Where("rm.role_id IN (?)", roleId).Order("Orden asc")
		}).Order("Orden asc").Find(&data).Error
	if err != nil {
		return err, nil
	}
	return nil, &data
}
