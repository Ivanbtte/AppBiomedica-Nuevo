import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {MaterialModule} from '../material.module';
import {RouterModule} from '@angular/router';
import {SharedModule} from '../shared/shared.module';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {FlexLayoutModule} from '@angular/flex-layout';
import {AgregarAreaServicioComponent } from './agregar-area-servicio/agregar-area-servicio.component';
import {AreaServicioRoutes} from './area-servicios.routing';

@NgModule({
  declarations: [AgregarAreaServicioComponent],
  imports: [
    CommonModule,
    MaterialModule,
    RouterModule.forChild(AreaServicioRoutes),
    SharedModule,
    FormsModule,
    ReactiveFormsModule,
    FlexLayoutModule,
  ]
})
export class AreaServiciosModule { }
