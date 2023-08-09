package ModeloUnidadM

import (
	"appbiomedica/Conexion"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

/* *********************    Funcion que regresa todas las unidades medicas de oaxaca     ********************* */
func ConsultaUnidadesMe(delegacion string) (error, *[]UnidadMed) {
	var datos []UnidadMed
	db := Conexion.GetDB()
	err := db.Where("delegaciones_clv_dele=? and estado = ?", delegacion, true).Order("denominacion_uni").Find(&datos).Error
	if err != nil {
		return err, nil
	}
	return nil, &datos

}

/* *********************     Funcion que regresa una sola unidad medica     ********************* */
func ConsultaUnaUnidadMed(Id int) (error, *UnidadMed) {
	var dato UnidadMed
	db := Conexion.GetDB()
	err := db.Where("id = ?", Id).First(&dato).Error
	if err != nil {
		return err, nil
	}
	return nil, &dato
}

/* ******************   Funcion que regresa una sola unidad medica buscando por num_ui_prei     ******************* */
func ConsultaUnidadMed(ui_prei string) (error, *UnidadMed) {
	var dato UnidadMed
	db := Conexion.GetDB()
	err := db.Where("ui_prei = ? and estado = ?", ui_prei, true).First(&dato).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No existe esta unidad medica"), nil
		}
		return err, nil
	}
	return nil, &dato
}

/*****  Funcion que busca una unidad medica por nombrey descripcion de servicio  *****/
func ConsultaUnidadNombre(nombre, servicio string) (bool, *UnidadMed) {
	var unidad UnidadMed
	db := Conexion.GetDB()
	//"LOWER (nombre_completo) LIKE ?", "%" + strings.ToLower(nombre) + "%"
	//db = db.Where("LOWER (nombre_completo) LIKE ?", "%"+strings.ToLower(nombre)+"%")
	err := db.Where("position(trim(lower(ubicacion_denom)) in ?) > 0", nombre).
		Where("position(trim(lower(num_unidad)) in ?) > 0", nombre).
		Where("position(trim(lower(descrip_servicio)) in ?) > 0", servicio).First(&unidad).Error
	if err != nil {
		fmt.Println("-*/-/-*/-/-*", err)
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		return false, nil
	}
	return true, &unidad
}

// Funcion que busca una unidad medica por nombre y ui
func ConsultaUnidadUi(nombre, ui string) *UnidadMed {
	fmt.Println("consultando unidad -------<<<<<<>>>>><",nombre,ui)
	var unidad UnidadMed
	db := Conexion.GetDB()
	err := db.Where("ui_prei =? and position(trim(lower(descrip_servicio)) in ?) > 0", ui, nombre).First(&unidad).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			fmt.Println("errorrrr")
			return nil
		}
		fmt.Println("errorrrr")

		return nil
	}
	fmt.Println(&unidad)
	return &unidad
}
