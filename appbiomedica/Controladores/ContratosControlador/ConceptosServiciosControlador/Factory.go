package ConceptosServiciosControlador

import "appbiomedica/Modelos/ContratoModel/ConceptosServiciosModelo"

func Factory(tipo string) ConceptosServiciosModelo.Concepto {
	switch tipo {
	case "5f587535d7be0a74bf7db093":
		return &ConceptosServiciosModelo.TrasladoPacientes{}
	case "5f587730d7be0a790dbbbe60":
		return &ConceptosServiciosModelo.EstudiosLaboratorio{}
	case "5f323c89d7be0a421a2305ff":
		return &ConceptosServiciosModelo.EquipoMedico{}
	default:
		return nil
	}
}

