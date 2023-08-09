import {Routes} from '@angular/router';
import {AuthGuard} from '../shared/modulos/auth.guard';
import {InicioComponent} from './inicio/inicio.component';

export const InicioRotues: Routes = [
  {
    path: '',
    children: [
      {
        path: '',
        redirectTo: 'inicio',
        pathMatch: 'full'
      },
      {
        path: '',
        component: InicioComponent
      }
    ],
    canActivate: [AuthGuard]
  }
];
