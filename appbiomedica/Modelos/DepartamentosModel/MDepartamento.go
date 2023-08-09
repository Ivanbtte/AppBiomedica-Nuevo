package DepartamentosModel

import (
	"appbiomedica/Conexion"
	"github.com/jinzhu/gorm"
)

func (d *Departamento) AgregarDepartamento(tx *gorm.DB) error {
	return tx.Create(d).Error
}

func (dt *DepartamentoTipo) AgregarDepartamentoTipo(tx *gorm.DB) error {
	return tx.Create(dt).Error
}

func ConsultaDepartamento(id string) (error, *Departamento) {
	var datos Departamento
	db := Conexion.GetDB()
	err := db.Where("id = ?", id).Preload("DepartamentoTipo").First(&datos).Error
	if err != nil {
		return err, nil
	}
	return nil, &datos
}
