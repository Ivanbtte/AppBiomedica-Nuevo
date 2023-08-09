package DepartamentosModel

import (
	"appbiomedica/Modelos/AreaModel"
	"appbiomedica/Modelos/ModeloUnidadM"
)

type OrganigramaUnidad struct {
	Id                string                   `gorm:"PRIMARY_KEY" json:"id,omitempty"`
	NivelJerarquico   int                      `json:"nivelJerarquico,omitempty"`
	DepartamentoId    string                   `json:"departamentoId,omitempty"`
	Departamento      *Departamento            `gorm:"foreignKey:DepartamentoId" json:"departamento,omitempty"`
	AreaId            string                   `json:"areaId,omitempty"`
	Area              *AreaModel.Area          `gorm:"foreignKey:Id" json:"area,omitempty"`
	UnidadMedId       int                      `json:"unidadMedId,omitempty"`
	UnidadMed         *ModeloUnidadM.UnidadMed `gorm:"foreignKey:UnidadMedId" json:"unidadMed,omitempty"`
	OrganigramaParent string                   `gorm:"default:null" json:"organigramaParent,omitempty"`
	Organigramas      []*OrganigramaUnidad     `gorm:"foreignKey:OrganigramaParent" json:"organigramas,omitempty"`
	Estado            bool                     `json:"estado,omitempty"`
}

type OrganigramaDelegacion struct {
	Id                  string                      `gorm:"PRIMARY_KEY" json:"id,omitempty"`
	NivelJerarquico     int                         `json:"nivelJerarquico,omitempty"`
	DepartamentoId      string                      `json:"departamentoId,omitempty"`
	Departamento        *Departamento               `gorm:"foreignKey:DepartamentoId" json:"departamento,omitempty"`
	AreaId              string                      `json:"areaId,omitempty"`
	Area                *AreaModel.Area             `gorm:"foreignKey:Id" json:"area,omitempty"`
	DelegacionesClvDele string                      `json:"delegacionesClvDele,omitempty"`
	Delegaciones        *ModeloUnidadM.Delegaciones `gorm:"foreignKey:DelegacionesClvDele" json:"delegaciones,omitempty"`
	OrganigramaParent   string                      `gorm:"default:null" json:"organigramaParent,omitempty"`
	Organigramas        []*OrganigramaDelegacion    `gorm:"foreignKey:OrganigramaParent" json:"organigramas,omitempty"`
	Estado              bool                        `json:"estado,omitempty"`
}
type OrganigramaUnidadMed struct {
	Id            string                   `gorm:"PRIMARY_KEY" json:"id,omitempty"`
	OrganigramaId string                   `json:"organigramaId,omitempty"`
	Organigrama   *OrganigramaUnidad       `gorm:"foreignKey:Id" json:"organigrama,omitempty"`
	UnidadMedId   int                      `json:"unidadMedId,omitempty"`
	UnidadMed     *ModeloUnidadM.UnidadMed `gorm:"foreignKey:UnidadMedId" json:"unidadMed,omitempty"`
	Estado        bool                     `json:"estado,omitempty"`
}
type DepartamentoPadre struct {
	OrganigramaChildId  string               `json:"organigramaChildId,omitempty"`
	OrganigramaParentId string               `json:"organigramaParentId,omitempty" gorm:"default:null"`
	DepartamentoPadre   []*DepartamentoPadre `gorm:"foreignKey:OrganigramaParentId" json:"departamentoPadre,omitempty"`
	Estado              bool                 `json:"estado,omitempty"`
}

