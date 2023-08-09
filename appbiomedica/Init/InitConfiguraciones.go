package Init

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/FechaLimiteModel"
	"appbiomedica/Modelos/PresupuestoModel"
)

func InitConfig() error {
	db := Conexion.GetDB()
	if !db.HasTable(&FechaLimiteModel.FechaLimite{}) {
		if err := db.CreateTable(&FechaLimiteModel.FechaLimite{}).Error; err != nil {
			return err
		}
	}
	if !db.HasTable(&PresupuestoModel.Presupuesto{}) {
		if err := db.CreateTable(&PresupuestoModel.Presupuesto{}).Error; err != nil {
			return err
		}
	}
	return nil
}
