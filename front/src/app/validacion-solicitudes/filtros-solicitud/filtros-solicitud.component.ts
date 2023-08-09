import {Component, Inject} from '@angular/core';
import {MatDialogRef, MAT_DIALOG_DATA} from '@angular/material/dialog';

@Component({
  selector: 'app-filtros-solicitud',
  templateUrl: './filtros-solicitud.component.html',
  styleUrls: ['./filtros-solicitud.component.css']
})
export class FiltrosSolicitudComponent {

  constructor(
    public dialogRef: MatDialogRef<FiltrosSolicitudComponent>,
    @Inject(MAT_DIALOG_DATA) public data: any
  ) {
  }

/*  onNoClick(arr: any[]): void {
    const seleccionados = arr.map(it => it.value);
    this.dialogRef.close(seleccionados);
  }*/

  onNoClick() {
    this.dialogRef.close();
  }

}
