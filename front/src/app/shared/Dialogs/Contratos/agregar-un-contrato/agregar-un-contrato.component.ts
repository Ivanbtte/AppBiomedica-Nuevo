import {Component, Inject, OnInit} from '@angular/core';
import {MatDatepickerInputEvent} from '@angular/material/datepicker';
import {MAT_DIALOG_DATA, MatDialog, MatDialogRef} from '@angular/material/dialog';
import {DateAdapter, MAT_DATE_FORMATS, MAT_DATE_LOCALE} from '@angular/material/core';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {SeleccionarProveedorRepresentanteComponent} from '../seleccionar-proveedor-representante/seleccionar-proveedor-representante.component';
import {MAT_MOMENT_DATE_FORMATS, MomentDateAdapter} from '@angular/material-moment-adapter';
import {Contrato, UnContrato} from '../../../../core/models/contrato.interface';
import Swal from 'sweetalert2';
import moment from 'moment';
import {RepresentanteLegalComponent} from '../representante-legal/representante-legal.component';
import jwt_decode from 'jwt-decode';
import {DelegacionesService} from '../../../../core/servicios/delegaciones.service';
import {ContratoService} from '../../../../core/servicios/ContratosServicios/contrato.service';
import {RepresentanteLegalService} from '../../../../core/servicios/ContratosServicios/representante-legal.service';
import {ContratoDelegacionService} from '../../../../core/servicios/ContratosServicios/contrato-delegacion.service';
import {ContratoPuestoService} from '../../../../core/servicios/ContratosServicios/contrato-puesto.service';

export interface DialogData {
  contrato: any;
  representante: any;
  proveedor: any;
  editar: boolean;
}

@Component({
  selector: 'app-agregar-un-contrato',
  templateUrl: './agregar-un-contrato.component.html',
  styleUrls: ['./agregar-un-contrato.component.css'],
  providers: [
    {provide: MAT_DATE_LOCALE, useValue: 'ja-JP'},
    {provide: DateAdapter, useClass: MomentDateAdapter, deps: [MAT_DATE_LOCALE]},
    {provide: MAT_DATE_FORMATS, useValue: MAT_MOMENT_DATE_FORMATS},
  ],
})
export class AgregarUnContratoComponent implements OnInit {
  public formulario_contrato!: FormGroup;
  representante: any;
  proveedor: any;
  fechaInico = '';
  fechaFin = '';
  fechaFallo = '';
  fechaActual = new Date();
  m = this.fechaActual.getUTCMonth();
  d = this.fechaActual.getDate();
  y = this.fechaActual.getFullYear();
  minDate = new Date(this.y - 10, this.m, this.d);
  maxDate = new Date(this.y + 20, 0, 1);
  minDate2: any;
  delegacionesArray: any[] = [];
  // public form_delegaciones = new FormControl();
  array_delegaciones_admin: any[] = [];
  contrato_nuevo!: UnContrato;
  /*variables para editar*/
  titulo = 'Agregar Contrato';
  contrato: any;
  fecha_inicio: any;
  fecha_fin: any;
  fecha_fallo: any;
  tipo_contrato: any[] = [];

  constructor(
    public dialogRef: MatDialogRef<AgregarUnContratoComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
    private fb: FormBuilder,
    public dialog: MatDialog,
    private adapter: DateAdapter<any>,
    private _delegaciones: DelegacionesService,
    private _contrato: ContratoService,
    private _delegacion_contrato: ContratoDelegacionService,
    private _representante: RepresentanteLegalService,
    private contrato_puesto: ContratoPuestoService,
  ) {
    this.adapter.setLocale('es-mx');
    this.contrato = this.data.contrato;
    if (this.data.editar) {
      this.titulo = 'Editar Contrato';
      this.proveedor = this.data.proveedor;
      this.consultarRepresentante();
    }
    this.proveedor = this.data.proveedor;
    this.representante = this.data.representante;
    console.log(data.contrato);
  }

  ngOnInit() {
    this.fecha_inicio = moment(this.data.contrato.InicioContrato, 'DD-MM-YYYY');
    this.fecha_fin = moment(this.data.contrato.FinContrato, 'DD-MM-YYYY');
    this.fecha_fallo = moment(this.data.contrato.FechaFallo, 'DD-MM-YYYY');
    this.fechaInico = this.data.contrato.InicioContrato;
    this.fechaFin = this.data.contrato.FinContrato;
    this.fechaFallo = this.data.contrato.FechaFallo;
    this.crear_formulario_contratos();
    this._delegaciones.ConsultarDelegaciones().subscribe(resultado => {
      this.delegacionesArray = resultado.Data;
      console.log(this.delegacionesArray);
    });
    this.consultarDelegaciones();
    this.consultarTipoContrato();
  }

  consultarTipoContrato() {
    const decode: any = jwt_decode(<string> localStorage.getItem('token'));
    this.contrato_puesto.Consultarcontrato_puesto(decode.PuestoId).subscribe(res => {
      this.tipo_contrato = res.Data.map((ite: any) => {
        return {
          Id: ite.TipoContrato.Id,
          TipoContrato: ite.TipoContrato.Tipo
        };
      });
      console.log(this.tipo_contrato);
    });
  }

  consultarDelegaciones() {
    this._delegacion_contrato.ConsultarDelegacionContrato(this.data.contrato.NumeroContrato).subscribe(res => {
      const seleccionados = res.Data.map((it: any) => it.Delegaciones);
      console.log(seleccionados);
      this.formulario_contrato.controls['form_delegaciones'].setValue(seleccionados);
      this.array_delegaciones_admin = seleccionados;
    });
  }

  compareDeleg(c1: { NombreDele: string }, c2: { NombreDele: string }) {
    return c1 && c2 && c1.NombreDele === c2.NombreDele;
  }

  onNoClick(): void {
    this.dialogRef.close();
  }

  consultarRepresentante() {
    this._representante.ConsultaRepresentante(this.contrato.NumeroContrato).subscribe(resul => {
      this.representante = resul.Data;
      console.log(this.representante);
      this.formulario_contrato.controls['NombreCompleto'].setValue(this.representante.NombreCompleto);
      this.formulario_contrato.controls['AliasEmpresa'].setValue(this.proveedor.AliasEmpresa);
    });
  }

  // Workaround for angular component issue #13870
  // tslint:disable-next-line:member-ordering
  disableAnimation = true;

  // tslint:disable-next-line:use-life-cycle-interface
  ngAfterViewInit(): void {
    // timeout required to avoid the dreaded 'ExpressionChangedAfterItHasBeenCheckedError'
    setTimeout(() => this.disableAnimation = false);
  }

  crear_formulario_contratos() {
    this.formulario_contrato = this.fb.group({
      NumeroContrato: [this.contrato.NumeroContrato, Validators.compose([Validators.required])],
      InicioContrato: [this.fecha_inicio, Validators.compose([Validators.required])],
      FinContrato: [this.fecha_fin, Validators.compose([Validators.required])],
      ProcedContratacion: [this.contrato.ProcedContratacion, Validators.compose([Validators.required])],
      TipoProcedContratacion: [this.contrato.TipoProcedContratacion, Validators.compose([Validators.required])],
      FechaFallo: [this.fecha_fallo, Validators.compose([Validators.required])],
      MontoTotal: [this.contrato.MontoTotal, Validators.compose([Validators.required])],
      TipoContrato: [this.contrato.TipoContrato, Validators.compose([Validators.required])],
      NombreCompleto: [this.representante.NombreCompleto, Validators.compose([Validators.required])],
      AliasEmpresa: [this.proveedor.AliasEmpresa, Validators.compose([Validators.required])],
      form_delegaciones: [null, Validators.compose([])],
    });
  }

  compare(c1: { TipoContrato: string }, c2: { TipoContrato: string }) {
    console.log(c1, c2);
    return c1 && c2 && c1.TipoContrato === c2.TipoContrato;
  }

  DialogProveedor(): void {
    const dialogRef = this.dialog.open(SeleccionarProveedorRepresentanteComponent, {
      height: 'auto',
      width: '89%',
      disableClose: true,
      data: {
        proveedor: ''
      }
    });
    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.proveedor = result;
        console.log(this.proveedor);
        this.formulario_contrato.controls['AliasEmpresa'].setValue(this.proveedor.AliasEmpresa);
        if (this.data.editar) {
          this.formulario_contrato.controls['AliasEmpresa'].markAsDirty();
        }
      }
    });
  }

  DialogRepresentante(): void {
    const dialogRef = this.dialog.open(RepresentanteLegalComponent, {
      height: 'auto',
      width: '65%',
      disableClose: true,
      data: {
        representante: this.representante,
        num_contrato: this.data.contrato.NumeroContrato,
        editar: this.data.editar,
      }
    });
    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        if (this.data.editar) {
          this.consultarRepresentante();
        } else {
          this.representante = result;
          console.log(this.representante);
          this.formulario_contrato.controls['NombreCompleto'].setValue(this.representante.NombreCompleto);
          if (this.data.editar) {
            this.formulario_contrato.controls['NombreCompleto'].markAsDirty();
          }
        }
      }
    });
  }

  valorFecha(event: MatDatepickerInputEvent<Date>) {
    console.log(event.value);
    // this.fechaInico = event.targetElement['value'];
    // const dia = event.value['_i'].date;
    // const mes = event.value['_i'].month;
    // const anio = event.value['_i'].year;
    // this.minDate2 = new Date(anio, mes, dia + 1);
  }

  valorFechaLimite(event: MatDatepickerInputEvent<Date>) {
    console.log(event.value);
    // this.fechaFin = event.targetElement['value'];
    // const dia = event.value['_i'].date;
    // const mes = event.value['_i'].month;
    // const anio = event.value['_i'].year;
    // this.maxDate = new Date(anio, mes, dia - 1);
  }

  valorFechaFallo(event: MatDatepickerInputEvent<Date>) {
    // this.fechaFallo = event.targetElement['value'];
    console.log(event.value);

  }

  // ********************************************************************************************************
  /* *****  opcines Seleccionadas  *****   */
  seleccionados() {
    this.array_delegaciones_admin = this.formulario_contrato.controls['form_delegaciones'].value;
  }

  // ********************************************************************************************************
  /* *****  Funcion para agregar el contrato a la base de datos  *****   */
  agregar_contrato() {
    const clv_deles = this.array_delegaciones_admin.map(function(dele) {
      return dele.ClvDele;
    });
    const contrato_previo = <Contrato> this.formulario_contrato.getRawValue();
    contrato_previo.InicioContrato = this.fechaInico;
    contrato_previo.FechaFallo = this.fechaFallo;
    contrato_previo.FinContrato = this.fechaFin;
    contrato_previo.ProveedorNProvImss = this.proveedor.NProvImss;
    this.contrato_nuevo = {
      Contrato: contrato_previo,
      Delegaciones: clv_deles,
      Representante: this.representante
    };
    console.log(this.contrato_nuevo, contrato_previo.ProveedorNProvImss);
    Swal.fire({
      title: 'Confirmacion',
      text: 'Agregar Contrato' + contrato_previo.NumeroContrato,
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      confirmButtonText: 'Si Agregar'
    }).then((result) => {
      if (result.value) {
        this._contrato.AgregarUnContrato(this.contrato_nuevo).subscribe(resultado => {
          console.log(resultado);
          Swal.fire({
            position: 'center',
            title: 'Correcto',
            text: resultado.Mensaje,
            icon: 'success',
            showConfirmButton: true,
          });
          this.dialogRef.close(true);
        });
      }
    });
  }

  // ********************************************************************************************************
  /* *****  Funcion para editar un contrato  *****   */
  editarContrato() {
    const contrato_prev = <Contrato> this.formulario_contrato.getRawValue();
    contrato_prev.InicioContrato = this.fechaInico;
    contrato_prev.FechaFallo = this.fechaFallo;
    contrato_prev.FinContrato = this.fechaFin;
    contrato_prev.ProveedorNProvImss = this.proveedor.NProvImss;
    console.log(contrato_prev);
    Swal.fire({
      title: 'Actualizacion',
      text: '¿Estas seguro de editar la informacion?',
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#a2a2a2',
      confirmButtonText: 'Si, editar',
      cancelButtonText: 'Cancelar',
      reverseButtons: true
    }).then((result) => {
      if (result.value) {
        this._contrato.ActualizarContrato(this.contrato.Id, contrato_prev).subscribe(res => {
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
