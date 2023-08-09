import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {ValidarComponent} from './validar/validar.component';
import {VerComponent} from './ver/ver.component';
import {MaterialModule} from '../material.module';
import {RouterModule} from '@angular/router';
import {SharedModule} from '../shared/shared.module';
import {ValidacionRoutes} from './validacion-solicitudes.routing';
import {DetalleValidarComponent} from './detalle-validar/detalle-validar.component';
import {FiltrosSolicitudComponent} from './filtros-solicitud/filtros-solicitud.component';
import {FlexLayoutModule} from '@angular/flex-layout';
import {ConsultasAnaliticasComponent} from './consultas-analiticas/consultas-analiticas.component';
import {CambiarEstatusComponent} from './cambiar-estatus/cambiar-estatus.component';
import { SolicitudFinalComponent } from './solicitud-final/solicitud-final.component';

@NgModule({
  entryComponents: [FiltrosSolicitudComponent, ConsultasAnaliticasComponent, CambiarEstatusComponent],
  declarations: [ValidarComponent, VerComponent, DetalleValidarComponent, FiltrosSolicitudComponent, ConsultasAnaliticasComponent,
    CambiarEstatusComponent,
    SolicitudFinalComponent],
  imports: [
    CommonModule,
    MaterialModule,
    RouterModule.forChild(ValidacionRoutes),
    SharedModule,
    FlexLayoutModule
  ]
})
export class ValidacionSolicitudesModule {
}
