import {Component, OnInit} from '@angular/core';
import {MatDialogRef} from '@angular/material/dialog';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {UsuarioService} from '../../../../authentication/servicios/usuario.service';
import Swal from 'sweetalert2';
import {RespuestaPeticion} from '../../../../core/models/estructuras_generales';
import {FirmaSolicitud} from '../../../../core/models/ServiciosSubrogados/trasladoPacientes/firma_solicitud.interface';
import {FirmService} from '../../../../core/servicios/ServiciosSubrogados/traslados/firm.service';

export interface Type {
  value: string;
  tipo: string;
}

@Component({
  selector: 'app-dialog-new-firm',
  templateUrl: './dialog-new-firm.component.html',
  styleUrls: ['./dialog-new-firm.component.css']
})
export class DialogNewFirmComponent implements OnInit {
  formFirm!: FormGroup;
  unitId = 0;
  pattern = '^[a-zA-ZáéíóúÁÉÍÓÚñÑ\.]+( [a-zA-ZáéíóúÁÉÍÓÚñÑ\.]+)*$';
  types: Type[] = [];
  firmObject!: Partial<FirmaSolicitud>;

  constructor(
    private fb: FormBuilder,
    public dialogRef: MatDialogRef<DialogNewFirmComponent>,
    private usuarioService: UsuarioService,
    private firmService: FirmService,
  ) {
    this.types = [
      {value: '1', tipo: 'Vo. Bo.'},
      {value: '2', tipo: 'Autoriza'},
    ];
    this.createForm();
  }

  ngOnInit() {
  }

  getUnit(id: string): boolean {
    let resultPet = true;
    this.usuarioService.unidadUsuario(id).subscribe(res => {
      console.log(res.Data)
      this.unitId = res.Data.unidadMedId;
      this.firmObject.usuarioId = id;
      this.firmObject.unidadMedId = this.unitId;
      console.log(this.unitId, this.firmObject)
    }, err => {
      if (err.error !== '' || err.error !== undefined) {
        const erores = <RespuestaPeticion> err.error;
        Swal.fire({
          icon: 'error',
          title: 'Error',
          text: erores.Mensaje,
          showConfirmButton: true,
        });
        resultPet = false;
      }
    });
    return resultPet;
  }

  createForm() {
    this.formFirm = this.fb.group({
      nombre: ['', Validators.compose([Validators.required, Validators.pattern(this.pattern)])],
      cargo: ['', Validators.compose([Validators.required, Validators.pattern(this.pattern)])],
      matricula: ['', Validators.compose([Validators.required, Validators.pattern('^([0-9])*$')])],
      tipo: ['', Validators.compose([Validators.required])],
    });

  }

  onNoClick(): void {
    this.dialogRef.close();
  }

  saveFirm() {
    const idUser = <string>localStorage.getItem('idUser');
    if (this.getUnit(idUser)) {
      this.firmObject = this.formFirm.getRawValue();
      const tempt = this.formFirm.get('tipo')!.value;
      this.firmObject.tipo = parseInt(tempt.value, 10);
      console.log(this.firmObject);
      Swal.fire({
        title: 'Estas seguro de agregar estos datos',
        showCancelButton: true,
        icon: 'warning',
        confirmButtonColor: '#00b162',
        cancelButtonColor: '#d33',
        confirmButtonText: 'Enviar',
        cancelButtonText: 'Cancelar'
      }).then((result) => {
        if (result.value) {
          this.firmService.AgregarFirma(this.firmObject).subscribe(resp => {
            console.log(resp.Mensaje);
            Swal.fire({
              position: 'center',
              title: resp.Mensaje,
              icon: 'success',
              showConfirmButton: true,
            });
            this.dialogRef.close(true);
          }, err => {
            if (err.error !== '' || err.error !== undefined) {
              const erores = <RespuestaPeticion> err.error;
              Swal.fire({
                position: 'center',
                title: 'Duplicado',
                text: erores.Mensaje,
                icon: 'error',
                showConfirmButton: true,
              });
            }
          });
        }
      });
    }
  }
}
