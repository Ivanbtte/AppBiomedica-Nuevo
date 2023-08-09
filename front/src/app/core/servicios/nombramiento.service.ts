import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {RespuestaPeticion} from '../models/estructuras_generales';
import {environment} from '../../../environments/environment';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class NombramientoService {

  constructor(
    private http: HttpClient,
  ) {
  }

  nombramiento(ID: string) {
    return this.http.get<RespuestaPeticion>(`${environment.host}/nombramiento/${ID}`).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
