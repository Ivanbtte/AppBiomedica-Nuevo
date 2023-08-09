package Sesion

import (
	"appbiomedica/Modelos/MenuModel"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"
)

//UsuarioLogeado Estrcutura de control de usuarios logueados.
type UsuarioLogeado struct {
	UsuarioId           string
	Correo              string
	Tipo                string
	Puesto              string
	Nombre              string
	AppellidoP          string
	ApellidoM           string
	Matricula           string
	PersonaId           string
	PuestoId            string
	PuestoOrganigramaId string
	MenuItem            []MenuModel.Menu
	jwt.StandardClaims
}

var key = []byte("secret")

//CreateTokenString función que crea el token de tipo string a partir de los datos de un usuario.
func CreateTokenString(idUser, correo, tipo, puesto, nombre, apeP, apeM, Matr, persona,
	puestoId string, organigramaId string, menuItem []MenuModel.Menu) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := UsuarioLogeado{}
	claims.ExpiresAt = time.Now().UTC().Add(time.Hour * time.Duration(5)).Unix()
	//claims.ExpiresAt = time.Now().UTC().Add(time.Minute * time.Duration(1)).Unix()
	claims.IssuedAt = time.Now().UTC().Unix()
	claims.UsuarioId = idUser
	claims.Correo = correo
	claims.Tipo = tipo
	claims.Puesto = puesto
	claims.Nombre = nombre
	claims.AppellidoP = apeP
	claims.ApellidoM = apeM
	claims.Matricula = Matr
	claims.PersonaId = persona
	claims.PuestoId = puestoId
	claims.PuestoOrganigramaId = organigramaId
	claims.MenuItem = menuItem
	token.Claims = claims
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//ValidaToken un token para su vigencia
func ValidaToken(r *http.Request) (bool, error) {
	header := r.Header.Get("Authorization")
	if header == "" {
		return false, errors.New("No hay cabecera  Authorization")
	}
	tokens := strings.Split(header, " ")[1]
	var usuario UsuarioLogeado
	token, err := jwt.ParseWithClaims(tokens, &usuario, func(token *jwt.Token) (interface{}, error) { return key, nil })
	if err != nil {
		return false, err
	}
	if usuario.ExpiresAt < time.Now().Unix() {
		return false, err
	}
	return token.Valid, nil
}
