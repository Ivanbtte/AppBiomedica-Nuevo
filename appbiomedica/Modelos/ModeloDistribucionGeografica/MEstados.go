package ModeloDistribucionGeografica

import "appbiomedica/Conexion"

//********************************************************************************************************
//-----------------------Consultar todos los Estados existentes

func ConsultaEstados() (error, *[] Estados) {
	var datos [] Estados
	db := Conexion.GetDB()
	err := db.Find(&datos).Error
	if err != nil {
		return err, nil
	}
	return nil, &datos

}
func ConsultaEstadosMunicipios() (error, *[] Estados) {
	var datos [] Estados
	db := Conexion.GetDB()
	err := db.Preload("Municipio").Find(&datos).Error
	if err != nil {
		return err, nil
	}
	return nil, &datos

}
