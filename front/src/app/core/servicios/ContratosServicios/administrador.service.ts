import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {RespuestaPeticion} from '../../models/estructuras_generales';
import {environment} from '../../../../environments/environment';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class AdministradorService {

  constructor(
    private _http: HttpClient
  ) {
  }

  /**************  Servicion para consultar administradores o auxiliares de un contrato  **************/
  ConsultarAdministradores(numero_contrato: string) {
    const queryParams = new HttpParams().set('numero_contrato', numero_contrato);
    return this._http.get<RespuestaPeticion>(`${environment.host}/administrador_aux`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
