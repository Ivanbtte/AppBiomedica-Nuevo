import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {RouterModule} from '@angular/router';
import {MaterialModule} from '../material.module';
import {SolicitudesRoutes} from './solicitudes.routing';
import {EqCocinaComponent} from './eq-cocina/eq-cocina.component';
import {EqMedicoComponent} from './eq-medico/eq-medico.component';
import {InstrumentalComponent} from './instrumental/instrumental.component';
import {EditarSolicComponent} from './editar-solic/editar-solic.component';
import {SeguimientoSolicComponent} from './seguimiento-solic/seguimiento-solic.component';
import {SharedModule} from '../shared/shared.module';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import { ConsumoComponent } from './consumo/consumo.component';
import { DetalleSeguimientoComponent } from './detalle-seguimiento/detalle-seguimiento.component';
import { DetalleEditarComponent } from './detalle-editar/detalle-editar.component';


@NgModule({
  entryComponents: [DetalleSeguimientoComponent],
  declarations: [
    EqCocinaComponent,
    EqMedicoComponent,
    InstrumentalComponent,
    EditarSolicComponent,
    SeguimientoSolicComponent,
    ConsumoComponent,
    DetalleSeguimientoComponent,
    DetalleEditarComponent,
  ],
  imports: [
    CommonModule,
    MaterialModule,
    RouterModule.forChild(SolicitudesRoutes),
    SharedModule,
    FormsModule,
    ReactiveFormsModule
  ]
})
export class SolicitudesModule {
}
