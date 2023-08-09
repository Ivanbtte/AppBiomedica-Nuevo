package Init

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/ModeloDistribucionGeografica"
)

func InitDistGeog() error {
	db := Conexion.GetDB()
	if !db.HasTable(&ModeloDistribucionGeografica.Estados{}) {

		if err := db.CreateTable(&ModeloDistribucionGeografica.Estados{}).Error; err != nil {
			return err
		}
	}
	if !db.HasTable(&ModeloDistribucionGeografica.Municipios{}) {

		if err := db.CreateTable(&ModeloDistribucionGeografica.Municipios{}).Error; err != nil {
			return err
		}
	}
	return nil
}
