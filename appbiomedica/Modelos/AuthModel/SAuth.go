package AuthModel

import (
	"appbiomedica/Modelos/UsuarioModel"
)

type Credenciales struct {
	ID        string
	UsuarioID string
	Usuario   *UsuarioModel.Usuario
	Correo    string `json:"correo"`
	Password  string `json:"password"`
	Estado    bool
}
