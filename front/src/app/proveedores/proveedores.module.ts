import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ProveedorComponent } from './proveedores/proveedor.component';
import {MaterialModule} from '../material.module';
import {RouterModule} from '@angular/router';
import {SharedModule} from '../shared/shared.module';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {FlexLayoutModule} from '@angular/flex-layout';
import {ProveedoresRoutes} from './proveedores.routing';

@NgModule({
  declarations: [ProveedorComponent],
  imports: [
    CommonModule,
    MaterialModule,
    RouterModule.forChild(ProveedoresRoutes),
    SharedModule,
    FormsModule,
    ReactiveFormsModule,
    FlexLayoutModule
  ]
})
export class ProveedoresModule { }
