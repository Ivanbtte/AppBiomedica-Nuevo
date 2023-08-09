package PermisosContratoModelo

import (
	"appbiomedica/Conexion"
	"github.com/jinzhu/gorm"
)

func (contrato_puesto *TipoContratoPuesto) AgregarTipoContratoPuesto(tx *gorm.DB) error {
	return tx.Create(contrato_puesto).Error
}

//********************************************************************************************************
/* *********************    Funcion que regresa todos los permisos de contratos disponibles     ********************* */
func ConsultaPuestosPermisocontrato(puesto_id string) (error, *[]TipoContratoPuesto) {
	var datos []TipoContratoPuesto
	db := Conexion.GetDB()
	err := db.Where("catalogo_puesto_id=?", puesto_id).Preload("TipoContrato").Find(&datos).Error
	if err != nil {
		return err, nil
	}
	return nil, &datos
}
