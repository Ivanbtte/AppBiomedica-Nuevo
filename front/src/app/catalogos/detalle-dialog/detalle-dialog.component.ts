import {Component, Inject} from '@angular/core';
import {MatDialogRef, MAT_DIALOG_DATA} from '@angular/material/dialog';
import {Prei} from '../../core/models/interfaceCatalogos/prei.interface';
import {PreiService} from '../servicios/prei.service';

export interface DialogData {
  id: string;
}

@Component({
  selector: 'app-detalle-dialog',
  templateUrl: './detalle-dialog.component.html',
  styleUrls: ['./detalle-dialog.component.css']
})
export class DetalleDialogComponent {
  public clave!: Prei;
  claveCuadro = '';

  constructor(
    private _prei: PreiService,
    public dialogRef: MatDialogRef<DetalleDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData) {
    this._prei.getClavePrei(data.id, '').subscribe(res => {
      this.clave = res.Data;
      this.claveCuadro = this.clave.Grupo + '.' + this.clave.Generico + '.' + this.clave.Especifico + '.' + this.clave.Diferenciador
        + '.' + this.clave.Variable;
    }, err => {
      console.log(err);
    });
  }

  onNoClick(): void {
    this.dialogRef.close();
  }

}
