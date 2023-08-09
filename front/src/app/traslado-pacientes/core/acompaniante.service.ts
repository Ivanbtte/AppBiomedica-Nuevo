import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { RespuestaPeticion } from '../../core/models/estructuras_generales';
import { environment } from 'src/environments/environment';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class AcompanianteService {

  constructor(private httpClient:HttpClient) { }
  ConsultarUltimoAcompaniante(unidadMedId: string) {
    const queryParams = new HttpParams().set('unidadMedId', unidadMedId);
    return this.httpClient.get<RespuestaPeticion>(`${environment.host}/acompaniante`, {params: queryParams}).pipe(map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
