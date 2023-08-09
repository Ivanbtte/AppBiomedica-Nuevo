import {Component, OnInit} from '@angular/core';
import {ClavesService} from '../../core/servicios/solicitud/claves.service';
import {MatTableDataSource} from '@angular/material/table';

@Component({
  selector: 'app-solicitud-final',
  templateUrl: './solicitud-final.component.html',
  styleUrls: ['./solicitud-final.component.css']
})
export class SolicitudFinalComponent implements OnInit {
  displayedColumns: string[] = ['Clave', 'Descripcion', 'Servicio', 'Unidad', 'Precio', 'Aprobada', 'Total'];
  dataSource: any;
  serviciosSelect: any;
  valorServicio = '';
  datosClaves: any;

  constructor(
    private _conceptosSolicitud: ClavesService,
  ) {
  }

  ngOnInit() {
    this.llenarTabla();
  }

  llenarTabla() {
    this._conceptosSolicitud.ConsultarClaves('', '', '', 'true', this.valorServicio).subscribe(res => {
      console.log(res.Data);
      this.datosClaves = res.Data;
      this.dataSource = new MatTableDataSource();
      this.dataSource.data = res.Data;
      this.agruparServicios(this.datosClaves);
    });
  }

  valorSelecServicios(valor: string) {
    console.log(valor);
    this._conceptosSolicitud.ConsultarClaves('', '', '', 'true', this.valorServicio).subscribe(res => {
      this.dataSource = new MatTableDataSource();
      this.dataSource.data = res.Data;
    });
  }

  agruparServicios(arr: any[]) {
    const servicios = arr.map(ite => {
      return {
        id: ite.ServiciosProformaId,
        servicio: ite.ServiciosProforma.Descripcion
      };
    });
    // var matriz = {};
    //
    // servicios.forEach(function (registro) {
    //   var pais = registro['id'];
    //   var descripcion = registro['servicio'];
    //   matriz[pais] = matriz[pais] ? matriz[pais] : descripcion;
    // });
    // matriz = Object.keys(matriz).map(function (pais) {
    //   return {id: pais, servicio: matriz[pais]};
    // });
    // this.serviciosSelect = matriz;
    // console.log(matriz);
    // console.log(servicios);
  }
}
