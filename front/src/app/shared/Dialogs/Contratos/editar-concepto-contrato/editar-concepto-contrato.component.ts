import {Component, Inject, OnInit} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialog, MatDialogRef} from '@angular/material/dialog';
import {DateAdapter, MAT_DATE_FORMATS, MAT_DATE_LOCALE,} from '@angular/material/core';
import {MatDatepickerInputEvent} from '@angular/material/datepicker';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {MAT_MOMENT_DATE_FORMATS, MomentDateAdapter} from '@angular/material-moment-adapter';
import moment from 'moment';
import Swal from 'sweetalert2';
import {RespuestaPeticion} from '../../../../core/models/estructuras_generales';
import {ConceptosContratoService} from '../../../../core/servicios/ContratosServicios/conceptos-contrato.service';

export interface DialogData {
  conceptoEditar: any;
}

@Component({
  selector: 'app-editar-concepto-contrato',
  templateUrl: './editar-concepto-contrato.component.html',
  styleUrls: ['./editar-concepto-contrato.component.css'],
  providers: [
    {provide: MAT_DATE_LOCALE, useValue: 'ja-JP'},
    {provide: DateAdapter, useClass: MomentDateAdapter, deps: [MAT_DATE_LOCALE]},
    {provide: MAT_DATE_FORMATS, useValue: MAT_MOMENT_DATE_FORMATS},
  ],
})
export class EditarConceptoContratoComponent implements OnInit {
  conceptoUpdate: any;
  public formulario_conceptos!: FormGroup;
  descripcion = '';
  fechaMaxima: any;
  fechaNueva = '';

  constructor(
    public dialogRef: MatDialogRef<EditarConceptoContratoComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
    private fb: FormBuilder,
    private adapter: DateAdapter<any>,
    public dialog: MatDialog,
    private _conceptos: ConceptosContratoService,
  ) {
    this.adapter.setLocale('es-mx');
    this.conceptoUpdate = data.conceptoEditar;
    this.descripcion = data.conceptoEditar.Prei.Descripcion;
    this.fechaMaxima = moment(data.conceptoEditar.FechaMaxEntrega, 'DD-MM-YYYY');
    this.fechaNueva = data.conceptoEditar.FechaMaxEntrega;
  }

  ngOnInit() {
    console.log(this.data.conceptoEditar);
    this.crear_formulario_conceptos();
  }

  crear_formulario_conceptos() {
    this.formulario_conceptos = this.fb.group({
      PrecioUniSnIva: [this.conceptoUpdate.PrecioUniSnIva, Validators.compose([Validators.required])],
      CantidadConcepto: [this.conceptoUpdate.CantidadConcepto, Validators.compose([Validators.required,
        Validators.pattern('^([0-9])*$')])],
      Marca: [this.conceptoUpdate.Marca, Validators.compose([Validators.required])],
      Modelo: [this.conceptoUpdate.Modelo, Validators.compose([Validators.required])],
      FechaMaxEntrega: [this.fechaMaxima, Validators.compose([Validators.required])],
      GarantiaBienes: [this.conceptoUpdate.GarantiaBienes, Validators.compose([Validators.required,
        Validators.pattern('^([0-9])*$')])],
      PreiIdArticulo: [this.conceptoUpdate.PreiIdArticulo, Validators.compose([Validators.required])],
    });
  }

  onNoClick(): void {
    this.dialogRef.close();
  }

  actualizar() {
    const datos: any = this.formulario_conceptos.getRawValue();
    datos.FechaMaxEntrega = this.fechaNueva;
    console.log(datos);
    Swal.fire({
      title: 'Actualizacion',
      text: '¿Estas seguro que deseas actualizar?',
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      confirmButtonText: 'Si, Actualizar',
      cancelButtonText: 'Cancelar'
    }).then((result) => {
      if (result.value) {
        this._conceptos.ActualizarConceptoContrato(this.conceptoUpdate.Id, datos,
          this.conceptoUpdate.ContratoNumeroContrato).subscribe(res => {
          Swal.fire(
            'Correcto',
            res.Mensaje,
            'success'
          );
          this.dialogRef.close(true);
        }, err => {
          if (err.error !== '' || err.error !== undefined) {
            const erores = <RespuestaPeticion> err.error;
            console.log(erores);
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
    });
  }

  valorFecha(event: MatDatepickerInputEvent<Date>) {
    // console.log(event.targetElement['value']);
    // this.fechaNueva = event.targetElement['value'];
    console.log(this.fechaMaxima, event);
  }
}
