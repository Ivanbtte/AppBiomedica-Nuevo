import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {RespuestaPeticion} from '../../models/estructuras_generales';
import {environment} from '../../../../environments/environment';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class ContactosContratoService {

  constructor(
    private _http: HttpClient
  ) {
  }

// ********************************************************************************************************
  /* *****  Correos de contratos  *****   */
  ConsultaCorreosContrato(n_contrato: string) {
    const queryParams = new HttpParams().set('n_contrato', n_contrato);
    return this._http.get<RespuestaPeticion>(`${environment.host}/correos_contrato`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  EliminarCorreoContrato(id_correo: string) {
    const queryParams = new HttpParams().set('id_correo', id_correo);
    return this._http.delete<RespuestaPeticion>(`${environment.host}/correos_contrato`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

// ********************************************************************************************************
  /* *****  Telefonos de contratos  *****   */
  ConsultaTelefonosContrato(n_contrato: string) {
    const queryParams = new HttpParams().set('n_contrato', n_contrato);
    return this._http.get<RespuestaPeticion>(`${environment.host}/telefonos_contrato`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  EliminarTelefonoContrato(id_telefono: string) {
    const queryParams = new HttpParams().set('id_telefono', id_telefono);
    return this._http.delete<RespuestaPeticion>(`${environment.host}/telefonos_contrato`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
