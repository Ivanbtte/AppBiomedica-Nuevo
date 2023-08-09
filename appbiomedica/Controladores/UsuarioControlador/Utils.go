package UsuarioControlador

import (
	"appbiomedica/Modelos/DepartamentosModel"
)

func ConsultRolesIdsByUser(usuarioPuestoId string) (error, []string) {
	arrayUserRole := make([]string, 0)
	arrayPuestoRole := make([]string, 0)
	arrayRoleIds := make([]string, 0)
	/***** buscar los roles que tiene asignado el puesto *****/
	err, puestoOrganigramaRole := DepartamentosModel.ConsultPuestoOrganigramaRoleByPuesto(usuarioPuestoId)
	if len(puestoOrganigramaRole) > 0 {
		for _, item := range puestoOrganigramaRole {
			arrayPuestoRole = append(arrayPuestoRole, item.RoleId)
		}
	}
	if err != nil {
		return err, arrayRoleIds
	}
	/***** buscar los roles que tiene asignado el usuario en el puesto *****/
	err, usuarioPuestoRole := DepartamentosModel.ConsultUsuarioPuestoRole(usuarioPuestoId)
	if len(usuarioPuestoRole) > 0 {
		for _, item := range usuarioPuestoRole {
			arrayUserRole = append(arrayUserRole, item.RoleId)
		}
	}
	if err != nil {
		return err, arrayRoleIds

	}
	arrayRoleIds = append(arrayPuestoRole, arrayUserRole...)
	return nil, arrayRoleIds
}
