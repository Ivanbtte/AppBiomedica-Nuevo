import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {SolicitudTrasladoComponent} from './components/solicitud-traslado/solicitud-traslado.component';
import {MaterialModule} from '../material.module';
import {RouterModule} from '@angular/router';
import {SharedModule} from '../shared/shared.module';
import {TrasladosRoutes} from './traslado-pacientes.routing';
import {FlexLayoutModule} from '@angular/flex-layout';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {DialogAutorizaAcompComponent} from './components/dialogs/dialog-autoriza-acomp/dialog-autoriza-acomp.component';
import {VistaPreviaSolicitudComponent} from './components/vista-previa-solicitud/vista-previa-solicitud.component';
import {DialogExcepcionesComponent} from './components/dialogs/dialog-excepciones/dialog-excepciones.component';
import {SolicitudesRealizadasComponent} from './components/solicitudes-realizadas/solicitudes-realizadas.component';
import {SolicitudesCanceladasComponent} from './components/solicitudes-canceladas/solicitudes-canceladas.component';
import {FirmasComponent} from './components/firmas/firmas.component';
import {DialogNewFirmComponent} from './components/dialogs/dialog-new-firm/dialog-new-firm.component';
import {DialogEditFirmComponent} from './components/dialogs/dialog-edit-firm/dialog-edit-firm.component';
import {PerfectScrollbarModule} from 'ngx-perfect-scrollbar';
import {CustomFormsModule} from 'ng2-validation';
import {TransferManagerComponent} from './components/transfer-manager/transfer-manager.component';
import {AddPermitModuleComponent} from './components/dialogs/add-permit-module/add-permit-module.component';
import { ViewPermitByUserComponent } from './components/view-permit-by-user/view-permit-by-user.component';
import { DialogDeleteUserComponent } from './components/dialogs/dialog-delete-user/dialog-delete-user.component';
import { AdministrarUsuariosComponent } from './components/administrar-usuarios/administrar-usuarios.component';
import { CargarFtpComponent } from './components/cargar-ftp/cargar-ftp.component';
import { TableroTrasladosComponent } from './components/tablero-traslados/tablero-traslados.component';
import { DetallePorContratoComponent } from './components/detalle-por-contrato/detalle-por-contrato.component';

@NgModule({
  declarations: [SolicitudTrasladoComponent, DialogAutorizaAcompComponent, VistaPreviaSolicitudComponent,
    DialogExcepcionesComponent, SolicitudesRealizadasComponent, SolicitudesCanceladasComponent, FirmasComponent, DialogNewFirmComponent,
    DialogEditFirmComponent, TransferManagerComponent, AddPermitModuleComponent, ViewPermitByUserComponent, DialogDeleteUserComponent, AdministrarUsuariosComponent, CargarFtpComponent, TableroTrasladosComponent, DetallePorContratoComponent,],
  entryComponents: [DialogAutorizaAcompComponent, VistaPreviaSolicitudComponent, DialogExcepcionesComponent, DialogNewFirmComponent, DialogEditFirmComponent, AddPermitModuleComponent, DialogDeleteUserComponent],
  imports: [
    CommonModule,
    RouterModule.forChild(TrasladosRoutes),
    FlexLayoutModule,
    FormsModule,
    ReactiveFormsModule,
    SharedModule,
    MaterialModule,
    PerfectScrollbarModule,
    CustomFormsModule,
  ],
  exports: [RouterModule]
})
export class TrasladoPacientesModule {
}
