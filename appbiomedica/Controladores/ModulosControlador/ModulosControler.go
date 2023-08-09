package ModulosControlador

import (
	"appbiomedica/Modelos/MenuModel"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetModulesByModuleParent(ctx *gin.Context) {
	menuParent := ctx.DefaultQuery("menuParent", "")
	//userId := ctx.DefaultQuery("userId", "")
	err, result := MenuModel.ConsultModuleChildByParent(menuParent)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los modulos",
			err.Error(), ctx)
		return
	}
	/*****  Consult permit by module  *****/
	//module := make([]Module, 0)
	//for _, item := range *result {
	//	_, permits := DepartamentosModel.SearchPermitByModuleByUser(userId, item.ModuloId)
	//	if permits != nil {
	//		module = append(module, Module{Menu: item, Crear: permits.Crear, Editar: permits.Editar,
	//			Visualizar: permits.Visualizar, Eliminar: permits.Eliminar})
	//	} else {
	//		module = append(module, Module{Menu: item, Crear: false, Editar: false,
	//			Visualizar: false, Eliminar: false})
	//	}
	//	fmt.Println(module)
	//}
	Estructuras.Responder(http.StatusOK, "Modulos consultados", result, ctx)
}
