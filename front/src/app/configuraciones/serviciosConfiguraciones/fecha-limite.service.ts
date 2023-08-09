import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {RespuestaPeticion} from '../../core/models/estructuras_generales';
import {environment} from '../../../environments/environment';
import {map} from 'rxjs/operators';
import {Fecha} from '../../core/models/fechaLimite.interface';

@Injectable({
  providedIn: 'root'
})
export class FechaLimiteService {

  constructor(
    private _http: HttpClient
  ) {
  }

  ConsultaFecha() {
    return this._http.get<RespuestaPeticion>(`${environment.host}/fecha`).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  AgregarFecha(obj: Fecha) {
    return this._http.post<RespuestaPeticion>(`${environment.host}/fecha`, obj).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  ActualizarFecha(obj: Fecha) {
    return this._http.put<RespuestaPeticion>(`${environment.host}/fecha`, obj).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
