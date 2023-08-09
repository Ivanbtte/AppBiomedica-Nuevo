import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {DatosUsuario, Passwords} from '../interfaces/usuario.interface';
import {environment} from '../../../environments/environment';
import {RespuestaPeticion} from '../../core/models/estructuras_generales';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class UsuarioService {

  constructor(
    private http: HttpClient,
  ) {
  }

  AddNewUser(obj: DatosUsuario) {
    return this.http.post<RespuestaPeticion>(`${environment.host}/usuarios`, obj).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  ActualizarUsuario(ID: string, obj: Passwords) {
    return this.http.put<RespuestaPeticion>(`${environment.host}/usuarios/${ID}`, obj).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  unidadUsuario(usuario_id: string) {
    const queryParams = new HttpParams().set('usuario_id', usuario_id);
    return this.http.get<RespuestaPeticion>(`${environment.host}/usuarios/unidad`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  unidades(delegacion: string) {
    const queryParams = new HttpParams().set('delegacion', delegacion);
    return this.http.get<RespuestaPeticion>(`${environment.host}/unidades`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  ConsultarUsuarios(estado: string) {
    const httpParams = new HttpParams().set('activos', estado);
    return this.http.get<RespuestaPeticion>(`${environment.host}/usuarios`, {params: httpParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  ValidarUsuario(ID: string, estado: string, obj: any) {
    const queryParams = new HttpParams().set('id', ID).set('valida', estado);
    return this.http.put<RespuestaPeticion>(`${environment.host}/usuarios`, obj, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  rechazarUsuario(IdUsuario: string, idPersona: string, mensaje: string) {
    let queryParams = new HttpParams();
    queryParams = queryParams.set('idUsuario', IdUsuario);
    queryParams = queryParams.set('idPersona', idPersona);
    queryParams = queryParams.set('mensaje', mensaje);
    return this.http.delete<RespuestaPeticion>(`${environment.host}/usuarios`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  ConsultUsersByDepartment(departmentId, userId: string) {
    const queryParams = new HttpParams().set('departmentId', departmentId).set('userId', userId);
    return this.http.get<RespuestaPeticion>(`${environment.host}/usuarios/department`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  ConsultUSerByPosition(userId: string) {
    const queryParams = new HttpParams().set('userId', userId);
    return this.http.get<RespuestaPeticion>(`${environment.host}/usuarios/puesto`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  CheckResetPassword(email, dateBirth: string) {
    const queryParams = new HttpParams().set('email', email).set('date', dateBirth);
    return this.http.get<RespuestaPeticion>(`${environment.host}/usuarios/reset`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  ResetPassword(id: string, obj: Partial<Passwords>) {
    const queryParams = new HttpParams().set('idUser', id);
    return this.http.put<RespuestaPeticion>(`${environment.host}/usuarios/reset`, obj, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
