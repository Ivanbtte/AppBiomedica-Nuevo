package ModeloUnidadM

import "appbiomedica/Conexion"

//********************************************************************************************************
/* *********************   Funcion que regresa todas las delegaciones     ********************* */

func ConsultsDelegations(typeDelegation int) (error, *[]Delegaciones) {
	var datos []Delegaciones
	db := Conexion.GetDB()
	if typeDelegation != 0 {
		db = db.Where("tipo = ?", typeDelegation)
	}
	err := db.Order("clv_dele").Find(&datos).Error
	if err != nil {
		return err, nil
	}
	return nil, &datos

}
