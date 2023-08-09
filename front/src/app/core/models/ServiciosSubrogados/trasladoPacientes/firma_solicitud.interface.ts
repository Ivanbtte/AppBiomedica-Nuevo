export interface FirmaSolicitud {
  id: string;
  nombre: string;
  cargo: string;
  matricula: string;
  tipo: number;
  fecha: Date;
  unidadMedId: number;
  unidadMed: any;
  usuarioId: string;
  usuario: any;
  estado: boolean;
}
