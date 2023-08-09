import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {RespuestaPeticion} from '../../models/estructuras_generales';
import {environment} from '../../../../environments/environment';
import {map} from 'rxjs/operators';
import {Solicitud} from '../../models/solicitudes.interface';

@Injectable({
  providedIn: 'root'
})
export class SolicitudService {

  constructor(
    private http: HttpClient) {
  }

  AgregarSolicitud(obj: Solicitud) {
    return this.http.post<RespuestaPeticion>(`${environment.host}/solicitud`, obj).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  Actualizarolicitud(Id: string, comentario: string, total: string, estatus: string, usuarioId: string, obj: Solicitud) {
    let parametro = new HttpParams().set('comentario', comentario).set('id', Id);
    if (total !== '') {
      console.log('esto trae total', total);
      parametro = parametro.set('total', total);
    }
    if (estatus !== '') {
      parametro = parametro.set('estatus', estatus).set('usuarioId', usuarioId);
    }
    return this.http.put<RespuestaPeticion>(`${environment.host}/solicitud`, obj, {params: parametro}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  ConsultarTodasSolicitudes(Id: string, valido: string) {
    const parametro = new HttpParams().set('id', Id).set('validadas', valido);
    return this.http.get<RespuestaPeticion>(`${environment.host}/solicitud`, {params: parametro}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  /* *****  Este servicio consulta si un usuario ya tiene una solicitud  *****   */
  ConsultarSolicitud(Id: string, Tipo: string, IdSolicitud: string) {
    let parametro = new HttpParams();
    if (Id !== '') {
      parametro = parametro.set('id', Id);
    }
    if (Tipo !== '') {
      parametro = parametro.set('tipo', Tipo);
    }
    if (IdSolicitud !== '') {
      parametro = parametro.set('solicitud', IdSolicitud);
    }
    return this.http.get<RespuestaPeticion>(`${environment.host}/solicitud/detalle`, {params: parametro}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
