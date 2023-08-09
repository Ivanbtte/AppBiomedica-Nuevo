package Init

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/ModeloUnidadM"
)

func InitUnidadMedica() error {
	db := Conexion.GetDB()
	//Creando tabla Region
	if !db.HasTable(&ModeloUnidadM.Region{}) {

		if err := db.CreateTable(&ModeloUnidadM.Region{}).Error; err != nil {
			return err
		}
	}
	//Creando tabla Delegaciones
	if !db.HasTable(&ModeloUnidadM.Delegaciones{}) {

		if err := db.CreateTable(&ModeloUnidadM.Delegaciones{}).Error; err != nil {
			return err
		}
	}
	//Creando tabla Catalogo nivel atencion
	if !db.HasTable(&ModeloUnidadM.CatalogoNivelAtn{}) {

		if err := db.CreateTable(&ModeloUnidadM.CatalogoNivelAtn{}).Error; err != nil {
			return err
		}
	}
	//Creando tabla tipo vialidad
	if !db.HasTable(&ModeloUnidadM.TipoVialidad{}) {

		if err := db.CreateTable(&ModeloUnidadM.TipoVialidad{}).Error; err != nil {
			return err
		}
	}
	//Creando tabla tipo asentamiento
	if !db.HasTable(&ModeloUnidadM.TipoAsentamiento{}) {

		if err := db.CreateTable(&ModeloUnidadM.TipoAsentamiento{}).Error; err != nil {
			return err
		}
	}
	//Creando tabla jurisdiccion
	if !db.HasTable(&ModeloUnidadM.Jurisdiccion{}) {

		if err := db.CreateTable(&ModeloUnidadM.Jurisdiccion{}).Error; err != nil {
			return err
		}
	}
	if !db.HasTable(&ModeloUnidadM.UnidadMed{}) {

		if err := db.CreateTable(&ModeloUnidadM.UnidadMed{}).AddForeignKey("delegaciones_clv_dele",
			"delegaciones(clv_dele)",
			"RESTRICT", "CASCADE").AddForeignKey("clv_nivel_atn_id",
			"catalogo_nivel_atn(clv_nivel_atn)",
			"RESTRICT", "CASCADE").AddForeignKey("clv_vialidad",
			"tipo_vialidad(clv_vialidad)",
			"RESTRICT", "CASCADE").AddForeignKey("clv_tipo_asent",
			"tipo_asentamiento(clv_tipo_asent)",
			"RESTRICT", "CASCADE").AddForeignKey("id_jurisdiccion",
			"jurisdiccion(id_jurisdiccion)",
			"RESTRICT", "CASCADE").Error; err != nil {
			return err
		}
	}
	return nil
}
