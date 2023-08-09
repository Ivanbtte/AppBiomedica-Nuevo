import {Injectable} from '@angular/core';
import {RespuestaPeticion} from '../../core/models/estructuras_generales';
import {environment} from '../../../environments/environment';
import {map} from 'rxjs/operators';
import {HttpClient, HttpParams} from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class PreiService {
  constructor(
    private http: HttpClient,
  ) {
  }

  getPrei(cuadro: string, nombre: string, num_pagina: string, num_registro: string, grupo: string, generico: string, especifico: string) {
    let params1 = new HttpParams().set('nombre', nombre).set('pagina', num_pagina).set('num_registros', num_registro);
    if (cuadro !== '') {
      params1 = params1.set('cuadro_id', cuadro);
    }
    if (grupo !== '') {
      params1 = params1.set('grupo', grupo);
    }
    if (generico !== '') {
      params1 = params1.set('generico', generico);
    }
    if (especifico !== '') {
      params1 = params1.set('especifico', especifico);
    }
    return this.http.get<RespuestaPeticion>(`${environment.host}/prei`, {params: params1}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  getClavePrei(Id: string, id_articulo: string) {
    let query_params = new HttpParams();
    if (Id !== '') {
      query_params = query_params.set('id', Id);
    }
    if (id_articulo !== '') {
      query_params = query_params.set('id_articulo', id_articulo);
    }
    return this.http.get<RespuestaPeticion>(`${environment.host}/prei/clave`, {params: query_params}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  getGrupoPrei(ID: string) {
    const params1 = new HttpParams().set('grupo', ID);
    return this.http.get<RespuestaPeticion>(`${environment.host}/prei`, {params: params1}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  getCuadroBasico() {
    return this.http.get<RespuestaPeticion>(`${environment.host}/cuadro_basico`).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  consultarFiltrosPrei(grupo: string, generico: string, grupos: boolean) {
    let queryParams = new HttpParams();
    if (grupo !== '') {
      queryParams = queryParams.set('grupo', grupo);
    }
    if (generico !== '') {
      queryParams = queryParams.set('generico', generico);
    }
    if (grupos !== false) {
      const valorParse: string = String(grupos);
      queryParams = queryParams.set('grupos', valorParse);
    }
    return this.http.get<RespuestaPeticion>(`${environment.host}/prei/filtro`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
