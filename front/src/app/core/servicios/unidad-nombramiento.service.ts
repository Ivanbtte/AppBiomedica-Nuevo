import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {RespuestaPeticion} from '../models/estructuras_generales';
import {environment} from '../../../environments/environment';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class UnidadNombramientoService {

  constructor(
    private http: HttpClient,
  ) {
  }

  unidadNombramiento(ID: string) {
    return this.http.get<RespuestaPeticion>(`${environment.host}/unidad_nomb/${ID}`).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  consultaTodasUnidades(id: string) {
    const param1 = new HttpParams().set('id', id);
    return this.http.get<RespuestaPeticion>(`${environment.host}/unidad_nomb`, {params: param1}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
