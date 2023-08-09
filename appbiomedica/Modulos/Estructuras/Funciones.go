package Estructuras

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"regexp"
	"strings"
)

const Texto = `^[a-zA-Z_]+(\s[a-zA-Z_]+)*$`

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("is-string", ValidaString)
	validate.RegisterValidation("Alpha2", ValidaAlpha)
}
func ValidaString(texto validator.FieldLevel) bool {
	validatext := regexp.MustCompile(`^[0-9a-zA-Z_áéíóúÁÉÍÓÚñÑºª]+(\s[0-9a-zA-Z_áéíóúÁÉÍÓÚñÑºª]+)*$`)
	if validatext.MatchString(texto.Field().String()) {
		return true
	}
	return false
}
func ValidaAlpha(texto validator.FieldLevel) bool {
	validatext := regexp.MustCompile(`^[a-zA-ZáéíóúÁÉÍÓÚñÑ]+(\s[a-zA-Z_áéíóúÁÉÍÓÚñÑ]+)*$`)
	if validatext.MatchString(texto.Field().String()) {
		return true
	}
	return false
}
func Responder(Codigo int, Mensaje string, Data interface{}, ctx *gin.Context) {
	ctx.JSON(Codigo, RespuestaGeneral{Mensaje: Mensaje, Data: Data})
	return

}
func ValidaEstructura(estructura interface{}) ([]EstructuraValidador, error) {
	errors := make([]EstructuraValidador, 0)
	err := validate.Struct(estructura)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return errors, err
		}
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Tag())
			errors = append(errors, EstructuraValidador{Campo: err.Field(), Error: err.Tag()})
		}
	}
	return errors, nil
}
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
//********************************************************************************************************
/*Funcion que elimina espacios en blanco al inicio y final de un string, ademas que le pone mayusculas a cada palabra*/

func FormatoTexto(texto string) string {
	return strings.TrimSpace(strings.Title(strings.ToLower(texto)))
}

/*****  Funcion que elimina espacios en blanco y poner todo en minuscula  *****/

func ConvertLowercase(text string) string {
	return strings.TrimSpace(strings.ToLower(text))
}

//********************************************************************************************************
/* **   Funcion que elimina espacios en blanco al inicio y al final de un string y convierte en mayusculas todo ** */
func FormatoTextoMayusculas(texto string) string {
	return strings.TrimSpace(strings.ToUpper(strings.ToLower(texto)))
}
