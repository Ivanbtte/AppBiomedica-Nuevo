package PermisosContratoModelo

type TipoContrato struct {
	Id   string `gorm:"PRIMARY_KEY"`
	Tipo string
}
type SubTipo struct {
	Id             string `gorm:"PRIMARY_KEY"`
	SubTipo        string
	TipoContratoId string
	TipoContrato   *TipoContrato
	Estado         bool
}
