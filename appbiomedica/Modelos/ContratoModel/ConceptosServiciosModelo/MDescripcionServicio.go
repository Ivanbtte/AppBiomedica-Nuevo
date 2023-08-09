package ConceptosServiciosModelo

import (
	"appbiomedica/Conexion"
	"github.com/jinzhu/gorm"
)

func (ds *DescripcionServicio) AgregarDescServicios(tx *gorm.DB) error {
	return tx.Create(ds).Error
}
//Funcion para buscar servicio por nombre y especialidad
func BuscarDescServicio(servicio,especialidad string) (*DescripcionServicio, error) {
	var resultado DescripcionServicio
	db := Conexion.GetDB()
	err := db.Joins("inner join especialidades on descripcion_servicio.especialidades_id = especialidades.id").
	Where("trim(lower(descripcion)) = ?", servicio).First(&resultado).Error
	if err != nil {
		return nil, err
	}
	return &resultado, nil
}