import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {RouterModule} from '@angular/router';
import {MaterialModule} from '../material.module';
import {SharedModule} from '../shared/shared.module';
import {FechaLimiteComponent} from './fecha-limite/fecha-limite.component';
import {PresupuestoComponent} from './presupuesto/presupuesto.component';
import {ConfiguracionesRoutes} from './configuraciones.routing';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {FlexLayoutModule} from '@angular/flex-layout';


@NgModule({
  declarations: [
    FechaLimiteComponent,
    PresupuestoComponent,
  ],
  imports: [
    CommonModule,
    MaterialModule,
    RouterModule.forChild(ConfiguracionesRoutes),
    SharedModule,
    FormsModule,
    ReactiveFormsModule,
    FlexLayoutModule
  ]
})
export class ConfiguracionesModule {
}
