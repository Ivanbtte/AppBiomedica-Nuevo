import {MediaMatcher} from '@angular/cdk/layout';
import {AfterViewInit, ChangeDetectorRef, Component, OnDestroy, OnInit, ViewChild} from '@angular/core';
import {PerfectScrollbarConfigInterface} from 'ngx-perfect-scrollbar';
import {UsuarioLogeado} from '../../core/models/UsuarioLogin.interface';
import {NombramientoService} from '../../core/servicios/nombramiento.service';
import {AbrirCatalogoSaiService} from '../../core/servicios/abrir-catalogo-sai.service';
import {Observable, Subscription} from 'rxjs';
import {Router} from '@angular/router';
import jwt_decode from 'jwt-decode';
import {MatSidenav} from '@angular/material/sidenav';
import {UsuarioService} from '../../authentication/servicios/usuario.service';

/** @title Responsive sidenav */
@Component({
  selector: 'app-full-layout',
  templateUrl: 'full.component.html',
  styleUrls: []
})
export class FullComponent implements OnInit, OnDestroy, AfterViewInit {
  mobileQuery: MediaQueryList;
  public usuarioLog: UsuarioLogeado;
  public unidadActual = '';
  dir = 'ltr';
  green = true;
  blue = false;
  dark = false;
  minisidebar!: boolean;
  boxed!: boolean;
  danger = false;
  showHide!: boolean;
  estadoBtn$!: Observable<boolean>;
  ruta_btn!: string;
  public config: PerfectScrollbarConfigInterface = {};
  private readonly _mobileQueryListener: () => void;
  subscription2!: Subscription;
  sidebarOpened!: boolean;
  @ViewChild('snav', {static: false}) public snav!: MatSidenav;

  constructor(
    changeDetectorRef: ChangeDetectorRef,
    media: MediaMatcher,
    private activa_btn_cancelar: AbrirCatalogoSaiService,
    private router: Router,
    private usuarioService: UsuarioService
  ) {
    this.mobileQuery = media.matchMedia('(min-width: 768px)');
    this._mobileQueryListener = () => changeDetectorRef.detectChanges();
    this.mobileQuery.addListener(this._mobileQueryListener);
    this.usuarioLog = jwt_decode(<string> localStorage.getItem('token'));
  }

  ngOnInit() {
    //this.consultUserPosition(this.usuarioLog.UsuarioId);
    this.unidadActual = <string> localStorage.getItem('unidad');
    this.estadoBtn$ = this.activa_btn_cancelar.getEstadoBtn();
    this.subscription2 = this.activa_btn_cancelar.getRuta().subscribe(resultado => {
      this.ruta_btn = resultado;
    });
  }

  ngOnDestroy(): void {
    this.mobileQuery.removeListener(this._mobileQueryListener);
    this.subscription2.unsubscribe();
  }

  ngAfterViewInit() {
  }

  closeNav(state: boolean) {
    if (state) {
      this.snav.close();
    }
  }

  inicio() {
    this.router.navigate(['/inicio']);
  }

  // Mini sidebar
}
