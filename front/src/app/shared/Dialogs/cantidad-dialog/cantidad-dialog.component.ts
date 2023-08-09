import {Component, Inject, OnInit} from '@angular/core';
import {MatDialogRef, MAT_DIALOG_DATA} from '@angular/material/dialog';
import {
  FormBuilder,
  FormGroup,
  Validators,
} from '@angular/forms';
import {SaiService} from '../../../catalogos/servicios/sai.service';
import {UsuarioLogeado} from '../../../core/models/UsuarioLogin.interface';
import jwt_decode from 'jwt-decode';
import Swal from 'sweetalert2';
import {RespuestaPeticion} from '../../../core/models/estructuras_generales';
import {ClavesService} from '../../../core/servicios/solicitud/claves.service';

export interface DialogData {
  clave: any;
  aprobar: boolean;
  agregar: boolean;
  editar: boolean;
}

@Component({
  selector: 'app-cantidad-dialog',
  templateUrl: './cantidad-dialog.component.html',
  styleUrls: ['./cantidad-dialog.component.css']
})
export class CantidadDialogComponent implements OnInit {
  producto: any;
  public form!: FormGroup;
  public total = 0;
  public cantidad = 0;
  public cantidadSemestral = 0;
  public cantidadBimestral = 0;
  public cantidadAnual = 0;
  public cantidadMes = 0;
  public claveSai: any;
  public usuarioLog!: UsuarioLogeado;
  public editar = false;
  public aprobar = false;
  public agregar = false;
  public claveConsulta: any;

  constructor(
    private fb: FormBuilder,
    public dialogRef: MatDialogRef<CantidadDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
    private _claveSai: SaiService,
    private _conceptosSolicitud: ClavesService,
  ) {
    this.editar = data.editar;
    this.agregar = data.agregar;
    this.aprobar = data.aprobar;
  }

  ngOnInit() {
    this.form = this.fb.group({
      Cantidad: [null!, Validators.compose([Validators.required, Validators.pattern('^[1-9][0-9]*$')])],
      CantidadSemestre: [null!, Validators.compose([Validators.required, Validators.pattern('^[1-9][0-9]*.?[0-9]*$')])],
      CantidadBimestre: [null, Validators.compose([Validators.required, Validators.pattern('^[1-9][0-9]*.?[0-9]*$')])],
      CantidadMes: [null, Validators.compose([Validators.required, Validators.pattern('^[1-9][0-9]*.?[0-9]*$')])],
      Comentario: [null, Validators.compose([Validators.required,
        Validators.pattern('^[a-zA-ZáéíóúÁÉÍÓÚñÑ]+(\\s[a-zA-Z_áéíóúÁÉÍÓÚñÑ]+)*$')])],
    });
    this.usuarioLog = jwt_decode(<string>localStorage.getItem('token'));
    this.producto = this.data.clave;
    this.total = this.producto.Total;
    this.cantidad = this.producto.CantidadSolicitada;
    console.log(this.producto);
    this.consultarClaveSai();
    this.quitarControlComentario();
  }

  quitarControlComentario() {
    if (this.agregar) {
      this.form.removeControl('Comentario');
      console.log(this.producto.claveId);
    }
  }

  consultarClaveSai() {
    if (this.agregar) {
      this.claveConsulta = this.producto.Grupo + this.producto.Generico + this.producto.Especifico + this.producto.Diferenciador +
        this.producto.Variable;
    } else {
      this.claveConsulta = this.producto.ClaveId;
    }
    this._claveSai.getClaveSai('', this.claveConsulta).subscribe(res => {
      console.log(res.Data);
      this.claveSai = res.Data;
    });
  }

  onNoClick(): void {
    this.dialogRef.close(false);
  }

  inputCantidad() {
    if (this.form.controls['Cantidad'].valid) {
      const cantidadTemporal = this.form.get('Cantidad')!.value ;
      this.cantidad = cantidadTemporal;
      this.cantidadAnual = cantidadTemporal;
      this.cantidadSemestral = cantidadTemporal / 2;
      this.cantidadBimestral = (cantidadTemporal / 6);
      this.cantidadMes = cantidadTemporal / 12;
      this.form.controls['CantidadSemestre'].setValue(this.cantidadSemestral);
      this.form.controls['CantidadBimestre'].setValue(this.cantidadBimestral);
      this.form.controls['CantidadMes'].setValue(this.cantidadMes);
    }
  }

  inputCantidadSemestre() {
    if (this.form.controls['CantidadSemestre'].valid) {
      const cantidadSem = this.form.get('CantidadSemestre')!.value;
      this.cantidad = cantidadSem * 2;
      this.cantidadAnual = cantidadSem * 2;
      this.cantidadSemestral = cantidadSem;
      this.cantidadBimestral = cantidadSem / 3;
      this.cantidadMes = cantidadSem / 6;
      this.form.controls['Cantidad'].setValue(this.cantidadAnual);
      this.form.controls['CantidadBimestre'].setValue(this.cantidadBimestral);
      this.form.controls['CantidadMes'].setValue(this.cantidadMes);
    }
  }

  inputCantidadBimestre() {
    if (this.form.controls['CantidadBimestre'].valid) {
      const cantidadBiMes = this.form.get('CantidadBimestre')!.value;
      this.cantidad = cantidadBiMes * 6;
      this.cantidadAnual = cantidadBiMes * 6;
      this.cantidadSemestral = cantidadBiMes * 3;
      this.cantidadBimestral = cantidadBiMes;
      this.cantidadMes = cantidadBiMes / 2;
      this.form.controls['Cantidad'].setValue(this.cantidadAnual);
      this.form.controls['CantidadSemestre'].setValue(this.cantidadSemestral);
      this.form.controls['CantidadMes'].setValue(this.cantidadMes);
    }
  }

  inputCantidadMes() {
    if (this.form.controls['CantidadMes'].valid) {
      const cantidadMes = this.form.get('CantidadMes')!.value;
      this.cantidad = cantidadMes * 12;
      this.cantidadAnual = cantidadMes * 12;
      this.cantidadSemestral = cantidadMes * 6;
      this.cantidadBimestral = cantidadMes * 2;
      this.cantidadMes = cantidadMes;
      this.form.controls['Cantidad'].setValue(this.cantidadAnual);
      this.form.controls['CantidadSemestre'].setValue(this.cantidadSemestral);
      this.form.controls['CantidadBimestre'].setValue(this.cantidadBimestral);
    }
  }

  actualizarCantidad() {
    const comentario = this.form.get('Comentario')!.value;
    const cantidad = parseInt((this.form.get('Cantidad')!.value), 10);
    const total = cantidad * this.producto.Precio;
    this._conceptosSolicitud.ActualizarClave(this.producto.Id, comentario, '', this.usuarioLog.UsuarioId, '', {
      SolicitudId: this.producto.SolicitudId,
      ClaveId: this.producto.ClaveId,
      Descripcion: this.producto.Descripcion,
      CantidadSolicitada: cantidad,
      Precio: this.producto.Precio,
      Total: total,
      ServiciosProformaId: this.producto.ServicioProformaId,
      EstatusId: this.producto.EstatusId,
      Fecha: this.producto.Fecha
    }).subscribe(res => {
      Swal.fire({
        icon: 'success',
        title: 'Correcto',
        text: res.Mensaje,
      });
      this.dialogRef.close(res.Data);
    }, err => {
      if (err.error !== '' || err.error !== undefined) {
        const erores = <RespuestaPeticion>err.error;
        Swal.fire({
          icon: 'error',
          title: 'Error',
          text: erores.Mensaje,
        });
      }
      console.log('ocurrio un error', err);
    });
  }

  aprobarCantidad() {
    const comentario = this.form.get('Comentario')!.value;
    const cantidad = (this.form.get('Cantidad')!.value);
    this._conceptosSolicitud.ActualizarClave(this.producto.Id, comentario, cantidad, '', '', {
      SolicitudId: this.producto.SolicitudId,
      ClaveId: this.producto.ClaveId,
      Descripcion: this.producto.Descripcion,
      CantidadSolicitada: this.producto.CantidadSolicitada,
      Precio: this.producto.Precio,
      Total: this.producto.Precio * cantidad,
      ServiciosProformaId: this.producto.ServiciosProformaId,
      EstatusId: this.producto.EstatusId,
      Fecha: this.producto.Fecha
    }).subscribe(res => {
      console.log(res.Data);
      Swal.fire({
        icon: 'success',
        title: 'Correcto',
        text: res.Mensaje,
      });
      this.dialogRef.close(res.Data);
    });
  }

  agregarCantidad() {
    this.dialogRef.close(this.form.get('Cantidad')!.value);
  }
}
