package ConceptosServiciosModelo

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/ModeloUnidadM"
	"appbiomedica/Modulos/Estructuras"
	"fmt"
	"github.com/jinzhu/gorm"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

//********************************************************************************************************
/* *********************   Funcion para agregar un concepto de un contrato     ********************* */
func (concepto *TrasladoPacientes) AgregarConcepto(tx *gorm.DB) error {
	return tx.Create(concepto).Error
}

/*****  Funcion para consultar los diferentes destinos disponibles por unidad medica y numero de contrato  *****/
func ConsultarDestinos(num_contrato string, unidad_id int) (error, *[]TrasladoPacientes) {
	db := Conexion.GetDB()
	var traslados []TrasladoPacientes
	err := db.Select("distinct  on (destino)*").
		Joins("inner join conceptos c on traslado_pacientes.conceptos_id = c.id").
		Joins("inner join contrato c2 on c.contrato_numero_contrato = c2.numero_contrato").
		Where("numero_contrato = ?", num_contrato).Where("unidad_med_id = ?", unidad_id).Find(&traslados).Error
	if err != nil {
		return err, nil
	}
	return nil, &traslados
}

/***  Funcion para consultar los diferentes origenes disponibles por unidad medica y numero de contrato y destino ***/
func ConsultarOrigenes(num_contrato, destino string, unidad_id int) (error, *[]TrasladoPacientes) {
	db := Conexion.GetDB()
	var traslados []TrasladoPacientes
	err := db.Select("distinct  on (origen)*").
		Joins("inner join conceptos c on traslado_pacientes.conceptos_id = c.id").
		Joins("inner join contrato c2 on c.contrato_numero_contrato = c2.numero_contrato").
		Where("numero_contrato = ?", num_contrato).Where("unidad_med_id = ?", unidad_id).
		Where("destino = ?", destino).Find(&traslados).Error
	if err != nil {
		return err, nil
	}
	return nil, &traslados
}

//********************************************************************************************************
/* *********************   Funcion que regresa la lista de conceptos en el contrato     ********************* */
func (concepto *TrasladoPacientes) ConsultarConceptos(num_contrato string, numPag, numReg int) (error, *[]Conceptos,
	float64, int) {
	var totalReg int
	var datos []Conceptos
	type tot struct {
		Total float64
	}
	var total tot
	db := Conexion.GetDB()
	if numReg != 0 && numPag != 0 {
		fmt.Println("entre con paginador")
		err := db.Where("contrato_numero_contrato = ?", num_contrato).Preload("TrasladoPacientes", func(db *gorm.DB) *gorm.DB {
			return db.Offset((numPag - 1) * numReg).Limit(numReg)
		}).Preload("TrasladoPacientes.UnidadMed").Find(&datos).Error
		if err != nil {
			return err, nil, total.Total, totalReg
		}
	} else {
		fmt.Println("sin paginador")
		err := db.Where("contrato_numero_contrato = ?", num_contrato).Preload("TrasladoPacientes").
			Preload("TrasladoPacientes.UnidadMed").Find(&datos).Error
		if err != nil {
			return err, nil, total.Total, totalReg
		}
	}
	err := db.Find(&[]TrasladoPacientes{}).Joins("inner join conceptos c on traslado_pacientes.conceptos_id = c.id").
		Where("contrato_numero_contrato = ?", num_contrato).Count(&totalReg).Error
	/*****    *****/
	err = db.Table("traslado_pacientes").Joins("inner join conceptos c on traslado_pacientes.conceptos_id = c.id").
		Select("sum(importe_max) as total").Where("contrato_numero_contrato = ?", num_contrato).Scan(&total).Error
	if err != nil {
		return err, nil, total.Total, totalReg
	}
	return nil, &datos, total.Total, totalReg
}

/*****  Funcion para definir la estructura de un traslado de paciente *****/

func (tp *TrasladoPacientes) DefinirEstructura(datos []string, id_concepto string) Concepto {
	var conceptos TrasladoPacientes
	_, unidad := ModeloUnidadM.ConsultaUnidadNombre(Estructuras.ConvertLowercase(datos[2]),
		Estructuras.ConvertLowercase(datos[1]))
	canMin, _ := strconv.Atoi(datos[8])
	canMax, _ := strconv.Atoi(datos[9])
	precio, _ := strconv.ParseFloat(datos[10], 64)
	impMin, _ := strconv.ParseFloat(datos[11], 64)
	impMax, _ := strconv.ParseFloat(datos[12], 64)
	conceptos.Id = bson.NewObjectId().Hex()
	conceptos.Servicio = datos[3]
	conceptos.Origen = datos[4]
	conceptos.Destino = datos[5]
	conceptos.UnidadMedida = datos[6]
	conceptos.LugarOtorgamiento = datos[7]
	conceptos.CantidadMin = canMin
	conceptos.CantidadMax = canMax
	conceptos.PrecioOfertado = precio
	conceptos.ImporteMin = impMin
	conceptos.ImporteMax = impMax
	conceptos.UnidadMedId = unidad.Id
	conceptos.ConceptosId = id_concepto
	return &conceptos
}


func ConsultBudget(contractNumber string, unitId int) (error, float64) {
	type Budget struct {
		Sum float64
	}
	var result Budget
	db := Conexion.GetDB()
	err := db.Table("traslado_pacientes").Select("sum(importe_max) as sum").
		Joins("inner join conceptos c on traslado_pacientes.conceptos_id = c.id").
		Where("c.contrato_numero_contrato = ?", contractNumber).Where("unidad_med_id = ?", unitId).Scan(&result).Error
	if err != nil {
		return err, 0
	}
	return nil, result.Sum
}
