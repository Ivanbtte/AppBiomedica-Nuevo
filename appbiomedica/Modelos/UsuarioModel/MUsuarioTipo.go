package UsuarioModel

import "appbiomedica/Conexion"

//********************************************************************************************************
/* *********************    Funcion para consultar los tipos de usuarios     ********************* */

func ConsultaUsuarioTipo() (error, *[]UsuarioTipo) {
	var datos []UsuarioTipo
	db := Conexion.GetDB()
	err := db.Order("id").Find(&datos).Error
	if err != nil {
		return err, nil
	}
	return nil, &datos

}
