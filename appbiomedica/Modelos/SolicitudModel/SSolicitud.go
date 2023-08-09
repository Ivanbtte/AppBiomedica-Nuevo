package SolicitudModel

import (
	"appbiomedica/Modelos/EstatusModel"
	"appbiomedica/Modelos/ModeloUnidadM"
	"appbiomedica/Modelos/UsuarioModel"
)

type Solicitud struct {
	Id              string `gorm:"PRIMARY_KEY"`
	UsuarioId       string
	Usuario         *UsuarioModel.Usuario
	UnidadMedId     int
	UnidadMed       *ModeloUnidadM.UnidadMed
	TipoSolicitudId int
	TipoSolicitud   *TipoSolicitud
	EstatusId       int
	Estatus         *EstatusModel.Estatus
	FechaCreacion   string
	Total           float64
	Correo          string
	Estado          bool
}
