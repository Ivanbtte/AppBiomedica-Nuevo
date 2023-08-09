package ContratoModelo

import (
	"appbiomedica/Modelos/ContratoModel/PermisosContratoModelo"
	"appbiomedica/Modelos/ModeloUnidadM"
	"appbiomedica/Modelos/ProveedoresModel"
)

type Contrato struct {
	Id                     string `gorm:"PRIMARY_KEY"`
	NumeroContrato         string `gorm:"unique;not null"`
	InicioContrato         string
	FinContrato            string
	ProcedContratacion     string
	TipoProcedContratacion string
	FechaFallo             string
	MontoTotal             float64
	TipoContratoId         string
	TipoContrato           *PermisosContratoModelo.TipoContrato
	ProveedorNProvImss     string
	Proveedor              *ProveedoresModel.Proveedor `gorm:"unique;not null;foreignkey:n_prov_imss;association_foreignkey:proveedor_n_prov_imss"`
	Estado                 bool                        `gorm:"default:true"`
}
type Contrato_Delegaciones struct {
	ContratoNumeroContrato string
	Contrato               *Contrato `gorm:"unique;not null;foreignkey:numero_contrato;association_foreignkey:contrato_numero_contrato"`
	DelegacionesClvDele    string
	Delegaciones           *ModeloUnidadM.Delegaciones
}
type RespuestaContrato struct {
	NumContra string
	Error     string
	Columna   string
	Linea     int
	Tipo      int
}

/*****  Estructura para relacionar un subtipo de contrato con un contrato  *****/
type SubTipoContrato struct {
	SubTipoId              string
	SubTipo                *PermisosContratoModelo.SubTipo
	ContratoNumeroContrato string
	Contrato               *Contrato `gorm:"unique;not null;foreignkey:numero_contrato;association_foreignkey:contrato_numero_contrato"`
	Estado                 bool
}
