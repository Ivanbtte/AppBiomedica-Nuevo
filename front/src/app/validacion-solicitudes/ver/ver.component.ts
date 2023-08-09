import {Component, OnInit} from '@angular/core';
import {SolicitudService} from '../../core/servicios/solicitud/solicitud.service';
import {MatDialog} from '@angular/material/dialog';
import {MatTableDataSource} from '@angular/material/table';
import {CambiarEstatusComponent} from '../cambiar-estatus/cambiar-estatus.component';

@Component({
  selector: 'app-ver',
  templateUrl: './ver.component.html',
  styleUrls: ['./ver.component.css']
})
export class VerComponent implements OnInit {
  displayedColumns: string[] = ['Nombre', 'TipoSolicitud', 'UnidadMedi', 'Estatus', 'Validar2'];
  dataSource: any;
  resultado: any[] = [];

  constructor(
    private _solicitudes: SolicitudService,
    public dialog: MatDialog,
  ) {
  }

  ngOnInit() {
    this.consultarSolicitudes();
  }

  consultarSolicitudes() {
    this._solicitudes.ConsultarTodasSolicitudes('', 'false').subscribe(res => {
      this.resultado = res.Data;
      this.dataSource = new MatTableDataSource();
      this.dataSource.data = this.resultado;
      console.log(this.resultado);
    });
  }

  cambiarEstatus(clave: any): void {
    const dialogRef = this.dialog.open(CambiarEstatusComponent, {
        height: 'auto',
        width: 'auto',
        disableClose: true,
        data: clave
      })
    ;

    dialogRef.afterClosed().subscribe(result => {
      this.consultarSolicitudes();
      // this.llenarTablaSolicitudes();
      /*      this.chipsVer = result;
            if (this.chipsVer.length !== 0) {
              this.buscarPorTema();
            }*/
    });
  }
}
