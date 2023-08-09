package TrasladoPacientesModel

import "appbiomedica/Modelos/ContratoModel/ConceptosServiciosModelo"

type RutasConsumidas struct {
	Id                  string                                      `gorm:"primaryKey;unique;not null" json:"id,omitempty"`
	Cantidad            int                                         `json:"cantidad,omitempty"`
	BoletosId           string                                      `json:"boletosId,omitempty"`
	Boletos             *Boletos                                    `json:"boletos,omitempty"`
	TrasladoPacientesId string                                      `json:"trasladoPacientesId,omitempty"`
	TrasladoPacientes   *ConceptosServiciosModelo.TrasladoPacientes `json:"trasladoPacientes,omitempty"`
	Estado              bool                                        `json:"estado,omitempty"`
}
type Boletos struct {
	Id                     string             `gorm:"primaryKey;unique;not null" json:"id,omitempty"`
	Ida                    bool               `json:"ida,omitempty"`
	Regreso                bool               `json:"regreso,omitempty"`
	Total                  int                `json:"total,omitempty"`
	Observaciones          string             `json:"observaciones,omitempty"`
	Tipo                   string             `json:"tipo,omitempty"`
	TipoPasajeroId         int                `json:"tipoPasajeroId,omitempty"`
	TipoPasajero           *TipoPasajero      `json:"tipoPasajero,omitempty"`
	SolicitudTrasladoFolio string             `json:"solicitudTrasladoFolio,omitempty"`
	SolicitudTraslado      *SolicitudTraslado `gorm:"foreignKey:Folio" json:"solicitudTraslado,omitempty"`
	Estado                 bool               `json:"estado,omitempty"`
}
type TipoPasajero struct {
	Id   uint   `gorm:"primaryKey;unique;not null;autoIncrement" json:"id,omitempty"`
	Tipo string `json:"tipo,omitempty"`
}
