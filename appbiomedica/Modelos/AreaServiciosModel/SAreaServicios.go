package AreaServiciosModel

import (
	"appbiomedica/Modelos/CatalogosModel"
	"appbiomedica/Modelos/ModeloUnidadM"
)

type AreaServicioUnidad struct {
	Id                  string `gorm:"PRIMARY_KEY"`
	UnidadMedId         int
	UnidadMed           *ModeloUnidadM.UnidadMed
	ServiciosProformaId string
	ServiciosProforma   *CatalogosModel.ServiciosProforma
	Area                string
	CentroCosto         float64 `gorm:"-"`
	Estado              bool
}
