import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {RespuestaPeticion} from '../../models/estructuras_generales';
import {environment} from '../../../../environments/environment';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class ContratoDelegacionService {

  constructor(
    private _http: HttpClient
  ) {
  }

  ConsultarDelegacionContrato(num_contrato: string) {
    const queryParams = new HttpParams().set('num_contrato', num_contrato);
    return this._http.get<RespuestaPeticion>(`${environment.host}/contrato/delegacion`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
