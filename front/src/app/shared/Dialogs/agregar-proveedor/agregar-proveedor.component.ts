import {Component, Inject, OnInit} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import Swal from 'sweetalert2';
import {ProveedoresInterface} from '../../../core/models/proveedores.interface';
import {RespuestaPeticion} from '../../../core/models/estructuras_generales';
import {ProveedoresService} from '../../../core/servicios/ProveedoresServicios/proveedores.service';
import {EntidadFederativaService} from '../../../core/servicios/entidad-federativa.service';

export interface DialogDataProveedores {
  editar: boolean;
  proveedor: ProveedoresInterface;
}

@Component({
  selector: 'app-agregar-proveedor',
  templateUrl: './agregar-proveedor.component.html',
  styleUrls: ['./agregar-proveedor.component.css']
})
export class AgregarProveedorComponent implements OnInit {
  public formulario_proveedor!: FormGroup;
  titulo: string='';
  datos!: ProveedoresInterface;
  entidades_arreglo: any;

  constructor(
    public dialogRef: MatDialogRef<AgregarProveedorComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogDataProveedores,
    private fb: FormBuilder,
    private _proveedor: ProveedoresService,
    private _entidades: EntidadFederativaService
  ) {
  }

  ngOnInit() {
    this.datos = this.data.proveedor;
    console.log(this.datos);
    this.crear_formulario_proveedores();
    if (this.data.editar) {
      this.titulo = 'Editar un Proveedor';
    } else {
      this.titulo = 'Agregar un nuevo Proveedor';
    }
    this._entidades.ConsultaEntidades().subscribe(res => {
      this.entidades_arreglo = res.Data;
      console.log(this.entidades_arreglo);
    });
  }

  crear_formulario_proveedores() {
    this.formulario_proveedor = this.fb.group({
      AliasEmpresa: [this.datos.AliasEmpresa, Validators.compose([Validators.pattern('^[a-zA-ZáéíóúÁÉÍÓÚñÑ\\s]*$'), Validators.required])],
      NombreEmpresa: [this.datos.NombreEmpresa, Validators.compose([Validators.required,
        Validators.pattern('^[a-zA-ZáéíóúÁÉÍÓÚñÑ\\s]|[\.]*$')])],
      Direccion: [this.datos.Direccion],
      Correo: [this.datos.Correo, Validators.compose([Validators.required, Validators.email])],
      RFC: [this.datos.RFC],
      NProvImss: this.datos.NProvImss,
      Telefono: [this.datos.Telefono, Validators.compose([Validators.required, Validators.pattern('^([0-9])*$')])],
      EstadosId: [this.datos.EstadosId, Validators.compose([Validators.required])],
      Municipio: [this.datos.Municipio],
      CP: [this.datos.CP]
    });
  }

  cerrarDialog() {
    this.dialogRef.close();
  }

  agregarProveedor() {
    const datos = <ProveedoresInterface>this.formulario_proveedor.getRawValue();
    console.log(datos);
    datos.NProvImss = datos.NProvImss.toString();
    datos.CP = datos.CP.toString();
    Swal.fire({
      title: 'Estas seguro',
      text: 'Deseas agregar este proveedor',
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      confirmButtonText: 'Si Agregar'
    }).then((result) => {
      if (result.value) {
        this._proveedor.AgregarProveedores(datos).subscribe(resultado => {
          console.log(resultado);
          Swal.fire(
            'Agregado',
            'El proveedor ha sido agregado correctamente',
            'success'
          );
          this.dialogRef.close(true);
        }, err => {
          if (err.error !== '' || err.error !== undefined) {
            const erores = <RespuestaPeticion>err.error;
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

  EditarProveedor() {
    const datos = <ProveedoresInterface>this.formulario_proveedor.getRawValue();
    datos.NProvImss = datos.NProvImss.toString();
    datos.CP = datos.CP.toString();
    console.log(datos);
    Swal.fire({
      title: 'Estas seguro',
      text: 'Deseas Actualizar este proveedor',
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      confirmButtonText: 'Si Actualizar'
    }).then((result) => {
      if (result.value) {
        this._proveedor.ActualizarProveedores(datos, this.datos.Id).subscribe(resultado => {
          console.log(resultado);
          Swal.fire(
            'Actualizado',
            'El proveedor ha sido actualizado correctamente',
            'success'
          );
          this.dialogRef.close(true);
        }, err => {
          if (err.error !== '' || err.error !== undefined) {
            const erores = <RespuestaPeticion>err.error;
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
}
