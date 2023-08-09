import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {InfoCuentaComponent} from './info-cuenta/info-cuenta.component';
import {MaterialModule} from '../material.module';
import {CuentaRoutes} from './cuenta.routing';
import {RouterModule} from '@angular/router';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {ConfiguracionesComponent } from './configuraciones/configuraciones.component';
import {FlexLayoutModule} from '@angular/flex-layout';
import {CoreModule} from '../core/core.module';

@NgModule({
  declarations: [InfoCuentaComponent, ConfiguracionesComponent],
    imports: [
        CommonModule,
        MaterialModule,
        RouterModule.forChild(CuentaRoutes),
        FormsModule,
        ReactiveFormsModule,
        FlexLayoutModule,
        CoreModule,
    ]
})
export class CuentaModule {
}
