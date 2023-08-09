package TrasladosControlador

import (
	"appbiomedica/Modelos/ContratoModel/ConceptosServiciosModelo"
	"appbiomedica/Modelos/DepartamentosModel"
	"appbiomedica/Modelos/ModeloUnidadM"
	"appbiomedica/Modelos/ServiciosSubrogados/TrasladoPacientesModel"
	"appbiomedica/Modulos/Estructuras"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetDestinos(ctx *gin.Context) {
	numero_contrato := ctx.DefaultQuery("numero_contrato", "")
	unidad_med := ctx.DefaultQuery("unidad_id", "")
	unidad, err := strconv.Atoi(unidad_med)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Numero de unidad incorrecto", err.Error(), ctx)
		fmt.Println(err.Error())
		return
	}
	err, destinos := ConceptosServiciosModelo.ConsultarDestinos(numero_contrato, unidad)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar destinos", err.Error(), ctx)
		fmt.Println(err.Error())

		return
	}
	Estructuras.Responder(http.StatusOK, "Destinos Consultados", destinos, ctx)
}

func GetOrigenes(ctx *gin.Context) {
	numero_contrato := ctx.DefaultQuery("numero_contrato", "")
	unidad_med := ctx.DefaultQuery("unidad_id", "")
	unidad, err := strconv.Atoi(unidad_med)
	destino := ctx.DefaultQuery("destino", "")
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Numero de unidad incorrecto", err.Error(), ctx)
		fmt.Println(err.Error())
		return
	}
	err, origenes := ConceptosServiciosModelo.ConsultarOrigenes(numero_contrato, destino, unidad)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los origenes", err.Error(), ctx)
		return
	}
	Estructuras.Responder(http.StatusOK, "Origenes consultados", origenes, ctx)
}
func GetBudget(ctx *gin.Context) {
	contractNumber := ctx.DefaultQuery("contractNumber", "")
	unitMedical := ctx.DefaultQuery("unitId", "")
	unitId, err := strconv.Atoi(unitMedical)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Numero de unidad medica incorrecto", err.Error(), ctx)
		return
	}
	err, budget := ConceptosServiciosModelo.ConsultBudget(contractNumber, unitId)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar el presupuesto", err.Error(), ctx)
		return
	}
	Estructuras.Responder(http.StatusOK, "Presupuesto", budget, ctx)
}
func GetEarn(ctx *gin.Context) {
	contractNumber := ctx.DefaultQuery("contractNumber", "")
	unitMedical := ctx.DefaultQuery("unitId", "")
	state := ctx.DefaultQuery("state", "")
	stateRequest, err := strconv.ParseBool(state)
	unitId := 0
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error estado no valido", err.Error(), ctx)
		return
	}
	if unitMedical != "" {
		unitId, err = strconv.Atoi(unitMedical)
		if err != nil {
			Estructuras.Responder(http.StatusBadRequest, "Numero de unidad medica incorrecto", err.Error(), ctx)
			return
		}
	}
	err, earn := TrasladoPacientesModel.ConsultEarn(stateRequest, unitId, contractNumber)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar el presupuesto", err.Error(), ctx)
		return
	}
	Estructuras.Responder(http.StatusOK, "Presupuesto", earn, ctx)
}
func GetTicketGoingAvailable(ctx *gin.Context) {
	contractNumber := ctx.DefaultQuery("contractNumber", "")
	unitMedical := ctx.DefaultQuery("unitId", "")
	state := ctx.DefaultQuery("state", "")
	destination := ctx.DefaultQuery("destination", "")
	typeTicket := ctx.DefaultQuery("typeTicket", "")
	stateRequest, err := strconv.ParseBool(state)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error estado no valido", err.Error(), ctx)
		return
	}
	unitId, err := strconv.Atoi(unitMedical)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Numero de unidad medica incorrecto", err.Error(), ctx)
		return
	}
	err, available := TrasladoPacientesModel.ConsultTicketGoingAvailable(destination, typeTicket, contractNumber,
		stateRequest, unitId)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar boletos disponibles", err.Error(),
			ctx)
		return
	}
	Estructuras.Responder(http.StatusOK, "Boletos Disponibles", available, ctx)
}

type AdminTraslado struct {
	Unidad  ModeloUnidadM.UnidadMed                  `json:"unidad,omitempty"`
	Datos   DepartamentosModel.PuestoOrganigramaRole `json:"datos,omitempty"`
	Usuario DepartamentosModel.UsuarioPuesto         `json:"usuario,omitempty"`
	Estado  bool                                     `json:"estado"`
}

func GetAdministradoresUnidad(ctx *gin.Context) {
	clvDele := ctx.DefaultQuery("clvDele", "")
	roleId := ctx.DefaultQuery("roleId", "")
	data := make([]AdminTraslado, 0)
	//Buscamos las unidades medicas de la delegacion.
	err, unidades := ModeloUnidadM.ConsultaUnidadesMe(clvDele)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar las unidades medicas", err.Error(), ctx)
		return
	}
	//Iteramos las unidades medicas para buscar en la tabla de puesto organigrama role si el rol que buscamos esta asignada a esa unidad medica.
	for _, row := range *unidades {
		fmt.Println(row.DenominacionUni)
		result, status, err := DepartamentosModel.ConsultarPuestoRole(row.Id, roleId)
		if err != nil {
			fmt.Println("-----------", err.Error())
			Estructuras.Responder(http.StatusInternalServerError, "Error al consultar el rol en la unidad", err.Error(), ctx)
			return
		}
		var temp AdminTraslado
		if status && result != nil {
			res ,_ := DepartamentosModel.ConsultarUsuarioPuesto(result.PuestoOrganigramaId)
			if res!= nil {
			temp.Usuario = *res
			}
			temp.Unidad = row
			temp.Datos = *result
			temp.Estado = true
			
			data = append(data, temp)
		} else {
			var vacio DepartamentosModel.PuestoOrganigramaRole
			temp.Unidad = row
			temp.Datos = vacio
			temp.Estado =false
			data = append(data, temp)
		}
	}
	Estructuras.Responder(http.StatusOK, "Administradores consultados", data, ctx)
}
