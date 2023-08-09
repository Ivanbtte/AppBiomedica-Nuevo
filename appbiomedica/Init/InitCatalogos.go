package Init

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/AreaServiciosModel"
	"appbiomedica/Modelos/CatalogosModel"
	"appbiomedica/Modelos/FiltrosModel"
	"appbiomedica/Modelos/ServiciosUnidadesModel"
)

func InitCatalogos() error {
	db := Conexion.GetDB()

	//Creando tabla de catalogo_unidades_medidas
	if !db.HasTable(&CatalogosModel.UnidadesMedidas{}) {

		if err := db.CreateTable(&CatalogosModel.UnidadesMedidas{}).Error; err != nil {
			return err
		}
	}
	//Creando tabla de Catalogo_cuadro_basico
	if !db.HasTable(&CatalogosModel.CuadroBasico{}) {

		if err := db.CreateTable(&CatalogosModel.CuadroBasico{}).Error; err != nil {
			return err
		}
	}
	//Creando tabla de catalogo_prei
	if !db.HasTable(&CatalogosModel.Prei{}) {

		if err := db.CreateTable(&CatalogosModel.Prei{}).AddForeignKey("unidades_medidas_id", "unidades_medidas(id)",
			"RESTRICT", "CASCADE").AddForeignKey("cuadro_basico_id", "cuadro_basico(id)",
			"RESTRICT", "CASCADE").Error; err != nil {
			return err
		}
	}
	//Creando tabla de temas
	if !db.HasTable(&FiltrosModel.Tema{}) {

		if err := db.CreateTable(&FiltrosModel.Tema{}).Error; err != nil {
			return err
		}
	}
	//Creando tabla de grupos_temas
	if !db.HasTable(&FiltrosModel.GrupoTema{}) {

		if err := db.CreateTable(&FiltrosModel.GrupoTema{}).AddForeignKey("tema_id", "tema(id)",
			"RESTRICT", "CASCADE").Error; err != nil {
			return err
		}
	}
	//Creando tabla de descripcion_grupo
	if !db.HasTable(&CatalogosModel.DescripcionGrupo{}) {

		if err := db.CreateTable(&CatalogosModel.DescripcionGrupo{}).Error; err != nil {
			return err
		}
	}
	//Creando tabla de unidad_presentacion
	if !db.HasTable(&CatalogosModel.UnidadPresentacion{}) {

		if err := db.CreateTable(&CatalogosModel.UnidadPresentacion{}).Error; err != nil {
			return err
		}
	}
	//Creando tabla de tipo_presentacion
	if !db.HasTable(&CatalogosModel.TipoPresentacion{}) {

		if err := db.CreateTable(&CatalogosModel.TipoPresentacion{}).Error; err != nil {
			return err
		}
	}
	//Creando tabla de sai
	if !db.HasTable(&CatalogosModel.Sai{}) {

		if err := db.CreateTable(&CatalogosModel.Sai{}).AddForeignKey("grupo", "grupo_tema(grupo)",
			"RESTRICT", "CASCADE").AddForeignKey("descripcion_grupo_id", "descripcion_grupo(id)",
			"RESTRICT", "CASCADE").AddForeignKey("unidad_presentacion_id", "unidad_presentacion(id)",
			"RESTRICT", "CASCADE").AddForeignKey("tipo_presentacion_id", "tipo_presentacion(id)",
			"RESTRICT", "CASCADE").Error; err != nil {
			return err
		}
	}
	//Creando tabla de servicios_proforma
	if !db.HasTable(&CatalogosModel.ServiciosProforma{}) {

		if err := db.CreateTable(&CatalogosModel.ServiciosProforma{}).Error; err != nil {
			return err
		}
	}
	//Creando tabla de proforma
	if !db.HasTable(&CatalogosModel.Proforma{}) {

		if err := db.CreateTable(&CatalogosModel.Proforma{}).AddForeignKey("id_servicio", "servicios_proforma(id)",
			"RESTRICT", "CASCADE").AddForeignKey("prei_id_articulo", "prei(id_articulo)",
			"RESTRICT", "CASCADE").Error; err != nil {
			return err
		}
	}
	/* *********************    creando tabla de los servicios por unidad     ********************* */
	if !db.HasTable(&ServiciosUnidadesModel.ServiciosPorUnidad{}) {

		if err := db.CreateTable(&ServiciosUnidadesModel.ServiciosPorUnidad{}).
			AddForeignKey("id_unidad_med", "unidad_med(id)", "RESTRICT", "CASCADE").
			AddForeignKey("id_servicios_proforma", "servicios_proforma(id)", "RESTRICT", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	/* *********************    creando tabla para las areas de los servicios por unidades medicas     ********************* */
	if !db.HasTable(&AreaServiciosModel.AreaServicioUnidad{}) {

		if err := db.CreateTable(&AreaServiciosModel.AreaServicioUnidad{}).
			AddForeignKey("unidad_med_id", "unidad_med(id)", "RESTRICT", "CASCADE").
			AddForeignKey("servicios_proforma_id", "servicios_proforma(id)", "RESTRICT", "CASCADE").
			Error; err != nil {
			return err
		}
	}
	return nil
}
