import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { RespuestaPeticion } from '../../core/models/estructuras_generales';
import { environment } from 'src/environments/environment';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class AdministradoresTrasladosService {

  constructor(private httpClient: HttpClient) { }

  ConsultAdministradoresUnidad(clvDele, roleId: string) {
    const queryParams = new HttpParams().set('clvDele', clvDele).set('roleId', roleId);
    return this.httpClient.get<RespuestaPeticion>(`${environment.host}/traslados/administradores`, {params: queryParams}).pipe(map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
