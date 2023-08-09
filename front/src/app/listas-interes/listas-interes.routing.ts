import {Routes} from '@angular/router';
import {AuthGuard} from '../shared/modulos/auth.guard';
import {ListasComponent} from './listas/listas.component';

export const ListasRoutes: Routes = [
  {
    path: '',
    component: ListasComponent,
    canActivate: [AuthGuard]
  }
];
