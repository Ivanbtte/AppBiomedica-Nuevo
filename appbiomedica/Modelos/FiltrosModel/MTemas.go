package FiltrosModel

import "appbiomedica/Conexion"

//********************************************************************************************************
//-----------------------Consultar todos los temas

func ConsultaTemas() (error, *[] Tema) {
	var datos [] Tema
	db := Conexion.GetDB()
	err := db.Find(&datos).Error
	if err != nil {
		return err, nil
	}
	return nil, &datos
}
