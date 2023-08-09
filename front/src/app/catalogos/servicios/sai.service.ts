import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {RespuestaPeticion} from '../../core/models/estructuras_generales';
import {environment} from '../../../environments/environment';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class SaiService {
  constructor(
    private http: HttpClient,
  ) {
  }

  getSai(grupo: string, temas: string, nombre: string, num_pagina: string, num_registro: string, cuadro: string, generico: string,
         especifico: string) {
    let params1 = new HttpParams().set('nombre', nombre).set('pagina', num_pagina).set('num_registros', num_registro);
    if (temas !== '') {
      params1 = params1.set('tema', temas);

    }
    if (grupo !== '') {
      params1 = params1.set('grupo', grupo);
    }
    if (cuadro !== '') {
      params1 = params1.set('cuadro', cuadro);
    }
    if (generico !== '') {
      params1 = params1.set('generico', generico);
    }
    if (especifico !== '') {
      params1 = params1.set('especifico', especifico);
    }
    return this.http.get<RespuestaPeticion>(`${environment.host}/sai`, {params: params1}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  getClaveSai(ID: string, clave: string) {
    let queryparams = new HttpParams();
    if (ID !== '') {
      queryparams = queryparams.set('id', ID);
    }
    if (clave !== '') {
      queryparams = queryparams.set('clave', clave);
    }
    return this.http.get<RespuestaPeticion>(`${environment.host}/sai/clave`, {params: queryparams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
