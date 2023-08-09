package DepartamentosModel

import (
	"appbiomedica/Conexion"
	"github.com/jinzhu/gorm"
)

func (dp *DepartamentoPadre) AddDepartamentoPadre(tx *gorm.DB) error {
	return tx.Create(dp).Error
}

func ConsultDepartmentUnit(id string) (error, *[]OrganigramaUnidad) {
	var data []OrganigramaUnidad
	db := Conexion.GetDB()
	err := db.
		Where("organigrama_unidad.id = ?", id).
		Preload("Departamento").
		Preload("Organigramas").
		Preload("Organigramas.Departamento").
		Preload("Organigramas.Organigramas").
		Preload("Organigramas.Organigramas.Departamento").
		Preload("Organigramas.Organigramas.Organigramas").
		Preload("Organigramas.Organigramas.Organigramas.Departamento"). 
		Preload("Organigramas.Organigramas.Organigramas.Organigramas").
		Preload("Organigramas.Organigramas.Organigramas.Organigramas.Departamento").
		Find(&data).Error 
	if err != nil {
		return err, nil
	}
	return nil, &data
}
func ConsultParentUnitId(unitId int) (error, *OrganigramaUnidad) {
	var data OrganigramaUnidad
	db := Conexion.GetDB()
	err := db.Where("unidad_med_id = ? AND nivel_jerarquico = ?", unitId, 1).
		First(&data).Error
	if err != nil {
		return err, nil
	}
	return nil, &data
}
func ConsultParentDelegationId(clvDele string) (error, *OrganigramaDelegacion) {
	var data OrganigramaDelegacion
	db := Conexion.GetDB()
	err := db.Where("delegaciones_clv_dele = ? AND nivel_jerarquico = ?", clvDele, 1).
		First(&data).Error
	if err != nil {
		return err, nil
	}
	return nil, &data
}
func ConsultDepartmentDelegation(id string) (error, *[]OrganigramaDelegacion) {
	var data []OrganigramaDelegacion
	db := Conexion.GetDB()
	err := db.Where("id = ?", id).Preload("Departamento").
		Preload("Organigramas").
		Preload("Organigramas.Departamento").
		Preload("Organigramas.Organigramas").
		Preload("Organigramas.Organigramas.Departamento").
		Preload("Organigramas.Organigramas.Organigramas").
		Preload("Organigramas.Organigramas.Organigramas.Departamento").
		Preload("Organigramas.Organigramas.Organigramas.Organigramas").
		Preload("Organigramas.Organigramas.Organigramas.Organigramas.Departamento").
		Find(&data).Error
	if err != nil {
		return err, nil
	}
	return nil, &data
}
