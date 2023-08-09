package ArchivosContratosControlador

import (
	"appbiomedica/Modulos/Estructuras"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

type Archivos struct {
	Nombre string
	Ruta   string
}
type Carpeta struct {
	Nombre string
	Ruta   string
}

//********************************************************************************************************
/* *********************    Funcion para buscar solo Carpetas     ********************* */
func BusquedaCarpetas(c *gin.Context) {
	carpetas := make([]Carpeta, 0)
	ruta_path := c.DefaultQuery("path", "")
	path := "ArchivosContratos"
	if ruta_path != "" {
		path += "/" + ruta_path
	}
	archivos, err := ioutil.ReadDir(path)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al acceder al directorio", path, c)
		return
	}
	for _, archivo := range archivos {
		if archivo.IsDir() {
			c := Carpeta{Nombre: archivo.Name(), Ruta: path}
			carpetas = append(carpetas, c)
		}
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", carpetas, c)
	return
}

//********************************************************************************************************
/* *********************    Funcion para buscar solo archivos dentro de una carpeta     ********************* */
func BusquedaArchivos(c *gin.Context) {
	archivos := make([]Archivos, 0)
	ruta_path := c.DefaultQuery("path", "")
	num_contrato := c.DefaultQuery("num_contrato", "")
	path := "ArchivosContratos"
	if ruta_path != "" {
		path += "/" + ruta_path
	}
	info, err := ioutil.ReadDir(path)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al acceder al directorio", path, c)
		return
	}
	for _, archivo := range info {
		if !archivo.IsDir() && strings.Contains(archivo.Name(), num_contrato) {
			a := Archivos{Nombre: archivo.Name(), Ruta: path}
			archivos = append(archivos, a)
		}
	}
	Estructuras.Responder(http.StatusOK, "Datos consultados correctamente", archivos, c)
	return
}
