package DepartamentosModel

import (
	"appbiomedica/Modelos/UsuarioModel"
)

/* ** Este modelo nos servira para almacenar si una jefatura o coordinacion depende de alguna otra ** */

type Departamento struct {
	Id                 string `gorm:"PRIMARY_KEY"`
	Nombre             string
	DepartamentoTipoId string
	DepartamentoTipo   *DepartamentoTipo
	UsuarioTipoId      string `validate:"required"`
	UsuarioTipo        *UsuarioModel.UsuarioTipo
	Estado             bool
}


/*****  Esta estructura nos servira para saber que tipo de departamento agregaremos  *****/

type DepartamentoTipo struct {
	Id     string
	Tipo   string
	Estado bool
}
