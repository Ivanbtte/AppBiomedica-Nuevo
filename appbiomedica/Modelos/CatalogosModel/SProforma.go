package CatalogosModel

/* *********************    Estructura de catalogo proforma     ********************* */
type Proforma struct {
	Id             int `gorm:"PRIMARY_KEY"`
	PreiIdArticulo string
	Pre            *Prei `gorm:"unique;not null;foreignkey:id_articulo;association_foreignkey:prei_id_articulo"`
	IdServicio     string
	ServiciosPro   *ServiciosProforma `gorm:"foreignkey:id_servicio"`
	NivelAtn       string
	Descripcion    string
	ClaveCB        string
	Estado         bool
}

/* *********************   Estructura de los servicos proforma     ********************* */
type ServiciosProforma struct {
	Id          string `gorm:"PRIMARY_KEY"`
	Descripcion string
	Estado      bool
}
