package SolicitudModel

import (
	"appbiomedica/Conexion"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

//********************************************************************************************************
/* *********************    Funcion que agrega una solicitud      ********************* */
func (p *Solicitud) AgregarSolicitud() error {
	db := Conexion.GetDB()
	tx := db.Begin()
	err := tx.Create(p).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

//********************************************************************************************************
/* *********************    Funcion que regresa las solicitudes  por usuario     ********************* */
func ConsultaSolicitudes(Id, tipo, idSolicitud string) (error, *Solicitud) {
	var datos Solicitud
	db := Conexion.GetDB()
	if tipo != "0" {
		db = db.Where("tipo_solicitud_id = ?", tipo)
	}
	if Id != "0" {
		db = db.Where("usuario_id = ?", Id)
	}
	if idSolicitud != "0" {
		db = db.Where("id = ?", idSolicitud)
	}
	err := db.Preload("Usuario").Preload("Usuario.Persona").Preload("TipoSolicitud").
		Preload("Estatus").Preload("UnidadMed").Find(&datos).Error
	if err != nil {
		return err, nil
	}
	fmt.Println("resultado", &datos)
	return nil, &datos

}

//********************************************************************************************************
/* *********************    Funcion que regresa todas las solicitudes  por usuario     ********************* */
func ConsultaTodasSolicitudes(Id, validados string) (error, *[]Solicitud) {
	var datos []Solicitud
	db := Conexion.GetDB()
	if Id != "" {
		db = db.Where("usuario_id = ?", Id)
	}
	if validados == "true" {
		db = db.Where("estatus_id IN (?)", []string{"2", "3", "4", "5"})
	}
	if validados == "false" {
		db = db.Where("estatus_id IN (?)", []string{"6", "7", "8", "9", "10", "11", "12"})
	}
	err := db.Where("estado = ?", true).Preload("Usuario").Preload("Usuario.Persona").
		Preload("TipoSolicitud").
		Preload("Estatus").Preload("UnidadMed").Find(&datos).Error
	if err != nil {
		return err, nil
	}
	return nil, &datos

}

//********************************************************************************************************
/* *********************    Funcion para actualizar una solicitud     ********************* */
func (a *Solicitud) ActualizarSolicitud(id string, tx *gorm.DB) error {
	var claves Solicitud
	fmt.Println("esto llega al actualizar weee", a)
	err := tx.Where("id = ?", id).First(&claves).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro el id")
		}
		return err
	}
	if a.EstatusId == 2 {
		tx = tx.Model(&a).Where("id=?", id).Update(Solicitud{FechaCreacion: a.FechaCreacion})
	}
	fmt.Println("estado:  ", a.Estado)
	return tx.Model(&a).Where("id=?", id).Updates(map[string]interface{}{"estatus_id": a.EstatusId,
		"estado": a.Estado, "total": a.Total}).Error
}

//********************************************************************************************************
/* *********************    Funcion para actualizar solo el total de una solicitud     ********************* */
func ActualizarTotalSolicitud(id string, tx *gorm.DB, total float64) error {
	var claves Solicitud
	err := tx.Where("id = ?", id).First(&claves).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro el id")
		}
		return err
	}
	return tx.Model(Solicitud{}).Where("id=?", id).Update(Solicitud{Total: total}).Error
}

//********************************************************************************************************
/* *********************    Funcion para actualizar solo el estatus de una solicitud     ********************* */
func ActualizarEstatuslSolicitud(id string, tx *gorm.DB, estatus int) error {
	var claves Solicitud
	err := tx.Where("id = ?", id).First(&claves).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro el id")
		}
		return err
	}
	return tx.Model(Solicitud{}).Where("id=?", id).Update(Solicitud{EstatusId: estatus}).Error
}
