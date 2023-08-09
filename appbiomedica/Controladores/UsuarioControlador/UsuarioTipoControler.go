package UsuarioControlador

import (
	"appbiomedica/Modelos/UsuarioModel"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"net/http"
)

//********************************************************************************************************
/* *********************    Controlador para consultar los tipos de usuarios     ********************* */
func GetUsuarioTipos(c *gin.Context) {
	err, resultado := UsuarioModel.ConsultaUsuarioTipo()
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}
