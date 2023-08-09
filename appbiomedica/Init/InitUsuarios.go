package Init

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/AuthModel"
	"appbiomedica/Modelos/DepartamentosModel"
	"appbiomedica/Modelos/PermisoModel"
	"appbiomedica/Modelos/PersonaModel"
	"appbiomedica/Modelos/UsuarioModel"
)

func InitUsuarios() error {
	db := Conexion.GetDB()

	//Creando tabla de personas
	if !db.HasTable(&PersonaModel.Persona{}) {

		if err := db.CreateTable(&PersonaModel.Persona{}).Error; err != nil {
			return err
		}
	}
	//Creando tabla de permiso
	if !db.HasTable(&PermisoModel.Permiso{}) {

		if err := db.CreateTable(&PermisoModel.Permiso{}).Error; err != nil {
			return err
		}
	}

	/*****  Tabla de tipo de usuario  *****/
	if !db.HasTable(&UsuarioModel.UsuarioTipo{}) {
		if err := db.CreateTable(&UsuarioModel.UsuarioTipo{}).Error; err != nil {
			return err
		}
	}
	//Creando tabla de usuarios
	if !db.HasTable(&UsuarioModel.Usuario{}) {
		if err := db.CreateTable(&UsuarioModel.Usuario{}).
			AddForeignKey("persona_id", "persona(id)", "CASCADE", "CASCADE").
			AddForeignKey("usuario_tipo_id", "usuario_tipo(id)", "CASCADE", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	//Creando tabla de credenciales
	if !db.HasTable(&AuthModel.Credenciales{}) {

		if err := db.CreateTable(&AuthModel.Credenciales{}).AddForeignKey("usuario_id", "usuario(id)",
			"RESTRICT", "CASCADE").Error; err != nil {
			return err
		}
	}
	//********************************************************************************************************
	/* *********************    Creando Temporalmente la tabla de usuario puesto     ********************* */
	if !db.HasTable(&DepartamentosModel.UsuarioPuesto{}) {

		if err := db.CreateTable(&DepartamentosModel.UsuarioPuesto{}).
			AddForeignKey("usuario_id", "usuario(id)", "CASCADE", "CASCADE").
			AddForeignKey("puesto_organigrama_id", "puesto_organigrama(id)", "CASCADE", "CASCADE").
			Error; err != nil {
			return err
		}
	}

	/***** Tabla para asignar a un usuario en un puesto un rol *****/
	if !db.HasTable(&DepartamentosModel.UsuarioPuestoRole{}) {
		if err := db.CreateTable(&DepartamentosModel.UsuarioPuestoRole{}).
			AddForeignKey("usuario_puesto_id", "usuario_puesto(id)", "SET NULL", "CASCADE").
			AddForeignKey("role_id", "role(id)", "SET NULL", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	/***** Tabla para quitar o poner permisos a una accion de un módulo*****/
	if !db.HasTable(&DepartamentosModel.ActionExceptionByUser{}) {
		if err := db.CreateTable(&DepartamentosModel.ActionExceptionByUser{}).
			AddForeignKey("usuario_puesto_id", "usuario_puesto(id)", "SET NULL", "CASCADE").
			AddForeignKey("action_id", "action(id)", "CASCADE", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	if !db.HasTable(&DepartamentosModel.TipoModificacion{}) {
		if err := db.CreateTable(&DepartamentosModel.TipoModificacion{}).Error; err != nil {
			return err
		}
	}
	/***** Tabla para almacenar la bitacora de las validaciones de usuarios en puestos *****/
	if !db.HasTable(&DepartamentosModel.BitacoraUsuarioPuesto{}) {

		if err := db.CreateTable(&DepartamentosModel.BitacoraUsuarioPuesto{}).
			AddForeignKey("usuario_modificado_id", "usuario_puesto(id)", "SET NULL", "CASCADE").
			AddForeignKey("usuario_modificante_id", "usuario_puesto(id)", "SET NULL", "CASCADE").
			AddForeignKey("tipo_modificacion_id", "tipo_modificacion(id)", "SET NULL", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	return nil
}
