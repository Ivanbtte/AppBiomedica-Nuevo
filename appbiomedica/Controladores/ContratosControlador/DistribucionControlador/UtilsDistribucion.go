package DistribucionControlador

import (
	"appbiomedica/Controladores/ContratosControlador/ConceptosControlador"
	"appbiomedica/Modelos/CatalogosModel"
	"appbiomedica/Modelos/ContratoModel/ConceptoContratoModelo"
	"appbiomedica/Modelos/ContratoModel/ContratoModelo"
	"appbiomedica/Modelos/ContratoModel/DistribucionModelo"
	"appbiomedica/Modelos/ModeloUnidadM"
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/asaskevich/govalidator"
	"io"
	"strconv"
)

type Prei_Contr struct {
	NumContrato string
	IdPrei      string
}
type Prei_ContrCan struct {
	NumContrato string
	IdPrei      string
	Precio      float32
	Marca       string
	Modelo      string
	Fecha       string
	Garantia    int
	Cantidad    int
}
type ContratoDelegacion struct {
	NumContrato  string
	DelegacionId string
}

func ValidarDistribucionConConceptos(file io.Reader) ([]ConceptoContratoModelo.RespuestaDistribucion, error, bool) {
	errores := make([]ConceptoContratoModelo.RespuestaDistribucion, 0)
	encabezado := make([]string, 0)
	archivo, err := excelize.OpenReader(file)
	if err != nil {
		return errores, err, false
	}
	rows, err := archivo.GetRows("Distribucion")
	if err != nil {
		return errores, errors.New("Error al leer el archivo o no es el archivo original"), false
	}
	for num, row := range rows {
		if num > 0 {
			/*Validando que el contrato exista en la base de datos*/
			if !ContratoModelo.CompruebaContrato(row[0]) {
				mensaje := "Error este contrato no ha sido registrado"
				n := ConceptoContratoModelo.RespuestaDistribucion{NumContrato: row[0], IdPrei: row[1], Error: mensaje,
					Columna: encabezado[0], Fila: num + 1, Tipo: 1}
				errores = append(errores, n)
			}
			/*Validando que el id prei (equipo) ya se encuentre registrado en el mismo oontrato*/
			duplicado, equipo := ConceptoContratoModelo.VerificaEquipo(row[1], row[0])
			if duplicado == false {
				mensaje := "Este Id Prei no ha sido registrado en el contrato"
				n := ConceptoContratoModelo.RespuestaDistribucion{NumContrato: row[0], IdPrei: row[1], Error: mensaje,
					Columna: encabezado[1], Fila: num + 1, Tipo: 3}
				errores = append(errores, n)
			} else {
				/*Validando que la cantidad de equipos sea un numero*/
				_, err := strconv.Atoi(row[3])
				if err != nil {
					mensaje := "Cantidad no es un numero"
					n := ConceptoContratoModelo.RespuestaDistribucion{NumContrato: row[0], IdPrei: row[1], Error: mensaje,
						Columna: encabezado[3], Fila: num + 1, Tipo: 3}
					errores = append(errores, n)
				}

				/*Validando que la unidad medica exista*/
				err, unidadMed := ModeloUnidadM.ConsultaUnidadMed(row[2])
				if err != nil {
					/*No encuentra la unidad medica*/
					mensaje := "No existe esta unidad medica"
					n := ConceptoContratoModelo.RespuestaDistribucion{NumContrato: row[0], IdPrei: row[1], Error: mensaje,
						Columna: encabezado[2], Fila: num + 1, Tipo: 3}
					errores = append(errores, n)
				} else {
					/*Si encuentra la unidad medica*/
					/*Validar que no se repita la distribucion  del id_prei para la misma unidad medica y el contrato */
					fmt.Println("Lo encontre:----", unidadMed.ClvPresupuestal, equipo.Id)
					estado, _ := DistribucionModelo.ValidaRepetidos(unidadMed.ClvPresupuestal, equipo.Id)
					if estado {
						mensaje := "Ya existe una distribucion para esta unidad y esta clave"
						n := ConceptoContratoModelo.RespuestaDistribucion{NumContrato: row[0], IdPrei: row[1], Error: mensaje,
							Columna: encabezado[2], Fila: num + 1, Tipo: 2, Concepto: equipo}
						errores = append(errores, n)
					}
				}
			}
		} else {
			encabezado = row
		}
	}
	resultado := FiltraArrayprei(rows, 1)
	for _, id := range resultado {
		res := FiltraArrayUnidad(rows, id.NumContrato+id.IdPrei, 2)
		if res {
			mensaje := "Esta unidad no se puede agregar 2 veces al mismo contrato y el mismo id prei"
			n := ConceptoContratoModelo.RespuestaDistribucion{NumContrato: id.NumContrato, IdPrei: id.IdPrei, Error: mensaje,
				Columna: encabezado[1], Fila: 0, Tipo: 0,}
			errores = append(errores, n)
		}
	}
	/***************************************************/
	/**************************************************/
	resultado2 := FiltraArrayprei(rows, 1)
	fmt.Println("filtroooo", &resultado2)
	/**************************************************/
	for _, registro := range resultado2 {
		_, equipo := ConceptoContratoModelo.VerificaEquipo(registro.IdPrei, registro.NumContrato)
		if equipo != nil {
			/*Validando que la cantidad a ingresar coincida con la cantidad total del concepto*/
			_, total := DistribucionModelo.VerificarCantidades(equipo.Id)
			sumatoria_can := FiltraPreiCantidad(rows, 1, 3, 0, registro.IdPrei, registro.NumContrato)
			res := equipo.CantidadConcepto - total
			if res < sumatoria_can {
				mensaje := "La cantidad que deseas ingresar supera la cantidad disponible por distribuir"
				n := ConceptoContratoModelo.RespuestaDistribucion{NumContrato: registro.NumContrato, IdPrei: registro.IdPrei, Error: mensaje,
					Columna: encabezado[3], Tipo: 2, Cantidad: sumatoria_can, Concepto: equipo}
				errores = append(errores, n)
			}
		}
	}
	if len(errores) > 0 {
		return errores, nil, false
	}
	return errores, nil, true
}
func ValidarDistribucionSinConceptos(file io.Reader) ([]ConceptoContratoModelo.RespuestaDistribucion, error, bool) {
	errores := make([]ConceptoContratoModelo.RespuestaDistribucion, 0)
	encabezado := make([]string, 0)
	archivo, err := excelize.OpenReader(file)
	if err != nil {
		println(err.Error())
		return errores, err, false
	}
	rows, err := archivo.GetRows("DistribucionConceptos")
	if err != nil {
		return errores, errors.New("Error al leer el archivo o no es el archivo original"), false
	}
	for num, row := range rows {
		fmt.Println(row)
		if num > 0 {
			/*Validando que el contrato exista en la base de datos*/
			if !ContratoModelo.CompruebaContrato(row[0]) {
				mensaje := "Error este contrato no ha sido registrado"
				n := ConceptoContratoModelo.RespuestaDistribucion{NumContrato: row[0], IdPrei: row[1], Error: mensaje,
					Columna: encabezado[0], Fila: num + 1, Tipo: 1}
				errores = append(errores, n)
			}
			/*Validando que el id prei (equipo) ya se encuentre registrado en la bd en el mismo oontrato*/
			duplicado, equipo := ConceptoContratoModelo.VerificaEquipo(row[1], row[0])
			if duplicado == true {
				prec, err := strconv.ParseFloat(row[2], 64)
				if err == nil {
					if float64(prec) != equipo.PrecioUniSnIva {
						mensaje := "Este precio de este id no coincide con el registrado en el contrato"
						n := ConceptoContratoModelo.RespuestaDistribucion{NumContrato: row[0], IdPrei: row[1], Error: mensaje,
							Columna: encabezado[2], Fila: num + 1, Tipo: 5, Concepto: equipo}
						errores = append(errores, n)
					}
				}

				err, unidad := ModeloUnidadM.ConsultaUnidadMed(row[7])
				if unidad != nil {
					estado, _ := DistribucionModelo.ValidaRepetidos(unidad.ClvPresupuestal, equipo.Id)
					if estado {
						mensaje := "Ya existe una distribucion para esta unidad y esta clave"
						n := ConceptoContratoModelo.RespuestaDistribucion{NumContrato: row[0], IdPrei: row[1], Error: mensaje,
							Columna: encabezado[7], Fila: num + 1, Tipo: 2, Concepto: equipo}
						errores = append(errores, n)
					}
				}
			}
			/*Validando que la unidad medica exista*/
			err, _ := ModeloUnidadM.ConsultaUnidadMed(row[7])
			if err != nil {
				mensaje := "No existe esta unidad medica"
				n := ConceptoContratoModelo.RespuestaDistribucion{NumContrato: row[0], IdPrei: row[1], Error: mensaje,
					Columna: encabezado[7], Fila: num + 1, Tipo: 3}
				errores = append(errores, n)
			}

			/*Validando que el id prei (equipo) exista en el catalogo del PREI*/
			e, _ := CatalogosModel.ConsultaUnaClave("", row[1])
			if e != nil {
				mensaje := "Este Id Prei No existe en el catalogo"
				n := ConceptoContratoModelo.RespuestaDistribucion{NumContrato: row[0], IdPrei: row[1], Error: mensaje,
					Columna: encabezado[1], Fila: num + 1, Tipo: 0}
				errores = append(errores, n)
			}
			/*Validando que la columna del precio sea un numero o un numero flotante*/
			_, err = strconv.ParseFloat(row[2], 32)
			if err != nil {
				mensaje := "Precio Unitario formato incorrecto"
				n := ConceptoContratoModelo.RespuestaDistribucion{NumContrato: row[0], IdPrei: row[1], Error: mensaje,
					Columna: encabezado[2], Fila: num + 1, Tipo: 0}
				errores = append(errores, n)
			}
			/*Validando que el campo de fecha sea el formato correcto*/
			if !govalidator.IsTime(row[5], "2/1/2006") {
				mensaje := "Fecha de Entrega, formato incorrecto tiene que ser " + "(dd/mm/aaaa)"
				n := ConceptoContratoModelo.RespuestaDistribucion{NumContrato: row[0], IdPrei: row[1], Error: mensaje,
					Columna: encabezado[5], Fila: num + 1, Tipo: 0}
				errores = append(errores, n)
			}
			/*Validando que la garantia  sea un numero*/
			_, err = strconv.Atoi(row[6])
			if err != nil {
				mensaje := "La garantia debe de ser un numero"
				n := ConceptoContratoModelo.RespuestaDistribucion{NumContrato: row[0], IdPrei: row[1], Error: mensaje,
					Columna: encabezado[6], Fila: num + 1, Tipo: 0}
				errores = append(errores, n)
			}
			/*Validando que la cantidad de equipos sea un numero*/
			_, err = strconv.Atoi(row[8])
			if err != nil {
				mensaje := "Cantidad no es un numero"
				n := ConceptoContratoModelo.RespuestaDistribucion{NumContrato: row[0], IdPrei: row[1], Error: mensaje,
					Columna: encabezado[8], Fila: num + 1, Tipo: 3}
				errores = append(errores, n)
			}
		} else {
			encabezado = row
		}
	}
	/*********************validando monto total***********************************************/
	resultado := ConceptosControlador.FiltraArray(rows, 0)
	for _, n_contrato := range resultado {
		sumatoria := ConceptosControlador.VerificaPrecioConceptos(rows, n_contrato, 0, 2, 8)
		contrato := ContratoModelo.RegresaContrato(n_contrato)
		if contrato != nil {
			_, _, suma := ConceptoContratoModelo.ConsultaConceptos(n_contrato)
			sumatoriaTotal := sumatoria + suma
			if sumatoriaTotal > contrato.MontoTotal {
				mensaje := "Error este contrato excede el Monto total del contrato"
				n := ConceptoContratoModelo.RespuestaDistribucion{NumContrato: n_contrato, IdPrei: "0", Error: mensaje,
					Columna: "numero_contrato", Fila: 0, Suma: sumatoria, Tipo: 4}
				errores = append(errores, n)
			}

		}
	}
	/*Validando que no se dupliquen una distribuion de un id_prei en el mismo contrato y la misma unidad medica*/
	resultadoArr := FiltraArrayprei(rows, 1)
	for _, id := range resultadoArr {
		res := FiltraArrayUnidad(rows, id.NumContrato+id.IdPrei, 7)
		if res {
			mensaje := "Este Id Prei no se puede distribuir dos veces en el mismo contrato y en la misma unidad medica"
			n := ConceptoContratoModelo.RespuestaDistribucion{NumContrato: id.NumContrato, IdPrei: id.IdPrei,
				Error: mensaje, Columna: encabezado[1], Fila: 0, Tipo: 0,}
			errores = append(errores, n)
		}
	}
	/*Validando que los precios coincidan con los id prei*/
	for _, idPrei := range resultadoArr {
		precios, posci := VerificarPrecioPrei(rows, idPrei.NumContrato+idPrei.IdPrei, 2)
		if precios {
			mensaje := "Este Id Prei tiene dos precios diferentes"
			n := ConceptoContratoModelo.RespuestaDistribucion{NumContrato: idPrei.NumContrato, IdPrei: idPrei.IdPrei,
				Error: mensaje, Columna: encabezado[2], Fila: posci + 1, Tipo: 0,}
			errores = append(errores, n)
		}
	}
	if len(errores) > 0 {
		return errores, nil, false
	}
	return errores, nil, true
}

//********************************************************************************************************
/********************************Funcion que regresa un arreglo de no repetidos*************************************/
/*   Funcion que recibe una matriz (lectura de un archivo excel)  y la posicion de la columna que deseamos filtrar  */
func FiltraPreiCantidad(arreglo [][]string, pos_prei, pos_cant, pos_contrato int, id_prei, num_contrato string) int {
	var total int
	for num, row := range arreglo {
		if num > 0 {
			if row[pos_prei] == id_prei && row[pos_contrato] == num_contrato {
				cantidad, _ := strconv.Atoi(row[pos_cant])
				total = total + cantidad
			}
		}
	}
	return total
}
func FiltraArrayprei(arreglo [][]string, posicion int) []Prei_Contr {
	var unique []Prei_Contr
	m := map[string]bool{}
	for num, row := range arreglo {
		if num > 0 {
			if !m[row[posicion]+row[0]] {
				m[row[posicion]+row[0]] = true
				ma := Prei_Contr{NumContrato: row[0], IdPrei: row[posicion]}
				unique = append(unique, ma)
			}
		}
	}
	return unique
}
func FiltraArrayUnidad(arreglo [][]string, id_prei string, pos_uiPrei int) bool {
	m := map[string]bool{}
	var resultado bool
	for num, row := range arreglo {
		if num > 0 {
			if id_prei == row[0]+row[1] {
				if !m[id_prei+row[pos_uiPrei]] {
					m[id_prei+row[pos_uiPrei]] = true
					resultado = false
				} else {
					resultado = true
					return resultado
				}
			}
		}
	}
	return resultado
}
func FiltraDelegaciones(arreglo [][]string, unidad_inf int) []ContratoDelegacion {
	var unique []ContratoDelegacion
	m := map[string]bool{}
	for num, row := range arreglo {
		if num > 0 {
			if !m[row[0]+row[unidad_inf][:2]] {
				m[row[0]+row[unidad_inf][:2]] = true
				ma := ContratoDelegacion{NumContrato: row[0], DelegacionId: row[unidad_inf][:2]}
				unique = append(unique, ma)
			}
		}
	}
	return unique
}
func FiltraArraypreiCantidad(arreglo [][]string, posicion int) []Prei_ContrCan {
	var unique []Prei_ContrCan
	m := map[string]bool{}
	for num, row := range arreglo {
		if num > 0 {
			if !m[row[posicion]+row[0]] {
				m[row[posicion]+row[0]] = true
				precio, _ := strconv.ParseFloat(row[2], 32)
				garantia, _ := strconv.Atoi(row[6])
				ma := Prei_ContrCan{NumContrato: row[0], IdPrei: row[posicion], Precio: float32(precio), Marca: row[3],
					Modelo: row[4], Fecha: row[5], Garantia: garantia}
				unique = append(unique, ma)
			}
		}
	}
	for i, concepto := range unique {
		var suma int
		for num, fila := range arreglo {
			cantidad, _ := strconv.Atoi(fila[8])
			if num > 0 {
				if concepto.NumContrato == fila[0] && concepto.IdPrei == fila[1] {
					suma = cantidad + suma

					fmt.Println("**********", concepto.NumContrato, "***", concepto.IdPrei, suma)
					unique[i].Cantidad = suma
				}
			}
		}
	}
	return unique
}
func VerificarPrecioPrei(arreglo [][]string, id_prei string, pos_precio int) (bool, int) {
	m := map[string]bool{}
	var resultado bool
	contador := 0
	for num, row := range arreglo {
		if num > 0 {
			if id_prei == row[0]+row[1] {
				if !m[id_prei+row[pos_precio]] {
					m[id_prei+row[pos_precio]] = true
					resultado = false
					contador = contador + 1
				}
				if contador > 1 {
					return true, num
				}
			}
		}
	}
	return resultado, 0
}
