package DepartamentosModel

import (
	"appbiomedica/Conexion"
	"github.com/jinzhu/gorm"
)

func (po *PuestoOrganigrama) AddPositionOrganization(tx *gorm.DB) error {
	return tx.Create(po).Error
}

func ConsultPositionOrganization(departmentUnitId, departmentDelegationId string) (error, *[]PuestoOrganigrama) {
	var data []PuestoOrganigrama
	db := Conexion.GetDB()
	if departmentUnitId != "" {
		db = db.Where("organigrama_unidad_id = ? ", departmentUnitId)
	}
	if departmentDelegationId != "" {
		db = db.Where("organigrama_delegacion_id = ?", departmentDelegationId)
	}
	err := db.Preload("CatalogoPuesto").
		Find(&data).Error
	if err != nil {
		return err, nil
	}
	return nil, &data
}

func CountPositionsBusy(positionId string) (error, int) {
	var total int
	db := Conexion.GetDB()
	err := db.Model(&UsuarioPuesto{}).Where("puesto_organigrama_id = ? and estado = ?", positionId, true).Count(&total).Error
	if err != nil {
		return err, 0
	}
	return nil, total
}
