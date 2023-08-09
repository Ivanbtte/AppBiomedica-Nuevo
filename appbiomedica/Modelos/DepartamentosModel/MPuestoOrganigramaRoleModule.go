package DepartamentosModel

import (
	"appbiomedica/Conexion"
	"github.com/jinzhu/gorm"
)

func (po *PuestoOrganigramaRole) AddPuestoOrganigramaRole(tx *gorm.DB) error {
	return tx.Create(po).Error
}
func ConsultPuestoOrganigramaRoleByPuesto(userPositionId string) (error, []*PuestoOrganigramaRole) {
	datos := make([]*PuestoOrganigramaRole, 0)
	db := Conexion.GetDB()
	err := db.Joins("inner join usuario_puesto up on puesto_organigrama_role.puesto_organigrama_id = up.puesto_organigrama_id").
		Where("up.id = ? and up.estado = ? and up.valid = ?", userPositionId, true, true).Find(&datos).Error
	if err != nil {
		return err, nil
	}
	return nil, datos
}
func CheckPermitRoleByOrganigrama(idRole, idPosition string) bool {
	var results PuestoOrganigramaRole
	db := Conexion.GetDB()
	err := db.Where("puesto_organigrama_id = ? AND role_id = ? AND status = ?", idPosition, idRole, true).First(
		&results).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false
		}
		return false
	}
	return true
}
func ConsultarPuestoRole(unidadId int, roleId string) (*PuestoOrganigramaRole, bool, error) {
	var item PuestoOrganigramaRole
	db := Conexion.GetDB()
	err := db.Joins("inner join puesto_organigrama po on puesto_organigrama_role.puesto_organigrama_id = po.id").
		Joins("inner join organigrama_unidad ou on po.organigrama_unidad_id = ou.id").
		Where("ou.unidad_med_id = ? and role_id = ?", unidadId, roleId).
		Preload("PuestoOrganigrama").
		Preload("PuestoOrganigrama.CatalogoPuesto").
		Preload("PuestoOrganigrama.OrganigramaUnidad.Departamento").
		Find(&item).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, true, nil
		}
		return nil, false, err
	}
	return &item, true, nil
}
