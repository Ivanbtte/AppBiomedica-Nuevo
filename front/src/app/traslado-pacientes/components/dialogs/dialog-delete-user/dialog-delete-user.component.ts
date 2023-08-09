import { Component, OnInit } from "@angular/core";
import { ServicioUsuarioPuestoService } from "../../../../core/servicios/ServiciosUsuarios/servicio-usuario-puesto.service";
import { FormGroup, FormBuilder, Validators } from "@angular/forms";
import { MatDialogRef } from "@angular/material/dialog";
import Swal from "sweetalert2";

@Component({
  selector: "app-dialog-delete-user",
  templateUrl: "./dialog-delete-user.component.html",
  styleUrls: ["./dialog-delete-user.component.css"],
})
export class DialogDeleteUserComponent implements OnInit {
  public form: FormGroup;
  tiposModificaciones: any[];
  constructor(
    private usuarioPuesto: ServicioUsuarioPuestoService,
    private fb: FormBuilder,
    public dialogRef: MatDialogRef<DialogDeleteUserComponent>
  ) {}

  ngOnInit(): void {
    this.consultTipoModificaciones();
    this.form = this.fb.group({
      tipoModificacion: [null, Validators.compose([Validators.required])],
      justificacion: [
        null,
        Validators.compose([Validators.required, Validators.minLength(20)]),
      ],
    });
  }

  consultTipoModificaciones() {
    this.usuarioPuesto
      .GetTipoModificaciones("61f1a754d7be0a2b052e887e")
      .subscribe((results) => {
        console.log(results.Data);
        this.tiposModificaciones = results.Data;
      });
  }
  onSubmit() {
    Swal.fire({
      title: "Confirmación de baja",
      text: "¿Estas seguro de realizar esta baja?",
      icon: "warning",
      showCancelButton: true,
      confirmButtonColor: "#3085d6",
      cancelButtonColor: "#d33",
      confirmButtonText: "Si, dar de baja",
      cancelButtonText: "Cancelar",
      reverseButtons: true,
    }).then((result) => {
      if (result.isConfirmed) {
        this.dialogRef.close(this.form.value);
      }
    });
  }
}
