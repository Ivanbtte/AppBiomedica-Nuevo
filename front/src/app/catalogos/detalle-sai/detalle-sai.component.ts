import {Component, Inject} from '@angular/core';
import {MatDialogRef, MAT_DIALOG_DATA} from '@angular/material/dialog';
import {SaiService} from '../servicios/sai.service';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';

export interface DialogData {
  id: string;
  estadoAgregar: boolean;
  estadoCantidad: boolean;
}

@Component({
  selector: 'app-detalle-sai',
  templateUrl: './detalle-sai.component.html',
  styleUrls: ['./detalle-sai.component.css']
})
export class DetalleSaiComponent {
  public clave: any;
  public hideBoton: boolean;
  public hideInpu: boolean;
  firstFormGroup: FormGroup;

  constructor(
    private _sai: SaiService,
    public dialogRef: MatDialogRef<DetalleSaiComponent>,
    private fb: FormBuilder,
    @Inject(MAT_DIALOG_DATA) public data: DialogData) {
    this.firstFormGroup = this.fb.group({
      CantidadSolicitada: [null, Validators.compose([Validators.pattern('^([0-9])*$')])],
    });
    this.hideBoton = data.estadoAgregar;
    this.hideInpu = data.estadoCantidad;
    this._sai.getClaveSai(data.id, '').subscribe(res => {
      this.clave = res.Data;
      console.log(this.clave);
    }, err => {
      console.log(err);
    });
  }

  onNoClick(): void {
    this.dialogRef.close();
  }

}
