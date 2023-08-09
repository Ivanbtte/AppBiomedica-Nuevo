package ContratoControlador

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/ContratoModel/AdminAuxModelo"
	"appbiomedica/Modelos/ContratoModel/ContactoModelo"
	"appbiomedica/Modelos/ContratoModel/ContratoModelo"
	"appbiomedica/Modelos/ContratoModel/FianzaModelo"
	"appbiomedica/Modelos/ContratoModel/PenasModelo"
	"appbiomedica/Modelos/ContratoModel/RepresentanteModel"
	"appbiomedica/Modelos/CorreoTelefonoModel"
	"appbiomedica/Modelos/ModeloUnidadM"
	"appbiomedica/Modelos/ProveedoresModel"
	"appbiomedica/Modelos/ServiciosSubrogados/TrasladoPacientesModel"
	"appbiomedica/Modulos/Estructuras"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/mgo.v2/bson"
)

var validate *validator.Validate

type ContratoNuevo struct {
	Contrato            ContratoModelo.Contrato
	Delegaciones        []string                          `gorm:"-"`
	AdminAuxContrato    []AdminAuxModelo.AdminAuxContrato `gorm:"-"`
	PenasConvencionales PenasModelo.PenasConvencionales
	RepresentanteLegal  RepresentanteModel.RepresentanteLegal
	PenasDeductivas     []PenasModelo.PenasDeductivas `gorm:"-"`
	Fianza              FianzaModelo.Fianza           `json:"-"`
}

type UnContrato struct {
	Contrato      ContratoModelo.Contrato
	Delegaciones  []string `gorm:"-"`
	Representante struct {
		NombreCompleto string
		Correos        []struct {
			Correo     string
			TipoCorreo string
		} `gorm:"-"`
		Telefonos []struct {
			Extension    string
			Telefono     string
			TipoTelefono string
		} `gorm:"-"`
	}
}

type PromedioAcompaniante struct {
	UnidadMedica ModeloUnidadM.UnidadMed
	Promedio     TrasladoPacientesModel.Promedio
}

//********************************************************************************************************
/* *********************    Controlador para consultar uno o todos los contratos     ********************* */
func GetContratos(c *gin.Context) {
	id_contrato := c.DefaultQuery("id_contrato", "")
	filtros := c.DefaultQuery("filtro", "")
	proveedores := c.DefaultQuery("proveedores", "")
	tipo := c.DefaultQuery("tipo", "")
	delegaciones := c.DefaultQuery("delegaciones", "")
	num_pagina := c.DefaultQuery("pagina", "1")
	num_registro := c.DefaultQuery("num_registros", "15")
	arrayFiltro := []string{}
	arreglo_provee := []string{}
	arreglo_tipo := []string{}
	arreglo_deleg := []string{}
	num_pag, err := strconv.Atoi(num_pagina)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Numero de pagina incorrecto", err.Error(), c)
		return
	}
	num_reg, err := strconv.Atoi(num_registro)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Numero de registros incorrecto", err.Error(), c)
		return
	}
	if filtros != "" {
		arrayFiltro = strings.Split(filtros, " ")
	}
	if proveedores != "" {
		arreglo_provee = strings.Split(proveedores, ",")
	}
	if tipo != "" {
		arreglo_tipo = strings.Split(tipo, ",")
	}
	if delegaciones != "" {
		arreglo_deleg = strings.Split(delegaciones, ",")
	}
	err, resultado, total := ContratoModelo.ConsultarContrato(id_contrato, arreglo_tipo, arrayFiltro, arreglo_provee,
		arreglo_deleg,
		num_pag, num_reg)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", gin.H{"total": total, "registros": resultado}, c)
	return
}

/*****  Controlador para consultar contratos que tengan servicios subrogados de traslado de pacientes  *****/
func GetContratosTraslado(ctx *gin.Context) {
	unidad_id := ctx.DefaultQuery("unidad_id", "")
	fecha_cita := ctx.DefaultQuery("fecha_cita", "")
	contratos, err := ContratoModelo.RegresaContratoTraslado(unidad_id)
	var resultado []ContratoModelo.Contrato
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), ctx)
		return
	}
	/*****  Validar la fecha de la cita con la del contrato  *****/
	temporalDate, _ := time.Parse(time.RFC3339, fecha_cita)
	medicalDate, _ := time.Parse("02/01/2006", temporalDate.Format("02/01/2006"))
	fmt.Println("-------", medicalDate, temporalDate.Hour())
	for _, fecha := range *contratos {
		contractEndDate, _ := time.Parse("02/01/2006", fecha.FinContrato)
		contractStartDate, _ := time.Parse("02/01/2006", fecha.InicioContrato)
		if (medicalDate.Before(contractEndDate) || medicalDate.Equal(contractEndDate)) &&
			(medicalDate.After(contractStartDate) || medicalDate.Equal(contractStartDate)) {
			resultado = append(resultado, fecha)
		}
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, ctx)
}

type MayorSolicitud struct {
	Unidad     ModeloUnidadM.UnidadMed `json:"unidad"`
	Porcentaje float32                 `json:"porcentaje"`
}
type ContratoDevengo struct {
	Contrato                       ContratoModelo.Contrato `json:"contrato,omitempty"`
	Devengo                        float64                 `json:"devengo,omitempty"`
	PorcentajeDevengo              float64                 `json:"porcentajeDevengo,omitempty"`
	Solicitudes                    int64                   `json:"solicitudes"`
	ConAcompaniante                int32                   `json:"conAcompaniante"`
	SinAcompaniante                int32                   `json:"sinAcompaniante"`
	PromedioAcompaniante           PromedioAcompaniante    `json:"promedioAcompaniante"`
	PorcentajeAcompanianteContrato float32                 `json:"porcentajeAcompanianteContrato"`
	MayorSolicitud                 MayorSolicitud          `json:"mayorSolicitud"`
}

func GetContratosTrasladoSubtipo(ctx *gin.Context) {
	subTipo := ctx.DefaultQuery("subTipo", "")
	contratos, err := ContratoModelo.ConsultarContratoSubTipo(subTipo)
	contratosVigentes := make([]ContratoDevengo, 0)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "error al consultar los contratos", err, ctx)
		return
	}
	if len(*contratos) == 0 {
		Estructuras.Responder(http.StatusInternalServerError, "no hay contratos vigentes", errors.New("no hay contratos"), ctx)
		return
	}
	fechaActual, _ := time.Parse("02/01/2006", time.Now().Format("02/01/2006"))
	for _, contrato := range *contratos {
		fechaFin, _ := time.Parse("02/01/2006", contrato.Contrato.FinContrato)
		fechaInicio, _ := time.Parse("02/01/2006", contrato.Contrato.InicioContrato)
		// Verificar los contratos vigentes a la fecha actual
		if (fechaActual.Before(fechaFin) || fechaActual.Equal(fechaFin)) &&
			(fechaActual.After(fechaInicio) || fechaActual.Equal(fechaInicio)) {
			// Consultar el devengo por contrato y calcular el iva
			var promedioDevengo float64
			var porcentajeAcompanianteContrato float32
			err, earn := TrasladoPacientesModel.ConsultEarn(true, 0, contrato.ContratoNumeroContrato)
			if err != nil {
				Estructuras.Responder(http.StatusInternalServerError, "error al consultar el devengo", err.Error(), ctx)
				return
			}
			earn = earn * 1.16
			promedioDevengo = (earn / (contrato.Contrato.MontoTotal * 1.16)) * 100
			// Consultar las solicitudes realizadas por contrato
			solTotal, err := TrasladoPacientesModel.ContarSolicitudesRealizadas(contrato.ContratoNumeroContrato, true)
			if err != nil {
				Estructuras.Responder(http.StatusInternalServerError, "error al consultar el total de solicitudes", err.Error(), ctx)
				return
			}
			// Consultar las solicitudes realizadas con acompañantes por contrato
			Acompaniante, err := TrasladoPacientesModel.ContarSolicitudesAcompaniante(contrato.ContratoNumeroContrato, true, true)
			if err != nil {
				Estructuras.Responder(http.StatusInternalServerError, "error al consultar el total de solicitudes con acompañante", err.Error(), ctx)
				return
			}
			if solTotal != 0 {
				porcentajeAcompanianteContrato = float32(Acompaniante) / float32(solTotal)
			}
			// Consultar las solicitudes realizadas sin acompañantes por contrato
			sinAcompaniante, err := TrasladoPacientesModel.ContarSolicitudesAcompaniante(contrato.ContratoNumeroContrato, false, true)
			if err != nil {
				Estructuras.Responder(http.StatusInternalServerError, "error al consultar el total de solicitudes sin acompañante", err.Error(), ctx)
				return
			}
			// Consultar la unidad medica con mayor porcentaje de solicitudes con acompañantes
			result, err := TrasladoPacientesModel.PorcentajeMayorAcompaniante(contrato.ContratoNumeroContrato)
			if err != nil {
				Estructuras.Responder(http.StatusInternalServerError, "error al consultar datos",
					err.Error(), ctx)
				return
			}
			var unidadResult PromedioAcompaniante
			if len(result) > 0 {
				err, Unidad := ModeloUnidadM.ConsultaUnaUnidadMed(result[0].UnidadMedId)
				if err != nil {
					Estructuras.Responder(http.StatusInternalServerError, "error al consultar datos",
						err.Error(), ctx)
					return
				}

				unidadResult.UnidadMedica = *Unidad
				unidadResult.Promedio = result[0]
			}
			mayorSolicitud, err := TrasladoPacientesModel.UnidadMayorSolicitudes(contrato.ContratoNumeroContrato, true)
			if err != nil {
				Estructuras.Responder(http.StatusInternalServerError, "error al consultar datos",
					err.Error(), ctx)
				return
			}
			fmt.Println(mayorSolicitud)
			var mayor MayorSolicitud
			if mayorSolicitud.Unidad != 0 {
				err, Unidad := ModeloUnidadM.ConsultaUnaUnidadMed(mayorSolicitud.Unidad)
				if err != nil {
					Estructuras.Responder(http.StatusInternalServerError, "error al consultar datos",
						err.Error(), ctx)
					return
				}
				mayor.Unidad = *Unidad
				mayor.Porcentaje = float32(mayorSolicitud.Total) / float32(solTotal)
			}
			temp := ContratoDevengo{Contrato: *contrato.Contrato, Devengo: earn, PorcentajeDevengo: promedioDevengo, Solicitudes: solTotal, ConAcompaniante: Acompaniante, SinAcompaniante: sinAcompaniante, PromedioAcompaniante: unidadResult, PorcentajeAcompanianteContrato: porcentajeAcompanianteContrato, MayorSolicitud: mayor}
			contratosVigentes = append(contratosVigentes, temp)
		}
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", contratosVigentes, ctx)
}

type DetalleTablero struct {
	Devengo      []TrasladoPacientesModel.DevengoPorMes        `json:"devengo"`
	Acompaniante []TrasladoPacientesModel.AcompañantePorUnidad `json:"acompaniante"`
}

func GetDetallesContrato(ctx *gin.Context) {
	contratos := ctx.DefaultQuery("contratos", "")
	fmt.Println(contratos)
	var resultado DetalleTablero
	result, err := TrasladoPacientesModel.ConsultarDevengoPorMes(true, 0, contratos)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "error al consultar datos",
			err.Error(), ctx)
		return
	}
	resultado.Devengo = result
	resultAcompaniante, err := TrasladoPacientesModel.ContarSolicitudesAcompaniantePorUnidad(contratos, true, true)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "error al consultar datos",
			err.Error(), ctx)
		return
	}
	if len(resultAcompaniante) > 0 {
		for index, item := range resultAcompaniante {
			err, Unidad := ModeloUnidadM.ConsultaUnaUnidadMed(item.Numero)
			if err != nil {
				Estructuras.Responder(http.StatusInternalServerError, "error al consultar datos",
					err.Error(), ctx)
				return
			}
			resultAcompaniante[index].Unidad = *Unidad
		}
	}
	resultado.Acompaniante = resultAcompaniante
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, ctx)
}

//********************************************************************************************************
/* *********************    Controlador para consultar los proveedores de los contratos     ********************* */
func GetProveedoresContrato(c *gin.Context) {
	resultado, err := ContratoModelo.RegresaProveedoresContratos()
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
}

//********************************************************************************************************
/* *********************    Controlador para validar y verificar un archivo de carga    ********************* */
func VerificarContratoFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	f, err := file.Open()
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al abrir el archivo", err, c)
		return
	}
	defer f.Close()
	archivo, err := excelize.OpenReader(f)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al leer el archivo", err, c)
		return
	}
	rows, _ := archivo.GetRows("Contratos")
	fmt.Println("--------", rows)
	//if err != nil {
	//	Estructuras.Responder(http.StatusInternalServerError, "No se pudo leer el archivo o no tiene el formato correcto", err, c)
	//	return
	//}
	var errores []ContratoModelo.RespuestaContrato
	var cabeceras []string
	for num, line := range rows {
		fmt.Print("**-*-", line)
		if num == 0 {
			cabeceras = line
		}
		if num > 0 {
			fmt.Println("Tamaño del arreglo.......----{{", len(line), "-*******")
			if ContratoModelo.CompruebaContrato(line[0]) {
				mns := "Error este contrato ya ha sido registrado"
				n := ContratoModelo.RespuestaContrato{
					NumContra: line[0], Error: mns, Columna: cabeceras[0], Linea: num + 1, Tipo: 1}
				errores = append(errores, n)
			}
			if !ProveedoresModel.VerificaProveedor(line[7]) {
				mns := "Este proveedor no esta registrado"
				n := ContratoModelo.RespuestaContrato{
					NumContra: line[0], Error: mns, Columna: cabeceras[7], Linea: num + 1, Tipo: 2}
				errores = append(errores, n)
			}
			_, err := strconv.ParseFloat(line[6], 32)
			if err != nil {
				fmt.Println("error monto", err)
				mensaje := "Monto Total formato incorrecto"
				n := ContratoModelo.RespuestaContrato{
					NumContra: line[0], Error: mensaje, Columna: cabeceras[6], Linea: num + 1, Tipo: 0}
				errores = append(errores, n)
			}
			fmt.Println("Fecha inicio", line[1])
			if !govalidator.IsTime(line[1], "02/01/2006") {
				mensaje := "Inicio de contrato, formato incorrecto tiene que ser " + "(dd/mm/aaaa)"
				n := ContratoModelo.RespuestaContrato{
					NumContra: line[0], Error: mensaje, Columna: cabeceras[1], Linea: num + 1, Tipo: 0}
				errores = append(errores, n)
			}
			fmt.Println("Fecha fin", line[2])
			if !govalidator.IsTime(line[2], "02/01/2006") {
				mensaje := "Fin de contrato, formato incorrecto tiene que ser " + "(dd/mm/aaaa)"
				n := ContratoModelo.RespuestaContrato{
					NumContra: line[0], Error: mensaje, Columna: cabeceras[2], Linea: num + 1, Tipo: 0}
				errores = append(errores, n)
			}
			fmt.Println("Fecha fallo", line[5])
			if !govalidator.IsTime(line[5], "02/01/2006") {
				mensaje := "Fecha de fallo, formato incorrecto tiene que ser " + "(dd/mm/aaaa)"
				n := ContratoModelo.RespuestaContrato{
					NumContra: line[0], Error: mensaje, Columna: cabeceras[5], Linea: num + 1, Tipo: 0}
				errores = append(errores, n)
			}
			/*			validate = validator.New()
						if line[9] == "" {
							mensaje := "El nombre del representante legal no puede estar vacio"
							n := ContratoModelo.RespuestaContrato{
								NumContra: line[0], Error: mensaje, Columna: cabeceras[9], Linea: num + 1, Tipo: 0}
							errores = append(errores, n)
						}
						correos := strings.Split(line[12], ",")
						for _, correo := range correos {
							errs := validate.Var(strings.TrimSpace(correo), "email")
							if errs != nil {
								fmt.Println(errs, correo) // output: Key: "" Error:Field validation for "" failed on the "email" tag
								mensaje := "Formato de correo invalido"
								n := ContratoModelo.RespuestaContrato{
									NumContra: line[0], Error: mensaje, Columna: cabeceras[12], Linea: num + 1, Tipo: 0}
								errores = append(errores, n)
							}
						}
						telefonos := strings.Split(line[10], ",")
						for _, tel := range telefonos {
							if tel != "" {
								errs := validate.Var(strings.TrimSpace(tel), "numeric")
								if errs != nil {
									fmt.Println(errs, tel) // output: Key: "" Error:Field validation for "" failed on the "email" tag
									mensaje := "Formato de telefono invalido"
									n := ContratoModelo.RespuestaContrato{
										NumContra: line[0], Error: mensaje, Columna: cabeceras[10], Linea: num + 1, Tipo: 0}
									errores = append(errores, n)
								}
							}
						}*/
		}
	}
	fmt.Println("encabezados-.-.-.-.-.-.", cabeceras, "...-.", len(cabeceras))
	if len(errores) > 0 {
		Estructuras.Responder(http.StatusInternalServerError, "Verifica tus datos", errores, c)
		return
	} else {
		Estructuras.Responder(http.StatusOK, "Validacion Exitosa", errores, c)
		return
	}
}

//********************************************************************************************************
/* **************    Controlador para GUARDAR EN LA BASE DE DATOS LOS CONTRATOS DESDE UN ARCHIVO   **************** */
func PostContratoFile(c *gin.Context) {
	tipo := c.DefaultQuery("tipo_contrato", "")
	file, _ := c.FormFile("file")
	f, err := file.Open()
	if err != nil {
		return
	}
	defer f.Close()
	fmt.Println(file.Filename)
	archivo, err := excelize.OpenReader(f)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al leer el archivo", err, c)
		return
	}
	rows, err := archivo.GetRows("Contratos")
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "No se pudo leer el archivo o no tiene el formato correcto", err, c)
		return
	}
	//********************************************************************************************************
	/* *********************    inicio mi transaccion de mi base de datos     ********************* */
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al iniciar la transaccion", tx.Error, c)
		return
	}
	/*Declarando las variables de las estructuras y las tablas las cuales se ocuparan o se registraran*/
	var contrato ContratoModelo.Contrato
	/*	var representante RepresentanteModel.RepresentanteLegal
		var correo CorreoTelefonoModel.Correos
		var telefono CorreoTelefonoModel.Telefonos
		var correContrato ContactoModelo.CorreosContrato
		var telefonoContrato ContactoModelo.TelefonoContrato*/
	/*for para iterar al archivo fila por fila e ir agregando los registros necesaerios en cada tabla*/
	for num, line := range rows {
		if num > 0 {
			/*Verificamos que no se duplique el numero de contrato*/
			if ContratoModelo.CompruebaContrato(Estructuras.ConvertLowercase(line[0])) {
				Estructuras.Responder(http.StatusInternalServerError, "Error este contrato ya ha sido registrado", err.Error(), c)
				return
			}
			/*Agregar el contrato primero*/
			total, _ := strconv.ParseFloat(line[6], 32)
			contrato.Id = bson.NewObjectId().Hex()
			contrato.NumeroContrato = Estructuras.FormatoTextoMayusculas(line[0])
			contrato.InicioContrato = line[1]
			contrato.FinContrato = line[2]
			contrato.ProcedContratacion = line[3]
			contrato.TipoProcedContratacion = line[4]
			contrato.FechaFallo = line[5]
			contrato.MontoTotal = float64(total)
			contrato.TipoContratoId = tipo
			contrato.ProveedorNProvImss = line[7]
			contrato.Estado = true
			fmt.Println(contrato)
			err = contrato.AgregarContrato(tx)
			if err != nil {
				tx.Rollback()
				Estructuras.Responder(http.StatusInternalServerError, "Error al agregar el contrato", err.Error(), c)
				return
			}
			/*Agregar el representante legal*/
			/*			representante.Id = bson.NewObjectId().Hex()
						representante.NombreCompleto = strings.Title(strings.ToLower(line[9]))
						representante.ContratoNumeroContrato = contrato.NumeroContrato
						err = representante.AgregarRepresentanteLegal(tx)
						if err != nil {
							tx.Rollback()
							Estructuras.Responder(http.StatusInternalServerError, "Error al agregar el representante legal", err.Error(), c)
							return
						}*/
			/*Agregar los correos */
			//ListaCorreos := strings.Split(line[12], ",")
			//for _, item := range ListaCorreos {
			/*Busco el correo en la bd*/
			//_, resultado := CorreoTelefonoModel.BuscarCorreo(strings.TrimSpace(item))
			/*Si lo encuentro */
			/*				if resultado != nil {
							correContrato.CorreosId = resultado.Id
						} else {*/
			/*Si no lo encuentro lo agrego*/
			/*				correo.Correo = item
							correo.Id = bson.NewObjectId().Hex()
							err = correo.AgregaCorreo(tx)
							if err != nil {
								tx.Rollback()
								Estructuras.Responder(http.StatusInternalServerError, "Error al agregar el correo en la base de datos", err.Error(), c)
								return
							}
							correContrato.CorreosId = correo.Id
						}*/
			/*Agrego el correo al contrato*/
			/*				correContrato.Id = bson.NewObjectId().Hex()
							correContrato.TipoCorreo = "Representante Legal"
							correContrato.ContratoNumeroContrato = contrato.NumeroContrato
							err = correContrato.AgregaCorreoContrato(tx)
							if err != nil {
								tx.Rollback()
								Estructuras.Responder(http.StatusInternalServerError, "Error al asignar los correos al contrato", err.Error(), c)
								return
							}
						}*/
			/*Agregar los telefonos */
			/*			ListaTelefonos := strings.Split(line[10], ",")
						for _, item := range ListaTelefonos {
							_, resultado := CorreoTelefonoModel.BuscarTelefono(strings.TrimSpace(item))
							if resultado != nil {
								telefonoContrato.TelefonosId = resultado.Id
							} else {
								telefono.Telefono = item
								telefono.Id = bson.NewObjectId().Hex()
								err = telefono.AgregaTelefono(tx)
								if err != nil {
									tx.Rollback()
									Estructuras.Responder(http.StatusInternalServerError, "Error al agregar el telefono en la base de datos", err.Error(), c)
									return
								}
								telefonoContrato.TelefonosId = telefono.Id
							}*/
			/*Asigno el telefono al correo*/
			/*				telefonoContrato.Id = bson.NewObjectId().Hex()
							telefonoContrato.TipoTelefono = "Representante Legal"
							telefonoContrato.ContratoNumeroContrato = contrato.NumeroContrato
							//telefonoContrato.Extension
							err = telefonoContrato.AgregaTelContrato(tx)
							if err != nil {
								tx.Rollback()
								Estructuras.Responder(http.StatusInternalServerError, "Error al asignar el telefono al contrato", err.Error(), c)
								return
							}*/
		}
	}
	//********************************************************************************************************
	/* *********************   finalizar insercion y hacer el commit     ********************* */
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al realizar el commit", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos agregados correctamente", contrato, c)
	return
}

//********************************************************************************************************
/* *********************    Controlador para Agregar un nuevo contrato     ********************* */
func PostContrato(c *gin.Context) {
	var contrato ContratoNuevo
	err := c.BindJSON(&contrato)
	fmt.Println(contrato)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error al castear la estructura", err.Error(), c)
		return
	}
	fmt.Println(&contrato)
	//********************************************************************************************************
	/* *********************    inicio mi transaccion de mi base de datos     ********************* */
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al iniciar la transaccion", tx.Error, c)
		return
	}
	//********************************************************************************************************
	/* *********************   Validar que el contrato no exista     ********************* */
	if ContratoModelo.CompruebaContrato(contrato.Contrato.NumeroContrato) {
		Estructuras.Responder(http.StatusInternalServerError, "Error este contrato ya ha sido registrado", err.Error(), c)
		return
	}
	//********************************************************************************************************
	/* *********************    creo mi variabale de contrato para agregar en la bd     ********************* */
	var contratos ContratoModelo.Contrato
	contratos = contrato.Contrato
	contratos.Id = bson.NewObjectId().Hex()
	contratos.Estado = true
	err = contratos.AgregarContrato(tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al agregar el contrato", err.Error(), c)
		return
	}
	//********************************************************************************************************
	/* *********************    insertar delegaciones por contrato    ********************* */

	var delegaciones ContratoModelo.Contrato_Delegaciones
	err = delegaciones.AgregarDelegacionContrato(contratos.NumeroContrato, contrato.Delegaciones, tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al agregar delegaciones contrato", err.Error(), c)
		return
	}
	//********************************************************************************************************
	/* *********************    Insertar administradores y auxiliares     ********************* */
	var administrador AdminAuxModelo.AdminAuxContrato
	for _, admi := range contrato.AdminAuxContrato {
		administrador.ContratoNumeroContrato = contratos.NumeroContrato
		administrador.Estado = true
		administrador.Id = bson.NewObjectId().Hex()
		administrador.Cargo = admi.Cargo
		administrador.Responsabilidad = admi.Responsabilidad
		administrador.PersonaId = admi.PersonaId
		administrador.DelegacionesClvDele = admi.DelegacionesClvDele
		err = administrador.AgregarAdminAux(tx)
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "Error al agregar administrador o auxiliar", err.Error(), c)
			return
		}
	}
	//********************************************************************************************************
	/* *********************    Insertar Pena convencional      ********************* */
	var penaConvencional PenasModelo.PenasConvencionales
	penaConvencional = contrato.PenasConvencionales
	penaConvencional.Id = bson.NewObjectId().Hex()
	penaConvencional.ContratoNumeroContrato = contratos.NumeroContrato
	penaConvencional.Estado = true
	err = penaConvencional.AgregarPenaConv(tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al agregar la pena convencional", err.Error(), c)
		return
	}
	//********************************************************************************************************
	/* *********************    Insertar Penas Deductivas     ********************* */
	if len(contrato.PenasDeductivas) != 0 {
		var penasDeductivas PenasModelo.PenasDeductivas
		for _, pena := range contrato.PenasDeductivas {
			penasDeductivas.ContratoNumeroContrato = contratos.NumeroContrato
			penasDeductivas.Estado = true
			penasDeductivas.Id = bson.NewObjectId().Hex()
			penasDeductivas.ConceptoObligacion = pena.ConceptoObligacion
			penasDeductivas.NivelServicio = pena.NivelServicio
			penasDeductivas.UnidadMedida = pena.UnidadMedida
			penasDeductivas.Deduccion = pena.Deduccion
			penasDeductivas.DescripcionDeduccion = pena.DescripcionDeduccion
			penasDeductivas.LimiteIncumplimiento = pena.LimiteIncumplimiento
			err = penasDeductivas.AgregarPenaDed(tx)
			if err != nil {
				tx.Rollback()
				Estructuras.Responder(http.StatusInternalServerError, "Error al agregar la pena deductiva", err.Error(), c)
				return
			}
		}
	}
	//********************************************************************************************************
	/* *********************   Insertar fianza     ********************* */
	if contrato.Fianza.Afianzadora != "" {
		var fianza FianzaModelo.Fianza
		fianza = contrato.Fianza
		fianza.ContratoNumeroContrato = contratos.NumeroContrato
		fianza.Id = bson.NewObjectId().Hex()
		fianza.Estado = true
		err = fianza.Agregarfianza(tx)
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "Error al agregar la fianza", err.Error(), c)
			return
		}
	}
	//********************************************************************************************************
	/* *********************    Insertar al representante legal     ********************* */
	if contrato.RepresentanteLegal.ContratoNumeroContrato != "" {
		var representante RepresentanteModel.RepresentanteLegal
		representante = contrato.RepresentanteLegal
		representante.ContratoNumeroContrato = contratos.NumeroContrato
		strings.Title(strings.ToLower(representante.NombreCompleto))
		representante.Id = bson.NewObjectId().Hex()
		err = representante.AgregarRepresentanteLegal(tx)
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "Error al agregar el representante legal", err.Error(), c)
			return
		}
	}
	//********************************************************************************************************
	/* *********************   finalizar insercion y hacer el commit     ********************* */
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al realizar el commit", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos agregados correctamente", contrato, c)
}

//********************************************************************************************************
/* *********************    Controlador para actualizar un contrato ya registrado     ********************* */
func UpdateContrato(c *gin.Context) {
	var contrato ContratoModelo.Contrato
	id_contrato := c.DefaultQuery("id_contrato", "")
	err := c.BindJSON(&contrato)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error al castear la estructura", err.Error(), c)
		return
	}
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al conectar a la base de datos", tx.Error, c)
		return
	}
	if err = contrato.ActualizarContrato(id_contrato, tx); err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al actualizar los datos", err.Error(), c)
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al actualizar el proveedor", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos actualizados correctamente", contrato, c)

}

//********************************************************************************************************
/* *********************    Controlador para eliminar un contrato     ********************* */
func DeleteContrato(c *gin.Context) {
	id_contrato := c.DefaultQuery("id_contrato", "")
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al conectar a la base de datos", tx.Error, c)
		return
	}
	err, resultado := ContratoModelo.EliminarContrato(id_contrato, tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al eliminar los datos", err.Error(), c)
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al eliminar el contrato", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Contrato Eliminado Correctamente", resultado, c)

}

//********************************************************************************************************
/* *********************    Controlador para Agregar un nuevo contrato     ********************* */
func PostUnContrato(c *gin.Context) {
	var contrato UnContrato
	err := c.BindJSON(&contrato)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error al castear la estructura", err.Error(), c)
		return
	}
	fmt.Println(&contrato.Contrato)
	fmt.Println(&contrato.Delegaciones)
	fmt.Println(&contrato.Representante)
	//********************************************************************************************************
	/* *********************    inicio mi transaccion de mi base de datos     ********************* */
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al iniciar la transaccion", tx.Error, c)
		return
	}
	//********************************************************************************************************
	/* *********************    creo mi variabale de contrato para agregar en la bd     ********************* */
	var contratos ContratoModelo.Contrato

	contratos = contrato.Contrato
	contratos.Id = bson.NewObjectId().Hex()
	contratos.Estado = true
	if ContratoModelo.CompruebaContrato(contratos.NumeroContrato) {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Este Numero de contrato ya esta registrado",
			contratos, c)
		return
	}
	err = contratos.AgregarContrato(tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al agregar el contrato", err.Error(), c)
		return
	}
	//********************************************************************************************************
	/* *********************    insertar delegaciones por contrato    ********************* */

	var delegaciones ContratoModelo.Contrato_Delegaciones
	err = delegaciones.AgregarDelegacionContrato(contratos.NumeroContrato, contrato.Delegaciones, tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al agregar delegaciones contrato", err.Error(), c)
		return
	}
	//********************************************************************************************************
	/* *********************    Insertar al representante legal     ********************* */
	var representante RepresentanteModel.RepresentanteLegal
	representante.ContratoNumeroContrato = contratos.NumeroContrato
	representante.NombreCompleto = strings.Title(strings.ToLower(contrato.Representante.NombreCompleto))
	representante.Id = bson.NewObjectId().Hex()
	err = representante.AgregarRepresentanteLegal(tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al agregar el representante", err.Error(), c)
		return
	}
	//********************************************************************************************************
	/* *********************    Insertar el correo o los correos electronicos  al contrato   ********************* */
	var correo_contrato ContactoModelo.CorreosContrato
	var correo CorreoTelefonoModel.Correos
	for _, email := range contrato.Representante.Correos {
		/*Busco si el correo ya existe en la base de datos*/
		repetido, resultado := CorreoTelefonoModel.BuscarCorreo(email.Correo)
		/*Si lo encuentro ya no lo agrego y le asigno su id para relacionarlo al contrato*/
		if repetido {
			correo_contrato.CorreosId = resultado.Id
		} else {
			/*Si no lo encuentro lo agrego*/
			correo.Correo = email.Correo
			correo.Id = bson.NewObjectId().Hex()
			err = correo.AgregaCorreo(tx)
			if err != nil {
				tx.Rollback()
				Estructuras.Responder(http.StatusInternalServerError, "Error al agregar el correo en la base de datos", err.Error(), c)
				return
			}
			correo_contrato.CorreosId = correo.Id
		}
		/*Agrego el correo al contrato*/
		correo_contrato.Id = bson.NewObjectId().Hex()
		correo_contrato.TipoCorreo = email.TipoCorreo
		correo_contrato.ContratoNumeroContrato = contratos.NumeroContrato
		err = correo_contrato.AgregaCorreoContrato(tx)
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "Error al asignar los correos al contrato", err.Error(), c)
			return
		}
	}
	//********************************************************************************************************
	/* *********************    Insertar el telefono o telefonos al contrato     ********************* */
	var telefono_contrato ContactoModelo.TelefonoContrato
	var telefono CorreoTelefonoModel.Telefonos
	for _, tel := range contrato.Representante.Telefonos {
		/*Busco si el numero de telefono ya existe en la base de datos*/
		repetido, resultado := CorreoTelefonoModel.BuscarTelefono(tel.Telefono)
		/*Si lo encuentro ya no lo agrego y le asigno su id para relacionarlo al contrato*/
		if repetido {
			telefono_contrato.TelefonosId = resultado.Id
		} else {
			/*Si no lo encuentro lo agrego*/
			telefono.Telefono = tel.Telefono
			telefono.Id = bson.NewObjectId().Hex()
			err = telefono.AgregaTelefono(tx)
			if err != nil {
				tx.Rollback()
				Estructuras.Responder(http.StatusInternalServerError,
					"Error al agregar el telefono en la base de datos", err.Error(), c)
				return
			}
			telefono_contrato.TelefonosId = telefono.Id
		}
		/*Agrego el telefono al contrato*/
		telefono_contrato.Id = bson.NewObjectId().Hex()
		telefono_contrato.TipoTelefono = tel.TipoTelefono
		telefono_contrato.ContratoNumeroContrato = contratos.NumeroContrato
		telefono_contrato.Extension = tel.Extension
		err = telefono_contrato.AgregaTelContrato(tx)
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "Error al asignar los telefonos al contrato",
				err.Error(), c)
			return
		}
	}
	//********************************************************************************************************
	/* *********************   finalizar insercion y hacer el commit     ********************* */
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al realizar el commit", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos agregados correctamente", contrato, c)
}

// Funcion que regresa la unidad medica con el mayor porcentaje de solicitudes con acompañante realizada
func GetPorcentajeMayorAcompaniante(ctx *gin.Context) {
	numContrato := ctx.DefaultQuery("numContrato", "")
	result, err := TrasladoPacientesModel.PorcentajeMayorAcompaniante(numContrato)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "error al consultar datos",
			err.Error(), ctx)
		return
	}
	err, Unidad := ModeloUnidadM.ConsultaUnaUnidadMed(result[0].UnidadMedId)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "error al consultar datos",
			err.Error(), ctx)
		return
	}
	unidadResult := PromedioAcompaniante{
		UnidadMedica: *Unidad,
		Promedio:     result[0],
	}
	Estructuras.Responder(http.StatusOK, "datos consultados", unidadResult, ctx)
}
