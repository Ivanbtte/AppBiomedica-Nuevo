package ServiciosUnidadesModel

import (
	"appbiomedica/Modelos/CatalogosModel"
	"appbiomedica/Modelos/ModeloUnidadM"
)

type ServiciosPorUnidad struct {
	IdUnidadMed         int
	UnidadMed           *ModeloUnidadM.UnidadMed `gorm:"foreignkey:id;association_foreignkey:id_unidad_med"`
	IdServiciosProforma string
	ServiciosProforma   *CatalogosModel.ServiciosProforma `gorm:"foreignkey:id;association_foreignkey:id_servicios_proforma"`
}
