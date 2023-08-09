import {Component, OnInit} from '@angular/core';
import {Router} from '@angular/router';
import Swal from 'sweetalert2';
import jwt_decode from 'jwt-decode';
import {FormBuilder, FormGroup, Validators,} from '@angular/forms';
import {Credenciales} from '../interfaces/usuario.interface';
import {LoginService} from '../servicios/login.service';
import {RespuestaPeticion} from '../../core/models/estructuras_generales';
import {UnidadNombramientoService} from '../../core/servicios/unidad-nombramiento.service';
import {UsuarioLogeado} from '../../core/models/UsuarioLogin.interface';
import {UsuarioService} from '../servicios/usuario.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})

export class LoginComponent implements OnInit {
  hide = true;
  public form!: FormGroup;
  public usuarioLog: UsuarioLogeado;

  constructor(
    private fb: FormBuilder,
    private router: Router,
    private _login: LoginService,
    private _unidadNombrami: UnidadNombramientoService,
    private usuarioService: UsuarioService
  ) {
  }

  ngOnInit() {
    this.form = this.fb.group({
      Correo: [
        null,
        Validators.compose([Validators.required, Validators.email])
      ],
      password: [null, Validators.compose([Validators.required])]
    });
  }

  onSubmit() {
    const datos = <Credenciales> this.form.getRawValue();
    this._login.login(datos).subscribe(res => {
        if (res.Data['status'] === true) {
          Swal.fire({
            icon: 'warning',
            title: 'No tienes registrado un puesto',
            text: '¿Deseas agregarlo?',
            showConfirmButton: true,
            confirmButtonText: 'Si, Agregar',
            confirmButtonColor: '#3085d6',
            showCancelButton: true,
            cancelButtonText: 'Después',
            reverseButtons: true
          }).then((result) => {
            if (result.isConfirmed) {
              this.router.navigate(['/cuenta/info'], {queryParams: {index: 1}});
            }
          });
        } else {
          Swal.fire({
            icon: 'success',
            title: res.Mensaje,
            showConfirmButton: false,
            timer: 2000,
          });
        }
        localStorage.setItem('token', res.Data['token']);
        const temp = localStorage.getItem('token');
        this.usuarioLog = jwt_decode(temp);
        const nameModule = res.Data['permits'].map(item => item.modulo.name);
        localStorage.setItem('permisos', JSON.stringify(nameModule));
        localStorage.setItem('idUser', this.usuarioLog.UsuarioId);
        localStorage.setItem('menu', JSON.stringify(this.usuarioLog.MenuItem));
        //this.consultar_permisos();
        this.consultUserPosition(this.usuarioLog.UsuarioId);
        this.router.navigate(['/inicio']);
      }, err => {
        if (err.error !== '' || err.error !== undefined) {
          const errores = <RespuestaPeticion> err.error;
          Swal.fire({
            icon: 'error',
            title: errores.Data,
            text: errores.Mensaje,
          });
        }
      }
    );
  }

  consultUserPosition(id: string) {
    this.usuarioService.ConsultUSerByPosition(id).subscribe(result => {
      localStorage.setItem('user', JSON.stringify(result.Data));
      localStorage.setItem('position', result.Data.PuestoOrganigrama.catalogoPuesto.Puesto);
    });
  }
}
