package Init

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/BitacoraModel"
	"appbiomedica/Modelos/EstatusModel"
	"appbiomedica/Modelos/SolicitudModel"
)

func InitSolicitud() error {
	db := Conexion.GetDB()
	/* *********************   Creandio la tabla Estatus     ********************* */
	if !db.HasTable(&EstatusModel.Estatus{}) {

		if err := db.CreateTable(&EstatusModel.Estatus{}).Error; err != nil {
			return err
		}
	}
	/* *********************    Creando la tabla Tipo de Solicitud     ********************* */
	if !db.HasTable(&SolicitudModel.TipoSolicitud{}) {

		if err := db.CreateTable(&SolicitudModel.TipoSolicitud{}).Error; err != nil {
			return err
		}
	}
	/* *********************    Creando la tabla de la solicitudes     ********************* */
	if !db.HasTable(&SolicitudModel.Solicitud{}) {

		if err := db.CreateTable(&SolicitudModel.Solicitud{}).
			AddForeignKey("usuario_id", "usuario(id)", "RESTRICT", "CASCADE").
			AddForeignKey("unidad_med_id", "unidad_med(id)", "RESTRICT", "CASCADE").
			AddForeignKey("tipo_solicitud_id", "tipo_solicitud(id)", "RESTRICT", "CASCADE").
			AddForeignKey("estatus_id", "estatus(id)", "RESTRICT", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	/* *********************    Creando la tabla de las claves de las solicitudes     ********************* */
	if !db.HasTable(&SolicitudModel.ClavesSolicitud{}) {

		if err := db.CreateTable(&SolicitudModel.ClavesSolicitud{}).
			AddForeignKey("solicitud_id", "solicitud(id)", "RESTRICT", "CASCADE").
			AddForeignKey("servicios_proforma_id", "servicios_proforma(id)", "RESTRICT", "CASCADE").
			AddForeignKey("estatus_id", "estatus(id)", "RESTRICT", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	/* *********************    Creando la tabla de la bitacora de solicitud     ********************* */
	if !db.HasTable(&BitacoraModel.BitacoraSolicitud{}) {

		if err := db.CreateTable(&BitacoraModel.BitacoraSolicitud{}).
			AddForeignKey("solicitud_id", "solicitud(id)", "RESTRICT", "CASCADE").
			AddForeignKey("estatus_id", "estatus(id)", "RESTRICT", "CASCADE").
			AddForeignKey("usuario_id", "usuario(id)", "RESTRICT", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	/* *********************    Creando la tabla de la bitacora de claves   ********************* */
	if !db.HasTable(&BitacoraModel.BitacoraClave{}) {

		if err := db.CreateTable(&BitacoraModel.BitacoraClave{}).
			AddForeignKey("usuario_id", "usuario(id)", "RESTRICT", "CASCADE").
			AddForeignKey("estatus_id", "estatus(id)", "RESTRICT", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	return nil
}
