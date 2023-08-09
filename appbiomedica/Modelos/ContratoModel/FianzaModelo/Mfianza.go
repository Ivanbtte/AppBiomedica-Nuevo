package FianzaModelo

import (
	"appbiomedica/Conexion"
	"errors"
	"github.com/jinzhu/gorm"
)

func (fianza *Fianza) Agregarfianza(tx *gorm.DB) error {
	return tx.Create(fianza).Error
}

//********************************************************************************************************
/* ***********    Funcion para consultar una fianza buscandolo por numero de contrato     ************** */
func ConsultarFianza(numero_contrato string) (error, *Fianza) {
	var dato Fianza
	db := Conexion.GetDB()
	err := db.Where("contrato_numero_contrato = ?", numero_contrato).First(&dato).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro la fianza"), nil
		}
		return err, nil
	}
	return nil, &dato
}