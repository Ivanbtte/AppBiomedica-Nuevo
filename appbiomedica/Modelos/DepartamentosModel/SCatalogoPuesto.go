package DepartamentosModel

import "appbiomedica/Modelos/UsuarioModel"

type CatalogoPuesto struct {
	Id            string `gorm:"PRIMARY_KEY"`
	Puesto        string
	UsuarioTipoId string `validate:"required"`
	UsuarioTipo   *UsuarioModel.UsuarioTipo
	Estado        bool
}

