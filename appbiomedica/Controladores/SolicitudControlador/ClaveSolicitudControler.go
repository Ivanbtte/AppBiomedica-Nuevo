package SolicitudControlador

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/BitacoraModel"
	"appbiomedica/Modelos/SolicitudModel"
	"appbiomedica/Modulos/Estructuras"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
	"time"
)

//*****************************************************************************************************************
/* *********************    Controlador para consultar una clave    ********************* */
func GetClave(c *gin.Context) {
	id := c.Param("id")
	err, resultado := SolicitudModel.ConcultaUnConceptoId(id)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}

//*****************************************************************************************************************
/* *********************    Controlador para consultar las solicitudes por usuario     ********************* */
func GetClavesSolicitud(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	clave := c.DefaultQuery("clave", "")
	unidad := c.DefaultQuery("unidad", "")
	final := c.DefaultQuery("final", "")
	idServicio := c.DefaultQuery("servicio", "")
	err, resultado := SolicitudModel.ConsultaConceptos(id, clave, unidad, final, idServicio)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}

//*****************************************************************************************************************
/* *********************    Controlador para agregar una clave a una solicitud       ********************* */
func PostClaveSolicitud(c *gin.Context) {
	var clave_solicitud SolicitudModel.ClavesSolicitud
	unidadId := c.DefaultQuery("unidadId", "0")
	idUnidad, err := strconv.Atoi(unidadId)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Id unidad Incorrecto", err.Error(), c)
		return
	}
	err = c.BindJSON(&clave_solicitud)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error al castear la estructura", err.Error(), c)
		return
	}
	fmt.Print(clave_solicitud)
	estado, resultado := SolicitudModel.BuscaClave(clave_solicitud.ClaveId, clave_solicitud.ServiciosProformaId)
	fmt.Println("esto regrese:----", resultado)
	if estado == true {
		fmt.Println("regrese estado true")
		if resultado == idUnidad {
			fmt.Println("si es lo mismo", resultado, idUnidad)
			Estructuras.Responder(http.StatusBadRequest, "Esta clave ya ha sido agregada", nil, c)
			return
		} else {
			/* *********************    entonces si agrego la clave      ********************* */
			clave_solicitud.Id = bson.NewObjectId().Hex()
			t := time.Now().Format("02/01/2006")
			clave_solicitud.Fecha = t
			clave_solicitud.Estado = false
			err = clave_solicitud.AgregarClavesSolicitud()
			if err != nil {
				Estructuras.Responder(http.StatusInternalServerError, "Error al insertar la clave", err.Error(), c)
				return
			}
			Estructuras.Responder(http.StatusOK, "Clave agregada correctamente", clave_solicitud, c)
			return
		}
	} else {
		clave_solicitud.Id = bson.NewObjectId().Hex()
		clave_solicitud.Estado = false
		t := time.Now().Format("02/01/2006")
		clave_solicitud.Fecha = t
		err = clave_solicitud.AgregarClavesSolicitud()
		if err != nil {
			Estructuras.Responder(http.StatusInternalServerError, "Error al insertar la clave", err.Error(), c)
			return
		}
		Estructuras.Responder(http.StatusOK, "Clave agregada correctamente", clave_solicitud, c)
		return
	}
}

//*****************************************************************************************************************
/* *********************    Controlador para actuzalizar una clave y hacer registro a la bitacora     ********************* */
func UpdateClave(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	comentario := c.DefaultQuery("comentario", " ")
	cantidadAprobada := c.DefaultQuery("cantidad_apro", "")
	aprobarEstado := c.DefaultQuery("aprobar", "")
	usuarioId := c.DefaultQuery("usuarioId", "")
	var claves SolicitudModel.ClavesSolicitud
	err := c.BindJSON(&claves)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error en la estructura", err.Error(), c)
	}
	if id == "" {
		Estructuras.Responder(http.StatusBadRequest, "Debes ingresar el id", nil, c)
		return
	}
	fmt.Println(claves)
	//inicio transaccion en la base de datos
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al iniciar la transaccion", tx.Error, c)
		return
	}
	if cantidadAprobada != "" {
		canti, err := strconv.Atoi(cantidadAprobada)
		if err != nil {
			Estructuras.Responder(http.StatusBadRequest, "Cantidad incorrecta", err.Error(), c)
			return
		}
		err, resultado := claves.AprobarConcepto(id, canti, tx)
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
		Estructuras.Responder(http.StatusOK, "Cantidad aprobada correctamente", resultado, c)
		return
	}
	if aprobarEstado != "" {
		err := claves.ApruebaEstadoclave(id, tx)
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
		Estructuras.Responder(http.StatusOK, "Cantidad aprobada correctamente", nil, c)
		return
	}
	err, resultado := claves.ActualizarClaveConcepto(id, tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al actualizar datos", err.Error(), c)
		return
	}
	fmt.Println(resultado)
	var bitacora BitacoraModel.BitacoraClave
	t := time.Now().Format("02/01/2006")
	h := time.Now().Format("3:04PM")
	bitacora.Id = bson.NewObjectId().Hex()
	bitacora.Accion = "Actualizar"
	bitacora.ClaveId = claves.ClaveId
	bitacora.EstatusId = 1
	bitacora.Fecha = t
	bitacora.Hora = h
	bitacora.Comentario = comentario
	bitacora.UsuarioId = usuarioId
	bitacora.Estado = true
	err = bitacora.AgregaBitacoraClave(tx)
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
	Estructuras.Responder(http.StatusOK, "Datos actualizados correctamente", resultado, c)
	return
}

//*****************************************************************************************************************
/* *********************    Controlador para eliminar una clave y hacer registro a la bitacora     ********************* */
func DeleteClave(c *gin.Context) {
	id := c.Param("id")
	comentario := c.DefaultQuery("comentario", " ")
	var claves SolicitudModel.ClavesSolicitud
	if id == "" {
		Estructuras.Responder(http.StatusBadRequest, "Debes ingresar el id", nil, c)
		return
	}
	//inicio transaccion en la base de datos
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al iniciar la transaccion", tx.Error, c)
		return
	}
	err, resultado := claves.EliminarClaveConcepto(id, tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al actualizar datos", err.Error(), c)
		return
	}
	fmt.Println("esta es el resultado liminado", resultado)
	var bitacora BitacoraModel.BitacoraClave
	t := time.Now().Format("02/01/2006")
	bitacora.Id = bson.NewObjectId().Hex()
	bitacora.Accion = "Eliminar"
	bitacora.ClaveId = resultado.ClaveId
	bitacora.EstatusId = 1
	bitacora.Fecha = t
	bitacora.Comentario = comentario
	bitacora.UsuarioId = resultado.Solicitud.UsuarioId
	bitacora.Estado = true
	err = bitacora.AgregaBitacoraClave(tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al insertar los datos en la bitacora", err.Error(), c)
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al Eliminar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "clave eliminada correctamente", claves, c)
	return
}
