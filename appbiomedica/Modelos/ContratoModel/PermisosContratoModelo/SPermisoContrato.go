package PermisosContratoModelo

type PermisoContrato struct {
	Id      string `gorm:"PRIMARY_KEY"`
	Permiso string
}
