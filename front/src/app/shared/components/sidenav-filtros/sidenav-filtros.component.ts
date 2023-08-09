import { Component, OnInit } from '@angular/core';
import {UsuarioLogeado} from '../../../core/models/UsuarioLogin.interface';
import {MenuItems} from '../../../shared/menu-items/menu-items';
import jwt_decode from 'jwt-decode';
@Component({
  selector: 'app-sidenav-filtros',
  templateUrl: './sidenav-filtros.component.html',
  styleUrls: ['./sidenav-filtros.component.css']
})
export class SidenavFiltrosComponent implements OnInit {
  public usuarioLog: UsuarioLogeado;

  constructor(
    public menu: MenuItems,)
  {
    this.usuarioLog = jwt_decode(<string> localStorage.getItem('token'));
  }

  ngOnInit(): void {
  }

}
