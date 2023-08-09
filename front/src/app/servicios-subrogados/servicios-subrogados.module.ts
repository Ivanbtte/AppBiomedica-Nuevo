import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {FlexLayoutModule} from '@angular/flex-layout';
import { ServiciosSubrogadosRoutes } from './servicios-subrogados.routing';
import { RouterModule } from '@angular/router';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {SharedModule} from '../shared/shared.module';
import {MaterialModule} from '../material.module';
import {PerfectScrollbarModule} from 'ngx-perfect-scrollbar';
import {CustomFormsModule} from 'ng2-validation';
import { SolicitudSubrogadoComponent } from './solicitud-subrogado/solicitud-subrogado.component';


@NgModule({
  declarations: [
    SolicitudSubrogadoComponent
  ],
  imports: [
    CommonModule,
    RouterModule.forChild(ServiciosSubrogadosRoutes),
    FlexLayoutModule,
    FormsModule,
    ReactiveFormsModule,
    SharedModule,
    MaterialModule,
    PerfectScrollbarModule,
    CustomFormsModule,
  ],
  exports:[RouterModule]
})
export class ServiciosSubrogadosModule { }
