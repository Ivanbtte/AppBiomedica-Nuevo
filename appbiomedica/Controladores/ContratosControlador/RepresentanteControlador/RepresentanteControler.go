package RepresentanteControlador

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/ContratoModel/ContactoModelo"
	"appbiomedica/Modelos/ContratoModel/RepresentanteModel"
	"appbiomedica/Modelos/CorreoTelefonoModel"
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strings"
)

type RepresentanteContrato struct {
	NombreCompleto string
	Correos        []struct {
		Id         string
		Correo     string
		TipoCorreo string
	} `gorm:"-"`
	Telefonos []struct {
		Id           string
		Extension    string
		Telefono     string
		TipoTelefono string
	} `gorm:"-"`
}

//********************************************************************************************************
/* ****************    Controlador para consultar al representante legal de un contrato		*************** */
func GetRepresentante(c *gin.Context) {
	n_contrato := c.DefaultQuery("n_contrato", "")
	resultado, err := RepresentanteModel.ConsultarRepresentante(n_contrato)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", resultado, c)
	return
}

//********************************************************************************************************
/* *********************    Controlador para editar la informacion de un representante legal asi como los correos y
************************************numeros telefonicos asociados a un contrato     ********************* */
func EditarRepresentante(c *gin.Context) {
	num_contrato := c.DefaultQuery("num_contrato", "")
	var representante RepresentanteContrato
	var correo CorreoTelefonoModel.Correos
	var telefono CorreoTelefonoModel.Telefonos
	var correContrato ContactoModelo.CorreosContrato
	var telefonoContrato ContactoModelo.TelefonoContrato
	err := c.BindJSON(&representante)
	if err != nil {
		Estructuras.Responder(http.StatusBadRequest, "Error al castear la estructura", err.Error(), c)
		return
	}
	resultado, err := RepresentanteModel.ConsultarRepresentante(num_contrato)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al consultar los datos", err.Error(), c)
		return
	}
	/******inicio transaccion en la base de datos******/
	tx := Conexion.GetDB().Begin()
	if tx.Error != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al conectar a la base de datos", tx.Error, c)
		return
	}
	/*****  Verificar que el nombre del representante cambio o se modifio  *****/
	if strings.TrimSpace(resultado.NombreCompleto) != Estructuras.FormatoTexto(representante.NombreCompleto) {
		/* *********************     Si el nombre es diferente actualizar en la db    ********************* */
		repre := RepresentanteModel.RepresentanteLegal{
			NombreCompleto: Estructuras.FormatoTexto(representante.NombreCompleto),
		}
		if err = repre.ActualizarRepresentante(num_contrato, tx); err != nil {
			tx.Rollback()
			Estructuras.Responder(http.StatusInternalServerError, "Error al actualizar el nombre", err.Error(), c)
			return
		}
	}
	/*****  Verificar los correos   *****/
	for _, email := range representante.Correos {
		if email.Id == "" {
			/*****  Este correo es nuevo   *****/
			/*Busco el correo en la bd*/
			_, resultado := CorreoTelefonoModel.BuscarCorreo(strings.TrimSpace(email.Correo))
			/*Si lo encuentro */
			if resultado != nil {
				correContrato.CorreosId = resultado.Id
			} else {
				/*Si no lo encuentro lo agrego*/
				correo.Correo = Estructuras.FormatoTexto(email.Correo)
				correo.Id = bson.NewObjectId().Hex()
				err = correo.AgregaCorreo(tx)
				if err != nil {
					tx.Rollback()
					Estructuras.Responder(http.StatusInternalServerError, "Error al agregar el correo en la base de datos", err.Error(), c)
					return
				}
				correContrato.CorreosId = correo.Id
			}
			/*Agrego el correo al contrato*/
			correContrato.Id = bson.NewObjectId().Hex()
			correContrato.TipoCorreo = Estructuras.FormatoTexto(email.TipoCorreo)
			correContrato.ContratoNumeroContrato = num_contrato
			err = correContrato.AgregaCorreoContrato(tx)
			if err != nil {
				tx.Rollback()
				Estructuras.Responder(http.StatusInternalServerError, "Error al asignar los correos al contrato", err.Error(), c)
				return
			}
		} else {
			item, _ := ContactoModelo.ConsultarUnCorreo(email.Id)
			if strings.TrimSpace(email.Correo) != item.Correos.Correo {
				/* *********************    Actualizar el correo electronico     ********************* */
				correo.Correo = strings.TrimSpace(email.Correo)
				if err = correo.ActualizarCorreo(item.CorreosId, tx); err != nil {
					tx.Rollback()
					Estructuras.Responder(http.StatusInternalServerError, "Error al actualizar el correo",
						err.Error(), c)
					return
				}
			}
			if strings.TrimSpace(email.TipoCorreo) != item.TipoCorreo {
				/*****  Actualizar el tipo de correo  *****/
				correContrato.TipoCorreo = strings.TrimSpace(email.TipoCorreo)
				if err = correContrato.ActualizarTipoCorreo(email.Id, tx); err != nil {
					tx.Rollback()
					Estructuras.Responder(http.StatusInternalServerError, "Error al actualizar el tipo de correo",
						err.Error(), c)
					return
				}
			}
		}
	}
	/*****  Verificar los numeros de telefono *****/
	for _, tel := range representante.Telefonos {
		if tel.Id == "" {
			/*****  Agregar un nuevo telefono  *****/
			_, resultado := CorreoTelefonoModel.BuscarTelefono(strings.TrimSpace(tel.Telefono))
			if resultado != nil {
				telefonoContrato.TelefonosId = resultado.Id
			} else {
				telefono.Telefono = strings.TrimSpace(tel.Telefono)
				telefono.Id = bson.NewObjectId().Hex()
				err = telefono.AgregaTelefono(tx)
				if err != nil {
					tx.Rollback()
					Estructuras.Responder(http.StatusInternalServerError, "Error al agregar el telefono en la base de datos", err.Error(), c)
					return
				}
				telefonoContrato.TelefonosId = telefono.Id
			}
			/*Asigno el telefono al contrato*/
			telefonoContrato.Id = bson.NewObjectId().Hex()
			telefonoContrato.TipoTelefono = Estructuras.FormatoTexto(tel.TipoTelefono)
			telefonoContrato.Extension = tel.Extension
			telefonoContrato.ContratoNumeroContrato = num_contrato
			err = telefonoContrato.AgregaTelContrato(tx)
			if err != nil {
				tx.Rollback()
				Estructuras.Responder(http.StatusInternalServerError, "Error al asignar el telefono al contrato", err.Error(), c)
				return
			}
		} else {
			item, _ := ContactoModelo.ConsultarUnTelefono(tel.Id)
			if strings.TrimSpace(tel.Telefono) != item.Telefonos.Telefono {
				/* *********************    Actualizar el numero telefonico     ********************* */
				telefono.Telefono = strings.TrimSpace(tel.Telefono)
				if err = telefono.ActualizarTelefono(item.TelefonosId, tx); err != nil {
					tx.Rollback()
					Estructuras.Responder(http.StatusInternalServerError, "Error al actualizar el telefono",
						err.Error(), c)
					return
				}
			}
			if strings.TrimSpace(tel.TipoTelefono) != item.TipoTelefono ||
				strings.TrimSpace(tel.Extension) != item.Extension {
				/*****  Actualizar el tipo o la extension de un telefono de un contrato  *****/
				telefonoContrato.TipoTelefono = strings.TrimSpace(tel.TipoTelefono)
				telefonoContrato.Extension = strings.TrimSpace(tel.Extension)
				if err = telefonoContrato.ActualizarTelefonoContrato(tel.Id, tx); err != nil {
					tx.Rollback()
					Estructuras.Responder(http.StatusInternalServerError,
						"Error al actualizar el tipo o extension del telefono",
						err.Error(), c)
					return
				}
			}
		}
	}
	//********************************************************************************************************
	/* *********************   finalizar actualizacion y hacer el commit     ********************* */
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		Estructuras.Responder(http.StatusInternalServerError, "Error al realizar el commit", err.Error(), c)
		return
	}
	Estructuras.Responder(http.StatusOK, "Datos actualizados correctamente", nil, c)
	return
}
