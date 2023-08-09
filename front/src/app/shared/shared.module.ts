import {NgModule} from '@angular/core';
import {AuthGuard} from './modulos/auth.guard';
import {MaterialModule} from '../material.module';
import {MenuItems} from './menu-items/menu-items';
import {AccordionAnchorDirective, AccordionDirective, AccordionLinkDirective} from './accordion';
import {CommonModule} from '@angular/common';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {FlexLayoutModule} from '@angular/flex-layout';
import {ReturnTypeFirmPipe} from './pipes/return-type-firm.pipe';
import {ReturnStateFirmPipe} from './pipes/return-state-firm.pipe';
import {ReturnConditionFirmPipe} from './pipes/return-condition-firm.pipe';
import {CantidadDialogComponent} from './Dialogs/cantidad-dialog/cantidad-dialog.component';
import {AgregarProveedorComponent} from './Dialogs/agregar-proveedor/agregar-proveedor.component';
import {ContactosProveedoresComponent} from './Dialogs/contactos-proveedores/contactos-proveedores.component';
import {DetalleProveedorComponent} from './Dialogs/detalle-proveedor/detalle-proveedor.component';
import {DetalleContactoProveedorComponent} from './Dialogs/detalle-contacto-proveedor/detalle-contacto-proveedor.component';
import {AgregarContactoProveedorComponent} from './Dialogs/agregar-contacto-proveedor/agregar-contacto-proveedor.component';
import {CrudPersonasComponent} from './Dialogs/crud-personas/crud-personas.component';
import {SeleccionarProveedorRepresentanteComponent} from './Dialogs/Contratos/seleccionar-proveedor-representante/seleccionar-proveedor-representante.component';
import {ConsultarClavePreiComponent} from './Dialogs/consultar-clave-prei/consultar-clave-prei.component';
import {DistribucionConceptosComponent} from './Dialogs/Contratos/distribucion-conceptos/distribucion-conceptos.component';
import {EditarConceptoContratoComponent} from './Dialogs/Contratos/editar-concepto-contrato/editar-concepto-contrato.component';
import {CargaMasivaContratosComponent} from './Dialogs/Contratos/carga-masiva-contratos/carga-masiva-contratos.component';
import {ConsultarContratoComponent} from './Dialogs/Contratos/consultar-contrato/consultar-contrato.component';
import {EnviarArchivoConceptComponent} from './Dialogs/Contratos/carga-masiva-conceptos/carga-masiva-conceptos.component';
import {CargaMasivaDistribucionComponent} from './Dialogs/Contratos/carga-masiva-distribucion/carga-masiva-distribucion.component';
import {VerificarDistribucionComponent} from './Dialogs/Contratos/verificar-distribucion/verificar-distribucion.component';
import {ExaminarMontoContratoComponent} from './Dialogs/Contratos/examinar-monto-contrato/examinar-monto-contrato.component';
import {DetalleContacContrComponent} from './Dialogs/Contratos/detalle-contac-contr/detalle-contac-contr.component';
import {EditarPrecioPreiComponent} from './Dialogs/Contratos/editar-precio-prei/editar-precio-prei.component';
import {FiltrosContratosComponent} from './Dialogs/Contratos/filtros-contratos/filtros-contratos.component';
import {AgregarEditarDistrComponent} from './Dialogs/Contratos/agregar-editar-distr/agregar-editar-distr.component';
import {RepresentanteLegalComponent} from './Dialogs/Contratos/representante-legal/representante-legal.component';
import {AgregarUnContratoComponent} from './Dialogs/Contratos/agregar-un-contrato/agregar-un-contrato.component';
import {ArchivosContratoComponent} from './Dialogs/Contratos/archivos-contrato/archivos-contrato.component';
import {VisualizadorPDFComponent} from './Dialogs/Contratos/visualizador-pdf/visualizador-pdf.component';
import {SaiSolicitudComponent} from './Dialogs/sai-solicitud/sai-solicitud.component';
import {QRCodeModule} from 'angular2-qrcode';
import {AgregarEditarContratoComponent} from './Dialogs/Contratos/agregar-editar-contrato/agregar-editar-contrato.component';
import {NgxDropzoneModule} from 'ngx-dropzone';
import {SidenavFiltrosComponent} from './components/sidenav-filtros/sidenav-filtros.component';
import {CapitalizadoPipe} from './pipes/capitalizado.pipe';
import {ThousandsPipe} from './pipes/thousands-pipe.pipe';
import { NgChartsModule } from 'ng2-charts';

// In your App's module:
imports: [
  NgChartsModule
]
@NgModule({
  entryComponents: [
    CantidadDialogComponent,
    AgregarProveedorComponent,
    ContactosProveedoresComponent,
    DetalleProveedorComponent,
    DetalleContactoProveedorComponent,
    AgregarContactoProveedorComponent,
    AgregarEditarContratoComponent,
    CrudPersonasComponent,
    SeleccionarProveedorRepresentanteComponent,
    ConsultarClavePreiComponent,
    DistribucionConceptosComponent,
    EditarConceptoContratoComponent,
    CargaMasivaContratosComponent,
    ConsultarContratoComponent,
    EnviarArchivoConceptComponent,
    CargaMasivaDistribucionComponent,
    VerificarDistribucionComponent,
    ExaminarMontoContratoComponent,
    DetalleContacContrComponent,
    EditarPrecioPreiComponent,
    FiltrosContratosComponent,
    AgregarEditarDistrComponent,
    RepresentanteLegalComponent,
    AgregarUnContratoComponent,
    ArchivosContratoComponent,
    VisualizadorPDFComponent
  ],
  declarations: [
    AccordionAnchorDirective,
    AccordionLinkDirective,
    AccordionDirective,
    ReturnTypeFirmPipe,
    ReturnStateFirmPipe,
    ReturnConditionFirmPipe,
    CapitalizadoPipe,
    ThousandsPipe,
    SaiSolicitudComponent,
    CantidadDialogComponent,
    AgregarProveedorComponent,
    ContactosProveedoresComponent,
    DetalleProveedorComponent,
    DetalleContactoProveedorComponent,
    AgregarContactoProveedorComponent,
    AgregarEditarContratoComponent,
    CrudPersonasComponent,
    SeleccionarProveedorRepresentanteComponent,
    ConsultarClavePreiComponent,
    DistribucionConceptosComponent,
    EditarConceptoContratoComponent,
    CargaMasivaContratosComponent,
    ConsultarContratoComponent,
    EnviarArchivoConceptComponent,
    CargaMasivaDistribucionComponent,
    VerificarDistribucionComponent,
    ExaminarMontoContratoComponent,
    DetalleContacContrComponent,
    EditarPrecioPreiComponent,
    FiltrosContratosComponent,
    AgregarEditarDistrComponent,
    RepresentanteLegalComponent,
    AgregarUnContratoComponent,
    ArchivosContratoComponent,
    VisualizadorPDFComponent,
    SidenavFiltrosComponent
  ],
  exports: [
    AccordionAnchorDirective,
    AccordionLinkDirective,
    AccordionDirective,
    ReturnTypeFirmPipe,
    ReturnStateFirmPipe,
    CapitalizadoPipe,
    ThousandsPipe,
    SidenavFiltrosComponent,
    NgxDropzoneModule,
    NgChartsModule
  ],
  imports: [
    MaterialModule,
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    FlexLayoutModule,
    QRCodeModule,
    NgxDropzoneModule,
    NgChartsModule
  ],
  providers: [MenuItems, AuthGuard]
})
export class SharedModule {
}
