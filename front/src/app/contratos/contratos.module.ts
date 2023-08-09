import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {RouterModule} from '@angular/router';
import {MaterialModule} from '../material.module';
import {SharedModule} from '../shared/shared.module';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {FlexLayoutModule} from '@angular/flex-layout';
import {AdministrarContratosComponent} from './administrar-contratos/administrar-contratos.component';
import {SeguimientoContratosComponent} from './seguimiento-contratos/seguimiento-contratos.component';
import {ContratosRoutes} from './contratos.routing';
import {DetalleContratoComponent} from './detalle-contrato/detalle-contrato.component';
// import {Ng2SmartTableModule} from 'ng2-smart-table';
import {AgregarConceptosContratoComponent} from './agregar-conceptos-contrato/agregar-conceptos-contrato.component';
import {FianzaComponent} from './fianza/fianza.component';
import {TrasladoPacientesComponent} from './ConceptosServicios/traslado-pacientes/traslado-pacientes.component';
import {EquipoMedicoComponent} from './ConceptosServicios/equipo-medico/equipo-medico.component';
import { EstudiosLaboratorioComponent } from './ConceptosServicios/estudios-laboratorio/estudios-laboratorio.component';

@NgModule({
  entryComponents: [EquipoMedicoComponent, TrasladoPacientesComponent],
  declarations: [AdministrarContratosComponent, SeguimientoContratosComponent, DetalleContratoComponent,
    AgregarConceptosContratoComponent,
    FianzaComponent,
    TrasladoPacientesComponent,
    EquipoMedicoComponent,
    EstudiosLaboratorioComponent],
  imports: [
    CommonModule,
    MaterialModule,
    RouterModule.forChild(ContratosRoutes),
    SharedModule,
    FormsModule,
    ReactiveFormsModule,
    FlexLayoutModule,
    // Ng2SmartTableModule,
  ]
})
export class ContratosModule {
}
