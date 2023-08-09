package AdminAuxModelo

import (
	"appbiomedica/Conexion"
	"github.com/jinzhu/gorm"
)

func (admin *AdminAuxContrato) AgregarAdminAux(tx *gorm.DB) error {
	return tx.Create(admin).Error
}

//********************************************************************************************************
/* **************    Funcion para consultar administradores o auxiliares buscandolo por numero de contrato
************** */
func ConsultarAdministrador(numero_contrato string) (error, *[]AdminAuxContrato) {
	var datos []AdminAuxContrato
	db := Conexion.GetDB()
	err := db.Where("contrato_numero_contrato = ?", numero_contrato).Find(&datos).Error
	if err != nil {
		return err, nil
	}
	return nil, &datos
}