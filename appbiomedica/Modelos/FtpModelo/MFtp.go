package FtpModelo

import (
	"appbiomedica/Conexion"
	"errors"
	"github.com/jinzhu/gorm"
)

/*****  Funcion para consultar un numero de folio ftp  *****/
func ConsultaFolioFtp(folio string) (*Ftp, bool, error) {
	var datos Ftp
	db := Conexion.GetDB()
	err := db.Where("folio_tp = ? ", folio).First(&datos).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, true, errors.New("no se encontro este folio")
		}
		return nil, false, err
	}
	return &datos, false, nil
}
func (f *Ftp) AgregarFolioTp(tx *gorm.DB) error {
	return tx.Create(f).Error
}
