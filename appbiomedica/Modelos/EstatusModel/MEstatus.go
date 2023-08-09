package EstatusModel

import "appbiomedica/Conexion"

func ConsultaEstatus() (error, *[]Estatus) {
	db := Conexion.GetDB()
	var datos []Estatus
	err := db.Where("id IN (?)", []string{"7", "8", "9", "10", "11", "12", "13", "14", "15", "16"}).Find(&datos).Error
	if err != nil {
		return err, nil
	}
	return nil, &datos
}
