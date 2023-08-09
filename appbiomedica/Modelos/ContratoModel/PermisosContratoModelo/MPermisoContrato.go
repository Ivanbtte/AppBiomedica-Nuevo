package PermisosContratoModelo

import (
	"appbiomedica/Conexion"
	"github.com/jinzhu/gorm"
)

func (permiso *PermisoContrato) AgregarPermisoContrato(tx *gorm.DB) error {
	return tx.Create(permiso).Error
}


//********************************************************************************************************
/* *********************    Funcion que regresa      ********************* */
func ConsultaPuestosPermiso() (error, *[]PermisoContrato) {
	var datos []PermisoContrato
	db := Conexion.GetDB()
	err := db.Find(&datos).Error
	if err != nil {
		return err, nil
	}
	return nil, &datos
}