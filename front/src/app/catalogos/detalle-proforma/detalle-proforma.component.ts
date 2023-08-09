import {Component, Inject} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {ProformaService} from '../servicios/proforma.service';
import {ServicioProformaService} from '../servicios/servicio-proforma.service';
import {ServicioProforma} from '../../core/models/interfaceCatalogos/servicioPro.interface';

export interface DialogData {
  id: string;
}

@Component({
  selector: 'app-detalle-proforma',
  templateUrl: './detalle-proforma.component.html',
  styleUrls: ['./detalle-proforma.component.css']
})
export class DetalleProformaComponent {
  public clave: any;
  public servicios: ServicioProforma[] = [];
  public nombreServicio = '';

  constructor(
    private _proforma: ProformaService,
    private _servicios: ServicioProformaService,
    public dialogRef: MatDialogRef<DetalleProformaComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData) {
    this._proforma.getClaveProforma(data.id).subscribe(res => {
      this.clave = res.Data;
      console.log(this.clave);
    }, err => {
      console.log(err);
    });

    this._servicios.getClaveProformaServicios(data.id).subscribe(res => {
      this.servicios = res.Data;
      const nomb = this.servicios.map(it => it.Descripcion);
      this.nombreServicio = nomb.join();
      console.log(this.servicios);
    }, err => {
      console.log(err);
    });
  }

  onNoClick(): void {
    this.dialogRef.close();
  }
}
