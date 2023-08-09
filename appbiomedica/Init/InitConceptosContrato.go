package Init

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/ContratoModel/ConceptosServiciosModelo"
)

func InitConceptos() error {
	db := Conexion.GetDB()
	/*****  Tabla para almacenar las cabeceras de los conceptos  *****/
	if !db.HasTable(ConceptosServiciosModelo.Conceptos{}) {
		if err := db.CreateTable(ConceptosServiciosModelo.Conceptos{}).
			AddForeignKey("contrato_numero_contrato", "contrato(numero_contrato)", "CASCADE", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	/*****  Tabla del catalogo de especialidades  *****/
	if !db.HasTable(&ConceptosServiciosModelo.Especialidades{}) {
		if err := db.CreateTable(&ConceptosServiciosModelo.Especialidades{}).
			Error; err != nil {
			return err
		}
	}
	/*****  Tabla del catalogo de servicios de estudios  *****/
	if !db.HasTable(&ConceptosServiciosModelo.DescripcionServicio{}) {
		if err := db.CreateTable(&ConceptosServiciosModelo.DescripcionServicio{}).
		AddForeignKey("especialidades_id", "especialidades(id)", "SET NULL", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	/*****  Tabla para los conceptos de tarslado de pacientes  *****/
	if !db.HasTable(&ConceptosServiciosModelo.TrasladoPacientes{}) {
		if err := db.CreateTable(&ConceptosServiciosModelo.TrasladoPacientes{}).
			AddForeignKey("unidad_med_id", "unidad_med(id)", "CASCADE", "CASCADE").
			AddForeignKey("conceptos_id", "conceptos(id)", "CASCADE", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	/*****  Tabla para los conceptos de estudios de laboratorio  *****/
	if !db.HasTable(&ConceptosServiciosModelo.EstudiosLaboratorio{}) {
		if err := db.CreateTable(&ConceptosServiciosModelo.EstudiosLaboratorio{}).
			AddForeignKey("unidad_med_id", "unidad_med(id)", "SET NULL", "CASCADE").
			AddForeignKey("conceptos_id", "conceptos(id)", "SET NULL", "CASCADE").
			AddForeignKey("descripcion_servicio_id", "descripcion_servicio(id)", "SET NULL", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	return nil
}
