export interface PenasDeductivas {
  ConceptoObligacion: string;
  NivelServicio: string;
  UnidadMedida: string;
  Deduccion: number;
  DescripcionDeduccion: string;
  LimiteIncumplimiento: string;
}

export interface PenasConvencionales {
  Descripcion: string;
  Porcentaje: number;
  NotaCreditoOficio: string;
  ReferenciaCobro: string;
  FechaCobro: string;
}
