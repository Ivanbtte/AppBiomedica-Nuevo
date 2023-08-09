package ModeloDistribucionGeografica

type Estados struct {
	Clave     string `gorm:"PRIMARY_KEY"`
	Nombre    string
	Municipio []*Municipios `gorm:"foreignkey:clave_estado"`
	Estado    bool
}
type Municipios struct {
	Clave       string `gorm:"PRIMARY_KEY"`
	Nombre      string
	ClaveEstado string `gorm:"PRIMARY_KEY"`
	Estado      string
}
