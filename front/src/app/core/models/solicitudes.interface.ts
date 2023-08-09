export interface Conceptos {
  SolicitudId: string;
  ClaveId: string;
  Descripcion: string;
  CantidadSolicitada: number;
  Precio: number;
  Total: number;
  ServiciosProformaId: string;
  EstatusId: number;
  Fecha: string;
}

export interface Solicitud {
  Id: string;
  UsuarioId: string;
  UnidadMedId: number;
  TipoSolicitudId: number;
  EstatusId: number;
  FechaCreacion: string;
  Total: number;
  Correo: string;
  Estado: boolean;
}

