package CatalogosModel

import "appbiomedica/Conexion"

//********************************************************************************************************
//-----------------------Consultar todos los cuadro basico

func ConsultaCuadroBasico() (error, *[] CuadroBasico) {
	var datos [] CuadroBasico
	db := Conexion.GetDB()
	err := db.Find(&datos).Error
	if err != nil {
		return err, nil
	}
	return nil, &datos

}
