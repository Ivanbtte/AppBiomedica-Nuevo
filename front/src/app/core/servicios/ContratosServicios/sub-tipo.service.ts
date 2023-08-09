import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {RespuestaPeticion} from '../../models/estructuras_generales';
import {environment} from '../../../../environments/environment';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class SubTipoService {

  constructor(
    private _http: HttpClient
  ) {
  }

  /* ******  Servicio para consultar los subtipos de contratos  ******   */
  ConsultaSubTipo(tipo_contrato: string) {
    const queryParams = new HttpParams().set('tipo_contrato', tipo_contrato);
    return this._http.get<RespuestaPeticion>(`${environment.host}/sub_tipo`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
