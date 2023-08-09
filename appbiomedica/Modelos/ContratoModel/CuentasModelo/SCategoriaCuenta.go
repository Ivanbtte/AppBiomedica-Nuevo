package CuentasModelo

type CategoriaCuenta struct {
	Id        string `gorm:"PRIMARY_KEY"`
	Categoria string
}
