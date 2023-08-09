package Init

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/ContratoModel/AdminAuxModelo"
	"appbiomedica/Modelos/ContratoModel/ConceptoContratoModelo"
	"appbiomedica/Modelos/ContratoModel/ContactoModelo"
	"appbiomedica/Modelos/ContratoModel/ContratoModelo"
	"appbiomedica/Modelos/ContratoModel/CuentasModelo"
	"appbiomedica/Modelos/ContratoModel/DistribucionModelo"
	"appbiomedica/Modelos/ContratoModel/FianzaModelo"
	"appbiomedica/Modelos/ContratoModel/PenasModelo"
	"appbiomedica/Modelos/ContratoModel/PermisosContratoModelo"
	"appbiomedica/Modelos/ContratoModel/RepresentanteModel"
	"appbiomedica/Modelos/CorreoTelefonoModel"
	"appbiomedica/Modelos/DepartamentosModel"
)

func InitContrato() error {
	db := Conexion.GetDB()
	/*****  Creando la tabla donde se almacenaran las categorias de las cuentas  *****/
	if !db.HasTable(&CuentasModelo.CategoriaCuenta{}) {
		if err := db.CreateTable(&CuentasModelo.CategoriaCuenta{}).Error; err != nil {
			return err
		}
	}
	/*****  Creando la tabla donde se almacenara el catalogo de cueentas existentes  *****/
	if !db.HasTable(&CuentasModelo.Cuentas{}) {
		if err := db.CreateTable(&CuentasModelo.Cuentas{}).
			AddForeignKey("categoria_cuenta_id", "categoria_cuenta(id)", "CASCADE", "CASCADE").
			Error; err != nil {
			return err
		}
	}

	/*****  Creando tabla para almacenar los puestos  *****/
	if !db.HasTable(&DepartamentosModel.CatalogoPuesto{}) {
		if err := db.CreateTable(&DepartamentosModel.CatalogoPuesto{}).
			AddForeignKey("usuario_tipo_id", "usuario_tipo(id)", "CASCADE", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	/*****  Creando tabla para alamacenar los permisos de los contratos  *****/
	if !db.HasTable(&PermisosContratoModelo.PermisoContrato{}) {
		if err := db.CreateTable(&PermisosContratoModelo.PermisoContrato{}).Error; err != nil {
			return err
		}
	}
	/*****  Creando la tabla para almacenar los tipos de contratos  *****/
	if !db.HasTable(&PermisosContratoModelo.TipoContrato{}) {
		if err := db.CreateTable(&PermisosContratoModelo.TipoContrato{}).Error; err != nil {
			return err
		}
	}
	/*****  Creando la tabla para almacenar los subtipos de contratos  *****/
	if !db.HasTable(&PermisosContratoModelo.SubTipo{}) {
		if err := db.CreateTable(&PermisosContratoModelo.SubTipo{}).
			AddForeignKey("tipo_contrato_id", "tipo_contrato(id)", "CASCADE", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	/*****  Tabla para relacionar los subtipos de contratos con un contrato  *****/
	if !db.HasTable(&ContratoModelo.SubTipoContrato{}) {
		if err := db.CreateTable(&ContratoModelo.SubTipoContrato{}).
			AddForeignKey("sub_tipo_id", "sub_tipo(id)", "CASCADE", "CASCADE").
			AddForeignKey("contrato_numero_contrato", "contrato(numero_contrato)", "CASCADE", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	/*****  Creando la tabla para alamecenar los permisos de los puestos para los contratos  *****/
	if !db.HasTable(&PermisosContratoModelo.TipoContratoPuesto{}) {
		if err := db.CreateTable(&PermisosContratoModelo.TipoContratoPuesto{}).
			AddForeignKey("catalogo_puesto_id", "catalogo_puesto(id)", "CASCADE", "CASCADE").
			AddForeignKey("tipo_contrato_id", "tipo_contrato(id)", "CASCADE", "CASCADE").
			AddForeignKey("permiso_contrato_id", "permiso_contrato(id)", "CASCADE", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	if !db.HasTable(&ContratoModelo.Contrato{}) {

		if err := db.CreateTable(&ContratoModelo.Contrato{}).
			AddForeignKey("proveedor_n_prov_imss", "proveedor(n_prov_imss)", "RESTRICT", "CASCADE").
			AddForeignKey("tipo_contrato_id", "tipo_contrato(id)", "CASCADE", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	if !db.HasTable(&ContratoModelo.Contrato_Delegaciones{}) {

		if err := db.CreateTable(&ContratoModelo.Contrato_Delegaciones{}).
			AddForeignKey("delegaciones_clv_dele", "delegaciones(clv_dele)", "CASCADE", "CASCADE").
			AddForeignKey("contrato_numero_contrato", "contrato(numero_contrato)", "CASCADE", "CASCADE").
			Error; err != nil {
			return err
		}
	}

	if !db.HasTable(&FianzaModelo.Fianza{}) {

		if err := db.CreateTable(&FianzaModelo.Fianza{}).
			AddForeignKey("contrato_numero_contrato", "contrato(numero_contrato)", "CASCADE", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	if !db.HasTable(&RepresentanteModel.RepresentanteLegal{}) {

		if err := db.CreateTable(&RepresentanteModel.RepresentanteLegal{}).
			AddForeignKey("contrato_numero_contrato", "contrato(numero_contrato)", "CASCADE", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	if !db.HasTable(&AdminAuxModelo.AdminAuxContrato{}) {

		if err := db.CreateTable(&AdminAuxModelo.AdminAuxContrato{}).
			AddForeignKey("persona_id", "persona(id)", "RESTRICT", "CASCADE").
			AddForeignKey("delegaciones_clv_dele", "delegaciones(clv_dele)", "RESTRICT", "CASCADE").
			AddForeignKey("contrato_numero_contrato", "contrato(numero_contrato)", "CASCADE", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	if !db.HasTable(&PenasModelo.PenasConvencionales{}) {

		if err := db.CreateTable(&PenasModelo.PenasConvencionales{}).
			AddForeignKey("contrato_numero_contrato", "contrato(numero_contrato)", "CASCADE", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	if !db.HasTable(&PenasModelo.PenasDeductivas{}) {

		if err := db.CreateTable(&PenasModelo.PenasDeductivas{}).
			AddForeignKey("contrato_numero_contrato", "contrato(numero_contrato)", "CASCADE", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	if !db.HasTable(&ConceptoContratoModelo.ConceptoContrato{}) {

		if err := db.CreateTable(&ConceptoContratoModelo.ConceptoContrato{}).
			AddForeignKey("prei_id_articulo", "prei(id_articulo)", "RESTRICT", "CASCADE").
			AddForeignKey("contrato_numero_contrato", "contrato(numero_contrato)", "RESTRICT", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	if !db.HasTable(&DistribucionModelo.DistribucionUnidadMed{}) {

		if err := db.CreateTable(&DistribucionModelo.DistribucionUnidadMed{}).
			AddForeignKey("unidad_med_clv_presupuestal", "unidad_med(clv_presupuestal)", "CASCADE", "CASCADE").
			AddForeignKey("concepto_contrato_id", "concepto_contrato(id)", "CASCADE", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	if !db.HasTable(&DistribucionModelo.SeguimientoClave{}) {

		if err := db.CreateTable(&DistribucionModelo.SeguimientoClave{}).
			AddForeignKey("distribucion_unidad_med_id", "distribucion_unidad_med(id)", "CASCADE", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	/*Creando tabla para almacenar correos electronicos*/
	if !db.HasTable(&CorreoTelefonoModel.Correos{}) {
		if err := db.CreateTable(&CorreoTelefonoModel.Correos{}).Error; err != nil {
			return err
		}
	}
	/*Creando tabla para almacenar correos electronicos y asociarlos con el representante legal*/
	if !db.HasTable(&ContactoModelo.CorreosContrato{}) {

		if err := db.CreateTable(&ContactoModelo.CorreosContrato{}).
			AddForeignKey("correos_id", "correos(id)", "CASCADE", "CASCADE").
			AddForeignKey("contrato_numero_contrato", "contrato(numero_contrato)", "CASCADE", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	/*Creando tabla para almacenar numeros telefonicos*/
	if !db.HasTable(&CorreoTelefonoModel.Telefonos{}) {

		if err := db.CreateTable(&CorreoTelefonoModel.Telefonos{}).Error; err != nil {
			return err
		}
	}
	/*Creando tabla para almacenar numeros telefonicos y asociarlos con el representante legal*/
	if !db.HasTable(&ContactoModelo.TelefonoContrato{}) {

		if err := db.CreateTable(&ContactoModelo.TelefonoContrato{}).
			AddForeignKey("telefonos_id", "telefonos(id)", "CASCADE", "CASCADE").
			AddForeignKey("contrato_numero_contrato", "contrato(numero_contrato)", "CASCADE", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	return nil
}
