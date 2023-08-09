import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {RespuestaPeticion} from '../../models/estructuras_generales';
import {environment} from '../../../../environments/environment';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class ContratoPuestoService {

  constructor(
    private _http: HttpClient
  ) {
  }

  /**************  Servicion para consultar tipo de contrato por puesto **************/
  Consultarcontrato_puesto(puesto_id: string) {
    const queryParams = new HttpParams().set('puesto_id', puesto_id);
    return this._http.get<RespuestaPeticion>(`${environment.host}/contrato_puesto`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
