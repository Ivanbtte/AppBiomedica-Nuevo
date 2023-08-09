import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
import {RespuestaPeticion} from '../../core/models/estructuras_generales';
import {environment} from '../../../environments/environment';
import {map} from 'rxjs/operators';
import {Fecha} from '../../core/models/fechaLimite.interface';
import {Presupuesto} from '../../core/models/presupuesto';

@Injectable({
  providedIn: 'root'
})
export class PresupuestoService {

  constructor(
    private _http: HttpClient
  ) {
  }

  ConsultaPresupuesto() {
    return this._http.get<RespuestaPeticion> (`${environment.host}/presupuesto`).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  AgregarPresupuesto(obj: Presupuesto) {
    return this._http.post<RespuestaPeticion> (`${environment.host}/presupuesto`, obj).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  ActualizarPresupuesto(obj: Presupuesto) {
    return this._http.put<RespuestaPeticion> (`${environment.host}/presupuesto`, obj).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
