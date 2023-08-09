package TrasladosControlador

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/ServiciosSubrogados/TrasladoPacientesModel"
	"appbiomedica/Modulos/Estructuras"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

type Boletos struct {
	Ida           bool
	Observaciones string
	Regreso       bool
	Tipo          int
	TipoBoleto    int
	Total         int
}
type SolicitudtrasladoBoletos struct {
	Solicitud                TrasladoPacientesModel.SolicitudTraslado
	Boletos                  []Boletos
	Acompaniante             TrasladoPacientesModel.Acompaniante
	AutorizacionesEspeciales TrasladoPacientesModel.AutorizacionesEspeciales
}

func ReturnTicketType(numberTicket int) string {
	if numberTicket != 0 {
		if numberTicket == 1 {
			return "Entero"
		} else {
			return "Medio"
		}
	} else {
		return "Sin Boleto"
	}
}
func ReturnObservation(obs string) string {
	if obs == "" {
		return "---"
	} else {
		return obs
	}
}

func AgregarSolicitud(solicitudBoletos SolicitudtrasladoBoletos, fechaCita, uiPrei string, firms []string, folios string) (int, string, interface{}, error) {
	solicitudBoletos.Solicitud.FechaExpedicion = time.Now()
	t, _ := time.Parse(time.RFC3339, fechaCita)
	solicitudBoletos.Solicitud.FechaCita = t
	nuevoAnio := t.Year()
	digitos := strconv.Itoa(nuevoAnio)
	/*****  Funcion para saber que numero de folio consecutivo sigue  *****/
	err, consecutivo := TrasladoPacientesModel.BuscarFolioConsecutivo(solicitudBoletos.Solicitud.UnidadMedId, nuevoAnio)
	if err != nil {
		return http.StatusInternalServerError, "Error al consultar el folio consecutivo", nil, err
	}
	//Buscar que el folio FTP no se encuentre registrado y si se encuentra registrado que la solicitud este cancelada
	if folios == "" {
		duplicado, _ := TrasladoPacientesModel.BuscarFolioFtp(solicitudBoletos.Solicitud.Ftp01, "true")
		if duplicado {
			return http.StatusInternalServerError,
				"Este folio FTP01: " + solicitudBoletos.Solicitud.Ftp01 + " ya se encuentra registrado", nil, err
		}
	}
	/*****  Inicializar la transaccion de la BD  *****/
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		return http.StatusInternalServerError, "Error al iniciar la transaccion", nil, tx.Error
	}
	/*****  Agregar la solicitud de traslado  *****/
	folio_sig := strconv.Itoa(consecutivo + 1)
	solicitudBoletos.Solicitud.Folio = uiPrei + "-" + digitos[2:] + "-" + folio_sig
	solicitudBoletos.Solicitud.Estado = true
	err = solicitudBoletos.Solicitud.AgregarSolicitud(tx)
	if err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, "Error al insertar los datos", nil, err
	}
	/*****  Agregar los boletos   *****/
	var boletos TrasladoPacientesModel.Boletos
	var rutas TrasladoPacientesModel.RutasConsumidas
	for _, registro := range solicitudBoletos.Boletos {
		boletos.Id = bson.NewObjectId().Hex()
		boletos.Ida = registro.Ida
		boletos.Regreso = registro.Regreso
		boletos.Total = registro.Total
		boletos.Observaciones = ReturnObservation(registro.Observaciones)
		boletos.Tipo = ReturnTicketType(registro.TipoBoleto)
		boletos.TipoPasajeroId = registro.Tipo
		boletos.SolicitudTrasladoFolio = solicitudBoletos.Solicitud.Folio
		boletos.Estado = true
		err = boletos.AgregarBoletosTraslado(tx)
		if err != nil {
			tx.Rollback()
			return http.StatusInternalServerError, "Error al agregar los boletos", nil, err
		}
		/*****  Validar y agregar la ruta consumida  *****/
		if registro.TipoBoleto != 0 {
			if registro.Ida {
				_, res := TrasladoPacientesModel.SearchTicket(solicitudBoletos.Solicitud.UnidadMedId,
					solicitudBoletos.Solicitud.Origen, solicitudBoletos.Solicitud.Destino,
					ReturnTicketType(registro.TipoBoleto), solicitudBoletos.Solicitud.ContratoNumeroContrato)
				if res != nil {
					rutas.TrasladoPacientesId = res.Id
				} else {
					tx.Rollback()
					return http.StatusInternalServerError, "No se tiene contratado el tipo de boleto (" + ReturnTicketType(
						registro.TipoBoleto) + ") a esta ruta -->" + solicitudBoletos.Solicitud.Origen + " --- " +
						solicitudBoletos.Solicitud.Destino, nil, errors.New("no se encuentra el boleto")
				}
				rutas.Id = bson.NewObjectId().Hex()
				rutas.Cantidad = 1
				rutas.BoletosId = boletos.Id
				rutas.Estado = true
				err = rutas.AgregarRutaConsumida(tx)
				if err != nil {
					tx.Rollback()
					return http.StatusInternalServerError, "Error al agregar la ruta", nil, err
				}
			}
			if registro.Regreso {
				_, res := TrasladoPacientesModel.SearchTicket(solicitudBoletos.Solicitud.UnidadMedId,
					solicitudBoletos.Solicitud.Destino, solicitudBoletos.Solicitud.Origen,
					ReturnTicketType(registro.TipoBoleto), solicitudBoletos.Solicitud.ContratoNumeroContrato)
				if res != nil {
					rutas.TrasladoPacientesId = res.Id
				} else {
					tx.Rollback()
					return http.StatusInternalServerError, "No se tiene contratado el tipo de boleto (" + ReturnTicketType(
						registro.TipoBoleto) + ") a esta ruta -->" + solicitudBoletos.Solicitud.Destino + " --- " +
						solicitudBoletos.Solicitud.Origen, nil, errors.New("no se encuentra el boleto")
				}
				rutas.Id = bson.NewObjectId().Hex()
				rutas.Cantidad = 1
				rutas.BoletosId = boletos.Id
				rutas.Estado = true
				err = rutas.AgregarRutaConsumida(tx)
				if err != nil {
					tx.Rollback()
					return http.StatusInternalServerError, "Error al agregar la ruta", nil, err
				}
			}
		}
	}
	if solicitudBoletos.Solicitud.Acompañante {
		var datos TrasladoPacientesModel.Acompaniante
		datos.Id = bson.NewObjectId().Hex()
		datos.Cargo = solicitudBoletos.Acompaniante.Cargo
		datos.Justificacion = solicitudBoletos.Acompaniante.Justificacion
		datos.Matricula = solicitudBoletos.Acompaniante.Matricula
		datos.Motivo = solicitudBoletos.Acompaniante.Motivo
		datos.NombreAutoriza = solicitudBoletos.Acompaniante.NombreAutoriza
		datos.TarjetaInapan = solicitudBoletos.Acompaniante.TarjetaInapan
		datos.SolicitudTrasladoFolio = solicitudBoletos.Solicitud.Folio
		err = datos.AgregarAutorizacionAcomp(tx)
		if err != nil {
			tx.Rollback()
			return http.StatusInternalServerError, "Error al agregar los datos", nil, err
		}
	}
	var firmsNew TrasladoPacientesModel.FirmaTraslado
	for _, resultado := range firms {
		firmsNew.Id = bson.NewObjectId().Hex()
		firmsNew.FirmaSolicitudId = resultado
		firmsNew.SolicitudTrasladoFolio = solicitudBoletos.Solicitud.Folio
		firmsNew.Estado = true
		err = firmsNew.AgregarFirmaTraslado(tx)
		if err != nil {
			tx.Rollback()
			return http.StatusInternalServerError, "Error al asociar la firma con la solicitud", nil, err
		}
	}
	if solicitudBoletos.AutorizacionesEspeciales.Justificacion != "" {
		var autorizacionEsp TrasladoPacientesModel.AutorizacionesEspeciales
		autorizacionEsp.Id = bson.NewObjectId().Hex()
		autorizacionEsp.Justificacion = solicitudBoletos.AutorizacionesEspeciales.Justificacion
		autorizacionEsp.TipoAutorizacion = solicitudBoletos.AutorizacionesEspeciales.TipoAutorizacion
		autorizacionEsp.SolicitudTrasladoFolio = solicitudBoletos.Solicitud.Folio
		autorizacionEsp.Estado = true
		err = autorizacionEsp.AgregarAutorizacionEsp(tx)
		if err != nil {
			tx.Rollback()
			return http.StatusInternalServerError, "error al registrar la autorizacion especial", nil, err
		}
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, "Error al realizar el commit", nil, err
	}
	return http.StatusOK, "Solicitud realizada correctamente", solicitudBoletos.Solicitud, nil
}

func PostSolicitud(c *gin.Context) {
	var solicitudBoletos SolicitudtrasladoBoletos
	fechaCita := c.DefaultQuery("fecha", "")
	uiPrei := c.DefaultQuery("uiPrei", "")
	firm := c.DefaultQuery("firmsId", "")
	folioSeparado := c.DefaultQuery("folioSeparado", "")
	err := c.BindJSON(&solicitudBoletos)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error al castear la estructura", err.Error(), c)
		return
	}
	/*****  parsear string a array  *****/
	firms := make([]string, 0)
	if firm != "" {
		firms = strings.Split(firm, ",")
	}
	if folioSeparado == "True" {
		boletos := make([]Boletos, 0)
		boletos2 := make([]Boletos, 0)
		solicitudBoletos.Solicitud.Version = 2
		for _, b := range solicitudBoletos.Boletos {
			b.Regreso = false
			b.Ida = true
			b.Total = 1
			boletos = append(boletos, b)
		}
		solicitudBoletos.Boletos = boletos
		codigo, mns, data, err := AgregarSolicitud(solicitudBoletos, fechaCita, uiPrei, firms, folioSeparado)
		if err != nil {
			Estructuras.Responder(codigo, mns, err.Error(), c)
			return
		}
		for _, b2 := range solicitudBoletos.Boletos {
			b2.Ida = false
			b2.Regreso = true
			b2.Total = 1
			boletos2 = append(boletos2, b2)
		}
		solicitudBoletos.Boletos = boletos2
		codigo2, mns2, data2, err := AgregarSolicitud(solicitudBoletos, fechaCita, uiPrei, firms, folioSeparado)
		if err != nil {
			Estructuras.Responder(codigo2, mns2, err.Error(), c)
			return
		}
		fmt.Print("se generaran dos folios", solicitudBoletos, len(solicitudBoletos.Boletos))
		Estructuras.Responder(codigo2, mns2, gin.H{"folio1": data, "folio2": data2}, c)
		return
	}
	if folioSeparado == "" {
		solicitudBoletos.Solicitud.Version = 1
		fmt.Print("solo un folio")
		codigo, mns, data, err := AgregarSolicitud(solicitudBoletos, fechaCita, uiPrei, firms, folioSeparado)
		if err != nil {
			Estructuras.Responder(codigo, mns, err.Error(), c)
			return
		}
		Estructuras.Responder(codigo, mns, data, c)
		return
	}
	// solicitudBoletos.Solicitud.FechaExpedicion = time.Now()
	// t, _ := time.Parse(time.RFC3339, fechaCita)
	// solicitudBoletos.Solicitud.FechaCita = t
	// nuevoAnio := t.Year()
	// digitos := strconv.Itoa(nuevoAnio)
	// /*****  Funcion para saber que numero de folio consecutivo sigue  *****/
	// err, consecutivo := TrasladoPacientesModel.BuscarFolioConsecutivo(solicitudBoletos.Solicitud.UnidadMedId, nuevoAnio)
	// if err != nil {
	// 	Estructuras.Responder(http.StatusInternalServerError, "Error al consultar el folio consecutivo", err.Error(), c)
	// 	return
	// }
	// /**  Buscar que el folio FTP no se encuentre registrado y si se encuentra registrado que la solicitud este cancelada ***/
	// duplicado, _ := TrasladoPacientesModel.BuscarFolioFtp(solicitudBoletos.Solicitud.Ftp01, "true")
	// if duplicado {
	// 	Estructuras.Responder(http.StatusInternalServerError,
	// 		"Este folio FTP01: "+solicitudBoletos.Solicitud.Ftp01+" ya se encuentra registrado", solicitudBoletos.Solicitud, c)
	// 	return
	// }
	// /*****  Inicializar la transaccion de la BD  *****/
	// tx := Conexion.GetDB().Begin()
	// if tx.Error != nil {
	// 	Estructuras.Responder(http.StatusInternalServerError, "Error al iniciar la transaccion", tx.Error, c)
	// 	return
	// }
	// /*****  Agregar la solicitud de traslado  *****/
	// folio_sig := strconv.Itoa(consecutivo + 1)
	// solicitudBoletos.Solicitud.Folio = uiPrei + "-" + digitos[2:] + "-" + folio_sig
	// solicitudBoletos.Solicitud.Version = 1
	// solicitudBoletos.Solicitud.Estado = true
	// err = solicitudBoletos.Solicitud.AgregarSolicitud(tx)
	// if err != nil {
	// 	tx.Rollback()
	// 	Estructuras.Responder(http.StatusInternalServerError, "Error al insertar los datos", err.Error(), c)
	// 	return
	// }
	// /*****  Agregar los boletos   *****/
	// var boletos TrasladoPacientesModel.Boletos
	// var rutas TrasladoPacientesModel.RutasConsumidas
	// for _, registro := range solicitudBoletos.Boletos {
	// 	boletos.Id = bson.NewObjectId().Hex()
	// 	boletos.Ida = registro.Ida
	// 	boletos.Regreso = registro.Regreso
	// 	boletos.Total = registro.Total
	// 	boletos.Observaciones = ReturnObservation(registro.Observaciones)
	// 	boletos.Tipo = ReturnTicketType(registro.TipoBoleto)
	// 	boletos.TipoPasajeroId = registro.Tipo
	// 	boletos.SolicitudTrasladoFolio = solicitudBoletos.Solicitud.Folio
	// 	boletos.Estado = true
	// 	err = boletos.AgregarBoletosTraslado(tx)
	// 	if err != nil {
	// 		tx.Rollback()
	// 		Estructuras.Responder(http.StatusInternalServerError, "Error al agregar los boletos", err.Error(), c)
	// 		return
	// 	}
	// 	/*****  Validar y agregar la ruta consumida  *****/
	// 	if registro.TipoBoleto != 0 {
	// 		if registro.Ida {
	// 			_, res := TrasladoPacientesModel.SearchTicket(solicitudBoletos.Solicitud.UnidadMedId,
	// 				solicitudBoletos.Solicitud.Origen, solicitudBoletos.Solicitud.Destino,
	// 				ReturnTicketType(registro.TipoBoleto), solicitudBoletos.Solicitud.ContratoNumeroContrato)
	// 			if res != nil {
	// 				rutas.TrasladoPacientesId = res.Id
	// 			} else {
	// 				tx.Rollback()
	// 				Estructuras.Responder(http.StatusInternalServerError,
	// 					"No se tiene contratado el tipo de boleto ("+ReturnTicketType(registro.TipoBoleto)+
	// 						") a esta ruta -->"+solicitudBoletos.Solicitud.Origen+" --- "+solicitudBoletos.Solicitud.Destino,
	// 					errors.New("no se encuentra el boleto"), c)
	// 				return
	// 			}
	// 			rutas.Id = bson.NewObjectId().Hex()
	// 			rutas.Cantidad = 1
	// 			rutas.BoletosId = boletos.Id
	// 			rutas.Estado = true
	// 			err = rutas.AgregarRutaConsumida(tx)
	// 			if err != nil {
	// 				tx.Rollback()
	// 				Estructuras.Responder(http.StatusInternalServerError, "Error al agregar la ruta", err.Error(), c)
	// 				return
	// 			}
	// 		}
	// 		if registro.Regreso {
	// 			_, res := TrasladoPacientesModel.SearchTicket(solicitudBoletos.Solicitud.UnidadMedId,
	// 				solicitudBoletos.Solicitud.Destino, solicitudBoletos.Solicitud.Origen,
	// 				ReturnTicketType(registro.TipoBoleto), solicitudBoletos.Solicitud.ContratoNumeroContrato)
	// 			if res != nil {
	// 				rutas.TrasladoPacientesId = res.Id
	// 			} else {
	// 				tx.Rollback()
	// 				Estructuras.Responder(http.StatusInternalServerError,
	// 					"No se tiene contratado el tipo de boleto ("+ReturnTicketType(registro.TipoBoleto)+
	// 						") a esta ruta -->"+solicitudBoletos.Solicitud.Destino+" --- "+solicitudBoletos.Solicitud.Origen,
	// 					errors.New("no se encuentra el boleto"), c)
	// 				return
	// 			}
	// 			rutas.Id = bson.NewObjectId().Hex()
	// 			rutas.Cantidad = 1
	// 			rutas.BoletosId = boletos.Id
	// 			rutas.Estado = true
	// 			err = rutas.AgregarRutaConsumida(tx)
	// 			if err != nil {
	// 				tx.Rollback()
	// 				Estructuras.Responder(http.StatusInternalServerError, "Error al agregar la ruta", err.Error(), c)
	// 				return
	// 			}
	// 		}
	// 	}
	// }
	// if solicitudBoletos.Solicitud.Acompañante {
	// 	var datos TrasladoPacientesModel.Acompaniante
	// 	datos.Id = bson.NewObjectId().Hex()
	// 	datos.Cargo = solicitudBoletos.Acompaniante.Cargo
	// 	datos.Justificacion = solicitudBoletos.Acompaniante.Justificacion
	// 	datos.Matricula = solicitudBoletos.Acompaniante.Matricula
	// 	datos.Motivo = solicitudBoletos.Acompaniante.Motivo
	// 	datos.NombreAutoriza = solicitudBoletos.Acompaniante.NombreAutoriza
	// 	datos.TarjetaInapan = solicitudBoletos.Acompaniante.TarjetaInapan
	// 	datos.SolicitudTrasladoFolio = solicitudBoletos.Solicitud.Folio
	// 	err = datos.AgregarAutorizacionAcomp(tx)
	// 	if err != nil {
	// 		tx.Rollback()
	// 		Estructuras.Responder(http.StatusInternalServerError, "Error al agregar los datos", err.Error(), c)
	// 		return
	// 	}
	// }
	// var firmsNew TrasladoPacientesModel.FirmaTraslado
	// for _, resultado := range firms {
	// 	firmsNew.Id = bson.NewObjectId().Hex()
	// 	firmsNew.FirmaSolicitudId = resultado
	// 	firmsNew.SolicitudTrasladoFolio = solicitudBoletos.Solicitud.Folio
	// 	firmsNew.Estado = true
	// 	err = firmsNew.AgregarFirmaTraslado(tx)
	// 	if err != nil {
	// 		tx.Rollback()
	// 		Estructuras.Responder(http.StatusInternalServerError, "Error al asociar la firma con la solicitud",
	// 			err.Error(), c)
	// 		return
	// 	}
	// }
	// if solicitudBoletos.AutorizacionesEspeciales.Justificacion != "" {
	// 	var autorizacionEsp TrasladoPacientesModel.AutorizacionesEspeciales
	// 	autorizacionEsp.Id = bson.NewObjectId().Hex()
	// 	autorizacionEsp.Justificacion = solicitudBoletos.AutorizacionesEspeciales.Justificacion
	// 	autorizacionEsp.TipoAutorizacion = solicitudBoletos.AutorizacionesEspeciales.TipoAutorizacion
	// 	autorizacionEsp.SolicitudTrasladoFolio = solicitudBoletos.Solicitud.Folio
	// 	autorizacionEsp.Estado = true
	// 	err = autorizacionEsp.AgregarAutorizacionEsp(tx)
	// 	if err != nil {
	// 		tx.Rollback()
	// 		Estructuras.Responder(http.StatusInternalServerError, "error al registrar la autorizacion especial", err.Error(), c)
	// 		return
	// 	}
	// }
	// err = tx.Commit().Error
	// if err != nil {
	// 	tx.Rollback()
	// 	Estructuras.Responder(http.StatusInternalServerError, "Error al realizar el commit", err.Error(), c)
	// 	return
	// }
	// Estructuras.Responder(http.StatusOK, "Solicitud realizada correctamente", solicitudBoletos.Solicitud, c)
}

/*****  Controlador para consultar las solicitudes realizadas por unidad medica  *****/
func GetRequestsMade(c *gin.Context) {
	unitMedical := c.DefaultQuery("unitMedical", "")
	numberPageTmp := c.DefaultQuery("numberPage", "1")
	pageSizeTmp := c.DefaultQuery("pageSize", "25")
	filter := c.DefaultQuery("filter", "")
	state := c.DefaultQuery("state", "")
	filterCompanion := c.DefaultQuery("companion", "")
	filterOrigin := c.DefaultQuery("origin", "")
	filterDestination := c.DefaultQuery("destination", "")
	numberContract := c.DefaultQuery("numberContract", "")
	dateStart := c.DefaultQuery("dateStart", "")
	dateEnd := c.DefaultQuery("dateEnd", "")
	startDate := ""
	endDate := ""
	if dateStart != "" && dateEnd != "" {
		temporalDate, _ := time.Parse(time.RFC3339, dateStart)
		startDate = temporalDate.Truncate(24 * time.Hour).Format("2006-01-02 15:04:05")

		temporalDate2, _ := time.Parse(time.RFC3339, dateEnd)
		endDate = temporalDate2.Add(24 * time.Hour).Truncate(24 * time.Hour).Format("2006-01-02 15:04:05")

	}
	unitMedicalId, err := strconv.Atoi(unitMedical)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Numero de Unidad Medica incorrecto", err.Error(), c)
		return
	}
	numberPage, err := strconv.Atoi(numberPageTmp)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Numero de pagina incorrecto", err.Error(), c)
		return
	}
	pageSize, err := strconv.Atoi(pageSizeTmp)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Numero de registros incorrecto", err.Error(), c)
		return
	}
	var contractsNumbers []string
	if numberContract != "" {
		contractsNumbers = strings.Split(numberContract, ",")
	}
	err, result, total := TrasladoPacientesModel.ConsultarSolicitudes(unitMedicalId, numberPage, pageSize, filter, state, filterCompanion, filterOrigin, filterDestination, contractsNumbers, startDate, endDate)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", gin.H{"total": total, "registros": result}, c)
}

/*****  Controlador para consultar una solicitud por folio  *****/
func GetFolioSolicitud(ctx *gin.Context) {
	folio := ctx.DefaultQuery("folio", "")
	var folioCancelar []string
	resul, solicitud := TrasladoPacientesModel.BuscarFolioSolicitud(folio)
	if !resul {
		Estructuras.Responder(http.StatusInternalServerError,
			"No se encontro una solicitud con este numero de folio: "+folio, nil, ctx)
		return
	}
	if !solicitud.Estado {
		Estructuras.Responder(http.StatusInternalServerError,
			"Esta solicitud con numero de folio: "+folio+" ya se encuentra cancelado", nil, ctx)
		return
	}
	if solicitud.Version == 2 {
		_, busquedaFolio := TrasladoPacientesModel.BuscarFolioSolicitudes(solicitud.Ftp01)
		for _, item := range *busquedaFolio {
			folioCancelar = append(folioCancelar, item.Folio)
		}
		Estructuras.Responder(http.StatusOK, "Solicitud de traslado encontrado", gin.H{"solicitud": solicitud, "folios": folioCancelar}, ctx)
		return
	}
	Estructuras.Responder(http.StatusOK, "Solicitud de traslado encontrado", gin.H{"solicitud": solicitud}, ctx)
}

/*****  Controlador para cancelar una solicitud  *****/
func CancelFolioSolicitud(ctx *gin.Context) {
	folio := ctx.DefaultQuery("folio", "")
	boletoSeparados:= ctx.DefaultQuery("boletoSeparados","")
	var cancelacion TrasladoPacientesModel.SolicitudesCanceladas
	var solicitud TrasladoPacientesModel.SolicitudTraslado
	err := ctx.BindJSON(&cancelacion)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error al castear la estructura", err.Error(), ctx)
		return
	}
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al iniciar la transaccion", tx.Error, ctx)
		return
	}
	if boletoSeparados != "" {
		 foliosTmp := strings.Split(folio, ",")
		 for _, item := range foliosTmp{
			err = solicitud.CancelarSolicitud(item, tx)
			if err != nil {
				tx.Rollback()
				Estructuras.Responder(http.StatusInternalServerError,
					"Error al cancelar la solicitud con este numero de folio: "+item, err.Error(), ctx)
				return
			}
			cancelacion.Id = bson.NewObjectId().Hex()
			cancelacion.SolicitudTrasladoFolio = item
			cancelacion.FechaCancelacion = time.Now()
			cancelacion.Estado = true
			err = cancelacion.AgregarBitacoraCancelacion(tx)
			if err != nil {
				tx.Rollback()
				Estructuras.Responder(http.StatusInternalServerError, "Error al insertar los datos en la bitacora",
					err.Error(), ctx)
				return
			}
		 }
	}else{
		err = solicitud.CancelarSolicitud(folio, tx)
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError,
				"Error al cancelar la solicitud con este numero de folio: "+folio, err.Error(), ctx)
			return
		}
		cancelacion.Id = bson.NewObjectId().Hex()
		cancelacion.FechaCancelacion = time.Now()
		cancelacion.Estado = true
		err = cancelacion.AgregarBitacoraCancelacion(tx)
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "Error al insertar los datos en la bitacora",
				err.Error(), ctx)
			return
		}
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al realizar el commit", err.Error(), ctx)
		return
	}
	Estructuras.Responder(http.StatusOK, "Solicitud de traslado cancelada", solicitud, ctx)
}
func GetOriginFilter(ctx *gin.Context) {
	unitId := ctx.DefaultQuery("unitId", "")
	state := ctx.DefaultQuery("state", "")
	destination := ctx.DefaultQuery("destination", "")
	companion := ctx.DefaultQuery("companion", "")
	numberContract := ctx.DefaultQuery("numberContract", "")
	dateStart := ctx.DefaultQuery("startDate", "")
	dateEnd := ctx.DefaultQuery("endDate", "")
	startDate := ""
	endDate := ""
	if dateStart != "" && dateEnd != "" {
		temporalDate, _ := time.Parse(time.RFC3339, dateStart)
		startDate = temporalDate.Truncate(24 * time.Hour).Format("2006-01-02 15:04:05")
		temporalDate2, _ := time.Parse(time.RFC3339, dateEnd)
		endDate = temporalDate2.Add(24 * time.Hour).Truncate(24 * time.Hour).Format("2006-01-02 15:04:05")
	}
	stateRequest, err := strconv.ParseBool(state)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error estado no valido", err.Error(), ctx)
		return
	}
	idUnit, err := strconv.Atoi(unitId)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Numero de Unidad Medica incorrecto", err.Error(), ctx)
		return
	}
	var contractsNumbers []string
	if numberContract != "" {
		contractsNumbers = strings.Split(numberContract, ",")
	}
	results, err := TrasladoPacientesModel.SearchOriginFilter(idUnit, stateRequest, destination, companion,
		contractsNumbers, startDate, endDate)
	Estructuras.Responder(http.StatusOK, "Datos consultados", results, ctx)
	return
}
func GetDestinationFilter(ctx *gin.Context) {
	unitId := ctx.DefaultQuery("unitId", "")
	state := ctx.DefaultQuery("state", "")
	origin := ctx.DefaultQuery("origin", "")
	companion := ctx.DefaultQuery("companion", "")
	numberContract := ctx.DefaultQuery("numberContract", "")
	dateStart := ctx.DefaultQuery("startDate", "")
	dateEnd := ctx.DefaultQuery("endDate", "")
	startDate := ""
	endDate := ""
	if dateStart != "" && dateEnd != "" {
		temporalDate, _ := time.Parse(time.RFC3339, dateStart)
		startDate = temporalDate.Truncate(24 * time.Hour).Format("2006-01-02 15:04:05")
		temporalDate2, _ := time.Parse(time.RFC3339, dateEnd)
		endDate = temporalDate2.Add(24 * time.Hour).Truncate(24 * time.Hour).Format("2006-01-02 15:04:05")
	}
	stateRequest, err := strconv.ParseBool(state)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error estado no valido", err.Error(), ctx)
		return
	}
	idUnit, err := strconv.Atoi(unitId)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Numero de Unidad Medica incorrecto", err.Error(), ctx)
		return
	}
	var contractsNumbers []string
	if numberContract != "" {
		contractsNumbers = strings.Split(numberContract, ",")
	}
	results, err := TrasladoPacientesModel.SearchDestinationFilter(idUnit, stateRequest, origin, companion,
		contractsNumbers, startDate, endDate)
	Estructuras.Responder(http.StatusOK, "Datos consultados", results, ctx)
	return
}
