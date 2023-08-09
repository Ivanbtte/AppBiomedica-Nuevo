import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {Observable} from 'rxjs';
import {RespuestaPeticion} from '../../core/models/estructuras_generales';
import {environment} from '../../../environments/environment';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class FiltrosSaiService {

  constructor(
    private http: HttpClient
  ) {
  }

  consultarFiltrosGrupos(grupo: string, generico: string, todos_genericos: boolean, todos_especificos: boolean) {
    let queryParams = new HttpParams();
    if (grupo !== '') {
      queryParams = queryParams.set('grupo', grupo);
    }
    if (generico !== '') {
      queryParams = queryParams.set('generico', generico);
    }
    if (todos_genericos) {
      const valorParse: string = String(todos_genericos);
      queryParams = queryParams.set('todos_genericos', valorParse);
    }
    if (todos_especificos) {
      const valorParse: string = String(todos_especificos);
      queryParams = queryParams.set('todos_especificos', valorParse);
    }
    return this.http.get<RespuestaPeticion>(`${environment.host}/sai/filtros`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
