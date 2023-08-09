package DepartamentosModel

import (
	"appbiomedica/Conexion"
	"github.com/jinzhu/gorm"
)

func (puesto *CatalogoPuesto) AgregarPuesto(tx *gorm.DB) error {
	return tx.Create(puesto).Error
}

/* ********************* Funcion que regresa todos los puestos disponibles ********************* */

func ConsultaPuestos() (*[]CatalogoPuesto, error) {
	var datos []CatalogoPuesto
	db := Conexion.GetDB()
	err := db.Find(&datos).Error
	if err != nil {
		return nil, err
	}
	return &datos, nil
}
