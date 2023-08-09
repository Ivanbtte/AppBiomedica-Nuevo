package Init

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/ProveedoresModel"
)

func InitProveedor() error {
	db := Conexion.GetDB()
	if !db.HasTable(&ProveedoresModel.Proveedor{}) {
		if err := db.CreateTable(&ProveedoresModel.Proveedor{}).
			AddForeignKey("estados_id", "estados(clave)", "RESTRICT", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	if !db.HasTable(&ProveedoresModel.ContactoProveedor{}) {
		if err := db.CreateTable(&ProveedoresModel.ContactoProveedor{}).
			AddForeignKey("proveedor_n_prov_imss", "proveedor(n_prov_imss)", "CASCADE", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	return nil
}
