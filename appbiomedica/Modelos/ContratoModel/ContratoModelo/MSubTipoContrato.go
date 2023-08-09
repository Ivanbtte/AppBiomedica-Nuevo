package ContratoModelo

import (
	"appbiomedica/Conexion"
	"errors"
	"github.com/jinzhu/gorm"
)

/*****  Funcion para crear una nueva relacion de un subtipo con un contrato*****/

func (sub_contrato *SubTipoContrato) AgregarSubTipoContrato(tx *gorm.DB) error {
	return tx.Create(sub_contrato).Error
}

/*****  Funcion para consultar una relacion de un subtipo con un contrato   *****/
func ConsultarSubTipoContrato(numero_contrato string) (*SubTipoContrato, error) {
	var sub_tipo SubTipoContrato
	db := Conexion.GetDB()
	err := db.Where("contrato_numero_contrato = ?", numero_contrato).First(&sub_tipo).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("no se ha registrado ningun tipo se servicio")
		}
		return nil, err
	}
	return &sub_tipo, nil
}

//Funcion para consultar los contratos por sub_tipo

func ConsultarContratoSubTipo(subTipo string) (*[]SubTipoContrato, error) {
	var datos []SubTipoContrato
	db := Conexion.GetDB()
	err := db.Joins("inner join contrato c on sub_tipo_contrato.contrato_numero_contrato = c.numero_contrato").
		Where("sub_tipo_id = ? and c.estado =?", subTipo,true).Preload("Contrato").Find(&datos).Error
	if err != nil {
		return nil, err
	}
	return &datos, nil
}
