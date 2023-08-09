package PuestosControlador

import (
	"appbiomedica/Modelos/DepartamentosModel"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"net/http"
)

//********************************************************************************************************
/* *********************    Funcion para consultar todos los puestos      ********************* */
func GetPuestos(c *gin.Context) {
	resultado, err := DepartamentosModel.ConsultaPuestos()
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}