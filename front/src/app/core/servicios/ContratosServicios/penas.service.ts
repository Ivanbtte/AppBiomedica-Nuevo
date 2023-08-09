import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {RespuestaPeticion} from '../../models/estructuras_generales';
import {environment} from '../../../../environments/environment';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class PenasService {

  constructor(
    private _http: HttpClient
  ) {
  }

  /**************  Servicion para penas deductivas  **************/
  ConsultarPenaDeductiva(numero_contrato: string) {
    const queryParams = new HttpParams().set('numero_contrato', numero_contrato);
    return this._http.get<RespuestaPeticion>(`${environment.host}/pena_deductiva`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  /************** Servicion para penas convencionales  **************/
  ConsultarPenaConvencional(numero_contrato: string) {
    const queryParams = new HttpParams().set('numero_contrato', numero_contrato);
    return this._http.get<RespuestaPeticion>(`${environment.host}/pena_convencional`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
