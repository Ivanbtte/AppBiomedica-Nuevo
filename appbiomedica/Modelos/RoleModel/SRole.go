package RoleModel

import (
	"appbiomedica/Modelos/ModuloModel"
)

type Role struct {
	Id          string `gorm:"primaryKey;unique;not null" json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Hierarchy   int    `json:"hierarchy,omitempty"`
	Status      bool   `json:"status"`
}
type RoleModule struct {
	Id       string              `gorm:"primaryKey;unique;not null" json:"id,omitempty"`
	RoleId   string              `json:"roleId,omitempty"`
	Role     *Role               `gorm:"foreignKey:RoleId" json:"role,omitempty"`
	ModuloId string              `json:"moduloId,omitempty"`
	Modulo   *ModuloModel.Modulo `gorm:"foreignKey:ModuloId" json:"modulo,omitempty"`
	Status   bool                `json:"status"`
}
type RoleModuleAction struct {
	Id           string              `gorm:"primaryKey;unique;not null" json:"id,omitempty"`
	RoleModuleId string              `json:"roleModuleId,omitempty"`
	RoleModule   *RoleModule         `gorm:"foreignKey:RoleModuleId" json:"roleModule,omitempty"`
	ActionId     string              `json:"actionId,omitempty"`
	Action       *ModuloModel.Action `gorm:"foreignKey:ActionId" json:"action,omitempty"`
	Status       bool                `json:"status"`
}
