package ModeloDistribucionGeografica

import "appbiomedica/Conexion"

//********************************************************************************************************
//-----------------------Consultar todos los municipios existentes

func ConsultaMunicipios() (error, *[] Municipios) {
	var datos [] Municipios
	db := Conexion.GetDB()
	err := db.Find(&datos).Error
	if err != nil {
		return err, nil
	}
	return nil, &datos

}
