package TrasladoPacientesModel

import (
	"appbiomedica/Modelos/ModeloUnidadM"
	"appbiomedica/Modelos/UsuarioModel"
	"time"
)

type FirmaTraslado struct {
	Id                     string             `json:"id,omitempty"`
	FirmaSolicitudId       string             `json:"firmaSolicitudId,omitempty"`
	FirmaSolicitud         *FirmaSolicitud    `json:"firmaSolicitud,omitempty"`
	SolicitudTrasladoFolio string             `json:"solicitudTrasladoFolio,omitempty"`
	SolicitudTraslado      *SolicitudTraslado `gorm:"foreignKey:Folio" json:"solicitudTraslado,omitempty"`
	Estado                 bool               `json:"estado,omitempty"`
}
type FirmaSolicitud struct {
	Id          string                   `gorm:"primaryKey;unique;not null" json:"id,omitempty"`
	Nombre      string                   `json:"nombre,omitempty"`
	Cargo       string                   `json:"cargo,omitempty"`
	Matricula   string                   `json:"matricula,omitempty"`
	Tipo        int                      `json:"tipo,omitempty"`
	Fecha       time.Time                `json:"fecha,omitempty"`
	UnidadMedId int                      `json:"unidadMedId,omitempty"`
	UnidadMed   *ModeloUnidadM.UnidadMed `json:"unidadMed,omitempty"`
	UsuarioId   string                   `json:"usuarioId,omitempty"`
	Usuario     *UsuarioModel.Usuario    `json:"usuario,omitempty"`
	Estado      bool                     `json:"estado,omitempty"`
}
type RecordUpdateFirm struct {
	Id               string                   `gorm:"primaryKey;unique;not null" json:"id,omitempty"`
	Nombre           string                   `json:"nombre,omitempty"`
	Cargo            string                   `json:"cargo,omitempty"`
	Matricula        string                   `json:"matricula,omitempty"`
	Tipo             int                      `json:"tipo,omitempty"`
	Fecha            time.Time                `json:"fecha,omitempty"`
	UsuarioId        string                   `json:"usuarioId,omitempty"`
	Usuario          *UsuarioModel.Usuario    `json:"usuario,omitempty"`
	UnidadMedId      int                      `json:"unidadMedId,omitempty"`
	UnidadMed        *ModeloUnidadM.UnidadMed `json:"unidadMed,omitempty"`
	FirmaSolicitudId string                   `json:"firmaSolicitudId,omitempty"`
	Estado           bool                     `json:"estado,omitempty"`
}
type RecordDisableFirm struct {
	Id               string                   `gorm:"primaryKey;unique;not null" json:"id,omitempty"`
	Opcion           string                   `json:"opcion,omitempty"`
	Fecha            time.Time                `json:"fecha,omitempty"`
	UsuarioId        string                   `json:"usuarioId,omitempty"`
	Usuario          *UsuarioModel.Usuario    `json:"usuario,omitempty"`
	UnidadMedId      int                      `json:"unidadMedId,omitempty"`
	UnidadMed        *ModeloUnidadM.UnidadMed `json:"unidadMed,omitempty"`
	FirmaSolicitudId string                   `json:"firmaSolicitudId,omitempty"`
	Estado           bool                     `json:"estado,omitempty"`
}
