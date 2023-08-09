import {Component, Inject, OnInit} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {FormArray, FormBuilder, Validators} from '@angular/forms';
import {CorreoContrato} from '../../../../core/models/representante_legal.interface';
import Swal from 'sweetalert2';
import {RespuestaPeticion} from '../../../../core/models/estructuras_generales';
import {RepresentanteLegalService} from '../../../../core/servicios/ContratosServicios/representante-legal.service';
import {ContactosContratoService} from '../../../../core/servicios/ContratosServicios/contactos-contrato.service';

export interface DialogData {
  representante: any;
  num_contrato: string;
  editar: boolean;
}

@Component({
  selector: 'app-representante-legal',
  templateUrl: './representante-legal.component.html',
  styleUrls: ['./representante-legal.component.css']
})
export class RepresentanteLegalComponent implements OnInit {
  titulo = 'Agregar Representante Legal';
  representante!: any;
  correo_contrato!: CorreoContrato;

  constructor(
    public dialogRef: MatDialogRef<RepresentanteLegalComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
    private  _representante: RepresentanteLegalService,
    private formBuilder: FormBuilder,
    private  _contactoCont: ContactosContratoService,
  ) {
    if (data.editar) {
      this.titulo = 'Editar Representante Legal';
      this.consultarRepresentante();
      this.consultarCorreos();
      this.consultarTelefonos();
    }
    this.representante = this.data.representante;
    console.log('Esto llega en el data', this.representante);
  }

  get Correos() {
    return this.representanteForm.get('Correos') as FormArray;
  }

  get Telefonos() {
    return this.representanteForm.get('Telefonos') as FormArray;
  }

  representanteForm = this.formBuilder.group({
    NombreCompleto: [null, Validators.compose([Validators.required])],
    Correos: this.formBuilder.array([]),
    Telefonos: this.formBuilder.array([])
  });

  ngOnInit() {
    if (!this.data.editar && this.representante) {
      this.representanteForm.controls['NombreCompleto'].setValue(this.representante.NombreCompleto);
      for (const email of this.representante.Correos) {
        console.log(email);
        const correoFormGroup = this.formBuilder.group({
          Correo: [email.Correo, Validators.compose([Validators.email, Validators.required])],
          TipoCorreo: [email.TipoCorreo, Validators.compose([Validators.required])],
        });
        this.Correos.push(correoFormGroup);
      }
      for (const tel of this.representante.Telefonos) {
        const telefonoFormGroup = this.formBuilder.group({
          Telefono: [tel.Telefono, Validators.compose([Validators.required, Validators.pattern('^([0-9])*$')])],
          Extension: [tel.Extension, Validators.compose([Validators.pattern('^([0-9])*$')])],
          TipoTelefono: [tel.TipoTelefono, Validators.compose([Validators.required])]
        });
        this.Telefonos.push(telefonoFormGroup);
      }
    }
    if (!this.data.editar && this.representante == null || this.representante == '') {
      console.log('Mejecuteeee');
      this.agregarCorreo();
      this.agregarTelefonos();
    }
  }

  agregarCorreo() {
    const correoFormGroup = this.formBuilder.group({
      Correo: [null, Validators.compose([Validators.email, Validators.required])],
      TipoCorreo: [null, Validators.compose([Validators.required])],
    });
    this.Correos.push(correoFormGroup);
  }

  borrarCorreo(indice: number) {
    const elemento = this.Correos.at(indice).value;
    console.log(elemento);
    if (elemento.Id) {
      Swal.fire({
        title: 'Eliminar',
        text: '¿Estas seguro de eliminar este correo?',
        icon: 'error',
        showCancelButton: true,
        confirmButtonColor: '#3085d6',
        cancelButtonColor: '#a2a2a2',
        confirmButtonText: 'Si, Eliminar',
        cancelButtonText: 'Cancelar',
        reverseButtons: true
      }).then((result) => {
        if (result.value) {
          console.log('Se tiene que eliminar de la base de datos');
          this._contactoCont.EliminarCorreoContrato(elemento.Id).subscribe(res => {
            Swal.fire(
              'Correcto',
              'El correo ha sido eliminado',
              'success'
            );
            this.Correos.removeAt(indice);
          }, err => {
            if (err.error !== '' || err.error !== undefined) {
              const erores = <RespuestaPeticion>err.error;
              Swal.fire({
                icon: 'error',
                title: 'Ocurrio un error',
                text: erores.Mensaje,
              });
            }
          });
        }
      });
    } else {
      console.log('Solo eliminar del arreglo');
      this.Correos.removeAt(indice);
    }
  }

  agregarTelefonos() {
    const telefonoFormGroup = this.formBuilder.group({
      Telefono: [null, Validators.compose([Validators.required, Validators.pattern('^([0-9])*$')])],
      Extension: [null, Validators.compose([Validators.pattern('^([0-9])*$')])],
      TipoTelefono: [null, Validators.compose([Validators.required])]
    });
    this.Telefonos.push(telefonoFormGroup);
  }

  borrarTelefono(indice: number) {
    const elemento = this.Telefonos.at(indice).value;
    console.log(elemento);
    if (elemento.Id) {
      Swal.fire({
        title: 'Eliminar',
        text: '¿Estas seguro de eliminar este telefono?',
        icon: 'error',
        showCancelButton: true,
        confirmButtonColor: '#3085d6',
        cancelButtonColor: '#a2a2a2',
        confirmButtonText: 'Si, Eliminar',
        cancelButtonText: 'Cancelar',
        reverseButtons: true
      }).then((result) => {
        if (result.value) {
          console.log('Se tiene que eliminar de la base de datos');
          this._contactoCont.EliminarTelefonoContrato(elemento.Id).subscribe(res => {
            Swal.fire(
              'Correcto',
              'El telefono ha sido eliminado',
              'success'
            );
            this.Telefonos.removeAt(indice);
          }, err => {
            if (err.error !== '' || err.error !== undefined) {
              const erores = <RespuestaPeticion>err.error;
              Swal.fire({
                icon: 'error',
                title: 'Ocurrio un error',
                text: erores.Mensaje,
              });
            }
          });
        }
      });
    } else {
      console.log('Solo eliminar del arreglo');
      this.Telefonos.removeAt(indice);
    }
  }

  consultarRepresentante() {
    this._representante.ConsultaRepresentante(this.data.num_contrato).subscribe(resul => {
      this.representante = resul.Data;
      console.log(this.representante);
      this.representanteForm.controls['NombreCompleto'].setValue(this.representante.NombreCompleto);
    });
  }

  onNoClick(): void {
    this.dialogRef.close();
  }

  guardar() {
    console.log(this.representanteForm.value);
    this.dialogRef.close(this.representanteForm.value);
  }

  editarRepresentante() {
    Swal.fire({
      title: 'Actualizacion',
      text: '¿Estas seguro de actualizar los datos?',
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#a2a2a2',
      reverseButtons: true,
      confirmButtonText: 'Si, Actualizar',
      cancelButtonText: 'Cancelar'
    }).then((result) => {
      if (result.value) {
        this._representante.EditarRepresentante(this.data.num_contrato, this.representanteForm.value).subscribe(res => {
          Swal.fire(
            'Correcto',
            'Los datos han sido actualizados',
            'success'
          );
          this.dialogRef.close(true);
        }, err => {
          if (err.error !== '' || err.error !== undefined) {
            const erores = <RespuestaPeticion>err.error;
            Swal.fire({
              icon: 'error',
              title: 'Ocurrio un error',
              text: erores.Mensaje,
            });
          }
        });
      }
    });
  }

  consultarCorreos() {
    this._contactoCont.ConsultaCorreosContrato(this.data.num_contrato).subscribe(res => {
      this.correo_contrato = res.Data;
      console.log('Correos', this.correo_contrato);
      const correo = res.Data.map((ite:any) => {
        return {
          id: ite.Id,
          correo: ite.Correos.Correo,
          tipo: ite.TipoCorreo
        };
      });
      for (const email of correo) {
        const correoFormGroup = this.formBuilder.group({
          Id: [email.id],
          Correo: [email.correo, Validators.compose([Validators.email, Validators.required])],
          TipoCorreo: [email.tipo, Validators.compose([Validators.required])],
        });
        this.Correos.push(correoFormGroup);
      }
    });
  }

  consultarTelefonos() {
    this._contactoCont.ConsultaTelefonosContrato(this.data.num_contrato).subscribe(res => {
      console.log('Telefonosssss', res.Data);
      const telefono = res.Data.map((ite:any) => {
        return {
          id: ite.Id,
          telefono: ite.Telefonos.Telefono,
          tipo: ite.TipoTelefono,
          extension: ite.Extension
        };
      });
      console.log(telefono);
      for (const tel of telefono) {
        const telefonoFormGroup = this.formBuilder.group({
          Id: [tel.id],
          Telefono: [tel.telefono, Validators.compose([Validators.required, Validators.pattern('^([0-9])*$')])],
          Extension: [tel.extension, Validators.compose([Validators.pattern('^([0-9])*$')])],
          TipoTelefono: [tel.tipo, Validators.compose([Validators.required])]
        });
        this.Telefonos.push(telefonoFormGroup);
      }
    });
  }
}
