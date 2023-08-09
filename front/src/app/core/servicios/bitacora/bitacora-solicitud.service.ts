import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {RespuestaPeticion} from '../../models/estructuras_generales';
import {environment} from '../../../../environments/environment';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class BitacoraSolicitudService {

  constructor(
    private http: HttpClient
  ) {
  }

  ConsultarBitacoraSolicitud(Id: string) {
    const parametro = new HttpParams().set('id_solicitud', Id);
    return this.http.get<RespuestaPeticion>(`${environment.host}/bitacora_solicitud`, {params: parametro}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
