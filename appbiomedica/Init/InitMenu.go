package Init

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/MenuModel"
	"appbiomedica/Modelos/ModuloModel"
)

func InitMenu() error {
	db := Conexion.GetDB()
	if !db.HasTable(&ModuloModel.Modulo{}) {
		if err := db.CreateTable(&ModuloModel.Modulo{}).Error; err != nil {
			return err
		}
	}
	if !db.HasTable(&MenuModel.Menu{}) {
		if err := db.CreateTable(&MenuModel.Menu{}).
			AddForeignKey("menu_parent", "menu(id)", "RESTRICT", "CASCADE").
			AddForeignKey("modulo_id", "modulo(id)", "RESTRICT", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	if !db.HasTable(&ModuloModel.Action{}) {
		if err := db.CreateTable(&ModuloModel.Action{}).
			AddForeignKey("modulo_id", "modulo(id)", "RESTRICT", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	return nil
}
