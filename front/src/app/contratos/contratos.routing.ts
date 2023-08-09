import {Routes} from '@angular/router';
import {AuthGuard} from '../shared/modulos/auth.guard';
import {AdministrarContratosComponent} from './administrar-contratos/administrar-contratos.component';
import {SeguimientoContratosComponent} from './seguimiento-contratos/seguimiento-contratos.component';
import {DetalleContratoComponent} from './detalle-contrato/detalle-contrato.component';
import {AgregarConceptosContratoComponent} from './agregar-conceptos-contrato/agregar-conceptos-contrato.component';
import {AgregarUnContratoComponent} from '../shared/Dialogs/Contratos/agregar-un-contrato/agregar-un-contrato.component';
import {FianzaComponent} from './fianza/fianza.component';
import {PermisoAdministradorContratosGuard} from '../shared/modulos/permiso-administrador-contratos.guard';

export const ContratosRoutes: Routes = [
  {
    path: '',
    children: [
      {
        path: 'administrar',
        component: AdministrarContratosComponent,
        // canActivate: [PermisoAdministradorContratosGuard]
      },
      {
        path: 'seguimiento',
        component: SeguimientoContratosComponent
      },
      {
        path: 'detalle/:contrato_id',
        component: DetalleContratoComponent
      },
      {
        path: 'agregar/:contrato_id',
        component: AgregarConceptosContratoComponent
      },
      {
        path: 'agregar/contrato:numero_contrato',
        component: AgregarUnContratoComponent
      },
      {
        path: 'fianza/:numero_contrato',
        component: FianzaComponent,
        canActivate: [PermisoAdministradorContratosGuard]
      },
      {
        path: 'fianza',
        component: FianzaComponent,
        canActivate: [PermisoAdministradorContratosGuard]
      }
    ],
    canActivate: [AuthGuard]
  }
];
