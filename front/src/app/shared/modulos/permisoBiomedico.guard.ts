import {Injectable} from '@angular/core';
import {CanActivate, Router} from '@angular/router';
import {Observable} from 'rxjs';

@Injectable({
  providedIn: 'root'
})

export class PermisoBiomedicoGuard implements CanActivate {
  public permisos: string[] = [];

  constructor(
    private _router: Router
  ) {
  }

  canActivate(): Observable<boolean> | Promise<boolean> | boolean {
    return this.havePermission();
  }

  havePermission(): boolean {
    this.permisos = JSON.parse(<string> localStorage.getItem('permisos'));
    console.log(this.permisos);
    const have = this.permisos.filter(permit => permit === 'Jefe Biomedico');
    return have.length !== 0;
  }
}
