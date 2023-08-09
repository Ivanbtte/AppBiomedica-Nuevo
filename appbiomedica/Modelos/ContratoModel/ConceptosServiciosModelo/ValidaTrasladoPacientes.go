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
func (concepto TrasladoPacientes) ValidaConceptosArchivo(file io.Reader) (bool, []RespuestConceptoArchivo, error) {
	var errores []RespuestConceptoArchivo
	var cabeceras []string
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
	/*****    *****/
	for num, row := range rows {
		if num == 0 {
			cabeceras = row
		}
		if len(row) != 13 {
			return false, errores, errors.New("error en el numero de columnas del archivo")
		}
		if len(row) == 13 && num > 0 {
			/*Validando que el contrato exista en la base de datos*/
			if !ContratoModelo.CompruebaContrato(Estructuras.ConvertLowercase(row[0])) {
				mensaje := "Error este contrato no ha sido registrado"
				n := RespuestConceptoArchivo{NumContrato: row[0], Error: mensaje,
					Columna: cabeceras[0], Fila: num + 1, Tipo: 1}
				errores = append(errores, n)
			}
			/*****  Validando que la unidad medica exista  *****/
			respuesta, _ := ModeloUnidadM.ConsultaUnidadNombre(Estructuras.ConvertLowercase(row[2]),
				Estructuras.ConvertLowercase(row[1]))
			if !respuesta {
				mensaje := "Error no existe esta unidad medica"
				n := RespuestConceptoArchivo{NumContrato: row[0], Error: mensaje, Columna: cabeceras[2], Fila: num + 1,
					Tipo: 0}
				errores = append(errores, n)
			}
			/*Validando que la cantidad minima sea un numero*/
			_, err = strconv.Atoi(row[8])
			if err != nil {
				mensaje := "Cantidad minima no es un numero"
				n := RespuestConceptoArchivo{NumContrato: row[0], Error: mensaje, Columna: cabeceras[7], Fila: num + 1,
					Tipo: 0}
				errores = append(errores, n)
			}
			/*Validando que la cantidad maxima sea un numero*/
			_, err = strconv.Atoi(row[9])
			if err != nil {
				mensaje := "Cantidad maxima no es un numero"
				n := RespuestConceptoArchivo{NumContrato: row[0], Error: mensaje, Columna: cabeceras[8], Fila: num + 1,
					Tipo: 0}
				errores = append(errores, n)
			}
			/*Validando que la columna del precio sea un numero flotante*/
			_, err := strconv.ParseFloat(row[10], 32)
			if err != nil {
				mensaje := "Precio ofertado formato incorrecto"
				n := RespuestConceptoArchivo{NumContrato: row[0], Error: mensaje, Columna: cabeceras[9], Fila: num + 1,
					Tipo: 0}
				errores = append(errores, n)
			}
			/*Validando que la columna del importe minimo sea un numero flotante*/
			_, err = strconv.ParseFloat(row[11], 32)
			if err != nil {
				mensaje := "Importe minimo formato incorrecto"
				n := RespuestConceptoArchivo{NumContrato: row[0], Error: mensaje, Columna: cabeceras[10], Fila: num + 1,
					Tipo: 0}
				errores = append(errores, n)
			}
			/*Validando que la columna del importe maximo sea un numero flotante*/
			_, err = strconv.ParseFloat(row[12], 32)
			if err != nil {
				mensaje := "Importe maximo formato incorrecto"
				n := RespuestConceptoArchivo{NumContrato: row[0], Error: mensaje, Columna: cabeceras[11], Fila: num + 1,
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
		sumatoria := VerificaPrecioConceptos(rows, n_contrato, 0, 12)
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
	if len(errores) > 0 {
		return false, errores, nil
	}
	return true, errores, nil
}

//********************************************************************************************************
/********************************Funcion que regresa un arreglo de no repetidos*************************************/
/*   Funcion que recibe una matriz (lectura de un archivo excel)  y la posicion de la columna que deseamos filtrar  */
func FiltraArray(arreglo [][]string, posicion int) []string {
	var unique []string
	m := map[string]bool{}
	for num, row := range arreglo {
		if num > 0 {
			if !m[row[posicion]] {
				m[row[posicion]] = true
				unique = append(unique, row[posicion])
			}
		}
	}
	return unique
}

//********************************************************************************************************
/* *****************    Funcion para comprobar la sumatopria de los precios de los conceptos      ****************** */
func VerificaPrecioConceptos(datos [][]string, n_contrato string, pos_cont, pos_prec int) float64 {
	var sumatoria float64
	for linea, registro := range datos {
		if linea > 0 {
			if registro[pos_cont] == n_contrato {
				precio, _ := strconv.ParseFloat(registro[pos_prec], 64)
				sumatoria = sumatoria + float64(precio)
			}
		}
	}
	return sumatoria
}
