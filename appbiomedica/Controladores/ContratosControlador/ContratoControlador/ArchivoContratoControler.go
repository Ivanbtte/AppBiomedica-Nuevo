package ContratoControlador

import (
	"appbiomedica/Modulos/Estructuras"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

//********************************************************************************************************
/* *********************    Funcion para guardar los contratos en pdf     ********************* */
func GuardarArchivoContrato(c *gin.Context) {
	num_contrato := c.DefaultQuery("num_contrato", "")
	file, _ := c.FormFile("file")
	f, err := file.Open()
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al leer el archivo", err, c)
		return
	}
	file.Filename = num_contrato
	defer f.Close()
	fmt.Println(file.Header.Get("Content-Type"))
	data, err := ioutil.ReadAll(f)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "Error al procesar el archivo", err, c)
		return
	}
	err = ioutil.WriteFile("./ArchivosContratos/Contratos/"+num_contrato+".pdf", data, 0666)
	if err != nil {
		fmt.Println(err.Error())
		Estructuras.Responder(http.StatusInternalServerError, "Error al guardar el archivo", err, c)

		return
	}
}
