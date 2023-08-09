package ConceptosServiciosModelo

import (
	"errors"
	"fmt"
	"io"
)

//********************************************************************************************************
/* *********************    funcion para validar un archivo de excel       ********************* */
func (concepto EquipoMedico) ValidaConceptosArchivo(file io.Reader) (bool, []RespuestConceptoArchivo, error) {
	fmt.Println("estoy validando un archivo que tiene equipo medico",concepto)
	return false, nil, errors.New("Este tipo de conceptos aun no se pueden agregar: Equipo Medico")
}