export interface Tema {
  Id: number;
  Tema: string;
}

export interface GrupoTema {
  Grupo: string;
  TemaId: number;
}

export interface TemaGrupo {
  Grupo: string;
  TemaId: number;
  Tema: {
    Id: number;
    Tema: string;
  };
}
