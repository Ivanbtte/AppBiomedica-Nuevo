import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {RespuestaPeticion} from '../../core/models/estructuras_generales';
import {environment} from '../../../environments/environment';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class ModulesService {

  constructor(private httpClient: HttpClient) {
  }

  ConsultModules(menuParent: string) {
    const queryParams = new HttpParams().set('menuParent', menuParent);
    return this.httpClient.get<RespuestaPeticion>(`${environment.host}/menu/module`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  ConsultActionsByModule(moduleId, userId: string) {
    const queryParams = new HttpParams().set('moduleId', moduleId).set('userId', userId);
    return this.httpClient.get<RespuestaPeticion>(`${environment.host}/module/actions`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
  ConsultarPermisosUnidad(unitId, moduleId: string){
    const queryParams = new HttpParams().set('unitId',unitId).set('moduleId',moduleId);
    return this.httpClient.get<RespuestaPeticion>(`${environment.host}/acciones`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
