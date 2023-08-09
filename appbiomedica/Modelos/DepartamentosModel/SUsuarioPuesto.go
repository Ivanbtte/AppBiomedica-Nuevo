package DepartamentosModel

import (
	"appbiomedica/Modelos/RoleModel"
	"appbiomedica/Modelos/UsuarioModel"
	"time"
)

type UsuarioPuesto struct {
	Id                  string `gorm:"PRIMARY_KEY"`
	UsuarioId           string
	Usuario             *UsuarioModel.Usuario
	PuestoOrganigramaId string
	PuestoOrganigrama   *PuestoOrganigrama `gorm:"foreignKey:PuestoOrganigramaId"`
	FechaRegistro       time.Time
	Valid               bool
	Estado              bool
}

type UsuarioPuestoRole struct {
	Id              string          `gorm:"PRIMARY_KEY" json:"id,omitempty"`
	UsuarioPuestoId string          `json:"usuarioPuestoId,omitempty"`
	UsuarioPuesto   *UsuarioPuesto  `gorm:"foreignKey:UsuarioPuestoId" json:"usuarioPuesto,omitempty"`
	RoleId          string          `json:"roleId,omitempty"`
	Role            *RoleModel.Role `gorm:"foreignKey:RoleId" json:"role,omitempty"`
	Status          bool            `json:"status,omitempty"`
}
type BitacoraUsuarioPuesto struct {
	Id                   string           `gorm:"PRIMARY_KEY" json:"id,omitempty"`
	UsuarioModificadoId  string           `json:"usuarioModificadoId,omitempty"`
	UsuarioModificado    *UsuarioPuesto   `gorm:"foreignKey:UsuarioModificadoId" json:"usuarioModificado,omitempty"`
	UsuarioModificanteId string           `json:"usuarioModificanteId,omitempty"`
	UsuarioModificante   *UsuarioPuesto   `gorm:"foreignKey:UsuarioModificanteId" json:"usuarioModificante,omitempty"`
	TipoModificacionId   string           `json:"tipoModificacionId,omitempty"`
	TipoModificacion     TipoModificacion `json:"tipoModificacion,omitempty"`
	Fecha                time.Time        `json:"fecha,omitempty"`
	Comentario           string           `json:"comentario,omitempty"`
	Status               bool             `json:"status,omitempty"`
}
type TipoModificacion struct {
	Id          string `gorm:"PRIMARY_KEY" json:"id,omitempty"`
	Tipo        string `json:"tipo,omitempty"`
	Descripcion string `json:"descripcion,omitempty"`
	Estado      bool   `json:"estado,omitempty"`
}
