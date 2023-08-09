import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {RespuestaPeticion} from '../models/estructuras_generales';
import {environment} from '../../../environments/environment';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class EntidadFederativaService {

  constructor(
    private http: HttpClient
  ) {
  }

  ConsultaEntidades() {
    return this.http.get<RespuestaPeticion>(`${environment.host}/estados`).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
