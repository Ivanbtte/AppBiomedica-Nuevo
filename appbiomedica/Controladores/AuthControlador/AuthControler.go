package AuthControlador

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/AuthModel"
	"appbiomedica/Modelos/DepartamentosModel"
	"appbiomedica/Modelos/MenuModel"
	"appbiomedica/Modelos/RoleModel"
	"appbiomedica/Modelos/UsuarioModel"
	"appbiomedica/Modulos/Estructuras"
	"appbiomedica/Modulos/Sesion"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/tools/go/ssa/interp/testdata/src/errors"
)

//*****************************************************************************************************************
//login

func Login(c *gin.Context) {
	var credenciales AuthModel.Credenciales
	var status = false
	var menuRole *[]MenuModel.Menu
	var nameModule *[]RoleModel.RoleModule
	arrayUserRole := make([]string, 0)
	arrayPuestoRole := make([]string, 0)
	arrayRoleIds := make([]string, 0)
	err := c.BindJSON(&credenciales)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "error al analizar los datos", err.Error(), c)
		return
	}
	err, result := credenciales.Login()
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "error al consultar los datos de las credenciales",
			err.Error(), c)
		return
	}
	if !result.Estado {
		Estructuras.Responder(http.StatusInternalServerError, "Lo siento este usuario se ha dado de baja del sistema",
			errors.New("Error").Error(), c)
		return
	}
	datoPuesto, err := DepartamentosModel.ConsultaUnUsuarioPuesto(result.UsuarioID)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos usuario ", err.Error(), c)
		return
	}

	/***** buscar los roles que tiene asignado el puesto *****/
	err, puestoOrganigramaRole := DepartamentosModel.ConsultPuestoOrganigramaRoleByPuesto(datoPuesto.Id)
	if len(puestoOrganigramaRole) > 0 {
		for _, item := range puestoOrganigramaRole {
			arrayPuestoRole = append(arrayPuestoRole, item.RoleId)
		}
	}
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos usuario ", err.Error(), c)
		return
	}
	/***** buscar los roles que tiene asignado el usuario en el puesto *****/
	err, usuarioPuestoRole := DepartamentosModel.ConsultUsuarioPuestoRole(datoPuesto.Id)
	if len(usuarioPuestoRole) > 0 {
		fmt.Println("Usuario puesto por rol", usuarioPuestoRole)
		for _, item := range usuarioPuestoRole {
			arrayUserRole = append(arrayUserRole, item.RoleId)
		}
	}
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos usuario ", err.Error(), c)
		return
	}

	/*****  *****/
	arrayRoleIds = append(arrayPuestoRole, arrayUserRole...)
	if !datoPuesto.Estado && len(arrayRoleIds) == 0 {
		fmt.Println("este usuario no tiene puesto")
		//arrayRoleIds = append(arrayRoleIds, "614c97bbd7be0a789bb6f02c")
		status = true
	}
	_, menuRole = MenuModel.ConsultMenuByRole(arrayRoleIds)
	err, nameModule = RoleModel.ConsultNameModule(arrayRoleIds)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "error al consultar permisos", err.Error(), c)
		return
	}
	//err, menu := DepartamentosModel.ConsultPermitByRole(datoPuesto.PuestoOrganigramaId)
	//crear token
	var auth Sesion.UsuarioLogeado
	auth.UsuarioId = result.UsuarioID
	auth.Correo = result.Correo
	auth.Tipo = datoPuesto.Usuario.UsuarioTipoId
	auth.Puesto = datoPuesto.PuestoOrganigrama.CatalogoPuesto.Puesto
	auth.Nombre = datoPuesto.Usuario.Persona.Nombre
	auth.AppellidoP = datoPuesto.Usuario.Persona.ApellidoPat
	auth.ApellidoM = datoPuesto.Usuario.Persona.ApellidoMat
	auth.Matricula = datoPuesto.Usuario.Persona.Matricula
	auth.PersonaId = datoPuesto.Usuario.PersonaId
	auth.PuestoId = datoPuesto.PuestoOrganigrama.CatalogoPuestoId
	auth.PuestoOrganigramaId = datoPuesto.PuestoOrganigramaId
	auth.MenuItem = *menuRole
	token, err := Sesion.CreateTokenString(auth.UsuarioId, auth.Correo, auth.Tipo, auth.Puesto,
		auth.Nombre, auth.AppellidoP, auth.ApellidoM, auth.Matricula, auth.PersonaId, auth.PuestoId,
		auth.PuestoOrganigramaId, auth.MenuItem)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Ocurrio un problema al generar su token intente de nuevo.",
			err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Sesión iniciada", gin.H{"token": token, "status": status, "permits": nameModule}, c)
}

func Auth(c *gin.Context) {
	valido, err := Sesion.ValidaToken(c.Request)
	if !valido {
		Estructuras.Responder(http.StatusUnauthorized, "No has iniciado sesión", err, c)
		c.Abort()
		return
	}
	c.Next()
}

//Restablecer contraseña de un usuario

func ResetPassword(c *gin.Context) {
	var nuevasCredenciales struct {
		Password       string `json:"nuevoPassword"`
		RepitePassword string `json:"repitePassword"`
	}
	err := c.BindJSON(&nuevasCredenciales)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error al obtener los datos de la nueva contraseña", err.Error(), c)
		return
	}
	if nuevasCredenciales.Password != nuevasCredenciales.RepitePassword {
		Estructuras.Responder(http.StatusBadRequest, "La contraseña y su confirmación no coinciden", "", c)
		return
	}
	id := c.Param("idUsuario")
	if id == "" {
		Estructuras.Responder(http.StatusBadRequest, "El id del usuario es inválido", nil, c)
		return
	}
	_, err = UsuarioModel.ConsultaUnUsuario(id)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar la información del usuario", err.Error(), c)
		return
	}
	credenciales := AuthModel.Credenciales{
		Password: nuevasCredenciales.Password,
	}
	//inicio transaccion en la base de datos
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al conectar a la base de datos", tx.Error, c)
		return
	}
	if err = credenciales.ActualizarCredencial(id, tx); err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al actualizar la contraseña", err.Error(), c)
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al actualizar las credenciales", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Contraseña actualizada correctamente", "", c)
}

/*func ReturnPermitByUserPermit(permits []DepartamentosModel.UsuarioPuestoPermiso) []Sesion.Permits {
	permit := make([]Sesion.Permits, 0)
	for _, item := range permits {
		permit = append(permit, Sesion.Permits{Permiso: item.Permiso.Descripcion})
	}
	return permit
}
func ReturnPermitByRolePosition(permits []DepartamentosModel.RolPuesto) []Sesion.Permits {
	permit := make([]Sesion.Permits, 0)
	for _, item := range permits {
		permit = append(permit, Sesion.Permits{Permiso: item.Permiso.Descripcion})
	}
	return permit
}*/
