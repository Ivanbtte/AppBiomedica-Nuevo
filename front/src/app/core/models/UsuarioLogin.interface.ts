import {Menu} from '../../shared/menu-items/menu-items';

export interface UsuarioLogeado {
  UsuarioId: string;
  Correo: string;
  Tipo: string;
  Puesto: string;
  Nombre: string;
  AppellidoP: string;
  ApellidoM: string;
  Matricula: string;
  PersonaId: string;
  PuestoId: string;
  PuestoOrganigramaId: string;
  MenuItem: Menu[];
  exp: number;
  iat: number;
}

export interface Personaw {
  id: string;
  Persona: {
    id: string;
  };
}
