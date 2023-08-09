package TrasladosControlador

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/ServiciosSubrogados/TrasladoPacientesModel"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
	"time"
)

/*****  Controlador para agregar los datos de la firma nueva  *****/
func PostFirmaSolicitud(c *gin.Context) {
	var firmas TrasladoPacientesModel.FirmaSolicitud
	err := c.BindJSON(&firmas)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error al castear la estructura", err.Error(), c)
		return
	}
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al iniciar la transaccion", tx.Error, c)
		return
	}
	firmas.Id = bson.NewObjectId().Hex()
	firmas.Fecha = time.Now()
	firmas.Estado = true
	duplicate := TrasladoPacientesModel.FindDuplicate(firmas.Matricula, firmas.UnidadMedId)
	if duplicate {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Esta matricula ya esta registrada", nil, c)
		return
	}
	err = firmas.AgregarFirma(tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al agregar los datos de las firmas",
			err.Error(), c)
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al realizar el commit", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos guardados correctamente", firmas, c)
	return
}

/*****  Controlador para consultar una firma por unidad medica  *****/
func GetFirmaSolicitud(ctx *gin.Context) {
	unidad_id := ctx.DefaultQuery("unidad_id", "")
	tipo := ctx.DefaultQuery("tipo", "")
	state := ctx.DefaultQuery("state", "")
	/*****    *****/
	unidad, err := strconv.Atoi(unidad_id)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Numero de unidad medica incorrecto", err.Error(), ctx)
		return
	}
	tpo_id, err := strconv.Atoi(tipo)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Tipo incorrecto", err.Error(), ctx)
		return
	}
	stateFirm, err := strconv.ParseBool(state)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al parsear booleano", err.Error(), ctx)
		return
	}
	err, firma := TrasladoPacientesModel.ReturnAllFirm(unidad, tpo_id, stateFirm)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Datos no consultados", err.Error(), ctx)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", firma, ctx)
	return
}

/*****  Controlador para actualizar un nombre y cargo de una firma  *****/
func EnableFirm(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "")
	estado := ctx.DefaultQuery("state", "")
	unitId := ctx.DefaultQuery("unitId", "")
	userId := ctx.DefaultQuery("userId", "")
	state, err := strconv.ParseBool(estado)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al parsear booleano", err.Error(), ctx)
		return
	}
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al iniciar la transaccion", tx.Error, ctx)
		return
	}
	err = TrasladoPacientesModel.ActualizarFirma(id, !state, tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al habilitar o inhabilitar la firma",
			err.Error(), ctx)
		return
	}
	option := ""
	if state {
		option = "Inactivar"
	} else {
		option = "Activar"
	}
	idUnit, _ := strconv.Atoi(unitId)
	var record TrasladoPacientesModel.RecordDisableFirm
	record.Id = bson.NewObjectId().Hex()
	record.Opcion = option
	record.Fecha = time.Now()
	record.UsuarioId = userId
	record.UnidadMedId = idUnit
	record.FirmaSolicitudId = id
	record.Estado = true
	err = record.AddUpdateFirm(tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al agregar los datos en la bitacora",
			err.Error(), ctx)
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al realizar el commit", err.Error(), ctx)
		return
	}
	Estructuras.Responder(http.StatusOK, "Firma actualizada correctamente", nil, ctx)
	return
}

/*****  Controlador para consultar todas las firmas por unidad medica  *****/
func GetFirmsByUnit(ctx *gin.Context) {
	unitId := ctx.DefaultQuery("unitId", "")
	id, err := strconv.Atoi(unitId)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Numero de unidad medica invalido", err.Error(), ctx)
		return
	}
	err, units := TrasladoPacientesModel.ConsultarFirmasUnidad(id, "")
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar las firmas", err.Error(), ctx)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados", units, ctx)
	return
}

/*****  Controlador para actualizar una firma de solicitud  *****/
func PutFirm(ctx *gin.Context) {
	firmId := ctx.DefaultQuery("id", "")
	unitId := ctx.DefaultQuery("unitId", "")
	matric := ctx.DefaultQuery("matric", "")
	userId := ctx.DefaultQuery("userId", "")
	var firm TrasladoPacientesModel.FirmaSolicitud
	err := ctx.BindJSON(&firm)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error al castear la estructura", err.Error(), ctx)
		return
	}
	idUnit, _ := strconv.Atoi(unitId)
	if matric != "" {
		duplicate := TrasladoPacientesModel.FindDuplicate(firm.Matricula, idUnit)
		if duplicate {
			Estructuras.Responder(http.StatusInternalServerError, "Esta matricula ya esta registrada", nil, ctx)
			return
		}
	}
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al iniciar la transaccion", tx.Error, ctx)
		return
	}
	err, result := TrasladoPacientesModel.ReturnFirm(firmId)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos ", err.Error(), ctx)
		return
	}
	var recordUpdate TrasladoPacientesModel.RecordUpdateFirm
	recordUpdate.Id = bson.NewObjectId().Hex()
	recordUpdate.Nombre = result.Nombre
	recordUpdate.Cargo = result.Cargo
	recordUpdate.Matricula = result.Matricula
	recordUpdate.Tipo = result.Tipo
	recordUpdate.Fecha = time.Now()
	recordUpdate.UsuarioId = userId
	recordUpdate.UnidadMedId = idUnit
	recordUpdate.FirmaSolicitudId = result.Id
	recordUpdate.Estado = true
	err = recordUpdate.AddUpdateFirm(tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al guardar en la bitacora ", err.Error(), ctx)
		return
	}
	err = firm.UpdateFirm(firmId, tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al intentar actualizar los datos", err.Error(),
			ctx)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos actualizados", firm, ctx)
	return
}
