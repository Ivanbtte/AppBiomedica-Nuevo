import {Injectable} from '@angular/core';
import {RespuestaPeticion} from '../../core/models/estructuras_generales';
import {environment} from '../../../environments/environment';
import {map} from 'rxjs/operators';
import {HttpClient, HttpParams} from '@angular/common/http';
import {Personas} from '../interfaces/usuario.interface';


@Injectable({
  providedIn: 'root'
})
export class PersonaService {

  constructor(
    private http: HttpClient,
  ) {
  }

  AgregarPersona(obj: Personas) {
    return this.http.post<RespuestaPeticion>(`${environment.host}/personas`, obj).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  getAllPersona(num_pagina: string, num_registro: string) {
    const queryParams = new HttpParams().set('pagina', num_pagina).set('num_registros', num_registro);
    return this.http.get<RespuestaPeticion>(`${environment.host}/personas`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
          return res;
        }
      )
    );
  }

  getPersona(ID: string) {
    return this.http.get<RespuestaPeticion>(`${environment.host}/personas/${ID}`).pipe(
      map((res: RespuestaPeticion) => {
          return res;
        }
      )
    );
  }

  putPersona(obj: Personas, ID: string) {
    return this.http.put<RespuestaPeticion>(`${environment.host}/personas/${ID}`, obj).pipe(
      map((res: RespuestaPeticion) => {
          return res;
        }
      )
    );
  }

  checkEnrollment(enrollment: string) {
    const queryParams = new HttpParams().set('enrollment', enrollment);
    return this.http.get(`${environment.host}/person/check/available`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
          return res;
        }
      )
    );
  }
}
