import {Component, OnInit} from '@angular/core';
import {NombramientoService} from '../../core/servicios/nombramiento.service';
import jwt_decode from 'jwt-decode';
import {UsuarioLogeado} from '../../core/models/UsuarioLogin.interface';
import {ServiciosUsuariosService} from '../../core/servicios/ServiciosUsuarios/servicios-usuarios.service';
import Swal from 'sweetalert2';
import {SaiComponent} from '../../catalogos/sai/sai.component';
import {MatDialog} from '@angular/material/dialog';
import {AbrirCatalogoSaiService} from '../../core/servicios/abrir-catalogo-sai.service';
import {AreaServicioService} from '../../core/servicios/AreasServicios/area-servicio.service';
import {RespuestaPeticion} from '../../core/models/estructuras_generales';

@Component({
  selector: 'app-enfermeria-fondo-fijo',
  templateUrl: './enfermeria-fondo-fijo.component.html',
  styleUrls: ['./enfermeria-fondo-fijo.component.css']
})
export class EnfermeriaFondoFijoComponent implements OnInit {
  public usuarioLog!: UsuarioLogeado;
  public unidadesMedicas: any[]=[];
  public serviciosUsuarioResul: any[]=[];
  public areaServicios: any[]=[];
  public id_unidad: any;

  constructor(
    private _nombramiento: NombramientoService,
    private  _serviciosUsuarios: ServiciosUsuariosService,
    public dialog: MatDialog,
    private _servicioComunicacion: AbrirCatalogoSaiService,
    private _area_servicio: AreaServicioService
  ) {
  }

  ngOnInit() {
    this.usuarioLog = jwt_decode(<string>localStorage.getItem('token'));
    this.obtenerUnidadesMedicas();
  }

  valorSelectUnidad(value: string) {
    console.log(value);
    this.id_unidad = value;
    this.obtenerServiciosUsuarios(value, this.usuarioLog.UsuarioId);
  }

  valorSelectServicio(value: string) {
    this.obtenerAreaServicio(this.id_unidad, value);
  }

  obtenerUnidadesMedicas() {
    this._nombramiento.nombramiento(this.usuarioLog.PersonaId).subscribe(res => {
      const datos = res.Data;
      this.unidadesMedicas = datos.UnidadMed;
      console.log(this.unidadesMedicas);
    });
  }

  obtenerServiciosUsuarios(idUnidad: string, idUsuario: string) {
    this._serviciosUsuarios.serviciosUsuarios(idUnidad, idUsuario).subscribe(res => {
      console.log(res.Data);
      this.serviciosUsuarioResul = res.Data;
    });
  }

  obtenerAreaServicio(idUnidad: string, idServicio: string) {
    this._area_servicio.ConsultarAreaServicio(idUnidad, idServicio, '').subscribe(res => {
      console.log(res.Data);
      this.areaServicios = res.Data;
    }, err => {
      if (err.error !== '' || err.error !== undefined) {
        const erores = <RespuestaPeticion>err.error;
        this.areaServicios = [];
        console.log(erores.Data, erores.Mensaje);
        Swal.fire({
          title: 'Error aun no han registrado ninguna Area en este servicio',
          text: erores.Mensaje,
          icon: 'error',
        });
      }
      console.log(err.error);
    });
  }

  openDialog(): void {
    // localStorage.setItem('servicioId', this.valorServicio.toString());
    // localStorage.setItem('unidadId', this.valor.toString());
    Swal.fire({
      icon: 'info',
      text: 'Recuerda que no puedes agregar dos veces la misma clave por servicio',
    });
    // this.consultarServicio(this.valorServicio);
    this._servicioComunicacion.enviarEstadoComoponente(true);
    this._servicioComunicacion.enviarCuadroComponente('SI');
    // this.agregarSolicitud();
    const dialogRef = this.dialog.open(SaiComponent, {
      height: '80%',
      width: 'auto',
      disableClose: false,
    });
    dialogRef.afterClosed().subscribe(result => {
      // this.llenarTablaSolicitudes();
      this._servicioComunicacion.enviarEstadoComoponente(false);
      this._servicioComunicacion.enviarCuadroComponente('');
    });
  }
}
