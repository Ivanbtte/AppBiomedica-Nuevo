package Utileria

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"time"
)

func CrearPdfFinal() {
	fecha := time.Now().Format("02/01/2006")
	pdf := gofpdf.New("L", "mm", "Legal", "")
	_, mright, _, _ := pdf.GetMargins()
	fmt.Println(mright)
	pdf.AddPage()
	pdf.Image("Modelos/SolicitudModel/Utileria/imssLogoFinal.png", 0, 0, 12, 0, false, "PNG", 0, "")
	pdf.Image("Modelos/SolicitudModel/Utileria/gobierno.png", 270, -5, 80, 0, false, "PNG", 0, "")
	pdf.SetFont("Arial", "B", 10)
	pdf.MoveTo(0, 25)
	pdf.CellFormat(100, 5, "0", "0", 1, "L", false, 0, "")

	pdf.MoveTo(26, 10)
	pdf.CellFormat(100, 5, "INSTITUTO MEXICANO DEL SEGURO SOCIAL", "0", 1, "L", false, 0, "")
	pdf.MoveTo(26, 15)
	pdf.CellFormat(100, 5, "DELEGACION ESTATAL OAXACA", "0", 1, "L", false, 0, "")
	pdf.MoveTo(26, 20)
	pdf.CellFormat(100, 5, "JEFATURA DE SERVICIOS ADMINISTRATIVOS", "0", 1, "L", false, 0, "")
	pdf.MoveTo(26, 25)
	pdf.CellFormat(100, 5, "COORDINACION DE ABASTECIMIENTO Y EQUIPAMIENTO", "0", 1, "L", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(170, 5, "Adquisicion Local.", "0", 0, "L", false, 0, "")
	pdf.CellFormat(100, 5, "No. De Requisicion:", "0", 0, "R", false, 0, "")
	pdf.CellFormat(100, 5, "21 379 19 0232", "0", 1, "L", false, 0, "")
	pdf.CellFormat(170, 5, "Almacen General Delegacional.", "0", 0, "L", false, 0, "")
	pdf.CellFormat(100, 5, "Fecha de Elaboracion:", "0", 0, "R", false, 0, "")
	pdf.CellFormat(100, 5, fecha, "0", 1, "L", false, 0, "")
	pdf.CellFormat(170, 5, "Boulevard Guadalupe Hinojosa de Murat.", "0", 0, "L", false, 0, "")
	pdf.CellFormat(100, 5, "Fecha requerida:", "0", 0, "R", false, 0, "")
	pdf.CellFormat(100, 5, "Conforma a pedido", "0", 1, "L", false, 0, "")
	pdf.CellFormat(170, 5, "Unidad Emisora:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(100, 5, "Lugar de entrega:", "0", 0, "R", false, 0, "")
	pdf.CellFormat(100, 5, "Almacen Delegacional", "0", 1, "L", false, 0, "")
	pdf.WriteAligned(0, 35, "This text is aligned Center", "C")

	err := pdf.OutputFileAndClose("prueba1.pdf")
	if err != nil {
		fmt.Println(err)
	}
}
