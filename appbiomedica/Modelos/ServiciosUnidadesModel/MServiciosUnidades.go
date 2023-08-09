package ServiciosUnidadesModel

import (
	"appbiomedica/Conexion"
)

//********************************************************************************************************
/* *********************    Funcion que regresa todos los servicios por una unidad medica     ********************* */

func ConsultaServiciosUnidades(Id [] string) (error, *[] ServiciosPorUnidad) {
	var datos [] ServiciosPorUnidad
	db := Conexion.GetDB()
	if len(Id) != 0 {
		db = db.Where("id_unidad_med IN (?)", Id)
		err := db.Preload("ServiciosProforma").Preload("UnidadMed").Find(&datos).Error
		if err != nil {
			return err, nil
		}
	}
	return nil, &datos
}
