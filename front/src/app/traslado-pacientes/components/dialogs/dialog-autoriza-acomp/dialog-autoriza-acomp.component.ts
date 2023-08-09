import { Component, OnInit, Inject } from "@angular/core";
import { MatDialogRef, MAT_DIALOG_DATA } from "@angular/material/dialog";
import { FormBuilder, FormGroup, Validators } from "@angular/forms";
import { Acompaniante } from "../../../../core/models/ServiciosSubrogados/trasladoPacientes/solicitud_traslado.interface";
import { AcompanianteService } from "../../../core/acompaniante.service";

export interface Justificacion {
  id: string;
  justificacion: string;
}
export interface DialogData {
  edad: number;
  unidadMed: string;
}

@Component({
  selector: "app-dialog-autoriza-acomp",
  templateUrl: "./dialog-autoriza-acomp.component.html",
  styleUrls: ["./dialog-autoriza-acomp.component.css"],
})
export class DialogAutorizaAcompComponent implements OnInit {
  datosAutorizacion!: Acompaniante;
  justificaciones: Justificacion[] = [
    {
      id: "Paciente mayor que no puede valerse por sí mismo",
      justificacion: "Paciente mayor que no puede valerse por sí mismo",
    },
    {
      id: "Paciente con padecimientos neuropsiquiátricos",
      justificacion: "Paciente con padecimientos neuropsiquiátricos",
    },
    {
      id: "Paciente con padecimientos invalidantes",
      justificacion: "Paciente con padecimientos invalidantes",
    },
    {
      id: "Paciente con una urgencia médica",
      justificacion: "Paciente con una urgencia médica",
    },
    {
      id: "Paciente programado para cirugía de alta especialidad",
      justificacion: "Paciente programado para cirugía de alta especialidad",
    },
    {
      id: "Paciente con padecimientos que pongan en riesgo su vida",
      justificacion: "Paciente con padecimientos que pongan en riesgo su vida",
    },
    {
      id: "**No identificable en FTP-01**",
      justificacion: "**No identificable en FTP-01**",
    },
    {
      id: "Paciente menor de edad",
      justificacion: "Paciente menor de edad",
    },
  ];
  form!: FormGroup;

  constructor(
    public dialogRef: MatDialogRef<DialogAutorizaAcompComponent>,
    private fb: FormBuilder,
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
    private acompanianteService: AcompanianteService
  ) {
    console.log(data.edad);
  }

  ngOnInit() {
    this.form = this.fb.group({
      motivo: [null, Validators.compose([Validators.required])],
      justificacion: [null, Validators.compose([Validators.required])],
      nombreAutoriza: [null, Validators.compose([Validators.required])],
      cargo: [null, Validators.compose([])],
      matricula: [null, Validators.compose([Validators.pattern("^([0-9])*$")])],
    });
    if (this.data.edad < 16) {
      this.form.get("motivo").setValue("Paciente menor de edad");
      this.justificaciones.shift();
    } else if (this.data.edad > 65) {
      this.form.get("motivo").setValue("Paciente mayor que no puede valerse por sí mismo");
      this.justificaciones.pop();
    } else {
      this.justificaciones.pop();
      this.justificaciones.shift();
    }
    this.consultarAcompaniante();
  }
  consultarAcompaniante() {
    this.acompanianteService.ConsultarUltimoAcompaniante(this.data.unidadMed).subscribe((result) => {
        this.form.get('nombreAutoriza').setValue(result.Data.nombreAutoriza);
        this.form.get('cargo').setValue(result.Data.cargo);
        this.form.get('matricula').setValue(result.Data.matricula);
      });
  }
  onNoClick(): void {
    this.dialogRef.close(false);
  }

  save() {
    this.datosAutorizacion = this.form.getRawValue();
    this.dialogRef.close({
      autorizacion: this.datosAutorizacion,
    });
  }
}
