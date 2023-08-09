import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {RespuestaPeticion} from '../../models/estructuras_generales';
import {environment} from '../../../../environments/environment';
import {map} from 'rxjs/operators';
import {ProveedoresInterface} from '../../models/proveedores.interface';

@Injectable({
  providedIn: 'root'
})
export class ProveedoresService {

  constructor(
    private http: HttpClient
  ) {
  }

  ConsultarProveedores(id: string, nombre: string, num_pagina: string, num_registro: string) {
    let queryParams = new HttpParams().set('pagina', num_pagina).set('num_registros', num_registro);
    if (id !== '') {
      queryParams = queryParams.set('id_proveedor', id);
    }
    if (nombre !== '') {
      queryParams = queryParams.set('nombre_proveedor', nombre);
    }
    return this.http.get<RespuestaPeticion>(`${environment.host}/proveedores`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  // ********************************************************************************************************
  /* *****  funcion para consultar proveedores por varios parametros  *****   */
  ConsultarProveedoresFiltro(filtros: string) {
    let queryParams = new HttpParams();
    if (filtros !== '') {
      queryParams = queryParams.set('filtro', filtros);
    }
    return this.http.get<RespuestaPeticion>(`${environment.host}/proveedores/nombre`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  AgregarProveedores(obj: ProveedoresInterface) {
    return this.http.post<RespuestaPeticion>(`${environment.host}/proveedores`, obj).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  ActualizarProveedores(obj: ProveedoresInterface, Id_proveedor: string) {
    const queryParams = new HttpParams().set('id_proveedor', Id_proveedor);
    return this.http.put<RespuestaPeticion>(`${environment.host}/proveedores`, obj, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  EliminarProveedor(id_proveedor: string) {
    const queryParams = new HttpParams().set('id_proveedor', id_proveedor);
    return this.http.delete<RespuestaPeticion>(`${environment.host}/proveedores`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
