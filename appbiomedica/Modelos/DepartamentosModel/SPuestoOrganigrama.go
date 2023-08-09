package DepartamentosModel

import (
	"appbiomedica/Modelos/ModuloModel"
	"appbiomedica/Modelos/RoleModel"
)

type PuestoOrganigrama struct {
	Id                      string                 `gorm:"PRIMARY_KEY" json:"id,omitempty"`
	PuestoJefe              bool                   `json:"puestoJefe"`
	NumeroPlazas            int                    `json:"numeroPlazas,omitempty"`
	OrganigramaUnidadId     string                 `gorm:"default:null" json:"organigramaUnidadId"`
	OrganigramaUnidad       *OrganigramaUnidad     `gorm:"foreignKey:OrganigramaUnidadId" json:"organigramaUnidad"`
	OrganigramaDelegacionId string                 `gorm:"default:null" json:"organigramaDelegacionId"`
	OrganigramaDelegacion   *OrganigramaDelegacion `gorm:"foreignKey:OrganigramaDelegacionId" json:"organigramaDelegacion"`
	CatalogoPuestoId        string                 `json:"catalogoPuestoId,omitempty"`
	CatalogoPuesto          *CatalogoPuesto        `gorm:"foreignKey:CatalogoPuestoId" json:"catalogoPuesto,omitempty"`
	Estado                  bool                   `json:"estado,omitempty"`
}

type PuestoOrganigramaRole struct {
	Id                  string             `gorm:"PRIMARY_KEY" json:"id,omitempty"`
	PuestoOrganigramaId string             `json:"puestoOrganigramaId,omitempty"`
	PuestoOrganigrama   *PuestoOrganigrama `gorm:"foreignKey:PuestoOrganigramaId" json:"puestoOrganigrama,omitempty"`
	RoleId              string             `json:"roleId,omitempty"`
	Role                *RoleModel.Role    `gorm:"foreignKey:RoleId" json:"role,omitempty"`
	Status              bool               `json:"status,omitempty"`
}
type ActionExceptionByUser struct {
	Id              string              `gorm:"primaryKey;unique;not null" json:"id,omitempty"`
	ActionId        string              `json:"actionId,omitempty"`
	Action          *ModuloModel.Action `gorm:"foreignKey:ActionId" json:"action,omitempty"`
	UsuarioPuestoId string              `json:"usuarioPuestoId,omitempty"`
	UsuarioPuesto   *UsuarioPuesto      `gorm:"foreignKey:UsuarioPuestoId" json:"usuarioPuesto,omitempty"`
	Status          bool                `json:"status"`
}

// type Turno struct {
// 	Clave       int    `gorm:"primaryKey;unique;not null" json:"clave,omitempty"`
// 	Descripcion string `json:"descripcion,omitempty"`
// 	Estado      bool   `json:"estado,omitempty"`
// }

// type PuestoOrganigramaTurno struct {
// 	Id                  string             `gorm:"primaryKey;unique;not null" json:"id,omitempty"`
// 	PuestoOrganigramaId string             `json:"puestoOrganigramaId,omitempty"`
// 	PuestoOrganigrama   *PuestoOrganigrama `gorm:"foreignKey:PuestoOrganigramaId" json:"puestoOrganigrama,omitempty"`
// 	TurnoClave          int                `json:"turnoClave,omitempty"`
// 	Turno               Turno              `gorm:"foreignKey:TurnoClave" json:"turno,omitempty"`
// 	Estado              bool               `json:"estado,omitempty"`
// }	

