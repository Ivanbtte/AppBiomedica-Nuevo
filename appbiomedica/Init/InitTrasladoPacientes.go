package Init

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/FtpModelo"
	"appbiomedica/Modelos/ServiciosSubrogados/TrasladoPacientesModel"
)

func InitTrasladoPacientes() error {
	db := Conexion.GetDB()
	/*****  Tabla de folios ftp  *****/
	if !db.HasTable(&FtpModelo.Ftp{}) {
		if err := db.CreateTable(&FtpModelo.Ftp{}).Error; err != nil {
			return err
		}
	}
	/*****  Tabla de tipos de pasajeros  *****/
	if !db.HasTable(&TrasladoPacientesModel.TipoPasajero{}) {
		if err := db.CreateTable(&TrasladoPacientesModel.TipoPasajero{}).Error; err != nil {
			return err
		}
	}
	/*****  Tabla de personas y cargos que firman una solicitud  *****/
	if !db.HasTable(&TrasladoPacientesModel.FirmaSolicitud{}) {
		if err := db.CreateTable(&TrasladoPacientesModel.FirmaSolicitud{}).
			AddForeignKey("unidad_med_id", "unidad_med(id)", "RESTRICT", "CASCADE").
			AddForeignKey("usuario_id", "usuario(id)", "RESTRICT", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	/*****  Tabla de Solicitudes de traslado  *****/
	if !db.HasTable(&TrasladoPacientesModel.SolicitudTraslado{}) {
		if err := db.CreateTable(&TrasladoPacientesModel.SolicitudTraslado{}).
			AddForeignKey("contrato_numero_contrato", "contrato(numero_contrato)", "CASCADE", "CASCADE").
			AddForeignKey("unidad_med_id", "unidad_med(id)", "CASCADE", "CASCADE").
			AddForeignKey("usuario_id", "usuario(id)", "CASCADE", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	/*****  Tabla de boletos de una solicitud de traslado de pacientes  *****/
	if !db.HasTable(&TrasladoPacientesModel.Boletos{}) {
		if err := db.CreateTable(&TrasladoPacientesModel.Boletos{}).
			AddForeignKey("tipo_pasajero_id", "tipo_pasajero(id)", "CASCADE", "CASCADE").
			AddForeignKey("solicitud_traslado_folio", "solicitud_traslado(folio)", "RESTRICT", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	/*****  Tabla de boletos consumidos del contrato de una solicitud de traslado de pacientes  *****/
	if !db.HasTable(&TrasladoPacientesModel.RutasConsumidas{}) {
		if err := db.CreateTable(&TrasladoPacientesModel.RutasConsumidas{}).
			AddForeignKey("traslado_pacientes_id", "traslado_pacientes(id)", "RESTRICT", "CASCADE").
			AddForeignKey("boletos_id", "boletos(id)", "RESTRICT", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	/*****  Tabla para saber que firmas y que nombres llevara la solicitud  *****/
	if !db.HasTable(&TrasladoPacientesModel.FirmaTraslado{}) {
		if err := db.CreateTable(&TrasladoPacientesModel.FirmaTraslado{}).
			AddForeignKey("firma_solicitud_id", "firma_solicitud(id)", "CASCADE", "CASCADE").
			AddForeignKey("solicitud_traslado_folio", "solicitud_traslado(folio)", "RESTRICT", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	/*****  Tabla para saber los datos de la autorizacion del acompaniante  *****/
	if !db.HasTable(&TrasladoPacientesModel.Acompaniante{}) {
		if err := db.CreateTable(&TrasladoPacientesModel.Acompaniante{}).
			AddForeignKey("solicitud_traslado_folio", "solicitud_traslado(folio)", "RESTRICT", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	/*****  Tabla de bitacora de solicitudes canceladaas  *****/
	if !db.HasTable(&TrasladoPacientesModel.SolicitudesCanceladas{}) {
		if err := db.CreateTable(&TrasladoPacientesModel.SolicitudesCanceladas{}).
			AddForeignKey("solicitud_traslado_folio", "solicitud_traslado(folio)", "RESTRICT", "CASCADE").
			AddForeignKey("usuario_id", "usuario(id)", "RESTRICT", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	/*****  Tabla de bitacora para actualizaciones de firmas  *****/
	if !db.HasTable(&TrasladoPacientesModel.RecordUpdateFirm{}) {
		if err := db.CreateTable(&TrasladoPacientesModel.RecordUpdateFirm{}).
			AddForeignKey("unidad_med_id", "unidad_med(id)", "RESTRICT", "CASCADE").
			AddForeignKey("usuario_id", "usuario(id)", "RESTRICT", "CASCADE").
			AddForeignKey("firma_solicitud_id", "firma_solicitud(id)", "CASCADE", "CASCADE").

			Error; err != nil {
			return err
		}
	}
	/*****  Tabla de bitacora para habilitar o deshabilitar firmas  *****/
	if !db.HasTable(&TrasladoPacientesModel.RecordDisableFirm{}) {
		if err := db.CreateTable(&TrasladoPacientesModel.RecordDisableFirm{}).
			AddForeignKey("unidad_med_id", "unidad_med(id)", "RESTRICT", "CASCADE").
			AddForeignKey("usuario_id", "usuario(id)", "RESTRICT", "CASCADE").
			AddForeignKey("firma_solicitud_id", "firma_solicitud(id)", "CASCADE", "CASCADE").

			Error; err != nil {
			return err
		}
	}

		/*****  Tabla de bitacora de solicitudes con autorizaciones especiales  *****/
		if !db.HasTable(&TrasladoPacientesModel.AutorizacionesEspeciales{}) {
			if err := db.CreateTable(&TrasladoPacientesModel.AutorizacionesEspeciales{}).
				AddForeignKey("solicitud_traslado_folio", "solicitud_traslado(folio)", "SET NULL", "CASCADE").
				Error; err != nil {
				return err
			}
		}


		/*****  Tabla para asignar aciones especiales de un modulo a una unidad medica  *****/
		if !db.HasTable(&TrasladoPacientesModel.PermisosPorUnidad{}) {
			if err := db.CreateTable(&TrasladoPacientesModel.PermisosPorUnidad{}).
				AddForeignKey("action_id", "action(id)", "SET NULL", "CASCADE").
				AddForeignKey("unidad_med_id", "unidad_med(id)", "SET NULL", "CASCADE").
				Error; err != nil {
				return err
			}
		}
	return nil
}
