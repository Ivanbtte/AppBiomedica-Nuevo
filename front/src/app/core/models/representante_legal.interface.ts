export interface RepresentanteLegal {
  Id: string;
  ContratoNumeroContrato: string;
  Contrato: any;
  NombreCompleto: string;
}

export interface CorreoContrato {
  Id: string;
  TipoCorreo: string;
  CorreosId: string;
  Correos: any;
  ContratoNumeroContrato: string;
  Contrato: any;
}
