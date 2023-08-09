package DepartamentosModel

import (
	"appbiomedica/Conexion"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

func (p *UsuarioPuesto) AgregaUsuarioPuesto(tx *gorm.DB) error {
	return tx.Create(p).Error
}

func ConsultaUnUsuarioPuesto(Id string) (*UsuarioPuesto, error) {
	var item UsuarioPuesto
	db := Conexion.GetDB()
	err := db.Where("usuario_id = ?", Id).
		Preload("Usuario").
		Preload("Usuario.Persona").
		Preload("PuestoOrganigrama").
		Preload("PuestoOrganigrama.CatalogoPuesto").
		Preload("PuestoOrganigrama.OrganigramaUnidad").
		Preload("PuestoOrganigrama.OrganigramaUnidad.Departamento").
		Preload("PuestoOrganigrama.OrganigramaUnidad.UnidadMed").
		Preload("PuestoOrganigrama.OrganigramaUnidad.UnidadMed.Delegacion").
		Preload("PuestoOrganigrama.OrganigramaDelegacion").
		Preload("PuestoOrganigrama.OrganigramaDelegacion.Departamento").
		Preload("PuestoOrganigrama.OrganigramaDelegacion.Delegaciones").
		Last(&item).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, err
		}
		return nil, err
	}
	return &item, nil
}

func ConsultUserByDepartment(departmentId, userId string) (error, *[]UsuarioPuesto) {
	var results []UsuarioPuesto
	db := Conexion.GetDB()
	err := db.
		Joins("inner join puesto_organigrama po on po.id = usuario_puesto.puesto_organigrama_id").
		Joins("inner join organigrama_unidad ou on ou.id = po.organigrama_unidad_id").
		Joins("inner join usuario u on u.id = usuario_id").
		Joins("inner join persona p on p.id = u.persona_id").
		Where("(ou.organigrama_parent = ? or ou.id = ?) and usuario_puesto.estado = ?", departmentId, departmentId,
			true).
		Not("usuario_id = ?", userId).Preload("Usuario").
		Preload("Usuario.Persona").
		Preload("PuestoOrganigrama").
		Preload("PuestoOrganigrama.CatalogoPuesto").
		Preload("PuestoOrganigrama.OrganigramaUnidad.Departamento").
		Order("p.nombre asc").
		//	Group("puesto_organigrama_id").
		Find(&results).Error
	if err != nil {
		return err, nil
	}
	return nil, &results
}

func ConsultUsersNotValid(status, valid bool, unitMedId int, clvDele, department string) (error, *[]UsuarioPuesto) {
	fmt.Println(status, valid)
	var results []UsuarioPuesto
	db := Conexion.GetDB()
	db = db.Joins("inner join usuario u on u.id = usuario_puesto.usuario_id").
		Joins("inner join puesto_organigrama po on po.id = usuario_puesto.puesto_organigrama_id")
	if unitMedId != 0 {
		db = db.Joins("inner join organigrama_unidad ou on ou.id = po.organigrama_unidad_id").Where("ou."+
			"unidad_med_id = ? and (ou.id = ? or ou.organigrama_parent = ?)", unitMedId, department, department).
			Preload("PuestoOrganigrama.OrganigramaUnidad.Departamento").
			Preload("PuestoOrganigrama.OrganigramaUnidad.UnidadMed").
			Preload("PuestoOrganigrama.OrganigramaUnidad.UnidadMed.Delegacion")
	}
	if clvDele != "" {
		db = db.Joins("inner join organigrama_delegacion od on od.id = po.organigrama_delegacion_id").Where("od."+
			"delegaciones_clv_dele = ? and (od.id = ? or od.organigrama_parent = ?)", clvDele, department, department).
			Preload("PuestoOrganigrama.OrganigramaDelegacion.Departamento").
			Preload("PuestoOrganigrama.OrganigramaDelegacion.Delegaciones")
	}
	err := db.Where("usuario_puesto.estado = ? and valid = ?", status, valid).
		Preload("Usuario").
		Preload("Usuario.Persona").
		Preload("PuestoOrganigrama").
		Preload("PuestoOrganigrama.CatalogoPuesto").
		Find(&results).Error
	if err != nil {
		return err, nil
	}
	return nil, &results
}
func CheckNotRepeatUsuarioPuesto(usuarioId string) (bool, *UsuarioPuesto, error) {
	data := UsuarioPuesto{}
	db := Conexion.GetDB()
	err := db.Where("usuario_id = ? and estado = ?", usuarioId, true).First(&data).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil, nil
		}
		return false, &data, err
	}
	return true, &data, nil
}

func CambiarValidacionUsuarioPuesto(idUsuarioPuesto string, status bool, tx *gorm.DB) error {
	var usuarioPuesto UsuarioPuesto
	err := tx.Where("id = ?", idUsuarioPuesto).First(&usuarioPuesto).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("no se encontró el id")
		}
		return err
	}
	return tx.Model(&UsuarioPuesto{}).Where("id=?", idUsuarioPuesto).Updates(map[string]interface{}{"valid": status}).Error
}

func CambiarEstadoUsuarioPuesto(idUsuarioPuesto string, estado bool, tx *gorm.DB) error {
	var usuarioPuesto UsuarioPuesto
	err := tx.Where("id = ?", idUsuarioPuesto).First(&usuarioPuesto).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("no se encontró el id")
		}
		tx.Rollback()
		return err
	}
	return tx.Model(&UsuarioPuesto{}).Where("id=?", idUsuarioPuesto).Updates(map[string]interface{}{"estado": estado}).Error
}
func ConsultarUsuarioPuesto(puestoOrgId string)(*UsuarioPuesto, error){
	data := UsuarioPuesto{}
	db := Conexion.GetDB()
	err := db.Where("puesto_organigrama_id = ? and estado = ?", puestoOrgId, true).
	Preload("Usuario").
	Preload("Usuario.Persona").
	First(&data).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, err
		}
		return  nil, err
	}
	return  &data, nil
}