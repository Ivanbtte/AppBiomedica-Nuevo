import {Injectable} from '@angular/core';
import {RespuestaPeticion} from '../../core/models/estructuras_generales';
import {environment} from '../../../environments/environment';
import {map} from 'rxjs/operators';
import {HttpClient, HttpParams} from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class FiltrosService {

  constructor(
    private http: HttpClient
  ) {
  }

  getTemas() {
    return this.http.get<RespuestaPeticion>(`${environment.host}/temas`).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  getGrupoTemas(todos: string, temas: string) {
    let paprams1 = new HttpParams().set('all', todos);
    if (temas !== '') {
      paprams1 = paprams1.set('tema', temas);
    }
    return this.http.get<RespuestaPeticion>(`${environment.host}/grupo/temas`, {params: paprams1}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}

