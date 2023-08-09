package PenasModelo

import "appbiomedica/Modelos/ContratoModel/ContratoModelo"

type PenasConvencionales struct {
	Id                     string `gorm:"PRIMARY_KEY"`
	Descripcion            string
	Porcentaje             float32
	NotaCreditoOficio      string
	ReferenciaCobro        string
	FechaCobro             string
	ContratoNumeroContrato string
	Contrato               *ContratoModelo.Contrato `gorm:"unique;not null;foreignkey:numero_contrato;association_foreignkey:contrato_numero_contrato"`
	Estado                 bool
}
