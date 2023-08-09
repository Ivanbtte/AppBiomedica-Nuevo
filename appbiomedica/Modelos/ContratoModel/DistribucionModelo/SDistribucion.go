package DistribucionModelo

import (
	"appbiomedica/Modelos/ContratoModel/ConceptoContratoModelo"
	"appbiomedica/Modelos/ModeloUnidadM"
)

type DistribucionUnidadMed struct {
	Id                       string `gorm:"PRIMARY_KEY"`
	Cantidad                 int
	UnidadMedClvPresupuestal string
	UnidadMed                *ModeloUnidadM.UnidadMed `gorm:"unique;not null;foreignkey:clv_presupuestal;association_foreignkey:unidad_med_clv_presupuestal"`
	ConceptoContratoId       string
	ConceptoContrato         *ConceptoContratoModelo.ConceptoContrato
}
type SeguimientoClave struct {
	Id                      string
	NumSerie                string
	FechaInstalacion        string
	FinGarantiaBienes       string
	ServicioUbicacionFinal  string
	DistribucionUnidadMedId string
	DistribucionUnidadMed   *DistribucionUnidadMed
}