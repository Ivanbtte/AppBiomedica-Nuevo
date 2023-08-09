import {Organigrama} from './departament.interface';

export interface CompanyPosition {
  id: string;
  puestoJefe: boolean;
  numeroPlazas: number;
  organigramaId: string;
  organigrama: Organigrama;
  catalogoPuestoId: string;
  catalogoPuesto: any;
  estado: boolean;
}
