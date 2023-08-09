import {Component, Inject, OnInit} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';

export interface DialogData {
  id_concepto: string;
}

@Component({
  selector: 'app-verificar-distribucion',
  templateUrl: './verificar-distribucion.component.html',
  styleUrls: ['./verificar-distribucion.component.css']
})
export class VerificarDistribucionComponent implements OnInit {

  constructor(
    public dialogRef: MatDialogRef<VerificarDistribucionComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData
  ) {
  }

  ngOnInit() {
  }

  onNoClick(): void {
    this.dialogRef.close();
  }

}
