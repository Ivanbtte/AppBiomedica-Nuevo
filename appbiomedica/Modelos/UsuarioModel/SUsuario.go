package UsuarioModel

import (
	"appbiomedica/Modelos/PersonaModel"
	"time"
)

type Usuario struct {
	Id            string `gorm:"PRIMARY_KEY"`
	Correo        string
	Contraseña    string `validate:"required"`
	UsuarioTipoId string `validate:"required"`
	UsuarioTipo   *UsuarioTipo
	PersonaId     string
	Persona       *PersonaModel.Persona
	FechaRegistro time.Time
	Estado        bool
}
type Passwords struct {
	Password    string
	NewPassword string
}

type UsuarioTipo struct {
	Id     string `gorm:"PRIMARY_KEY"`
	Tipo   string
	Estado bool
}

