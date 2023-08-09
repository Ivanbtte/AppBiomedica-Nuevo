import {Routes} from '@angular/router';
import {AuthGuard} from '../shared/modulos/auth.guard';
import {PermisoSolicitudGuard} from '../shared/modulos/permiso-solicitud.guard';
import {EqCocinaComponent} from './eq-cocina/eq-cocina.component';
import {EqMedicoComponent} from './eq-medico/eq-medico.component';
import {InstrumentalComponent} from './instrumental/instrumental.component';
import {EditarSolicComponent} from './editar-solic/editar-solic.component';
import {SeguimientoSolicComponent} from './seguimiento-solic/seguimiento-solic.component';
import {ConsumoComponent} from './consumo/consumo.component';
import {DetalleEditarComponent} from './detalle-editar/detalle-editar.component';

export const SolicitudesRoutes: Routes = [
  {
    path: '',
    children: [
      {
        path: 'nueva/consumo',
        component: ConsumoComponent,
      },
      {
        path: 'nueva/cocina',
        component: EqCocinaComponent,
      },
      {
        path: 'nueva/medico',
        component: EqMedicoComponent,
      },
      {
        path: 'nueva/instrumental',
        component: InstrumentalComponent,
      },
      {
        path: 'editar',
        component: EditarSolicComponent,
      },
      {
        path: 'editar/folio/:id',
        component: DetalleEditarComponent
      },
      {
        path: 'seguimiento',
        component: SeguimientoSolicComponent,
      },
    ],
    canActivate: [AuthGuard, PermisoSolicitudGuard]
  }
];
