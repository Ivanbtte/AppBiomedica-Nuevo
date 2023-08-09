package SolicitudControlador

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/BitacoraModel"
	"appbiomedica/Modelos/FechaLimiteModel"
	"appbiomedica/Modelos/SolicitudModel"
	"appbiomedica/Modulos/Estructuras"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
	"time"
)

//*****************************************************************************************************************
/* *********************    Controlador para consultar tolas las  solicitudes por usuario     ********************* */
func GetTodasSolicitud(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	validado := c.DefaultQuery("validadas", "")
	err, resultado := SolicitudModel.ConsultaTodasSolicitudes(id, validado)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}

//*****************************************************************************************************************
/* *********************    Controlador para consultar las solicitudes por usuario     ********************* */
func GetUnaSolicitud(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	tipo := c.DefaultQuery("tipo", "0")
	idsolicitud := c.DefaultQuery("solicitud", "0")
	err, resultado := SolicitudModel.ConsultaSolicitudes(id, tipo, idsolicitud)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	if resultado.EstatusId == 17 {
		Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", nil, c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}

//*****************************************************************************************************************
/* *********************    Controlador para agregar una solicitud       ********************* */
func PostSolicitud(c *gin.Context) {
	var solicitud SolicitudModel.Solicitud
	err := c.BindJSON(&solicitud)
	fmt.Print(solicitud)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error al castear la estructura", err.Error(), c)
		return
	}

	solicitud.Id = bson.NewObjectId().Hex()
	solicitud.Estado = false
	t := time.Now().Format("02/01/2006")
	solicitud.FechaCreacion = t
	err = solicitud.AgregarSolicitud()
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al insertar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos agregados correctamente", solicitud, c)
	return
}

//*****************************************************************************************************************
/* *********************    Controlador para actuzalizar una solicitud y hacer registro a la bitacora     ********************* */
func UpdateSolicitud(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	comentario := c.DefaultQuery("comentario", " ")
	estatus := c.DefaultQuery("estatus", "0")
	total := c.DefaultQuery("total", "0")
	usuarioId := c.DefaultQuery("usuarioId", "0")
	totalSol, err := strconv.ParseFloat(total, 64)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Numero de registros incorrecto", err.Error(), c)
		return
	}
	var solicitud SolicitudModel.Solicitud
	err = c.BindJSON(&solicitud)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error en la estructura", err.Error(), c)
	}
	result, err := FechaLimiteModel.VerificaFecha()
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Aun no puedes mandar una solicitud", err.Error(), c)
		return
	}
	tActual := time.Now().Format("2/1/2006")
	ta, _ := time.Parse("2/1/2006", tActual)
	fLimi, _ := time.Parse("2/1/2006", result.FechaLimite)
	fIni, _ := time.Parse("2/1/2006", result.FechaInicio)
	fmt.Println(fIni, ta, fLimi)
	if !ta.After(fIni) || !ta.Before(fLimi) {
		Estructuras.Responder(http.StatusBadRequest, "No puedes mandar una solicitud", errors.New("Estas fuera de fechas").Error(), c)
		return
	}
	fmt.Println(solicitud)
	//inicio transaccion en la base de datos
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al iniciar la transaccion", tx.Error, c)
		return
	}
	if total != "0" {
		err = SolicitudModel.ActualizarTotalSolicitud(id, tx, totalSol)
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "Error al actualizar datos", err.Error(), c)
			return
		}
		err = tx.Commit().Error
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "Error al actualizar los datos", err.Error(), c)
			return
		}
		Estructuras.Responder(http.StatusOK, "Datos actualizados correctamente", solicitud, c)
		return
	}
	if estatus != "0" {
		status, err := strconv.Atoi(estatus)
		if err != nil {
			Estructuras.Responder(http.StatusBadRequest, "estatus incorrecto", err.Error(), c)
			return
		}
		err = SolicitudModel.ActualizarEstatuslSolicitud(solicitud.Id, tx, status)
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "Error al actualizar datos", err.Error(), c)
			return
		}
		var bitacora BitacoraModel.BitacoraSolicitud
		t := time.Now().Format("02/01/2006")
		h := time.Now().Format("3:04PM")
		bitacora.Id = bson.NewObjectId().Hex()
		bitacora.SolicitudId = solicitud.Id
		bitacora.EstatusId = status
		bitacora.Fecha = t
		bitacora.Hora = h
		bitacora.Comentario = comentario
		bitacora.Estado = true
		bitacora.UsuarioId = usuarioId
		err = bitacora.AgregaBitacoraSolicitud(tx)
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "Error al insertar los datos", err.Error(), c)
			return
		}
		err = tx.Commit().Error
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "Error al actualizar los datos", err.Error(), c)
			return
		}
		Estructuras.Responder(http.StatusOK, "Datos actualizados correctamente", solicitud, c)
		return
	}
	fechaEnviada := time.Now().Format("02/01/2006")
	solicitud.FechaCreacion = fechaEnviada
	err = solicitud.ActualizarSolicitud(id, tx)

	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al actualizar datos", err.Error(), c)
		return
	}
	var bitacora BitacoraModel.BitacoraSolicitud
	t := time.Now().Format("02/01/2006")
	h := time.Now().Format("3:04PM")
	bitacora.Id = bson.NewObjectId().Hex()
	bitacora.SolicitudId = solicitud.Id
	bitacora.EstatusId = solicitud.EstatusId
	bitacora.Fecha = t
	bitacora.Hora = h
	bitacora.Comentario = comentario
	bitacora.Estado = true
	bitacora.UsuarioId = solicitud.UsuarioId
	err = bitacora.AgregaBitacoraSolicitud(tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al insertar los datos", err.Error(), c)
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al actualizar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos actualizados correctamente", solicitud, c)
	return
}
