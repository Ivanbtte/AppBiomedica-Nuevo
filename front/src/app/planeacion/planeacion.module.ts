import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { EnfermeriaFondoFijoComponent } from './con-cuadro-basico/enfermeria-fondo-fijo.component';
import {MaterialModule} from '../material.module';
import {RouterModule} from '@angular/router';
import {SharedModule} from '../shared/shared.module';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {PlaneacionRoutes} from './planeacion.routing';
import { SinCuadroBasicoComponent } from './sin-cuadro-basico/sin-cuadro-basico.component';
import {FlexLayoutModule} from '@angular/flex-layout';

@NgModule({
  declarations: [EnfermeriaFondoFijoComponent, SinCuadroBasicoComponent],
  imports: [
    CommonModule,
    MaterialModule,
    RouterModule.forChild(PlaneacionRoutes),
    SharedModule,
    FormsModule,
    ReactiveFormsModule,
    FlexLayoutModule
  ]
})
export class PlaneacionModule { }
