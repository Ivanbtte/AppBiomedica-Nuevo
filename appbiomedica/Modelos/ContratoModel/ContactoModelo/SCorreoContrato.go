package ContactoModelo

import (
	"appbiomedica/Modelos/ContratoModel/ContratoModelo"
	"appbiomedica/Modelos/CorreoTelefonoModel"
)

type CorreosContrato struct {
	Id                     string `gorm:"PRIMARY_KEY"`
	TipoCorreo             string `gorm:"DEFAULT:'Personal'"`
	CorreosId              string
	Correos                *CorreoTelefonoModel.Correos
	ContratoNumeroContrato string
	Contrato               *ContratoModelo.Contrato `gorm:"unique;not null;foreignkey:numero_contrato;association_foreignkey:contrato_numero_contrato"`
}
