import {Component, Inject, OnInit, ViewEncapsulation} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialog, MatDialogRef} from '@angular/material/dialog';
import {MatTableDataSource} from '@angular/material/table';
import {RespuestaPeticion} from '../../../../core/models/estructuras_generales';
import Swal from 'sweetalert2';
import {AgregarProveedorComponent} from '../../agregar-proveedor/agregar-proveedor.component';
import {ConsultarContratoComponent} from '../consultar-contrato/consultar-contrato.component';
import {saveAs} from 'file-saver';
import jwt_decode from 'jwt-decode';
import {FormControl} from '@angular/forms';
import {ContratoService} from '../../../../core/servicios/ContratosServicios/contrato.service';
import {ContratoPuestoService} from '../../../../core/servicios/ContratosServicios/contrato-puesto.service';

export interface DialogData {
  animal: string;
  name: string;
}

@Component({
  selector: 'app-carga-csv',
  templateUrl: './carga-masiva-contratos.component.html',
  styleUrls: ['./carga-masiva-contratos.component.css'],
  encapsulation: ViewEncapsulation.None,


})
export class CargaMasivaContratosComponent implements OnInit {
  files: File[] = [];
  Errores: any[] = [];
  displayedColumns: string[] = ['NumContrato', 'Columna', 'Linea', 'Error', 'Tipo'];
  dataSource: any;
  public validado = false;
  public color = '';
  tipo_contrato: any[] = [];
  tipo = new FormControl();
  valor_tipo = '';

  constructor(
    private _contrato: ContratoService,
    public dialog: MatDialog,
    private contrato_puesto: ContratoPuestoService,
    public dialogRef: MatDialogRef<CargaMasivaContratosComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData) {
  }

  onNoClick(): void {
    this.dialogRef.close();
  }

  ngOnInit() {
    const decode: any = jwt_decode(<string> localStorage.getItem('token'));
    this.contrato_puesto.Consultarcontrato_puesto(decode.PuestoId).subscribe(res => {
      this.tipo_contrato = res.Data.map((ite: any) => {
        return {
          id: ite.TipoContrato.Id,
          tipo: ite.TipoContrato.Tipo
        };
      });
    });
    this.dataSource = new MatTableDataSource();
  }

  SubirArchivo() {
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
        this._contrato.SubirArchivo(formData, this.valor_tipo).subscribe(resul => {
          console.log(resul);
          Swal.fire({
            position: 'center',
            title: resul.Mensaje,
            icon: 'success',
            showConfirmButton: true,
          });
          this.dialogRef.close(true);
        }, err => {
          if (err.error !== '' || err.error !== undefined) {
            const erores = <RespuestaPeticion> err.error;
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

  ValidarArchivo() {
    this.validado = false;
    this.Errores = [];
    this.dataSource.data = this.Errores;
    const formData = new FormData();
    formData.append('file', this.files[0]);
    this._contrato.ValidarArchivo(formData).subscribe(resul => {
      Swal.fire({
        position: 'center',
        title: resul.Mensaje,
        icon: 'success',
        showConfirmButton: true,
      });
      this.validado = true;
      this.color = '#18903f';
    }, err => {
      if (err.error !== '' || err.error !== undefined) {
        const erores = <RespuestaPeticion> err.error;
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

  cambio(value: any) {
    this.valor_tipo = value;
  }

  // ********************************************************************************************************
  /* *****  Funcion para determinar que accion realizar dependiendo del error  *****   */
  AbrirAccion(row: any) {
    if (row.Tipo === 2) {
      const index1 = this.dataSource.data.indexOf(row);
      console.log(index1);
      const dialogRef = this.dialog.open(AgregarProveedorComponent, {
        height: 'auto',
        width: '80%',
        disableClose: true,
        data: {
          editar: false,
          proveedor: ''
        }
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
    if (row.Tipo === 1) {
      const dialogRef = this.dialog.open(ConsultarContratoComponent, {
        height: 'auto',
        width: '50%',
        disableClose: true,
        data: {
          num_contrato: row.NumContra,
        }
      });
      dialogRef.afterClosed().subscribe(result => {
      });
    }
  }

  // ********************************************************************************************************
  /* *****  Funcion para cambiar el mensaje del boton de acciones  *****   */
  MensajeBoton(codigo: number): string {
    if (codigo === 1) {
      return 'Consultar';
    }
    if (codigo === 2) {
      return 'Agregar';
    }
    return 'Ver';
  }

  // ********************************************************************************************************
  /* *****   descargar archivo *****   */
  descargar() {
    this._contrato.DescargarPlantillaContrato().subscribe(res => {
      console.log(res);
      if (File) {
        saveAs(res, 'plantilla_contrato.xlsx');
      }
    });
  }

  // ********************************************************************************************************
  /* *****  Funcion para agregar un archivo a el arreglo   *****   */
  onSelect(event: any) {
    if (this.files.length < 1) {
      console.log(...event.addedFiles);
      this.files.push(...event.addedFiles);
    }
  }

  // ********************************************************************************************************
  /* *****  Funcion para eliminar el archivo  *****   */
  onRemove(event: any) {
    console.log(event);
    this.files.splice(this.files.indexOf(event), 1);
    this.validado = false;
    this.Errores = [];
    this.dataSource.data = this.Errores;
    this.color = '#0089e1';
  }
}
