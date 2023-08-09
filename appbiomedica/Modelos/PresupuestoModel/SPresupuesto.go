package PresupuestoModel

type Presupuesto struct {
	Id          string `gorm:"PRIMARY_KEY"`
	Presupuesto float64
	Estado      bool
}
