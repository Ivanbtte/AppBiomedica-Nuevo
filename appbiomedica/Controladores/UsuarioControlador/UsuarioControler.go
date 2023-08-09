package UsuarioControlador

import (
	"appbiomedica/Conexion"
	"appbiomedica/Controladores/PersonaControlador/Utilerias"
	"appbiomedica/Modelos/AuthModel"
	"appbiomedica/Modelos/DepartamentosModel"
	"appbiomedica/Modelos/PersonaModel"
	"appbiomedica/Modelos/UsuarioModel"
	"appbiomedica/Modulos/Estructuras"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

//consulta todos los usuarios

func GetUsuarios(c *gin.Context) {
	estado := c.DefaultQuery("activos", "")
	b, err := strconv.ParseBool(estado)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al parsear booleano", err.Error(), c)
		return
	}
	err, resultado := UsuarioModel.ConsultaUsuario(b)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}

/* *********************    controlador para agregar un nuevo usuario     ********************* */

type DatosUsuario struct {
	Persona             PersonaModel.Persona
	Usuario             UsuarioModel.Usuario
	PuestoOrganigramaId string
}

func AgregarUsuario(c *gin.Context) {
	var DateUser DatosUsuario
	var personas PersonaModel.Persona
	var usuarios UsuarioModel.Usuario
	var credenciales AuthModel.Credenciales
	err := c.BindJSON(&DateUser)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error al analizar la estructura", err.Error(), c)
		return
	}
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al iniciar la base de datos", tx.Error, c)
		return
	}
	if DateUser.Usuario.PersonaId != "" && DateUser.Persona.Id != "" {
		fmt.Println("actualizar")
		err = DateUser.Persona.ActualizarPersona(DateUser.Persona.Id, tx)
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "Error al actualizar datos", err.Error(), c)
			return
		}
		err, credential := AuthModel.GetOneCredentialByPersonId(DateUser.Persona.Id)
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "error al buscar datos del usuario", err.Error(), c)
			return
		}
		DateUser.Usuario.Id = credential.UsuarioID
		DateUser.Usuario.Estado = true
		usuarios.Id = credential.UsuarioID
		err = DateUser.Usuario.UpdateUserInactive(credential.UsuarioID, tx)
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "error al actualizar el usuario", err.Error(), c)
			return
		}
		credenciales.Correo = DateUser.Usuario.Correo
		credenciales.Password = Estructuras.GetMD5Hash(DateUser.Usuario.Contraseña)
		credenciales.Estado = true
		err = credenciales.UpdateCredentialInactive(credential.ID, tx)
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "error al actualizar las credenciales",
				err.Error(), c)
			return
		}
	} else {
		personas = DateUser.Persona
		if PersonaModel.VerificaCorreoPersona(personas.Correo) {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "Este correo ya ha sido registrado", nil, c)
			return
		}
		personas.Id = bson.NewObjectId().Hex()
		personas.Estado = true
		err = personas.AgregaPersona(tx)
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "Error al insertar los datos", err.Error(), c)
			return
		}
		usuarios = DateUser.Usuario
		usuarios.Id = bson.NewObjectId().Hex()
		usuarios.PersonaId = personas.Id
		usuarios.FechaRegistro = time.Now()
		usuarios.Estado = true
		err = usuarios.AgregaUsuario(tx)
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "Error al insertar los datos", err.Error(), c)
			return
		}
		credenciales.ID = bson.NewObjectId().Hex()
		credenciales.UsuarioID = usuarios.Id
		credenciales.Correo = usuarios.Correo
		credenciales.Password = usuarios.Contraseña
		credenciales.Estado = true
		err = credenciales.AgregaCredencial(tx)
		if err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "Error al insertar los datos", err.Error(), c)
			return
		}
	}
	/*****  agregar un puesto a un usuario  *****/
	var usuarioPuesto DepartamentosModel.UsuarioPuesto
	usuarioPuesto.Id = bson.NewObjectId().Hex()
	usuarioPuesto.UsuarioId = usuarios.Id
	usuarioPuesto.PuestoOrganigramaId = DateUser.PuestoOrganigramaId
	usuarioPuesto.Valid = false
	usuarioPuesto.FechaRegistro = time.Now()
	usuarioPuesto.Estado = true
	err = usuarioPuesto.AgregaUsuarioPuesto(tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al asociar el puesto al usuario", err.Error(), c)
		return
	}
	/***** asignar el permiso por default al usuario *****/
	var usuarioPuestoRole DepartamentosModel.UsuarioPuestoRole
	usuarioPuestoRole.Id = bson.NewObjectId().Hex()
	usuarioPuestoRole.UsuarioPuestoId = usuarioPuesto.Id
	usuarioPuestoRole.RoleId = "614c97bbd7be0a789bb6f02c" //Role Consulta por default
	usuarioPuestoRole.Status = true
	err = usuarioPuestoRole.AddUsuarioPuestoRoleModule(tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al asignar el permiso por default", err.Error(), c)
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al insertar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Usuario registrado correctamente", usuarios, c)
}

//Actualiza un usuario

func UpdateUsuario(c *gin.Context) {
	id := c.Param("id")
	var usuarios UsuarioModel.Usuario
	var password UsuarioModel.Passwords
	err := c.BindJSON(&password)
	if id == "" {
		Estructuras.Responder(http.StatusBadRequest, "Debes ingresar el id", nil, c)
		return
	}
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error en la estructura", err.Error(), c)
	}
	fmt.Println(password)
	resultado, err := UsuarioModel.ConsultaUnUsuario(id)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	fmt.Println(resultado.Contraseña)
	if resultado.Contraseña != Estructuras.GetMD5Hash(password.Password) {
		Estructuras.Responder(http.StatusInternalServerError, "Error en la contraseña actual",
			errors.New("contraseña incorrecta"), c)
		return
	}
	//inicio transacción en la base de datos
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al iniciar la base de datos", tx.Error, c)
		return
	}
	usuarios.Contraseña = password.NewPassword
	err = usuarios.ActualizarUsuario(id, tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al actualizar datos", err.Error(), c)
		return
	}
	var credenciales AuthModel.Credenciales
	credenciales.Password = password.NewPassword
	err = credenciales.ActualizarCredencial(id, tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al actualizar los datos", err.Error(), c)
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al actualizar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos actualizados correctamente", usuarios, c)
}

/* *********************    Controlador para validar o bloquear un usuario     ********************* */

func ValidaUsuario(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	validar := c.DefaultQuery("valida", "")
	if validar == "" {
		Estructuras.Responder(http.StatusBadRequest, "Debes ingresar el id", nil, c)
		return
	}
	b, err := strconv.ParseBool(validar)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al parsear booleano", err.Error(), c)
		return
	}
	var usuarios UsuarioModel.Usuario
	err = c.BindJSON(&usuarios)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error en la estructura", err.Error(), c)
	}
	//inicio transacción en la base de datos
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al iniciar la transacción", tx.Error, c)
		return
	}
	err = UsuarioModel.ValidarUsuario(id, b, tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al validar usuario", err.Error(), c)
		return
	}
	err = AuthModel.ValidaCredencial(id, b, tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al actualizar los datos", err.Error(), c)
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al actualizar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos actualizados correctamente", usuarios, c)
}

//Elimina un usuario

func DeleteUsuario(c *gin.Context) {
	/* *********************    Declarando dos queryParams     ********************* */
	idUsuario := c.DefaultQuery("idUsuario", "0")
	idPersona := c.DefaultQuery("idPersona", "0")
	mensaje := c.DefaultQuery("mensaje", "0")
	if idUsuario == "" || idPersona == "" {
		Estructuras.Responder(http.StatusBadRequest, "Debes ingresar el id", nil, c)
		return
	}
	/* *********************    Inicializo la base de datos     ********************* */
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al iniciar la transacción", tx.Error, c)
		return
	}
	/* *********************    inicio el orden para eliminar los registros del usuario     ********************* */

	/* *********************    3.- Elimino de la tabla credenciales     ********************* */
	err, _ := AuthModel.EliminarCredenciales(idUsuario, tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al eliminar credenciales", err.Error(), c)
		return
	}
	/* *********************    4.- Elimino de la tabla usuarios     ********************* */
	err, _ = UsuarioModel.EliminarUsuario(idUsuario, tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al eliminar usuario", err.Error(), c)
		return
	}
	/* *********************    5.- Elimino de la tabla personas     ********************* */
	err, persona := PersonaModel.EliminarPersona(idPersona, tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al eliminar persona", err.Error(), c)
		return
	}
	/* ********************* si ninguna función de eliminar arrojo un error hago el commit ******************* */
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al eliminar los datos", err.Error(), c)
		return
	}
	err = Utilerias.EnviarCorreo(mensaje, persona.Correo, "Tu ingreso ha sido rechazado")
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al enviar el correo", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Usuario eliminado correctamente", persona, c)
	return
}

func CheckResetPassword(ctx *gin.Context) {
	dateBirth := ctx.DefaultQuery("date", "")
	email := ctx.DefaultQuery("email", "")
	err, result, data := UsuarioModel.CheckResetPassword(email, dateBirth)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error con el servidor", err.Error(), ctx)
		return
	}
	if !result {
		Estructuras.Responder(http.StatusInternalServerError, "Error al actualizar los datos", err.Error(), ctx)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", gin.H{"status": result, "data": data}, ctx)
	return
}
func ResetPasswordUser(ctx *gin.Context) {
	idUser := ctx.DefaultQuery("idUser", "")
	var password UsuarioModel.Passwords
	var usuarios UsuarioModel.Usuario
	err := ctx.BindJSON(&password)

	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error en la estructura", err.Error(), ctx)
	}
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al iniciar la transacción", tx.Error, ctx)
		return
	}
	fmt.Println(password)
	usuarios.Contraseña = password.NewPassword
	err = usuarios.ActualizarUsuario(idUser, tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al actualizar datos", err.Error(), ctx)
		return
	}
	var credenciales AuthModel.Credenciales
	credenciales.Password = password.NewPassword
	err = credenciales.ActualizarCredencial(idUser, tx)
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError,
			"Error al actualizar la contraseña de las credenciales", err.Error(), ctx)
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al cerrar la transaction", err.Error(), ctx)
		return
	}
	Estructuras.Responder(http.StatusOK, "Contraseña actualizada correctamente", nil, ctx)
}
