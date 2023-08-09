import {Component, Inject} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {SolicitudService} from '../../core/servicios/solicitud/solicitud.service';
import {UsuarioLogeado} from '../../core/models/UsuarioLogin.interface';
import jwt_decode from 'jwt-decode';
import {BitacoraSolicitudService} from '../../core/servicios/bitacora/bitacora-solicitud.service';

export interface DialogData {
  id: any;
}

@Component({
  selector: 'app-detalle-seguimiento',
  templateUrl: './detalle-seguimiento.component.html',
  styleUrls: ['./detalle-seguimiento.component.css']
})
export class DetalleSeguimientoComponent {
  usuarioLog: UsuarioLogeado;
  solicitud: any;
  bitacora: any;
  visibleBitacora: boolean;

  constructor(
    public dialogRef: MatDialogRef<DetalleSeguimientoComponent>,
    private _solicitudes: SolicitudService,
    private _bitacoraSolicitud: BitacoraSolicitudService,
    @Inject(MAT_DIALOG_DATA) public data: DialogData
  ) {
    this.visibleBitacora = false;
    this.usuarioLog = jwt_decode(<string>localStorage.getItem('token'));
    this.solicitud = data.id;
    console.log(this.solicitud);
    this.consultarSeguimientoSolicitud();
    this.consultaBitacoraSolicitud();

  }

  consultaBitacoraSolicitud() {
    this._bitacoraSolicitud.ConsultarBitacoraSolicitud(this.solicitud.Id).subscribe(res => {
      this.bitacora = res.Data;
      console.log(res.Data);
    });
  }

  consultarSeguimientoSolicitud() {
    this._solicitudes.ConsultarSolicitud(this.solicitud.UsuarioId, this.solicitud.TipoSolicitudId, '').subscribe(res => {
      console.log(res.Data);
    });
  }

  formatLabel(value: number | null) {
    if (!value) {
      return 0;
    }
    if (value >= 1) {
      return value;
    }
    return value;
  }
}
