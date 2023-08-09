import { Justificacion } from '../../../../traslado-pacientes/components/dialogs/dialog-autoriza-acomp/dialog-autoriza-acomp.component';
export interface Solicitudtraslado {
  id: number;
  folio: string;
  version: number;
  representanteLegal: string;
  nombrePaciente: string;
  nss: string;
  agregadoMedico: string;
  fechaCita?: string;
  origen: string;
  destino: string;
  acompañante: Boolean;
  ftp01: string;
  fechaExpedicion: string;
  contratoNumeroContrato: string;
  unidadMedId: number;
  unidadMed: any;
  usuarioId: string;
  usuario: any;
  estado: Boolean;
}

export interface SolicitudtrasladoBoletos {
  Solicitud: Solicitudtraslado;
  Boletos: Boletos[];
  Acompaniante: Acompaniante;
  AutorizacionesEspeciales: AutorizacionesEspeciales;
}

export interface Boletos {
  Ida: boolean;
  Observaciones: string;
  Regreso: boolean;
  Tipo: number;
  TipoBoleto: number;
  Total: number;
}

export interface Acompaniante {
  id: string;
  motivo: string;
  justificacion: string;
  nombreAutoriza: string;
  cargo: string;
  matricula: string;
  tarjetaInapan: boolean;
  solicitudTrasladoFtp01: string;

}

export interface SolicitudesCanceladas {
  id: string;
  motivo: string;
  fechaCancelacion: string;
  solicitudTrasladoFtp01: string;
  solicitudTraslado: any;
  usuarioId: string;
  usuario: any;
  estado: boolean;
}
export interface AutorizacionesEspeciales {
  id: string;
  justificacion: string;
  tipoAutorizacion: string;
  solicitudTrasladoFolio: string;
  solicitudTraslado: Solicitudtraslado;
  estado: boolean;
}