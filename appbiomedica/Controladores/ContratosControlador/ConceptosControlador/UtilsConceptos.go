package ConceptosControlador

import (
	"appbiomedica/Modelos/CatalogosModel"
	"appbiomedica/Modelos/ContratoModel/ConceptoContratoModelo"
	"appbiomedica/Modelos/ContratoModel/ContratoModelo"
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/asaskevich/govalidator"
	"io"
	"strconv"
)

//********************************************************************************************************
/* *********************    Funcion que valida un archivo de carga de conceptos     ********************* */
func ValidarArchivo(file io.Reader) (bool, []ConceptoContratoModelo.RespuestaArchivo, error) {
	m := map[string]bool{}
	archivo, err := excelize.OpenReader(file)
	var errores []ConceptoContratoModelo.RespuestaArchivo
	fmt.Println("Leyendo Archivo verificando")
	if err != nil {
		println(err.Error())
		return false, errores, errors.New("Error al leer el archivo")
	}
	var cabeceras []string
	// Get all the rows in the Sheet1.
	rows, _ := archivo.GetRows("Conceptos")
	//if err != nil {
	//	println(err.Error())
	//	return false, errores, errors.New("No se pudo leer el archivo o no tiene el formato correcto")
	//}
	for num, row := range rows {
		if num == 0 {
			cabeceras = row
		}
		if len(row) == 8 && num > 0 {
			/*Validando que el contrato exista en la base de datos*/
			if !ContratoModelo.CompruebaContrato(row[7]) {
				mensaje := "Error este contrato no ha sido registrado"
				n := ConceptoContratoModelo.RespuestaArchivo{NumContrato: row[7], IdPrei: row[6], Error: mensaje,
					Columna: cabeceras[7], Fila: num, Tipo: 1}
				errores = append(errores, n)
			}
			/*Validando que el id prei (equipo) exista en el catalogo del PREI*/
			e, _ := CatalogosModel.ConsultaUnaClave("", row[6])
			if e != nil {
				mensaje := "Este Id Prei No existe en el catalogo prei"
				n := ConceptoContratoModelo.RespuestaArchivo{NumContrato: row[7], IdPrei: row[6], Error: mensaje,
					Columna: cabeceras[6], Fila: num, Tipo: 0}
				errores = append(errores, n)
			}
			/*Validando que el id prei (equipo) ya se encuentre registrado en el mismo oontrato*/
			duplicado, _ := ConceptoContratoModelo.VerificaEquipo(row[6], row[7])
			if duplicado == true {
				mensaje := "Este Id Prei ya ha sido registrado en este contrato"
				n := ConceptoContratoModelo.RespuestaArchivo{NumContrato: row[7], IdPrei: row[6], Error: mensaje,
					Columna: cabeceras[6], Fila: num, Tipo: 0}
				errores = append(errores, n)
			}
			/*Validando que la columna del precio sea un numero o un numero flotante*/
			_, err := strconv.ParseFloat(row[0], 32)
			if err != nil {
				mensaje := "Precio Unitario formato incorrecto"
				n := ConceptoContratoModelo.RespuestaArchivo{NumContrato: row[7], IdPrei: row[6], Error: mensaje,
					Columna: cabeceras[0], Fila: num, Tipo: 0}
				errores = append(errores, n)
			}
			/*Validando que la cantidad de equipos sea un numero*/
			_, err = strconv.Atoi(row[1])
			if err != nil {
				mensaje := "Cantidad no es un numero"
				n := ConceptoContratoModelo.RespuestaArchivo{NumContrato: row[7], IdPrei: row[6], Error: mensaje,
					Columna: cabeceras[1], Fila: num, Tipo: 0}
				errores = append(errores, n)
			}
			/*Validando que la garantia  sea un numero*/
			_, err = strconv.Atoi(row[5])
			if err != nil {
				mensaje := "La garantia debe de ser un numero"
				n := ConceptoContratoModelo.RespuestaArchivo{NumContrato: row[7], IdPrei: row[6], Error: mensaje,
					Columna: cabeceras[5], Fila: num, Tipo: 0}
				errores = append(errores, n)
			}
			/*Validando que el campo de fechq sea el formato correcto*/
			if !govalidator.IsTime(row[4], "2/1/2006") {
				mensaje := "Fecha de Entrega, formato incorrecto tiene que ser " + "(dd/mm/aaaa)"
				n := ConceptoContratoModelo.RespuestaArchivo{NumContrato: row[7], IdPrei: row[6], Error: mensaje,
					Columna: cabeceras[4], Fila: num, Tipo: 0}
				errores = append(errores, n)
			}
			/*Validando archivo que no se repitan los id prei*/
			if !m[row[7]+row[6]] {
				m[row[7]+row[6]] = true
			} else {
				mensaje := "Este id prei se duplica en el mismo contrato"
				np := ConceptoContratoModelo.RespuestaArchivo{NumContrato: row[7], IdPrei: row[6], Error: mensaje,
					Columna: cabeceras[6], Fila: num + 1, Suma: 0, Tipo: 0}
				errores = append(errores, np)
			}
		}
	}
	fmt.Println(cabeceras)
	fmt.Println(errores)
	/*********************validando monto total***********************************************/
	resultado := FiltraArray(rows, 7)
	fmt.Println("sin repetidos: ", resultado)
	/*****************************************************************************************************************/
	for _, n_contrato := range resultado {
		sumatoria := VerificaPrecioConceptos(rows, n_contrato, 7, 0, 1)
		fmt.Println("Suma total por contrato: Contrato No. : ", n_contrato, "Total $$: ", sumatoria)
		fmt.Printf("%f\n", sumatoria)
		contrato := ContratoModelo.RegresaContrato(n_contrato)
		if contrato != nil {
			_, _, suma := ConceptoContratoModelo.ConsultaConceptos(n_contrato)
			sumatoriaTotal := sumatoria + suma
			if sumatoriaTotal > contrato.MontoTotal {
				//s := strconv.FormatFloat(float64(sumatoria), 'f', 3, 64)
				mensaje := "Error este contrato excede el Monto total del contrato"
				n := ConceptoContratoModelo.RespuestaArchivo{NumContrato: n_contrato, IdPrei: "0", Error: mensaje,
					Columna: "numero_contrato", Fila: 0, Suma: sumatoria, Tipo: 2}
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
func VerificaPrecioConceptos(datos [][]string, n_contrato string, pos_cont, pos_prec, pos_cant int) float64 {
	var sumatoria, total float64
	for linea, registro := range datos {
		if linea > 0 {
			if registro[pos_cont] == n_contrato {
				precio, _ := strconv.ParseFloat(registro[pos_prec], 32)
				cantidad, _ := strconv.Atoi(registro[pos_cant])
				total = float64(precio) * float64(cantidad)
				sumatoria = sumatoria + total
			}
		}
	}
	return sumatoria
}
