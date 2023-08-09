package DepartamentosModel

import (
	"appbiomedica/Conexion"
	"github.com/jinzhu/gorm"
)

func (bup *BitacoraUsuarioPuesto) AddBitacoraUsuarioPuesto(tx *gorm.DB) error {
	return tx.Create(bup).Error
}

//Modelo de tipo de modificacion
func (tm *TipoModificacion) AddTipoModificacion(tx *gorm.DB) error {
	return tx.Create(tm).Error
}

func ConsultTipoModificacion(notId string) (*[]TipoModificacion, error) {
	var results []TipoModificacion
	db := Conexion.GetDB()
	err := db.Not("id = ?", notId).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return &results, nil
}
