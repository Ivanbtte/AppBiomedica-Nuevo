package FianzaModelo

import "appbiomedica/Modelos/ContratoModel/ContratoModelo"

type Fianza struct {
	Id                     string `gorm:"PRIMARY_KEY"`
	Afianzadora            string
	NumPoliza              string
	Tipo                   string
	MontoFianza            float32 `json:"MontoFianza,omitempty";gorm:"-";gorm:"default:0.0"`
	ContratoNumeroContrato string
	Contrato               *ContratoModelo.Contrato `gorm:"unique;not null;foreignkey:numero_contrato;association_foreignkey:contrato_numero_contrato"`
	Estado                 bool
}
