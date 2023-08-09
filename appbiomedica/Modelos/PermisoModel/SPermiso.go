package PermisoModel

type Permiso struct {
	Id          string `gorm:"PRIMARY_KEY" json:"id,omitempty"`
	Descripcion string `validate:"required" json:"descripcion,omitempty"`
	Prioridad   int    `json:"prioridad,omitempty"`
	Estado      bool   `json:"estado,omitempty"`
}
