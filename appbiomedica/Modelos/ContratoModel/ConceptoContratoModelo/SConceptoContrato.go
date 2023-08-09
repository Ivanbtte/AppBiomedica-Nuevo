package ConceptoContratoModelo

import (
	"appbiomedica/Modelos/CatalogosModel"
	"appbiomedica/Modelos/ContratoModel/ContratoModelo"
)

type ConceptoContrato struct {
	Id                     string `gorm:"PRIMARY_KEY"`
	PrecioUniSnIva         float64
	CantidadConcepto       int
	CantidadAmpliada       int
	CantidadDistribuida    int
	Marca                  string
	Modelo                 string
	ObjetoContratacion     string
	FechaMaxEntrega        string
	GarantiaBienes         int
	PreiIdArticulo         string
	Prei                   *CatalogosModel.Prei `gorm:"unique;not null;foreignkey:id_articulo;association_foreignkey:prei_id_articulo"`
	ContratoNumeroContrato string
	Contrato               *ContratoModelo.Contrato `gorm:"unique;not null;foreignkey:numero_contrato;association_foreignkey:contrato_numero_contrato"`
}
type ConceptoContratoTemp struct {
	PrecioUniSnIva     float32
	CantidadConcepto   int
	Marca              string
	Modelo             string
	ObjetoContratacion string
	FechaMaxEntrega    string
	GarantiaBienes     int
	PreiIdArticulo     string
}
type RespuestaArchivo struct {
	NumContrato string
	IdPrei      string
	Error       string
	Columna     string
	Fila        int
	Suma        float64
	Tipo        int
}
type RespuestaDistribucion struct {
	NumContrato string
	IdPrei      string
	Error       string
	Columna     string
	Fila        int
	Tipo        int
	Cantidad    int
	Suma        float64
	Concepto    *ConceptoContrato
}
