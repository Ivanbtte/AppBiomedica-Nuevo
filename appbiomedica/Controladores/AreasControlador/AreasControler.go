package AreasControlador

import (
	"appbiomedica/Modelos/AreaModel"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ConsultAreas(ctx *gin.Context) {
	err, areas := AreaModel.GetAreas()
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", areas, ctx)
		return
	}
	Estructuras.Responder(http.StatusOK, "Areas Consultadas", areas, ctx)
	return
}
