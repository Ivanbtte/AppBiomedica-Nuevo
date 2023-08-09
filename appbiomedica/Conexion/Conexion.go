package Conexion

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// db variable global de conexión a Base de Datos en Sql
var db *gorm.DB

func Init() error {
	//***************************conexion servidor pruebaaaaa
	// conn, err := gorm.Open("postgres", "host=localhost port=5432 user=dbbiomedica dbname=biomedicaPrueba "+
	// "password=serverSCC2020 sslmode=disable")
	//***************************conexion servidor
	conn, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=biomedica "+
		"password=guille sslmode=disable")
	//**************************conexion pc oficina
	// conn, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=biomedica "+
	// 	"password=123qwe123 sslmode=disable")
	//********************************conexion lap
	//conn, err := gorm.Open("postgres", "host=localhost port=5433 user=postgres dbname=biomedica "+
	//	"password=rogelioSCC2610 sslmode=disable")
	if err != nil {
		return err
	}
	db = conn
	return nil
}

func GetDB() *gorm.DB {
	return db
}
