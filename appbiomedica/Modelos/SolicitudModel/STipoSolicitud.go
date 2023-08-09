package SolicitudModel

type TipoSolicitud struct {
	Id   int `gorm:"PRIMARY_KEY"`
	Tipo string
}
