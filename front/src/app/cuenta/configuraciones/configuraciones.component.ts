import {Component, OnInit} from '@angular/core';
import {UsuarioLogeado} from '../../core/models/UsuarioLogin.interface';
import jwt_decode from 'jwt-decode';
import {NombramientoService} from '../../core/servicios/nombramiento.service';
import {UnidadNombramientoService} from '../../core/servicios/unidad-nombramiento.service';
import {UnidadesMedicasService} from '../../core/servicios/unidades-medicas.service';
import {Router} from '@angular/router';

@Component({
  selector: 'app-configuraciones',
  templateUrl: './configuraciones.component.html',
  styleUrls: ['./configuraciones.component.css']
})
export class ConfiguracionesComponent implements OnInit {
  public usuarioLog!: UsuarioLogeado;
  public nombramiento: any;
  public unidades: any[] = [];
  public unidadActual = '';
  public selected = 0;
  public visible!: boolean;
  public visibleSelec = true;
  public unidadMed: any;

  constructor(
    private _nombramiento: NombramientoService,
    private _unidadNombrami: UnidadNombramientoService,
    private _unidades: UnidadesMedicasService,
    private router: Router,
  ) {
  }

  ngOnInit() {
    this.usuarioLog = jwt_decode(<string> localStorage.getItem('token'));
    this.unidadActual = <string>localStorage.getItem('unidad');
    console.log(this.usuarioLog);
    this._nombramiento.nombramiento(this.usuarioLog.PersonaId).subscribe(res => {
      this.nombramiento = res.Data;
      this.unidades = this.nombramiento.UnidadMed;
      console.log(this.unidades);
      if (this.unidades.length > 1) {
        this.visible = false;
      } else {
        this.visible = true;
      }
      this.selected = this.unidades[0].Id;
    }, err => {
      console.log(err);
    });
  }

  selectUnidades(value: string) {
    console.log('cambie:_ ', value);
    this._unidades.unaUnidad(value).subscribe(res => {
      this.unidadMed = res.Data;
      localStorage.setItem('unidad', this.unidadMed.DenominacionUni);
      console.log(localStorage.getItem('unidad'));
      this.router.navigate(['/cuenta/configuracion']);
    }, err => {
      console.log(err);
    });
  }

  desactivarSelect() {
    this.visibleSelec = !this.visibleSelec;
  }

}
