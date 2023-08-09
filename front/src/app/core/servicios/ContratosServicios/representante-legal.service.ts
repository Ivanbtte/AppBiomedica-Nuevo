import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {RespuestaPeticion} from '../../models/estructuras_generales';
import {environment} from '../../../../environments/environment';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class RepresentanteLegalService {

  constructor(
    private _http: HttpClient
  ) {
  }

  ConsultaRepresentante(n_contrato: string) {
    const queryParams = new HttpParams().set('n_contrato', n_contrato);
    return this._http.get<RespuestaPeticion>(`${environment.host}/representante`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  EditarRepresentante(n_contrato: string, representante: any) {
    const queryParams = new HttpParams().set('num_contrato', n_contrato);
    return this._http.put<RespuestaPeticion>(`${environment.host}/representante`, representante, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
