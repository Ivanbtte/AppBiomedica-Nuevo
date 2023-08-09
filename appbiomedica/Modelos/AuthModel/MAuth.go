package AuthModel

import "C"
import (
	"appbiomedica/Conexion"
	"appbiomedica/Modulos/Estructuras"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

//********************************************************************************************************
/* *********************    Función para el login     ********************* */

func (c *Credenciales) Login() (error, *Credenciales) {
	var credenciales Credenciales
	db := Conexion.GetDB()
	err := db.Where("correo=?", c.Correo).Last(&credenciales).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("este correo no ha sido registrado"), nil
		}
		return err, nil
	}
	fmt.Print(credenciales.Password, "------------", Estructuras.GetMD5Hash(c.Password))
	if credenciales.Password != Estructuras.GetMD5Hash(c.Password) {
		return errors.New("contraseña incorrecta, inténtalo de nuevo"), nil
	}
	return nil, &credenciales
}

/* *********************    Función que agrega un nuevo registro de credenciales     ********************* */

func (c *Credenciales) AgregaCredencial(tx *gorm.DB) error {
	return tx.Create(c).Error
}

/* *********************    Función que actualiza la contraseña de una credencial     ********************* */

func (c *Credenciales) ActualizarCredencial(id string, tx *gorm.DB) error {
	return tx.Model(&c).Where("usuario_id=?", id).Update(Credenciales{Password: Estructuras.GetMD5Hash(c.Password)}).
		Error
}

/* *********************    Función que valida una credencial     ********************* */

func ValidaCredencial(id string, estado bool, tx *gorm.DB) error {
	return tx.Model(&Credenciales{}).Where("usuario_id=?", id).Updates(map[string]interface{}{"estado": estado}).Error
}

/* *********************    Función para eliminar credenciales   ********************* */

func EliminarCredenciales(id string, tx *gorm.DB) (error, *Credenciales) {
	var item Credenciales
	err := tx.Where("usuario_id = ?", id).Find(&item).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("no se encontró el id"), nil
		}
		return err, nil
	}
	return tx.Model(&Credenciales{}).Where("usuario_id=?", id).Delete(&item).Error, &item
}
func GetOneCredentialByPersonId(personId string) (error, *Credenciales) {
	var credenciales Credenciales
	db := Conexion.GetDB()
	err := db.Joins("inner join usuario u on u.id = credenciales.usuario_id").
		Joins("inner join persona p on p.id = u.persona_id").
		Where("p.id = ?", personId).Preload("Usuario").Preload("Usuario.Persona").
		Find(&credenciales).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("no se encontró el id"), nil
		}
		return err, nil
	}
	return nil, &credenciales
}
func (c Credenciales) UpdateCredentialInactive(id string, tx *gorm.DB) error {
	return tx.Model(&c).Where("id=?", id).Update(Credenciales{Correo: c.Correo,
		Password: c.Password, Estado: c.Estado}).Error
}
