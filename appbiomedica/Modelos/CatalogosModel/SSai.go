package CatalogosModel

type Sai struct {
	Id                   int    `gorm:"PRIMARY_KEY"`
	Grupo                string `gorm:"index:grupo"`
	Generico             string
	Especifico           string
	Diferenciador        string
	Variable             string
	CuadroBasico         string
	DescripcionGrupoId   int
	DescripcionGrupo     *DescripcionGrupo
	Descripcion          string
	UnidadPresentacionId int
	UnidadPresentacion   *UnidadPresentacion
	CantidadPresentacion float32
	TipoPresentacionId   int
	TipoPresentacion     *TipoPresentacion
	Precio               float64
	PartidaPresupuestal  int
	FechaRegistro        string
	Inventariables       string
	NivelCompra          int
	Linea                int
	NumRegistro          int
	Estado               bool
}
type TipoPresentacion struct {
	Id           int `gorm:"PRIMARY_KEY"`
	Presentacion string
	Estado       bool
}
type UnidadPresentacion struct {
	Id                 int `gorm:"PRIMARY_KEY"`
	UnidadPresentacion string
	Estado             bool
}
type DescripcionGrupo struct {
	Id          int `gorm:"PRIMARY_KEY"`
	Descripcion string
	Estado      bool
}
