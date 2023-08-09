import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {RespuestaPeticion} from '../../core/models/estructuras_generales';
import {environment} from '../../../environments/environment';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class ProformaService {

  constructor(
    private http: HttpClient,
  ) {
  }

  /* *****  Servicio para consultar todos el catalogo proforma  *****   */
  getProforma(servicio: string, nombre: string, num_pagina: string, num_registro: string) {
    let parametros = new HttpParams().set('nombre', nombre).set('pagina', num_pagina).set('num_registros',
      num_registro);
    if (servicio !== '') {
      parametros = parametros.set('servicio', servicio);
    }
    return this.http.get<RespuestaPeticion>(`${environment.host}/proforma`, {params: parametros}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  /* *****  Servicio para consultar una sola clave proforma recibe un id  *****   */
  getClaveProforma(ID: string) {
    return this.http.get<RespuestaPeticion>(`${environment.host}/proforma/${ID}`).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
