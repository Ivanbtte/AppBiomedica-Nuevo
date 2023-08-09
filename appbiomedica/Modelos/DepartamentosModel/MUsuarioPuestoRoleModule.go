package DepartamentosModel

import (
	"appbiomedica/Conexion"
	"errors"

	"github.com/jinzhu/gorm"
)

func (up *UsuarioPuestoRole) AddUsuarioPuestoRoleModule(tx *gorm.DB) error {
	return tx.Create(up).Error
}
func ConsultUsuarioPuestoRole(usuarioPuestoId string) (error, []*UsuarioPuestoRole) {
	datos := make([]*UsuarioPuestoRole, 0)
	db := Conexion.GetDB()
	err := db.Joins("inner join usuario_puesto up on up.id = usuario_puesto_role.usuario_puesto_id").
		Where("up.id = ? and up.estado =? and status = ?", usuarioPuestoId, true, true).Find(&datos).Error
	if err != nil {
		return err, nil
	}
	return nil, datos
}

func CheckPermitByRole(idRole, idUserPosition string, status bool) bool {
	var results UsuarioPuestoRole
	db := Conexion.GetDB()
	err := db.Where("usuario_puesto_id = ? AND role_id = ? AND status = ?", idUserPosition, idRole, status).First(
		&results).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false
		}
		return false
	}
	return true
}
func CambiarEstadoUsuarioRole(id string, roleId string, status bool, tx *gorm.DB) error {
	var usuarioRole UsuarioPuestoRole
	err := tx.Where("usuario_puesto_id = ?", id).First(&usuarioRole).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("no se encontró el id")
		}
		return err
	}
	return tx.Model(&UsuarioPuestoRole{}).Where("usuario_puesto_id=? and role_id =? and status =?", id, roleId, true).Updates(map[string]interface{}{"status": status}).Error
}

func ValidarUsuarioRole(usuarioPuestoId, roleId string, status bool, tx *gorm.DB) error {
	var usuarioRole UsuarioPuestoRole
	err := tx.Where("usuario_puesto_id = ? AND role_id = ?", usuarioPuestoId, roleId).First(&usuarioRole).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("no se encontró el id")
		}
		return err
	}
	return tx.Model(&UsuarioPuestoRole{}).Where("usuario_puesto_id =? AND role_id =? AND status =?", usuarioPuestoId, roleId, false).Updates(map[string]interface{}{"status": status}).Error
}
