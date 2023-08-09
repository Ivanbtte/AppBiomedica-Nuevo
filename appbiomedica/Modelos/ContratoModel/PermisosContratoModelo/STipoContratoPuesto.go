package PermisosContratoModelo

import (
	"appbiomedica/Modelos/DepartamentosModel"
)

type TipoContratoPuesto struct {
	Id                string `gorm:"PRIMARY_KEY"`
	CatalogoPuestoId  string
	CatalogoPuesto    *DepartamentosModel.CatalogoPuesto
	TipoContratoId    string
	TipoContrato      *TipoContrato
	PermisoContratoId string
	PermisoContrato   *PermisoContrato
}
