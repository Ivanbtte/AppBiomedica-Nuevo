import {Routes} from '@angular/router';
import {SaiComponent} from './sai/sai.component';
import {PreiComponent} from './prei/prei.component';
import {AuthGuard} from '../shared/modulos/auth.guard';
import {ProformaComponent} from './proforma/proforma.component';

export const CatalogosRoutes: Routes = [
  {
    path: '',
    children: [
      {
        path: 'sai',
        component: SaiComponent,
        canActivate: [AuthGuard]
      },
      {
        path: 'prei',
        component: PreiComponent,
        canActivate: [AuthGuard]
      }, {
        path: 'proforma',
        component: ProformaComponent,
        canActivate: [AuthGuard]
      }
    ],
    canActivate: [AuthGuard]
  }
];
