import { Routes } from '@angular/router';
import { AuthGuard } from '../shared/modulos/auth.guard';
import { SolicitudSubrogadoComponent } from './solicitud-subrogado/solicitud-subrogado.component';
import { AdministrarUsuariosSubComponent } from './administrar-usuarios-sub/administrar-usuarios-sub.component';

export const ServiciosSubrogadosRoutes: Routes = [
  {
    path: '',
    children: [
      {path: 'solicitud', component: SolicitudSubrogadoComponent},
      {path: 'administrar', component: AdministrarUsuariosSubComponent},
    ],
    canActivate: [AuthGuard]
  }
];