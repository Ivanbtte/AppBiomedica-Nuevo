package ConceptosServiciosModelo

import (
	"appbiomedica/Conexion"
)

//********************************************************************************************************
/* *********************   Funcion para agregar un concepto de un contrato     ********************* */
func (concepto *Conceptos) AgregarConceptoContrato() error {
	db := Conexion.GetDB()
	err := db.Create(concepto).Error
	if err != nil {
		return err
	}
	return nil
}

/*****  Funcion para consultar un concepto  *****/
func BuscarConceptoContrato(n_contrato string) (Conceptos, bool) {
	var concepto Conceptos
	db := Conexion.GetDB()
	err := db.Where("trim(lower(contrato_numero_contrato)) = ?", n_contrato).First(&concepto).Error
	if err != nil {
		return concepto, false
	}
	return concepto, true
}
