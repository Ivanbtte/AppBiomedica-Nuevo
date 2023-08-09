import {Routes} from '@angular/router';
import {FullComponent} from './layouts/full/full.component';
import {AppBlankComponent} from './layouts/blank/blank.component';
import {PermisoBiomedicoGuard} from './shared/modulos/permisoBiomedico.guard';
import { ServiciosSubrogadosModule } from './servicios-subrogados/servicios-subrogados.module';

export const AppRoutes: Routes = [
  {
    path: '',
    component: FullComponent,
    children: [
      {
        path: '',
        pathMatch: 'full',
        redirectTo: 'inicio'
      },
      {
        path: 'inicio',
        loadChildren: () => import('./inicio/inicio.module').then(m => m.InicioModule)
      },
      {
        path: 'area',
        loadChildren: () => import('./area-servicios/area-servicios.module').then(m => m.AreaServiciosModule)
      },
      {
        path: 'catalogos',
        loadChildren: () => import('./catalogos/catalogos.module').then(m => m.CatalogosModule)
      },
      {
        path: 'configuraciones',
        loadChildren: () => import('./configuraciones/configuraciones.module').then(m => m.ConfiguracionesModule)
      },
      {
        path: 'contratos',
        loadChildren: () => import('./contratos/contratos.module').then(m => m.ContratosModule)
      },
      {
        path: 'cuenta',
        loadChildren: () => import('./cuenta/cuenta.module').then(m => m.CuentaModule)
      },
      {
        path: 'listas',
        loadChildren: () => import('./listas-interes/listas-interes.module').then(m => m.ListasInteresModule)
      },
      {
        path: 'planeacion',
        loadChildren: () => import('./planeacion/planeacion.module').then(m => m.PlaneacionModule)
      },
      {
        path: 'proveedores',
        loadChildren: () => import('./proveedores/proveedores.module').then(m => m.ProveedoresModule)
      },
      {
        path: 'solicitudes',
        loadChildren: () => import('./solicitudes/solicitudes.module').then(m => m.SolicitudesModule)
      },
      {
        path: 'usuarios',
        loadChildren: () => import('./usuarios/usuarios.module').then(m => m.UsuariosModule),
        // canActivate: [PermisoBiomedicoGuard]
      },
      {
        path: 'traslado',
        loadChildren: () => import('./traslado-pacientes/traslado-pacientes.module').then(m => m.TrasladoPacientesModule)
      },
      {
        path: 'requerimientos',
        loadChildren: () => import('./validacion-solicitudes/validacion-solicitudes.module').then(m => m.ValidacionSolicitudesModule)
      },
      {
        path: 'subrogados',
        loadChildren: () => import('./servicios-subrogados/servicios-subrogados.module').then(m => m.ServiciosSubrogadosModule)
      },
    ]
  },
  {
    path: '',
    component: AppBlankComponent,
    children: [
      {
        path: 'authentication',
        loadChildren: () => import('./authentication/authentication.module').then(m => m.AuthenticationModule)
      }
    ]
  },
  {
    path: '**',
    redirectTo: 'authentication/404'
  }
];
