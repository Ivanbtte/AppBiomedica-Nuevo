package AdminAuxModelo

import (
	"appbiomedica/Modelos/ContratoModel/ContratoModelo"
	"appbiomedica/Modelos/ModeloUnidadM"
	"appbiomedica/Modelos/PersonaModel"
)

type AdminAuxContrato struct {
	Id                     string `gorm:"PRIMARY_KEY"`
	Cargo                  string
	Responsabilidad        string
	PersonaId              string
	Persona                *PersonaModel.Persona
	DelegacionesClvDele    string
	Delegaciones           *ModeloUnidadM.Delegaciones
	ContratoNumeroContrato string
	Contrato               *ContratoModelo.Contrato `gorm:"unique;not null;foreignkey:numero_contrato;association_foreignkey:contrato_numero_contrato"`
	Estado                 bool
}
