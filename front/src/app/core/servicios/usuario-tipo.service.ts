import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {RespuestaPeticion} from '../models/estructuras_generales';
import {environment} from '../../../environments/environment';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class UsuarioTipoService {

  constructor(
    private _http: HttpClient
  ) {
  }

  checkUserType() {
    return this._http.get<RespuestaPeticion>(`${environment.host}/usuario_tipo/sn`,).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
