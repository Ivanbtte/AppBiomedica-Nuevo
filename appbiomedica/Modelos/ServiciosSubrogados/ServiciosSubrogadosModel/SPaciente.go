package serviciossubrogadosmodel

type Paciente struct {
	Id        string `gorm:"primaryKey;unique;not null" json:"id,omitempty"`
	Nombre    string `json:"nombre,omitempty"`
	Apellidos string `json:"apellidos,omitempty"`
	Sexo      string `json:"sexo,omitempty"`
	Nss       int    `json:"nss,omitempty"`
	Agregado  string `json:"agregado,omitempty"`
	Curp      string `json:"curp,omitempty"`
}
