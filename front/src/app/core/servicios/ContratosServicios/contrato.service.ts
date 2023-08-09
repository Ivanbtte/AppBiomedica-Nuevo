import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {RespuestaPeticion} from '../../models/estructuras_generales';
import {environment} from '../../../../environments/environment';
import {map} from 'rxjs/operators';
import {Contrato, ContratoNuevo, UnContrato} from '../../models/contrato.interface';

@Injectable({
  providedIn: 'root'
})
export class ContratoService {

  constructor(
    private _http: HttpClient
  ) {
  }

  ConsultarContratos(id_contrato: string, tipo: string, filtros: string, proveedores: string, delegaciones: string, unidad_med: string, estatus: string,
                     clv_dele: string, num_pagina: string, num_registro: string,) {
    let queryParams = new HttpParams();
    if (num_pagina !== '') {
      queryParams = queryParams.set('pagina', num_pagina);
    }
    if (num_registro !== '') {
      queryParams = queryParams.set('num_registros', num_registro);
    }
    if (id_contrato !== '') {
      queryParams = queryParams.set('id_contrato', id_contrato);
    }
    if (filtros !== '') {
      queryParams = queryParams.set('filtro', filtros);
    }
    if (proveedores !== '') {
      queryParams = queryParams.set('proveedores', proveedores);
    }
    if (delegaciones !== '') {
      queryParams = queryParams.set('delegaciones', delegaciones);
    }
    if (clv_dele !== '') {
      queryParams = queryParams.set('clv_dele', clv_dele);
    }
    if (tipo !== '') {
      queryParams = queryParams.set('tipo', tipo);
    }
    return this._http.get<RespuestaPeticion>(`${environment.host}/contrato`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  AgregarContrato(obj: ContratoNuevo) {
    return this._http.post<RespuestaPeticion>(`${environment.host}/contrato`, obj).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  AgregarUnContrato(obj: UnContrato) {
    return this._http.post<RespuestaPeticion>(`${environment.host}/contrato/solo`, obj).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  ValidarArchivo(data: FormData) {
    return this._http.post<RespuestaPeticion>(`${environment.host}/contrato/verificar`, data).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  SubirArchivo(data: FormData, tipo_contrato: string) {
    const query_params = new HttpParams().set('tipo_contrato', tipo_contrato);
    return this._http.post<RespuestaPeticion>(`${environment.host}/contrato/subir`, data, {params: query_params}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  ActualizarContrato(id_contrato: string, obj: Contrato) {
    const queryParams = new HttpParams().set('id_contrato', id_contrato);
    return this._http.put<RespuestaPeticion>(`${environment.host}/contrato`, obj, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  DescargarPlantillaContrato() {
    return this._http.get(`${environment.host}/plantillas/plantilla_contrato.xlsx`, {responseType: 'blob'});
  }

  ConsultarContratoTraslado(unidad_id: string, fecha: string) {
    const queryParams = new HttpParams().set('unidad_id', unidad_id).set('fecha_cita', fecha);
    return this._http.get<RespuestaPeticion>(`${environment.host}/contrato/traslados`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  ConsultarContratoSubTipo(subTipo: string) {
    const queryParams = new HttpParams().set('subTipo', subTipo);
    return this._http.get<RespuestaPeticion>(`${environment.host}/contrato/subTipo`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
  ConsultarPorcentajeAcompaniante(numContrato: string) {
    const queryParams = new HttpParams().set('numContrato', numContrato);
    return this._http.get<RespuestaPeticion>(`${environment.host}/contrato/promedio/mayor/acompaniante`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
  ConsultarDetalleContrato(contratos: string) {
    const queryParams = new HttpParams().set('contratos', contratos);
    return this._http.get<RespuestaPeticion>(`${environment.host}/contrato/detalle/contrato`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
