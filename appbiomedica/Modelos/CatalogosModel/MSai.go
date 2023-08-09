package CatalogosModel

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modulos/Query"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

func ConsultaSai(IdGrupo string, tema []string, Nombre []string, numPagina, numRegistro int, cuadro string,
	generico string, especifico string) (error, *[] Sai, int) {
	var datos [] Sai
	var total int
	db := Conexion.GetDB()
	if IdGrupo != "0" && len(tema) == 0 {
		fmt.Println("solo grupo")
		db = db.Where("grupo = ?", IdGrupo)
	}
	if generico != "" {
		db = db.Where("generico = ?", generico)
	}
	if especifico != "" {
		db = db.Where("especifico = ?", especifico)
	}
	if len(Nombre) != 0 {
		db = Query.ContruirQuery(db, Nombre)
	}
	if len(tema) != 0 && IdGrupo == "0" {
		fmt.Println("solo temas")
		db = db.Joins("JOIN grupo_tema  On grupo_tema.grupo = sai.grupo").
			Joins("JOIN tema  On grupo_tema.tema_id = tema.id").
			Where("tema.id IN (?)", tema)
	}
	if len(tema) != 0 && IdGrupo != "0" {
		fmt.Println("temas y grupo")
		db = db.Joins("INNER JOIN grupo_tema  On grupo_tema.grupo = sai.grupo").
			Where("tema_id IN (?) AND grupo_tema.grupo = ?", tema, IdGrupo)
	}
	if cuadro != "" {
		db = db.Where("cuadro_basico = ?", cuadro)
	}
	err := db.Find(&[]Sai{}).Count(&total).Error
	if err != nil {
		return err, nil, total
	}
	err = db.Offset((numPagina - 1) * numRegistro).Limit(numRegistro).Order("grupo").Preload("DescripcionGrupo").Preload("UnidadPresentacion").
		Preload("TipoPresentacion").Find(&datos).Error
	if err != nil {
		return err, nil, total
	}
	fmt.Println("Total de datos consultados", datos)
	return nil, &datos, total
}

func ConsultarGenericosEspecifico(grupo [] string, generico string, todosGenericos, todosEspecificos bool) (error, *[] Sai) {
	var datos [] Sai
	db := Conexion.GetDB()
	if len(grupo) != 0 && generico == "" {
		db = db.Select("distinct generico").Where("grupo IN (?)", grupo).Order("generico")
	}
	if generico != "" && len(grupo) != 0 {
		db = db.Select("distinct especifico").Where("grupo IN (?) AND generico = ?", grupo, generico).Order("especifico")
	}
	if todosGenericos {
		db = db.Select("distinct generico").Order("generico")
	}
	if todosEspecificos {
		db = db.Select("distinct especifico").Order("especifico")
	}
	err := db.Find(&datos).Error
	if err != nil {
		return err, nil
	}
	fmt.Println(grupo, generico, todosGenericos, todosEspecificos)
	return nil, &datos
}

//********************************************************************************************************
//-----------------------Consultar un sola clave

func ConsultaClaveSai(ID, grupo, generico, especifico, diferenciador, variable string) (error, *Sai) {
	var datos Sai
	db := Conexion.GetDB()
	if ID != "0" {
		db = db.Where("id=?", ID)
	}
	if grupo != "" && generico != "" && especifico != "" {
		db = db.Where("grupo = ? AND generico=? AND especifico =? AND diferenciador =? AND variable =?",
			grupo, generico, especifico, diferenciador, variable)
	}

	err := db.Preload("DescripcionGrupo").Preload("UnidadPresentacion").
		Preload("TipoPresentacion").Find(&datos).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("Este id no existe"), nil
		}
		return err, nil
	}
	return nil, &datos

}
