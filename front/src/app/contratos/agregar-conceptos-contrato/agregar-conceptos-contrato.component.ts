import {Component, OnInit} from '@angular/core';
import {RespuestaPeticion} from '../../core/models/estructuras_generales';
import Swal from 'sweetalert2';
import {ConceptoContrato, ConceptoContratoTemp, ConceptosNuevos} from '../../core/models/concepto_contrato.interface';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
// import {LocalDataSource} from 'ng2-smart-table';
import {ActivatedRoute, Router} from '@angular/router';
import {ContratoService} from '../../core/servicios/ContratosServicios/contrato.service';
import {MatDatepickerInputEvent} from '@angular/material/datepicker';
import {MatDialog} from '@angular/material/dialog';
import {MatTableDataSource} from '@angular/material/table';
import {DateAdapter, MAT_DATE_FORMATS, MAT_DATE_LOCALE} from '@angular/material/core';
import {ConceptosContratoService} from '../../core/servicios/ContratosServicios/conceptos-contrato.service';
import {AbrirCatalogoSaiService} from '../../core/servicios/abrir-catalogo-sai.service';
import {Contrato} from '../../core/models/contrato.interface';
import {ConsultarClavePreiComponent} from '../../shared/Dialogs/consultar-clave-prei/consultar-clave-prei.component';
import {MAT_MOMENT_DATE_FORMATS, MomentDateAdapter} from '@angular/material-moment-adapter';

@Component({
  selector: 'app-agregar-conceptos-contrato',
  templateUrl: './agregar-conceptos-contrato.component.html',
  styleUrls: ['./agregar-conceptos-contrato.component.css'],
  providers: [
    {provide: MAT_DATE_LOCALE, useValue: 'ja-JP'},
    {provide: DateAdapter, useClass: MomentDateAdapter, deps: [MAT_DATE_LOCALE]},
    {provide: MAT_DATE_FORMATS, useValue: MAT_MOMENT_DATE_FORMATS},
  ],
})
export class AgregarConceptosContratoComponent implements OnInit {
  conceptos_contrato: ConceptoContratoTemp[] = [];
  conceptoFinal!: ConceptosNuevos;
  conceptosAgregados: ConceptoContrato[] = [];
  fechaMaxima = '';
  clave_seleccionada: any;
  contrato: Contrato[] = [];
  descripcion = '';
  public formulario_conceptos!: FormGroup;
  datosArray: Array<any>[] = [];
  contrato_id: any;
  dataSource: any;
  estadoDescripcion = true;
  displayedColumns: string[] = ['id_prei', 'clave_cuadro', 'descripcion_prei', 'marca', 'modelo', 'fecha_max', 'precio'];
  configuraciones = {
    columns: {
      PreiIdArticulo: {
        title: 'Id PREI',
        filter: true,
        width: '10%'
      },
      Descripcion: {
        title: 'Descripcion',
        filter: true,
        width: '40%'
      },
      Marca: {
        title: 'Marca',
        filter: true,
        width: '20%'
      },
      Modelo: {
        title: 'Modelo',
        filter: true,
        width: '20%'
      },
      CantidadConcepto: {
        title: 'Cantidad',
        filter: true,
        width: '10%'
      },
    },
    mode: 'external',    // Use external functions to Edit/Delete, etc.
    actions: {
      add: false,
      edit: false,
      delete: false,
      custom: [
        {
          name: 'delete',
          title: '<i class="material-icons text-danger">delete</i>'
        },
        {
          name: 'edit',
          title: '<i class="material-icons text-info">edit</i>'
        },
      ]
    },
    noDataMessage: 'Sin Equipos por incluir',

  };
  // source: LocalDataSource;
  MontoContratado!: number;
  MontoRestante!: number;

  constructor(
    private _route: ActivatedRoute,
    private _contrato: ContratoService,
    public dialog: MatDialog,
    private fb: FormBuilder,
    private adapter: DateAdapter<any>,
    private _conceptos: ConceptosContratoService,
    private _activar_btn_cancelar: AbrirCatalogoSaiService,
    private router: Router,
  ) {
    this.adapter.setLocale('es-mx');
    // this.source = new LocalDataSource(this.datosArray);
  }

  ngOnInit() {
    this.contrato_id = this._route.snapshot.paramMap.get('contrato_id');
    this.crear_formulario_conceptos();
    this.consultarContrato();
    this._activar_btn_cancelar.setEstadoBtn(true);
    this._activar_btn_cancelar.setRuta('/contratos/detalle/' + this.contrato_id);
  }

  consultarContrato() {
    this._contrato.ConsultarContratos(this.contrato_id, '', '', '', '', '', '', '',
      '', '').subscribe(resultado => {
      console.log(resultado.Data);
      this.contrato = resultado.Data['registros'];
      this.consultarConceptos();
    });
  }

  consultarConceptos() {
    const id_contrato = this.contrato[0].NumeroContrato;
    console.log(id_contrato);
    this._conceptos.ConsultarConceptosContrato(id_contrato).subscribe(resultado => {
      console.log(resultado.Data);
      this.conceptosAgregados = resultado.Data['registros'];
      this.dataSource = new MatTableDataSource();
      this.dataSource.data = resultado.Data['registros'];
      this.MontoContratado = resultado.Data['total'];
      this.MontoRestante = this.contrato[0].MontoTotal - this.MontoContratado;
    });
  }

  crear_formulario_conceptos() {
    this.formulario_conceptos = this.fb.group({
      PrecioUniSnIva: [null, Validators.compose([Validators.required])],
      CantidadConcepto: [null, Validators.compose([Validators.required, Validators.pattern('^([0-9])*$')])],
      Marca: [null, Validators.compose([Validators.required])],
      Modelo: [null, Validators.compose([Validators.required])],
      FechaMaxEntrega: [null, Validators.compose([Validators.required])],
      GarantiaBienes: [null, Validators.compose([Validators.required, Validators.pattern('^([0-9])*$')])],
      PreiIdArticulo: [null, Validators.compose([Validators.required])],
    });
  }

  openDialogAgregarConcepto(): void {
    const id_prei = this.formulario_conceptos.controls['PreiIdArticulo'].value;
    if (id_prei !== null) {
      if (!this.verificarClaveRepetida(id_prei)) {
        const dialogRef = this.dialog.open(ConsultarClavePreiComponent, {
          height: 'auto',
          width: '70%',
          disableClose: true,
          data: {
            id_prei: id_prei
          }
        });
        dialogRef.afterClosed().subscribe((result: any) => {
          if (result) {
            this.descripcion = result.Descripcion;
            this.estadoDescripcion = false;
            this.clave_seleccionada = result;
          }
        });
      } else {
        Swal.fire({
          title: 'Oops...',
          text: 'Esta clave ya esta agregada',
          icon: 'error',
        });
      }
    }
  }

  verificarClaveRepetida(id_prei: string): boolean {
    const i = this.conceptosAgregados.findIndex(obj => obj.PreiIdArticulo === id_prei);
    if (i === -1) {
      const d = this.conceptos_contrato.findIndex(obj => obj.PreiIdArticulo === id_prei);
      return d !== -1;
    } else {
      return true;
    }
  }

  agregarConceptos() {
    const id_contrato = this.contrato[0].NumeroContrato;
    this.conceptoFinal = {
      ConceptoContrato: this.conceptos_contrato
    };
    console.log(this.conceptoFinal);
    Swal.fire({
      title: 'Estas Seguro',
      text: 'Deseas agregar estos equipos',
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      confirmButtonText: 'Si ,agregar'
    }).then((result) => {
      if (result.value) {
        this._conceptos.AgregarConceptoContrato(this.conceptoFinal, id_contrato).subscribe(resultado => {
          Swal.fire({
            position: 'center',
            title: 'Correcto',
            text: resultado.Mensaje,
            icon: 'success',
            showConfirmButton: true,
          });
          this.router.navigate(['/contratos/detalle', this.contrato_id]);
        }, err => {
          if (err.error !== '' || err.error !== undefined) {
            const erores = <RespuestaPeticion> err.error;
            console.log(erores.Data);
            Swal.fire({
              position: 'center',
              title: 'Id Prei: ' + erores.Data.PreiIdArticulo,
              text: this.regresaDescripcion(erores.Data.Prei.Descripcion),
              footer: erores.Mensaje,
              icon: 'error',
              showConfirmButton: true,
            });
          }
          console.log(err.error);
        });
      }
    });
  }

  eliminarConceptoArray(event: any) {
    const i = this.datosArray.indexOf(event.data);
    console.log(i);
    this.datosArray.splice(i, 1);
    // this.source.refresh();
  }

  onCustomAction(event: any) {
    switch (event.action) {
      case 'delete':
        this.eliminarConceptoArray(event);
        break;
      case 'edit':
        console.log('editar');
    }
  }

  valorFecha(event: MatDatepickerInputEvent<Date>) {
    // this.fechaMaxima = event.targetElement['value'];
    console.log(this.fechaMaxima);
  }

  AgregarItemArray() {
    const datos = this.formulario_conceptos.getRawValue();
    console.log('datos', datos);
    const verificar_Precio = datos.PrecioUniSnIva * datos.CantidadConcepto;
    console.log(verificar_Precio, '-', this.MontoRestante);
    if (verificar_Precio <= this.MontoRestante) {
      console.log('datos despues del if', datos);
      datos.Descripcion = this.regresaDescripcion(this.clave_seleccionada.Descripcion);
      datos.Prei = this.clave_seleccionada;
      datos.Fecha = this.fechaMaxima;
      this.datosArray.push(datos);
      // this.source.refresh();
      const concepto: ConceptoContratoTemp = {
        PrecioUniSnIva: datos.PrecioUniSnIva,
        CantidadConcepto: datos.CantidadConcepto,
        Marca: datos.Marca,
        Modelo: datos.Modelo,
        ObjetoContratacion: datos.Prei.CuadroBasico.Grupo,
        FechaMaxEntrega: datos.Fecha,
        GarantiaBienes: datos.GarantiaBienes,
        PreiIdArticulo: datos.PreiIdArticulo
      };
      console.log(concepto);
      this.MontoContratado = this.MontoContratado + (concepto.CantidadConcepto * concepto.PrecioUniSnIva);
      this.conceptos_contrato.push(concepto);
      this.formulario_conceptos.reset();
      this.descripcion = '';
      this.estadoDescripcion = true;
      this.MontoRestante = this.contrato[0].MontoTotal - this.MontoContratado;
    } else {
      Swal.fire({
        title: 'Oops...',
        text: 'No se puede agregar esta cantidad superas el limite de monto contratado',
        icon: 'error',
      });
    }
  }

  regresaDescripcion(valor: string): string {
    const descripcion = valor.split('.', 1);
    return descripcion.toString();
  }

  estaVacioArreglo(): boolean {
    return this.datosArray.length <= 0;
  }
}
