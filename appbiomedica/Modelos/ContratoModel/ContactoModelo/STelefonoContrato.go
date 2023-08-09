package ContactoModelo

import (
	"appbiomedica/Modelos/ContratoModel/ContratoModelo"
	"appbiomedica/Modelos/CorreoTelefonoModel"
)

type TelefonoContrato struct {
	Id                     string `gorm:"PRIMARY_KEY"`
	TipoTelefono           string `gorm:"DEFAULT:'Personal'"`
	TelefonosId            string
	Telefonos              *CorreoTelefonoModel.Telefonos
	Extension              string
	ContratoNumeroContrato string
	Contrato               *ContratoModelo.Contrato `gorm:"unique;not null;foreignkey:numero_contrato;association_foreignkey:contrato_numero_contrato"`
}
