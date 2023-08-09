import {Routes} from '@angular/router';
import {AuthGuard} from '../shared/modulos/auth.guard';
import {InfoCuentaComponent} from './info-cuenta/info-cuenta.component';
import {ConfiguracionesComponent} from './configuraciones/configuraciones.component';

export const CuentaRoutes: Routes = [
  {
    path: '',
    children: [
      {
        path: 'info',
        component: InfoCuentaComponent
      },
      {
        path: 'configuracion',
        component: ConfiguracionesComponent
      }
    ],
    canActivate: [AuthGuard]
  }
];
