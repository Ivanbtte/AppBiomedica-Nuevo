package AccederControlador

import (
	"appbiomedica/Modulos/Estructuras"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/gin-gonic/gin"
)

type RespuestaAcceder struct {
	UsrMsg string
	Stat   bool
	Data   interface{}
}

func GetDatosAcceder(ctx *gin.Context) {
	nss := ctx.DefaultQuery("nss", "")
	clienteHttp := &http.Client{}
	url := "http://11.1.21.70/scace/Clases/getAccederData.php?act=getAccederInfo&mode=basic&storeInSes=1"
	peticion, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// Maneja el error de acuerdo a tu situación
		Estructuras.Responder(http.StatusInternalServerError, "error creando peticion", err, ctx)
		return 
	}
	params := peticion.URL.Query()
	params.Add("NSS", nss)
	peticion.URL.RawQuery = params.Encode()
	respuesta, err := clienteHttp.Do(peticion)
	if err != nil {
		// Maneja el error de acuerdo a tu situación
		fmt.Println(err,respuesta)
		Estructuras.Responder(http.StatusInternalServerError, "error haciendo peticion, no pudimos conectar con el servidor", err, ctx)
		return
	}
	// No olvides cerrar el cuerpo al terminar
	defer respuesta.Body.Close()

	cuerpoRespuesta, err := ioutil.ReadAll(respuesta.Body)
	if err != nil {
		Estructuras.Responder(http.StatusInternalServerError, "error leyendo respuesta", err, ctx)
		return
	}

	respuestaString := string(cuerpoRespuesta)
	data := RespuestaAcceder{}
	json.Unmarshal([]byte(respuestaString), &data)
	fmt.Println(data.Data)
	// if err != nil {
	// 	Estructuras.Responder(http.StatusInternalServerError, "error codificando respuesta", err, ctx)
	// 	return
	// }
	// Aquí puedes decodificar la respuesta si es un JSON, o convertirla a cadena
	Estructuras.Responder(http.StatusOK, "datos consultados correctamente", data, ctx)
}
