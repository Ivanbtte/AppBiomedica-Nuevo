package SolicitudModel

import (
	"appbiomedica/Conexion"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

//********************************************************************************************************
/* *********************    Funcion que agrega una clave a una  solicitud      ********************* */
func (p *ClavesSolicitud) AgregarClavesSolicitud() error {
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
/* *********************    Funcion que busca una clave repetida     ********************* */
func BuscaClave(Id string, servicio string) (bool, int) {
	var datos ClavesSolicitud
	var id int
	db := Conexion.GetDB()
	err := db.Preload("Solicitud").Where("clave_id = ? AND servicios_proforma_id = ?", Id, servicio).Find(&datos).Error
	if err != nil {
		return false, 0
	}
	id = datos.Solicitud.UnidadMedId
	fmt.Println(id)
	return true, id

}

//********************************************************************************************************
/* *********************    Funcion que regresa una clave    ********************* */
func ConcultaUnConceptoId(id string) (error, *ClavesSolicitud) {
	var item ClavesSolicitud
	db := Conexion.GetDB()
	err := db.Where("id = ?", id).Preload("ServiciosProforma").Preload("Solicitud").Preload("Solicitud.UnidadMed").First(&item).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro el id"), nil
		}
		return err, nil
	}
	return nil, &item
}

//********************************************************************************************************
/* *********************    Funcion que regresa todos los conceptos de una solicitud     ********************* */
func ConsultaConceptos(Id, clave, unidad, final, servicioId string) (error, *[] ClavesSolicitud) {
	var datos [] ClavesSolicitud
	db := Conexion.GetDB()
	if Id != "" {
		db = db.Where("solicitud_id = ? ", Id)
	}
	if clave != "" {
		db = db.Joins("JOIN solicitud  On solicitud.id = claves_solicitud.solicitud_id").
			Where("clave_id = ?", clave)
	}
	if unidad != "" {
		db = db.Joins("JOIN solicitud  On solicitud.id = claves_solicitud.solicitud_id").
			Where("unidad_med_id = ?", unidad)
	}
	if final != "" {
		db = db.Where("estado = ?", true)
	}
	if servicioId != "" {
		db = db.Where("servicios_proforma_id = ?", servicioId)
	}
	err := db.Preload("ServiciosProforma").Preload("Solicitud").Preload("Solicitud.UnidadMed").Find(&datos).Error
	if err != nil {
		return err, nil
	}
	return nil, &datos

}

//********************************************************************************************************
/* *********************    Funcion para actualizar una clave de una solicitud     ********************* */
func (a *ClavesSolicitud) ActualizarClaveConcepto(id string, tx *gorm.DB) (error, *ClavesSolicitud) {
	var claves ClavesSolicitud
	err := tx.Where("id = ?", id).Preload("ServiciosProforma").Preload("Solicitud").Preload("Solicitud.UnidadMed").First(&claves).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro el id"), nil
		}
		return err, nil
	}
	return tx.Model(&a).Where("id=?", id).Update(ClavesSolicitud{CantidadSolicitada: a.CantidadSolicitada, Total: a.Total}).Error, &claves
}

//********************************************************************************************************
/* *********************    Funcion para actualizar la cantidad aprobada   ********************* */
func (a *ClavesSolicitud) AprobarConcepto(id string, cantidad int, tx *gorm.DB) (error, *ClavesSolicitud) {
	var claves ClavesSolicitud
	err := tx.Where("id = ?", id).Preload("ServiciosProforma").Preload("Solicitud").Preload("Solicitud.UnidadMed").First(&claves).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro el id"), nil
		}
		return err, nil
	}
	return tx.Model(&a).Where("id=?", id).Update(ClavesSolicitud{CantidadAprobada: cantidad, Total: a.Total}).Error, a
}

//********************************************************************************************************
/* *********************    Funcion que actualiza el estado de una clave    ********************* */
func (p *ClavesSolicitud) ApruebaEstadoclave(id string, tx *gorm.DB) error {
	var claves ClavesSolicitud
	err := tx.Where("id = ?", id).First(&claves).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro el id")
		}
		return err
	}
	if claves.CantidadAprobada == 0 {
		tx = tx.Model(&p).Where("id=?", id).Update(ClavesSolicitud{CantidadAprobada: p.CantidadSolicitada})
	}
	return tx.Model(&p).Where("id=?", id).Update(ClavesSolicitud{Estado: true}).Error
}

//********************************************************************************************************
/* *********************    Funcion para eliminar una clave de una solicitud     ********************* */
func (a *ClavesSolicitud) EliminarClaveConcepto(id string, tx *gorm.DB) (error, *ClavesSolicitud) {
	var claves ClavesSolicitud
	err := tx.Where("id = ?", id).Preload("Solicitud").Find(&claves).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro el id"), nil
		}
		return err, nil
	}
	return tx.Model(&a).Where("id=?", id).Delete(&claves).Error, &claves
}
