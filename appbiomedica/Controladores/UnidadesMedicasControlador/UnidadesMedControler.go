package UnidadesMedicasControlador

import (
	"appbiomedica/Modelos/ModeloUnidadM"
	"appbiomedica/Modulos/Estructuras"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//*****************************************************************************************************************
//consulta todos las unidades medicas
func GetUnidadesMedi(c *gin.Context) {
	delegacion:= c.DefaultQuery("delegacion","")
	err, resultado := ModeloUnidadM.ConsultaUnidadesMe(delegacion)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	fmt.Print(resultado)
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}

/* *********************    Consulta una unidad medica     ********************* */
func GetUnaUnidad(c *gin.Context)  {
	id := c.Param("id")
	Id, err := strconv.Atoi(id)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Id del grupo incorrecto", err.Error(), c)
		return
	}
	err, unidad := ModeloUnidadM.ConsultaUnaUnidadMed(Id)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", unidad, c)
	return
}