package UsuarioModel

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modulos/Estructuras"
	"errors"
	"github.com/jinzhu/gorm"
)

func ConsultaUnUsuario(Id string) (*Usuario, error) {
	var item Usuario
	db := Conexion.GetDB()
	err := db.Where("id=?", Id).Find(&item).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("no se encontró el id")
		}
		return nil, err
	}
	return &item, nil
}

//-----------------------Consultar todos los usuarios

func ConsultaUsuario(estado bool) (error, *[]Usuario) {
	var datos []Usuario
	db := Conexion.GetDB()
	err := db.Where("estado = ?", estado).Preload("Persona").Find(&datos).Error
	if err != nil {
		return err, nil
	}
	return nil, &datos
}

/* *********************    Función para agregar un usuario     ********************* */

func (u *Usuario) AgregaUsuario(tx *gorm.DB) error {
	u.Contraseña = Estructuras.GetMD5Hash(u.Contraseña)
	return tx.Create(u).Error
}

//-----------------------Actualizar un usuario

func (u *Usuario) ActualizarUsuario(id string, tx *gorm.DB) error {
	var usuarios Usuario
	err := tx.Where("id = ?", id).First(&usuarios).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("no se encontró el id")
		}
		return err
	}
	return tx.Model(&u).Where("id=?", id).Update(Usuario{Contraseña: Estructuras.GetMD5Hash(u.Contraseña)}).Error
}

/* ********************* Validar usuario ********************* */

func ValidarUsuario(id string, estado bool, tx *gorm.DB) error {
	var usuarios Usuario
	err := tx.Where("id = ?", id).Last(&usuarios).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("no se encontró el id")
		}
		return err
	}
	return tx.Model(&Usuario{}).Where("id=?", id).Updates(map[string]interface{}{"estado": estado}).Error
}

/* *********************  Función para eliminar usuarios   ********************* */

func EliminarUsuario(id string, tx *gorm.DB) (error, *Usuario) {
	var item Usuario
	err := tx.Where("id = ?", id).Find(&item).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("no se encontró el id"), nil
		}
		return err, nil
	}
	return tx.Model(&Usuario{}).Where("id=?", id).Delete(&item).Error, &item
}
func CheckResetPassword(email, dateOfBirth string) (error, bool, *Usuario) {
	var data Usuario
	db := Conexion.GetDB()
	err := db.Joins("inner join persona p on p.id = usuario.persona_id").Where("p.correo = ? and p.fecha_nac = ?",
		email, dateOfBirth).Preload("Persona").Find(&data).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("no se encontró registro con los datos proporcionados"), false, nil
		}
		return err, false, nil
	}
	return nil, true, &data
}

func ResetUserPassword(id, newPassword string, tx *gorm.DB) error {
	var usuarios Usuario
	err := tx.Where("id = ?", id).First(&usuarios).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("no se encontró el usuario")
		}
		return err
	}
	return tx.Model(&usuarios).Where("id=?", id).Update(Usuario{Contraseña: Estructuras.GetMD5Hash(newPassword)}).Error
}

func (u Usuario) UpdateUserInactive(id string, tx *gorm.DB) error {
	return tx.Model(&u).Where("id=?", id).Update(Usuario{Correo: u.Correo,
		Contraseña: u.Contraseña, UsuarioTipoId: u.UsuarioTipoId, Estado: u.Estado}).Error
}

func (u Usuario) UpdateUserType(id string, tx *gorm.DB) error {
	var usuarios Usuario
	err := tx.Where("id = ?", id).First(&usuarios).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("no se encontró el id")
		}
		return err
	}
	return tx.Model(&u).Where("id=?", id).Update(Usuario{UsuarioTipoId: u.UsuarioTipoId}).Error
}