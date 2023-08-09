import { Component, OnInit } from '@angular/core';
import { UsuarioService } from '../../authentication/servicios/usuario.service';
import { FechaLimiteService } from '../../configuraciones/serviciosConfiguraciones/fecha-limite.service';
import { UsuarioLogeado } from '../../core/models/UsuarioLogin.interface';
import jwt_decode from 'jwt-decode';

@Component({
  selector: 'app-inicio',
  templateUrl: './inicio.component.html',
  styleUrls: ['./inicio.component.scss']
})
export class InicioComponent implements OnInit {
  userLog: UsuarioLogeado;

  constructor(
    private _usuario: UsuarioService,
  ) {
    this.userLog = jwt_decode(<string>localStorage.getItem('token'));
  }
  ngOnInit() {
    if (this.userLog.Tipo === '3' || this.userLog.Tipo === '4') {
      this.getUnit();
    }
  }
  getUnit() {
    this._usuario.unidadUsuario(<string>localStorage.getItem('idUser')).subscribe(res => {
      localStorage.setItem('unit', JSON.stringify(res.Data));
    });
  }

}
