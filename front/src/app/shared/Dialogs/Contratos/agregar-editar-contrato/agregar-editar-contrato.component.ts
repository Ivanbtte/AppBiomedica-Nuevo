import {Component, Inject, OnInit, ViewChild} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialog, MatDialogRef} from '@angular/material/dialog';
import {FormBuilder, FormControl, FormGroup, Validators} from '@angular/forms';
import Swal from 'sweetalert2';
import {Contrato, ContratoNuevo, Fianza, RepresentanteLegal} from '../../../../core/models/contrato.interface';
import {ProveedoresService} from '../../../../core/servicios/ProveedoresServicios/proveedores.service';
import {ContactoProveedor, ProveedoresInterface} from '../../../../core/models/proveedores.interface';
import {ContactoProveedoresService} from '../../../../core/servicios/ProveedoresServicios/contacto-proveedores.service';
import {DateAdapter, MAT_DATE_FORMATS, MAT_DATE_LOCALE} from '@angular/material/core';
import {MAT_MOMENT_DATE_FORMATS, MomentDateAdapter} from '@angular/material-moment-adapter';
import {MatDatepickerInputEvent} from '@angular/material/datepicker';
import {CrudPersonasComponent} from '../../crud-personas/crud-personas.component';
import {PenasConvencionales, PenasDeductivas} from '../../../../core/models/penas_deductivas.interface';
import {SeleccionarProveedorRepresentanteComponent} from '../seleccionar-proveedor-representante/seleccionar-proveedor-representante.component';
import {AdminAux} from '../../../../core/models/admin_aux.interface';
import {ContratoService} from '../../../../core/servicios/ContratosServicios/contrato.service';
import {RespuestaPeticion} from '../../../../core/models/estructuras_generales';
import {DelegacionesService} from '../../../../core/servicios/delegaciones.service';

export interface DialogData {
  editar: boolean;
  contrato: Contrato;
}

@Component({
  selector: 'app-agregar-editar-contrato',
  templateUrl: './agregar-editar-contrato.component.html',
  styleUrls: ['./agregar-editar-contrato.component.css'],
  providers: [
    {provide: MAT_DATE_LOCALE, useValue: 'ja-JP'},
    {provide: DateAdapter, useClass: MomentDateAdapter, deps: [MAT_DATE_LOCALE]},
    {provide: MAT_DATE_FORMATS, useValue: MAT_MOMENT_DATE_FORMATS},
  ],
})
export class AgregarEditarContratoComponent implements OnInit {
  public form_delegaciones = new FormControl();
  public formulario_contrato!: FormGroup;
  public formulario_penas_conv!: FormGroup;
  public formulario_penas_ded!: FormGroup;
  public formulario_fianza!: FormGroup;
  public formulario_admin_aux!: FormGroup;
  fechaInico = '';
  fechaFin = '';
  fechaFallo = '';
  titulo = '';
  datos!: Contrato;
  representante!: ContactoProveedor;
  repre_legal!: RepresentanteLegal;
  array_penas_deductivas: PenasDeductivas[] = [];
  array_admin: AdminAux[] = [];
  delegacionesArray: any[] = [];
  proveedoresArray: ProveedoresInterface[] = [];
  lastFilter = '';
  date = new FormControl(new Date());
  fechaActual = new Date();
  m = this.fechaActual.getUTCMonth();
  d = this.fechaActual.getDate();
  y = this.fechaActual.getFullYear();
  minDate = new Date(this.y - 10, this.m, this.d);
  maxDate = new Date(this.y + 20, 0, 1);
  minDate2: any;
  nombre_compelto = '';
  array_delegaciones_admin: any[] = [];
  contrato_nuevo!: ContratoNuevo;
  index_tab = 0;
  @ViewChild('tabGroup') tabGroup! : any;

  constructor(
    public dialogRef: MatDialogRef<AgregarEditarContratoComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
    private fb: FormBuilder,
    private _delegaciones: DelegacionesService,
    private _proveedores: ProveedoresService,
    private _contacto_proveedor: ContactoProveedoresService,
    public dialog: MatDialog,
    private adapter: DateAdapter<any>,
    private _contrato: ContratoService,
  ) {
    this.adapter.setLocale('es-mx');
  }

  ngOnInit() {
    console.log(this.data.editar);
    this.datos = this.data.contrato;
    this.crear_formulario_contratos();
    this.crear_formulario_penas_conv();
    this.crear_formulario_penas_ded();
    this.crear_formulario_fianza();
    this.crear_formulario_admin_aux();
    // this.crear_formulario_concepto_contrato();
    if (this.data.editar) {
      this.titulo = 'Editar un Contrato';
    } else {
      this.titulo = 'Agregar un nuevo Contrato';
    }
    this._delegaciones.ConsultarDelegaciones().subscribe(resultado => {
      this.delegacionesArray = resultado.Data;
    });
  }

  crear_formulario_contratos() {
    this.formulario_contrato = this.fb.group({
      NumeroContrato: [null, Validators.compose([Validators.required])],
      InicioContrato: [null, Validators.compose([Validators.required])],
      FinContrato: [null, Validators.compose([Validators.required])],
      ProcedContratacion: [null, Validators.compose([Validators.required])],
      TipoProcedContratacion: [null, Validators.compose([Validators.required])],
      FechaFallo: [null, Validators.compose([Validators.required])],
      MontoTotal: [null, Validators.compose([Validators.required])],
      TipoContrato: [null, Validators.compose([Validators.required])],
      ContactoProveedorId: [null, Validators.compose([Validators.required])],
      ProveedorNProvImss: [null, Validators.compose([Validators.required])],
    });
  }

  crear_formulario_penas_conv() {
    this.formulario_penas_conv = this.fb.group({
      Descripcion: ['', Validators.compose([Validators.required])],
      Porcentaje: ['', Validators.compose([Validators.required])],
    });
  }

  crear_formulario_penas_ded() {
    this.formulario_penas_ded = this.fb.group({
      ConceptoObligacion: ['', Validators.compose([Validators.required])],
      NivelServicio: ['', Validators.compose([Validators.required])],
      UnidadMedida: ['', Validators.compose([Validators.required])],
      Deduccion: ['', Validators.compose([Validators.required])],
      DescripcionDeduccion: ['', Validators.compose([Validators.required])],
      LimiteIncumplimiento: ['', Validators.compose([Validators.required])],
    });
  }

  crear_formulario_fianza() {
    this.formulario_fianza = this.fb.group({
      Afianzadora: ['', Validators.compose([Validators.required])],
      NumPoliza: ['', Validators.compose([Validators.required])],
      Tipo: ['', Validators.compose([Validators.required])],
      MontoFianza: ['', Validators.compose([Validators.required])],
    });
  }

  crear_formulario_admin_aux() {
    this.formulario_admin_aux = this.fb.group({
      DelegacionesClvDele: ['', Validators.compose([Validators.required])],
      Cargo: ['', Validators.compose([Validators.required])],
      Responsabilidad: ['', Validators.compose([Validators.required])],
      PersonaId: ['', Validators.compose([Validators.required])],
    });
  }

  cerrarDialog() {
    this.dialogRef.close();
  }

  // Workaround for angular component issue #13870
  disableAnimation = true;

  ngAfterViewInit(): void {
    // timeout required to avoid the dreaded 'ExpressionChangedAfterItHasBeenCheckedError'
    setTimeout(() => this.disableAnimation = false);
  }

  openDialogRepresentante(): void {
    const dialogRef = this.dialog.open(SeleccionarProveedorRepresentanteComponent, {
      height: 'auto',
      width: '88%',
      disableClose: true,
      data: {
        dato: ''
      }
    });
    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.representante = result;
        this.formulario_contrato.controls['ContactoProveedorId'].setValue(this.representante.Id);
        this.formulario_contrato.controls['ProveedorNProvImss'].setValue(this.representante.ProveedorNProvImss);
        console.log(this.representante);
      }
    });
  }

  openDialogPersona(): void {
    const dialogRef = this.dialog.open(CrudPersonasComponent, {
      autoFocus: true,
      height: 'auto',
      width: '90%',
      disableClose: true,
      data: {}
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        console.log(result);
        this.nombre_compelto = result.Nombre + ' ' + result.ApellidoPat + ' ' + result.ApellidoMat;
        this.formulario_admin_aux.controls['PersonaId'].setValue(result);
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
    // this.fechaFin = event.targetElement['value'];
    // const dia = event.value['_i'].date;
    // const mes = event.value['_i'].month;
    // const anio = event.value['_i'].year;
    // this.maxDate = new Date(anio, mes, dia - 1);
    console.log(event.value);
  }

  valorFechaFallo(event: MatDatepickerInputEvent<Date>) {
    // this.fechaFallo = event.targetElement['value'];
    console.log(event.value);
  }

  // ********************************************************************************************************
  /* *****  Funcion que inserta una nueva pena deductiva a un arreglo  *****   */
  agregar_pena_deductiva() {
    const datos = this.formulario_penas_ded.getRawValue();
    console.log(datos);
    this.array_penas_deductivas.push(datos);
    console.log(this.array_penas_deductivas);
    this.formulario_penas_ded.reset();
  }

  // ********************************************************************************************************
  /* *****  Funcion que elimina una pena de la lista   *****   */
  eliminar_pena_deductiva(item: PenasDeductivas) {
    const i = this.array_penas_deductivas.indexOf(item);
    console.log(i);
    this.array_penas_deductivas.splice(i, 1);
  }

  // ********************************************************************************************************
  /* *****  Funcion que inserta un administrador o un auxiliar al contrato  *****   */
  agregar_admin_aux() {
    const datos = <AdminAux> this.formulario_admin_aux.getRawValue();
    const delegaciones = this.formulario_admin_aux.controls['DelegacionesClvDele'].value;
    const personas = this.formulario_admin_aux.controls['PersonaId'].value;
    datos.DelegacionesClvDele = delegaciones.ClvDele;
    datos.Delegaciones = delegaciones;
    datos.PersonaId = personas.Id;
    datos.Persona = personas;
    console.log(datos);
    this.array_admin.push(datos);
    console.log(this.array_admin);
    this.formulario_admin_aux.reset();
    this.nombre_compelto = '';
  }

  // ********************************************************************************************************
  /* *****  Funcion que elimina una asignacion de la lista   *****   */
  eliminar_admin(item: AdminAux) {
    const i = this.array_admin.indexOf(item);
    console.log(i);
    this.array_admin.splice(i, 1);
  }

  // ********************************************************************************************************
  /* *****  opcines Seleccionadas  *****   */
  seleccionados() {
    this.array_delegaciones_admin = this.form_delegaciones.value;
    console.log(this.array_delegaciones_admin);
  }

// ********************************************************************************************************
  /* *****  Funcion para agregar el contrato a la base de datos  *****   */
  agregar_contrato() {
    const clv_deles = this.array_delegaciones_admin.map(function(dele) {
      return dele.ClvDele;
    });
    const contrato_previo = <Contrato> this.formulario_contrato.getRawValue();
    const pena_convencional = <PenasConvencionales> this.formulario_penas_conv.getRawValue();
    const fianza = <Fianza> this.formulario_fianza.getRawValue();
    contrato_previo.InicioContrato = this.fechaInico;
    contrato_previo.FechaFallo = this.fechaFallo;
    contrato_previo.FinContrato = this.fechaFin;
    this.repre_legal = {
      Id: '',
      ContratoId: '',
      ContactoProveedorId: contrato_previo.ContactoProveedorId
    };
    this.contrato_nuevo = {
      Contrato: contrato_previo,
      Delegaciones: clv_deles,
      AdminAuxContrato: this.array_admin,
      PenasConvencionales: pena_convencional,
      RepresentanteLegal: this.repre_legal,
      PenasDeductivas: this.array_penas_deductivas,
      Fianza: fianza
    };
    console.log(this.contrato_nuevo);
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
        this._contrato.AgregarContrato(this.contrato_nuevo).subscribe(resultado => {
          console.log(resultado);
          Swal.fire({
            position: 'center',
            title: 'Correcto',
            text: resultado.Mensaje,
            icon: 'success',
            showConfirmButton: true,
          });
          this.dialogRef.close(true);
        }, err => {
          if (err.error !== '' || err.error !== undefined) {
            const erores = <RespuestaPeticion> err.error;
            Swal.fire({
              position: 'center',
              title: erores.Mensaje,
              icon: 'error',
              showConfirmButton: true,
            });
          }
        });
      }
    });
  }

  // ********************************************************************************************************
  /* *****  Funcion para avanzar al tab siguiente  *****   */
  siguiente() {
    this.index_tab = this.tabGroup.selectedIndex + 1;
  }

  verificarArrayAdmin(): boolean {
    return this.array_admin.length > 0;
  }

  verificarArrayPenas(): boolean {
    return this.array_penas_deductivas.length > 0;
  }
}

