export interface Credenciales {
  correo: string;
  password: string;
}

export interface Personas {
  Id: string;
  Nombre: string;
  ApellidoPat: string;
  ApellidoMat: string;
  FechaNac: string;
  Genero: string;
  Correo: string;
  Telefono: string;
  Extension: number;
  Celular: string;
  Matricula: string;
  Estado: boolean;
}

export interface UnidadesM {
  Id: number;
  RegionId: string;
  Regiones: any;
  Clues: string;
  ClvPersonal: string;
  UiPrei: string;
  ClvUbicacion: string;
  ClvPresupuestal: string;
  DelegacionesClvDele: string;
  Delegacion: any;
  UniPresupuestal: number;
  ClvNivelAtnId: number;
  Nivel: any;
  DenominacionUni: string;
  TipoServicio: string;
  DescripServicio: string;
  NumUnidad: string;
  NomUnidad: string;
  UbicacionDenom: string;
  Direccion: string;
  ClvVialidad: number;
  Vialidad: any;
  NombVialidad: string;
  NumExterior: string;
  ClvTipoAsent: number;
  Asentamiento: any;
  NombAsentmt: string;
  Cp: number;
  ClvJurisdiccion: number;
  IdJurisdiccion: number;
  ClvEntidad: number;
  ClvMunicipio: number;
  ClvLocalidad: number;
  Localidad: string;
  Jurisdiccion: any;
  Latitud: number;
  Longitud: number;
  InicioProduct: string;
  Estado: boolean;

}

export interface Areas {
  valor: string;
  valorMostrado: string;
}

export interface TiposUsuario {
  valor: string;
  valorMostrado: string;
}

export interface Cargo {
  Id: string;
  Nombre: string;
  TipoUsuarioId: string;
  Estado: boolean;
}

export interface DatosUsuario {
  Persona: Personas;
  Usuario: Usuario;
  PuestoOrganigramaId: string;
}

export interface Nombramiento {
  Area: string;
  Unidades: number [];
  CargoId: string;
  PersonaId: string;
}

export interface Usuario {
  Id?: string;
  Correo: string;
  Contraseña: string;
  UsuarioTipoId: string;
  PersonaId: string;
  FechaRegistro?: Date;
  Estado?: boolean;
}

export interface UsuarioPersona {
  Id: string;
  Correo: string;
  TipoUsuario: string;
  PersonaId: string;
  Persona: {
    Nombre: string;
    ApellidoPat: string;
    ApellidoMat: string;
    FechaNac: string;
    Genero: string;
    Correo: string;
    Telefono: string;
    Extension: number;
    Celular: string;
    Matricula: string;
  };
}

export interface Passwords {
  Password: string;
  NewPassword: string;
}

export interface Person {
  status: boolean;
  personStatus: boolean;
  data: Personas;
}
