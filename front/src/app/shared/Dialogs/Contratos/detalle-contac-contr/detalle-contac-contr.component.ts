import {Component, Inject, OnInit} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {ContactosContratoService} from '../../../../core/servicios/ContratosServicios/contactos-contrato.service';

export interface DialogData {
  n_contrato: string;
  representante: any;
}

@Component({
  selector: 'app-detalle-contac-contr',
  templateUrl: './detalle-contac-contr.component.html',
  styleUrls: ['./detalle-contac-contr.component.css']
})
export class DetalleContacContrComponent implements OnInit {
  correos: any[]=[];
  correoTodos='';
  todos = false;
  telefonos: any[]=[];

  constructor(
    public dialogRef: MatDialogRef<DetalleContacContrComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
    private  _contactoCont: ContactosContratoService,
  ) {
  }

  onNoClick(): void {
    this.dialogRef.close();
  }

  ngOnInit() {
    this.consultarContactos();
    this.ConsultarTelefonos();
  }

  consultarContactos() {
    this._contactoCont.ConsultaCorreosContrato(this.data.n_contrato).subscribe(res => {
      const correo = res.Data.map((ite:any) => {
        return {
          correo: ite.Correos.Correo
        };
      });
      this.correos = correo;
      const ids = this.correos.map(it => it.correo);
      this.correoTodos = ids.join(';');
      if (this.correos.length > 1) {
        this.todos = true;
      }
    });
  }

  ConsultarTelefonos() {
    this._contactoCont.ConsultaTelefonosContrato(this.data.n_contrato).subscribe(res => {
      console.log(res.Data);
      const telef = res.Data.map((ite:any) => {
        return {
          telefono: ite.Telefonos.Telefono,
          extension: ite.Extension
        };
      });
      this.telefonos = telef;
    });
  }
}
