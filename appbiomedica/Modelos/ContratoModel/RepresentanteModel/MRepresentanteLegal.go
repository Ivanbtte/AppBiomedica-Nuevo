package RepresentanteModel

import (
	"appbiomedica/Conexion"
	"github.com/jinzhu/gorm"
)

//********************************************************************************************************
/* *********************   Funcion para agregar al representante legal     ********************* */
func (rl *RepresentanteLegal) AgregarRepresentanteLegal(tx *gorm.DB) error {
	return tx.Create(rl).Error
}

//********************************************************************************************************
/* *********************    Funcion para consultar representantes legales por contrato     ********************* */
func ConsultarRepresentante(n_contrato string) (*RepresentanteLegal, error) {
	var representante RepresentanteLegal
	db := Conexion.GetDB()
	err := db.Where("contrato_numero_contrato = ?", n_contrato).First(&representante).Error
	if err != nil {
		return nil, err
	}
	return &representante, nil
}

//********************************************************************************************************
/* *********************    Funcion para actualizar el nombre del representante legal     ********************* */
func (rl *RepresentanteLegal) ActualizarRepresentante(num_contrato string, tx *gorm.DB) error {
	return tx.Model(&rl).Where("contrato_numero_contrato=?", num_contrato).
		Update(RepresentanteLegal{NombreCompleto: rl.NombreCompleto}).Error
}
