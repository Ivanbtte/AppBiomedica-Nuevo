package ConceptosServiciosModelo

import (
	"appbiomedica/Modelos/CatalogosModel"
	"appbiomedica/Modelos/ContratoModel/ContratoModelo"
	"appbiomedica/Modelos/ModeloUnidadM"
	"io"

	"github.com/jinzhu/gorm"
)

type Concepto interface {
	AgregarConcepto(tx *gorm.DB) error
	ValidaConceptosArchivo(file io.Reader) (bool, []RespuestConceptoArchivo, error)
	DefinirEstructura(datos []string, id_concepto string) Concepto
	ConsultarConceptos(num_contrato string, numPag, numReg int) (error, *[]Conceptos, float64, int)
}

// `gorm:"ForeignKey:RecordId;AssociationForeignKey:sys_id"`
/*****  Cabecera de conceptos  *****/
type Conceptos struct {
	Id                     string                   `json:"id,omitempty"`
	ContratoNumeroContrato string                   `json:"contratoNumeroContrato,omitempty"`
	Contrato               *ContratoModelo.Contrato `json:"contrato,omitempty" ;gorm:"unique;not null foreignkey:numero_contrato;association_foreignkey:contrato_numero_contrato"`
	TrasladoPacientes      []TrasladoPacientes      `json:"trasladoPacientes,omitempty"`
	EstudiosLaboratorio    []EstudiosLaboratorio    `json:"estudiosLaboratorio,omitempty"`
	EquipoMedico           []EquipoMedico           `json:"equipoMedico,omitempty"`
}
type TrasladoPacientes struct {
	Id                string                   `json:"id,omitempty" ;gorm:"primaryKey;unique;notnull"`
	Servicio          string                   `json:"servicio,omitempty"`
	Origen            string                   `json:"origen,omitempty"`
	Destino           string                   `json:"destino,omitempty"`
	UnidadMedida      string                   `json:"unidadMedida,omitempty"`
	LugarOtorgamiento string                   `json:"lugarOtorgamiento,omitempty"`
	CantidadMin       int                      `json:"cantidadMin,omitempty"`
	CantidadMax       int                      `json:"cantidadMax,omitempty"`
	PrecioOfertado    float64                  `json:"precioOfertado,omitempty"`
	ImporteMin        float64                  `json:"importeMin,omitempty"`
	ImporteMax        float64                  `json:"importeMax,omitempty"`
	UnidadMedId       int                      `json:"unidadMedId,omitempty"`
	UnidadMed         *ModeloUnidadM.UnidadMed `json:"unidadMed,omitempty"`
	ConceptosId       string                   `json:"conceptosId,omitempty"`
}

type EstudiosLaboratorio struct {
	Id                    string                   `json:"id,omitempty" ;gorm:"primaryKey;unique"`
	Partida               string                   `json:"partida,omitempty"`
	PartidaPresupuestal   string                   `json:"partidaPresupuestal,omitempty"`
	CentroCostos          string                   `json:"centroCostos,omitempty"`
	UnidadMedId           int                      `json:"unidadMedId,omitempty"`
	UnidadMed             *ModeloUnidadM.UnidadMed `json:"unidadMed,omitempty"`
	DescripcionServicioId string                   `json:"descripcionServicioId,omitempty"`
	DescripcionServicio   DescripcionServicio      `json:"descripcionServicio,omitempty"`
	DescripcionServ       string                   `json:"descripcionServ,omitempty"`
	UnidadMedida          string                   `json:"unidadMedida,omitempty"`
	PrecioReferencia      float32                  `json:"precioReferencia,omitempty"`
	EstudiosMin           int                      `json:"estudiosMin,omitempty"`
	EstudiosMax           int                      `json:"estudiosMax,omitempty"`
	ImporteMin            float32                  `json:"importeMin,omitempty"`
	ImporteMax            float32                  `json:"importeMax,omitempty"`
	Dias                  string                   `json:"dias,omitempty"`
	Horario               string                   `json:"horario,omitempty"`
	LugarOtorgamiento     string                   `json:"lugarOtorgamiento,omitempty"`
	ConceptosId           string                   `json:"conceptosId,omitempty"`
}
type EquipoMedico struct {
	Id                  string `gorm:"PRIMARY_KEY"`
	PrecioUniSnIva      float32
	CantidadConcepto    int
	CantidadAmpliada    int
	CantidadDistribuida int
	Marca               string
	Modelo              string
	ObjetoContratacion  string
	FechaMaxEntrega     string
	GarantiaBienes      int
	PreiIdArticulo      string
	Prei                *CatalogosModel.Prei `gorm:"unique;not null;foreignkey:id_articulo;association_foreignkey:prei_id_articulo"`
	ConceptosId         string
}
type RespuestConceptoArchivo struct {
	NumContrato string
	Error       string
	Columna     string
	Fila        int
	Suma        float64
	Tipo        int
}
type Especialidades struct {
	Id     int    `json:"id,omitempty" ;gorm:"primaryKey;unique"`
	Nombre string `json:"nombre,omitempty"`
	Estado bool   `json:"estado"`
}
type DescripcionServicio struct {
	Id               string         `gorm:"primarykey;unique" ;json:"id,omitempty"`
	Descripcion      string         `json:"descripcion,omitempty"`
	EspecialidadesId int            `json:"especialidadesId,omitempty"`
	Especialidades   Especialidades `json:"especialidades,omitempty"`
	Estado           bool           `json:"estado"`
}
