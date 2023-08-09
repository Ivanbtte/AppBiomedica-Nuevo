import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {RespuestaPeticion} from '../../models/estructuras_generales';
import {environment} from '../../../../environments/environment';
import {map} from 'rxjs/operators';
import {DistribucionUnidadMed} from '../../models/concepto_contrato.interface';

@Injectable({
  providedIn: 'root'
})
export class DistribucionConceptosService {

  constructor(
    private _http: HttpClient
  ) {
  }

  ConsultarDistribucion(id_concepto: string) {
    const queryParams = new HttpParams().set('id_concepto', id_concepto);
    return this._http.get<RespuestaPeticion>(`${environment.host}/distribucion`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  DescargarPlantillaDistribucion(nombre: string) {
    return this._http.get(`${environment.host}/plantillas/plantilla_${nombre}.xlsx`, {responseType: 'blob'});
  }

  ValidarArchivo(data: FormData, conceptos: string) {
    let queryParams = new HttpParams();
    if (conceptos !== '') {
      queryParams = queryParams.set('distribucion', conceptos);
    }
    return this._http.post<RespuestaPeticion>(`${environment.host}/distribucion/verificar`, data, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  AgregarArchivo(data: FormData, conceptos: string) {
    let queryParams = new HttpParams();
    if (conceptos !== '') {
      queryParams = queryParams.set('distribucion', conceptos);
    }
    return this._http.post<RespuestaPeticion>(`${environment.host}/distribucion/agregar`, data, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  AgregaDistribucion(distribucion: DistribucionUnidadMed, cantidad: string, contrato: string, clv_dele: string) {
    let queryParams = new HttpParams();
    if (cantidad !== '') {
      queryParams = queryParams.set('CantidadConcepto', cantidad);
    }
    if (contrato !== '') {
      queryParams = queryParams.set('NumContrato', contrato);
    }
    if (clv_dele !== '') {
      queryParams = queryParams.set('ClvDele', clv_dele);
    }
    return this._http.post<RespuestaPeticion>(`${environment.host}/distribucion`, distribucion, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  EliminarDistribucion(id_distribucion: string, num_contrato: string, clv_dele: string) {
    const queryParams = new HttpParams().set('id', id_distribucion).set('n_contrato', num_contrato).set('clv', clv_dele);
    return this._http.delete<RespuestaPeticion>(`${environment.host}/distribucion`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  ActualizarDistribucion(obj: DistribucionUnidadMed, num_contrato: string, clv_dele: string) {
    const queryParams = new HttpParams().set('NumContrato', num_contrato).set('ClvDele', clv_dele);
    return this._http.put<RespuestaPeticion>(`${environment.host}/distribucion`, obj, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
