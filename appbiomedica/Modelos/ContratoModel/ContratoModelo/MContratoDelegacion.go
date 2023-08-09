package ContratoModelo

import (
	"appbiomedica/Conexion"
	"errors"
	"github.com/jinzhu/gorm"
)

//********************************************************************************************************
/* *********************    Funcion para agregar una delegacion a  un contrato       ********************* */
func (p *Contrato_Delegaciones) AgregarDelegContrato(tx *gorm.DB) error {
	return tx.Create(p).Error
}

//********************************************************************************************************
/* *****************  Funcion que regresa la lista de delegaciones por contrato  ********************* */
func ConsultaDelegContrato(num_contrato string) (error, *[]Contrato_Delegaciones) {
	var datos []Contrato_Delegaciones
	db := Conexion.GetDB()
	err := db.Where("contrato_numero_contrato = ?", num_contrato).Preload("Delegaciones").Find(&datos).Error
	if err != nil {
		return err, nil
	}
	return nil, &datos

}

//********************************************************************************************************
/* *********************    Funcion para eliminar una delegacion de un contrato     ********************* */
func EliminarDeleContr(num_contrato, clv_dele string, tx *gorm.DB) (error, *Contrato_Delegaciones) {
	var item Contrato_Delegaciones
	err := tx.Where("contrato_numero_contrato = ? AND delegaciones_clv_dele = ?", num_contrato,
		clv_dele).Find(&item).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro esta delegacion del contrato"), nil
		}
		return err, nil
	}
	return tx.Model(&Contrato_Delegaciones{}).Where("contrato_numero_contrato = ? AND delegaciones_clv_dele = ?",
		num_contrato, clv_dele).Delete(&item).Error, &item
}
