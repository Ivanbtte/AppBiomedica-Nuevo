package PenasModelo

import (
	"appbiomedica/Conexion"
	"errors"
	"github.com/jinzhu/gorm"
)

func (pena *PenasConvencionales) AgregarPenaConv(tx *gorm.DB) error {
	return tx.Create(pena).Error
}

//********************************************************************************************************
/* ***********    Funcion para consultar una pena convencional buscandolo por numero de contrato     ************** */
func ConsultarPenaConvencional(numero_contrato string) (error, *PenasConvencionales) {
	var dato PenasConvencionales
	db := Conexion.GetDB()
	err := db.Where("contrato_numero_contrato = ?", numero_contrato).First(&dato).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro la pena"), nil
		}
		return err, nil
	}
	return nil, &dato
}
