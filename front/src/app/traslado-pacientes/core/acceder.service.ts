import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { map } from 'rxjs/operators';
import { RespuestaPeticion } from '../../core/models/estructuras_generales';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class AccederService {

  constructor(private httpClient: HttpClient) { }

  ConsultarDerechoabiente(nss: string) {
    const queryParams = new HttpParams().set('nss', nss);
    return this.httpClient.get<RespuestaPeticion>(`${environment.host}/acceder`, { params: queryParams }).pipe(map((res: RespuestaPeticion) => {
      return res;
    })
    );
  }
}
