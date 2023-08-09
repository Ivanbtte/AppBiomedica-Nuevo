import {Routes} from '@angular/router';
import {AuthGuard} from '../shared/modulos/auth.guard';
import {PermisoBiomedicoGuard} from '../shared/modulos/permisoBiomedico.guard';
import {FechaLimiteComponent} from './fecha-limite/fecha-limite.component';
import {PresupuestoComponent} from './presupuesto/presupuesto.component';

export const ConfiguracionesRoutes: Routes = [
  {
    path: '',
    children: [
      {
        path: 'fecha',
        component: FechaLimiteComponent,
      },
      {
        path: 'presupuesto',
        component: PresupuestoComponent
      }
    ],
    canActivate: [AuthGuard, PermisoBiomedicoGuard]
  }
];
