package RepresentanteModel

import "appbiomedica/Modelos/ContratoModel/ContratoModelo"

type RepresentanteLegal struct {
	Id                     string `gorm:"PRIMARY_KEY"`
	ContratoNumeroContrato string
	Contrato               *ContratoModelo.Contrato `gorm:"unique;not null;foreignkey:numero_contrato;association_foreignkey:contrato_numero_contrato"`
	NombreCompleto         string
}