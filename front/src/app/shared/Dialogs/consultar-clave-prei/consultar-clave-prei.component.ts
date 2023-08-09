import {Component, Inject, OnInit} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {PreiService} from '../../../catalogos/servicios/prei.service';
import {Prei} from '../../../core/models/interfaceCatalogos/prei.interface';

export interface DialogData {
  id_prei: string;
}

@Component({
  selector: 'app-agregar-conceptos-contrato',
  templateUrl: './consultar-clave-prei.component.html',
  styleUrls: ['./consultar-clave-prei.component.css']
})
export class ConsultarClavePreiComponent implements OnInit {
  public id_prei='';
  public clave!: Prei;
  claveCuadro = '';
  error = false;

  constructor(
    public dialogRef: MatDialogRef<ConsultarClavePreiComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
    private _prei: PreiService,
  ) {
  }

  ngOnInit() {
    this.id_prei = this.data.id_prei;
    this._prei.getClavePrei('', this.id_prei).subscribe(res => {
      this.clave = res.Data;
      this.claveCuadro = this.clave.Grupo + '.' + this.clave.Generico + '.' + this.clave.Especifico + '.' + this.clave.Diferenciador
        + '.' + this.clave.Variable;
    }, err => {
      this.error = true;
    });
  }

  onNoClick(): void {
    this.dialogRef.close(false);
  }

  regresaDescripcion(valor: string): string {
    const descripcion = valor.split('.', 1);
    return descripcion.toString();
  }
}
