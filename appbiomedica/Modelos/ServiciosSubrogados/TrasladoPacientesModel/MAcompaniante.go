package TrasladoPacientesModel

import (
	"appbiomedica/Conexion"
	"errors"

	"github.com/jinzhu/gorm"
)

//********************************************************************************************************
/* *********************   Funcion para agregar datos de la autorizacion del acompaniante     ********************* */
func (a *Acompaniante) AgregarAutorizacionAcomp(tx *gorm.DB) error {
	return tx.Create(a).Error
}
//********************************************************************************************************
/* *********************    Funcion que busca los datos de la autorizacion del acompañanate     ********************* */
func BuscarDatosAcompaniante(folio string) (bool, *Acompaniante) {
	var acompaniante Acompaniante
	db := Conexion.GetDB()
	err := db.Where("solicitud_traslado_folio = ?", folio).First(&acompaniante).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		return false, nil
	}
	// repetido error
	return true, &acompaniante
}

func BuscarAcompaniantePorUnidad(unidadMed int)(*Acompaniante, error){
	var acompaniante Acompaniante
	db := Conexion.GetDB()
	err:= db.Joins("inner join solicitud_traslado st on acompaniante.solicitud_traslado_folio = st.folio").
	Where("st.unidad_med_id = ?", unidadMed).Last(&acompaniante).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err){
			return nil, errors.New("no se encontro ningun registro para esta unidad")
		}
		return nil, err
	}
	return &acompaniante, nil
}