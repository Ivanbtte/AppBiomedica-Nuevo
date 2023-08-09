package CatalogosModel

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modulos/Query"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

//********************************************************************************************************
//-----------------------Consulta todos el prei

func ConsultaPrei(id_Cuadro int, Nombre []string, grupo, generico, especfico string, numPagina,
	numRegistros int) (error, *[] Prei, int64) {
	var datos [] Prei
	var total int64
	db := Conexion.GetDB()
	if id_Cuadro != 0 {
		db = db.Where("cuadro_basico_id=?", id_Cuadro)
	}
	if len(Nombre) != 0 {
		db = Query.ContruirQuery(db, Nombre)
	}
	if grupo != "" {
		db = db.Where("grupo = ?", grupo)
	}
	if generico != "" {
		db = db.Where("generico = ?", generico)
	}
	if especfico != "" {
		db = db.Where("especifico = ?", especfico)
	}
	err := db.Find(&[]Prei{}).Count(&total).Error
	if err != nil {
		return err, nil, total
	}
	err = db.Offset((numPagina - 1) * numRegistros).Limit(numRegistros).Preload("UnidadesMedidas").
		Order("grupo").Find(&datos).Error
	if err != nil {
		return err, nil, total
	}

	fmt.Println("Total de datos consultados: ", total)
	return nil, &datos, total
}

/* *********************    Funcion que regresa los grupos, genericos y especificos     ********************* */

func ConsultarGruposPrei(grupo string, generico string, grupos bool) (error, *[] Prei) {
	var datos [] Prei
	db := Conexion.GetDB()
	if grupos {
		db = db.Select("distinct grupo").Order("grupo")
	}
	if grupo != "" && generico == "" {
		db = db.Select("distinct generico").Where("grupo = ?", grupo).Order("generico")
	}
	if grupo != "" && generico != "" {
		db = db.Select("distinct especifico").Where("grupo = ? AND generico = ?", grupo, generico).
			Order("especifico")
	}
	err := db.Find(&datos).Error
	if err != nil {
		return err, nil
	}
	return nil, &datos
}

//********************************************************************************************************
//-----------------------Consultar un sola clave

func ConsultaUnaClave(ID, prei_id string) (error, *Prei) {
	var datos Prei
	db := Conexion.GetDB()
	if prei_id != "" {
		db = db.Where("id_articulo = ?", prei_id)
	}
	if ID != "" {
		db = db.Where("id=?", ID)
	}
	err := db.Preload("UnidadesMedidas").Preload("CuadroBasico").Find(&datos).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("Este id no existe"), nil
		}
		return err, nil
	}
	return nil, &datos

}
