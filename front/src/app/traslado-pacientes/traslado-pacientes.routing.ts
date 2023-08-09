import {Routes} from '@angular/router';
import {AuthGuard} from '../shared/modulos/auth.guard';
import {PermisoSolicitudTrasladoGuard} from '../shared/modulos/permiso-solicitud-traslado.guard';
import {SolicitudTrasladoComponent} from './components/solicitud-traslado/solicitud-traslado.component';
import {SolicitudesRealizadasComponent} from './components/solicitudes-realizadas/solicitudes-realizadas.component';
import {SolicitudesCanceladasComponent} from './components/solicitudes-canceladas/solicitudes-canceladas.component';
import {FirmasComponent} from './components/firmas/firmas.component';
import {TransferManagerComponent} from './components/transfer-manager/transfer-manager.component';
import {PermisoAdministradorTrasladosGuard} from '../shared/modulos/permiso-administrador-traslados.guard';
import {ViewPermitByUserComponent} from './components/view-permit-by-user/view-permit-by-user.component';
import { AdministrarUsuariosComponent } from './components/administrar-usuarios/administrar-usuarios.component';
import { CargarFtpComponent } from './components/cargar-ftp/cargar-ftp.component';
import { TableroTrasladosComponent } from './components/tablero-traslados/tablero-traslados.component';
import { DetallePorContratoComponent } from './components/detalle-por-contrato/detalle-por-contrato.component';

export const TrasladosRoutes: Routes = [
  {
    path: '',
    children: [
      {path: 'solicitud', component: SolicitudTrasladoComponent, canActivate: [PermisoSolicitudTrasladoGuard]},
      {path: 'realizadas', component: SolicitudesRealizadasComponent, canActivate: [PermisoSolicitudTrasladoGuard]},
      {path: 'cancelar', component: SolicitudesCanceladasComponent, canActivate: [PermisoSolicitudTrasladoGuard]},
      {path: 'firma', component: FirmasComponent, canActivate: [PermisoAdministradorTrasladosGuard]},
      {path: 'manager', component: TransferManagerComponent, canActivate: [PermisoAdministradorTrasladosGuard]},
      {path: 'permit/:id/:unitMedic', component: ViewPermitByUserComponent, canActivate: [PermisoAdministradorTrasladosGuard]},
      {path: 'administradores', component: AdministrarUsuariosComponent},
      {path: 'ftp', component:CargarFtpComponent},
      {path: 'tablero', component:TableroTrasladosComponent},
      {path:'detalle/contrato',component:DetallePorContratoComponent}

    ],
    canActivate: [AuthGuard]
  }
];
