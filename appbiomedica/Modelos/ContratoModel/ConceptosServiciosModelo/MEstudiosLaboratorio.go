package ConceptosServiciosModelo

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/ModeloUnidadM"
	"appbiomedica/Modulos/Estructuras"
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
	"gopkg.in/mgo.v2/bson"
)

//********************************************************************************************************
/* *********************   Funcion para agregar un concepto de un contrato     ********************* */
func (concepto *EstudiosLaboratorio) AgregarConcepto(tx *gorm.DB) error {
	return tx.Create(concepto).Error
}

//********************************************************************************************************
/* *********************   Funcion que regresa la lista de conceptos en el contrato     ********************* */
func (concepto *EstudiosLaboratorio) ConsultarConceptos(num_contrato string, numPag, numReg int) (error,
	*[]Conceptos, float64, int) {
	var datos []Conceptos
	var totalReg int
	type tot struct {
		Total float64
	}
	var total tot
	db := Conexion.GetDB()
	if numReg != 0 && numPag != 0 {
		err := db.Where("contrato_numero_contrato = ?", num_contrato).Preload("EstudiosLaboratorio", func(db *gorm.DB) *gorm.DB {
			return db.Offset((numPag - 1) * numReg).Limit(numReg)
		}).Preload("EstudiosLaboratorio.DescripcionServicio").
		Preload("EstudiosLaboratorio.DescripcionServicio.Especialidades").
		Preload("EstudiosLaboratorio.UnidadMed").Find(&datos).Error
		if err != nil {
			return err, nil, total.Total, totalReg
		}
	} else {
		err := db.Where("contrato_numero_contrato = ?", num_contrato).Preload("EstudiosLaboratorio").
			Preload("EstudiosLaboratorio.DescripcionServicio").
			Preload("EstudiosLaboratorio.DescripcionServicio.Especialidades").
			Preload("EstudiosLaboratorio.UnidadMed").Find(&datos).Error
		if err != nil {
			return err, nil, total.Total, totalReg
		}
	}
	err := db.Find(&[]EstudiosLaboratorio{}).Joins("inner join conceptos c on estudios_laboratorio.conceptos_id = c.id").
		Where("contrato_numero_contrato = ?", num_contrato).Count(&totalReg).Error
	/*****    *****/
	err = db.Table("estudios_laboratorio").Joins("inner join conceptos c on estudios_laboratorio.conceptos_id = c.id").
		Select("sum(importe_max) as total").Where("contrato_numero_contrato = ?", num_contrato).
		Scan(&total).Error
	if err != nil {
		return err, nil, total.Total, totalReg
	}
	return nil, &datos, total.Total, totalReg
}

/*****  Funcion para definir la estructura de un estudio de laboratorio *****/
func (tp *EstudiosLaboratorio) DefinirEstructura(datos []string, id_concepto string) Concepto {
	var conceptos EstudiosLaboratorio
	unidad := ModeloUnidadM.ConsultaUnidadUi(Estructuras.ConvertLowercase(datos[5]),
		Estructuras.ConvertLowercase(datos[3]))
	fmt.Println("unidad medica----", unidad)
	especialidad, _ := BuscarDescServicio(Estructuras.ConvertLowercase(datos[7]),Estructuras.ConvertLowercase(datos[6]))
	can_min, _ := strconv.Atoi(datos[10])
	can_max, _ := strconv.Atoi(datos[11])
	precio, _ := strconv.ParseFloat(datos[9], 64)
	imp_min, _ := strconv.ParseFloat(datos[12], 64)
	imp_max, _ := strconv.ParseFloat(datos[13], 64)
	conceptos.Id = bson.NewObjectId().Hex()
	conceptos.Partida = datos[1]
	conceptos.PartidaPresupuestal = datos[2]
	conceptos.CentroCostos = datos[4]
	conceptos.UnidadMedId = unidad.Id
	conceptos.DescripcionServicioId = especialidad.Id
	conceptos.DescripcionServ = datos[7]
	conceptos.UnidadMedida = datos[8]
	conceptos.PrecioReferencia = float32(precio)
	conceptos.EstudiosMin = can_min
	conceptos.EstudiosMax = can_max
	conceptos.ImporteMin = float32(imp_min)
	conceptos.ImporteMax = float32(imp_max)
	conceptos.Dias = datos[14]
	conceptos.Horario = datos[15]
	conceptos.LugarOtorgamiento = datos[16]
	conceptos.ConceptosId = id_concepto
	return &conceptos
}
