import {Injectable} from '@angular/core';
import {RespuestaPeticion} from '../models/estructuras_generales';
import {environment} from '../../../environments/environment';
import {map} from 'rxjs/operators';
import {HttpClient, HttpParams} from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class ServiciosUnidadesService {

  constructor(
    private http: HttpClient,
  ) {
  }

  serviciosUnidades(ID: string) {
    const queryParams = new HttpParams().set('id', ID);
    return this.http.get<RespuestaPeticion>(`${environment.host}/servicios_unidad`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
