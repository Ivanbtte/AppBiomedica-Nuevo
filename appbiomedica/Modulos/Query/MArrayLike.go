package Query

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

/* ****************   Esta funcion crea varios where like para hacer una busqueda mas avanzada     **************** */
func ContruirQuery(db *gorm.DB, arreglo []string) *gorm.DB {
	for _, filtro := range arreglo {
		db = db.Where("LOWER (descripcion) LIKE ?", "%"+strings.ToLower(filtro)+"%")
	}
	return db
}

/* ****************   Esta funcion crea varios where para buscar por varias palabras claves    **************** */
func ContruirQueryContrato(db *gorm.DB, arreglo []string) *gorm.DB {
	db = db.Joins("INNER JOIN proveedor p on contrato.proveedor_n_prov_imss = p.n_prov_imss")
	for _, filtro := range arreglo {
		db = db.Where("LOWER (numero_contrato) LIKE ? OR LOWER (proced_contratacion) LIKE ? OR LOWER (tipo_proced_contratacion) LIKE ?", "%"+strings.ToLower(filtro)+"%", "%"+strings.ToLower(filtro)+"%","%"+strings.ToLower(filtro)+"%")
		/*.Or("LOWER (proveedor_n_prov_imss) LIKE ?", "%"+strings.ToLower(filtro)+"%").
		Or("LOWER (p.nombre_empresa) LIKE ?", "%"+strings.ToLower(filtro)+"%")*/
	}
	return db
}

/* **********  Esta funcion crea varios where para buscar por varias palabras claves de un proveedor************ */
func ContruirQueryProveedor(db *gorm.DB, arreglo []string) *gorm.DB {
	for _, filtro := range arreglo {
		db = db.Where("LOWER (alias_empresa) LIKE ? OR LOWER (nombre_empresa) LIKE ? OR LOWER (n_prov_imss) LIKE ?", "%"+strings.ToLower(filtro)+"%", "%"+strings.ToLower(filtro)+"%", "%"+strings.ToLower(filtro)+"%")
	}
	return db
}
/* **********  Esta funcion crea varios where para buscar por varias palabras claves de una solicitud de traslado
************ */
func ContruirQuerySolicitudTraslado(db *gorm.DB, filtro string) *gorm.DB {
		db = db.Where("LOWER (folio) LIKE ? OR LOWER (nombre_paciente) LIKE ? OR LOWER (destino) LIKE ? OR LOWER (" +
			"origen) LIKE ? OR LOWER (ftp01) LIKE ?",
			"%"+strings.ToLower(filtro)+"%","%"+strings.ToLower(filtro)+"%","%"+strings.ToLower(filtro)+"%",
			"%"+strings.ToLower(filtro)+"%","%"+strings.ToLower(filtro)+"%")
		t, err := time.Parse("02-1-2006", filtro)
		if err != nil {
			fmt.Println(err)
		}else{
			db = db.Or("fecha_cita :: date = ?", t)
		}

	return db
}