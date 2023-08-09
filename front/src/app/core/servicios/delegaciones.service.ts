import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {RespuestaPeticion} from '../models/estructuras_generales';
import {environment} from '../../../environments/environment';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class DelegacionesService {

  constructor(
    private _http: HttpClient
  ) {
  }

  ConsultarDelegaciones() {
    return this._http.get<RespuestaPeticion>(`${environment.host}/delegaciones`,).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  /* ******  Consultar delegaciones sin token  ******   */
  consultDelegations(type: string) {
    let queryParams = new HttpParams();
    if (type !== '') {
      queryParams = queryParams.set('type', type);
    }
    return this._http.get<RespuestaPeticion>(`${environment.host}/delegaciones/registro`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
