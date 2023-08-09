import {Injectable} from '@angular/core';
import {Observable} from 'rxjs';
import {RespuestaPeticion} from '../../models/estructuras_generales';
import {environment} from '../../../../environments/environment';
import {map} from 'rxjs/operators';
import {HttpClient, HttpParams} from '@angular/common/http';
import {ConceptoContrato} from '../../models/concepto_contrato.interface';

@Injectable({
  providedIn: 'root'
})
export class ConceptosContratoService {

  constructor(
    private _http: HttpClient
  ) {
  }

  AgregarConceptoContrato(obj: any, id_contrato: string){
    let queryParams = new HttpParams();
    if (id_contrato !== '') {
      queryParams = queryParams.set('id_contrato', id_contrato);
    }
    return this._http.post<RespuestaPeticion>(`${environment.host}/contrato/conceptos`, obj, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  ActualizarConceptoContrato(id_concepto: string, obj: ConceptoContrato, numero_contrato: string) {
    let queryParams = new HttpParams().set('id_concepto', id_concepto);
    if (numero_contrato != '') {
      queryParams = queryParams.set('numero_contrato', numero_contrato);
    }
    return this._http.put<RespuestaPeticion>(`${environment.host}/contrato/conceptos`, obj, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  ConsultarConceptosContrato(id_contrato: string) {
    const queryParams = new HttpParams().set('id_contrato', id_contrato);
    return this._http.get<RespuestaPeticion>(`${environment.host}/contrato/conceptos`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  ConsultarConceptos(id_contrato: string, tipo: string, subTipo: string, num_pagina: string, num_registro: string){
    let queryParams = new HttpParams().set('num_contrato', id_contrato);
    if (tipo !== '') {
      queryParams = queryParams.set('tipo_contrato', tipo);
    }
    if (subTipo !== '') {
      queryParams = queryParams.set('sub_tipo_contrato', subTipo);
    }
    if (num_pagina !== '') {
      queryParams = queryParams.set('pagina', num_pagina);
    }
    if (num_registro !== '') {
      queryParams = queryParams.set('num_registros', num_registro);
    }
    return this._http.get<RespuestaPeticion>(`${environment.host}/contrato/concepto`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  DescargarArchivoPlantilla(nombre: string) {
    return this._http.get(`${environment.host}/plantillas/plantilla_${nombre}.xlsx`, {responseType: 'blob'});
  }

  ValidarConceptosFile(data: FormData, tipo_contrato: string, sub_tipo: string) {
    let queryParams = new HttpParams();
    if (tipo_contrato !== '') {
      queryParams = queryParams.set('tipo_contrato', tipo_contrato);
    }
    if (sub_tipo !== '') {
      queryParams = queryParams.set('sub_tipo_contrato', sub_tipo);
    }
    return this._http.post<RespuestaPeticion>(`${environment.host}/contrato/concepto/validar`, data, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  AgregarConceptosArchivo(data: FormData, tipo_contrato: string, sub_tipo: string) {
    let queryParams = new HttpParams();
    if (tipo_contrato !== '') {
      queryParams = queryParams.set('tipo_contrato', tipo_contrato);
    }
    if (sub_tipo !== '') {
      queryParams = queryParams.set('sub_tipo_contrato', sub_tipo);
    }
    return this._http.post<RespuestaPeticion>(`${environment.host}/contrato/concepto`, data, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  EliminarConcepto(id_concepto: string) {
    const queryParams = new HttpParams().set('id_concepto', id_concepto);
    return this._http.delete<RespuestaPeticion>(`${environment.host}/contrato/conceptos`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
