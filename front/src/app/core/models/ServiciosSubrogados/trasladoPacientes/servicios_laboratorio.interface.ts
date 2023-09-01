export interface SolicitudServLaboratorio {
  id: string;
  folio: string;
  fecha?: string;
  unidadMedicaId: number;
  pacienteId: string;
  tipodeServicio: string;
  motivoSubrogacion: string;
  ramo: string;
  tipo: string;
  vigenciadeDerechos: string;
  diagnostico: string;
  grupoSubrogarId: string;
  proveedorId: string;
  usuarioId: string;
}