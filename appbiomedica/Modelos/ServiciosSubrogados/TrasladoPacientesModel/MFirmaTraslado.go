package TrasladoPacientesModel

import (
	"appbiomedica/Conexion"
	"errors"
	"github.com/jinzhu/gorm"
)

/*****  Funcion para guardar un nombre y cargo de una solicitud  *****/
func (ft *FirmaTraslado) AgregarFirmaTraslado(tx *gorm.DB) error {
	return tx.Create(ft).Error
}

/*****  Funcion para consultar la firma asociada a la solicitud de traslado  *****/
func ConsultarFirmaTraslado(folio string, tipo int) (error, *FirmaTraslado) {
	var firma FirmaTraslado
	db := Conexion.GetDB()
	err := db.Joins("inner join firma_solicitud on firma_traslado.firma_solicitud_id = firma_solicitud.id").
		Where("solicitud_traslado_folio =? AND firma_solicitud.tipo=?", folio, tipo).Preload("FirmaSolicitud").
		First(&firma).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No hay datos relacionados con esta unidad"), nil
		}
		return err, nil
	}
	return nil, &firma
}
