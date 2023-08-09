package CatalogoControlador

import (
	"appbiomedica/Modelos/CatalogosModel"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"net/http"
)

//*****************************************************************************************************************
//consulta todos los Cuadro Basico
func GetCuadroBasico(c *gin.Context) {
	err, resultado := CatalogosModel.ConsultaCuadroBasico()
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}