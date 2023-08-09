import {Injectable} from '@angular/core';
import {RespuestaPeticion} from '../models/estructuras_generales';
import {environment} from '../../../environments/environment';
import {map} from 'rxjs/operators';
import {HttpClient} from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class AreasService {

  constructor(private httpClient: HttpClient) {
  }

  consultAreas() {
    return this.httpClient.get<RespuestaPeticion>(`${environment.host}/areas`,).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
