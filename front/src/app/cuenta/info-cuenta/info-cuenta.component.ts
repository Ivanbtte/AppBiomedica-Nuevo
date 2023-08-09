import { Component, OnInit } from "@angular/core";
import jwt_decode from "jwt-decode";
import { UsuarioLogeado } from "../../core/models/UsuarioLogin.interface";
import {
  FormBuilder,
  FormControl,
  FormGroup,
  Validators,
} from "@angular/forms";
import {
  Passwords,
  Personas,
  UnidadesM,
} from "../../authentication/interfaces/usuario.interface";
import { UsuarioService } from "../../authentication/servicios/usuario.service";
import { CustomValidators } from "ng2-validation";
import Swal from "sweetalert2";
import { PersonaService } from "../../authentication/servicios/persona.service";
import { Observable } from "rxjs";
import { RespuestaPeticion } from "../../core/models/estructuras_generales";
import { ActivatedRoute, Router } from "@angular/router";
import { MatDatepickerInputEvent } from "@angular/material/datepicker";
import {
  DateAdapter,
  MAT_DATE_FORMATS,
  MAT_DATE_LOCALE,
} from "@angular/material/core";
import {
  MAT_MOMENT_DATE_FORMATS,
  MomentDateAdapter,
} from "@angular/material-moment-adapter";
import * as _moment from "moment";
import "rxjs/add/operator/filter";
import { UsuarioTipoService } from "../../core/servicios/usuario-tipo.service";
import { UsuarioTipo } from "../../core/models/user.type.model";
import { Delegaciones } from "../../core/models/delegation.model";
import { DelegacionesService } from "../../core/servicios/delegaciones.service";
import { Area } from "../../core/models/areas";
import { AreasService } from "../../core/servicios/areas.service";
import { CompanyPosition } from "../../authentication/interfaces/company.position.interface";
import { DepartamentosService } from "../../core/servicios/DepartamentosSubdireccion/departamentos.service";
import { DepartmentTreeComponent } from "../../authentication/department-tree/department-tree.component";
import { MatDialog } from "@angular/material/dialog";
import { PuestosService } from "../../authentication/servicios/puestos.service";
import { UserPositionsService } from "../../core/servicios/userPositions/user-positions.service";
import { U } from "@angular/cdk/keycodes";

const moment = _moment;
const password = new FormControl("", Validators.compose([Validators.required]));
const passwordNew = new FormControl(
  "",
  Validators.compose([Validators.required])
);
const confirmPassword = new FormControl(
  "",
  Validators.compose([
    Validators.required,
    CustomValidators.equalTo(passwordNew),
  ])
);

@Component({
  selector: "app-info-cuenta",
  templateUrl: "./info-cuenta.component.html",
  styleUrls: ["./info-cuenta.component.css"],
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
export class InfoCuentaComponent implements OnInit {
  hide = true;
  hide2 = true;
  hide3 = true;
  firstFormGroup: FormGroup;
  form: FormGroup;
  positionForm: FormGroup;
  areaForm: FormGroup;
  usuarioLog: UsuarioLogeado;
  persona: Personas;
  masculino = false;
  femenino = false;
  otro = false;
  userPosition: any;
  dataSelected: string;
  minDate: Date;
  maxDate: Date;
  position = "";
  indexSelected = 0;
  userTypes: UsuarioTipo[] = [];
  patternString = "^[a-zA-ZáéíóúÁÉÍÓÚñÑ]+(\\s[a-zA-Z_áéíóúÁÉÍÓÚñÑ]+)*$";
  patternNumber = "^([0-9])*$";
  textTypeDelegations = "";
  delegations: Delegaciones[] = [];
  unitMedics: UnidadesM[] = [];
  areas: Area[] = [];
  positions: CompanyPosition[] = [];
  editPosition = false;
  edit = false;
  departmentSelected = "";
  userTypeUnit = false;
  department = "";
  editDataPerson = false;

  constructor(
    private fb: FormBuilder,
    private personaService: PersonaService,
    private router: Router,
    private usuarioService: UsuarioService,
    private adapter: DateAdapter<any>,
    private route: ActivatedRoute,
    private usuarioTipoService: UsuarioTipoService,
    private delegacionesService: DelegacionesService,
    private areasService: AreasService,
    private departamentosService: DepartamentosService,
    private dialog: MatDialog,
    private puestosService: PuestosService,
    private userPositionsService: UserPositionsService
  ) {
    this.adapter.setLocale("es-mx");
    const currentYear = new Date().getFullYear();
    this.minDate = new Date(currentYear - 85, 0, 1);
    this.maxDate = new Date(currentYear - 15, 11, 31);
  }

  ngOnInit() {
    this.checkParams();
    this.usuarioLog = jwt_decode(localStorage.getItem("token"));
    console.log(this.usuarioLog);
    this.consultUserPosition(this.usuarioLog.UsuarioId);
    this.createFormDataPerson();
    this.createFormDataUser();
  }

  checkParams() {
    this.route.queryParams.subscribe((params) => {
      if (params["index"] !== undefined) {
        this.indexSelected = params["index"];
        this.editPosition = true;
        this.consultUserType();
        this.createFormPosition();
        this.createFormArea();
        this.consultAreas();
      }
    });
  }

  createFormPosition() {
    this.positionForm = this.fb.group({
      userType: ["", Validators.required],
      delegation: ["", Validators.required],
      unitMedic: ["", Validators.required],
    });
  }

  createFormArea() {
    this.areaForm = this.fb.group({
      area: [{ value: "", disabled: true }, Validators.required],
      department: ["", Validators.required],
      position: ["", Validators.required],
    });
  }

  consultUserType() {
    this.usuarioTipoService.checkUserType().subscribe(
      (res) => {
        this.userTypes = res.Data;
        console.log(this.userTypes);
      },
      (err) => {
        console.log(err);
      }
    );
  }

  consultUserPosition(id: string) {
    this.usuarioService.ConsultUSerByPosition(id).subscribe((result) => {
      this.userPosition = result.Data;
      console.log(result.Data, "we");
      if (this.userPosition.Estado === false) {
        this.position = "Sin puesto";
        this.editPosition = true;
        this.indexSelected = 1;
        this.consultUserType();
        this.createFormPosition();
        this.createFormArea();
        this.consultAreas();
      } else {
        this.position =
          this.userPosition.PuestoOrganigrama.catalogoPuesto.Puesto;
        this.userTypeUnit =
          this.userPosition.Usuario.UsuarioTipoId === "3" ||
          this.userPosition.Usuario.UsuarioTipoId === "4";
        if (this.userTypeUnit) {
          this.department =
            this.userPosition.PuestoOrganigrama.organigramaUnidad.departamento.Nombre;
        } else {
          this.department =
            this.userPosition.PuestoOrganigrama.organigramaDelegacion.departamento.Nombre;
        }
      }
      this.setDataForm();
      if (this.userPosition.Usuario.Persona.Genero === "M") {
        this.masculino = true;
      }
      if (this.userPosition.Usuario.Persona.Genero === "F") {
        this.femenino = true;
      }
      if (this.userPosition.Usuario.Persona.Genero === "O") {
        this.otro = true;
      }
    });
  }

  createFormDataPerson() {
    this.firstFormGroup = this.fb.group({
      Nombre: [
        null,
        Validators.compose([
          Validators.pattern(this.patternString),
          Validators.required,
        ]),
      ],
      ApellidoPat: [
        null,
        Validators.compose([
          Validators.pattern(this.patternString),
          Validators.required,
        ]),
      ],
      ApellidoMat: [
        null,
        Validators.compose([
          Validators.pattern(this.patternString),
          Validators.required,
        ]),
      ],
      FechaNac: [null, Validators.compose([Validators.required])],
      Genero: [null, Validators.compose([Validators.required])],
      Telefono: [
        null,
        Validators.compose([
          Validators.pattern(this.patternNumber),
          Validators.minLength(10),
          Validators.maxLength(10),
        ]),
      ],
      Extension: [
        null,
        Validators.compose([Validators.pattern(this.patternNumber)]),
      ],
      Matricula: [
        null,
        Validators.compose([Validators.pattern(this.patternNumber)]),
      ],
      Celular: [
        null,
        Validators.compose([Validators.pattern(this.patternNumber)]),
      ],
    });
  }

  enableFirsForm() {
    this.editDataPerson = true;
  }

  cancelEditionOfPersonalData() {
    this.editDataPerson = false;
    this.createFormDataPerson();
    this.setDataForm();
  }
  enableEditPosition() {
    this.edit = true;
    console.log("editarrr", this.edit);
    this.editPosition = true;
    this.indexSelected = 1;
    this.consultUserType();
    this.createFormPosition();
    this.createFormArea();
    this.consultAreas();
  }

  cancelEditPositionForm() {
    this.editPosition = false;
    this.positionForm.reset();
    this.areaForm.reset();
    this.cleanFormArea();
    this.cleanFormUserType();
  }

  createFormDataUser() {
    this.form = this.fb.group({
      Password: password,
      NewPassword: passwordNew,
      confirmPassword: confirmPassword,
    });
  }

  selectedUserType(value: string) {
    switch (value) {
      case "1":
        this.userTypeSelectedNormative();
        break;
      case "2":
        this.userTypeSelectedDelegation();
        break;
      case "3":
        this.userTypeSelectedUMAE();
        break;
      case "4":
        this.userTypeSelectedUnitMedic();
        break;
    }
    this.cleanFormUserType();
  }

  cleanFormUserType() {
    this.positionForm.get("delegation").reset();
    this.delegations = [];
    this.positionForm.get("unitMedic").reset();
    this.unitMedics = [];
    this.cleanFormArea();
  }

  cleanFormArea() {
    this.areaForm.get("department").reset();
    this.departmentSelected = "";
    this.areaForm.get("position").reset();
    this.positions = [];
  }

  userTypeSelectedNormative() {
    this.delegations = [];
    this.textTypeDelegations = "";
    this.positionForm.get("delegation").clearValidators();
    this.positionForm.get("unitMedic").clearValidators();
  }

  userTypeSelectedDelegation() {
    this.consultDelegations("1");
    this.textTypeDelegations = "Delegación:";
    this.positionForm.get("delegation").setValidators([Validators.required]);
    this.positionForm.get("unitMedic").clearValidators();
  }

  userTypeSelectedUMAE() {
    this.consultDelegations("2");
    this.textTypeDelegations = "UMAE:";
    this.positionForm.get("unitMedic").setValidators([Validators.required]);
  }

  userTypeSelectedUnitMedic() {
    this.consultDelegations("1");
    this.textTypeDelegations = "Delegación:";
    this.positionForm.get("unitMedic").setValidators([Validators.required]);
  }

  consultDelegations(value: string) {
    this.delegacionesService.consultDelegations(value).subscribe(
      (res) => {
        this.delegations = res.Data;
      },
      (err) => {
        console.log(err);
      }
    );
  }

  selectedDelegation(value: Delegaciones) {
    const typeUser = this.positionForm.get("userType").value;
    if (typeUser !== "2") {
      this.usuarioService.unidades(value.ClvDele).subscribe(
        (res) => {
          this.unitMedics = res.Data;
        },
        (err) => {
          console.log(err);
        }
      );
    } else {
      this.clearFormUnitMedic();
    }
  }

  clearFormUnitMedic() {
    this.positionForm.controls["unitMedic"].reset();
    this.unitMedics = [];
    this.cleanFormArea();
  }

  changeSelectedUnitMedic() {
    this.cleanFormArea();
  }

  setDataForm() {
    this.dataSelected = this.userPosition.Usuario.Persona.FechaNac;
    this.firstFormGroup.setValue({
      Nombre: this.userPosition.Usuario.Persona.Nombre,
      ApellidoPat: this.userPosition.Usuario.Persona.ApellidoPat,
      ApellidoMat: this.userPosition.Usuario.Persona.ApellidoMat,
      FechaNac: this.formatDate(this.userPosition.Usuario.Persona.FechaNac),
      Genero: this.userPosition.Usuario.Persona.Genero,
      Telefono: this.userPosition.Usuario.Persona.Telefono,
      Extension: this.userPosition.Usuario.Persona.Extension,
      Matricula: this.userPosition.Usuario.Persona.Matricula,
      Celular: this.userPosition.Usuario.Persona.Celular,
    });
  }

  formatDate(value: string): Date {
    const day = moment(value, "DD/MM/YYYY");
    return day.toDate();
  }

  actualizarPersona() {
    const datos = <Personas>this.firstFormGroup.getRawValue();
    datos.Correo = this.usuarioLog.Correo;
    datos.FechaNac = this.dataSelected;
    Swal.fire({
      title: "Estás seguro de actualizar",
      text: "Puedes cancelar",
      icon: "warning",
      showCancelButton: true,
      confirmButtonColor: "#3085d6",
      cancelButtonColor: "#d33",
      confirmButtonText: "Si Actualizar",
    }).then((result) => {
      if (result.value) {
        let peticion: Observable<RespuestaPeticion>;
        peticion = this.personaService.putPersona(
          datos,
          this.usuarioLog.PersonaId
        );
        peticion.subscribe(
          (res) => {
            Swal.fire("Actualizado", res.Mensaje, "success");
            this.consultUserPosition(this.usuarioLog.UsuarioId);
            this.editDataPerson = false;
          },
          (err) => {
            if (err.error !== "" || err.error !== undefined) {
              const erores = <RespuestaPeticion>err.error;
              console.log(erores);
              Swal.fire({
                position: "center",
                title: erores.Mensaje,
                icon: "error",
                showConfirmButton: true,
              });
            }
          }
        );
      }
    });
  }

  actualizarUsuario() {
    const datos = <Passwords>this.form.getRawValue();
    Swal.fire({
      title: "Estás seguro de actualizar",
      text: "No quiero actualizar",
      icon: "warning",
      showCancelButton: true,
      confirmButtonColor: "#3085d6",
      cancelButtonColor: "#d33",
      confirmButtonText: "Si Actualizar!",
    }).then((result) => {
      if (result.value) {
        let peticion: Observable<RespuestaPeticion>;
        peticion = this.usuarioService.ActualizarUsuario(
          this.usuarioLog.UsuarioId,
          datos
        );
        peticion.subscribe(
          (res) => {
            Swal.fire("Actualizado!", res.Mensaje, "success");
            this.router.navigate(["authentication/login"]);
          },
          (err) => {
            if (err.error !== "" || err.error !== undefined) {
              const erores = <RespuestaPeticion>err.error;
              Swal.fire({
                position: "center",
                title: erores.Mensaje,
                icon: "error",
                showConfirmButton: false,
                timer: 1000,
              });
            }
          }
        );
      }
    });
  }

  selectedDate(event: MatDatepickerInputEvent<Date>) {
    this.dataSelected = event.targetElement["value"];
  }

  consultAreas() {
    this.areasService.consultAreas().subscribe((result) => {
      this.areas = result.Data;
    });
  }

  openDialogDepartment() {
    let clvDele = "";
    let Dele = "";
    let unitId = 0;
    let unit = "";
    const userType =
      this.positionForm.get("userType").value !== null
        ? this.positionForm.get("userType").value
        : "";
    if (userType === "2") {
      clvDele = this.positionForm.get("delegation").value.ClvDele;
      Dele = this.positionForm.get("delegation").value;
    }
    if (userType === "3" || userType === "4") {
      unitId = this.positionForm.get("unitMedic").value.Id;
      unit = this.positionForm.get("unitMedic").value;
    }
    const dialogRef = this.dialog.open(DepartmentTreeComponent, {
      height: "auto",
      width: "60%",
      disableClose: true,
      data: {
        unitMedId: unitId,
        UnitMed: unit,
        clvDelegation: clvDele,
        Delegation: Dele,
      },
    });
    dialogRef.afterClosed().subscribe((result) => {
      if (result !== "") {
        this.areaForm.get("area").setValue(result.areaId);
        this.areaForm.get("department").setValue(result.name);
        this.departmentSelected = result.name;
        if (userType === "2") {
          this.consultPositionDepartment("", result.id);
        }
        if (userType === "3" || userType === "4") {
          this.consultPositionDepartment(result.id, "");
        }
        this.areaForm.get("position").reset();
        this.positions = [];
      }
    });
  }

  consultPositionDepartment(departmentUnitId, departmentDelegationId: string) {
    this.puestosService
      .ConsultPositionOrganizations(departmentUnitId, departmentDelegationId)
      .subscribe((result) => {
        this.positions = result.Data;
      });
  }

  addNewUserPosition() {
    const userType = this.positionForm.get("userType").value;
    console.log(this.positionForm.value);
    console.log(this.areaForm.value);
    let update = ''
    if (this.edit) {
      update = 'si'
    }
    const userPosition = {
      UsuarioId: this.userPosition.UsuarioId,
      PuestoOrganigramaId: this.areaForm.get("position").value.id,
    };
    let message = "";
    if (this.positionForm.get("userType").value === "2") {
      message =
        "<b>Delegacion: </b>" +
        this.positionForm.get("delegation").value.NombreDele +
        "<br>";
    } else {
      message =
        "<b>Unidad Medica: </b>" +
        this.positionForm.get("unitMedic").value.DenominacionUni +
        "<br>";
    }
    console.log(userPosition);
    /* ******  swal  ******   */
    Swal.fire({
      title: "¿Confirmar si los datos del puesto son correctos?",
      icon: "warning",
      html: `<div align="left">
             ${message}
             <b>Departamento: </b>  ${
               this.areaForm.get("department").value
             } <br>
             <b>Puesto: </b> ${
               this.areaForm.get("position").value.catalogoPuesto.Puesto
             } <br>
             </div>`,
      showCancelButton: true,
      confirmButtonColor: "#3ba941",
      cancelButtonColor: "#d53147",
      confirmButtonText: "Si, Guardar",
      cancelButtonText: "No, Editar",
      reverseButtons: true,
    }).then((result) => {
      if (result.isConfirmed) {
        /* ******  servicio  ******   */
        this.userPositionsService
          .addNewUserPosition(userPosition, userType,update)
          .subscribe(
            (result) => {
              /* ******  actualizar el menu y permisos  ******   */
              const nameModule = result.Data["permits"].map(
                (item) => item.modulo.name
              );
              localStorage.setItem("permisos", JSON.stringify(nameModule));
              localStorage.setItem("menu", JSON.stringify(result.Data["menu"]));
              console.log(result.Data);
              Swal.fire({
                title: "Se agrego correctamente el puesto",
                text: "Deberas iniciar sesion para aplicar los cambios.",
                icon: "success",
                showConfirmButton: true,
              }).then((result) => {
                if (result.isConfirmed) {
                  this.router.navigate(["authentication/login"]);
                }
              });
            },
            (err) => {
              if (err.error !== "" || err.error !== undefined) {
                const errores = <RespuestaPeticion>err.error;
                Swal.fire({
                  icon: "error",
                  title: "Ocurrió un error",
                  text: errores.Mensaje,
                });
              }
            }
          );
      }
    });
  }
}
