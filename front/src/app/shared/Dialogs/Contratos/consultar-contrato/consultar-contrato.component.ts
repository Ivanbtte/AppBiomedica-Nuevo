import {Component, Inject, OnInit} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {ContratoService} from '../../../../core/servicios/ContratosServicios/contrato.service';
import {Contrato} from '../../../../core/models/contrato.interface';

export interface DialogData {
  num_contrato: string;
}

@Component({
  selector: 'app-consultar-contrato',
  templateUrl: './consultar-contrato.component.html',
  styleUrls: ['./consultar-contrato.component.css']
})
export class ConsultarContratoComponent implements OnInit {
  public Contrato!: Contrato;

  constructor(
    public dialogRef: MatDialogRef<ConsultarContratoComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
    private _contrato: ContratoService
  ) {
  }

  ngOnInit() {
    this._contrato.ConsultarContratos(this.data.num_contrato, '', '', '', '', '', '',
      '', '', '').subscribe(resultado => {
      this.Contrato = resultado.Data['registros'][0];
      console.log(this.Contrato);
    });
  }

  onNoClick(): void {
    this.dialogRef.close();
  }
}
