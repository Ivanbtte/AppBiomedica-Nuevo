package Pdf

import (
	"appbiomedica/Modelos/ServiciosSubrogados/TrasladoPacientesModel"
	"appbiomedica/Modulos/Estructuras"
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
)

/*****  Funcion para solicitar un PDF de una solicitud de traslado de pacientes sin datos  *****/
func GetPdfNoDates(ctx *gin.Context) {
	unitId := ctx.DefaultQuery("id", "")
	voboId := ctx.DefaultQuery("idVoBo", "")
	validateId := ctx.DefaultQuery("idAuthorize", "")
	id, err := strconv.Atoi(unitId)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error en el id de la unidad medica", err.Error(), ctx)
		return
	}
	mpdf, err := TrasladoPacientesModel.CreatePdfNoDate(id, voboId, validateId)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error al generar pdf", err.Error(), ctx)
		return
	}
	pdf, _ := base64.StdEncoding.DecodeString(mpdf)
	ctx.Writer.Header().Set("Content-Type", "application/pdf")
	buffer := bytes.NewBuffer([]byte(pdf))
	io.Copy(ctx.Writer, buffer)
	Estructuras.Responder(http.StatusOK, "Pdf generado correctamente", nil, ctx)
}

/*****  Funcion para solicitar un PDF solo los datos de una solicitud de traslado de pacientes  *****/
func GetPdfSoloDatos(ctx *gin.Context) {
	ftp := ctx.DefaultQuery("folio", "")
	reimpresion := ctx.DefaultQuery("reimpresion", "")
	folioSeparados := ctx.DefaultQuery("folioSeparados", "")
	fpdf := gofpdf.New("P", "mm", "Letter", "")

	if folioSeparados != "" {
		fmt.Println("Folios separados")
		folios := strings.Split(ftp, ",")
		pdf, err := TrasladoPacientesModel.CrearPdfConDatos(folios[0], reimpresion, fpdf)
		if err != nil {
			Estructuras.Responder(http.StatusInternalServerError, "Error al generar pdf", err.Error(), ctx)
			return
		}
		pdf2, err := TrasladoPacientesModel.CrearPdfConDatos(folios[1], reimpresion, pdf)
		if err != nil {
			Estructuras.Responder(http.StatusInternalServerError, "Error al generar pdf", err.Error(), ctx)
			return
		}
		err = pdf2.OutputFileAndClose("PlantillaSolicitudTraslado.pdf")
		if err != nil {
			Estructuras.Responder(http.StatusInternalServerError, "Error al generar pdf", err.Error(), ctx)
			return
		}
		data, err := ioutil.ReadFile("PlantillaSolicitudTraslado.pdf")
		if err != nil {
			Estructuras.Responder(http.StatusInternalServerError, "Error al generar pdf", err.Error(), ctx)
			return
		}
		os.Remove("PlantillaSolicitudTraslado.pdf")
		mpdf := base64.StdEncoding.EncodeToString(data)
		pdfF, _ := base64.StdEncoding.DecodeString(mpdf)
		ctx.Writer.Header().Set("Content-Type", "application/pdf")
		buffer := bytes.NewBuffer([]byte(pdfF))
		io.Copy(ctx.Writer, buffer)
		Estructuras.Responder(http.StatusOK, "Pdf generado correctamente", nil, ctx)
		return
	} else {
		fmt.Println("Un solo folio")
		pdf, err := TrasladoPacientesModel.CrearPdfConDatos(ftp, reimpresion, fpdf)
		if err != nil {
			Estructuras.Responder(http.StatusInternalServerError, "Error al generar pdf", err.Error(), ctx)
			return
		}
		err = pdf.OutputFileAndClose("PlantillaSolicitudTraslado.pdf")
		if err != nil {
			Estructuras.Responder(http.StatusInternalServerError, "Error al generar pdf", err.Error(), ctx)
			return
		}
		data, err := ioutil.ReadFile("PlantillaSolicitudTraslado.pdf")
		if err != nil {
			Estructuras.Responder(http.StatusInternalServerError, "Error al generar pdf", err.Error(), ctx)
			return
		}
		os.Remove("PlantillaSolicitudTraslado.pdf")
		mpdf := base64.StdEncoding.EncodeToString(data)
		pdf2, _ := base64.StdEncoding.DecodeString(mpdf)
		ctx.Writer.Header().Set("Content-Type", "application/pdf")
		buffer := bytes.NewBuffer([]byte(pdf2))
		io.Copy(ctx.Writer, buffer)
		Estructuras.Responder(http.StatusOK, "Pdf generado correctamente", nil, ctx)
		return
	}
}

/*****  Funcion para solicitar un PDF completo de una solicitud de traslado de pacientes  *****/
func GetPdfCompleto(ctx *gin.Context) {
	ftp := ctx.DefaultQuery("folio", "")
	reimpresion := ctx.DefaultQuery("reimpresion", "")
	folioSeparados := ctx.DefaultQuery("folioSeparados", "")
	fpdf := gofpdf.New("P", "mm", "Letter", "")
	if folioSeparados != "" {
		fmt.Println("Folios separados")
		folios := strings.Split(ftp, ",")
		pdf, err := TrasladoPacientesModel.CrearPdfCompleto(folios[0], reimpresion, folioSeparados, fpdf)
		if err != nil {
			Estructuras.Responder(http.StatusInternalServerError, "Error al generar pdf", err.Error(), ctx)
			return
		}
		pdf2, err := TrasladoPacientesModel.CrearPdfCompleto(folios[1], reimpresion, folioSeparados, pdf)
		if err != nil {
			Estructuras.Responder(http.StatusInternalServerError, "Error al generar pdf", err.Error(), ctx)
			return
		}
		err = pdf2.OutputFileAndClose("PlantillaSolicitudTraslado.pdf")
		if err != nil {
			Estructuras.Responder(http.StatusInternalServerError, "Error al generar pdf", err.Error(), ctx)
			return
		}
		data, err := ioutil.ReadFile("PlantillaSolicitudTraslado.pdf")
		if err != nil {
			Estructuras.Responder(http.StatusInternalServerError, "Error al generar pdf", err.Error(), ctx)
			return
		}
		os.Remove("PlantillaSolicitudTraslado.pdf")
		mpdf := base64.StdEncoding.EncodeToString(data)
		pdfF, _ := base64.StdEncoding.DecodeString(mpdf)
		ctx.Writer.Header().Set("Content-Type", "application/pdf")
		buffer := bytes.NewBuffer([]byte(pdfF))
		io.Copy(ctx.Writer, buffer)
		os.Remove("Modelos/ServiciosSubrogados/TrasladoPacientesModel/recursos/qr2.png")
		Estructuras.Responder(http.StatusOK, "Pdf generado correctamente", nil, ctx)
		return
	} else {
		fmt.Println("Un solo folio")
		pdf, err := TrasladoPacientesModel.CrearPdfCompleto(ftp, reimpresion, folioSeparados, fpdf)
		if err != nil {
			Estructuras.Responder(http.StatusInternalServerError, "Error al generar pdf", err.Error(), ctx)
			return
		}
		err = pdf.OutputFileAndClose("PlantillaSolicitudTraslado.pdf")
		if err != nil {
			Estructuras.Responder(http.StatusInternalServerError, "Error al generar pdf", err.Error(), ctx)
			return
		}
		data, err := ioutil.ReadFile("PlantillaSolicitudTraslado.pdf")
		if err != nil {
			Estructuras.Responder(http.StatusInternalServerError, "Error al generar pdf", err.Error(), ctx)
			return
		}
		os.Remove("PlantillaSolicitudTraslado.pdf")
		mpdf := base64.StdEncoding.EncodeToString(data)
		pdf2, _ := base64.StdEncoding.DecodeString(mpdf)
		ctx.Writer.Header().Set("Content-Type", "application/pdf")
		buffer := bytes.NewBuffer([]byte(pdf2))
		io.Copy(ctx.Writer, buffer)
		os.Remove("Modelos/ServiciosSubrogados/TrasladoPacientesModel/recursos/qr2.png")
		Estructuras.Responder(http.StatusOK, "Pdf generado correctamente", nil, ctx)
		return
	}
}
