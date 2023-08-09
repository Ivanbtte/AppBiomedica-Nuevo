package MenuModel

import (
	"appbiomedica/Modelos/ModuloModel"
)

type Menu struct {
	Id         string              `gorm:"PRIMARY_KEY" json:"id,omitempty"`
	State      string              `json:"state,omitempty"`
	Name       string              `json:"name,omitempty"`
	Type       string              `json:"type,omitempty"`
	Icon       string              `json:"icon,omitempty"`
	MenuParent string              `gorm:"default:null" json:"menuParent,omitempty"`
	Children   []Menu              `gorm:"foreignKey:MenuParent" json:"children,omitempty"`
	ModuloId   string              `json:"moduloId,omitempty"`
	Modulo     *ModuloModel.Modulo `json:"modulo,omitempty"`
	Orden      int                 `json:"orden,omitempty"`
	Status     bool                `json:"status,omitempty"`
}