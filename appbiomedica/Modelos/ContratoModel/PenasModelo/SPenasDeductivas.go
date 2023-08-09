package PenasModelo

import (
	"appbiomedica/Modelos/ContratoModel/ContratoModelo"
)

type PenasDeductivas struct {
	Id                     string `gorm:"PRIMARY_KEY"`
	ConceptoObligacion     string
	NivelServicio          string
	UnidadMedida           string
	Deduccion              float32
	DescripcionDeduccion   string
	LimiteIncumplimiento   string
	NotaCreditoOficio      string
	ReferenciaCobro        string
	FechaCobro             string
	ContratoNumeroContrato string
	Contrato               *ContratoModelo.Contrato `gorm:"unique;not null;foreignkey:numero_contrato;association_foreignkey:contrato_numero_contrato"`
	Estado                 bool
}
