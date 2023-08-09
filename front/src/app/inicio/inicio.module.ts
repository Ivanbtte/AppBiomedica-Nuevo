import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {InicioComponent} from './inicio/inicio.component';
import {MaterialModule} from '../material.module';
import {InicioRotues} from './inicio.routing';
import {RouterModule} from '@angular/router';
import {FlexLayoutModule} from '@angular/flex-layout';
// import {ChartsModule} from 'ng2-charts';

@NgModule({
  declarations: [InicioComponent],
  imports: [
    CommonModule,
    MaterialModule,
    FlexLayoutModule,
    // ChartsModule,
    RouterModule.forChild(InicioRotues)
  ],
  exports: [RouterModule],
})
export class InicioModule {
}
