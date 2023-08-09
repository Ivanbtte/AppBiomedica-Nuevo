package CatalogosModel

type Prei struct {
	Id                int `gorm:"PRIMARY_KEY"`
	IdArticulo        string `gorm:"unique;not null"`
	Grupo             string
	Generico          string
	Especifico        string
	Diferenciador     string
	Variable          string
	Descripcion       string
	Precio            float32
	UnidadesMedidasId int
	UnidadesMedidas   *UnidadesMedidas
	CuadroBasicoId    int
	CuadroBasico      *CuadroBasico
	Estado            bool
}
type UnidadesMedidas struct {
	Id     int `gorm:"PRIMARY_KEY"`
	Unidad string
	Estado bool
}
type CuadroBasico struct {
	Id     int `gorm:"PRIMARY_KEY"`
	Grupo  string
	Estado bool
}
