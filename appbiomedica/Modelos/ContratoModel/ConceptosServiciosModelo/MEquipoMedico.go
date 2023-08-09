package ConceptosServiciosModelo

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

//********************************************************************************************************
/* *********************   Funcion para agregar un concepto de un contrato     ********************* */
func (concepto *EquipoMedico) AgregarConcepto(tx *gorm.DB) error {
	return tx.Create(concepto).Error
}

//********************************************************************************************************
/* *********************   Funcion que regresa la lista de conceptos en el contrato     ********************* */
func (concepto *EquipoMedico) ConsultarConceptos(id_contrato string, numPag, numReg int) (error, *[]Conceptos,
	float64, int) {
	fmt.Println(id_contrato)
	//var datos []interface{}
	type tot struct {
		Total float64
	}
	var total tot
	/*var datos []TrasladoPacientes
	type tot struct {
		Total float32
	}
	var total tot
	db := Conexion.GetDB()
	err := db.Where("contrato_numero_contrato = ?", id_contrato).Preload("UnidadMed").Preload("Contrato").
		Find(&datos).Error
	if err != nil {
		return err, nil, total.Total
	}
	err = db.Table("traslado_pacientes").Select("sum(importe_max) as total").
		Where("contrato_numero_contrato = ?", id_contrato).Scan(&total).Error
	if err != nil {
		return err, nil, total.Total
	}*/
	return nil, &[]Conceptos{}, total.Total, 0
}

/*****  Funcion para definir la estructura de un equipo medico *****/
func (tp *EquipoMedico) DefinirEstructura(datos []string, id_concepto string) Concepto {
	var conceptos EquipoMedico
	fmt.Println(datos)
	/*	_, unidad := ModeloUnidadM.ConsultaUnidadNombre(Estructuras.ConvertirMinusculas(datos[2]),
			Estructuras.ConvertirMinusculas(datos[1]))
		can_min, _ := strconv.Atoi(datos[7])
		can_max, _ := strconv.Atoi(datos[8])
		precio, _ := strconv.ParseFloat(datos[9], 64)
		imp_min, _ := strconv.ParseFloat(datos[10], 64)
		imp_max, _ := strconv.ParseFloat(datos[11], 64)
		conceptos.Id = bson.NewObjectId().Hex()
		conceptos.Servicio = datos[3]
		conceptos.Ruta = datos[4]
		conceptos.UnidadMedida = datos[5]
		conceptos.LugarOtorgamiento = datos[6]
		conceptos.CantidadMin = can_min
		conceptos.CantidadMax = can_max
		conceptos.PrecioOfertado = precio
		conceptos.ImporteMin = imp_min
		conceptos.ImporteMax = imp_max
		conceptos.UnidadMedId = unidad.Id
		conceptos.ContratoNumeroContrato = datos[0]*/
	return &conceptos
}
