import {Routes} from '@angular/router';
import {AuthGuard} from '../shared/modulos/auth.guard';
import {PermisoSolicitudGuard} from '../shared/modulos/permiso-solicitud.guard';
import {AgregarAreaServicioComponent} from './agregar-area-servicio/agregar-area-servicio.component';
export const AreaServicioRoutes: Routes = [
  {
    path: '',
    children: [
      {
        path: 'nueva',
        component: AgregarAreaServicioComponent
      }
    ],
    canActivate: [AuthGuard, PermisoSolicitudGuard]
  }
];
