export interface TrasladoPacientes {
  id: string;
  servicio: string;
  origen: string;
  destino: string;
  unidadMedida: string;
  lugarOtorgamiento: string;
  cantidadMin: number;
  cantidadMax: number;
  precioOfertado: number;
  importeMin: number;
  importeMax: number;
  unidadMedId: number;
  unidadMed: any;
  conceptosId: string;
}
