import {UnidadesM} from '../../authentication/interfaces/usuario.interface';

export interface UserUnitMedical {
  Id: string;
  UsuarioPuestoId: string;
  UsuarioPuesto: any;
  UnidadMedId: number;
  UnidadMed: UnidadesM
  Estado: boolean;
}
