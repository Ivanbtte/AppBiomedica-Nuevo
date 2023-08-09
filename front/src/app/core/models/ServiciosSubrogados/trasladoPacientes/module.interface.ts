export interface Modulo {
  id: string;
  name: string;
  description: string;
  status: boolean;
}

export interface Actions {
  id: string;
  name: string;
  description: string;
  value: boolean;
  moduloId: string;
  modulo: Modulo;
  status: boolean;
}

export interface ActionsPermit {
  action: Actions;
  permit: boolean;
}
