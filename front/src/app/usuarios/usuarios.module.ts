import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {RouterModule} from '@angular/router';
import {MaterialModule} from '../material.module';
import {SharedModule} from '../shared/shared.module';
import {ValidarUsuariosComponent} from './validar-usuarios/validar-usuarios.component';
import {BloquearUsuariosComponent} from './bloquear-usuarios/bloquear-usuarios.component';
import {UsuariosRoutes} from './usuarios.routing';
import {DetalleUsuarioComponent} from './detalle-usuario/detalle-usuario.component';
import {FlexModule} from '@angular/flex-layout';

@NgModule({
  entryComponents: [
    DetalleUsuarioComponent
  ],
  declarations: [
    ValidarUsuariosComponent,
    BloquearUsuariosComponent,
    DetalleUsuarioComponent
  ],
    imports: [
        CommonModule,
        MaterialModule,
        RouterModule.forChild(UsuariosRoutes),
        SharedModule,
        FlexModule,
    ]
})
export class UsuariosModule {
}
