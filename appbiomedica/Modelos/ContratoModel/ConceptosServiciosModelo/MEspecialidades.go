package ConceptosServiciosModelo

import (
	"appbiomedica/Conexion"
	"github.com/jinzhu/gorm"
)

func (e *Especialidades) AgregarEspecialidad(tx *gorm.DB) error {
	return tx.Create(e).Error
}
func BuscarEspecialidad(especialidad string) (*Especialidades, error) {
	var resultado Especialidades
	db := Conexion.GetDB()
	err := db.Where("trim(lower(nombre)) = ?", especialidad).First(&resultado).Error
	if err != nil {
		return nil, err
	}
	return &resultado, nil
}
