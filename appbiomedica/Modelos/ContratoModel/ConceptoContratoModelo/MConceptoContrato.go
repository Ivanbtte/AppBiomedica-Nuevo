package ConceptoContratoModelo

import (
	"appbiomedica/Conexion"
	"errors"
	"github.com/jinzhu/gorm"
)

//********************************************************************************************************
/* *********************   Funcion para agregar un concepto de un contrato     ********************* */
func (concepto *ConceptoContrato) AgregarConceptoContrato(tx *gorm.DB) error {
	return tx.Create(concepto).Error
}

//********************************************************************************************************
/* *************funcion para verificar que un concepto no este registrado en el mismo ocntrato ********************* */
func VerificaEquipo(id_prei, contrato_id string) (bool, *ConceptoContrato) {
	var item ConceptoContrato
	db := Conexion.GetDB()
	err := db.Where("prei_id_articulo = ? AND contrato_numero_contrato  = ?", id_prei, contrato_id).
		Preload("Prei").Preload("Contrato").First(&item).Error
	if err != nil {
		return false, nil
	}
	return true, &item
}
//********************************************************************************************************
/* *********************    Funcion para cosnultar una distribucion     ********************* */
func RegresaConcepto(id string) (error, *ConceptoContrato) {
	var concepto ConceptoContrato
	db := Conexion.GetDB()
	err := db.Where("id = ?", id).First(&concepto).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro este concepto"), nil
		}
		return err, nil
	}
	return nil, &concepto
}
//********************************************************************************************************
/* *********************   Funcion que regresa la lista de conceptos en el contrato     ********************* */
func ConsultaConceptos(id_contrato string) (error, *[]ConceptoContrato, float64) {
	var datos []ConceptoContrato
	type tot struct {
		Total float64
	}
	var total tot
	db := Conexion.GetDB()
	err := db.Where("contrato_numero_contrato = ?", id_contrato).Preload("Prei").Preload("Contrato").
		Order("cantidad_distribuida desc").Find(&datos).Error
	if err != nil {
		return err, nil, total.Total
	}
	err = db.Table("concepto_contrato").Select("sum(precio_uni_sn_iva * cantidad_concepto) as total").
		Where("contrato_numero_contrato = ?", id_contrato).Scan(&total).Error
	if err != nil {
		return err, nil, total.Total
	}
	return nil, &datos, total.Total
}

//********************************************************************************************************
/* *********************    Funcion para actualizar un concepto     ********************* */
func (c ConceptoContrato) ActualizarConcepto(id string) error {
	db := Conexion.GetDB()
	var concepto ConceptoContrato
	tx := db.Begin()
	err := tx.Where("id = ?", id).First(&concepto).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro el id")
		}
		return err
	}
	err = tx.Model(ConceptoContrato{}).Where("id=?", id).Update(ConceptoContrato{PrecioUniSnIva: c.PrecioUniSnIva,
		CantidadConcepto: c.CantidadConcepto, Marca: c.Marca, Modelo: c.Modelo, FechaMaxEntrega: c.FechaMaxEntrega,
		GarantiaBienes: c.GarantiaBienes}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

//********************************************************************************************************
/* *********************    Funcion para actualizar un concepto     ********************* */
func ActualizarDistribucion(id string, cantidadDistribuida, cantidadContratada int) error {
	db := Conexion.GetDB()
	tx := db.Begin()
	var concepto ConceptoContrato
	err := tx.Where("id = ?", id).First(&concepto).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro el id")
		}
		return err
	}
	err = tx.Model(&concepto).Updates(map[string]interface{}{"cantidad_distribuida": cantidadDistribuida,
		"cantidad_concepto": cantidadContratada}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
//********************************************************************************************************
/* *********************    Funcion para eliminar un concepto de un contrato   ********************* */
func  EliminarConcepto(id_concepto string, tx *gorm.DB) (error, *ConceptoContrato) {
	var item ConceptoContrato
	err := tx.Where("id = ?", id_concepto).Find(&item).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro el id"), nil
		}
		return err, nil
	}
	return tx.Model(&ConceptoContrato{}).Where("id=?", id_concepto).Delete(&item).Error, &item
}