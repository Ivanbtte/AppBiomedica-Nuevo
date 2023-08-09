package SolicitudModel

import "appbiomedica/Modelos/CatalogosModel"

type ClavesSolicitud struct {
	Id                  string `gorm:"PRIMARY_KEY"`
	SolicitudId         string
	Solicitud           *Solicitud
	ClaveId             string
	Descripcion         string
	CantidadSolicitada  int
	CantidadRecibida    int
	CantidadAprobada    int
	CantidadEntregada   int
	Precio              float64
	Total               float64
	ServiciosProformaId string
	ServiciosProforma   *CatalogosModel.ServiciosProforma
	EstatusId           int
	Fecha               string
	Estado              bool
}
