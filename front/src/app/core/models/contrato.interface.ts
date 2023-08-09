import {AdminAux} from './admin_aux.interface';
import {PenasConvencionales, PenasDeductivas} from './penas_deductivas.interface';

export interface Contrato {
  Id: string;
  NumeroContrato: string;
  InicioContrato: string;
  FinContrato: string;
  ProcedContratacion: string;
  TipoProcedContratacion: string;
  FechaFallo: string;
  MontoTotal: number;
  TipoContratoId: string;
  TipoContrato: any;
  ProveedorNProvImss: string;
  ContactoProveedorId: string;
  Proveedor: any;
  Estado: boolean;
}

export interface ContratoNuevo {
  Contrato: Contrato;
  Delegaciones: any[];
  AdminAuxContrato: AdminAux [];
  PenasConvencionales: PenasConvencionales;
  RepresentanteLegal: RepresentanteLegal;
  PenasDeductivas: PenasDeductivas[];
  Fianza: Fianza;
}

export interface UnContrato {
  Contrato: Contrato;
  Delegaciones: any[];
  Representante: any;
}

export interface Fianza {
  Id: string;
  Afianzadora: string;
  NumPoliza: string;
  Tipo: string;
  MontoFianza: number;
}

export interface RepresentanteLegal {
  Id: string;
  ContratoId: string;
  ContactoProveedorId: string;
}
