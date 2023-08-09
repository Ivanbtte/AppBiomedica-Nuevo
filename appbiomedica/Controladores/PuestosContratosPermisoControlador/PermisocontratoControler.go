package PuestosContratosPermisoControlador

import (
	"appbiomedica/Modelos/ContratoModel/PermisosContratoModelo"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"net/http"
)

//********************************************************************************************************
/* ************    Funcion para consultar todos el tipo de contrato que puedee ver seguun su puesto  **************** */
func GetPuestosPermisoContrato(c *gin.Context) {
	puesto_id := c.DefaultQuery("puesto_id", "")
	err, resultado := PermisosContratoModelo.ConsultaPuestosPermisocontrato(puesto_id)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}