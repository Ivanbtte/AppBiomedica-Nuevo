import { Component, OnInit } from "@angular/core";
import { Router } from "@angular/router";
import {
  FormBuilder,
  FormControl,
  FormGroup,
  Validators,
} from "@angular/forms";
import {
  DateAdapter,
  MAT_DATE_FORMATS,
  MAT_DATE_LOCALE,
} from "@angular/material/core";
import {
  MAT_MOMENT_DATE_FORMATS,
  MomentDateAdapter,
} from "@angular/material-moment-adapter";
import { MatDatepickerInputEvent } from "@angular/material/datepicker";
import { UsuarioService } from "../servicios/usuario.service";
import { RespuestaPeticion } from "../../core/models/estructuras_generales";
import Swal from "sweetalert2";
import { CustomValidators } from "ng2-validation";
import { Passwords } from "../interfaces/usuario.interface";

const password = new FormControl(
  null,
  Validators.compose([Validators.required])
);
const confirmPassword = new FormControl(
  null,
  Validators.compose([Validators.required, CustomValidators.equalTo(password)])
);

@Component({
  selector: "app-forgot",
  templateUrl: "./forgot.component.html",
  styleUrls: ["./forgot.component.scss"],
  providers: [
    { provide: MAT_DATE_LOCALE, useValue: "es-mx" },
    {
      provide: DateAdapter,
      useClass: MomentDateAdapter,
      deps: [MAT_DATE_LOCALE],
    },
    { provide: MAT_DATE_FORMATS, useValue: MAT_MOMENT_DATE_FORMATS },
  ],
})
export class ForgotComponent implements OnInit {
  formForgot: FormGroup;
  formPassword: FormGroup;
  dataSelected: string;
  confirm = false;
  hide = true;
  hide2 = true;
  minDate: Date;
  maxDate: Date;
  data: any;

  constructor(
    private fb: FormBuilder,
    private router: Router,
    private adapter: DateAdapter<any>,
    private usuarioService: UsuarioService
  ) {
    this.adapter.setLocale("es-mx");
    this.formForgot = this.fb.group({
      email: [
        null,
        Validators.compose([Validators.required, Validators.email]),
      ],
      fecha: [null, Validators.compose([Validators.required])],
    });
    this.formPassword = this.fb.group({
      password: password,
      confirmPassword: confirmPassword,
    });
    this.formPassword.reset();
    const currentYear = new Date().getFullYear();
    this.minDate = new Date(currentYear - 85, 0, 1);
    this.maxDate = new Date(currentYear - 15, 11, 31);
  }

  ngOnInit() {}

  onSubmit() {
    this.usuarioService
      .CheckResetPassword(this.formForgot.get("email").value, this.dataSelected)
      .subscribe(
        (result) => {
          console.log(result.Data, "consultando");
          if (result.Data["status"] === true) {
            this.confirm = result.Data["status"];
            this.data = result.Data["data"];
          }
        },
        (err) => {
          if (err.error !== "" || err.error !== undefined) {
            const errores = <RespuestaPeticion>err.error;
            console.log(errores.Data);
            Swal.fire({
              position: "center",
              title: "Error",
              text: errores.Data,
              icon: "error",
              showConfirmButton: true,
            });
          }
        }
      );
  }

  selectedDate(event: MatDatepickerInputEvent<Date>) {
    console.log("cambiando we ");
    this.dataSelected = event.targetElement["value"];
  }

  resetPassword() {
    console.log(this.data.Id, this.formPassword.get("password").value);
    const tmp: Partial<Passwords> = {
      NewPassword: this.formPassword.get("password").value,
    };
    console.log(tmp)
    Swal.fire({
      title: "Actualizar contraseña",
      text: "¿Estas seguro de actualizar la contraseña?",
      icon: "warning",
      showCancelButton: true,
      confirmButtonColor: "#3085d6",
      cancelButtonColor: "#d33",
      cancelButtonText: "Cancelar",
      confirmButtonText: "Si, actualizar",
    }).then((result) => {
      if (result.isConfirmed) {
        this.usuarioService.ResetPassword(this.data.Id, tmp).subscribe(
          (result) => {
            console.log(result);
            Swal.fire("Éxito", "Tu contraseña ha sido actualizada", "success");
            this.router.navigate(["/authentication/login"]);
          },
          (err) => {
            if (err.error !== "" || err.error !== undefined) {
              const errores = <RespuestaPeticion>err.error;
              console.log(errores.Data);
              Swal.fire({
                position: "center",
                title: "Error",
                text: errores.Data,
                icon: "error",
                showConfirmButton: true,
              });
            }
          }
        );
      }
    });
  }
}
