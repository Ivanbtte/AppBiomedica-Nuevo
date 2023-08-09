export interface BitacoraUsuarioPuesto {
  id: string;
  usuarioModificadoId: string;
  usuarioModificado: any;
  usuarioModificanteId: string;
  usuarioModificante: any;
  tipoModificacionId: string;
  tipoModificacion: any;
  fecha: Date;
  comentario: string;
  status: boolean;
}

export interface UnsuscribeUser {
  usuarioPuestoId: string;
  tipoModificacionId: string;
  justificacion: string;
}
