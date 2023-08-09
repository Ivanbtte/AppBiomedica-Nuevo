package Init

import "appbiomedica/Conexion"

func Init() error {
	var db = Conexion.GetDB()
	db.LogMode(true)
	db.SingularTable(true)
	if err := InitMenu(); err != nil {
		return err
	}
	if err := InitDistGeog(); err != nil {
		return err
	}
	if err := InitUnidadMedica(); err != nil {
		return err
	}
	if err := InitDepartamentos(); err != nil {
		return err
	}
	if err := InitUsuarios(); err != nil {
		return err
	}
	if err := InitCatalogos(); err != nil {
		return err
	}
	if err := InitSolicitud(); err != nil {
		return err
	}
	if err := InitConfig(); err != nil {
		return err
	}
	if err := InitProveedor(); err != nil {
		return err
	}
	if err := InitContrato(); err != nil {
		return err
	}
	if err := InitTrasladoPacientes(); err != nil {
		return err
	}
	if err := InitConceptos(); err != nil {
		return err
	}
	return nil
}
