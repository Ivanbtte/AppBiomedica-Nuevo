package Sesion

type Permits struct {
	Permiso string `json:"permiso,omitempty"`
}
type SubChildren struct {
	State  string `json:"state,omitempty"`
	Name   string `json:"name,omitempty"`
	Type   string `json:"type,omitempty"`
	Permit string `json:"permit,omitempty"`
}
type ChildrenItems struct {
	State  string        `json:"state,omitempty"`
	Name   string        `json:"name,omitempty"`
	Type   string        `json:"type,omitempty"`
	Permit string        `json:"permit,omitempty"`
	Child  []SubChildren `json:"child,omitempty"`
}
type Menu struct {
	State     string          `json:"state,omitempty"`
	Name      string          `json:"name,omitempty"`
	Type      string          `json:"type,omitempty"`
	Icon      string          `json:"icon,omitempty"`
	Permit    string          `json:"permit,omitempty"`
	Children  []ChildrenItems `json:"children,omitempty"`
}

var MENUITEMS = []Menu{
	{
		State:  "inicio",
		Name:   "Inicio",
		Type:   "link",
		Icon:   "apps",
		Permit: "Consulta",
	},
	{
		State:  "catalogos",
		Name:   "Catálogos",
		Type:   "sub",
		Icon:   "assignment",
		Permit: "Consulta",
		Children: []ChildrenItems{
			{State: "sai", Name: "SAI", Type: "link"},
			{State: "prei", Name: "PREI", Type: "link"},
			{State: "proforma", Name: "PROFORMA", Type: "link"},
		},
	},
	{
		State:  "listas",
		Name:   "Sitios de Interés",
		Type:   "link",
		Icon:   "link",
		Permit: "Consulta",
	},
	{
		State:  "area",
		Name:   "Areas Servicios",
		Type:   "sub",
		Icon:   "business",
		Permit: "Enfermeras",
		Children: []ChildrenItems{
			{State: "nueva", Name: "Agregar", Type: "link"},
		},
	},
	{
		State:  "planeacion",
		Name:   "Fondo Fijo",
		Type:   "sub",
		Icon:   "today",
		Permit: "Enfermeras",
		Children: []ChildrenItems{
			{State: "enfermeria", Name: "Enfermeria", Type: "subchild", Child: []SubChildren{
				{State: "cuadrosi", Name: "En Cuadro Basico", Type: "link"},
				{State: "cuadrono", Name: "Fuera de Cuadro Basico", Type: "link"},
			},
			},
		},
	},
	{
		State:  "solicitudes",
		Name:   "Solicitudes",
		Type:   "sub",
		Icon:   "description",
		Permit: "Solicitud",
		Children: []ChildrenItems{
			{State: "nueva", Name: "Nueva", Type: "subchild", Child: []SubChildren{
				{State: "consumo", Name: "Bienes de Consumo", Type: "link"},},
			},
			{State: "seguimiento", Name: "Seguimiento", Type: "link"},
			{State: "editar", Name: "Editar", Type: "link"},
		},
	},
	{
		State:  "configuraciones",
		Name:   "Configuraciones",
		Type:   "sub",
		Icon:   "build",
		Permit: "Jefe Biomedico",
		Children: []ChildrenItems{
			{State: "fecha", Name: "Fecha Límite", Type: "link"},
			{State: "presupuesto", Name: "Presupuesto", Type: "link"},
		},
	},
	{
		State:  "usuarios",
		Name:   "Usuarios",
		Type:   "sub",
		Icon:   "supervised_user_circle",
		Permit: "Jefe Biomedico",
		Children: []ChildrenItems{
			{State: "validar", Name: "Validar Usuarios", Type: "link"},
			{State: "bloquear", Name: "Bloquear", Type: "link"},
		},
	},
	{
		State:  "requerimientos",
		Name:   "Validar Solicitudes",
		Type:   "sub",
		Icon:   "assignment_turned_in",
		Permit: "Jefe Biomedico",
		Children: []ChildrenItems{
			{State: "revision", Name: "Validar Solicitud", Type: "link"},
			{State: "validado", Name: "Solicitudes Aprobadas", Type: "link"},
			{State: "final", Name: "Solicitud Final", Type: "link"},
		},
	},
	{
		State:  "proveedores",
		Name:   "Proveedores",
		Type:   "sub",
		Icon:   "apartment",
		Permit: "Consulta",
		Children: []ChildrenItems{
			{State: "lista", Name: "Lista de Proveedores", Type: "link"},
		},
	},
	{
		State:  "contratos",
		Name:   "Contratos",
		Type:   "sub",
		Icon:   "receipt",
		Permit: "Administrador de contratos",
		Children: []ChildrenItems{
			{State: "administrar", Name: "Administrar", Type: "link"},
			{State: "seguimiento", Name: "Seguimiento", Type: "link"},
		},
	},
	{
		State:  "traslado",
		Name:   "Traslados de pacientes",
		Type:   "sub",
		Icon:   "receipt",
		Permit: "Solicitud traslado",
		Children: []ChildrenItems{
			{State: "solicitud", Name: "Nueva Solicitud", Type: "link"},
			{State: "realizadas", Name: "Solicitudes Realizadas", Type: "link"},
			{State: "cancelar", Name: "Cancelar Solicitud", Type: "link"},
		},
	},
	{
		State:  "traslado",
		Name:   "Traslados de pacientes",
		Type:   "sub",
		Icon:   "receipt",
		Permit: "Administrador Traslado",
		Children: []ChildrenItems{
			{State: "firma", Name: "Firmas", Type: "link"},
			{State: "solicitud", Name: "Nueva Solicitud", Type: "link"},
			{State: "realizadas", Name: "Solicitudes Realizadas", Type: "link"},
			{State: "cancelar", Name: "Cancelar Solicitud", Type: "link"},
			{State: "manager", Name: "Administrar Traslados", Type: "link"},
		},
	},
}

/*****  Funcion para filtrar el menu item segun los permisos del usuario  *****/
func FiltrarPermisos(permisos []Permits) (error, []Menu) {
	var menuItem []Menu
	for _, permis := range permisos {
		for _, item := range MENUITEMS {
			if permis.Permiso == item.Permit {
				menuItem = append(menuItem, item)
			}
		}
	}
	return nil, menuItem
}
