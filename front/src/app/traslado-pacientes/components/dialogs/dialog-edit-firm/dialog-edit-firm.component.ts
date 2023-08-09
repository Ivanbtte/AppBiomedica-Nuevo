import {Component, Inject, OnInit} from '@angular/core';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {UsuarioService} from '../../../../authentication/servicios/usuario.service';
import {FirmaSolicitud} from '../../../../core/models/ServiciosSubrogados/trasladoPacientes/firma_solicitud.interface';
import {ReturnTypeFirmPipe} from '../../../../shared/pipes/return-type-firm.pipe';
import {FirmService} from '../../../../core/servicios/ServiciosSubrogados/traslados/firm.service';
import Sweet from 'sweetalert2';
import {RespuestaPeticion} from '../../../../core/models/estructuras_generales';

export interface DialogData {
  firm: FirmaSolicitud
}

export interface Type {
  value: string;
  tipo: string;
}

@Component({
  selector: 'app-dialog-edit-firm',
  templateUrl: './dialog-edit-firm.component.html',
  styleUrls: ['./dialog-edit-firm.component.css']
})
export class DialogEditFirmComponent implements OnInit {
  formFirm!: FormGroup;
  unitId = 0;
  firmObj!: Partial<FirmaSolicitud>;
  pattern = '^[a-zA-ZáéíóúÁÉÍÓÚñÑ]+( [a-zA-Z_áéíóúÁÉÍÓÚñÑ]+)*$';
  types: Type[] = [
    {value: '1', tipo: 'Vo. Bo.'},
    {value: '2', tipo: 'Autoriza'},
  ];
  public changeDetect = false;

  constructor(
    private fb: FormBuilder,
    public dialogRef: MatDialogRef<DialogEditFirmComponent>,
    private usuarioService: UsuarioService,
    private firmService: FirmService,
    @Inject(MAT_DIALOG_DATA) public data: DialogData
  ) {
    this.createForm();
  }

  ngOnInit() {
  }

  createForm() {
    this.formFirm = this.fb.group({
      nombre: ['', Validators.compose([Validators.required, Validators.pattern(this.pattern)])],
      cargo: ['', Validators.compose([Validators.required, Validators.pattern(this.pattern)])],
      tipo: ['', Validators.compose([Validators.required])],
      matricula: ['', Validators.compose([Validators.required, Validators.pattern('^([0-9])*$')])],
    });
    this.formFirm.patchValue(this.data.firm);
    const returnTypeFirmPipe = new ReturnTypeFirmPipe();
    this.formFirm.controls['tipo'].setValue(returnTypeFirmPipe.transform(this.data.firm.tipo));
  }

  onNoClick(): void {
    this.dialogRef.close();
  }

  compareFn(c1: { tipo: number }, c2: number) {
    return c1.tipo === c2;
  }

  changeSelect() {
    this.changeDetect = true;
  }

  editFirm() {
    Sweet.fire({
      title: 'Estas seguro de agregar estos datos',
      showCancelButton: true,
      icon: 'warning',
      confirmButtonColor: '#00b162',
      cancelButtonColor: '#d33',
      confirmButtonText: 'Enviar',
      cancelButtonText: 'Cancelar'
    }).then((result) => {
      if (result.value) {
        this.firmObj = this.formFirm.getRawValue();
        const temp = this.formFirm.get('tipo')!.value;
        this.firmObj.tipo = parseInt(temp.value, 10);
        if (!this.changeDetect) {
          delete this.firmObj.tipo;
        }
        let matric = '';
        if (this.data.firm.matricula !== this.firmObj.matricula) {
          matric = 'si';
        }
        const userId = <string>localStorage.getItem('idUser');
        this.firmService.UpdateFirm(this.firmObj, this.data.firm.id, this.data.firm.unidadMedId.toString(), matric, userId).subscribe(resp => {
          Sweet.fire({
            position: 'center',
            title: resp.Mensaje,
            icon: 'success',
            showConfirmButton: true,
          });
          console.log(resp.Data);
          this.dialogRef.close(true);
        }, err => {
          if (err.error !== '' || err.error !== undefined) {
            const erores = <RespuestaPeticion> err.error;
            Sweet.fire({
              position: 'center',
              title: 'Error',
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
