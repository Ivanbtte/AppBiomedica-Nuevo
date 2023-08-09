package ConceptosServiciosModelo

import (
	"appbiomedica/Modelos/ContratoModel/ContratoModelo"
	"appbiomedica/Modelos/ModeloUnidadM"
	"appbiomedica/Modulos/Estructuras"
	"errors"
	"fmt"
	"io"
	"strconv"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

//********************************************************************************************************
/* *********************    funcion para validar un archivo de excel       ********************* */
func (concepto EstudiosLaboratorio) ValidaConceptosArchivo(file io.Reader) (bool, []RespuestConceptoArchivo, error) {
	var errores []RespuestConceptoArchivo
	var cabeceras []string
	var sumatoria float64
	/*****  Leyendo archivo  *****/
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
		if len(row) != 17 {
			return false, errores, errors.New("error en el numero de columnas del archivo")
		}
		if len(row) == 17 && num > 0 {
			/*Validando que el contrato exista en la base de datos*/
			if !ContratoModelo.CompruebaContrato(Estructuras.ConvertLowercase(row[0])) {
				mensaje := "Error este contrato no ha sido registrado"
				n := RespuestConceptoArchivo{NumContrato: row[0], Error: mensaje,
					Columna: cabeceras[0], Fila: num + 1, Tipo: 1}
				errores = append(errores, n)
			}
			/*****  Validando que la unidad medica exista  *****/
			respuesta := ModeloUnidadM.ConsultaUnidadUi(Estructuras.ConvertLowercase(row[5]),
				Estructuras.ConvertLowercase(row[3]))
			if respuesta == nil {
				mensaje := "Error no existe esta unidad medica"
				n := RespuestConceptoArchivo{NumContrato: row[0], Error: mensaje, Columna: cabeceras[5], Fila: num + 1,
					Tipo: 0}
				errores = append(errores, n)
			}
			// validando que el servicio exista en nuestro catalogo
			especialidad, _ := BuscarDescServicio(Estructuras.ConvertLowercase(row[7]),Estructuras.ConvertLowercase(row[6]))
			if especialidad == nil {
				mensaje := "Error no existe este servicio *" + row[6] + "* en nuestro catalogo"
				n := RespuestConceptoArchivo{NumContrato: row[0], Error: mensaje, Columna: cabeceras[7], Fila: num + 1}
				errores = append(errores, n)
			}
			// validando el formato de precio de referencia
			_, err := strconv.ParseFloat(row[9], 32)
			if err != nil {
				mensaje := "Precio ofertado formato incorrecto"
				n := RespuestConceptoArchivo{NumContrato: row[0], Error: mensaje, Columna: cabeceras[9], Fila: num + 1,
					Tipo: 0}
				errores = append(errores, n)
			}
			// validando la cantidad minima
			_, err = strconv.Atoi(row[10])
			if err != nil {
				mensaje := "Cantidad minima no es un numero"
				n := RespuestConceptoArchivo{NumContrato: row[0], Error: mensaje, Columna: cabeceras[10], Fila: num + 1,
					Tipo: 0}
				errores = append(errores, n)
			}
			// validando la cantidad maxima
			_, err = strconv.Atoi(row[11])
			if err != nil {
				mensaje := "Cantidad maxima no es un numero"
				n := RespuestConceptoArchivo{NumContrato: row[0], Error: mensaje, Columna: cabeceras[11], Fila: num + 1,
					Tipo: 0}
				errores = append(errores, n)
			}
			// validando el importe minimo
			_, err = strconv.ParseFloat(row[12], 32)
			if err != nil {
				mensaje := "Importe minimo formato incorrecto"
				n := RespuestConceptoArchivo{NumContrato: row[0], Error: mensaje, Columna: cabeceras[12], Fila: num + 1,
					Tipo: 0}
				errores = append(errores, n)
			}
			// validando el importe maximo
			_, err = strconv.ParseFloat(row[13], 32)
			if err != nil {
				mensaje := "Importe maximo formato incorrecto"
				n := RespuestConceptoArchivo{NumContrato: row[0], Error: mensaje, Columna: cabeceras[13], Fila: num + 1,
					Tipo: 0}
				errores = append(errores, n)
			}
		}
	}
		/*********************validando monto total***********************************************/
		resultado := FiltraArray(rows, 0)
		fmt.Println("sin repetidos: ", resultado)
		/*****    *****/
		for _, n_contrato := range resultado {
			sumatoria = VerificaPrecioConceptos(rows, n_contrato, 0, 13)
			fmt.Println("Suma total por contrato: Contrato No. : ", n_contrato, "Total $$: ", sumatoria)
			fmt.Printf("%f\n", sumatoria)
			contrato := ContratoModelo.RegresaContrato(n_contrato)
			if contrato != nil {
				_, _, suma, _ := concepto.ConsultarConceptos(n_contrato,0,0)
				sumatoriaTotal := sumatoria + suma
				if sumatoriaTotal > contrato.MontoTotal {
					mensaje := "Error este contrato excede el Monto total del contrato"
					n := RespuestConceptoArchivo{NumContrato: n_contrato, Error: mensaje, Columna: "Numero Contrato",
						Fila: 0, Suma: sumatoriaTotal, Tipo: 2}
					errores = append(errores, n)
				}
			}
		}
		fmt.Println("sumatoriaa:",sumatoria)
	if len(errores) > 0 {
		return false, errores, nil
	}
	return true, errores, nil
}
