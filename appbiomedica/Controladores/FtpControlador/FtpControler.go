package FtpControlador

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/FtpModelo"
	"appbiomedica/Modelos/ServiciosSubrogados/TrasladoPacientesModel"
	"appbiomedica/Modulos/Estructuras"
	"net/http"
	"strconv"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gin-gonic/gin"
)

func GetFolioFtp(ctx *gin.Context) {
	folio := ctx.DefaultQuery("folio", "")
	resultado, _, err := FtpModelo.ConsultaFolioFtp(folio)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, err.Error(), nil, ctx)
		return
	}
	if resultado.Transporte != "Transporte Subrogado" && resultado.ImporteEmitido != 0 {
		Estructuras.Responder(http.StatusInternalServerError, "Este folio FTP-01 no esta autorizado para transporte subrogado, ya que especifica", gin.H{"transporte": resultado.Transporte, "importe": resultado.ImporteEmitido}, ctx)
		return
	}
	duplicado, _ := TrasladoPacientesModel.BuscarFolioFtp(folio, "true")
	if duplicado {
		Estructuras.Responder(http.StatusInternalServerError,
			"Este folio FTP-01: "+folio+" ya se encuentra registrado", nil, ctx)
		return
	}
	Estructuras.Responder(http.StatusOK, "Folio Consultado", resultado, ctx)
}

func AgregarFoliosTP(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	f, err := file.Open()
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al abrir el archivo", err, ctx)
		return
	}
	defer f.Close()
	archivo, err := excelize.OpenReader(f)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al leer el archivo", err, ctx)
		return
	}
	rows, err := archivo.GetRows("Sheet1")
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error el archivo no es valido", err, ctx)
		return
	}
	var ftp FtpModelo.Ftp
	var contador int
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al iniciar la transaccion", tx.Error, ctx)
		return
	}
	for _, x := range rows {
		if len(x) > 0 {
			if len(x[0]) == 10 && len(x) == 17 {
				resultado, value, _ := FtpModelo.ConsultaFolioFtp(x[0])
				if resultado == nil && value {
					contador = contador + 1
					impEmit, _ := strconv.ParseFloat(x[14], 64)
					impPag, _ := strconv.ParseFloat(x[15], 64)
					ftp.FolioTp = x[0]
					ftp.TipoFolio = x[1]
					ftp.UnidadReceptora = x[2]
					ftp.Especialidad = x[3]
					ftp.FechaHora = x[4]
					ftp.Tipocita = x[5]
					ftp.Nss = x[6]
					ftp.Agregado = x[7]
					ftp.Paciente = x[8]
					ftp.Acompañante = x[9]
					ftp.Transporte = x[10]
					ftp.UsuarioEmitio = x[11]
					ftp.MotivoAjuste = x[12]
					ftp.FechaPago = x[13]
					ftp.ImporteEmitido = impEmit
					ftp.ImportePagado = impPag
					ftp.EstadoTp = x[16]
					err = ftp.AgregarFolioTp(tx)
					if err != nil {
						tx.Rollback()
						Estructuras.Responder(http.StatusInternalServerError, "Error al agregar el folio", err.Error(), ctx)
						return
					}
				}
			}
		}
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al realizar el commit", err.Error(), ctx)
		return
	}
	Estructuras.Responder(http.StatusOK, "Se han cargado correctamente", contador, ctx)
}
