import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {RespuestaPeticion} from '../../core/models/estructuras_generales';
import {environment} from '../../../environments/environment';
import {map} from 'rxjs/operators';
import { PuestoOrganigramaRole } from '../interfaces/puesto.role.interface';

@Injectable({
  providedIn: 'root'
})
export class PuestosService {

  constructor(
    private _http: HttpClient
  ) {
  }

  puestos() {
    return this._http.get<RespuestaPeticion>(`${environment.host}/puesto`).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  /* ******  Servicio para consultar puestos disponibles por departamento jefatura etc  ******   */
  ConsultarPuestosDispo(departamento_id: string, tipo_usuario: string, unidad_med: string) {
    let queryParam = new HttpParams().set('departamento_id', departamento_id);
    if (tipo_usuario !== '') {
      queryParam = queryParam.set('tipo', tipo_usuario);
    }
    if (unidad_med !== '') {
      queryParam = queryParam.set('unidad_id', unidad_med);
    }
    return this._http.get<RespuestaPeticion>(`${environment.host}/puesto/disponibles`, {params: queryParam}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  ConsultPositionOrganizations(departmentUnitId, departmentDelegationId: string) {
    let queryParams = new HttpParams();
    if (departmentUnitId !== '') {
      queryParams = queryParams.set('departmentUnitId', departmentUnitId);
    }
    if (departmentDelegationId !== '') {
      queryParams = queryParams.set('departmentDelegationId', departmentDelegationId);
    }
    return this._http.get<RespuestaPeticion>(`${environment.host}/puesto/final`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  AgregaPuestoRole(obj:Partial<PuestoOrganigramaRole>){
    return this._http.post<RespuestaPeticion>(`${environment.host}/puesto/role`, obj).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

}
