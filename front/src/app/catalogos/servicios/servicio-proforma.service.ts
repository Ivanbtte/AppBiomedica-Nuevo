import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {RespuestaPeticion} from '../../core/models/estructuras_generales';
import {environment} from '../../../environments/environment';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class ServicioProformaService {

  constructor(
    private http: HttpClient
  ) {
  }

  getServiciosProf() {
    return this.http.get<RespuestaPeticion>(`${environment.host}/servicios_proforma`).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  /* *****  Cambiar esto despues debe de ir en la funcion de arriba   *****   */
  getServicio(id: string) {
    const params1 = new HttpParams().set('id', id);
    return this.http.get<RespuestaPeticion>(`${environment.host}/servicios_proforma`, {params: params1}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  /* *****  Servicio para consultar los servicios de una sola clave proforma recibe un id  *****   */
  getClaveProformaServicios(ID: string) {
    return this.http.get<RespuestaPeticion>(`${environment.host}/servicios_proforma/${ID}`).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
