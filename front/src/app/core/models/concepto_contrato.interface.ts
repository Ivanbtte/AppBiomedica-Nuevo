export interface ConceptoContrato {
  Id: string;
  PrecioUniSnIva: number;
  CantidadConcepto: number;
  CantidadAmpliada: number;
  Marca: string;
  Modelo: string;
  ObjetoContratacion: string;
  FechaMaxEntrega: string;
  GarantiaBienes: string;
  PreiIdArticulo: string;
  Prei: any;
  ContratoNumeroContrato: string;
  Contrato: any;
}

export interface ConceptoContratoTemp {
  PrecioUniSnIva: number;
  CantidadConcepto: number;
  Marca: string;
  Modelo: string;
  ObjetoContratacion: string;
  FechaMaxEntrega: string;
  GarantiaBienes: number;
  PreiIdArticulo: string;
}

export interface ConceptosNuevos {
  ConceptoContrato: ConceptoContratoTemp[];
}

export interface DistribucionUnidadMed {
  Id: string;
  Cantidad: number;
  UnidadMedClvPresupuestal: string;
  ConceptoContratoId: string;
}
