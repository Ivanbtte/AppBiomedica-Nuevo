import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {RespuestaPeticion} from '../../models/estructuras_generales';
import {environment} from '../../../../environments/environment';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class ServiciosUsuariosService {

  constructor(
    private http: HttpClient,
  ) {
  }

  serviciosUsuarios(idUnidad: string, idUsuario: string) {
    let queryParams = new HttpParams();
    if (idUnidad !== '') {
      queryParams = queryParams.set('unidad_id', idUnidad);
    }
    if (idUsuario !== '') {
      queryParams = queryParams.set('usuario_id', idUsuario);
    }
    return this.http.get<RespuestaPeticion>(`${environment.host}/servicios_usuarios`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
