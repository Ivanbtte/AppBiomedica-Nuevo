package PenasModelo

import (
	"appbiomedica/Conexion"
	"github.com/jinzhu/gorm"
)

//********************************************************************************************************
/* *********************    Funcion para agregar una pena deductiva    *********** */
func (pena *PenasDeductivas) AgregarPenaDed(tx *gorm.DB) error {
	return tx.Create(pena).Error
}

//********************************************************************************************************
/* **************    Funcion para consultar una pena deductiva buscandolo por numero de contrato     ************** */
func ConsultarPenaDeductiva(numero_contrato string) (error, *[]PenasDeductivas, bool) {
	var datos []PenasDeductivas
	db := Conexion.GetDB()
	err := db.Where("contrato_numero_contrato = ?", numero_contrato).Find(&datos).Error
	if err != nil {
		return err, nil, false
	}
	return nil, &datos, true
}
