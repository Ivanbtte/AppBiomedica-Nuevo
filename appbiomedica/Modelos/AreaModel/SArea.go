package AreaModel

type Area struct {
	Id             string `gorm:"primaryKey;unique;not null" json:"id,omitempty"`
	Nombre         string `json:"nombre,omitempty"`
	OrdenJerarquia int    `json:"ordenJerarquia,omitempty"`
	Estado         bool   `json:"estado,omitempty"`
}
