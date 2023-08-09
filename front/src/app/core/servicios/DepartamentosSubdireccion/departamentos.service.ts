import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {RespuestaPeticion} from '../../models/estructuras_generales';
import {environment} from '../../../../environments/environment';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class DepartamentosService {

  constructor(private httpClient: HttpClient) {
  }

  ConsultOrganizationUnitMed(unitMedId,clvDele:string){
    let queryParams = new HttpParams()
    if (unitMedId !== ''){
      queryParams =queryParams.set('unitMedId', unitMedId)
    }
    if (clvDele !== ''){
      queryParams =queryParams.set('clvDele', clvDele)
    }
    return this.httpClient.get<RespuestaPeticion>(`${environment.host}/departamentos/sn/prueba`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
