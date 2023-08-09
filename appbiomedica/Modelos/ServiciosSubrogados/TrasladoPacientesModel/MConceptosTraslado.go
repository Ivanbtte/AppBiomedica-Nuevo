package TrasladoPacientesModel

import (
	"appbiomedica/Conexion"
	"github.com/jinzhu/gorm"
)

/* *********************   Funcion para agregar un boleto de un traslado     ********************* */
func (b *Boletos) AgregarBoletosTraslado(tx *gorm.DB) error {
	return tx.Create(b).Error
}
func BuscarBoletosTraslado(folio string) (bool, *[]Boletos) {
	var boleto []Boletos
	db := Conexion.GetDB()
	err := db.Where("solicitud_traslado_folio = ?", folio).Preload("TipoPasajero").Find(
		&boleto).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		return false, nil
	}
	return true, &boleto
}

