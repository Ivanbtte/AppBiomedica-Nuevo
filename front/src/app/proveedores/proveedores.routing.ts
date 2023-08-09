import {Routes} from '@angular/router';
import {AuthGuard} from '../shared/modulos/auth.guard';
import {ProveedorComponent} from './proveedores/proveedor.component';

export const ProveedoresRoutes: Routes = [
  {
    path: '',
    children: [
      {
        path: 'lista',
        component: ProveedorComponent
      }
    ],
    canActivate: [AuthGuard]
  }
];
