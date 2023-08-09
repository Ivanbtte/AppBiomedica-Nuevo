import {Injectable} from '@angular/core';
import {CanActivate, Router} from '@angular/router';
import {LoginService} from '../../authentication/servicios/login.service';
import Swal from 'sweetalert2';
import {Observable} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthGuard implements CanActivate {
  constructor(private _authService: LoginService,
              private _router: Router
  ) {
  }

  canActivate(): Observable<boolean> | Promise<boolean> | boolean {
    if (!this._authService.estaLogeado()) {
      this._router.navigate(['authentication/login']).then(r => {console.log(r)});
      return false;
    }
    if (this._authService.isExpiret()) {
      this._router.navigate(['authentication/login']).then(r=>{console.log(r)});
      Swal.fire({
        position: 'center',
        title: 'Tu Sesión ha Caducado',
        icon: 'error',
        showConfirmButton: true,
      });
      localStorage.clear();
      return false;
    }
    if (this._authService.estaLogeado() && !this._authService.isExpiret()) {
      return true;
    } else {
      Swal.fire({
        position: 'center',
        title: 'Inicia Sesion',
        icon: 'error',
        showConfirmButton: true,
      });
      this._router.navigate(['authentication/login']).then(r=>{});
      return false;
    }
  }
}
