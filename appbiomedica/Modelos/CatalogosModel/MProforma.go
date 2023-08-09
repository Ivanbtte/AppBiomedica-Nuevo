package CatalogosModel

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modulos/Query"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

/* *********************    funcion para consultar el catalogo proforma     ********************* */
func ConsultaProforma(IdServicio string, Nombre []string, NumPagina, NumRegistros int) (error, *[] Proforma, int64) {
	var datos [] Proforma
	type tot struct {
		Total int64
	}
	var total tot
	db := Conexion.GetDB()
	if IdServicio != "0" {
		db = db.Where("id_servicio = ? ", IdServicio)
	}
	if len(Nombre) != 0 {
		db = Query.ContruirQuery(db, Nombre)
	}

	err := db.Table("proforma").Select("count(distinct prei_id_articulo) as total").Scan(&total).Error
	if err != nil {
		return err, nil, 0
	}

	fmt.Print("total: ",total.Total)

	err = db.Table("proforma").Select("distinct on (prei_id_articulo)*").Offset((NumPagina - 1) * NumRegistros).Limit(NumRegistros).Scan(&datos).Error
	if err != nil {
		return err, nil, 0
	}
	fmt.Println("Total de datos consultados", total)
	return nil, &datos, total.Total

}

/* *********************   Funcion para consultar una sola clave proforma     ********************* */
func ConsultaClaveProforma(ID string) (error, *[]Proforma) {
	var datos []Proforma
	db := Conexion.GetDB()
	err := db.Where("prei_id_articulo=?", ID).Preload("Pre").Preload("ServiciosPro").
		Preload("Pre.UnidadesMedidas").Find(&datos).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("Este id no existe"), nil
		}
		return err, nil
	}
	return nil, &datos

}
