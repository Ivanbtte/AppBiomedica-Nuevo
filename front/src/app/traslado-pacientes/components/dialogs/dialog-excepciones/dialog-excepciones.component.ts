import {Component, Inject, OnInit} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {FormBuilder, FormControl, FormGroup, Validators} from '@angular/forms';

export interface DialogData {
  tipoBoleto: number;
  edad: number;
  exepcion: Excepciones;
}

export interface Excepciones {
  value: string;
  motivo: string;
}

@Component({
  selector: 'app-dialog-excepciones',
  templateUrl: './dialog-excepciones.component.html',
  styleUrls: ['./dialog-excepciones.component.css']
})
export class DialogExcepcionesComponent implements OnInit {
  excep: Excepciones[] = [
    {value: '1', motivo: 'Recoleccion de medicamentos'},
    {value: '2', motivo: 'Acompañante extra'},
    {value: '3', motivo: 'Medio boleto para niños menores de 5 años'}
  ];
  medicamentos = false;
  acompaniante = false;
  medioBoleto = false;
  controlMedicamentos = new FormControl('', Validators.required);
  controlBoleto = new FormControl('', Validators.required);
  formAcompaniante!: FormGroup;
  mayorEdad = true;
  eliminar = false;

  constructor(
    public dialogRef: MatDialogRef<DialogExcepcionesComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
    private fb: FormBuilder,
  ) {
    if (data.tipoBoleto != 0) {
      this.excep.pop();
    }
  }

  ngOnInit() {
    this.formAcompaniante = this.fb.group({
      justificacion: [null, Validators.compose([Validators.required])],
    });
    if (this.data.edad >= 18) {
      this.mayorEdad = false;
    }
    if (this.data.exepcion != null || this.data.exepcion != undefined) {
      this.eliminar = true;
    }
  }

  motivoSeleccionado(value: string) {
    this.medicamentos = false;
    this.acompaniante = false;
    this.medioBoleto = false;
    this.controlBoleto.reset();
    this.controlMedicamentos.reset();
    this.formAcompaniante.reset();
    const valor: number = parseInt(value, 10);
    switch (valor) {
      case 1:
        this.medicamentos = true;
        break;
      case 2:
        this.acompaniante = true;
        break;
      case 3:
        this.medioBoleto = true;
        break;
    }
  }

  onNoClick(): void {
    this.dialogRef.close();
  }

  habilitarBoton(){
    if (this.medicamentos || this.medioBoleto || this.acompaniante) {
      if (this.controlMedicamentos.invalid && this.medicamentos) {
        return true;
      }
      if (this.controlBoleto.invalid && this.medioBoleto) {
        return true;
      }
      if (this.formAcompaniante.invalid && this.acompaniante) {
        return true;
      }
    } else {
      return true;
    }

  }

  eliminarExcepciones() {
    this.dialogRef.close({valor: '', exepcion: 4});
  }

  guardar() {
    if (this.medicamentos) {
      this.dialogRef.close({valor: this.controlMedicamentos.value, exepcion: 1});
    }
    if (this.acompaniante) {
      this.dialogRef.close({valor: this.formAcompaniante.getRawValue(), exepcion: 2});
    }
    if (this.medioBoleto) {
      this.dialogRef.close({valor: this.controlBoleto.value, exepcion: 3});
    }
  }

}
