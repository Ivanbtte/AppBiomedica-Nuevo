package main

import (
	"appbiomedica/Conexion"
	"appbiomedica/Controladores/AccederControlador"
	"appbiomedica/Controladores/ActionModuleControlador"
	"appbiomedica/Controladores/ArchivosContratosControlador"
	"appbiomedica/Controladores/AreaServiciosControlador"
	"appbiomedica/Controladores/AreasControlador"
	"appbiomedica/Controladores/AuthControlador"
	"appbiomedica/Controladores/BitacorasControlador"
	"appbiomedica/Controladores/CatalogoControlador"
	"appbiomedica/Controladores/ContratosControlador/AdministradorControlador"
	"appbiomedica/Controladores/ContratosControlador/ConceptosControlador"
	"appbiomedica/Controladores/ContratosControlador/ConceptosServiciosControlador"
	"appbiomedica/Controladores/ContratosControlador/ContactoControlador"
	"appbiomedica/Controladores/ContratosControlador/ContratoControlador"
	"appbiomedica/Controladores/ContratosControlador/ContratoDelegacionControlador"
	"appbiomedica/Controladores/ContratosControlador/DistribucionControlador"
	"appbiomedica/Controladores/ContratosControlador/FianzaControlador"
	"appbiomedica/Controladores/ContratosControlador/PenasControlador"
	"appbiomedica/Controladores/ContratosControlador/RepresentanteControlador"
	"appbiomedica/Controladores/ContratosControlador/SubTipoControlador"
	"appbiomedica/Controladores/ContratosControlador/TrasladosControlador"
	"appbiomedica/Controladores/ContratosControlador/TrasladosControlador/Pdf"
	"appbiomedica/Controladores/DepartamentoControlador"
	"appbiomedica/Controladores/DistribucionGeoControlador"
	"appbiomedica/Controladores/EstatusControlador"
	"appbiomedica/Controladores/FechaLimiteControlador"
	"appbiomedica/Controladores/FiltrosControlador"
	"appbiomedica/Controladores/FtpControlador"
	"appbiomedica/Controladores/ModulosControlador"
	"appbiomedica/Controladores/PermisoControlador"
	"appbiomedica/Controladores/PersonaControlador"
	"appbiomedica/Controladores/PresupuestoControlador"
	"appbiomedica/Controladores/ProveedoresControlador"
	"appbiomedica/Controladores/PuestoRoleControlador"
	"appbiomedica/Controladores/PuestosContratosPermisoControlador"
	"appbiomedica/Controladores/PuestosControlador"
	"appbiomedica/Controladores/ServiciosUnidadesControlador"
	"appbiomedica/Controladores/SolicitudControlador"
	"appbiomedica/Controladores/UnidadesMedicasControlador"
	"appbiomedica/Controladores/UsuarioControlador"
	"appbiomedica/Init"
	_ "appbiomedica/Modelos/DepartamentosModel"
	"appbiomedica/Modulos/Config"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "gopkg.in/mgo.v2/bson"
)

func main() {
	//########################  Inicia instancia de postgres ##################################
	err := Conexion.Init()
	if err != nil {
		fmt.Println("Error al iniciar postgres", err)
		return
	}

	//########################  Inicia tablas de postgres ##################################
	err = Init.Init()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error al inicializar tablas")
		return
	}
	//PdfModel.ExampleFpdf_SplitLines_tables()
	r := gin.Default()
	r.Use(cors.New(Config.Cors))
	// Se levanta la carpeta de Public para servir los archivos de la aplicación de Angular (una vez construida)
	//r.Use(static.Serve("/", static.LocalFile("dist", true)))
	r.Use(static.Serve("/Fonts", static.LocalFile("Fonts", true)))
	r.Use(static.Serve("/plantillas", static.LocalFile("Plantillas", true)))
	r.Use(static.Serve("/ArchivosContratos", static.LocalFile("ArchivosContratos", true)))
	r.Use(static.Serve("/SolicitudDeTraslado", static.LocalFile("SolicitudDeTraslado", true)))
	//EndPoint que no necesitan el token para consultarse
	//EndPoint Auth
	login := r.Group("login")
	{
		login.POST("", AuthControlador.Login)
	}
	areas := r.Group("areas")
	{
		areas.GET("", AreasControlador.ConsultAreas)
	}
	credenciales := r.Group("credenciales")
	{
		credenciales.PUT("/:idUsuario/reset", AuthControlador.ResetPassword)
	}

	//EndPoint Unidades medicas
	unidades := r.Group("unidades")
	{
		unidades.GET("", UnidadesMedicasControlador.GetUnidadesMedi)
		unidades.GET(":id", UnidadesMedicasControlador.GetUnaUnidad)
	}
	//Endpoint Usuarios
	usuarios := r.Group("usuarios")
	{
		usuarios.POST("", UsuarioControlador.AgregarUsuario)
		usuarios.GET("/reset", UsuarioControlador.CheckResetPassword)
		usuarios.PUT("/reset", UsuarioControlador.ResetPasswordUser)
	}
	serviciosUnidades := r.Group("servicios_unidad")
	{
		serviciosUnidades.GET("", ServiciosUnidadesControlador.GetServiciosUnidades)
	}
	/*****EndPoint de puestos  *****/
	puesto := r.Group("puesto")
	{
		puesto.GET("/final", DepartamentoControlador.GetPositionOrganization)
		puesto.GET("", PuestosControlador.GetPuestos)
	}
	/*****EndPoint para consultar delegaciones sin tokens  *****/
	delegacionesRegistro := r.Group("delegaciones/registro")
	{
		delegacionesRegistro.GET("", UnidadesMedicasControlador.GetDelegations)
	}
	/*****EndPoint para consultar departamentos sin token *****/
	departamentosSn := r.Group("departamentos/sn")
	{
		departamentosSn.GET("/prueba", DepartamentoControlador.GetDepartmentHieracy)
		departamentosSn.GET("/solo", DepartamentoControlador.GetDepartamento)
	}
	/***** EndPoint para consultar los tipos de usuarios sin token *****/
	usuarioTipoSt := r.Group("usuario_tipo/sn")
	{
		usuarioTipoSt.GET("", UsuarioControlador.GetUsuarioTipos)
	}
	person := r.Group("person")
	{
		person.GET("check/available", PersonaControlador.CheckAvailableEnrollment)
	}
	/*****  -----------------------------------------------------------------------------------------------  *****/
	/*****  -----------------------------------------------------------------------------------------------  *****/
	//middleware que válida que todas las solicitudes tengan un token y que sea válido
	r.Use(AuthControlador.Auth)
	//EndPoint Personas
	personas2 := r.Group("personas")
	{
		personas2.GET("", PersonaControlador.GetPersonas)
		personas2.GET(":id", PersonaControlador.GetUnaPersona)
		personas2.PUT(":id", PersonaControlador.UpdatePersona)
		personas2.GET("/available", PersonaControlador.CheckAvailableEnrollment)
	}
	//Endpoint PREI
	prei := r.Group("prei")
	{
		prei.GET("", CatalogoControlador.GetPrei)
		prei.GET("/clave", CatalogoControlador.GetUnaClave)
		prei.GET("/filtro", CatalogoControlador.GetGruposFiltros)
	}
	//Endpoint Prei
	sai := r.Group("sai")
	{
		sai.GET("", CatalogoControlador.GetSai)
		sai.GET("/clave", CatalogoControlador.GetClaveSai)
		sai.GET("/filtros", CatalogoControlador.GetGenericosEspecificos)
	}

	//EndPoint Permisos
	permisos := r.Group("permisos")
	{
		permisos.GET("", PermisoControlador.GetPermisos)
		permisos.POST("", PermisoControlador.PostPermiso)
		permisos.PUT(":id", PermisoControlador.UpdatePermiso)
		permisos.DELETE(":id", PermisoControlador.DeletePermiso)
	}
	//Endpoint Usuarios
	usuarios2 := r.Group("usuarios")
	{
		usuarios2.GET("", UsuarioControlador.GetUsuarios)
		usuarios2.PUT(":id", UsuarioControlador.UpdateUsuario)
		usuarios2.PUT("", UsuarioControlador.ValidaUsuario)
		usuarios2.DELETE("", UsuarioControlador.DeleteUsuario)
		usuarios2.GET("/unidad", UsuarioControlador.GetUnidadUsuario)
		usuarios2.GET("/department", UsuarioControlador.GetUserByDepartment)
		usuarios2.GET("/puesto", UsuarioControlador.GetUserByPosition)
		usuarios2.GET("/valid", UsuarioControlador.GetUserValidation)
		usuarios2.GET("tipo/modificacion", UsuarioControlador.GetTipoModificacion)
		usuarios2.PUT("unsuscribe", UsuarioControlador.UnsuscribeUsuarioPuesto)
	}

	//EndPoint Municipios
	municipios := r.Group("municipios")
	{
		municipios.GET("", DistribucionGeoControlador.GetMunicipios)
	}
	//EndPoint Estados
	estados := r.Group("estados")
	{
		estados.GET("", DistribucionGeoControlador.GetEstados)
	}
	//EndPoint Temas
	temas := r.Group("temas")
	{
		temas.GET("", FiltrosControlador.GetTemas)
	}
	//EndPoint Grupo Temas
	grupoTemas := r.Group("grupo/temas")
	{
		grupoTemas.GET("", FiltrosControlador.GetGruposTemas)

	}
	//EndPoint Cuadro Basico
	cuadroBasico := r.Group("cuadro_basico")
	{
		cuadroBasico.GET("", CatalogoControlador.GetCuadroBasico)

	}
	/* *********************    EndPoint para servicios proforma     ********************* */
	serviciosProforma := r.Group("servicios_proforma")
	{
		serviciosProforma.GET("", CatalogoControlador.GetServiciosProforma)
		serviciosProforma.GET(":id", CatalogoControlador.GetClaveNombreServ)
	}
	/* *********************    EndPoint para catalogo proforma     ********************* */
	proforma := r.Group("proforma")
	{
		proforma.GET("", CatalogoControlador.GetProforma)
		proforma.GET(":id", CatalogoControlador.GetClaveProforma)
	}
	/* *********************    EndPoint para el catalogo de solicitudes      ********************* */
	solicitudes := r.Group("solicitud")
	{
		solicitudes.POST("", SolicitudControlador.PostSolicitud)
		solicitudes.GET("", SolicitudControlador.GetTodasSolicitud)
		solicitudes.GET("/detalle", SolicitudControlador.GetUnaSolicitud)
		solicitudes.PUT("", SolicitudControlador.UpdateSolicitud)
	}
	/* *********************    EndPoint para el catalogo claves Solicitud      ********************* */
	claveSolicitud := r.Group("clave")
	{
		claveSolicitud.POST("", SolicitudControlador.PostClaveSolicitud)
		claveSolicitud.GET("", SolicitudControlador.GetClavesSolicitud)
		claveSolicitud.GET(":id", SolicitudControlador.GetClave)
		claveSolicitud.PUT("", SolicitudControlador.UpdateClave)
		claveSolicitud.DELETE(":id", SolicitudControlador.DeleteClave)
	}
	bitacoraSolicitud := r.Group("bitacora_solicitud")
	{
		bitacoraSolicitud.GET("", BitacorasControlador.GetBitacoraSolicitud)
	}
	fechaLimite := r.Group("fecha")
	{
		fechaLimite.POST("", FechaLimiteControlador.PostFecha)
		fechaLimite.GET("", FechaLimiteControlador.GetFecha)
		fechaLimite.PUT("", FechaLimiteControlador.UpdateFecha)
	}
	presupuesto := r.Group("presupuesto")
	{
		presupuesto.GET("", PresupuestoControlador.GetPresupuesto)
		presupuesto.POST("", PresupuestoControlador.PostPresupuesto)
		presupuesto.PUT("", PresupuestoControlador.UpdatePresupuesto)
	}
	estatus := r.Group("estatus")
	{
		estatus.GET("", EstatusControlador.GetEstatus)
	}
	areaServicios := r.Group("area_servicio")
	{
		areaServicios.GET("", AreaServiciosControlador.GetAreasServicios)
		areaServicios.POST("", AreaServiciosControlador.PostAreaServicio)
	}
	//********************************************************************************************************
	/* *********************    EndPoint para los proveedores     ********************* */
	proveedores := r.Group("proveedores")
	{
		proveedores.GET("", ProveedoresControlador.GetProveedores)
		proveedores.GET("/nombre", ProveedoresControlador.GetProveedoresNombre)
		proveedores.POST("", ProveedoresControlador.PostProveedor)
		proveedores.PUT("", ProveedoresControlador.UpdateProveedor)
		proveedores.DELETE("", ProveedoresControlador.DeleteProveedor)
	}
	//********************************************************************************************************
	/* *********************    EndPoint para los contactos de los proveedores     ********************* */
	contactoProveedores := r.Group("proveedores/contacto")
	{
		contactoProveedores.GET("", ProveedoresControlador.GetContactoProveedores)
		contactoProveedores.POST("", ProveedoresControlador.PostContactoProveedor)
		contactoProveedores.PUT("", ProveedoresControlador.UpdateContactoProveedor)
		contactoProveedores.DELETE("", ProveedoresControlador.DeleteContactoProveedor)
	}
	//********************************************************************************************************
	/* *********************    EndPoint para los contratos     ********************* */
	contrato := r.Group("contrato")
	{
		contrato.GET("", ContratoControlador.GetContratos)
		contrato.GET("/proveedores", ContratoControlador.GetProveedoresContrato)
		contrato.GET("/traslados", ContratoControlador.GetContratosTraslado)
		contrato.GET("/subTipo", ContratoControlador.GetContratosTrasladoSubtipo)
		contrato.GET("/promedio/mayor/acompaniante", ContratoControlador.GetPorcentajeMayorAcompaniante)
		contrato.GET("/detalle/contrato", ContratoControlador.GetDetallesContrato)
		contrato.POST("", ContratoControlador.PostContrato)
		contrato.POST("/solo", ContratoControlador.PostUnContrato)
		contrato.POST("/verificar", ContratoControlador.VerificarContratoFile)
		contrato.POST("/subir", ContratoControlador.PostContratoFile)
		contrato.POST("/guardar", ContratoControlador.GuardarArchivoContrato)
		contrato.PUT("", ContratoControlador.UpdateContrato)
		contrato.DELETE("", ContratoControlador.DeleteContrato)
	}
	conceptosContrato := r.Group("contrato/conceptos")
	{
		conceptosContrato.POST("", ConceptosControlador.PostConceptoContrato)
		conceptosContrato.POST("/verificar", ConceptosControlador.ValidarConceptosArchivo)
		conceptosContrato.POST("/agregar", ConceptosControlador.AgreegarConceptosArchivo)
		conceptosContrato.GET("", ConceptosControlador.GetConceptos)
		conceptosContrato.PUT("", ConceptosControlador.UpdateConcepto)
		conceptosContrato.DELETE("", ConceptosControlador.DeleteConcepto)
	}
	concServContrato := r.Group("contrato/concepto")
	{
		concServContrato.POST("/validar", ConceptosServiciosControlador.PostConceptoServicioContrato)
		concServContrato.POST("", ConceptosServiciosControlador.AgregarConceptosArchivo)
		concServContrato.GET("", ConceptosServiciosControlador.GetConceptos)
	}
	delegaciones := r.Group("delegaciones")
	{
		delegaciones.GET("", UnidadesMedicasControlador.GetDelegations)
	}
	distribucionPorUnidad := r.Group("distribucion")
	{
		distribucionPorUnidad.GET("", DistribucionControlador.GetDistribucion)
		distribucionPorUnidad.POST("/verificar", DistribucionControlador.ValidarArchivoDistribucion)
		distribucionPorUnidad.POST("/agregar", DistribucionControlador.AgregarArchivoDistribucion)
		distribucionPorUnidad.POST("", DistribucionControlador.AgregaDistribucion)
		distribucionPorUnidad.DELETE("", DistribucionControlador.EliminaDistribucion)
		distribucionPorUnidad.PUT("", DistribucionControlador.EditarDistribucion)
	}
	contratoDeleg := r.Group("contrato/delegacion")
	{
		contratoDeleg.GET("", ContratoDelegacionControlador.GetContratoDeleg)
	}
	representante := r.Group("representante")
	{
		representante.GET("", RepresentanteControlador.GetRepresentante)
		representante.PUT("", RepresentanteControlador.EditarRepresentante)
	}
	correosContrato := r.Group("correos_contrato")
	{
		correosContrato.GET("", ContactoControlador.GetCorreosContrato)
		correosContrato.DELETE("", ContactoControlador.DeleteCorreoContrato)
	}
	telefonosContrato := r.Group("telefonos_contrato")
	{
		telefonosContrato.GET("", ContactoControlador.GetTelefonoContrato)
		telefonosContrato.DELETE("", ContactoControlador.DeleteTelefonoContrato)
	}
	penaDeductiva := r.Group("pena_deductiva")
	{
		penaDeductiva.GET("", PenasControlador.GetPenasDeductivas)
	}
	penaConvencional := r.Group("pena_convencional")
	{
		penaConvencional.GET("", PenasControlador.GetPenasConvencionales)
	}
	administradorAux := r.Group("administrador_aux")
	{
		administradorAux.GET("", AdministradorControlador.GetAdministradores)
	}
	fianza := r.Group("fianza")
	{
		fianza.GET("", FianzaControlador.GetFianza)
		fianza.POST("", FianzaControlador.PostFianza)
	}
	/***** EndPoint para consultar los tipos de contrato segun el puesto *****/
	tipoContratoPuesto := r.Group("contrato_puesto")
	{
		tipoContratoPuesto.GET("", PuestosContratosPermisoControlador.GetPuestosPermisoContrato)
	}
	departamentos := r.Group("departamentos")
	{
		departamentos.GET("/solo", DepartamentoControlador.GetDepartamento)
	}
	archivosContratos := r.Group("archivos/contrato")
	{
		archivosContratos.GET("", ArchivosContratosControlador.BusquedaCarpetas)
		archivosContratos.GET("/archivos", ArchivosContratosControlador.BusquedaArchivos)
	}
	/***** EndPoint relacionado con los subtipos de contratos *****/
	subTipo := r.Group("sub_tipo")
	{
		subTipo.GET("", SubTipoControlador.GetSubTipo)
	}
	/***** EndPoint relacionado con los subtipos que contiene un contrato *****/
	subTipoContrato := r.Group("sub_tipo_contrato")
	{
		subTipoContrato.GET("", SubTipoControlador.GetSubTipoContrato)
	}
	/***** EndPoint para los folios del ftp-01  *****/
	ftp := r.Group("ftp")
	{
		ftp.GET("", FtpControlador.GetFolioFtp)
		ftp.POST("", FtpControlador.AgregarFoliosTP)
	}
	/***** EndPoint para las firmas de las solicitudes *****/
	firma := r.Group("firma")
	{
		firma.POST("", TrasladosControlador.PostFirmaSolicitud)
		firma.GET("", TrasladosControlador.GetFirmaSolicitud)
		firma.GET("/all", TrasladosControlador.GetFirmsByUnit)
		firma.PUT("/enable", TrasladosControlador.EnableFirm)
		firma.PUT("", TrasladosControlador.PutFirm)
	}
	traslados := r.Group("traslados")
	{
		traslados.GET("destinos", TrasladosControlador.GetDestinos)
		traslados.GET("origenes", TrasladosControlador.GetOrigenes)
		traslados.GET("pdf/plantilla", Pdf.GetPdfNoDates)
		traslados.GET("pdf/datos", Pdf.GetPdfSoloDatos)
		traslados.GET("pdf/completo", Pdf.GetPdfCompleto)
		traslados.POST("/solicitud", TrasladosControlador.PostSolicitud)
		traslados.GET("/solicitados", TrasladosControlador.GetRequestsMade)
		traslados.GET("/folio", TrasladosControlador.GetFolioSolicitud)
		traslados.PUT("/cancelar", TrasladosControlador.CancelFolioSolicitud)
		traslados.GET("/filter/origin", TrasladosControlador.GetOriginFilter)
		traslados.GET("/filter/destination", TrasladosControlador.GetDestinationFilter)
		traslados.GET("/filter/companion", TrasladosControlador.GetCompanionFilter)
		traslados.GET("/budget", TrasladosControlador.GetBudget)
		traslados.GET("/earn", TrasladosControlador.GetEarn)
		traslados.GET("/ticketGoing", TrasladosControlador.GetTicketGoingAvailable)
		traslados.GET("administradores", TrasladosControlador.GetAdministradoresUnidad)
	}
	menu := r.Group("menu")
	{
		menu.GET("module", ModulosControlador.GetModulesByModuleParent)
	}
	usuarioPuesto := r.Group("usuarioPuesto")
	{
		usuarioPuesto.POST("", UsuarioControlador.AddNewUsuarioPuesto)
		usuarioPuesto.POST("bitacora", UsuarioControlador.ValidarUsuarioPuesto)
		usuarioPuesto.POST("role", UsuarioControlador.AddNewRolePuesto)
	}
	acompaniante := r.Group("acompaniante")
	{
		acompaniante.GET("", TrasladosControlador.GetAcompaniantePorUnidad)
	}
	actionsModule := r.Group("acciones")
	{
		actionsModule.GET("", ActionModuleControlador.GetActionsByUnit)
	}
	puestoRole := r.Group("puesto/role")
	{
		puestoRole.POST("", PuestoRoleControlador.PostPuestoRole)
	}
	acceder := r.Group("acceder")
	{
		acceder.GET("", AccederControlador.GetDatosAcceder)
	}

	/*	module := r.Group("module")
		{
			module.GET("actions", PermitUserFromModule.GetPermitsByUser)
		}*/

	//r.Run("192.168.1.64")
	//********************************************************************************************************

	//********************************************************************************************************
	/* *********************    para modo pruebas     ********************* */
	/*****    *****/

	// tx := Conexion.GetDB().Begin()
	// if tx.Error != nil {
	// 	return
	// }
	// role := DepartamentosModel.PuestoOrganigrama{
	// 	Id:               bson.NewObjectId().Hex(),
	// 	PuestoJefe: true,
	// 	NumeroPlazas: 1,
	// 	OrganigramaUnidadId: "6287e6e3d7be0a6501874fec",
	// 	CatalogoPuestoId: "62fd5b7feaacdd6bc3982b17",
	// 	Estado: true,

	// }
	// err = role.AddPositionOrganization(tx)
	// if err != nil {
	// 	tx.Rollback()
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// err = tx.Commit().Error
	// if err != nil {
	// 	tx.Rollback()
	// 	fmt.Println(err.Error())
	// 	return
	// }
	//fmt.Print(bson.NewObjectId().Hex())
	/*	err = Utilerias.EnviarCorreo("Enviando correo de prueba", "rmoncada10@gmail.com", "Prueba")
		if err != nil {
			fmt.Println(err)
		}*/
	/* *********************    este es para construir ya bien la app de go     ********************* */
	//Para modo prueba
	//r.Run("172.25.251.100:4200")
	//Para modo produccioncdzxxzaxcxxzx bcv cfn hjku-+
	//r.Run("172.25.251.100:80")
	// f, err := os.OpenFile("CatalogoEspecialidad.xlsx", os.O_RDWR, 0644)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// resultado, errores, err := ServiciosSubrogadosControlador.ValidarEstudiosDescripcion(f)
	// if err != nil {
	// 	fmt.Println(errores, err)
	// 	return
	// }
	// if !resultado {
	// 	fmt.Println("Errores de validacion", errores)
	// 	return
	// }
	// defer f.Close()
	// fa, err := os.OpenFile("CatalogoEspecialidad.xlsx", os.O_RDWR, 0644)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// err = ServiciosSubrogadosControlador.AgregarServicioDesc(fa)
	// if err != nil{
	// 	fmt.Println("Error al registar los estudios", err)
	// 	return
	// }

	//Funcion para poner los ceros a la izquierda del numero de proveedor
	// err, provedores, _ := ProveedoresModel.ConsultarProveedor("", []string{}, 1, 300)
	// if err != nil {
	// 	fmt.Println("eroor", err)
	// 	return
	// }
	// tx := Conexion.GetDB().Begin()
	// if tx.Error != nil {
	// 	fmt.Println("Error al iniciar la transaccion")
	// 	return
	// }
	// for _, row := range *provedores {
	// 	if len(row.NProvImss) != 10 {
	// 		num,_ := strconv.Atoi(row.NProvImss)
	// 		tmp := fmt.Sprintf("%010d", num)
	// 		row.NProvImss = tmp
	// 		fmt.Println(row.NProvImss)
	// 		err = row.ActualizarProveedor(row.Id, tx)
	// 		if err != nil {
	// 			tx.Rollback()
	// 			fmt.Println("Error al actualizar", err)
	// 			return
	// 		}
	// 	}
	// }
	// err = tx.Commit().Error
	// if err != nil {
	// 	tx.Rollback()
	// 	fmt.Println("Error al realizar el commit")
	// 	return
	// }
	r.Run()
}
