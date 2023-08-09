import { Routes } from '@angular/router';
import { AuthGuard } from '../shared/modulos/auth.guard';
import { SolicitudSubrogadoComponent } from './solicitud-subrogado/solicitud-subrogado.component';

export const ServiciosSubrogadosRoutes: Routes = [
  {
    path: '',
    children: [
      {path: 'solicitud', component: SolicitudSubrogadoComponent},
    ],
    canActivate: [AuthGuard]
  }
];