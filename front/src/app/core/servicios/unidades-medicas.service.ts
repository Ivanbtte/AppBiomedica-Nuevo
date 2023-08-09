import {Injectable} from '@angular/core';
import {RespuestaPeticion} from '../models/estructuras_generales';
import {environment} from '../../../environments/environment';
import {map} from 'rxjs/operators';
import {HttpClient} from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class UnidadesMedicasService {

  constructor(
    private http: HttpClient,
  ) {
  }

  unaUnidad(ID: string) {
    return this.http.get<RespuestaPeticion>(`${environment.host}/unidades/${ID}`).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
