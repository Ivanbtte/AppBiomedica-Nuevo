import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {RespuestaPeticion} from '../../models/estructuras_generales';
import {environment} from '../../../../environments/environment';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class ArchivosContratosService {

  constructor(
    private _http: HttpClient
  ) {
  }

  /* ******  Servicio para consultar las carpetas de un contrato  ******   */
  ConsultarCarpetas(path: string) {
    const queryParams = new HttpParams().set('path', path);
    return this._http.get<RespuestaPeticion>(`${environment.host}/archivos/contrato`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  /* ******  Servicio para consultar archivos dentro de una carpeta  ******   */
  ConsultarArchivos(numero_contrato: string, Nombre: string) {
    const queryParams = new HttpParams().set('num_contrato', numero_contrato).set('path', Nombre);
    return this._http.get<RespuestaPeticion>(`${environment.host}/archivos/contrato/archivos`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  /* ******  Servicio para visualizar un pdf  ******   */
  DescargarArchivo(path: string) {
    return this._http.get(`${environment.host}/${path}`, {responseType: 'blob'});
  }

  /* ******  Servicio para subir unh nuevo archivo  ******   */
  SubirArchivo(data: FormData, path: string) {
    const query_params = new HttpParams().set('path', path);
    return this._http.post<RespuestaPeticion>(`${environment.host}/${{path}}`, data, {params: query_params}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
