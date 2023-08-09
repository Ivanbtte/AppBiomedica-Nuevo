import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {Conceptos} from '../../models/solicitudes.interface';
import {RespuestaPeticion} from '../../models/estructuras_generales';
import {environment} from '../../../../environments/environment';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class ClavesService {

  constructor(
    private http: HttpClient
  ) {
  }

  AgregarClave(obj: Conceptos, unidadId: string) {
    const parametro = new HttpParams().set('unidadId', unidadId);
    return this.http.post<RespuestaPeticion>(`${environment.host}/clave`, obj, {params: parametro}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  ConsultarUnaClaves(Id: string) {
    return this.http.get<RespuestaPeticion>(`${environment.host}/clave/${Id}`).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  ConsultarClaves(Id: string, clave: string, unidad: string, final: string, servicio: string) {
    let parametro = new HttpParams();
    if (Id !== '') {
      parametro = parametro.set('id', Id);
    }
    if (clave !== '') {
      parametro = parametro.set('clave', clave);
    }
    if (unidad !== '') {
      parametro = parametro.set('unidad', unidad);
    }
    if (final !== '') {
      parametro = parametro.set('final', final);
    }
    if (servicio !== '') {
      parametro = parametro.set('servicio', servicio);
    }
    return this.http.get<RespuestaPeticion>(`${environment.host}/clave`, {params: parametro}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  ActualizarClave(Id: string, Comentario: string, CantidadApro: string, UsuarioId: string, aprobar: string, obj: Conceptos) {
    let parametro = new HttpParams().set('comentario', Comentario).set('id', Id);
    if (CantidadApro !== '') {
      parametro = parametro.set('cantidad_apro', CantidadApro);
    }
    if (UsuarioId !== '') {
      parametro = parametro.set('usuarioId', UsuarioId);
    }
    if (aprobar !== '') {
      parametro = parametro.set('aprobar', aprobar);
    }
    return this.http.put<RespuestaPeticion>(`${environment.host}/clave`, obj, {params: parametro}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  EliminarClave(Id: string, Comentario: string) {
    const parametro = new HttpParams().set('comentario', Comentario);
    return this.http.delete<RespuestaPeticion>(`${environment.host}/clave/${Id}`, {params: parametro}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
