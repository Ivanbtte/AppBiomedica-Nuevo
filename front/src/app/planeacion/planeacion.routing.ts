import {Routes} from '@angular/router';
import {AuthGuard} from '../shared/modulos/auth.guard';
import {PermisoSolicitudGuard} from '../shared/modulos/permiso-solicitud.guard';
import {EnfermeriaFondoFijoComponent} from './con-cuadro-basico/enfermeria-fondo-fijo.component';
import {SinCuadroBasicoComponent} from './sin-cuadro-basico/sin-cuadro-basico.component';

export const PlaneacionRoutes: Routes = [
  {
    path: '',
    children: [
      {
        path: 'enfermeria/cuadrosi',
        component: EnfermeriaFondoFijoComponent
      },
      {
        path: 'enfermeria/cuadrono',
        component: SinCuadroBasicoComponent
      }
    ],
    canActivate: [AuthGuard, PermisoSolicitudGuard]
  }
];
