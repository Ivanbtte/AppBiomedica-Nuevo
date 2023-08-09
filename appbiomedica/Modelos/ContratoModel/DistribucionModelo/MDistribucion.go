package DistribucionModelo

import (
	"appbiomedica/Conexion"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

//********************************************************************************************************
/* *********************   Funcion para agregar una distribucion de un equipo     ********************* */
func (concepto *DistribucionUnidadMed) AgregarDistribucion(tx *gorm.DB) error {
	return tx.Create(concepto).Error
}

//********************************************************************************************************
/* *****************  Funcion que regresa la lista de distribucion de un concepto por unidad   ********************* */
func ConsultaDistribucion(id_concepto string) (error, *[]DistribucionUnidadMed) {
	var datos []DistribucionUnidadMed
	db := Conexion.GetDB()
	err := db.Where("concepto_contrato_id = ?", id_concepto).Preload("ConceptoContrato").
		Preload("ConceptoContrato.Prei").Preload("UnidadMed").Preload("UnidadMed.Delegacion").
		Find(&datos).Error
	if err != nil {
		return err, nil
	}
	return nil, &datos

}

//********************************************************************************************************
/* *******************    Funcion que comprobara la cantidad de la distribucion por id prei     ******************* */
func VerificarCantidades(id_concepto string) (bool, int) {
	type tot struct {
		Total int
	}
	var total tot
	db := Conexion.GetDB()
	err := db.Table("distribucion_unidad_med").Select("sum(cantidad) as total").
		Where("concepto_contrato_id = ?", id_concepto).Scan(&total).Error
	if err != nil {
		return false, 0
	}
	return true, total.Total
}

//********************************************************************************************************
/* **************    Funcion que valida si ya existe una distribucion con la misma unidad medica     *********** */
/*Regresa true si ya existe y false si no existe ningun registro de esa distribucion*/
func ValidaRepetidos(clv_presupuestal, concepto_id string) (bool, *DistribucionUnidadMed) {
	var distribucion DistribucionUnidadMed
	db := Conexion.GetDB()
	err := db.Where("unidad_med_clv_presupuestal = ? and concepto_contrato_id = ?", clv_presupuestal, concepto_id).First(&distribucion).Error
	if err != nil {
		return false, nil
	}
	return true, &distribucion
}

//********************************************************************************************************
/* *********************    Funcion para Editar una distribucion     ********************* */
func (dis *DistribucionUnidadMed) ActualizarDistribucion(id string, tx *gorm.DB) error {
	var distribucion DistribucionUnidadMed
	err := tx.Where("id = ?", id).First(&distribucion).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro esta distribucion")
		}
		return tx.Rollback().Error
	}
	return tx.Model(&dis).Where("id=?", id).Update(DistribucionUnidadMed{Cantidad: dis.Cantidad,
		UnidadMedClvPresupuestal: dis.UnidadMedClvPresupuestal}).Error
}

//********************************************************************************************************
/* *********************    Funcion para consultar una distribucion     ********************* */
func ConsultarDistribucion(id_distribucion string) (error, *DistribucionUnidadMed) {
	var distribucion DistribucionUnidadMed
	db := Conexion.GetDB()
	err := db.Where("id = ?", id_distribucion).Preload("ConceptoContrato").Preload("UnidadMed").
		First(&distribucion).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro esta distribucion"), nil
		}
		return err, nil
	}
	return nil, &distribucion
}

//********************************************************************************************************
/* *********************    Funcion para eliminar una distribucion     ********************* */
func EliminarDistribucion(id string, tx *gorm.DB) (error, *DistribucionUnidadMed) {
	var item DistribucionUnidadMed
	err := tx.Where("id = ?", id).Find(&item).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro esta distribucion"), nil
		}
		return err, nil
	}
	return tx.Model(&DistribucionUnidadMed{}).Where("id=?", id).Delete(&item).Error, &item
}

//********************************************************************************************************
/* ******************    Funcion que regresa las distribuciones por contrato y por delegacion     ****************** */
func BuscarDistribucion(num_contrato, clv_dele string) (error, *[]DistribucionUnidadMed) {
	var distribucion []DistribucionUnidadMed
	db := Conexion.GetDB()
	err := db.
		Joins("inner join unidad_med um on distribucion_unidad_med.unidad_med_clv_presupuestal = um.clv_presupuestal").
		Joins("inner join concepto_contrato cc on distribucion_unidad_med.concepto_contrato_id = cc.id").
		Where("um.delegaciones_clv_dele = ?  AND cc.contrato_numero_contrato = ?", clv_dele,
			num_contrato).Find(&distribucion).Error
	if err != nil {
		fmt.Println("Ocurrio un error", err.Error())
		return err, nil
	}
	return nil, &distribucion
}
