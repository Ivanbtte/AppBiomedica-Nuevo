export interface SubTipo {
  Id: string;
  SubTipo: string;
  TipoContratoId: string;
  TipoContrato: any;
  Estado: boolean;
}

export interface SubTipoContrato {
  SubTipoId: string;
  SubTipo: any;
  ContratoNumeroContrato: string;
  Contrato: any;
  Estado: boolean;
}
