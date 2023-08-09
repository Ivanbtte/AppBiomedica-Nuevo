import {Injectable} from '@angular/core';
import {RespuestaPeticion} from '../models/estructuras_generales';
import {environment} from '../../../environments/environment';
import {map} from 'rxjs/operators';
import {HttpClient} from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class EstatusService {

  constructor(
    private http: HttpClient
  ) {
  }

  ConsultaEstatus() {
    return this.http.get<RespuestaPeticion>(`${environment.host}/estatus`).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
