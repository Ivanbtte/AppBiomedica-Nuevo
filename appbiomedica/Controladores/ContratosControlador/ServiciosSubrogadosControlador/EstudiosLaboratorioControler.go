package ServiciosSubrogadosControlador

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/ContratoModel/ConceptosServiciosModelo"
	"appbiomedica/Modulos/Estructuras"
	"errors"
	"io"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"gopkg.in/mgo.v2/bson"
)

func ValidarEstudiosDescripcion(file io.Reader) (bool, []ConceptosServiciosModelo.RespuestConceptoArchivo, error) {
	var errores []ConceptosServiciosModelo.RespuestConceptoArchivo
	var cabeceras []string
	archivo, err := excelize.OpenReader(file)
	if err != nil {
		return false, errores, errors.New("error al leer el archivo")
	}
	// Obtener todos los registros de la hoja conceptos.
	rows, err := archivo.GetRows("Conceptos")
	if err != nil {
		return false, errores, errors.New("no se pudo leer el archivo, error al obtener los datos de la hoja Conceptos")
	}
	for num, row := range rows {
		if num == 0 {
			cabeceras = row
		}
		if len(row) != 2 {
			return false, errores, errors.New("error en el numero de columnas del archivo")
		}
		if len(row) == 2 && num > 0 {
			// validando que el servicio exista en nuestro catalogo
			especialidad, _ := ConceptosServiciosModelo.BuscarEspecialidad(Estructuras.ConvertLowercase(row[1]))
			if especialidad == nil {
				mensaje := "Error no existe esta especialidad *" + row[1] + "* en nuestro catalogo"
				n := ConceptosServiciosModelo.RespuestConceptoArchivo{NumContrato: row[0], Error: mensaje, Columna: cabeceras[1], Fila: num + 1}
				errores = append(errores, n)
			}
			servicio, _ := ConceptosServiciosModelo.BuscarDescServicio(Estructuras.ConvertLowercase(row[0]), Estructuras.ConvertLowercase(row[1]))
			if servicio != nil {
				mensaje := "Error este servicio *" + row[0] + "* ya se encuentra registrado"
				n := ConceptosServiciosModelo.RespuestConceptoArchivo{NumContrato: row[0], Error: mensaje, Columna: cabeceras[0], Fila: num + 1}
				errores = append(errores, n)
			}
		}
	}
	if len(errores) > 0 {
		return false, errores, nil
	}
	return true, errores, nil
}
func AgregarServicioDesc(file io.Reader) error {
	var servicioDes ConceptosServiciosModelo.DescripcionServicio
	archivo, err := excelize.OpenReader(file)
	if err != nil {
		return errors.New("error al leer el archivo")
	}
	// Obtener todos los registros de la hoja conceptos.
	rows, err := archivo.GetRows("Conceptos")
	if err != nil {
		return errors.New("no se pudo leer el archivo, error al obtener los datos de la hoja Conceptos")
	}
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		return errors.New("error al iniciar la transaccion")
	}
	for num, row := range rows {
		if len(row) != 2 {
			return errors.New("error en el numero de columnas del archivo")
		}
		if len(row) == 2 && num > 0 {
			// validando que el servicio exista en nuestro catalogo
			especialidad, _ := ConceptosServiciosModelo.BuscarEspecialidad(Estructuras.ConvertLowercase(row[1]))
			if especialidad == nil {
				mensaje := "Error no existe esta especialidad *" + row[1] + "* en nuestro catalogo"
				return errors.New(mensaje)
			}
			servicio, _ := ConceptosServiciosModelo.BuscarDescServicio(Estructuras.ConvertLowercase(row[0]), Estructuras.ConvertLowercase(row[1]))
			if servicio != nil {
				mensaje := "Error este servicio *" + row[0] + "* ya se encuentra registrado"
				return errors.New(mensaje)
			}
			servicioDes.Id = bson.NewObjectId().Hex()
			servicioDes.Descripcion = row[0]
			servicioDes.EspecialidadesId = especialidad.Id
			servicioDes.Estado = true
			err := servicioDes.AgregarDescServicios(tx)
			if err != nil {
				tx.Rollback()
				return errors.New("error al insertar el concepto")
			}
		}
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return errors.New("error al realizar el commit")
	}
	return nil
}
