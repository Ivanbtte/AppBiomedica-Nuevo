import {Component, Inject, OnInit} from '@angular/core';
import {MatDialogRef, MAT_DIALOG_DATA} from '@angular/material/dialog';
import {EstatusService} from '../../core/servicios/estatus.service';
import Swal from 'sweetalert2';
import {SolicitudService} from '../../core/servicios/solicitud/solicitud.service';
import jwt_decode from 'jwt-decode';
import {UsuarioLogeado} from '../../core/models/UsuarioLogin.interface';

@Component({
  selector: 'app-cambiar-estatus',
  templateUrl: './cambiar-estatus.component.html',
  styleUrls: ['./cambiar-estatus.component.css']
})
export class CambiarEstatusComponent {
  resultad: any;
  valor = 0;
  resultado: any;
  public usuarioLog: UsuarioLogeado;

  constructor(
    public dialogRef: MatDialogRef<CambiarEstatusComponent>,
    @Inject(MAT_DIALOG_DATA) public data: any,
    private _estatus: EstatusService,
    private _solicitudes: SolicitudService,
  ) {
    this.usuarioLog  = jwt_decode(<string>localStorage.getItem('token'));
    this.resultado = data;
    console.log(this.resultado);
    this._estatus.ConsultaEstatus().subscribe(res => {
      this.resultad = res.Data;
      console.log(this.resultad);
    });
  }

  valorSeleccionado(value: any) {
    this.valor = value;
  }

  actualizarSolicitud() {
    Swal.fire({
      title: 'Confirmación',
      text: 'Estás seguro de cambiar el estatus',
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      cancelButtonText: 'Cancelar',
      confirmButtonText: 'Si, cambiar'
    }).then((result) => {
      if (result.value) {
        this._solicitudes.Actualizarolicitud(this.resultado.Id, 'aprobada', '', this.valor.toString(), this.usuarioLog.UsuarioId, {
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
          this.onNoClick();
        });

      }
    });
  }

  onNoClick(): void {
    this.dialogRef.close();
  }
}
