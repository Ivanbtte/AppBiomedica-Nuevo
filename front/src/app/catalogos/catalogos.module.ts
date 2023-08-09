import {NgModule} from '@angular/core';
import {RouterModule} from '@angular/router';
import {CommonModule} from '@angular/common';
import {SaiComponent} from './sai/sai.component';
import {PreiComponent} from './prei/prei.component';
import {CatalogosRoutes} from './catalogos.routing';
import {MaterialModule} from '../material.module';
import {PreiService} from './servicios/prei.service';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {DetalleDialogComponent} from './detalle-dialog/detalle-dialog.component';
import {BusquedaSelectorComponent} from './busqueda-selector/busqueda-selector.component';
import {DetalleSaiComponent} from './detalle-sai/detalle-sai.component';
import {ProformaComponent} from './proforma/proforma.component';
import {DetalleProformaComponent} from './detalle-proforma/detalle-proforma.component';
import {FlexLayoutModule} from '@angular/flex-layout';
import {MatPaginator} from '@angular/material/paginator'
import {PerfectScrollbarModule} from 'ngx-perfect-scrollbar';
import { SharedModule } from '../shared/shared.module';

@NgModule({
  entryComponents: [DetalleDialogComponent, BusquedaSelectorComponent, DetalleSaiComponent, DetalleProformaComponent],
  declarations: [SaiComponent, PreiComponent, DetalleDialogComponent, BusquedaSelectorComponent, DetalleSaiComponent, ProformaComponent,
    DetalleProformaComponent],
  imports: [
    CommonModule,
    MaterialModule,
    RouterModule.forChild(CatalogosRoutes),
    SharedModule,
    FormsModule,
    ReactiveFormsModule,
    FlexLayoutModule,
    PerfectScrollbarModule,
  ],
  exports:[
    MatPaginator
  ],
  providers: [
    PreiService
  ]
})
export class CatalogosModule {
}
