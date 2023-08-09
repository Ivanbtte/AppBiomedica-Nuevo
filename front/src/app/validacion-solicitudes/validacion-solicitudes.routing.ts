import {Routes} from '@angular/router';
import {AuthGuard} from '../shared/modulos/auth.guard';
import {PermisoBiomedicoGuard} from '../shared/modulos/permisoBiomedico.guard';
import {ValidarComponent} from './validar/validar.component';
import {VerComponent} from './ver/ver.component';
import {DetalleValidarComponent} from './detalle-validar/detalle-validar.component';
import {SolicitudFinalComponent} from './solicitud-final/solicitud-final.component';
export const ValidacionRoutes: Routes = [
  {
    path: '',
    children: [
      {
        path: 'revision',
        component: ValidarComponent,
      },
      {
        path: 'validado',
        component: VerComponent
      },
      {
        path: 'final',
        component: SolicitudFinalComponent
      },
      {
        path: 'revision/validar/:id',
        component: DetalleValidarComponent
      },
    ],
    canActivate: [AuthGuard, PermisoBiomedicoGuard]
  }
];
