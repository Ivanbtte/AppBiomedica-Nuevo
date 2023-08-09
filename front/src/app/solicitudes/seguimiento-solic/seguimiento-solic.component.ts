import {Component, OnInit} from '@angular/core';
import {SolicitudService} from '../../core/servicios/solicitud/solicitud.service';
import {UsuarioLogeado} from '../../core/models/UsuarioLogin.interface';
import jwt_decode from 'jwt-decode';
import {MatTableDataSource} from '@angular/material/table';
import {MatDialog} from '@angular/material/dialog';
import {DetalleSeguimientoComponent} from '../detalle-seguimiento/detalle-seguimiento.component';

@Component({
  selector: 'app-seguimiento-solic',
  templateUrl: './seguimiento-solic.component.html',
  styleUrls: ['./seguimiento-solic.component.css']
})
export class SeguimientoSolicComponent implements OnInit {
  usuarioLog: UsuarioLogeado;
  displayedColumns: string[] = ['Folio', 'Solicitante', 'TipoS', 'Estatus', 'Total', 'Fecha', 'Detalle'];
  dataSource: any;
  Solicitudes: any[] = [];

  constructor(
    private _Solicitudes: SolicitudService,
    public dialog: MatDialog,
  ) {
    this.usuarioLog = jwt_decode(<string> localStorage.getItem('token'));

  }

  ngOnInit() {
    this.llenarTabla();
  }

  filterEnviados(array: any[]): any[] | undefined {
    if (array.length > 0) {
      const arreFinal: any[] = array.filter(item => {
        return item.EstatusId !== 1;
      });
      this.Solicitudes = arreFinal;

      console.log(this.Solicitudes);
      return this.Solicitudes;
    }
  }

  openDialogClave(id: any): void {
    const dialogRef = this.dialog.open(DetalleSeguimientoComponent, {
      height: '90%',
      width: '70%',
      data: {
        id: id,
      }
    });

    dialogRef.afterClosed().subscribe(result => {
    });
  }

  llenarTabla() {
    this._Solicitudes.ConsultarTodasSolicitudes(this.usuarioLog.UsuarioId, '').subscribe(res => {
      console.log(res.Data);
      this.dataSource = new MatTableDataSource();
      this.dataSource.data = this.filterEnviados(res.Data);
    });
  }
}
