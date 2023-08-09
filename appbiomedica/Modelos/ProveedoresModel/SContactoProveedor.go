package ProveedoresModel

type ContactoProveedor struct {
	Id                 string `gorm:"PRIMARY_KEY"`
	NombreCompleto     string
	Departamento       string
	Cargo              string
	Telefono           string
	Extension          string
	Celular            string
	Correo             string
	Comentarios        string
	Asunto             string
	ProveedorNProvImss string
	Proveedor          *Proveedor `gorm:"unique;not null;foreignkey:n_prov_imss;association_foreignkey:proveedor_n_prov_imss"`
	Estado             bool
}
