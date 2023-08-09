import {Component, OnInit} from '@angular/core';
import {Router} from '@angular/router';
import {DatosUsuario, Person, Personas, UnidadesM, Usuario} from '../interfaces/usuario.interface';
import {UsuarioService} from '../servicios/usuario.service';
import {FormBuilder, FormControl, FormGroup, Validators} from '@angular/forms';
import {CustomValidators} from 'ng2-validation';
import {ServiciosUnidadesService} from '../../core/servicios/servicios-unidades.service';
import {PuestosService} from '../servicios/puestos.service';
import {DelegacionesService} from '../../core/servicios/delegaciones.service';
import {DepartamentosService} from '../../core/servicios/DepartamentosSubdireccion/departamentos.service';
import {UsuarioTipoService} from '../../core/servicios/usuario-tipo.service';
import Swal from 'sweetalert2';
import {RespuestaPeticion} from '../../core/models/estructuras_generales';
import {AreasService} from '../../core/servicios/areas.service';
import {Area} from '../../core/models/areas';
import {Organigrama} from '../interfaces/departament.interface';
import {CompanyPosition} from '../interfaces/company.position.interface';
import {MatDialog} from '@angular/material/dialog';
import {DepartmentTreeComponent} from '../department-tree/department-tree.component';
import {CapitalizadoPipe} from '../../shared/pipes/capitalizado.pipe';
import {MatDatepickerInputEvent} from '@angular/material/datepicker';
import {DateAdapter, MAT_DATE_FORMATS, MAT_DATE_LOCALE} from '@angular/material/core';
import {MAT_MOMENT_DATE_FORMATS, MomentDateAdapter} from '@angular/material-moment-adapter';
import {PersonaService} from '../servicios/persona.service';

/* *****  constantes que validan las contraseñas   *****   */
const password = new FormControl('', Validators.compose([Validators.required]));
const confirmPassword = new FormControl('', Validators.compose([Validators.required, CustomValidators.equalTo(password)]));

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss'],
  providers: [
    {provide: MAT_DATE_LOCALE, useValue: 'es-mx'},
    {provide: DateAdapter, useClass: MomentDateAdapter, deps: [MAT_DATE_LOCALE]},
    {provide: MAT_DATE_FORMATS, useValue: MAT_MOMENT_DATE_FORMATS},
  ],
})
export class RegisterComponent implements OnInit {
  hide = true;
  hide2 = true;
  user: Usuario;
  estadoBoton = false;
  mensajeErrorCorreo = 'example@imss.gob.mx';
  botonVisible = true;
  UM: UnidadesM [] = [];
  departamentos: any[] = [];
  puestos: any[] = [];
  usuario_tipo: any[] = [];
  delegaciones: any[] = [];
  isLinear = true;
  firstFormGroup: FormGroup;
  secondFormGroup: FormGroup;
  thirdFormGroup: FormGroup;
  fourthFormGroup: FormGroup;
  departmentSelected = '';
  textTypeDelegations = '';
  areas: Area[] = [];
  positions: CompanyPosition[] = [];
  dataSelected: string;
  personResponse: Person;
  minDate: Date;
  maxDate: Date;

  constructor(private fb: FormBuilder,
              private router: Router,
              private _unidades: UsuarioService,
              private _serviciosUnidades: ServiciosUnidadesService,
              private _puestos: PuestosService,
              private _departamentos: DepartamentosService,
              private _usuario_tipo: UsuarioTipoService,
              private _delegaciones: DelegacionesService,
              private areasService: AreasService,
              private dialog: MatDialog,
              private adapter: DateAdapter<any>,
              private personaService: PersonaService,
  ) {
    this.adapter.setLocale('es-mx');
    const currentYear = new Date().getFullYear();
    this.minDate = new Date(currentYear - 85, 0, 1);
    this.maxDate = new Date(currentYear - 15, 11, 31);
  }

  ngOnInit() {
    this.createFirstForm();
    this.createSecondForm();
    this.createThirdForm();
    this.createFourthForm();
    this.consultUserType();
    this.consultAreas();
    this.secondFormGroup.statusChanges.subscribe(result => {
      if (result === 'VALID') {
        console.log('validado', result);
      }
    });
  }

  createFirstForm() {
    this.firstFormGroup = this.fb.group({
      Nombre: [null, Validators.compose([Validators.pattern('^[a-zA-ZáéíóúÁÉÍÓÚñÑ\\s]*$'),
        Validators.required])],
      ApellidoPat: [null, Validators.compose([Validators.pattern('^[a-zA-ZáéíóúÁÉÍÓÚñÑ\\s]*$'),
        Validators.required])],
      ApellidoMat: [null, Validators.compose([Validators.pattern('^[a-zA-ZáéíóúÁÉÍÓÚñÑ\\s]*$')])],
      FechaNac: [{value: null, disabled: true}, Validators.compose([Validators.required])],
      Genero: [null, Validators.compose([Validators.required])],
      Correo: [null, Validators.compose([Validators.required, Validators.pattern('^[_a-z0-9-]+(.[_a-z0-9-]+)*@imss.gob.mx$')])
      ],
      Telefono: [null, Validators.compose([Validators.pattern('^([0-9])*$')])],
      Extension: [null, Validators.compose([Validators.pattern('^([0-9])*$')])],
      Matricula: [null, Validators.compose([Validators.pattern('^([0-9])*$'), Validators.required])],
      //         MyValidators.validateEnrollment(this.personaService)
      Celular: [null, Validators.compose([Validators.pattern('^([0-9])*$')])]
    }, {updateOn: 'blur'});
  }

  get enrollmentField() {
    return this.firstFormGroup.get('Matricula');
  }

  getErrorMessage() {
    if (this.enrollmentField.hasError('pattern')) {
      return 'Matricula incorrecta.';
    }
  }

  blurEvent() {
    console.log(this.enrollmentField.valid);
    if (this.enrollmentField.valid) {
      this.personaService.checkEnrollment(this.enrollmentField.value).subscribe(result => {
        const tempValue = this.enrollmentField.value;
        this.firstFormGroup.reset();
        console.log(result.Data);
        this.personResponse = result.Data;
        if (!this.personResponse.status && this.personResponse.personStatus) {
          this.enrollmentField.setErrors({'invalid': true});
          Swal.fire({
            title: '<strong style="background-color: #e8e9e9; color: #b8316f">Esta matrícula ya esta registrada </strong>',
            text: 'Si crees que alguien ha usurpado tu identidad repórtalo al administrador del sistema',
            icon: 'error',
            showCancelButton: true,
            showConfirmButton: false,
            cancelButtonColor: '#d33',
            cancelButtonText: 'Cerrar'
          }).then((result) => {
            if (result.isConfirmed) {
              Swal.fire(
                'Deleted!',
                'Your file has been deleted.',
                'success'
              );
            }
          });
        } else {
          this.firstFormGroup.patchValue(this.personResponse.data);
        }
        this.enrollmentField.setValue(tempValue);
      });
    }
  }

  createSecondForm() {
    this.secondFormGroup = this.fb.group({
      tipo_usuario: ['', Validators.required],
      delegaciones: ['', Validators.required],
      unidad_med: ['', Validators.required],
    });
  }

  createThirdForm() {
    this.thirdFormGroup = this.fb.group({
      area: [{value: '', disabled: true}, Validators.required],
      departamentos: ['', Validators.required],
      puesto: ['', Validators.required],
    });
  }

  createFourthForm() {
    this.fourthFormGroup = this.fb.group({
      password: password,
      confirmPassword: confirmPassword
    });
  }

  consultUserType() {
    this._usuario_tipo.checkUserType().subscribe(res => {
      this.usuario_tipo = res.Data;
    }, err => {
      console.log(err);
    });
  }

  consultAreas() {
    this.areasService.consultAreas().subscribe(result => {
      this.areas = result.Data;
    });
  }

  selectedUserType(value: string) {
    switch (value) {
      case '1':
        this.userTypeSelectedNormative();
        break;
      case '2':
        this.userTypeSelectedDelegation();
        break;
      case '3':
        this.userTypeSelectedUMAE();
        break;
      case '4':
        this.userTypeSelectedUnitMedic();
        break;
    }
    this.cleanFormUserType();
  }

  userTypeSelectedNormative() {
    this.delegaciones = [];
    this.textTypeDelegations = '';
    this.secondFormGroup.get('delegaciones').clearValidators();
    this.secondFormGroup.get('unidad_med').clearValidators();
  }

  userTypeSelectedDelegation() {
    this.consultDelegations('1');
    this.textTypeDelegations = 'Delegacion';
    this.secondFormGroup.get('delegaciones').setValidators([Validators.required]);
    this.secondFormGroup.get('unidad_med').clearValidators();
  }

  userTypeSelectedUMAE() {
    this.consultDelegations('2');
    this.textTypeDelegations = 'UMAE';
    this.secondFormGroup.get('unidad_med').setValidators([Validators.required]);
  }

  userTypeSelectedUnitMedic() {
    this.consultDelegations('1');
    this.textTypeDelegations = 'Delegacion';
    this.secondFormGroup.get('unidad_med').setValidators([Validators.required]);
  }

  consultDelegations(value: string) {
    this._delegaciones.consultDelegations(value).subscribe(res => {
      this.delegaciones = res.Data;
    }, err => {
      console.log(err);
    });
  }

  cleanFormUserType() {
    this.clearSecondForm();
    this.clearThirdForm();
    this.thirdFormGroup.get('area').reset();
  }

  clearSecondForm() {
    this.secondFormGroup.get('delegaciones').reset();
    this.delegaciones = [];
    this.secondFormGroup.get('unidad_med').reset();
    this.UM = [];
  }

  clearThirdForm() {
    this.thirdFormGroup.get('departamentos').reset();
    this.thirdFormGroup.get('puesto').reset();
    this.positions = [];
  }

  selectedDelegation(value: string) {
    const typeUser = this.secondFormGroup.get('tipo_usuario').value;
    if (typeUser !== '2') {
      this._unidades.unidades(this.secondFormGroup.controls['delegaciones'].value.ClvDele).subscribe(res => {
        this.UM = res.Data;
      }, err => {
        console.log(err);
      });
    } else {
      this.clearFormUnitMedic();
    }
  }

  clearFormUnitMedic() {
    this.secondFormGroup.controls['unidad_med'].reset();
    this.UM = [];
  }


  openDialogDepartment() {
    let clvDele = '';
    let Dele = '';
    let unitId = 0;
    let unit = '';
    const userType = this.secondFormGroup.get('tipo_usuario').value !== null ? this.secondFormGroup.get('tipo_usuario').value : '';
    if (userType === '2') {
      clvDele = this.secondFormGroup.get('delegaciones').value.ClvDele;
      Dele = this.secondFormGroup.get('delegaciones').value;
    }
    if (userType === '3' || userType === '4') {
      unitId = this.secondFormGroup.get('unidad_med').value.Id;
      unit = this.secondFormGroup.get('unidad_med').value;
    }
    const dialogRef = this.dialog.open(DepartmentTreeComponent, {
      height: 'auto',
      width: '60%',
      disableClose: true,
      data: {
        unitMedId: unitId,
        UnitMed: unit,
        clvDelegation: clvDele,
        Delegation: Dele,
      }
    });
    dialogRef.afterClosed().subscribe(result => {
      if (result !== '') {
        this.thirdFormGroup.get('area').setValue(result.areaId);
        this.thirdFormGroup.get('departamentos').setValue(result.name);
        this.departmentSelected = result.name;
        if (userType === '2') {
          this.consultPositionDepartment('', result.id);
        }
        if (userType === '3' || userType === '4') {
          this.consultPositionDepartment(result.id, '');
        }
        this.thirdFormGroup.get('puesto').reset();
        this.positions = [];
      }
    });
  }

  consultPositionDepartment(departmentUnitId, departmentDelegationId: string) {
    this._puestos.ConsultPositionOrganizations(departmentUnitId, departmentDelegationId).subscribe(result => {
      this.positions = result.Data;
    });
  }

  addNewUser() {
    console.log(this.personResponse);
    const datos = <Personas> this.firstFormGroup.getRawValue();
    datos.Nombre = this.capitalizeText(datos.Nombre);
    datos.ApellidoPat = this.capitalizeText(datos.ApellidoPat);
    datos.ApellidoMat = this.capitalizeText(datos.ApellidoMat);
    datos.Extension = parseInt(this.firstFormGroup.controls['Extension'].value, 10);
    datos.FechaNac = this.dataSelected;
    datos.Id = '';
    this.user = {
      Correo: this.firstFormGroup.controls['Correo'].value,
      Contraseña: this.fourthFormGroup.controls['password'].value,
      UsuarioTipoId: this.secondFormGroup.controls['tipo_usuario'].value,
      PersonaId: ''
    };
    if (!this.personResponse.status && !this.personResponse.personStatus) {
      console.log('Es una reactivacion');
      datos.Id = this.personResponse.data.Id;
      this.user.PersonaId = this.personResponse.data.Id;
      datos.FechaNac = this.personResponse.data.FechaNac;
      datos.Estado = true;
    }
    const obj: DatosUsuario = {
      Persona: datos,
      Usuario: this.user,
      PuestoOrganigramaId: this.thirdFormGroup.get('puesto').value.id
    };
    let message = '';
    if (this.user.UsuarioTipoId === '2') {
      message = '<b>Delegacion: </b>' + this.secondFormGroup.get('delegaciones').value.NombreDele + '<br>';
    } else {
      message = '<b>Unidad Medica: </b>' + this.secondFormGroup.get('unidad_med').value.DenominacionUni + '<br>';
    }
    const name = datos.Nombre + ' ' + datos.ApellidoPat + ' ' + datos.ApellidoMat;
    const department = this.thirdFormGroup.get('departamentos').value;
    const position = this.thirdFormGroup.get('puesto').value.catalogoPuesto.Puesto;
    const enrollment = this.firstFormGroup.get('Matricula').value;
    const correo = this.firstFormGroup.get('Correo').value;
    console.log(obj);
    Swal.fire({
      title: '¿Confirmar si los datos de registro son correctos?',
      icon: 'warning',
      html: '<div align="left">' +
        '<b>Nombre: </b>' + name + '<br>' +
        '<b>Matricula: </b>' + enrollment + '<br>' +
        '<b>Correo: </b>' + correo + '<br>' +
        message +
        '<b>Departamento: </b>' + department + '<br>' +
        '<b>Puesto: </b>' + position + '<br>'
        + '</div>',
      showCancelButton: true,
      confirmButtonColor: '#3ba941',
      cancelButtonColor: '#d53147',
      confirmButtonText: 'Si, Guardar',
      cancelButtonText: 'No, Editar',
      reverseButtons: true
    }).then((result) => {
      if (result.isConfirmed) {
        this._unidades.AddNewUser(obj).subscribe(res => {
          Swal.fire({
            icon: 'success',
            position: 'center',
            title: res.Mensaje,
            showConfirmButton: false,
            timer: 1000
          });
          this.router.navigate(['authentication/login']);
        }, err => {
          if (err.error !== '' || err.error !== undefined) {
            const erores = <RespuestaPeticion> err.error;
            console.log(erores.Data);
            Swal.fire({
              position: 'center',
              title: erores.Mensaje,
              icon: 'error',
              showConfirmButton: false,
              timer: 1000
            });
          }
          console.log(err);
        });
      }
    });
  }

  cambiarValidaCorreo() {
    this.firstFormGroup.removeControl('Correo');
    this.firstFormGroup.addControl('Correo', this.fb.control('', Validators.compose([Validators.required, Validators.email])));
    this.mensajeErrorCorreo = 'Correo invalido';
    this.botonVisible = false;
    this.estadoBoton = true;
  }

  exit() {
    Swal.fire({
      title: '¿Estas seguro de salir?',
      text: 'Perdera toda la informacion registrada!',
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3ba941',
      cancelButtonColor: '#d53147',
      confirmButtonText: 'Si, salir',
      cancelButtonText: 'No, continuar registro',
      reverseButtons: true
    }).then((result) => {
      if (result.isConfirmed) {
        this.router.navigate(['authentication/login']);
      }
    });
  }

  capitalizeText(value: string) {
    const returnTypeFirmPipe = new CapitalizadoPipe();
    return returnTypeFirmPipe.transform(value);
  }

  selectedDate(event: MatDatepickerInputEvent<Date>) {
    this.dataSelected = event.targetElement['value'];
  }
}
