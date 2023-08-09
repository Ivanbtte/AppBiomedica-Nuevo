import {Component, Inject, OnInit, ViewEncapsulation} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialog, MatDialogRef} from '@angular/material/dialog';
import {MatTableDataSource} from '@angular/material/table';
import {ConceptosContratoService} from '../../../../core/servicios/ContratosServicios/conceptos-contrato.service';
import {RespuestaPeticion} from '../../../../core/models/estructuras_generales';
import Swal from 'sweetalert2';
import {AgregarUnContratoComponent} from '../agregar-un-contrato/agregar-un-contrato.component';
import {ExaminarMontoContratoComponent} from '../examinar-monto-contrato/examinar-monto-contrato.component';
import {ContratoPuestoService} from '../../../../core/servicios/ContratosServicios/contrato-puesto.service';
import jwt_decode from 'jwt-decode';
import {SubTipoService} from '../../../../core/servicios/ContratosServicios/sub-tipo.service';
import {SubTipo} from '../../../../core/models/sub_tipo.interface';
import { saveAs } from 'file-saver';

export interface DialogData {
  animal: string;
}

@Component({
  selector: 'app-enviar-archivo-concept',
  templateUrl: './carga-masiva-conceptos.component.html',
  styleUrls: ['./carga-masiva-conceptos.component.css'],
  encapsulation: ViewEncapsulation.None,
})
export class EnviarArchivoConceptComponent implements OnInit {
  files: File[] = [];
  Errores: any[]=[];
  displayedColumns: string[] = ['NumContrato', 'Linea', 'Fila', 'Error', 'Tipo'];
  dataSource: any;
  public validado = false;
  public color = '';
  tipo_contrato: any[]=[];
  tipo_seleccionado = '';
  sub_tipo_seleccionado = '';
  nombre_plantilla = '';
  lista_sub_tipos: SubTipo[] = [];

  constructor(
    private _conceptos: ConceptosContratoService,
    public dialogRef: MatDialogRef<EnviarArchivoConceptComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
    public dialog: MatDialog,
    private contrato_puesto: ContratoPuestoService,
    private _sub_tipo: SubTipoService,
  ) {
  }

  ngOnInit() {
    const decode:any = jwt_decode(<string>localStorage.getItem('token'));
    this.contrato_puesto.Consultarcontrato_puesto(decode.PuestoId).subscribe(res => {
      console.log(res.Data)
      this.tipo_contrato = res.Data.map((ite:any) => {
        return {
          id: ite.TipoContrato.Id,
          tipo: ite.TipoContrato.Tipo
        };
      });
    });
    this.dataSource = new MatTableDataSource();

  }

  /* ******  Funcion para guardar el valor seleccionado del tipo de contrato  ******   */
  tipoSeleccionado(value:any) {
    this.tipo_seleccionado = '';
    this.nombre_plantilla = '';
    console.log(value);
    this.consultarSubTipo(value);
  }

  /* ******  Funcion para guardar el valor seleccionado del subipo de contrato  ******   */
  subTipoSeleccionado(value:any) {
    console.log(value);
    this.sub_tipo_seleccionado = value.Id;
    this.nombre_plantilla = value.SubTipo;
  }

  /* ******  Funcion para consultar los subtipos de contratos  ******   */
  consultarSubTipo(tipo_contrato:any) {
    this._sub_tipo.ConsultaSubTipo(tipo_contrato.id).subscribe(res => {
      this.lista_sub_tipos = res.Data;
      if (this.lista_sub_tipos.length === 0) {
        this.lista_sub_tipos = [];
        this.tipo_seleccionado = tipo_contrato.id;
        this.nombre_plantilla = tipo_contrato.tipo;
      }
    }, err => {
      if (err.error !== '' || err.error !== undefined) {
        const erores = <RespuestaPeticion>err.error;
        console.log(erores);
      }
    });
  }

  onRemove(event:any) {
    console.log(event);
    this.files.splice(this.files.indexOf(event), 1);
    this.validado = false;
    this.Errores = [];
    this.dataSource.data = this.Errores;
    this.color = '#0089e1';
  }

  onSelect(event:any) {
    if (this.files.length < 1) {
      console.log(...event.addedFiles);
      this.files.push(...event.addedFiles);
    }
  }

  onNoClick(): void {
    this.dialogRef.close();
  }

// ********************************************************************************************************
  /* ***** Funcion para validar que los campos del archivo esten correctos  *****   */
  ValidarArchivo() {
    const tipo = this.tipo_seleccionado !== null ? this.tipo_seleccionado : '';
    const sub_tipo = this.sub_tipo_seleccionado !== null ? this.sub_tipo_seleccionado : '';
    this.validado = false;
    this.Errores = [];
    this.dataSource.data = this.Errores;
    const formData = new FormData();
    formData.append('file', this.files[0]);
    this._conceptos.ValidarConceptosFile(formData, tipo, sub_tipo).subscribe(resultado => {
      console.log(resultado);
      Swal.fire({
        position: 'center',
        title: resultado.Mensaje,
        icon: 'success',
        showConfirmButton: true,
      });
      this.validado = true;
      this.color = '#18903f';
    }, err => {
      if (err.error !== '' || err.error !== undefined) {
        const erores = <RespuestaPeticion>err.error;
        this.Errores = erores.Data;
        this.dataSource.data = this.Errores;
        console.log(this.Errores);
        this.color = '#d4782c';
        Swal.fire({
          position: 'center',
          title: erores.Mensaje,
          icon: 'error',
          showConfirmButton: true,
        });
      }
    });
  }

  // ********************************************************************************************************
  /* ***** Funcion para validar que los campos del archivo esten correctos  *****   */
  AgregarConceptos() {
    const formData = new FormData();
    formData.append('file', this.files[0]);
    const tipo = this.tipo_seleccionado !== null ? this.tipo_seleccionado : '';
    const sub_tipo = this.sub_tipo_seleccionado !== null ? this.sub_tipo_seleccionado : '';
    Swal.fire({
      title: 'Estas seguro de agregar estos datos',
      showCancelButton: true,
      icon: 'warning',
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      confirmButtonText: 'Si, agregar'
    }).then((result) => {
      if (result.value) {
        this._conceptos.AgregarConceptosArchivo(formData, tipo, sub_tipo).subscribe(resultado => {
          Swal.fire(
            'Agregados!',
            'Los Datos se han guardado exitosamente',
            'success'
          );
          this.dialogRef.close(true);
        }, err => {
          if (err.error !== '' || err.error !== undefined) {
            const erores = <RespuestaPeticion>err.error;
            console.log(this.Errores);
            this.color = '#d4782c';
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
  /* *****   descargar archivo *****   */
  descargar() {
    console.log('Tipo', this.tipo_seleccionado, '-----nombre', this.nombre_plantilla);
    const nom_temp = this.nombre_plantilla.toLocaleLowerCase().split(' ');
    const nomb = nom_temp.join('_');
    console.log(nomb);
    this._conceptos.DescargarArchivoPlantilla(nomb).subscribe(res => {
      console.log(res);
      if (File) {
        console.log(File);
        saveAs(res, 'plantilla_' + nomb + '.xlsx');
      }
    }, err => {
      if (err.error !== '' || err.error !== undefined) {
        const erores = <RespuestaPeticion>err.error;
        console.log(this.Errores);
        Swal.fire({
          position: 'center',
          title: 'No se encontro el archivo',
          icon: 'error',
          showConfirmButton: true,
        });
      }
    });
  }

  // ********************************************************************************************************
  /* *****  Funcion para cambiar el mensaje del boton de acciones  *****   */
  MensajeBoton(codigo: number): string {
    if (codigo === 1) {
      return 'Agregar';
    }
    if (codigo === 2) {
      return 'Consultar';
    }
    return 'Ver';
  }

  // ********************************************************************************************************
  /* *****  Funcion para determinar que accion realizar dependiendo del error  *****   */
  AbrirAccion(row: any) {
    if (row.Tipo === 2) {
      const dialogRef = this.dialog.open(ExaminarMontoContratoComponent, {
        height: 'auto',
        width: '70%',
        disableClose: true,
        data: {
          num_contrato: row.NumContrato,
          total_contratado: row.Suma
        }
      });
      dialogRef.afterClosed().subscribe(result => {
        if (result === true) {
          this.ValidarArchivo();
        }
      });
    }

    if (row.Tipo === 1) {
      const dialogRef = this.dialog.open(AgregarUnContratoComponent, {
        height: 'auto',
        width: '90%',
        disableClose: true,
        data: {}
      });
      dialogRef.afterClosed().subscribe(result => {
        if (result === true) {
          this.ValidarArchivo();
        }
      });
    }
    if (row.Tipo === 0) {
      Swal.fire({
        position: 'center',
        title: 'Debes editar tu archivo  de carga',
        text: row.Error,
        icon: 'error',
        showConfirmButton: true,
      });
    }
  }
}
