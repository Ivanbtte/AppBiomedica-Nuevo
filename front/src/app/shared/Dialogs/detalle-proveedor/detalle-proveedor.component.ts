import {Component, Inject, OnInit} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {ProveedoresInterface} from '../../../core/models/proveedores.interface';

export interface DialogData {
  proveedor: ProveedoresInterface;
}

@Component({
  selector: 'app-detalle-proveedor',
  templateUrl: './detalle-proveedor.component.html',
  styleUrls: ['./detalle-proveedor.component.css']
})
export class DetalleProveedorComponent implements OnInit {
  detalle_proveedor!: ProveedoresInterface;
  vCardInfo = '';

  constructor(
    public dialogRef: MatDialogRef<DetalleProveedorComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
  ) {
  }

  ngOnInit() {
    this.detalle_proveedor = this.data.proveedor;
    console.log(this.detalle_proveedor);
    let name = this.detalle_proveedor.AliasEmpresa,
      surname = '',
      org = this.detalle_proveedor.NombreEmpresa,
      email = this.detalle_proveedor.Correo,
      tel = this.detalle_proveedor.Telefono,
      n_pro = this.detalle_proveedor.NProvImss;

    this.vCardInfo = `BEGIN:VCARD
VERSION:3.0
N:${surname};${name}
FN:${surname} ${name}
ORG:${org}
EMAIL:${email}
TEL;TYPE=voice,work,oref:${tel}
NOTE:Numero de proveedor imss  ${n_pro}
END:VCARD
`;
  }

  onNoClick(): void {
    this.dialogRef.close();
  }
}
