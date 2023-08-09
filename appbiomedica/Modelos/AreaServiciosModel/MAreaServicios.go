package AreaServiciosModel

import (
	"appbiomedica/Conexion"
	"errors"
	"github.com/jinzhu/gorm"
)

// AgregarAreaServicio /* *********************    Funcion para agregar una nueva area     ********************* */
func (as *AreaServicioUnidad) AgregarAreaServicio() error {
	db := Conexion.GetDB()
	tx := db.Begin()
	err := tx.Create(as).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// ConsultarAreasServicios /* *********************   Funcion para consultar las áreas por servicio y por unidad
//********************* */
func ConsultarAreasServicios(Id_unidad, Id_servicio, Id_usuario string) (*[]AreaServicioUnidad, error) {
	var datos []AreaServicioUnidad
	db := Conexion.GetDB()
	if Id_usuario != "" {
		db = db.Joins("JOIN servicios_usuarios  On area_servicio_unidad.unidad_med_id = servicios_usuarios.unidad_med_id and area_servicio_unidad.servicios_proforma_id = servicios_usuarios.servicios_proforma_id").
			Where("servicios_usuarios.usuario_id = ?", Id_usuario)
	} else {
		db = db.Where("unidad_med_id = ? AND servicios_proforma_id = ?", Id_unidad, Id_servicio)
	}
	err := db.Preload("UnidadMed").Preload("ServiciosProforma").Find(&datos).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return &datos, errors.New("No se encontro el id")
		}
		return &datos, err
	}
	return &datos, err
}
