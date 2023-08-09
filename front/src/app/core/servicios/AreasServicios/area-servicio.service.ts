import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {RespuestaPeticion} from '../../models/estructuras_generales';
import {environment} from '../../../../environments/environment';
import {map} from 'rxjs/operators';
import {AreaServicioUnidades} from '../../models/areaServicios';

@Injectable({
  providedIn: 'root'
})
export class AreaServicioService {

  constructor(
    private http: HttpClient
  ) {
  }

  ConsultarAreaServicio(id_unidad: string, id_servicio: string, id_usuario: string) {
    let query_params = new HttpParams();
    if (id_unidad !== '') {
      query_params = query_params.set('id_unidad', id_unidad);
    }
    if (id_servicio !== '') {
      query_params = query_params.set('id_servicio', id_servicio);
    }
    if (id_usuario !== '') {
      query_params = query_params.set('id_usuario', id_usuario);
    }
    return this.http.get<RespuestaPeticion>(`${environment.host}/area_servicio`, {params: query_params}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  AgregarAreaServicio(obj: AreaServicioUnidades) {
    return this.http.post<RespuestaPeticion>(`${environment.host}/area_servicio`, obj).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
