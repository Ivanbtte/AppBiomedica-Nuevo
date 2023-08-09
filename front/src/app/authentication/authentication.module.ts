import {NgModule} from '@angular/core';
import {RouterModule} from '@angular/router';
import {CommonModule} from '@angular/common';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {FlexLayoutModule} from '@angular/flex-layout';
import {UsuarioService} from './servicios/usuario.service';
import {SharedModule} from '../shared/shared.module';
import {AuthenticationRoutes} from './authentication.routing';
import {ErrorComponent} from './error/error.component';
import {ForgotComponent} from './forgot/forgot.component';
import {LoginComponent} from './login/login.component';
import {RegisterComponent} from './register/register.component';
import {MaterialModule} from '../material.module';
import {PerfectScrollbarModule} from 'ngx-perfect-scrollbar';
import {DepartmentTreeComponent} from './department-tree/department-tree.component';


@NgModule({
  entryComponents: [DepartmentTreeComponent],
  imports: [
    CommonModule,
    RouterModule.forChild(AuthenticationRoutes),
    FlexLayoutModule,
    FormsModule,
    ReactiveFormsModule,
    SharedModule,
    MaterialModule,
    PerfectScrollbarModule
  ],
  declarations: [
    ErrorComponent,
    ForgotComponent,
    LoginComponent,
    RegisterComponent,
    DepartmentTreeComponent
  ],
  providers: [
    UsuarioService
  ]
})
export class AuthenticationModule {
}
