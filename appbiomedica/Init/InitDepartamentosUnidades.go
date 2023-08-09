package Init

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/AreaModel"
	"appbiomedica/Modelos/DepartamentosModel"
	"appbiomedica/Modelos/RoleModel"
)

func InitDepartamentos() error {
	db := Conexion.GetDB()
	/*****  Creando la tabla de tipos de departamentos  *****/
	if !db.HasTable(&DepartamentosModel.DepartamentoTipo{}) {
		if err := db.CreateTable(&DepartamentosModel.DepartamentoTipo{}).Error; err != nil {
			return err
		}
	}
	/*****  Tabla de areas para departamentos y usuarios  *****/
	if !db.HasTable(&AreaModel.Area{}) {
		if err := db.CreateTable(AreaModel.Area{}).Error; err != nil {
			return err
		}
	}
	/* ***    Creando la tabla para los departamentos, coordinaciones, jefaturas etc.     *** */
	if !db.HasTable(&DepartamentosModel.Departamento{}) {
		if err := db.CreateTable(&DepartamentosModel.Departamento{}).
			AddForeignKey("departamento_tipo_id", "departamento_tipo(id)", "CASCADE", "CASCADE").
			AddForeignKey("usuario_tipo_id", "usuario_tipo(id)", "CASCADE", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	/*****  Tabla OrganigramaUnidad  *****/
	if !db.HasTable(&DepartamentosModel.OrganigramaUnidad{}) {
		if err := db.CreateTable(&DepartamentosModel.OrganigramaUnidad{}).
			AddForeignKey("departamento_id", "departamento(id)", "CASCADE", "CASCADE").
			AddForeignKey("area_id", "area(id)", "CASCADE", "CASCADE").
			AddForeignKey("unidad_med_id", "unidad_med(id)", "CASCADE", "CASCADE").
			AddForeignKey("organigrama_parent", "organigrama_unidad(id)", "CASCADE", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	/*****  Tabla OrganigramaUnidad para delegaciones  *****/
	if !db.HasTable(&DepartamentosModel.OrganigramaDelegacion{}) {
		if err := db.CreateTable(&DepartamentosModel.OrganigramaDelegacion{}).
			AddForeignKey("departamento_id", "departamento(id)", "CASCADE", "CASCADE").
			AddForeignKey("area_id", "area(id)", "CASCADE", "CASCADE").
			AddForeignKey("delegaciones_clv_dele", "delegaciones(clv_dele)", "CASCADE", "CASCADE").
			AddForeignKey("organigrama_parent", "organigrama_delegacion(id)", "CASCADE", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	/*****  Tabla Puestos Organigramas  *****/
	if !db.HasTable(&DepartamentosModel.PuestoOrganigrama{}) {
		if err := db.CreateTable(&DepartamentosModel.PuestoOrganigrama{}).
			AddForeignKey("organigrama_unidad_id", "organigrama_unidad(id)", "CASCADE", "CASCADE").
			AddForeignKey("organigrama_delegacion_id", "organigrama_delegacion(id)", "CASCADE", "CASCADE").
			AddForeignKey("catalogo_puesto_id", "catalogo_puesto(id)", "CASCADE", "CASCADE").
			Error; err != nil {
			return err
		}
	}

	/***** Tabla para almacenar los roles *****/
	if !db.HasTable(&RoleModel.Role{}) {
		if err := db.CreateTable(&RoleModel.Role{}).
			Error; err != nil {
			return err
		}
	}
	/***** Tabla para almacenar los módulos que pueden acceder los roles *****/
	if !db.HasTable(&RoleModel.RoleModule{}) {
		if err := db.CreateTable(&RoleModel.RoleModule{}).
			AddForeignKey("role_id", "role(id)", "SET NULL", "CASCADE").
			AddForeignKey("modulo_id", "modulo(id)", "SET NULL", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	/***** Tabla para almacenar las acciones por módulo que tiene permitido cada rol *****/
	if !db.HasTable(&RoleModel.RoleModuleAction{}) {
		if err := db.CreateTable(&RoleModel.RoleModuleAction{}).
			AddForeignKey("role_module_id", "role_module(id)", "SET NULL", "CASCADE").
			AddForeignKey("action_id", "action(id)", "SET NULL", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	/***** Tabla para guardar los roles predeterminados por puesto *****/
	if !db.HasTable(&DepartamentosModel.PuestoOrganigramaRole{}) {
		if err := db.CreateTable(&DepartamentosModel.PuestoOrganigramaRole{}).
			AddForeignKey("puesto_organigrama_id", "puesto_organigrama(id)", "SET NULL", "CASCADE").
			AddForeignKey("role_id", "role(id)", "SET NULL", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	return nil
}
