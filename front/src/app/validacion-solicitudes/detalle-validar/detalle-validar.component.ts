import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {MatDialog} from '@angular/material/dialog';
import {MatTableDataSource} from '@angular/material/table';
import {ClavesService} from '../../core/servicios/solicitud/claves.service';
import {SolicitudService} from '../../core/servicios/solicitud/solicitud.service';
import jwt_decode from 'jwt-decode';
import {UsuarioLogeado} from '../../core/models/UsuarioLogin.interface';
import {Solicitud} from '../../core/models/solicitudes.interface';
import {ConsultasAnaliticasComponent} from '../consultas-analiticas/consultas-analiticas.component';
import Swal from 'sweetalert2';

@Component({
  selector: 'app-detalle-validar',
  templateUrl: './detalle-validar.component.html',
  styleUrls: ['./detalle-validar.component.css']
})
export class DetalleValidarComponent implements OnInit {
   id: any;
   conceptosBase: any[]=[];
  resultado: any;
  informacion = true;
  textBtnInf = 'Solicitud';
  /*
    displayedColumns: string[] = ['Clave', 'Descripcion', 'Servicio', 'Precio', 'Cantidades', 'Total', 'Editar', 'Eliminar'];
  */
  displayedColumns: string[] = ['Validado', 'Clave', 'Descripcion', 'Servicio', 'Precio', 'Cantidades', 'Aprobada', 'CPM', 'Total', 'Editar'];
  dataSource: any;
   usuarioLog: UsuarioLogeado;
   solicitudResultado!: Solicitud;
  totalSolicitud: any;
  iva: any;
  subtotal: any;

  constructor(
    private _route: ActivatedRoute,
    private _conceptosSolicitud: ClavesService,
    private _solicitudes: SolicitudService,
    public dialog: MatDialog,
    private _router: Router,
  ) {
    this.usuarioLog = jwt_decode(<string>localStorage.getItem('token'));
  }

  ngOnInit() {
    this.id = this._route.snapshot.paramMap.get('id');
    this.consultarSolicitudes();
  }

  consultarSolicitudes() {
    this._solicitudes.ConsultarSolicitud('', '', this.id).subscribe(res => {
      this.resultado = res.Data;
      this.solicitudResultado = res.Data;
      console.log(this.resultado);
      console.log(this.solicitudResultado);
    });
    this.llenarTablaSolicitudes();
  }

  llenarTablaSolicitudes() {
    this._conceptosSolicitud.ConsultarClaves(this.id, '', '', '', '').subscribe(res => {
      console.log(res.Data);
      this.conceptosBase = res.Data;
      if (this.conceptosBase.length === 0) {
        this._solicitudes.Actualizarolicitud(this.resultado.Id, 'aprobada', '', '6', this.usuarioLog.UsuarioId, {
          Id: this.resultado.Id,
          UsuarioId: this.resultado.UsuarioId,
          UnidadMedId: this.resultado.UnidadMedId,
          TipoSolicitudId: this.resultado.TipoSolicitudId,
          EstatusId: this.resultado.EstatusId,
          FechaCreacion: this.resultado.FechaCreacion,
          Total: this.resultado.Total,
          Correo: this.resultado.Correo,
          Estado: this.resultado.Estado
        }).subscribe(res => {
          Swal.fire({
            icon: 'success',
            title: 'Correcto',
            text: 'Solicitud Aprobada',
          });
          this._router.navigate(['/requerimientos/revision']);
        });
      } else {
        this.dataSource = new MatTableDataSource();
        this.dataSource.data = this.conceptosBase;
        this.totalSolicitud = this.calcularIva();
        this.subtotal = this.getTotalCost();
        this.iva = this.getTotal();
      }

      console.log(this.conceptosBase);
    });
  }

  filtroNoAprobadas(arr: any[]): any[] {
    const filt = arr.filter(item => {
      return item.CantidadAprobada === 0;
    });
    return filt;
  }

  /* *****  Runcion para revisar unasolicitud  *****   */
  revisarSolicitud() {
    if (this.resultado.EstatusId === 2) {
      console.log('soy esttus 2');
      this._solicitudes.Actualizarolicitud('', 'Solicitud Revisada', '', '3', this.usuarioLog.UsuarioId, {
        Id: this.solicitudResultado.Id,
        UsuarioId: this.solicitudResultado.UsuarioId,
        UnidadMedId: this.solicitudResultado.UnidadMedId,
        TipoSolicitudId: this.solicitudResultado.TipoSolicitudId,
        EstatusId: this.solicitudResultado.EstatusId,
        FechaCreacion: this.solicitudResultado.FechaCreacion,
        Total: this.solicitudResultado.Total,
        Correo: this.solicitudResultado.Correo,
        Estado: this.solicitudResultado.Estado
      }).subscribe(res => {
        console.log(res.Data);
        this.consultarSolicitudes();
      });

    }
    this.informacion = !this.informacion;
    if (this.informacion) {
      this.textBtnInf = 'Solicitud';
    } else {
      this.textBtnInf = 'Información';
    }
  }

  aprobarSolicitud() {
    Swal.fire({
      title: 'Confirmación',
      text: 'Estás seguro de aprobar esta solicitud',
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      cancelButtonText: 'Cancelar',
      confirmButtonText: 'Si, aprobar'
    }).then((result) => {
      if (result.value) {
        this._solicitudes.Actualizarolicitud(this.resultado.Id, 'aprobada', '', '6', this.usuarioLog.UsuarioId, {
          Id: this.resultado.Id,
          UsuarioId: this.resultado.UsuarioId,
          UnidadMedId: this.resultado.UnidadMedId,
          TipoSolicitudId: this.resultado.TipoSolicitudId,
          EstatusId: this.resultado.EstatusId,
          FechaCreacion: this.resultado.FechaCreacion,
          Total: this.resultado.Total,
          Correo: this.resultado.Correo,
          Estado: this.resultado.Estado
        }).subscribe(res => {
          Swal.fire({
            icon: 'success',
            title: 'Correcto',
            text: res.Mensaje,
          });
          this._router.navigate(['/requerimientos/revision']);
        });

      }
    });
  }

  consultasAnaliticas(clave: string): void {
    const dialogRef = this.dialog.open(ConsultasAnaliticasComponent, {
        height: '95%',
        width: '95%',
        disableClose: true,
        data: {
          idSolicitud: this.id,
          clave: clave,
          clavesSolicitud: this.conceptosBase
        }
      })
    ;

    dialogRef.afterClosed().subscribe(result => {
      this.llenarTablaSolicitudes();
      /*      this.chipsVer = result;
            if (this.chipsVer.length !== 0) {
              this.buscarPorTema();
            }*/
    });
  }

  comprobarValidados(clave: any): boolean {
    return clave.Estado === true;
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
}
