import {Component, Inject, OnInit} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {Contrato} from '../../../../core/models/contrato.interface';
import Swal from 'sweetalert2';
import {FormControl, Validators} from '@angular/forms';
import {ContratoService} from '../../../../core/servicios/ContratosServicios/contrato.service';
import {ConceptosContratoService} from '../../../../core/servicios/ContratosServicios/conceptos-contrato.service';

export interface DialogData {
  num_contrato: string;
  total_contratado: number;
}

@Component({
  selector: 'app-examinar-monto-contrato',
  templateUrl: './examinar-monto-contrato.component.html',
  styleUrls: ['./examinar-monto-contrato.component.css']
})
export class ExaminarMontoContratoComponent implements OnInit {
  Contrato!: Contrato;
  MontoContratado!: number;
  Conceptos: any[]=[];
  Autorizado = false;
  MontoSugerido!: number;
  public Monto = new FormControl(null, Validators.required);

  constructor(
    public dialogRef: MatDialogRef<ExaminarMontoContratoComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
    private  _contrato: ContratoService,
    private _conceptos: ConceptosContratoService,
  ) {
  }

  ngOnInit() {
    console.log(this.data.num_contrato);
    this._contrato.ConsultarContratos(this.data.num_contrato, '','', '', '', '', '',
      '', '', '').subscribe(resultado => {
      this.Contrato = resultado.Data['registros'][0];
      console.log(this.Contrato);
    });
    this._conceptos.ConsultarConceptosContrato(this.data.num_contrato).subscribe(resultado => {
      this.Conceptos = resultado.Data['registros'];
      this.MontoContratado = resultado.Data['total'];
      console.log(this.Conceptos, this.MontoContratado);
    });
  }

  AutorizacionEdicion() {
    Swal.fire({
      title: '¿Estas Seguro?',
      text: 'Editaras el monto total del contrato',
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      confirmButtonText: 'Si, editar'
    }).then((result) => {
      if (result.value) {
        this.Autorizado = true;
        this.MontoSugerido = ((this.data.total_contratado + this.MontoContratado) - this.Contrato.MontoTotal) + this.Contrato.MontoTotal;
        this.Monto.setValue(this.MontoSugerido);
      }
    });
  }

  ActualizarMonto() {
    if (this.Monto.value !== null) {
      Swal.fire({
        title: 'Estas Seguro',
        text: 'Editaras el monto total del contrato',
        icon: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#3085d6',
        cancelButtonColor: '#d33',
        confirmButtonText: 'Si, editar'
      }).then((result) => {
        if (result.value) {
          const Monto = this.Monto.value;
          this.Contrato.MontoTotal = Monto;
          this._contrato.ActualizarContrato(this.Contrato.Id, this.Contrato).subscribe(res => {
            Swal.fire(
              'Actualizado',
              res.Mensaje,
              'success'
            );
            this.dialogRef.close(true);
          });
        }
      });
    }
  }

  onNoClick(): void {
    this.dialogRef.close();
  }
}
