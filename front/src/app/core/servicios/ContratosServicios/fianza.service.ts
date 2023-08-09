import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {RespuestaPeticion} from '../../models/estructuras_generales';
import {environment} from '../../../../environments/environment';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class FianzaService {

  constructor(
    private _http: HttpClient
  ) {
  }

  /**************  Servicion para consultar una fianza de un contrato  **************/
  Consultarfianza(numero_contrato: string) {
    const queryParams = new HttpParams().set('numero_contrato', numero_contrato);
    return this._http.get<RespuestaPeticion>(`${environment.host}/fianza`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  /* ******  Servicio para agregar una fianza  ******   */
  AgregarFianza(numero_contrato: string, fianza: any) {
    const queryParams = new HttpParams().set('numero_contrato', numero_contrato);
    return this._http.post<RespuestaPeticion>(`${environment.host}/fianza`, fianza, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
