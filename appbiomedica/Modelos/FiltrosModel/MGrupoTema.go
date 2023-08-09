package FiltrosModel

import (
	"appbiomedica/Conexion"
)

//********************************************************************************************************
//-----------------------Consultar todos los grupos_temas

func ConsultaGrupoTemas(todos string, temas []string) (error, *[] GrupoTema) {
	var datos [] GrupoTema
	db := Conexion.GetDB()
	if todos == "false" {
		db = db.Select("distinct tema_id")
	}
	if len(temas) != 0 {
		db = db.Where("tema_id IN (?)", temas)
	}
	err := db.Preload("Tema").Find(&datos).Error
	if err != nil {
		return err, nil
	}
	return nil, &datos
}
