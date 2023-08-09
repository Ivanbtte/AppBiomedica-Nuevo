package TrasladoPacientesModel

import (
	"appbiomedica/Modelos/ModeloUnidadM"
	"appbiomedica/Modelos/ModuloModel"
	"appbiomedica/Modelos/UsuarioModel"
	"time"
)

type SolicitudTraslado struct {
	Id                     uint                     `gorm:"primaryKey;unique;not null;autoIncrement" json:"id,omitempty"`
	Folio                  string                   `gorm:"unique;not null" json:"folio,omitempty"`
	Version                int                      `json:"version,omitempty"`
	RepresentanteLegal     string                   `json:"representanteLegal,omitempty"`
	NombrePaciente         string                   `json:"nombrePaciente,omitempty"`
	Nss                    string                   `json:"nss,omitempty"`
	AgregadoMedico         string                   `json:"agregadoMedico,omitempty"`
	FechaCita              time.Time                `json:"fechaCita,omitempty"`
	Origen                 string                   `json:"origen,omitempty"`
	Destino                string                   `json:"destino,omitempty"`
	Acompañante            bool                     `json:"acompañante"`
	Ftp01                  string                   `json:"ftp01,omitempty"`
	FechaExpedicion        time.Time                `json:"fechaExpedicion,omitempty"`
	ContratoNumeroContrato string                   `json:"contratoNumeroContrato,omitempty"`
	UnidadMedId            int                      `json:"unidadMedId,omitempty"`
	UnidadMed              *ModeloUnidadM.UnidadMed `json:"unidadMed,omitempty"`
	UsuarioId              string                   `json:"usuarioId,omitempty"`
	Usuario                *UsuarioModel.Usuario    `json:"usuario,omitempty"`
	Estado                 bool                     `json:"estado,omitempty"`
}
type Acompaniante struct {
	Id                     string             `gorm:"primaryKey;unique;not null" json:"id,omitempty"`
	Motivo                 string             `json:"motivo,omitempty"`
	Justificacion          string             `json:"justificacion,omitempty"`
	NombreAutoriza         string             `json:"nombreAutoriza,omitempty"`
	Cargo                  string             `json:"cargo,omitempty"`
	Matricula              string             `json:"matricula,omitempty"`
	TarjetaInapan          bool               `json:"tarjetaInapan,omitempty"`
	SolicitudTrasladoFolio string             `json:"solicitudTrasladoFolio,omitempty"`
	SolicitudTraslado      *SolicitudTraslado `gorm:"foreignKey:Folio" json:"solicitudTraslado,omitempty"`
}
type SolicitudesCanceladas struct {
	Id                     string                `gorm:"primaryKey;unique;not null" json:"id,omitempty"`
	Motivo                 string                `json:"motivo,omitempty"`
	FechaCancelacion       time.Time             `json:"fechaCancelacion,omitempty"`
	SolicitudTrasladoFolio string                `json:"solicitudTrasladoFolio,omitempty"`
	SolicitudTraslado      *SolicitudTraslado    `gorm:"foreignKey:Folio" json:"solicitudTraslado,omitempty"`
	UsuarioId              string                `json:"usuarioId,omitempty"`
	Usuario                *UsuarioModel.Usuario `json:"usuario,omitempty"`
	Estado                 bool                  `json:"estado,omitempty"`
}
type AutorizacionesEspeciales struct {
	Id                     string             `gorm:"primaryKey;unique;not null" json:"id,omitempty"`
	Justificacion          string             `json:"justificacion,omitempty"`
	TipoAutorizacion       string             `json:"tipoAutorizacion,omitempty"`
	SolicitudTrasladoFolio string             `json:"solicitudTrasladoFolio,omitempty"`
	SolicitudTraslado      *SolicitudTraslado `gorm:"foreignKey:Folio" json:"solicitudTraslado,omitempty"`
	Estado                 bool               `json:"estado,omitempty"`
}
type PermisosPorUnidad struct {
	Id          string                   `gorm:"primaryKey;unique;not null" json:"id,omitempty"`
	ActionId    string                   `json:"actionId,omitempty"`
	Action      *ModuloModel.Action      `json:"action,omitempty"`
	UnidadMedId int                      `json:"unidadMedId,omitempty"`
	UnidadMed   *ModeloUnidadM.UnidadMed `json:"unidadMed,omitempty"`
	Estado      bool                     `json:"estado,omitempty"`
}
