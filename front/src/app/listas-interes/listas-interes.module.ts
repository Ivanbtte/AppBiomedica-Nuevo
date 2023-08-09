import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {RouterModule} from '@angular/router';
import {MaterialModule} from '../material.module';
import {SharedModule} from '../shared/shared.module';
import {ListasComponent} from './listas/listas.component';
import {ListasRoutes} from './listas-interes.routing';
import {FlexLayoutModule} from '@angular/flex-layout';

@NgModule({
  declarations: [
    ListasComponent
  ],
  imports: [
    CommonModule,
    MaterialModule,
    RouterModule.forChild(ListasRoutes),
    SharedModule,
    FlexLayoutModule
  ]
})
export class ListasInteresModule {
}
