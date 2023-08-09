package DepartamentosModel

import (
	"appbiomedica/Conexion"
	"github.com/jinzhu/gorm"
)

func (o *OrganigramaUnidad) AddOrganigramaUnidad(tx *gorm.DB) error {
	return tx.Create(o).Error
}
/***** Función para consultar la unidad médica en la cual está asociado el usuario *****/

func ConsultUnitIdByUser(userId string) (error, *OrganigramaUnidad) {
	var resultado OrganigramaUnidad
	db := Conexion.GetDB()
	err := db.Joins("inner join puesto_organigrama po on organigrama_unidad.id = po.organigrama_unidad_id").
		Joins("inner join usuario_puesto up on po.id = up.puesto_organigrama_id").
		Where("up.usuario_id = ? and up.estado = true", userId).Preload("UnidadMed").First(&resultado).Error
	if err != nil {
		return err, nil
	}
	return nil, &resultado
}