package ModeloUnidadM

type UnidadMed struct {
	Id                  int `gorm:"PRIMARY_KEY"`
	RegionId            string
	Regiones            *Region `gorm:"foreignkey:RegionId"`
	Clues               string
	ClvPersonal         string
	UiPrei              string
	ClvUbicacion        string
	ClvPresupuestal     string `gorm:"unique;not null"`
	DelegacionesClvDele string
	Delegacion          *Delegaciones `gorm:"foreignkey:DelegacionesClvDele"`
	UniPresupuestal     int
	ClvNivelAtnId       int
	Nivel               *CatalogoNivelAtn `gorm:"foreignkey:ClvNivelAtnId"`
	DenominacionUni     string
	TipoServicio        string
	DescripServicio     string
	NumUnidad           string
	NomUnidad           string
	UbicacionDenom      string
	Direccion           string
	ClvVialidad         int
	Vialidad            *TipoVialidad
	NombVialidad        string
	NumExterior         string
	ClvTipoAsent        int
	Asentamiento        *TipoAsentamiento
	NombAsentmt         string
	Cp                  int
	ClvJurisdiccion     int
	IdJurisdiccion      int
	ClvEntidad          int
	ClvMunicipio        int
	ClvLocalidad        int
	Localidad           string
	Jurisdiccion        *Jurisdiccion
	Latitud             float32
	Longitud            float32
	InicioProduct       string
	Estado              bool
}
type Region struct {
	Id     string `gorm:"PRIMARY_KEY"`
	Nombre string
	Estado bool
}
type RespuestaUM struct {
	Id              string
	DenominacionUni string
}
type Delegaciones struct {
	ClvDele      string `gorm:"PRIMARY_KEY"`
	NombreDele   string
	RelacionDele string
	Tipo         int
	Estado       bool
}
type CatalogoNivelAtn struct {
	ClvNivelAtn   int `gorm:"PRIMARY_KEY"`
	NivelAtencion string
	Estado        bool
}
type TipoVialidad struct {
	ClvVialidad    int `gorm:"PRIMARY_KEY"`
	NombreVialidad string
	Estado         bool
}
type TipoAsentamiento struct {
	ClvTipoAsent int `gorm:"PRIMARY_KEY"`
	TipoAsent    string
	Estado       bool
}
type Jurisdiccion struct {
	IdJurisdiccion int `gorm:"PRIMARY_KEY"`
	NomJurisdic    string
	Estado         bool
}
