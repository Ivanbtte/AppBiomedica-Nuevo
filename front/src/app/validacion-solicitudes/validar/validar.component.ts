import {Component, OnInit} from '@angular/core';
import {SolicitudService} from '../../core/servicios/solicitud/solicitud.service';
import {MatDialog} from '@angular/material/dialog';
import {MatTableDataSource} from '@angular/material/table';
import {FiltrosSolicitudComponent} from '../filtros-solicitud/filtros-solicitud.component';

@Component({
  selector: 'app-validar',
  templateUrl: './validar.component.html',
  styleUrls: ['./validar.component.css']
})
export class ValidarComponent implements OnInit {
  displayedColumns: string[] = ['Nombre', 'TipoSolicitud', 'UnidadMedi', 'Estatus', 'Validar'];
  dataSource: any;
  resultado: any[]=[];

  constructor(
    private _solicitudes: SolicitudService,
    public dialog: MatDialog,
  ) {
  }

  ngOnInit() {
    this.consultarSolicitudes();
  }

  consultarSolicitudes() {
    this._solicitudes.ConsultarTodasSolicitudes('', 'true').subscribe(res => {
      this.resultado = res.Data;
      this.dataSource = new MatTableDataSource();
      this.dataSource.data = this.resultado;
      console.log(this.resultado);
    });
  }
  openDialog(): void {
    const dialogRef = this.dialog.open(FiltrosSolicitudComponent, {
      height: 'auto',
      width: 'auto',
      disableClose: true,
      data: {datos: ''}
    });

    dialogRef.afterClosed().subscribe(result => {
/*      this.chipsVer = result;
      if (this.chipsVer.length !== 0) {
        this.buscarPorTema();
      }*/
    });
  }
}
