import {Component, OnInit} from '@angular/core';
import {Router} from '@angular/router';
import {UsuarioLogeado} from '../../core/models/UsuarioLogin.interface';
import jwt_decode from 'jwt-decode';
import {SolicitudService} from '../../core/servicios/solicitud/solicitud.service';
import {ClavesService} from '../../core/servicios/solicitud/claves.service';
import Swal from 'sweetalert2';

@Component({
  selector: 'app-editar-solic',
  templateUrl: './editar-solic.component.html',
  styleUrls: ['./editar-solic.component.css']
})
export class EditarSolicComponent implements OnInit {
  public usuarioLog: UsuarioLogeado;
  public solicitudes: any[]=[];
  public conceptos: any;

  constructor(
    private router: Router,
    private _Solicitudes: SolicitudService,
    private _Conceptos: ClavesService
  ) {
    this.usuarioLog= jwt_decode(<string>localStorage.getItem('token'));
  }

  ngOnInit() {
    this.llenarDivs();
  }

  /*  filtrarEnviados(arreglo: any[]): any[] {
      if (arreglo.length > 0) {
        const arreFinal = arreglo.filter(item => {
          console.log(item.EstatusId);
          return item.EstatusId < 6;
        });
        console.log(arreFinal);
        return arreFinal;
      }
    }*/
  checarEstado(clave: any): string {
    console.log(clave);
    if (clave.EstatusId > 5) {
      return 'Ver solicitud aprobada';
    } else {
      return 'Editar solicitud';
    }
  }

  llenarDivs() {
    this._Solicitudes.ConsultarTodasSolicitudes(this.usuarioLog.UsuarioId, '').subscribe(res => {
      this.solicitudes = res.Data;
      console.log(this.solicitudes);
    });
  }

  darSeguimiento() {
    this.router.navigate(['/solicitudes/seguimiento']);
  }

  cancelarsolicitud(clave: any) {
    console.log(clave);
    Swal.fire({
      title: '¿Estás seguro?',
      text: '¡No podrás revertir esto!',
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      cancelButtonText: 'No cancelar',
      confirmButtonText: 'Si, cancelar'
    }).then((result) => {
      if (result.value) {
        this._Solicitudes.Actualizarolicitud(clave.Id, 'Solicitud cancelada', '', '', '', {
          Id: clave.Id,
          UsuarioId: clave.UsuarioId,
          UnidadMedId: 0,
          TipoSolicitudId: 0,
          EstatusId: 17,
          FechaCreacion: '',
          Total: clave.Total,
          Correo: '',
          Estado: false
        }).subscribe(res => {
          Swal.fire(
            'Cancelada',
            'Tu solicitud ha sido cancelada',
            'success'
          );
          this.llenarDivs();
        });
      }
    });
  }
}
