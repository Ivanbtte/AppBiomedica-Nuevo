import {ChangeDetectorRef, Component, EventEmitter, OnDestroy, OnInit, Output} from '@angular/core';
import {PerfectScrollbarConfigInterface} from 'ngx-perfect-scrollbar';
import {MediaMatcher} from '@angular/cdk/layout';
import jwt_decode from 'jwt-decode';
import {UsuarioLogeado} from '../../../core/models/UsuarioLogin.interface';
import {LoginService} from '../../../authentication/servicios/login.service';
import {Router} from '@angular/router';
import {MenuItems} from '../../../shared/menu-items/menu-items';
import {PersonaService} from '../../../authentication/servicios/persona.service';
import {Personas} from '../../../authentication/interfaces/usuario.interface';

export let usuarioLog: UsuarioLogeado;

@Component({
  selector: 'app-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: []
})
export class AppSidebarComponent implements OnDestroy, OnInit {
  public config: PerfectScrollbarConfigInterface = {};
  mobileQuery: MediaQueryList;
  private _mobileQueryListener: () => void;
  status = false;
  public usuarioLog: UsuarioLogeado;
  @Output() closeNav: EventEmitter<boolean> = new EventEmitter;
  public persona: Personas;

  clickEvent() {
    this.status = !this.status;
  }

  masculino = false;
  femenino = false;
  otro = false;

  constructor(
    changeDetectorRef: ChangeDetectorRef,
    media: MediaMatcher,
    public menu: MenuItems,
    private _login: LoginService,
    private router: Router,
    private _persona: PersonaService,

  ) {
    this.mobileQuery = media.matchMedia('(min-width: 768px)');
    this._mobileQueryListener = () => changeDetectorRef.detectChanges();
    this.mobileQuery.addListener(this._mobileQueryListener);
    this.usuarioLog = jwt_decode(<string> localStorage.getItem('token'));
  }

  ngOnInit() {
    this._persona.getPersona(this.usuarioLog.PersonaId).subscribe(result => {
      this.persona = result.Data;
      if (this.persona.Genero === 'M') {
        this.masculino = true;
      }
      if (this.persona.Genero === 'O') {
        this.otro = true;
      }
      if (this.persona.Genero === 'F') {
        this.femenino = true;
      }
    }, err => {
      console.log(err);
    });

  }

  closeSideNv() {
    this.closeNav.emit(true);
  }

  ngOnDestroy(): void {
    this.mobileQuery.removeListener(this._mobileQueryListener);
  }

  cerrarSesion() {
    this._login.logoutUser();
  }

  cuenta() {
    this.router.navigate(['cuenta/info']);
  }

  configuracion() {
    this.router.navigate(['cuenta/configuracion']);
  }
}
