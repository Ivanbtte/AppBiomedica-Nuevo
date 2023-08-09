export interface Prei {
  Id: number;
  IdArticulo: string;
  IdArticuloCorto: string;
  Grupo: string;
  Generico: string;
  Especifico: string;
  Diferenciador: string;
  Variable: string;
  NombreCorto: string;
  Descripcion: string;
  Precio: number;
  UnidadesMedidasId: number;
  UnidadesMedidas: {
    Id: number;
    Unidad: string;
  };
  CuadroBasicoId: number;
  CuadroBasico: {
    Id: number;
    Grupo: string;
  };
  Estado: boolean;
}

export interface CuadroBasico {
  Id: number;
  Grupo: string;
  Estado: boolean;
}
