package PersonaModel

import (
	"appbiomedica/Conexion"
	"errors"
	"github.com/jinzhu/gorm"
)

//********************************************************************************************************
//-----------------------Consultar todos las personas

func ConsultaPersonas(numPagina, numRegistros int) (error, *[]Persona, int) {
	var datos []Persona
	var total int
	db := Conexion.GetDB()
	err := db.Find(&[]Persona{}).Count(&total).Error
	if err != nil {
		return err, nil, total
	}
	err = db.Offset((numPagina - 1) * numRegistros).Limit(numRegistros).Find(&datos).Error
	if err != nil {
		return err, nil, total
	}
	return nil, &datos, total

}

//********************************************************************************************************
//-----------------------Consulta una sola persona

func ConsultaUnaPersona(ID string) (error, *Persona) {
	var datos Persona
	db := Conexion.GetDB()
	err := db.Where("id=?", ID).Find(&datos).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("Esta persona no esta registrada"), nil
		}
		return err, nil
	}
	return nil, &datos

}

//********************************************************************************************************
//-----------------------Verificar que el correo no exista en la base de datos

func VerificaCorreoPersona(mail string) bool {
	db := Conexion.GetDB()
	tx := db.Begin()
	err := tx.Where("correo = ? and estado = ?", mail, true).Last(&Persona{}).Error
	if err != nil {
		tx.Rollback()
		return false
	}
	return true
}

//********************************************************************************************************
/* *********************    Funcion para agregar una persona     ********************* */

func (p *Persona) AgregaPersona(tx *gorm.DB) error {
	return tx.Create(p).Error
}

//********************************************************************************************************
//-----------------------Actualizar una Persona

func (p Persona) ActualizarPersona(id string, tx *gorm.DB) error {
	return tx.Model(&p).Where("id=?", id).Update(Persona{Nombre: p.Nombre, ApellidoPat: p.ApellidoPat,
		ApellidoMat: p.ApellidoMat, FechaNac: p.FechaNac, Genero: p.Genero, Telefono: p.Telefono,
		Extension: p.Extension, Matricula: p.Matricula, Celular: p.Celular, Correo: p.Correo, Estado: p.Estado}).Error
}

// EliminarPersona /* *********************    Funcion para eliminar personas   ********************* */

func EliminarPersona(id string, tx *gorm.DB) (error, *Persona) {
	var item Persona
	err := tx.Where("id = ?", id).Find(&item).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro el id"), nil
		}
		return err, nil
	}
	return tx.Model(&Persona{}).Where("id=?", id).Delete(&item).Error, &item
}
func CheckAvailablePerson(enrollment string) (bool, *Persona) {
	var user Persona
	db := Conexion.GetDB()
	err := db.Where("matricula = ?", enrollment).First(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return true, nil
		}
		return true, nil
	}
	return false, &user
}
