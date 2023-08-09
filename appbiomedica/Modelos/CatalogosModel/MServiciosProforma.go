package CatalogosModel

import "appbiomedica/Conexion"

/* *********************    funcion para consultar todos los servicios proforma     ********************* */

func ConsultaServProforma() (error, *[] ServiciosProforma) {
	var datos [] ServiciosProforma
	db := Conexion.GetDB()
	err := db.Find(&datos).Error
	if err != nil {
		return err, nil
	}
	return nil, &datos

}
func ConsultarServicio(id string) (error, *ServiciosProforma) {
	var datos ServiciosProforma
	db := Conexion.GetDB()
	err := db.Where("id = ?", id).Find(&datos).Error
	if err != nil {
		return err, nil
	}
	return nil, &datos
}

/* *********************    funcion que regresa el nombre de los servicios a la que pertenece una clave proforma
 ********************* */
func ConsultaNombresServicio(Id string) (error, *[] ServiciosProforma) {
	var datos [] ServiciosProforma
	db := Conexion.GetDB()
	err := db.Joins("JOIN proforma ON proforma.id_servicio = servicios_proforma.id").
		Where("proforma.prei_id_articulo = ?", Id).Find(&datos).Error

	if err != nil {
		return err, nil
	}
	return nil, &datos

}
