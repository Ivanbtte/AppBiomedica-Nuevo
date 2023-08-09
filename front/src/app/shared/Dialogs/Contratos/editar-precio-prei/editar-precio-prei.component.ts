import {Component, Inject, OnInit} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {FormControl, Validators} from '@angular/forms';
import {CurrencyPipe} from '@angular/common';
import {ConceptoContrato} from '../../../../core/models/concepto_contrato.interface';
import {RespuestaPeticion} from '../../../../core/models/estructuras_generales';
import Swal from 'sweetalert2';
import {ConceptosContratoService} from '../../../../core/servicios/ContratosServicios/conceptos-contrato.service';

export interface DialogData {
  concepto: any;
}

@Component({
  selector: 'app-editar-precio-prei',
  templateUrl: './editar-precio-prei.component.html',
  styleUrls: ['./editar-precio-prei.component.css'],
  providers: [CurrencyPipe],
})
export class EditarPrecioPreiComponent implements OnInit {
  claveCuadro='';
  verMas = false;
  idPrei: ConceptoContrato;
  editar = false;
  control = new FormControl(null, Validators.compose([Validators.pattern('^([0-9])*\\.?[0-9]*$'), Validators.required]));

  constructor(
    public dialogRef: MatDialogRef<EditarPrecioPreiComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
    private curre: CurrencyPipe,
    private _conceptos: ConceptosContratoService
  ) {
    this.idPrei = data.concepto;
  }

  ngOnInit() {
    this.claveCuadro = this.data.concepto.Prei.Grupo + '.' + this.data.concepto.Prei.Generico + '.' + this.data.concepto.Prei.Especifico + '.' +
      this.data.concepto.Prei.Diferenciador + '.' + this.data.concepto.Prei.Variable;
  }

  onNoClick(): void {
    this.dialogRef.close();
  }

  regresaDescripcion(valor: string): string {
    const descripcion = valor.split('.', 1);
    return descripcion.toString();
  }

  actualizarPrecio() {
    this.idPrei.PrecioUniSnIva = parseFloat(this.control.value);
    this._conceptos.ActualizarConceptoContrato(this.idPrei.Id, this.idPrei, this.idPrei.ContratoNumeroContrato).subscribe(res => {
      Swal.fire({
        position: 'center',
        title: res.Mensaje,
        icon: 'success',
        showConfirmButton: true,
      });
      this.dialogRef.close(true);
    }, err => {
      if (err.error !== '' || err.error !== undefined) {
        this.idPrei.PrecioUniSnIva = this.data.concepto.PrecioUniSnIva;
        const erores = <RespuestaPeticion>err.error;
        Swal.fire({
          position: 'center',
          title: erores.Mensaje,
          text: erores.Data,
          icon: 'error',
          showConfirmButton: true,
        });
      }
    });
  }
}
