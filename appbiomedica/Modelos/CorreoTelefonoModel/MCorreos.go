package CorreoTelefonoModel

import (
	"appbiomedica/Conexion"
	"github.com/jinzhu/gorm"
)

//********************************************************************************************************
/* *********************    Funcion para agregar un correo       ********************* */
func (p *Correos) AgregaCorreo(tx *gorm.DB) error {
	return tx.Create(p).Error
}

//********************************************************************************************************
/* *********************    Funcion que busca un correo     ********************* */
func BuscarCorreo(mail string) (bool, *Correos) {
	var correo Correos
	db := Conexion.GetDB()
	err := db.Where("correo = ?", mail).First(&correo).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		return false, nil
	}
	// repetido error
	return true, &correo
}

//********************************************************************************************************
/* *****************    Funcion para actualizar un correo electronico     ***************** */
func (cc *Correos) ActualizarCorreo(id string, tx *gorm.DB) error {
	return tx.Model(&cc).Where("id=?", id).Update(Correos{Correo: cc.Correo}).Error
}
