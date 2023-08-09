import {Component, Inject} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {UserPositions} from '../../Utils/UserPositions';
import * as _moment from 'moment';
const moment = _moment;

export interface DialogData {
  user: any;
}

@Component({
  selector: 'app-detalle-usuario',
  templateUrl: './detalle-usuario.component.html',
  styleUrls: ['./detalle-usuario.component.css']
})
export class DetalleUsuarioComponent {

  constructor(
    public dialogRef: MatDialogRef<DetalleUsuarioComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
  ) {
    console.log(data.user);
  }

  returnDepartment(): any {
    return UserPositions.returnDepartment(this.data.user);
  }


  onNoClick(): void {
    this.dialogRef.close();
  }
  formatExpeditionDate(valor: string): string {
    return moment(valor).format('DD-MM-YYYY');
  }

}
