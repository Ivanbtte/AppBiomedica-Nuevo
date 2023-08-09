import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {RespuestaPeticion} from '../../models/estructuras_generales';
import {environment} from '../../../../environments/environment';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class SubTipoContratoService {

  constructor(private _http: HttpClient) {
  }

  /* ******  Servicio para consultar los subtipos que tiene un contrato  ******   */
  ConsultaSubTipoContrato(numero_contrato: string) {
    const queryParams = new HttpParams().set('numero_contrato', numero_contrato);
    return this._http.get<RespuestaPeticion>(`${environment.host}/sub_tipo_contrato`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
