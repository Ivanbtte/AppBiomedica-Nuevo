import {Component, OnInit, Inject} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {ContactoProveedor} from '../../../core/models/proveedores.interface';

export interface DialogData {
  contacto: ContactoProveedor;
}

@Component({
  selector: 'app-detalle-contacto-proveedor',
  templateUrl: './detalle-contacto-proveedor.component.html',
  styleUrls: ['./detalle-contacto-proveedor.component.css']
})
export class DetalleContactoProveedorComponent implements OnInit {
  contacto!: ContactoProveedor;
  vCardInfo='';

  constructor(
    public dialogRef: MatDialogRef<DetalleContactoProveedorComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData
  ) {
  }

  ngOnInit() {
    this.contacto = this.data.contacto;
    console.log(this.contacto);
    let name = this.contacto.NombreCompleto,
      surname = '',
      org = this.contacto.Proveedor.NombreEmpresa,
      email = this.contacto.Correo,
      nota = this.contacto.Comentarios,
      cargo = this.contacto.Cargo,
      tel = this.contacto.Celular;

    this.vCardInfo = `BEGIN:VCARD
VERSION:3.0
N:${surname};${name}
FN:${surname} ${name}
NOTE:${nota}
ORG:${org}
EMAIL:${email}
TITLE:${cargo}
TEL;TYPE=voice,work,oref:${tel}
END:VCARD
`;
  }

  onNoClick(): void {
    this.dialogRef.close();
  }
}
