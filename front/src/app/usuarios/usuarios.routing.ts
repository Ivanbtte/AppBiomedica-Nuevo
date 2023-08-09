import {Routes} from '@angular/router';
import {AuthGuard} from '../shared/modulos/auth.guard';
import {PermisoBiomedicoGuard} from '../shared/modulos/permisoBiomedico.guard';
import {ValidarUsuariosComponent} from './validar-usuarios/validar-usuarios.component';
import {BloquearUsuariosComponent} from './bloquear-usuarios/bloquear-usuarios.component';
export const UsuariosRoutes: Routes = [
  {
    path: '',
    children: [
      {
        path: 'validar',
        component: ValidarUsuariosComponent,
      },
      {
        path: 'bloquear',
        component: BloquearUsuariosComponent,
      }
    ],
    canActivate: [AuthGuard]
    // canActivate: [AuthGuard, PermisoBiomedicoGuard]
  }
];
