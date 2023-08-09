package FtpModelo

type Ftp struct {
	FolioTp         string  `gorm:"primaryKey;unique;not null" json:"folioTp,omitempty"`
	TipoFolio       string  `json:"tipoFolio,omitempty"`
	UnidadReceptora string  `json:"unidadReceptora,omitempty"`
	Especialidad    string  `json:"especialidad,omitempty"`
	FechaHora       string  `json:"fechaHora,omitempty"`
	Tipocita        string  `json:"tipocita,omitempty"`
	Nss             string  `json:"nss,omitempty"`
	Agregado        string  `json:"agregado,omitempty"`
	Paciente        string  `json:"paciente,omitempty"`
	Acompañante     string  `json:"acompañante,omitempty"`
	Transporte      string  `json:"transporte,omitempty"`
	UsuarioEmitio   string  `json:"usuarioEmitio,omitempty"`
	MotivoAjuste    string  `json:"motivoAjuste,omitempty"`
	FechaPago       string  `json:"fechaPago,omitempty"`
	ImporteEmitido  float64 `json:"importeEmitido,omitempty"`
	ImportePagado   float64 `json:"importePagado,omitempty"`
	EstadoTp        string  `json:"estadoTp,omitempty"`
}
