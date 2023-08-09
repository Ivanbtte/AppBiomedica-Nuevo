import {Component, Inject, OnInit, ViewEncapsulation} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialog, MatDialogRef} from '@angular/material/dialog';
import {MatTableDataSource} from '@angular/material/table';
import {saveAs} from 'file-saver';
import {AgregarUnContratoComponent} from '../agregar-un-contrato/agregar-un-contrato.component';
import Swal from 'sweetalert2';
import {RespuestaPeticion} from '../../../../core/models/estructuras_generales';
import {DistribucionConceptosComponent} from '../distribucion-conceptos/distribucion-conceptos.component';
import {ExaminarMontoContratoComponent} from '../examinar-monto-contrato/examinar-monto-contrato.component';
import {CurrencyPipe} from '@angular/common';
import {EditarPrecioPreiComponent} from '../editar-precio-prei/editar-precio-prei.component';
import {DistribucionConceptosService} from '../../../../core/servicios/ContratosServicios/distribucion-conceptos.service';

export interface DialogData {
  conceptos: boolean;
}

@Component({
  selector: 'app-carga-masiva-distribucion',
  templateUrl: './carga-masiva-distribucion.component.html',
  styleUrls: ['./carga-masiva-distribucion.component.css'],
  encapsulation: ViewEncapsulation.None,
  providers: [CurrencyPipe]
})
export class CargaMasivaDistribucionComponent implements OnInit {
  nombre_file='';
  files: File[] = [];
  Errores: any[]=[];
  displayedColumns: string[] = ['NumContrato', 'IdPrei', 'Columna', 'Fila', 'Error', 'Tipo'];
  dataSource: any;
  public validado = false;
  public color = '';

  constructor(
    public dialogRef: MatDialogRef<CargaMasivaDistribucionComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
    private _distribucion: DistribucionConceptosService,
    public dialog: MatDialog,
    private curre: CurrencyPipe
  ) {
  }

  ngOnInit() {
    this.dataSource = new MatTableDataSource();
    if (this.data.conceptos) {
      this.nombre_file = 'distribucion_con_conceptos';
    } else {
      this.nombre_file = 'distribucion_sin_concepto';
    }
  }

  // ********************************************************************************************************
  /* ***** Funcion para descargar la plantilla correspondiente  *****   */
  descargar() {
    this._distribucion.DescargarPlantillaDistribucion(this.nombre_file).subscribe(res => {
      if (File) {
        saveAs(res, 'plantilla_' + this.nombre_file + '.xlsx');
      }
    }, err => {
      if (err.error !== '' || err.error !== undefined) {
        Swal.fire({
          position: 'center',
          title: 'Error',
          text: 'No se pudo descargar la plantilla',
          icon: 'error',
          showConfirmButton: true,
        });
      }
    });
  }

  // ********************************************************************************************************
  /* ***** Funcion para cerrar la ventana  *****   */
  onNoClick(): void {
    this.dialogRef.close();
  }

  // ********************************************************************************************************
  /* ***** Funcion para Eliminar el archivo seleccionado  *****   */
  onRemove(event:any) {
    console.log(event);
    this.files.splice(this.files.indexOf(event), 1);
    this.validado = false;
    this.Errores = [];
    this.dataSource.data = this.Errores;
    this.color = '#0089e1';
  }

  // ********************************************************************************************************
  /* *****  Funcion para agregar el archivo  *****   */
  onSelect(event:any) {
    if (this.files.length < 1) {
      console.log(...event.addedFiles);
      this.files.push(...event.addedFiles);
    }
  }

// ********************************************************************************************************
  /* ***** Funcion para validar todos los campos del archivo de distribucion  *****   */
  ValidarArchivo() {
    this.validado = false;
    this.Errores = [];
    this.dataSource.data = this.Errores;
    const formData = new FormData();
    formData.append('file', this.files[0]);
    this._distribucion.ValidarArchivo(formData, this.nombre_file).subscribe(resultado => {
      console.log(resultado);
      this.validado = true;
      this.color = '#18903f';
    }, err => {
      if (err.error !== '' || err.error !== undefined) {
        const erores = <RespuestaPeticion> err.error;
        this.Errores = erores.Data;
        this.dataSource.data = this.Errores;
        console.log(erores.Data);
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
  /* ***** Funcion para validar todos los campos del archivo de distribucion  *****   */
  AgregarArchivo() {
    this.Errores = [];
    this.dataSource.data = this.Errores;
    const formData = new FormData();
    formData.append('file', this.files[0]);
    Swal.fire({
      title: 'Estas seguro de agregar estos datos',
      showCancelButton: true,
      icon: 'warning',
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      confirmButtonText: 'Si, agregar'
    }).then((result) => {
      if (result.value) {
        this._distribucion.AgregarArchivo(formData, this.nombre_file).subscribe(resultado => {
          Swal.fire(
            'Agregados!',
            'Los Datos se han guardado exitosamente',
            'success'
          );
          this.dialogRef.close(true);
        }, err => {
          if (err.error !== '' || err.error !== undefined) {
            const erores = <RespuestaPeticion> err.error;
            this.Errores = erores.Data;
            this.dataSource.data = this.Errores;
            console.log(erores.Data);
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
  /* *****  Funcion para cambiar el mensaje del boton de acciones  *****   */
  MensajeBoton(codigo: number): string {
    if (codigo === 1) {
      return 'Agregar';
    }
    if (codigo === 2 || codigo === 4) {
      return 'Verificar';
    }
    return 'Ver';
  }

  // ********************************************************************************************************
  /* *****  Funcion para determinar que accion realizar dependiendo del error  *****   */
  AbrirAccion(row: any) {

    if (row.Tipo === 1) {
      const dialogRef = this.dialog.open(AgregarUnContratoComponent, {
        height: 'auto',
        width: '90%',
        disableClose: true,
        data: {}
      });
      dialogRef.afterClosed().subscribe(result => {
        if (result === true) {
          const index = this.dataSource.data.indexOf(row);
          console.log(index);
          this.dataSource.data.splice(index, 1);
          this.dataSource = new MatTableDataSource(this.dataSource.data);
        }
      });
    }
    if (row.Tipo === 2) {
      const dialogRef = this.dialog.open(DistribucionConceptosComponent, {
        height: 'auto',
        width: '90%',
        disableClose: true,
        data: {
          concepto: row.Concepto,
          cantidad: row.Cantidad,
          archivo: true
        }
      });
      dialogRef.afterClosed().subscribe(result => {
        if (result === true) {
          this.ValidarArchivo();
        }
      });
    }
    if (row.Tipo === 0 || row.Tipo === 3) {
      Swal.fire({
        position: 'center',
        title: 'Debes editar tu archivo  de carga',
        text: row.Error,
        icon: 'error',
        showConfirmButton: true,
      });
    }
    if (row.Tipo === 4) {
      const dialogRef = this.dialog.open(ExaminarMontoContratoComponent, {
        height: 'auto',
        width: 'auto',
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
    if (row.Tipo === 5) {
      const dialogRef = this.dialog.open(EditarPrecioPreiComponent, {
        height: 'auto',
        width: '65%',
        disableClose: true,
        data: {
          concepto: row.Concepto,
        }
      });
      dialogRef.afterClosed().subscribe(result => {
        if (result === true) {
          this.ValidarArchivo();
        }
      });
    }
  }
}
