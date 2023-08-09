package BitacoraModel

import (
	"appbiomedica/Conexion"
	"github.com/jinzhu/gorm"
)

//********************************************************************************************************
/* *********************    Funcion para agregar un registro a la bitacora     ********************* */
func (p *BitacoraSolicitud) AgregaBitacoraSolicitud(tx *gorm.DB) error {
	return tx.Create(p).Error
}

//********************************************************************************************************
/* *********************    Funcion que regresa todos la bitacora de una solicitud     ********************* */
func ConsultaBitacoraSolicitud(Id string) (error, *[] BitacoraSolicitud) {
	var datos [] BitacoraSolicitud
	db := Conexion.GetDB()
	if Id != "0" {
		db = db.Where("solicitud_id = ? ", Id)
	}
	err := db.Preload("Estatus").Find(&datos).Error
	if err != nil {
		return err, nil
	}
	return nil, &datos

}
