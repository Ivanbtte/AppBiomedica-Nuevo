package PersonaModel

type Persona struct {
	Id          string `gorm:"PRIMARY_KEY"`
	Nombre      string `validate:"required"`
	ApellidoPat string `validate:"required"`
	ApellidoMat string
	FechaNac    string `validate:"required"`
	Genero      string `validate:"required"`
	Correo      string `validate:"required"`
	Telefono    string
	Extension   int
	Celular     string
	Matricula   string `validate:"required"`
	Estado      bool
}

