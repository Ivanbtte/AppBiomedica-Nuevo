package CuentasModelo

type Cuentas struct {
	Id                  string `gorm:"PRIMARY_KEY"`
	Naturaleza          string
	Estado              string
	FechaEfectiva       string
	CuentaPrei          string `gorm:"unique;not null"`
	CuentaConac         string
	DescripcionCorta    string
	Descripcion         string
	OrdinarioImssr      string
	UnidOperativasImssr string
	UnidInfoImssr       string
	CCostosImssr        string
	BienestarImssp      string
	UnidOperativasImssp string
	UnidInfoImssp       string
	CCostosImssp        string
	CategoriaCuentaId   string
	CategoriaCuenta     *CategoriaCuenta
	Comentarios         string
}

