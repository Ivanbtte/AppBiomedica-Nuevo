package AreaModel

import (
	"appbiomedica/Conexion"
	"github.com/jinzhu/gorm"
)

func (a *Area) AddArea(tx *gorm.DB) error {
	return tx.Create(a).Error
}
func GetAreas() (error, *[]Area) {
	var areas []Area
	db := Conexion.GetDB()
	err := db.Order("orden_jerarquia").Find(&areas).Error
	if err != nil {
		return err, nil
	}
	return nil, &areas
}
