package BitacoraModel

import "appbiomedica/Modelos/EstatusModel"

type BitacoraSolicitud struct {
	Id          string `gorm:"PRIMARY_KEY"`
	SolicitudId string
	EstatusId   int
	Estatus     *EstatusModel.Estatus
	Fecha       string
	Hora        string
	Comentario  string
	UsuarioId   string
	Estado      bool
}
