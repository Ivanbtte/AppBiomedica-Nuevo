import {Component, OnInit} from '@angular/core';
import {ActivatedRoute} from '@angular/router';
import {UsuarioLogeado} from '../../core/models/UsuarioLogin.interface';
import jwt_decode from 'jwt-decode';
import {SolicitudService} from '../../core/servicios/solicitud/solicitud.service';
import {ServiciosUnidadesService} from '../../core/servicios/servicios-unidades.service';
import {ClavesService} from '../../core/servicios/solicitud/claves.service';
import {MatDialog} from '@angular/material/dialog';
import {MatTableDataSource} from '@angular/material/table';
import Swal from 'sweetalert2';
import {SaiComponent} from '../../catalogos/sai/sai.component';
import {ServicioProformaService} from '../../catalogos/servicios/servicio-proforma.service';
import {AbrirCatalogoSaiService} from '../../core/servicios/abrir-catalogo-sai.service';
import {CantidadDialogComponent} from '../../shared/Dialogs/cantidad-dialog/cantidad-dialog.component';

@Component({
  selector: 'app-detalle-editar',
  templateUrl: './detalle-editar.component.html',
  styleUrls: ['./detalle-editar.component.css']
})
export class DetalleEditarComponent implements OnInit {
  id: any;
  usuarioLog: UsuarioLogeado;
  solicitud: any;
  selectServicios: any;
  valorServicio = '';
  botonAgregar = true;
  displayedColumns: string[] = ['Clave', 'Descripcion', 'Servicio', 'Precio', 'Cantidades', 'Aprobada', 'CPM', 'Total', 'Editar', 'Eliminar'];
  dataSource: any;
  conceptosBase: any[] = [];
  totalSolicitud: any;
  iva: any;
  subtotal: any;
  actualizar = false;
  ticket = false;

  constructor(
    private _route: ActivatedRoute,
    private _solicitudes: SolicitudService,
    private _serviciosUnidades: ServiciosUnidadesService,
    private _conceptosSolicitud: ClavesService,
    private _serviciosProforma: ServicioProformaService,
    private _servicioComunicacion: AbrirCatalogoSaiService,
    public dialog: MatDialog,
  ) {
    this.usuarioLog = jwt_decode(<string> localStorage.getItem('token'));
  }

  ngOnInit() {
    this.id = this._route.snapshot.paramMap.get('id');
    console.log(this.id);
    this.consultarSolicitud();
    this.llenarTablaSolicitudes();
  }

  llenarTablaSolicitudes() {
    this._conceptosSolicitud.ConsultarClaves(this.id, '', '', '', '').subscribe(res => {
      this.conceptosBase = res.Data;
      console.log(this.conceptosBase);
      this.totalSolicitud = this.calcularIva();
      this.subtotal = this.getTotalCost();
      this.iva = this.getTotal();
      this.dataSource = new MatTableDataSource();
      this.dataSource.data = this.conceptosBase;
      if (this.actualizar === true) {
        this.actualizarTotal();
      }
    });
    this.actualizar = false;
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

  consultarSolicitud() {
    this._solicitudes.ConsultarSolicitud('', '', this.id).subscribe(res => {
      this.solicitud = res.Data;
      this.consultarServicios();
      this.comprobarEstado();
      console.log('respuesta', this.solicitud);
    });
  }

  comprobarEstado() {
    if (this.solicitud.EstatusId > 5) {
      console.log(this.solicitud.EstatusId);
      this.ticket = true;
      this.displayedColumns.pop();
      this.displayedColumns.pop();
    }
  }

  consultarServicios() {
    this._serviciosUnidades.serviciosUnidades(this.solicitud.UnidadMedId).subscribe(res => {
      this.selectServicios = res.Data;
      console.log(this.selectServicios);
    });
  }

  valorSelecServicios(valor: string) {
    this.valorServicio = valor;
    if (this.valorServicio !== '') {
      this.botonAgregar = false;
    }
  }

  openDialogCantidad(clave: any): void {
    const dialogRef = this.dialog.open(CantidadDialogComponent, {
      height: 'auto',
      width: '65%',
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
        this.actualizar = true;
      }

    });
  }

  openDialog(): void {
    localStorage.setItem('servicioId', this.valorServicio.toString());
    localStorage.setItem('solicitudId', this.id);
    localStorage.setItem('unidadId', this.solicitud.UnidadMedId);
    Swal.fire({
      icon: 'info',
      text: 'Recuerda que no puedes agregar dos veces la misma clave por servicio',
    });
    this._servicioComunicacion.enviarEstadoComoponente(true);
    const dialogRef = this.dialog.open(SaiComponent, {
      height: '80%',
      width: 'auto',
      disableClose: false,
    });
    dialogRef.afterClosed().subscribe(result => {
      this.llenarTablaSolicitudes();
      this._servicioComunicacion.enviarEstadoComoponente(false);
      this.actualizar = true;
    });
  }

  async eliminarClave(clave: any) {
    console.log(clave);
    // const {value: comentario} = await Swal.fire({
    //   type: 'warning',
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
    //     this.actualizar = true;
    //     Swal.fire({
    //       type: 'success',
    //       title: 'Correcto',
    //       text: res.Mensaje,
    //     });
    //   }, err => {
    //     if (err.error !== '' || err.error !== undefined) {
    //       const erores = <RespuestaPeticion> err.error;
    //       Swal.fire({
    //         type: 'error',
    //         title: 'Error',
    //         text: erores.Mensaje,
    //       });
    //     }
    //   });
    // }
  }

  actualizarTotal() {
    console.log('esto es el total a actualizar we ', this.totalSolicitud);
    this._solicitudes.Actualizarolicitud(this.id, '', this.totalSolicitud.toString(), '', '', {
      Id: this.id,
      UsuarioId: '',
      UnidadMedId: 0,
      TipoSolicitudId: 0,
      EstatusId: 0,
      FechaCreacion: '',
      Total: this.totalSolicitud,
      Correo: '',
      Estado: false
    }).subscribe(res => {
      console.log(res.Data);
    });
  }

  comprobarAprobacion(value: any): boolean {
    return value.CantidadAprobada !== 0;
  }
}
