package TrasladoPacientesModel

import (
	"appbiomedica/Modelos/ContratoModel/ContratoModelo"
	"appbiomedica/Modelos/ModeloUnidadM"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/jung-kurt/gofpdf"
	"github.com/skip2/go-qrcode"
)

type boletoAutorizado struct {
	id, tipo, ida, regreso, boleto, total, observaciones string
}
type Encabezado struct {
	Ancho  float64
	Nombre string
}

func RegresarMjs(tipo bool) string {
	if tipo {
		return "Si"
	} else {
		return "No"
	}
}

type BoletosPrecio struct {
	Ruta       string  `json:"ruta"`
	TipoBoleto string  `json:"tipoBoleto"`
	Cantidad   int     `json:"cantidad"`
	Precio     float64 `json:"precio"`
}
type DataQr struct {
	Unidad  string          `json:"unidad"`
	Folio   string          `json:"folio"`
	Boletos []BoletosPrecio `json:"boletos"`
}

var header = []Encabezado{{Ancho: 26, Nombre: "TIPO"}, {Ancho: 15, Nombre: "IDA"}, {Ancho: 25, Nombre: "REGRESO"},
	{Ancho: 25, Nombre: "BOLETO"}, {Ancho: 20, Nombre: "TOTAL"}, {Ancho: 0, Nombre: "NOTAS"}}

/*****  Funcion para generar un PDF de una solicitud de traslado de pacientes (solo la plantilla)  *****/
func CreatePdfNoDate(unidad_id int, voboId, authorizeId string) (string, error) {
	err, unidad_medica := ModeloUnidadM.ConsultaUnaUnidadMed(unidad_id)
	if err != nil {
		return "", err
	}
	err, firmVoBo := ReturnFirm(voboId)
	if err != nil {
		return "", err
	}
	err, firmAuthorize := ReturnFirm(authorizeId)
	if err != nil {
		return "", err
	}
	pdf := gofpdf.New("P", "mm", "Letter", "")
	tr := pdf.UnicodeTranslatorFromDescriptor("")
	/*****  -*--------------------------  *****/
	pdf.AddPage()
	/*****  encabezado  *****/
	pdf.Image("Modelos/ServiciosSubrogados/TrasladoPacientesModel/recursos/LOGOFINALIMSS.png", 15, 21, 25, 30, false,
		"PNG", 0, "")
	pdf.SetFont("Arial", "B", 15)
	pdf.MoveTo(pdf.GetX(), pdf.GetY()+2)
	pdf.MultiCell(0, 4, "SOLICITUD DE SERVICIO DE TRASLADO DE PACIENTES", "", "C", false)
	pdf.Ln(4)
	pdf.SetFont("Arial", "", 12)
	pdf.WriteAligned(0, 0, "INSTITUTO MEXICANO DEL SEGURO SOCIAL", "C")
	pdf.Ln(4)
	pdf.WriteAligned(0, 0, "OOAD ESTATAL EN OAXACA", "C")
	/*****  datos del contrato  *****/
	/*****  numero de contrato  *****/
	pdf.MoveTo(45, 29)
	pdf.SetFont("Arial", "", 11)
	pdf.MultiCell(30.5, 4, tr("CONTRATO N°:"), "", "L", false)
	/*****  REPRESENTANTE LEGAL  *****/
	pdf.SetFont("Arial", "", 11)
	pdf.MoveTo(45, pdf.GetY()+1)
	pdf.MultiCell(51, 4, tr("REPRESENTANTE LEGAL:"), "", "L", false)
	/*****  numero de proveedor  *****/
	pdf.SetFont("Arial", "", 11)
	pdf.MoveTo(45, pdf.GetY()+1)
	pdf.MultiCell(40, 4, tr("N° DE PROVEEDOR:"), "", "L", false)
	/*****  proveedor  *****/
	pdf.SetFont("Arial", "", 11)
	pdf.MoveTo(45, pdf.GetY()+1)
	pdf.MultiCell(28, 4, tr("PROVEEDOR:"), "", "L", false)
	/*****  numero  *****/
	pdf.MoveTo(pdf.GetX(), 60)
	pdf.SetFont("Arial", "B", 11.5)
	pdf.WriteAligned(0, 0, tr("ANEXO  N° 3.- \"TRASLADO DE PACIENTES QUE GENERA GASTOS DE PASAJES Y VIATICOS\" "), "C")
	pdf.Ln(7)
	/*****  ------------------------------------  *****/
	/*****  unidad medica  *****/
	pdf.SetFont("Arial", "", 11)
	pdf.MultiCell(35, 3, tr("UNIDAD MEDICA:"), "", "L", false)
	/*****  numero de unidad *****/
	pdf.SetFont("Arial", "B", 11)
	pdf.MoveTo(pdf.GetX()+34, pdf.GetY()-3)
	pdf.MultiCell(0, 3, tr(strings.ToUpper(unidad_medica.DenominacionUni)+", "+
		""+strings.ToUpper(unidad_medica.Localidad)), "", "L", false)
	pdf.Ln(3)
	/*****  FOLIO DE SOLICITU  *****/
	pdf.SetFont("Arial", "", 11)
	pdf.MultiCell(79, 3, tr("TRASLADO DE PACIENTE CON FOLIO N°:"), "", "L", false)
	pdf.Ln(3)

	/*****  NOMBRE DEL PACIENTE  *****/
	pdf.MultiCell(50, 3, tr("NOMBRE DEL PACIENTE:"), "", "L", false)
	pdf.Ln(1.5)
	/*****  NUMERO DE SEGURIDAD  *****/
	pdf.MultiCell(50, 3, tr("NUMERO DE AFILIACIÓN:"), "", "L", false)
	/*****  AGREGADO MEDICO  *****/
	pdf.MoveTo(pdf.GetX()+78, pdf.GetY()-3)
	pdf.MultiCell(42.5, 3, tr("AGREGADO MEDICO:"), "", "L", false)
	pdf.Ln(3)
	/*****  FECHA DE LA CITA  *****/
	pdf.MultiCell(55, 3, tr("FECHA DE LA CITA MEDICA:"), "", "L", false)
	pdf.Ln(3)
	/*****  LUGAR DE ORIGEN *****/
	pdf.MultiCell(34, 3, tr("BOLETO DE IDA:"), "", "L", false)
	pdf.Ln(1.5)
	/*****  LUGAR DE DESTINO *****/
	pdf.MultiCell(46, 3, tr("BOLETO DE REGRESO:"), "", "L", false)
	pdf.Ln(3)
	/*****  AUTORIZAR ACOMPAÑANTE *****/
	pdf.MultiCell(60, 3, tr("SE AUTORIZA ACOMPAÑANTE:"), "", "L", false)
	/*****  BOLETOS AUTORIZADOS *****/
	pdf.MoveTo(pdf.GetX(), 117)
	pdf.SetFont("Arial", "BU", 12)
	pdf.MultiCell(0, 4, tr("BOLETOS AUTORIZADOS"), "", "C", false)
	/*****  tablaaaa  *****/
	pdf.SetFont("Arial", "B", 11)
	pdf.MoveTo(pdf.GetX(), 168.5)
	/*****  VIGENCIA FTP *****/
	pdf.MultiCell(0, 5, tr("VIGENTE EN SUS DERECHOS CONFORME A LA CERTIFICACION DE VIGENCIA DE DERECHOS EN LA FTP-01"+
		":"), "", "L", false)
	/*****  documentos de identidficacion  *****/
	pdf.Ln(2.5)
	pdf.SetFont("Arial", "", 11)
	/*****  fecha de expedicion *****/
	pdf.MultiCell(48, 4, tr("FECHA DE EXPEDICION:"), "", "L", false)
	pdf.Ln(3)
	/*****  Leyenda documentos de identificacion  *****/
	pdf.SetFont("Arial", "B", 10.5)
	pdf.MultiCell(69, 4, tr("*DOCUMENTOS DE IDENTIFICACION:"), "", "L", false)
	pdf.MoveTo(pdf.GetX()+68.5, pdf.GetY()-4)
	pdf.SetFont("Arial", "I", 10.5)
	pdf.MultiCell(0, 4, tr("CREDENCIAL DE ELECTOR, LICENCIA DE MANEJO, "+"TARJETA DE "), "", "J", false)
	pdf.MultiCell(0, 4, tr("CONTROL DE CITAS IMSS CON FOTOGRAFIA O TARJETA DE IDENTIFICACION COMO PENSIONADO IMSS."), "", "L", false)
	/*****  Recepcion de boletos *****/
	pdf.MoveTo(pdf.GetX(), 208)
	pdf.SetFont("Arial", "", 11)
	pdf.MultiCell(61, 4, tr("RECEPCION DE BOLETO (S):"), "", "L", false)
	/*****  nombre del paciente *****/
	pdf.MoveTo(pdf.GetX()+61, pdf.GetY())
	pdf.CellFormat(0, 6, "NOMBRE Y FIRMA", "T", 0, "C", false, 0, "")
	pdf.Ln(15)
	pdf.MultiCell(95, 4, tr("VO. BO."), "", "C", false)
	pdf.MoveTo(pdf.GetX()+95, pdf.GetY()-4)
	pdf.MultiCell(95, 4, "AUTORIZA", "", "C", false)
	pdf.Ln(3)
	ultimo := pdf.GetY()
	pdf.MultiCell(90, 5, tr(strings.ToUpper(firmVoBo.Nombre)), "", "C", false)
	pdf.MoveTo(pdf.GetX()+100, ultimo)
	pdf.MultiCell(90, 5, tr(strings.ToUpper(firmAuthorize.Nombre)), "", "C", false)
	pdf.MoveTo(pdf.GetX(), 249)
	cargo2 := pdf.GetY()
	pdf.MultiCell(90, 5, tr(strings.ToUpper(firmVoBo.Cargo)), "T", "C", false)
	pdf.MoveTo(pdf.GetX()+100, cargo2)
	pdf.MultiCell(90, 5, tr(strings.ToUpper(firmAuthorize.Cargo)), "T", "C", false)
	err = pdf.OutputFileAndClose("PlantillaSolicitudTraslado.pdf")
	if err != nil {
		return "", err

	}
	data, err := ioutil.ReadFile("PlantillaSolicitudTraslado.pdf")
	if err != nil {
		return "", err
	}
	os.Remove("PlantillaSolicitudTraslado.pdf")
	return base64.StdEncoding.EncodeToString(data), nil
}

/*****  Funcion para generar un PDF de una solicitud de traslado de pacientes (solo datos)  *****/
func CrearPdfConDatos(folio, reimpresion string,pdf *gofpdf.Fpdf) (*gofpdf.Fpdf, error) {
	PreciosQrBoletos := make([]BoletosPrecio, 0)
	bole := make([]boletoAutorizado, 0)
	bolean, res := BuscarBoletosTraslado(folio)
	if !bolean {
		/*****  mandar erro de que no se encontro este folio registrado  *****/
		return nil, errors.New("No se encontraron datos con este numero de folio" + folio)
	}
	resul, solicitud := BuscarFolioSolicitud(folio)
	if !resul {
		/*****  mandar erro de que no se encontro este folio registrado  *****/
		return nil, errors.New("No se encontraron datos con este numero de folio" + folio)
	}
	contrato := ContratoModelo.RegresaContrato(solicitud.ContratoNumeroContrato)
	resp, _ := BuscarDatosAcompaniante(folio)
	if !resp {
		/*****  mandar erro de que no se encontro este folio registrado  *****/
	}
	resultadoBusqueda, solicitudTraslado := BuscarFolioSolicitud(folio)
	if !resultadoBusqueda {
		return nil, errors.New("No se encontro ningun registro con este numero de folio: " + folio)
	}
	err, unidad_medica := ModeloUnidadM.ConsultaUnaUnidadMed(solicitudTraslado.UnidadMedId)
	if err != nil {
		return nil, err
	}
	preciosQr := BuscarBoletosTrasladoPrecio(folio)
	var tmpPrecios BoletosPrecio
	for _, d := range *preciosQr {
		tmpPrecios.Ruta = d.TrasladoPacientes.Origen + d.TrasladoPacientes.Destino
		tmpPrecios.TipoBoleto = d.Boletos.Tipo
		tmpPrecios.Cantidad = d.Cantidad
		tmpPrecios.Precio = d.TrasladoPacientes.PrecioOfertado
		PreciosQrBoletos = append(PreciosQrBoletos, tmpPrecios)
	}
	var tmp boletoAutorizado
	var ida bool
	var regreso bool
	for _, c := range *res {
		tmp.tipo = c.TipoPasajero.Tipo
		tmp.ida = RegresarMjs(c.Ida)
		tmp.regreso = RegresarMjs(c.Regreso)
		tmp.boleto = c.Tipo
		tmp.total = strconv.Itoa(c.Total)
		tmp.observaciones = c.Observaciones
		bole = append(bole, tmp)
		if c.Ida {
			ida = true
		}
		if c.Regreso {
			regreso = true
		}
	}
	/*****  codigo QR  *****/
	texto, _ := json.Marshal(DataQr{Unidad: unidad_medica.DenominacionUni, Folio: folio, Boletos: PreciosQrBoletos})
	err = qrcode.WriteFile(string(texto), qrcode.Medium, -5, "Modelos/ServiciosSubrogados/TrasladoPacientesModel/recursos/qr2.png")
	if err != nil {
		return nil, err
	}
	pdf.SetAutoPageBreak(false, 1.5)
	tr := pdf.UnicodeTranslatorFromDescriptor("")
	/*****  Variables para la tabla  *****/

	/*****  -*--------------------------  *****/
	pdf.AddPage()
	/*****  encabezado  *****/
	pdf.MoveTo(pdf.GetX(), pdf.GetY()+2)
	pdf.Ln(4)
	pdf.Ln(4)
	/*****  datos del contrato  *****/
	/*****  numero de contrato  *****/
	pdf.MoveTo(45, 29)
	/*****  numero  *****/
	pdf.SetFont("Arial", "B", 11)
	pdf.MoveTo(75, pdf.GetY())
	pdf.MultiCell(0, 4, tr(solicitud.ContratoNumeroContrato), "", "L", false)
	/*****  Representante legal  *****/
	pdf.MoveTo(45, pdf.GetY()+1)
	/*****  ------ Nombre  *****/
	pdf.SetFont("Arial", "B", 11)
	pdf.MoveTo(95, pdf.GetY())
	pdf.MultiCell(0, 4, tr(solicitudTraslado.RepresentanteLegal), "", "L", false)
	/*****  Numero de proveedor  *****/
	pdf.MoveTo(45, pdf.GetY()+1)
	/*****  ---------- Numero  *****/
	pdf.SetFont("Arial", "B", 11)
	pdf.MoveTo(84, pdf.GetY())
	pdf.MultiCell(0, 4, tr(contrato.ProveedorNProvImss), "", "L", false)
	/*****  Proveedor  *****/
	pdf.MoveTo(45, pdf.GetY()+1)
	/*****  -------- Nombre   *****/
	pdf.SetFont("Arial", "B", 9)
	pdf.MoveTo(72, pdf.GetY())
	pdf.MultiCell(0, 4, tr(contrato.Proveedor.NombreEmpresa), "", "L", false)

	pdf.MoveTo(pdf.GetX(), 60)
	pdf.SetFont("Arial", "B", 11.5)
	pdf.WriteAligned(0, 0, tr(""), "C")
	pdf.Ln(7)
	/*****  ------------------------------------  *****/
	/*****  unidad medica  *****/
	pdf.SetFont("Arial", "", 11)
	pdf.MultiCell(35, 3, tr(""), "", "L", false)
	pdf.Ln(3)
	/*****  FOLIO *****/
	pdf.SetFont("Arial", "BU", 13)
	pdf.MoveTo(pdf.GetX()+69, pdf.GetY()-3)
	pdf.MultiCell(0, 3, tr(solicitudTraslado.Folio), "", "L", false)
	if solicitudTraslado.Version ==2  {
		pdf.MoveTo(pdf.GetX()+111, pdf.GetY()-3)
		pdf.SetFont("Arial", "", 9)
		pdf.MultiCell(0, 3, tr("(Boletos separados)"), "", "L", false)
	}
	pdf.Ln(3)
	/*****  NOMBRE DEL PACIENTE  *****/
	/*****  NOMBRE *****/
	pdf.SetFont("Arial", "B", 11)
	pdf.MoveTo(pdf.GetX()+49, pdf.GetY())
	pdf.MultiCell(0, 3, tr(solicitud.NombrePaciente), "", "L", false)
	pdf.Ln(1.5)
	/*****  NUMERO *****/
	pdf.SetFont("Arial", "B", 11)
	pdf.MoveTo(pdf.GetX()+49, pdf.GetY())
	pdf.MultiCell(38, 3, tr(solicitud.Nss), "", "L", false)
	/*****  AGREGADO MEDICO  *****/
	/*****  AGREGADO *****/
	pdf.SetFont("Arial", "B", 11)
	pdf.MoveTo(pdf.GetX()+119, pdf.GetY()-3)
	pdf.MultiCell(0, 3, tr(solicitud.AgregadoMedico), "", "L", false)
	pdf.Ln(3)
	/*****  FECHA *****/
	pdf.SetFont("Arial", "B", 11)
	pdf.MoveTo(pdf.GetX()+54, pdf.GetY())
	pdf.MultiCell(0, 3, tr(solicitudTraslado.FechaCita.Format("02/01/2006")+"    "+
		solicitudTraslado.FechaCita.Format("15:04")+" hrs."), "", "L", false)
	pdf.Ln(3)

	/*****  LUGAR *****/
	pdf.SetFont("Arial", "B", 11)
	pdf.MoveTo(pdf.GetX()+34, pdf.GetY())
	if ida {
		pdf.MultiCell(0, 3, tr(solicitud.Origen+" - "+solicitud.Destino), "", "L", false)
	} else {
		pdf.MultiCell(0, 3, "NO APLICA", "", "L", false)
	}
	pdf.Ln(1.5)
	/*****  LUGAR *****/
	pdf.SetFont("Arial", "B", 11)
	pdf.MoveTo(pdf.GetX()+46, pdf.GetY())
	fmt.Println("-*-*-*-*-*-*-*-*-", solicitudTraslado)

	if regreso {
		if ida || solicitudTraslado.Version == 2 {
			pdf.MultiCell(0, 3, tr(solicitud.Destino+" - "+solicitud.Origen), "", "L", false)
		} else {
			pdf.MultiCell(0, 3, tr(solicitud.Origen+" - "+solicitud.Destino), "", "L", false)
		}
	} else {
		pdf.MultiCell(0, 3, "NO APLICA", "", "L", false)
	}
	pdf.Ln(3)
	/*****  LUGAR *****/
	pdf.SetFont("Arial", "BU", 12)
	pdf.MoveTo(pdf.GetX()+59, pdf.GetY())
	pdf.MultiCell(0, 3, tr(RegresarMjs(solicitud.Acompañante)), "", "L", false)
	pdf.Image(ImageFile("Modelos/ServiciosSubrogados/TrasladoPacientesModel/recursos/qr2.png"), 152, 64, 56, 0, false,
		"", 0, "")
	/*****  BOLETOS AUTORIZADOS *****/
	pdf.MoveTo(pdf.GetX(), 117)
	pdf.SetFont("Arial", "BU", 12)
	pdf.MultiCell(0, 4, tr(""), "", "C", false)
	/*****  tablaaaa  *****/
	pdf.Ln(4)
	pdf.SetX(pdf.GetX())
	pdf.SetFont("Arial", "B", 11)
	for _, str := range header {
		pdf.CellFormat(str.Ancho, 4.5, str.Nombre, "1", 0, "CM", false, 0, "")
	}
	pdf.SetFont("Arial", "", 11)
	pdf.Ln(-1)
	var observacionTmp string
	for _, c := range bole {
		pdf.SetX(pdf.GetX())
		pdf.MultiCell(26, 5, tr(c.tipo), "1", "L", false)
		pdf.MoveTo(pdf.GetX()+26, pdf.GetY()-10)
		pdf.MultiCell(15, 10, tr(c.ida), "1", "C", false)
		pdf.MoveTo(pdf.GetX()+41, pdf.GetY()-10)
		pdf.MultiCell(25, 10, tr(c.regreso), "1", "C", false)
		pdf.MoveTo(pdf.GetX()+66, pdf.GetY()-10)
		pdf.MultiCell(25, 10, tr(c.boleto), "1", "C", false)
		pdf.MoveTo(pdf.GetX()+91, pdf.GetY()-10)
		pdf.MultiCell(20, 10, tr(c.total), "1", "C", false)
		pdf.MoveTo(pdf.GetX()+111, pdf.GetY()-10)
		pdf.MultiCell(0, 10, tr(c.observaciones), "1", "C", false)
		observacionTmp = observacionTmp + "  " + c.observaciones
	}
	/*****  Observaciones *****/
	pdf.SetX(10.00125)
	pdf.MoveTo(pdf.GetX(), 162)
	/*****  -------- Observaciones  *****/
	pdf.SetFont("Arial", "B", 11)
	if resp {
		pdf.MultiCell(0, 3.5, tr("OBSERVACIONES: ")+tr(observacionTmp),
			"", "J", false)
	} else {
		pdf.MultiCell(0, 3.5, tr("SIN OBSERVACIONES"), "", "J", false)
	}
	pdf.MoveTo(pdf.GetX(), 168.5)

	/*****  VIGENCIA FTP *****/
	pdf.MultiCell(0, 5, tr(""), "", "L", false)
	pdf.MultiCell(0, 5, tr(""), "", "L", false)
	/*****  FTP *****/
	pdf.SetFont("Arial", "BU", 12)
	pdf.MoveTo(pdf.GetX()+16, pdf.GetY()-5)
	pdf.MultiCell(0, 5, tr(solicitud.Ftp01), "", "L", false)
	/*****  documentos de identidficacion  *****/
	pdf.Ln(2.5)
	/*****  LUGAR *****/
	pdf.SetFont("Arial", "B", 11)
	pdf.MoveTo(pdf.GetX()+47, pdf.GetY())
	pdf.MultiCell(0, 4, solicitudTraslado.FechaExpedicion.Format("02/01/2006")+"    "+
		solicitudTraslado.FechaExpedicion.Format("15:04:05")+" hrs.", "", "L", false)
	if reimpresion != "" {
		nombre := solicitudTraslado.Usuario.Persona.Nombre + " " + solicitudTraslado.Usuario.Persona.ApellidoPat + " " + solicitudTraslado.Usuario.Persona.ApellidoMat
		hora := time.Now()
		pdf.SetXY(10.00125, 261)
		pdf.SetFont("Arial", "B", 9)
		pdf.MultiCell(150, 4, "Reimprimio: "+nombre+"  Fecha: "+hora.Format("02/01/2006")+"  "+hora.Format("15:04:05")+" hrs.", "", "L", false)
	}
	return pdf, nil
}

/*****  Funcion para generar un PDf de una solicitud de traslado de pacientes (completa)  *****/
func CrearPdfCompleto(folio, reimpresion,folioSeparado string,pdf *gofpdf.Fpdf) (*gofpdf.Fpdf, error) {
	PreciosQrBoletos := make([]BoletosPrecio, 0)
	bolean, res := BuscarBoletosTraslado(folio)
	if !bolean {
		return nil, errors.New("No se encontro ningun registro con este numero de folio: " + folio)
	}
	resultadoBusqueda, solicitudTraslado := BuscarFolioSolicitud(folio)
	if !resultadoBusqueda {
		return nil, errors.New("No se encontro ningun registro con este numero de folio: " + folio)
	}
	err, unidad_medica := ModeloUnidadM.ConsultaUnaUnidadMed(solicitudTraslado.UnidadMedId)
	if err != nil {
		return nil, err
	}
	err, firmaVisto := ConsultarFirmaTraslado(folio, 1)
	if err != nil {
		return nil, err
	}
	err, firmaAutoriza := ConsultarFirmaTraslado(folio, 2)
	if err != nil {
		return nil, err
	}
	preciosQr := BuscarBoletosTrasladoPrecio(folio)
	var tmpPrecios BoletosPrecio
	for _, d := range *preciosQr {
		tmpPrecios.Ruta = d.TrasladoPacientes.Origen + d.TrasladoPacientes.Destino
		tmpPrecios.TipoBoleto = d.Boletos.Tipo
		tmpPrecios.Cantidad = d.Cantidad
		tmpPrecios.Precio = d.TrasladoPacientes.PrecioOfertado
		PreciosQrBoletos = append(PreciosQrBoletos, tmpPrecios)
	}

	bole := make([]boletoAutorizado, 0)
	contrato := ContratoModelo.RegresaContrato(solicitudTraslado.ContratoNumeroContrato)
	if contrato == nil {
		return nil, errors.New("El numero del contrato no se encontro en la base de datos: " +
			solicitudTraslado.ContratoNumeroContrato)
	}
	var ida bool
	var regreso bool
	resultAcomp, _ := BuscarDatosAcompaniante(folio)
	var tmp boletoAutorizado
	for _, c := range *res {
		tmp.id = c.Id
		tmp.tipo = c.TipoPasajero.Tipo
		tmp.ida = RegresarMjs(c.Ida)
		tmp.regreso = RegresarMjs(c.Regreso)
		tmp.boleto = c.Tipo
		tmp.total = strconv.Itoa(c.Total)
		tmp.observaciones = c.Observaciones
		bole = append(bole, tmp)
		if c.Ida {
			ida = true
		}
		if c.Regreso {
			regreso = true
		}
	}
	/*****  codigo QR  *****/
	texto, _ := json.Marshal(DataQr{Unidad: unidad_medica.DenominacionUni, Folio: folio, Boletos: PreciosQrBoletos})
	err = qrcode.WriteFile(string(texto), qrcode.Medium, -5, "Modelos/ServiciosSubrogados/TrasladoPacientesModel/recursos/qr2.png")
	if err != nil {
		return nil, err
	}
	//pdf := gofpdf.New("P", "mm", "Letter", "")
	pdf.SetAutoPageBreak(false, 1.5)
	tr := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.AddPage()
	if !solicitudTraslado.Estado {
		x := pdf.GetX()
		y := pdf.GetY()
		pdf.SetFont("Arial", "B", 65)
		pdf.SetTextColor(255, 0, 0)
		pdf.SetXY(x, y+20)
		pdf.TransformBegin()
		pdf.TransformRotate(45, 170, 100)
		pdf.CellFormat(0, 72, "C A N C E L A D O", "", 0, "C", false, 0, "")
		pdf.TransformEnd()
		pdf.SetXY(x, y)
	}

	/*****  Encabezado  *****/
	pdf.Image("Modelos/ServiciosSubrogados/TrasladoPacientesModel/recursos/LOGOFINALIMSS.png", 15, 21, 25, 30, false,
		"PNG", 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.SetFont("Arial", "B", 15)
	pdf.MoveTo(pdf.GetX(), pdf.GetY()+2)
	pdf.MultiCell(0, 4, "SOLICITUD DE SERVICIO DE TRASLADO DE PACIENTES", "", "C", false)
	pdf.Ln(4)
	pdf.SetFont("Arial", "", 12)
	pdf.WriteAligned(0, 0, "INSTITUTO MEXICANO DEL SEGURO SOCIAL", "C")
	pdf.Ln(4)
	pdf.WriteAligned(0, 0, "OOAD ESTATAL EN OAXACA", "C")
	/*****----- DATOS DEL CONTRATO  ------*****/
	/*****  Contrato N°  *****/
	pdf.MoveTo(45, 29)
	pdf.SetFont("Arial", "", 11)
	pdf.MultiCell(30.5, 4, tr("CONTRATO N°:"), "", "L", false)
	/*****  ------- Numero  *****/
	pdf.SetFont("Arial", "B", 11)
	pdf.MoveTo(75, pdf.GetY()-4)
	pdf.MultiCell(0, 4, tr(solicitudTraslado.ContratoNumeroContrato), "", "L", false)
	/*****  Representante legal  *****/
	pdf.SetFont("Arial", "", 11)
	pdf.MoveTo(45, pdf.GetY()+1)
	pdf.MultiCell(51, 4, tr("REPRESENTANTE LEGAL:"), "", "L", false)
	/*****  ------ Nombre  *****/
	pdf.SetFont("Arial", "B", 11)
	pdf.MoveTo(95, pdf.GetY()-4)
	pdf.MultiCell(0, 4, tr(solicitudTraslado.RepresentanteLegal), "", "L", false)
	/*****  Numero de proveedor  *****/
	pdf.SetFont("Arial", "", 11)
	pdf.MoveTo(45, pdf.GetY()+1)
	pdf.MultiCell(40, 4, tr("N° DE PROVEEDOR:"), "", "L", false)
	/*****  ---------- Numero  *****/
	pdf.SetFont("Arial", "B", 11)
	pdf.MoveTo(84, pdf.GetY()-4)
	pdf.MultiCell(0, 4, tr(contrato.ProveedorNProvImss), "", "L", false)
	/*****  Proveedor  *****/
	pdf.SetFont("Arial", "", 11)
	pdf.MoveTo(45, pdf.GetY()+1)
	pdf.MultiCell(28, 4, tr("PROVEEDOR:"), "", "L", false)
	/*****  -------- Nombre   *****/
	pdf.SetFont("Arial", "B", 9)
	pdf.MoveTo(72, pdf.GetY()-4)
	pdf.MultiCell(0, 4, tr(contrato.Proveedor.NombreEmpresa), "", "L", false)
	/*****  Leyenda anexo 3  *****/
	pdf.MoveTo(pdf.GetX(), 60)
	pdf.SetFont("Arial", "B", 11.5)
	pdf.WriteAligned(0, 0, tr("ANEXO  N° 3.- \"TRASLADO DE PACIENTES QUE GENERA GASTOS DE PASAJES Y VIATICOS\" "), "C")
	pdf.Ln(7)
	/*****  -------------------Datos del traslado-----------------  *****/
	/*****  Unidad medica  *****/
	pdf.SetFont("Arial", "", 11)
	pdf.MultiCell(35, 3, tr("UNIDAD MEDICA:"), "", "L", false)
	/*****  Numero de unidad *****/
	pdf.SetFont("Arial", "B", 11)
	pdf.MoveTo(pdf.GetX()+34, pdf.GetY()-3)
	pdf.MultiCell(0, 3, tr(strings.ToUpper(unidad_medica.DenominacionUni)+", "+
		""+strings.ToUpper(unidad_medica.Localidad)), "", "L", false)
	pdf.Ln(3)
	/*****  Folio de Solicitud  *****/
	pdf.SetFont("Arial", "", 11)
	pdf.MultiCell(79, 3, tr("TRASLADO DE PACIENTE FOLIO N°:"), "", "L", false)
	/*****  ------- Numero  *****/
	pdf.SetFont("Arial", "BU", 13)
	pdf.MoveTo(pdf.GetX()+69, pdf.GetY()-3)
	pdf.MultiCell(0, 3, tr(solicitudTraslado.Folio), "", "L", false)
	if solicitudTraslado.Version ==2  {
		pdf.MoveTo(pdf.GetX()+111, pdf.GetY()-3)
		pdf.SetFont("Arial", "", 9)
		pdf.MultiCell(0, 3, tr("(Boletos separados)"), "", "L", false)
	}
	pdf.Ln(3)
	/*****  Nombre del paciente  *****/
	pdf.SetFont("Arial", "", 11)
	pdf.MultiCell(50, 3, tr("NOMBRE DEL PACIENTE:"), "", "L", false)
	/*****  ------- Nombre  *****/
	pdf.SetFont("Arial", "B", 11)
	pdf.MoveTo(pdf.GetX()+49, pdf.GetY()-3)
	pdf.MultiCell(0, 3, tr(solicitudTraslado.NombrePaciente), "", "L", false)
	pdf.Ln(1.5)
	/*****  Numero de seguridad social  *****/
	pdf.SetFont("Arial", "", 11)
	pdf.MultiCell(50, 3, tr("NUMERO DE AFILIACIÓN:"), "", "L", false)
	/*****  ------ Numero  *****/
	pdf.SetFont("Arial", "B", 11)
	pdf.MoveTo(pdf.GetX()+49, pdf.GetY()-3)
	pdf.MultiCell(38, 3, tr(solicitudTraslado.Nss), "", "L", false)
	/*****  Agregado medico  *****/
	pdf.SetFont("Arial", "", 11)
	pdf.MoveTo(pdf.GetX()+78, pdf.GetY()-3)
	pdf.MultiCell(42.5, 3, tr("AGREGADO MEDICO:"), "", "L", false)
	/*****  -------- Numero  *****/
	pdf.SetFont("Arial", "B", 11)
	pdf.MoveTo(pdf.GetX()+119, pdf.GetY()-3)
	pdf.MultiCell(0, 3, tr(solicitudTraslado.AgregadoMedico), "", "L", false)
	pdf.Ln(3)
	/*****  Fecha de la cita  *****/
	pdf.SetFont("Arial", "", 11)
	pdf.MultiCell(55, 3, tr("FECHA DE LA CITA MEDICA:"), "", "L", false)
	/*****  ------ Fecha  *****/
	pdf.SetFont("Arial", "B", 11)
	pdf.MoveTo(pdf.GetX()+54, pdf.GetY()-3)
	pdf.MultiCell(0, 3, tr(solicitudTraslado.FechaCita.Format("02/01/2006")+"    "+
		solicitudTraslado.FechaCita.Format("15:04")+" hrs."), "", "L", false)
	pdf.Ln(3)
	/*****  Lugar de origen*****/
	pdf.SetFont("Arial", "", 11)
	pdf.MultiCell(34, 3, tr("BOLETO DE IDA:"), "", "L", false)
	/*****  -------- Lugar  *****/
	pdf.SetFont("Arial", "B", 11)
	if ida {
		pdf.MoveTo(pdf.GetX()+34, pdf.GetY()-3)
		pdf.MultiCell(0, 3, tr(solicitudTraslado.Origen+" - "+solicitudTraslado.Destino), "", "L", false)
	} else {
		pdf.MoveTo(pdf.GetX()+34, pdf.GetY()-3)
		pdf.MultiCell(0, 3, tr("NO APLICA"), "", "L", false)
	}
	pdf.Ln(1.5)
	/*****  Lugar de destino *****/
	pdf.SetFont("Arial", "", 11)
	pdf.MultiCell(46, 3, tr("BOLETO DE REGRESO:"), "", "L", false)
	/*****  ------- Lugar  *****/
	pdf.SetFont("Arial", "B", 11)
	if regreso {
		if ida || solicitudTraslado.Version == 2 {
			pdf.MoveTo(pdf.GetX()+46, pdf.GetY()-3)
			pdf.MultiCell(0, 3, tr(solicitudTraslado.Destino+" - "+solicitudTraslado.Origen), "", "L", false)
		} else {
			pdf.MoveTo(pdf.GetX()+46, pdf.GetY()-3)
			pdf.MultiCell(0, 3, tr(solicitudTraslado.Destino+" - "+solicitudTraslado.Origen), "", "L", false)
		}
	} else {
		pdf.MoveTo(pdf.GetX()+46, pdf.GetY()-3)
		pdf.MultiCell(0, 3, tr("NO APLICA"), "", "L", false)
	}
	pdf.Ln(3)
	/*****  Autorizar acompañante *****/
	pdf.SetFont("Arial", "", 11)
	pdf.MultiCell(60, 3, tr("SE AUTORIZA ACOMPAÑANTE:"), "", "L", false)
	/*****  ------------ Autorizar  *****/
	pdf.SetFont("Arial", "BU", 12)
	pdf.MoveTo(pdf.GetX()+59, pdf.GetY()-3)
	pdf.MultiCell(0, 3, tr(RegresarMjs(solicitudTraslado.Acompañante)), "", "L", false)
	pdf.Image(ImageFile("Modelos/ServiciosSubrogados/TrasladoPacientesModel/recursos/qr2.png"), 152, 64, 56, 0, false,
		"", 0, "")
	/*****  Leyenda boletos autorizados *****/
	pdf.MoveTo(pdf.GetX(), 117)
	pdf.SetFont("Arial", "BU", 12)
	pdf.MultiCell(0, 4, tr("BOLETOS AUTORIZADOS"), "", "C", false)
	/*****  Tabla de boletos  *****/
	pdf.Ln(4)
	pdf.SetX(pdf.GetX())
	pdf.SetFont("Arial", "B", 11)
	for _, str := range header {
		pdf.CellFormat(str.Ancho, 4.5, str.Nombre, "1", 0, "CM", false, 0, "")
	}
	pdf.SetFont("Arial", "", 11)
	pdf.Ln(-1)
	var observacionTmp string
	for _, c := range bole {
		pdf.SetX(pdf.GetX())
		pdf.MultiCell(26, 5, tr(c.tipo), "1", "L", false)
		pdf.MoveTo(pdf.GetX()+26, pdf.GetY()-10)
		pdf.MultiCell(15, 10, tr(c.ida), "1", "C", false)
		pdf.MoveTo(pdf.GetX()+41, pdf.GetY()-10)
		pdf.MultiCell(25, 10, tr(c.regreso), "1", "C", false)
		pdf.MoveTo(pdf.GetX()+66, pdf.GetY()-10)
		pdf.MultiCell(25, 10, tr(c.boleto), "1", "C", false)
		pdf.MoveTo(pdf.GetX()+91, pdf.GetY()-10)
		pdf.MultiCell(20, 10, tr(c.total), "1", "C", false)
		pdf.MoveTo(pdf.GetX()+111, pdf.GetY()-10)
		pdf.MultiCell(0, 10, tr(c.observaciones), "1", "C", false)
		observacionTmp = observacionTmp + "  " + c.observaciones
	}
	/*****  Observaciones *****/
	pdf.SetX(10.00125)
	pdf.MoveTo(pdf.GetX(), 162)
	/*****  -------- Observaciones  *****/
	pdf.SetFont("Arial", "B", 11)
	if resultAcomp {
		pdf.MultiCell(0, 3.5, tr("OBSERVACIONES: ")+tr(observacionTmp),
			"", "J", false)
	} else {
		pdf.MultiCell(0, 3.5, tr("SIN OBSERVACIONES"), "", "J", false)
	}
	pdf.MoveTo(pdf.GetX(), 168.5)
	/*****  VIGENCIA FTP *****/
	pdf.MultiCell(0, 5, tr("VIGENTE EN SUS DERECHOS CONFORME A LA CERTIFICACION DE VIGENCIA DE DERECHOS EN LA FTP-01"+
		":"), "", "L", false)
	/*****  ------------ Folio   *****/
	pdf.SetFont("Arial", "BU", 12)
	pdf.MoveTo(pdf.GetX()+16, pdf.GetY()-5)
	pdf.MultiCell(0, 5, tr(solicitudTraslado.Ftp01), "", "L", false)
	pdf.Ln(2.5)
	/*****  Fecha de expedicion *****/
	pdf.SetFont("Arial", "", 11)
	pdf.MultiCell(48, 4, tr("FECHA DE EXPEDICION:"), "", "L", false)
	/*****  --------- Fecha  *****/
	pdf.SetFont("Arial", "B", 11)
	pdf.MoveTo(pdf.GetX()+47, pdf.GetY()-4)
	pdf.MultiCell(0, 4, solicitudTraslado.FechaExpedicion.Format("02/01/2006")+"    "+
		solicitudTraslado.FechaExpedicion.Format("15:04:05")+" hrs.", "", "L", false)
	pdf.Ln(3)
	/*****  Leyenda documentos de identificacion  *****/
	pdf.SetFont("Arial", "B", 10.5)
	pdf.MultiCell(69, 4, tr("*DOCUMENTOS DE IDENTIFICACION:"), "", "L", false)
	pdf.MoveTo(pdf.GetX()+68.5, pdf.GetY()-4)
	pdf.SetFont("Arial", "I", 10.5)
	pdf.MultiCell(0, 4, tr("CREDENCIAL DE ELECTOR, LICENCIA DE MANEJO, "+"TARJETA DE "), "", "J", false)
	pdf.MultiCell(0, 4, tr("CONTROL DE CITAS IMSS CON FOTOGRAFIA O TARJETA DE IDENTIFICACION COMO PENSIONADO IMSS."), "", "L", false)
	/*****  Recepcion de boletos *****/
	pdf.MoveTo(pdf.GetX(), 208)
	pdf.SetFont("Arial", "", 11)
	pdf.MultiCell(61, 4, tr("RECEPCION DE BOLETO (S):"), "", "L", false)
	/*****  nombre del paciente *****/
	pdf.MoveTo(pdf.GetX()+61, pdf.GetY())
	pdf.CellFormat(0, 6, "NOMBRE Y FIRMA", "T", 0, "C", false, 0, "")
	pdf.Ln(15)
	pdf.MultiCell(95, 4, tr("VO. BO."), "", "C", false)
	pdf.MoveTo(pdf.GetX()+95, pdf.GetY()-4)
	pdf.MultiCell(95, 4, "AUTORIZA", "", "C", false)
	pdf.Ln(3)
	ultimo := pdf.GetY()
	pdf.MultiCell(90, 5, tr(strings.ToUpper(firmaVisto.FirmaSolicitud.Nombre)), "", "C", false)
	pdf.MoveTo(pdf.GetX()+100, ultimo)
	pdf.MultiCell(90, 5, tr(strings.ToUpper(firmaAutoriza.FirmaSolicitud.Nombre)), "", "C", false)
	pdf.MoveTo(pdf.GetX(), 249)
	cargo2 := pdf.GetY()
	pdf.MultiCell(90, 5, tr(strings.ToUpper(firmaVisto.FirmaSolicitud.Cargo)), "T", "C", false)
	pdf.MoveTo(pdf.GetX()+100, cargo2)
	pdf.MultiCell(90, 5, tr(strings.ToUpper(firmaAutoriza.FirmaSolicitud.Cargo)), "T", "C", false)
	if reimpresion != "" {
		nombre := solicitudTraslado.Usuario.Persona.Nombre + " " + solicitudTraslado.Usuario.Persona.ApellidoPat + " " + solicitudTraslado.Usuario.Persona.ApellidoMat
		hora := time.Now()
		pdf.Ln(7)
		fmt.Println(pdf.GetXY())
		pdf.SetFont("Arial", "B", 9)
		pdf.MultiCell(150, 4, "Reimprimio: "+nombre+"  Fecha: "+hora.Format("02/01/2006")+"  "+hora.Format("15:04:05")+" hrs.", "", "L", false)
	}
	// err = pdf.OutputFileAndClose("PlantillaSolicitudTraslado.pdf")
	// if err != nil {
	// 	return "", err
	// }
	// data, err := ioutil.ReadFile("PlantillaSolicitudTraslado.pdf")
	// if err != nil {
	// 	return "", err
	// }
	// os.Remove("PlantillaSolicitudTraslado.pdf")
	//return base64.StdEncoding.EncodeToString(data), nil
	return pdf, nil
}

var gofpdfDir string

func ImageFile(fileStr string) string {
	return filepath.Join(gofpdfDir, fileStr)
}
