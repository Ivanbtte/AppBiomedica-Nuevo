package CorreoTelefonoModel

type Correos struct {
	Id     string `gorm:"PRIMARY_KEY"`
	Correo string `validate:"email"`
}
