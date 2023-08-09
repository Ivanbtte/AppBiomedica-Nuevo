package ProveedoresModel

import "appbiomedica/Modelos/ModeloDistribucionGeografica"

type Proveedor struct {
	Id            string `gorm:"PRIMARY_KEY"`
	AliasEmpresa  string
	NombreEmpresa string
	Correo        string
	RFC           string
	NProvImss     string `gorm:"unique;not null"`
	Telefono      string
	Direccion     string
	EstadosId     string
	Estados       *ModeloDistribucionGeografica.Estados `gorm:"foreignkey:clave;association_foreignkey:estados_id"`
	Municipio     string
	CP            string
	Estado        bool
}
