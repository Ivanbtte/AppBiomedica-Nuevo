package serviciossubrogadosmodel

import (
	"appbiomedica/Modelos/ModeloUnidadM"
	"time"
)

type SolicitudSubrogados struct {
	Id          string                   `gorm:"PRIMARY_KEY"`
	Folio       string                   `gorm:"unique;not null" json:"folio,omitempty"`
	FechaCita   time.Time                `json:"fechaCita,omitempty"`
	UnidadMedId int                      `json:"unidadMedId,omitempty"`
	UnidadMed   *ModeloUnidadM.UnidadMed `json:"unidadMed,omitempty"`
}
