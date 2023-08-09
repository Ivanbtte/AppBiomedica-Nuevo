package BitacoraModel

import "appbiomedica/Modelos/UsuarioModel"

type BitacoraClave struct {
	Id         string `gorm:"PRIMARY_KEY"`
	Accion     string
	ClaveId    string
	EstatusId  int
	Fecha      string
	Hora       string
	Comentario string
	UsuarioId  string
	Usuario    *UsuarioModel.Usuario
	Estado     bool
}
