import {Injectable} from '@angular/core';
import {Credenciales} from '../interfaces/usuario.interface';
import {HttpClient} from '@angular/common/http';
import {Router} from '@angular/router';
import {RespuestaPeticion} from '../../core/models/estructuras_generales';
import {environment} from '../../../environments/environment';
import {map} from 'rxjs/operators';
import {JwtHelperService} from '@auth0/angular-jwt';
import Swal from 'sweetalert2';

@Injectable({
  providedIn: 'root'
})
export class LoginService {
  constructor(
    private http: HttpClient,
    private _router: Router
  ) {
  }

  login(obj: Credenciales) {
    return this.http.post<RespuestaPeticion>(`${environment.host}/login`, obj).pipe(
      map((res: RespuestaPeticion) => {
          return res;
        }
      )
    );
  }

  estaLogeado() {
    return !!localStorage.getItem('token');
  }

  getToken(): string {
    return <string> localStorage.getItem('token');
  }

  isExpiret(): boolean {
    const helper = new JwtHelperService();
    const isExpired = helper.isTokenExpired(this.getToken());
    if (isExpired) {
      localStorage.clear();
      Swal.fire({
        position: 'center',
        icon: 'error',
        title: 'Inicia Sesión',
        showConfirmButton: true,
      });
      this._router.navigate(['authentication/login']);
    }
    return isExpired;
  }

  logoutUser() {
    localStorage.clear();
    this._router.navigate(['authentication/login']);
  }

  sesionCaducada() {
    localStorage.clear();
    this._router.navigate(['authentication/login']);
    Swal.fire({
      icon: 'error',
      title: 'Tu Sesión ha Caducado',
      showConfirmButton: true,
    });
  }
}

