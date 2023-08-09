package TrasladoPacientesModel

import (
	"appbiomedica/Conexion"
	"errors"
	"github.com/jinzhu/gorm"
)

/*****  Funcion para guardar un nombre y cargo de una solicitud  *****/
func (fs *FirmaSolicitud) AgregarFirma(tx *gorm.DB) error {
	return tx.Create(fs).Error
}

/*****  Funcion para consultar un nombre y cargo de una solicitud por unidad medica  *****/
func ReturnAllFirm(unitId, typeFirm int, state bool) (error, *[]FirmaSolicitud) {
	var firm []FirmaSolicitud
	db := Conexion.GetDB()
	err := db.Where("unidad_med_id = ? AND estado =?", unitId, state).Where("tipo = ?", typeFirm).Find(&firm).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("no hay datos relacionados con esta unidad"), nil
		}
		return err, nil
	}
	return nil, &firm
}

/*****  Funcion para consultar una firma por id *****/
func ReturnFirm(id string) (error, *FirmaSolicitud) {
	var firm FirmaSolicitud
	db := Conexion.GetDB()
	err := db.Where("id = ? ", id).First(&firm).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("no hay datos relacionados con esta unidad"), nil
		}
		return err, nil
	}
	return nil, &firm
}

/*****  Funcion para actualizar los datos de una firma  *****/
func (fs FirmaSolicitud) UpdateFirm(id string, tx *gorm.DB) error {
	if fs.Tipo != 0 {
		err := tx.Model(&fs).Where("id=?", id).Update(FirmaSolicitud{Nombre: fs.Nombre, Cargo: fs.Cargo,
			Matricula: fs.Matricula, Tipo: fs.Tipo}).Error
		if err != nil {
			tx.Rollback()
			return err
		}
		return tx.Commit().Error
	} else {
		err := tx.Model(&fs).Where("id=?", id).Update(FirmaSolicitud{Nombre: fs.Nombre, Cargo: fs.Cargo,
			Matricula: fs.Matricula}).Error
		if err != nil {
			tx.Rollback()
			return err
		}
		return tx.Commit().Error
	}
}

/*****  Funcion para habilitar o deshabilitar una firma *****/
func ActualizarFirma(id string, state bool, tx *gorm.DB) error {
	return tx.Model(&FirmaSolicitud{}).Where("id=?", id).Update("estado", state).Error
}

/*****  Funcion para consultar las dos firmas de una unidad medica  *****/
func ConsultarFirmasUnidad(unidad_med int, estado string) (error, *[]FirmaSolicitud) {
	var firma []FirmaSolicitud
	db := Conexion.GetDB()
	if estado != "" {
		db = db.Where("estado =?", estado)
	}
	err := db.Where("unidad_med_id = ?", unidad_med).Order(" estado desc").Order("tipo desc").Find(&firma).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("no hay datos relacionados con esta unidad"), nil
		}
		return err, nil
	}
	return nil, &firma
}

func FindDuplicate(matricula string, unitId int) bool {
	var firma FirmaSolicitud
	db := Conexion.GetDB()
	err := db.Where("unidad_med_id = ? AND matricula =?", unitId, matricula).First(&firma).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false
		}
		return false
	}
	return true
}
