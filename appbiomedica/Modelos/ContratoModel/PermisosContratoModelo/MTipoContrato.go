package PermisosContratoModelo

import (
	"appbiomedica/Conexion"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

func (tipo *TipoContrato) AgregarTipoContrato(tx *gorm.DB) error {
	return tx.Create(tipo).Error
}

//********************************************************************************************************
/* *********************    Funciones de la estructura subtipo de contrato     ********************* */
/*****  Funcion para crear un nuevo subtipo  *****/

func (sub *SubTipo) AgregarSubTipo(tx *gorm.DB) error {
	return tx.Create(sub).Error
}

/*****  Funcion para consultar un subtipo de contrato si existe   *****/
func ConsultarSubTipo(tipo_contrato string) (*[]SubTipo, error) {
	var sub_tipo []SubTipo
	db := Conexion.GetDB()
	err := db.Where("tipo_contrato_id = ?", tipo_contrato).Find(&sub_tipo).Error
	fmt.Println(err)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("Este id no existe")
			fmt.Println("No se econtro nada")
		}
		return nil, err
	}
	return &sub_tipo, nil
}
