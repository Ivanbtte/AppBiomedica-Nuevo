package TrasladoPacientesModel

import (
	"appbiomedica/Conexion"

	"github.com/jinzhu/gorm"
)

func (ppu *PermisosPorUnidad) AgregarPermisoUnidad(tx *gorm.DB) error {
	return tx.Create(ppu).Error
}

func BuscarPermisosPorUnidad(unidadId int, moduloId string) (bool, error) {
	var item PermisosPorUnidad
	db := Conexion.GetDB()
	err := db.Joins("inner join action a on permisos_por_unidad.action_id = a.id").
		Where("unidad_med_id = ? and a.modulo_id = ? and estado = ?", unidadId, moduloId, true).Last(&item).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
