import {Component, Inject, OnInit} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {ContactoProveedor} from '../../../core/models/proveedores.interface';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import Swal from 'sweetalert2';
import {ContactoProveedoresService} from '../../../core/servicios/ProveedoresServicios/contacto-proveedores.service';

export interface DialogData {
  editar: boolean;
  contacto: ContactoProveedor;
  proveedor_id: string;
}

@Component({
  selector: 'app-agregar-contacto-proveedor',
  templateUrl: './agregar-contacto-proveedor.component.html',
  styleUrls: ['./agregar-contacto-proveedor.component.css']
})
export class AgregarContactoProveedorComponent implements OnInit {
  public formulario_conntacto!: FormGroup;
  public titulo: string = '';
  public datos!: ContactoProveedor;

  constructor(
    public dialogRef: MatDialogRef<AgregarContactoProveedorComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
    private fb: FormBuilder,
    private _contactos: ContactoProveedoresService
  ) {
  }

  ngOnInit() {
    this.datos = this.data.contacto;
    this.crear_formulario_contactos();
    if (this.data.editar) {
      this.titulo = 'Editar un Contacto';
    } else {
      this.titulo = 'Agregar un nuevo Contacto';
    }
  }

  crear_formulario_contactos() {
    this.formulario_conntacto = this.fb.group({
      NombreCompleto: [this.datos.NombreCompleto, Validators.compose([Validators.pattern('^[a-zA-ZáéíóúÁÉÍÓÚñÑ\\s]*$'),
        Validators.required])],
      Departamento: [this.datos.Departamento, Validators.compose([Validators.required,
        Validators.pattern('^[a-zA-ZáéíóúÁÉÍÓÚñÑ\\s]*$')])],
      Cargo: [this.datos.Cargo, Validators.compose([Validators.required,
        Validators.pattern('^[a-zA-ZáéíóúÁÉÍÓÚñÑ\\s]*$')])],
      Telefono: [this.datos.Telefono, Validators.compose([Validators.pattern('^([0-9])*$')])],
      Extension: [this.datos.Extension, Validators.compose([Validators.pattern('^([0-9])*$')])],
      Celular: [this.datos.Celular, Validators.compose([Validators.pattern('^([0-9])*$')])],
      Correo: [this.datos.Correo, Validators.compose([Validators.email])],
      Comentarios: [this.datos.Comentarios],
      Asunto: [this.datos.Asunto],
    });
  }

  agregarContacto() {
    const datos = <ContactoProveedor> this.formulario_conntacto.getRawValue();
    datos.ProveedorNProvImss = this.data.proveedor_id;
    console.log(datos);
    Swal.fire({
      title: 'Estas seguro',
      text: 'Deseas agregar este contacto',
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      confirmButtonText: 'Si Agregar'
    }).then((result) => {
      if (result.value) {
        this._contactos.AgregarContactoProveedor(datos).subscribe(resultado => {
          console.log(resultado);
          Swal.fire(
            'Agregado',
            'El contacto ha sido agregado correctamente',
            'success'
          );
          this.dialogRef.close(true);
        });
      }
    });
  }

  EditarContacto() {
    const datos = <ContactoProveedor> this.formulario_conntacto.getRawValue();
    datos.ProveedorNProvImss = this.data.proveedor_id;
    console.log(datos);
    Swal.fire({
      title: 'Estas seguro',
      text: 'Deseas Actualizar este contacto',
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      confirmButtonText: 'Si Actualizar'
    }).then((result) => {
      if (result.value) {
        this._contactos.ActualizarContactoProveedor(datos, this.datos.Id).subscribe(resultado => {
          console.log(resultado);
          Swal.fire(
            'Actualizado',
            'El contacto ha sido actualizado correctamente',
            'success'
          );
          this.dialogRef.close(true);
        });
      }
    });
  }

  cerrarDialog() {
    this.dialogRef.close();
  }
}
