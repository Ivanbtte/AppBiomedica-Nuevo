package AreaServiciosControlador

import (
	"appbiomedica/Modelos/AreaServiciosModel"
	"appbiomedica/Modulos/Estructuras"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

//*****************************************************************************************************************
/* *********************
    Controlador para consultar todas las areas de un servicio y unidad medica
	Recibe como query params el id de la unidad y el id del servicio y regresa un arreglo
********************* */
func GetAreasServicios(c *gin.Context) {
	id_unidad := c.DefaultQuery("id_unidad", "")
	id_servicio := c.DefaultQuery("id_servicio", "")
	id_usuario := c.DefaultQuery("id_usuario", "")
	resultado, err := AreaServiciosModel.ConsultarAreasServicios(id_unidad, id_servicio, id_usuario)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", resultado, c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}

//*****************************************************************************************************************
/* *********************    Controlador para agregar una nueva area por servicio y unidad medica       ********************* */
func PostAreaServicio(c *gin.Context) {
	var area AreaServiciosModel.AreaServicioUnidad
	err := c.BindJSON(&area)
	fmt.Print(area)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error al castear la estructura", err.Error(), c)
		return
	}

	area.Id = bson.NewObjectId().Hex()
	area.Estado = true
	err = area.AgregarAreaServicio()
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al insertar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos agregados correctamente", area, c)
	return
}
