import {Component, OnInit} from '@angular/core';
import {UsuarioLogeado} from '../../core/models/UsuarioLogin.interface';
import jwt_decode from 'jwt-decode';
import {NombramientoService} from '../../core/servicios/nombramiento.service';
import {MatDialog} from '@angular/material/dialog';
import {MatTableDataSource} from '@angular/material/table';
import {SaiComponent} from '../../catalogos/sai/sai.component';
import {AbrirCatalogoSaiService} from '../../core/servicios/abrir-catalogo-sai.service';
import {ServiciosUnidadesService} from '../../core/servicios/servicios-unidades.service';
import {ServicioProformaService} from '../../catalogos/servicios/servicio-proforma.service';
import Swal from 'sweetalert2';
import {SolicitudService} from '../../core/servicios/solicitud/solicitud.service';
import {ClavesService} from '../../core/servicios/solicitud/claves.service';
import {Router} from '@angular/router';
import {RespuestaPeticion} from '../../core/models/estructuras_generales';
import {Solicitud} from '../../core/models/solicitudes.interface';
import {CantidadDialogComponent} from '../../shared/Dialogs/cantidad-dialog/cantidad-dialog.component';

@Component({
  selector: 'app-consumo',
  templateUrl: './consumo.component.html',
  styleUrls: ['./consumo.component.css']
})
export class ConsumoComponent implements OnInit {
  usuarioLog!: UsuarioLogeado;
  unidadesMedicas!: any[];
  fecha: any;
  dataSource: any;
  displayedColumns: string[] = ['Clave', 'Descripcion', 'Servicio', 'Precio', 'Cantidades', 'Total', 'Editar', 'Eliminar'];
  valor = 0;
  valorServicio = '';
  servicios: any;
  servicioNomb = '';
  nombreSolic = '';
  botonAgregar = true;
  estatusCarrito = false;
  /* *****  este valor tiene que salir de la base de datos  *****   */
  folioSolicitud = 'Sin asignar';
  unidadId: any;
  estadoSolicitud = true;
  conceptosBase: any[] = [];
  totalSolicitud: any;
  iva: any;
  subtotal: any;
  resultado!: Solicitud;
  filtroServicios: any;

  constructor(
    private _nombramiento: NombramientoService,
    public dialog: MatDialog,
    private _servicioComunicacion: AbrirCatalogoSaiService,
    private _serviciosUnidades: ServiciosUnidadesService,
    private _serviciosProforma: ServicioProformaService,
    private _Solicitudes: SolicitudService,
    private _conceptosSolicitud: ClavesService,
    private _router: Router,
  ) {
    const date = new Date();
    console.log(date);
    this.fecha = date.getUTCDate() + '/' + (date.getMonth() + 1) + '/' + date.getFullYear();
    console.log('Fecha actual', this.fecha);
  }

  openDialog(): void {
    localStorage.setItem('servicioId', this.valorServicio.toString());
    localStorage.setItem('unidadId', this.valor.toString());
    Swal.fire({
      icon: 'info',
      text: 'Recuerda que no puedes agregar dos veces la misma clave por servicio',
    });
    this.consultarServicio(this.valorServicio);
    this._servicioComunicacion.enviarEstadoComoponente(true);
    this.agregarSolicitud();
    const dialogRef = this.dialog.open(SaiComponent, {
      height: '80%',
      width: 'auto',
      disableClose: false,
    });
    dialogRef.afterClosed().subscribe(result => {
      this.llenarTablaSolicitudes();
      this._servicioComunicacion.enviarEstadoComoponente(false);
    });
  }

  agregarSolicitud() {
    if (!this.estatusCarrito) {
      console.log('estado carrito', this.estatusCarrito);
      this._Solicitudes.AgregarSolicitud({
        Id: '',
        UsuarioId: this.usuarioLog.UsuarioId,
        UnidadMedId: this.valor,
        TipoSolicitudId: 1,
        EstatusId: 1,
        FechaCreacion: this.fecha,
        Total: this.totalSolicitud,
        Correo: this.usuarioLog.Correo,
        Estado: false
      }).subscribe(res => {
        const resultadoSolicitud = res.Data;
        this.folioSolicitud = resultadoSolicitud.Id;
        localStorage.setItem('solicitudId', resultadoSolicitud.Id);
        localStorage.setItem('unidadId', resultadoSolicitud.UnidadMedId);
        console.log(res);
      });
    }
    this.estatusCarrito = true;
  }

  consultarServicio(id: string) {
    this._serviciosProforma.getServicio(id).subscribe(res => {
      const servicioNombre = res.Data;
      this.servicioNomb = servicioNombre.Descripcion;
      console.log(this.servicioNomb);
    });
  }

  consultarSolictud() {
    this._Solicitudes.ConsultarSolicitud(this.usuarioLog.UsuarioId, '1', '').subscribe(res => {
      this.resultado = res.Data;
      console.log(this.resultado);
    });
  }

  ngOnInit() {
    this.usuarioLog = jwt_decode(<string>localStorage.getItem('token'));
    this._Solicitudes.ConsultarSolicitud(this.usuarioLog.UsuarioId, '1', '').subscribe(res => {
      this.resultado = res.Data;
      console.log(this.resultado);
      if (this.resultado) {
        this.valor = this.resultado.UnidadMedId;
        this.valorSelecUnidades(this.valor.toString());
        if (this.resultado.EstatusId === 1) {
          this.estatusCarrito = true;
          this.folioSolicitud = this.resultado.Id;
          localStorage.setItem('solicitudId', this.resultado.Id);
          this.llenarTablaSolicitudes();
        } else {
          this.estadoSolicitud = false;
          Swal.fire({
            icon: 'warning',
            title: 'Ya tienes una solicitud en proceso, si deseas acualizar da clic al boton Editar',
            showCancelButton: true,
            confirmButtonColor: '#3085d6',
            cancelButtonColor: '#d33',
            cancelButtonText: 'Cancelar',
            confirmButtonText: 'Editar'
          }).then((result) => {
            if (result.value) {
              this._router.navigate(['/solicitudes/editar']);
            }
          });
        }
      } else {
        this.valor = 0;
        console.log('Estoy vacio', res.Data);
      }
    });
    this.nombreSolic = this.usuarioLog.Nombre + ' ' + this.usuarioLog.AppellidoP + ' ' + this.usuarioLog.ApellidoM;
    this._nombramiento.nombramiento(this.usuarioLog.PersonaId).subscribe(res => {
      const datos = res.Data;
      this.unidadesMedicas = datos.UnidadMed;
      console.log(this.unidadesMedicas);
    });
  }

   eliminarClave(clave: any) {
    console.log(clave);
    // const { value: comentario }  =  Swal.fire({
    //   icon: 'warning',
    //   title: '¿Estás seguro?',
    //   text: '¡No podrás revertir esto!',
    //   input: 'text',
    //   inputPlaceholder: 'Motivo',
    //   showCancelButton: true,
    //   cancelButtonText: 'Cancelar',
    //   confirmButtonColor: '#3085d6',
    //   cancelButtonColor: '#d33',
    //   confirmButtonText: 'Si, borrar',
    //   inputValidator: (value) => {
    //     if (!value) {
    //       return 'Necesitas escribir un comentario';
    //     }
    //   }
    // });
    // if (comentario) {
    //   this._conceptosSolicitud.EliminarClave(clave.Id, comentario).subscribe(res => {
    //     this.llenarTablaSolicitudes();
    //     Swal.fire({
    //       icon: 'success',
    //       title: 'Correcto',
    //       text: res.Mensaje,
    //     });
    //   }, err => {
    //     if (err.error !== '' || err.error !== undefined) {
    //       const erores = <RespuestaPeticion> err.error;
    //       Swal.fire({
    //         icon: 'error',
    //         title: 'Error',
    //         text: erores.Mensaje,
    //       });
    //     }
    //   });
    // }
  }

  openDialogCantidad(clave: any): void {
    const dialogRef = this.dialog.open(CantidadDialogComponent, {
      height: '85%',
      width: '70%',
      disableClose: true,
      data: {
        clave: clave,
        agregar: false,
        editar: true,
        aprobar: false
      }
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.llenarTablaSolicitudes();
      }

    });
  }

  llenarTablaSolicitudes() {
    this._conceptosSolicitud.ConsultarClaves(this.folioSolicitud, '', '', '', '').subscribe(res => {
      this.conceptosBase = res.Data;
      console.log(this.conceptosBase);
      this.dataSource = new MatTableDataSource();
      this.dataSource.data = this.conceptosBase;
      this.mapeoServicios(this.conceptosBase);
      if (this.conceptosBase.length > 0) {
        this.totalSolicitud = this.calcularIva();
        this.subtotal = this.getTotalCost();
        this.iva = this.getTotal();
      }
    });
  }

  /* *****  funcion en la que mapeo los servicios que hay en esta solicitud  *****   */
  mapeoServicios(arr: any[]) {
    this.filtroServicios = arr.map(item => {
        return {
          ServicioId: item.ServiciosProformaId,
          Servicio: item.ServiciosProforma.Descripcion
        };
      }
    );
    console.log(this.filtroServicios);
  }

  getTotalCost() {
    return this.conceptosBase.map(t => t.Total).reduce((acc, value) => acc + value);
  }

  calcularIva() {
    return (1.16 * this.getTotalCost());
  }

  getTotal() {
    return this.calcularIva() - this.getTotalCost();
  }

  enviarSolicitud() {
    this.consultarSolictud();
    Swal.fire({
      title: '¿Estás seguro?',
      text: '¡No podrás revertir esto!',
      icon: 'warning',
      showCancelButton: true,
      cancelButtonText: 'Cancelar',
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      confirmButtonText: 'Si, Enviar'
    }).then((result) => {
      if (result.value) {
        this.resultado.EstatusId = 2;
        this.resultado.Estado = true;
        this.resultado.Total = this.totalSolicitud;
        console.log('esto se envia', this.resultado);
        this._Solicitudes.Actualizarolicitud(this.resultado.Id, 'Solicitud enviada', '', '', '', {
          Id: this.resultado.Id,
          UsuarioId: this.resultado.UsuarioId,
          UnidadMedId: this.resultado.UnidadMedId,
          TipoSolicitudId: this.resultado.TipoSolicitudId,
          EstatusId: 2,
          FechaCreacion: '',
          Total: this.totalSolicitud,
          Correo: '',
          Estado: true
        }).subscribe(res => {
          Swal.fire({
            icon: 'success',
            title: 'Correcto',
            text: 'Solicitud enviada',
          });
          this._router.navigate(['/solicitudes/seguimiento']);
        }, err => {
          if (err.error !== '' || err.error !== undefined) {
            const erores = <RespuestaPeticion> err.error;
            console.log(erores.Data);
            Swal.fire({
              position: 'center',
              title: erores.Data,
              text: 'Consulta las fechas publicadas en la pagina de inicio',
              icon: 'error',
              showConfirmButton: true,
            });
          }
          console.log(err.error);
        });
      }
    });
  }

  valorSelecUnidades(valor: string) {
    this._serviciosUnidades.serviciosUnidades(valor).subscribe(res => {
      const resultado = res.Data;
      this.servicios = resultado;
      console.log('servicios', this.servicios);
    });
    this.valor = parseInt(valor, 10);
    console.log('valor:', this.valor);
  }

  valorSelecServicios(valor: string) {
    this.valorServicio = valor;
    console.log('Estos son los valores :=>', this.valor, this.valorServicio);
    if (this.valor !== 0 && this.valorServicio !== '') {
      this.botonAgregar = false;
    }
  }
}
