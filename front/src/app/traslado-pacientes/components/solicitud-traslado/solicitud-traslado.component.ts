import { Component, OnInit, ViewEncapsulation, ViewChild, ElementRef, AfterViewInit } from "@angular/core";
import { FtpService } from "../../../core/servicios/ServiciosSubrogados/traslados/ftp.service";
import { Ftp } from "../../../core/models/ServiciosSubrogados/trasladoPacientes/ftp.interface";
import { FormBuilder, FormControl, FormGroup, Validators } from "@angular/forms";
import { RespuestaPeticion } from "../../../core/models/estructuras_generales";
import Swal from "sweetalert2";
import { Acompaniante, Boletos, Solicitudtraslado, AutorizacionesEspeciales } from '../../../core/models/ServiciosSubrogados/trasladoPacientes/solicitud_traslado.interface';
import { UsuarioLogeado } from "../../../core/models/UsuarioLogin.interface";
import jwt_decode from "jwt-decode";
import { DateAdapter, MAT_DATE_FORMATS, MAT_DATE_LOCALE } from "@angular/material/core";
import { MAT_MOMENT_DATE_FORMATS, MomentDateAdapter } from "@angular/material-moment-adapter";
import { MatDialog } from "@angular/material/dialog";
import { MatRadioChange } from "@angular/material/radio";
import * as _moment from "moment";
import { Moment } from "moment";
import { ContratoService } from "../../../core/servicios/ContratosServicios/contrato.service";
import { Contrato } from "../../../core/models/contrato.interface";
import { PatientTransferRoutesService } from "../../../core/servicios/ServiciosSubrogados/traslados/patient-transfer-routes.service";
import { RepresentanteLegalService } from "../../../core/servicios/ContratosServicios/representante-legal.service";
import { TrasladoPacientes } from "../../../core/models/ServiciosSubrogados/trasladoPacientes/traslado_pacientes.interface";
import { DialogAutorizaAcompComponent } from "../dialogs/dialog-autoriza-acomp/dialog-autoriza-acomp.component";
import { VistaPreviaSolicitudComponent } from "../vista-previa-solicitud/vista-previa-solicitud.component";
import { DialogExcepcionesComponent, Excepciones } from "../dialogs/dialog-excepciones/dialog-excepciones.component";
import { SolicitudTrasladoService } from "../../../core/servicios/SolicitudTraslados/solicitud-traslado.service";
import { FirmService } from "../../../core/servicios/ServiciosSubrogados/traslados/firm.service";
import { FirmaSolicitud } from "../../../core/models/ServiciosSubrogados/trasladoPacientes/firma_solicitud.interface";
import { ModulesService } from '../../core/modules.service';
import { AccederService } from '../../core/acceder.service';
const moment = _moment;
//tipoBoleto --- 0.- Sin boleto, 1.-Entero, 2.-Medio
@Component({
  selector: "app-solicitud-traslado",
  templateUrl: "./solicitud-traslado.component.html",
  styleUrls: ["./solicitud-traslado.component.css"],
  providers: [
    { provide: MAT_DATE_LOCALE, useValue: "en-GB" },
    {
      provide: DateAdapter, useClass: MomentDateAdapter,
      deps: [MAT_DATE_LOCALE],
    },
    { provide: MAT_DATE_FORMATS, useValue: MAT_MOMENT_DATE_FORMATS },
  ],
  encapsulation: ViewEncapsulation.None,
})
export class SolicitudTrasladoComponent implements OnInit, AfterViewInit {
  @ViewChild("folionu", { static: false, read: ElementRef }) input1: ElementRef;
  folio!: Ftp;
  solicitudTraslado!: Solicitudtraslado;
  formSolicitud!: FormGroup;
  usuarioLog!: UsuarioLogeado;
  today = new Date();
  year = this.today.getFullYear();
  tipo_paciente = "";
  tipo_boleto = "";
  edad!: number;
  fechaCita = "";
  validarFecha!: Moment;
  pasajeros: Boletos[] = [];
  contratos: Contrato[] = [];
  destinos: any[] = [];
  origenes: TrasladoPacientes[] = [];
  unidadMedica: any;
  especialidad = "";
  preguntarFechaNac = false;
  acompaniante = false;
  control_acompaniante = new FormControl();
  controlIda = new FormControl(true);
  controlRegreso = new FormControl(true);
  pacienteInapan = false;
  tipoBoleto = 2;
  observacionesAcompa = "";
  observacionesPac = "";
  exepcion: Excepciones | undefined;
  acompanianteExtra = false;
  recoleccionMed = 0;
  autorizacionAcompaniante!: Acompaniante;
  imprimirPlantilla = true;
  ValidarFirmas = false;
  controlVoBo = new FormControl("", [Validators.required]);
  firmsVoBo: FirmaSolicitud[] = [];
  controlValidate = new FormControl("", [Validators.required]);
  firmsValidate: FirmaSolicitud[] = [];
  budget = 0;
  earn = 0;
  available = 0;
  startDate: Date;
  maxDate: Date;
  fechaCitaMedica: Date;
  fechaNacimiento: Date;
  anioNacimiento: number;
  controlFechaCita = new FormControl();
  controlFechaNacimiento = new FormControl();
  autorizacionEspecial = false;
  autorizacionesEspeciales: Partial<AutorizacionesEspeciales>;
  folioSeparado = false;
  respuestaAcceder = false
  controlFoliosS = new FormControl(false);
  constructor(
    private _folioFtp: FtpService,
    private fb: FormBuilder,
    private adapter: DateAdapter<any>,
    private _contrato: ContratoService,
    private patientTransfer: PatientTransferRoutesService,
    private firmService: FirmService,
    private _representante: RepresentanteLegalService,
    public dialog: MatDialog,
    private _solicitud: SolicitudTrasladoService,
    private modules: ModulesService,
    private accederService: AccederService
  ) {
    this.adapter.setLocale("es-mx");
    this.today.setHours(0, 0, 0, 0);
  }
  ngAfterViewInit() {
    this.input1.nativeElement.focus();
  }
  ngOnInit() {
    this.usuarioLog = jwt_decode(<string>localStorage.getItem("token"));
    this.unidadMedica = JSON.parse(<string>localStorage.getItem("unit"));
    this.formularioSolicitud();
    this.consultarPermisoFolioSeparado(this.unidadMedica.unidadMed.Id, '61367026d7be0a42003137fb')
    this.formSolicitud.controls["ftp01"].valueChanges.subscribe((res) => {
      if (res === "") {
        this.borrarDatos();
      }
      if (this.formSolicitud.controls["ftp01"].valid && res.length === 10) {
        this.formSolicitud.controls["fechaCita"].enable();
      } else {
        this.formSolicitud.controls["fechaCita"].disable();
      }
    });
    this.formSolicitud.controls["unidadMedId"].setValue(this.unidadMedica.unidadMedId);
    this.fetchFirms();
  }

  consultarPermisoFolioSeparado(unit, modulo: string) {
    this.modules.ConsultarPermisosUnidad(unit, modulo).subscribe(result => {
      console.log(result.Data)
      this.folioSeparado = result.Data
    })
  }
  
  buscarFolio() {
    const folio: string = this.formSolicitud.controls["ftp01"].value !== null ? this.formSolicitud.controls["ftp01"].value : "";
    this.borrarDatos();
    if (folio.length < 10) {
      Swal.fire({
        position: "center",
        title: "Formato de folio incorrecto",
        text: `Debe contener 10 numeros`,
        icon: "warning",
        showConfirmButton: true,
      });
    } else {
      this._folioFtp.consultarFolioFtp(folio).subscribe(
        (res) => {
          this.folio = res.Data;
          const extraerFecha = res.Data.fechaHora.split(" ");
          const separarFecha = extraerFecha[0].split("-");
          const separarHora = extraerFecha[1].split(":");
          this.fechaCitaMedica = new Date(separarFecha[0], separarFecha[1] - 1, separarFecha[2], separarHora[0], separarHora[1]);
          if (this.fechaCitaMedica >= this.today) {
            this.asignarDatosFtp()
            Swal.fire({
              position: "center",
              title: 'Folio encontrado',
              icon: "success",
              showConfirmButton: false,
              timer: 900,
            });
          } else {
            Swal.fire({
              position: "center",
              title: "El folio contiene una fecha de cita anterior al dia actual.",
              text: `¿Deseas continuar?`,
              icon: "warning",
              showConfirmButton: true,
              showCancelButton: true,
              confirmButtonColor: '#3085d6',
              cancelButtonColor: '#d33',
              allowOutsideClick: false,
              confirmButtonText: 'Continuar',
              cancelButtonText: 'Cancelar'
            }).then((result) => {
              if (result.isConfirmed) {
                Swal.fire({
                  title: 'Generar solicitud',
                  inputLabel: 'Se guardaran registros de estas solicitudes atipicas.',
                  inputPlaceholder: 'Ingresa el motivo.',
                  input: 'textarea',
                  confirmButtonColor: '#3085d6',
                  allowOutsideClick: false,
                  cancelButtonColor: '#d33',
                  cancelButtonText: 'Cancelar',
                  showCancelButton: true,
                  confirmButtonText: 'Continuar',
                  preConfirm: (value) => {
                    if (!value) {
                      Swal.showValidationMessage(
                        'El motivo no puede ser vacio'
                      )
                    } else if (value.length < 20) {
                      Swal.showValidationMessage(
                        'Escribir minimo 20 caracteres.'
                      )
                    } else {
                      this.autorizacionEspecial = true;
                      this.autorizacionesEspeciales = {
                        justificacion: value,
                        tipoAutorizacion: 'Fecha de cita anterior',
                      }
                    }
                  }
                }).then((result) => {
                  if (result.isConfirmed) {
                    Swal.fire(
                      'Validado',
                      'Tu solicitud ha sido aprobada, ahora genera tu Solicitud de traslado.',
                      'success'
                    )
                    this.today = this.fechaCitaMedica
                    this.asignarDatosFtp();
                  }
                })
              }
            })
          }
        },
        (err) => {
          if (err.error !== "" || err.error !== undefined) {
            const erores = <RespuestaPeticion>err.error;
            if (erores.Data !== null) {
              Swal.fire({
                position: "center",
                title: erores.Mensaje,
                html:
                  "Transporte: <b>" +
                  erores.Data.transporte +
                  "</b>,<br>" +
                  "Importe emitido: $<b>" +
                  erores.Data.importe +
                  "</b> <br>",
                icon: "error",
                showConfirmButton: true,
              });
            } else {
              Swal.fire({
                position: "center",
                title: erores.Mensaje,
                icon: "error",
                showConfirmButton: true,
              });
            }
          }
        }
      );
    }
  }

  asignarDatosFtp() {
    this.fechaCita = this.fechaCitaMedica.toISOString();
    const date = moment(this.fechaCita).format("DD-MM-YYYY");
    this.validarFecha = moment(date, "DD-MM-YYYY");
    this.especialidad = this.folio.especialidad;
    this.formSolicitud.controls["fechaCita"].enable();
    this.formSolicitud.controls["fechaCita"].setValue(this.fechaCitaMedica);
    this.formSolicitud.controls["ftp01"].setValue(this.folio.folioTp);
    this.formSolicitud.controls["nombrePaciente"].setValue(this.folio.paciente);
    this.formSolicitud.controls["nss"].setValue(this.folio.nss);
    this.formSolicitud.controls["agregadoMedico"].setValue(this.folio.agregado);
    this.formSolicitud.controls["especialidad"].setValue(this.folio.especialidad);
    this.formSolicitud.controls["nombrePaciente"].disable();
    this.formSolicitud.controls["nss"].disable();
    this.formSolicitud.controls["agregadoMedico"].disable();
    this.formSolicitud.controls["usuarioId"].setValue(this.usuarioLog.UsuarioId);
    this.anioNacimiento = parseInt(this.folio.agregado.slice(2, 6));
    this.controlValidate.setValue(this.firmsValidate[0].id);
    this.controlVoBo.setValue(this.firmsVoBo[0].id);
    this.validarEdad(this.anioNacimiento);
    this.consultarContratos();
  }
  fetchFirms() {
    this.firmService
      .GetAllTypeFirms(this.unidadMedica.unidadMedId, "1", "true")
      .subscribe(
        (res) => {
          this.firmsVoBo = res.Data;
        },
        (err) => {
          if (err.error !== "" || err.error !== undefined) {
            const errors = <RespuestaPeticion>err.error;
            Swal.fire({
              position: "center",
              title: errors.Mensaje,
              icon: "error",
              showConfirmButton: true,
            });
          }
        }
      );
    this.firmService
      .GetAllTypeFirms(this.unidadMedica.unidadMedId, "2", "true")
      .subscribe(
        (res) => {
          this.firmsValidate = res.Data;
        },
        (err) => {
          if (err.error !== "" || err.error !== undefined) {
            const errors = <RespuestaPeticion>err.error;
            Swal.fire({
              position: "center",
              title: errors.Mensaje,
              icon: "error",
              showConfirmButton: true,
            });
          }
        }
      );
  }

  /* ******  Formulario reactivo para la solicitud  ******   */
  formularioSolicitud() {
    this.formSolicitud = this.fb.group({
      especialidad: [
        { value: "", disabled: true },
        Validators.compose([Validators.required]),
      ],
      representanteLegal: [
        { value: "", disabled: true },
        Validators.compose([Validators.required]),
      ],
      nombrePaciente: [
        { value: "", disabled: true },
        Validators.compose([Validators.required]),
      ],
      nss: [
        { value: "", disabled: true },
        Validators.compose([
          Validators.required,
          Validators.pattern("^([0-9])*$"),
          Validators.minLength(10),
        ]),
      ],
      agregadoMedico: [
        { value: "", disabled: true },
        Validators.compose([
          Validators.required,
          Validators.pattern(
            "^[1-6]{1}[mfMF]{1}[0-9]{4}(([oO]{1}[rR]{1})|([eE]{1}[tT]{1})|([eE]{1}[cC]{1})|([sS]{1}[aA]{1})|" +
            "([eE]{1}[sS]{1})|([sS]{1}[fF]{1})|([pP]{1}[eE]{1})|([nN]{1}[dD]{1}))"
          ),
        ]),
      ],
      fechaCita: [
        { value: "", disabled: true },
        Validators.compose([Validators.required]),
      ],
      origen: [
        { value: "", disabled: true },
        Validators.compose([Validators.required]),
      ],
      destino: [
        { value: "", disabled: true },
        Validators.compose([Validators.required]),
      ],
      acompañante: [null, Validators.compose([Validators.required])],
      ftp01: [
        null,
        Validators.compose([
          Validators.required,
          Validators.maxLength(10),
          Validators.minLength(10),
          Validators.pattern("^([0-9])*$"),
        ]),
      ],
      contratoNumeroContrato: [
        { value: "", disabled: true },
        Validators.compose([Validators.required]),
      ],
      unidadMedId: [null, Validators.compose([Validators.required])],
      usuarioId: [null, Validators.compose([Validators.required])],
    });
  }

  validarEdad(anioNac: number) {
    this.preguntarFechaNac = false;
    this.pacienteInapan = false;
    this.edad = this.validarFecha.year() - anioNac;
    this.tipo_paciente = "Niño";
    this.tipoBoleto = 2;
    const nss = this.formSolicitud.get('nss').value
    const agregado = this.formSolicitud.get('agregadoMedico').value
    if (this.edad < 5) {
      this.tipo_boleto = "Sin boleto";
      this.tipoBoleto = 0;
    }
    if (this.edad === 5) {
      this.consultarFechaNacAcceder(nss, agregado)
      this.asignarFechas();
    }
    if (this.edad > 5 && this.edad < 13) {
      this.tipo_boleto = " con Medio boleto";
      this.tipoBoleto = 2;
    }
    if (this.edad === 13) {
      this.consultarFechaNacAcceder(nss, agregado)
      this.asignarFechas();
    }
    if (this.edad > 13) {
      this.tipo_boleto = "con Boleto Entero";
      this.tipoBoleto = 1;
    }
    if (this.edad === 16) {
      this.consultarFechaNacAcceder(nss, agregado)
      this.asignarFechas();
    }
    if (this.edad > 17) {
      this.tipo_paciente = "Adulto";
    }
    if (this.edad >= 60) {
      this.pacienteInapan = true;
    }
  }
  asignarFechas() {
    this.startDate = new Date(this.anioNacimiento, 0, 1);
    this.maxDate = new Date(this.anioNacimiento, 11, 31);
  }
  valorAcompaniante(event: MatRadioChange) {
    if (event.value === "si") {
      const dialogRef = this.dialog.open(DialogAutorizaAcompComponent, {
        height: "auto",
        width: "45%",
        disableClose: true,
        data: { edad: this.edad, unidadMed: this.unidadMedica.unidadMed.Id },
      });
      dialogRef.afterClosed().subscribe((result) => {
        if (result === false) {
          this.control_acompaniante.setValue("no");
          this.formSolicitud.controls["acompañante"].setValue(false);
        } else {
          this.formSolicitud.controls["acompañante"].setValue(true);
          this.autorizacionAcompaniante = result.autorizacion;
        }
      });
    } else {
      this.formSolicitud.controls["acompañante"].setValue(false);
    }
  }
  consultarFechaNacAcceder(nss, agregado: string) {
    this.respuestaAcceder=true
    // Swal.fire({
    //   title: 'Uploading...',
    //   html: 'Please wait...',
    //   allowEscapeKey: false,
    //   allowOutsideClick: false,
    //   didOpen: () => {
    //     Swal.showLoading()
    //   }
    // });

    this.accederService.ConsultarDerechoabiente(nss).subscribe(result => {

      if (result.Data.Stat && result != null) {
        this.respuestaAcceder = false;
        const tmp = result.Data.Data.filter(item => {
          return item.AgregadoMedico === agregado;
        });
        console.log(tmp)
        const separarFecha = tmp[0].FechaNacimiento.split("/");
        console.log(separarFecha)
        this.fechaNacimiento = new Date(separarFecha[2], separarFecha[1]-1, separarFecha[0])
        console.log('Fechaa Nc',this.fechaNacimiento)
        if (Object.prototype.toString.call(this.fechaNacimiento) === '[object Date]') {
          if (isNaN(this.fechaNacimiento.getTime())) {
            Swal.fire({
              position: "center",
              title: 'Error al consultar la fecha de nacimiento del paciente',
              text: 'Ingresa la fecha de nacimiento de manera manual',
              icon: "error",
              showConfirmButton: true,
            });
            console.log('Fecha de nacimiento no valida')
            this.preguntarFechaNac = true;
            this.respuestaAcceder = false;
          }
        }
        this.validarFechaCitaFechaNac();
      } else {
        Swal.fire({
          position: "center",
          title: result.Data.Usrmsg,
          text: 'Ingresa la fecha de nacimiento de manera manual',
          icon: "error",
          allowOutsideClick: false,
          showConfirmButton: true,
        });
        this.preguntarFechaNac = true;
        console.log(result.Data.Usrmsg)
        this.respuestaAcceder = false;
      }
    },
      (err) => {
        if (err.error !== "" || err.error !== undefined) {
          const errors = <RespuestaPeticion>err.error;
          Swal.fire({
            position: "center",
            title: errors.Mensaje,
            allowOutsideClick: false,
            text: 'Ingresa la fecha de nacimiento de manera manual',
            icon: "error",
            showConfirmButton: true,
          });
          this.preguntarFechaNac = true;
          this.respuestaAcceder = false;
        }
      }
    );
  }
  fechaNacimientoSeleccionada(event: any) {
    this.fechaNacimiento = new Date(
      event.value._i.year,
      event.value._i.month,
      event.value._i.date
    );
    this.validarFechaCitaFechaNac();
  }
  validarFechaCitaFechaNac() {
    this.edad = this.validarFecha.year() - this.anioNacimiento;
    const mesCita = this.fechaCitaMedica.getMonth();
    const mesNacimiento = this.fechaNacimiento.getMonth();
    const diaCita = this.fechaCitaMedica.getDate();
    const diaNacimiento = this.fechaNacimiento.getDate();
    if (this.edad > 18) {
      this.tipo_paciente = 'Adulto';
    } else {
      this.tipo_paciente = 'Niño'
    }
    if (mesNacimiento < mesCita) {
      //ya cumplio años
      this.diferenciaMenorCero(this.edad);
    } else if (mesNacimiento > mesCita) {
      //cumplira años el proximo mes
      this.diferenciaMayorCero(this.edad);
    } else {
      //cumplira años el mismo mes de la cita verificar el dia
      if (diaNacimiento > diaCita) {
        this.diferenciaMayorCero(this.edad);
      } else {
        this.diferenciaMenorCero(this.edad);
      }
    }
  }

  diferenciaMenorCero(edad: number) {
    switch (edad) {
      case 5:
        this.tipo_boleto = "Medio Boleto";
        this.tipoBoleto = 2;
        break;
      case 13:
        this.tipo_boleto = "Boleto Entero";
        this.tipoBoleto = 1;
        break;
      case 16:
        this.acompaniante = true;
        this.control_acompaniante.setValue("");
    }
  }

  diferenciaMayorCero(edad: number) {
    switch (edad) {
      case 5:
        this.tipo_boleto = "Sin Boleto";
        this.tipoBoleto = 0;
        this.edad = this.edad - 1;
        break;
      case 13:
        this.tipo_boleto = "Medio Boleto";
        this.edad = this.edad - 1;
        this.tipoBoleto = 2;
        break;
      case 16:
        this.edad = this.edad - 1;
        this.acompaniante = true;
        break;
    }
  }

  tarjetaInapan(event: MatRadioChange) {
    if (event.value === "si") {
      this.tipo_boleto = "Medio Boleto INAPAN";
      this.observacionesPac = "Credencial INAPAN";
      this.tipoBoleto = 2;
    } else {
      this.tipo_boleto = "Boleto Entero";
      this.observacionesPac = "";
      this.tipoBoleto = 1;
    }
  }

  fechaCitaSeleccionada(event: any) {
    this.fechaCita = event.value.format();
    this.formSolicitud.controls["fechaCita"].setValue(event.value._d);
    const date = moment(this.fechaCita).format("DD-MM-YYYY");
    this.validarFecha = moment(date, "DD-MM-YYYY");
    this.fechaCitaMedica = new Date(
      event.value._i.year,
      event.value._i.month,
      event.value._i.date
    );
    this.consultarContratos();
    this.validarEdad(this.anioNacimiento);
  }

  numeroContratoSeleccionado(value: string) {
    this.consultDestination(value);
    this._representante.ConsultaRepresentante(value).subscribe(
      (res) => {
        if (res.Data !== null) {
          this.formSolicitud.controls["representanteLegal"].setValue(
            res.Data.NombreCompleto
          );
          this.formSolicitud.controls["representanteLegal"].disable();
        }
      },
      (err) => {
        if (err.error !== "" || err.error !== undefined) {
          const erores = <RespuestaPeticion>err.error;
          this.formSolicitud.controls["representanteLegal"].setValue("");
          this.formSolicitud.controls["representanteLegal"].enable();
        }
      }
    );
  }

  consultarContratos() {
    this._contrato
      .ConsultarContratoTraslado(this.unidadMedica.unidadMedId, this.fechaCita)
      .subscribe((res) => {
        if (res.Data !== null) {
          this.contratos = res.Data;
          this.formSolicitud.controls["contratoNumeroContrato"].enable();
          if (res.Data.length === 1) {
            this.formSolicitud.controls["contratoNumeroContrato"].setValue(res.Data[0].NumeroContrato);
            this.numeroContratoSeleccionado(res.Data[0].NumeroContrato);
          }
        } else {
          this.contratos = [];
          this.formSolicitud.controls["contratoNumeroContrato"].disable();
          void Swal.fire({
            icon: "error",
            title: "Sin Contrato vigente",
            text: "No hay un contrato vigente para la fecha de cita seleccionada, comunicalo con el jefe inmediato.",
          });
        }
      });
  }

  consultDestination(contractNumber: string) {
    this.patientTransfer
      .ConsultarDestinos(contractNumber, this.unidadMedica.unidadMedId)
      .subscribe((res) => {
        this.consultBudget(contractNumber, this.unidadMedica.unidadMedId);
        this.destinos = res.Data;
        this.formSolicitud.controls["destino"].enable();
        this.consultEarn("true", contractNumber, this.unidadMedica.unidadMedId);
      });
  }

  consultOrigin(value: string) {
    const contractNumber = this.formSolicitud.controls["contratoNumeroContrato"].value;
    this.patientTransfer
      .ConsultarOrigenes(contractNumber, this.unidadMedica.unidadMedId, value)
      .subscribe((res) => {
        this.origenes = res.Data;
        this.formSolicitud.controls["origen"].enable();
        this.formSolicitud.controls["origen"].setValue(this.origenes[0].origen);
        // this.consultTicketGoingAvailable(value,"entero",contractNumber,"true",this.unidadMedica.unidadMedId);
      });
  }

  consultBudget(contractNumber, unitId: string) {
    this.patientTransfer
      .ConsultBudget(contractNumber, unitId)
      .subscribe((result) => {
        //presupuesto
        this.budget = result.Data;
      });
  }

  consultEarn(state, contractNumber, unitId: string) {
    this.patientTransfer
      .ConsultEarn(state, contractNumber, unitId)
      .subscribe((result) => {
        this.earn = result.Data;
        this.available = this.budget - result.Data;
      });
  }

  // consultTicketGoingAvailable(
  //   destination,
  //   typeTicket,
  //   contractNumber,
  //   state,
  //   unitId: string
  // ) {
  //   this.patientTransfer
  //     .ConsultTicketGoingAvailable(
  //       destination,
  //       typeTicket,
  //       contractNumber,
  //       state,
  //       unitId
  //     )
  //     .subscribe((result) => {});
  // }

  borrarDatos() {
    const unidad = this.formSolicitud.controls["unidadMedId"].value;
    this.formSolicitud.reset();
    this.formSolicitud.markAsUntouched();
    this.especialidad = "";
    this.observacionesAcompa = "";
    this.tipoBoleto = 2;
    this.observacionesPac = "";
    this.controlVoBo.reset();
    this.controlValidate.reset();
    this.controlIda.setValue(true);
    this.controlRegreso.setValue(true);
    this.controlFoliosS.setValue(false);
    this.tipo_boleto = "";
    this.tipo_paciente = "";
    this.control_acompaniante.setValue("");
    this.formSolicitud.controls["unidadMedId"].setValue(unidad);
    this.autorizacionesEspeciales = {
      justificacion: '',
      tipoAutorizacion: '',
    }
  }

  vistaPreviaSolicitud() {
    this.solicitudTraslado = this.formSolicitud.getRawValue();
    this.solicitudTraslado.fechaCita = this.fechaCita;
    let cantidadBoletos = 0;
    if (this.controlIda.value) {
      cantidadBoletos += 1;
    }
    if (this.controlRegreso.value) {
      cantidadBoletos += 1;
    }
    /* ******  Si el paciente no tiene derecho a boletos  ******   */
    if (this.tipoBoleto === 0) {
      if (this.recoleccionMed != 2) {
        const paciente: Boletos = {
          Tipo: 1,
          Ida: this.controlIda.value,
          Regreso: this.controlRegreso.value,
          Total: 0,
          TipoBoleto: this.tipoBoleto,
          Observaciones:
            "Niño con " + this.edad + " años." + this.observacionesPac,
        };
        this.agregarPasajero(paciente);
      }
    }
    /* ******  Si el paciente tiene derecho a  medio boleto  ******   */
    if (this.tipoBoleto === 2 && this.edad < 13) {
      if (this.recoleccionMed != 2) {
        const paciente: Boletos = {
          Tipo: 1,
          Ida: this.controlIda.value,
          Regreso: this.controlRegreso.value,
          Total: cantidadBoletos,
          TipoBoleto: this.tipoBoleto,
          Observaciones:
            "Niño con " + this.edad + " años. " + this.observacionesPac,
        };
        this.agregarPasajero(paciente);
      }
    } else if (this.tipoBoleto === 2) {
      if (this.recoleccionMed != 2) {
        const paciente: Boletos = {
          Tipo: 2,
          Ida: this.controlIda.value,
          Regreso: this.controlRegreso.value,
          Total: cantidadBoletos,
          TipoBoleto: this.tipoBoleto,
          Observaciones: this.observacionesPac,
        };
        this.agregarPasajero(paciente);
      }
    }
    /* ******  Si el paciente tiene derecho a  boleto entero ******   */
    if (this.tipoBoleto === 1 && this.edad < 18) {
      if (this.recoleccionMed != 2) {
        const paciente: Boletos = {
          Tipo: 2,
          Ida: this.controlIda.value,
          Regreso: this.controlRegreso.value,
          Total: cantidadBoletos,
          TipoBoleto: this.tipoBoleto,
          Observaciones: "Niño con " + this.edad + " años. " + this.observacionesPac,
        };
        this.agregarPasajero(paciente);
      }
    } else if (this.tipoBoleto === 1) {
      if (this.recoleccionMed != 2) {
        const paciente: Boletos = {
          Tipo: 2,
          Ida: this.controlIda.value,
          Regreso: this.controlRegreso.value,
          Total: cantidadBoletos,
          TipoBoleto: this.tipoBoleto,
          Observaciones: this.observacionesPac,

        };
        this.agregarPasajero(paciente);
      }
    }
    if (this.formSolicitud.controls["acompañante"].value === true) {
      if (this.recoleccionMed != 1) {
        const acompaniante: Boletos = {
          Tipo: 3,
          Ida: this.controlIda.value,
          Regreso: this.controlRegreso.value,
          Total: cantidadBoletos,
          TipoBoleto: 1,
          Observaciones: this.observacionesAcompa,
        };
        this.agregarPasajero(acompaniante);
      }
    }
    if (this.acompanianteExtra) {
      let boletoTipo = 1;
      let observacion = "Acompañante extra ";
      const acompaniante: Boletos = {
        Tipo: 3,
        Ida: this.controlIda.value,
        Regreso: this.controlRegreso.value,
        Total: cantidadBoletos,
        TipoBoleto: boletoTipo,
        Observaciones: observacion,
      };
      this.agregarPasajero(acompaniante);
    }
    const firms = this.controlVoBo.value + "," + this.controlValidate.value;
    const dialogRef = this.dialog.open(VistaPreviaSolicitudComponent, {
      height: "auto",
      width: "80%",
      data: {
        solicitud: this.solicitudTraslado,
        boletos: this.pasajeros,
        unidad: this.unidadMedica.unidadMed.DenominacionUni,
        edad: this.edad,
        uiPrei: this.unidadMedica.unidadMed.UiPrei,
        autorizacion: this.autorizacionAcompaniante,
        firms: firms,
        autorizacionEsp: this.autorizacionesEspeciales,
        foliosSeparados: this.controlFoliosS.value,
      },
      disableClose: true,
    });
    dialogRef.afterClosed().subscribe((result) => {
      this.pasajeros = [];
      if (result) {
        this.borrarDatos();
        this.borrarExepciones();
        this.input1.nativeElement.focus();
      }
    });
  }

  agregarPasajero(pasajero: Boletos) {
    this.pasajeros.push(pasajero);
  }

  abrirDialogExcepciones() {
    const dialogRef = this.dialog.open(DialogExcepcionesComponent, {
      height: "auto",
      width: "40%",
      disableClose: true,
      data: {
        tipoBoleto: this.tipoBoleto,
        edad: this.edad,
        exepcion: this.exepcion,
      },
    });
    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        this.exepcion = result.exepcion;
        console.log(result)
        let justificacionMedi = ''
        switch (result.exepcion) {
          case 1:
            //si el resultado es 1 solo viaja el paciente si el resultado es 2 viaja solo el acompañante
            if (result.valor === "1") {
              this.recoleccionMed = 1;
              this.observacionesPac += " Recoleccion de medicamentos. ";
              justificacionMedi = 'Recoleccion de medicamentos, viaja el paciente';
            } else {
              this.recoleccionMed = 2;
              justificacionMedi = 'Recoleccion de medicamentos, viaja el acompañante';
              this.observacionesAcompa += " Recoleccion de medicamentos. ";
            }
            this.autorizacionesEspeciales = {
              justificacion: justificacionMedi,
              tipoAutorizacion: 'Recoleccion de medicamentos'
            }
            break;
          case 2:
            this.acompanianteExtra = true;
            this.autorizacionesEspeciales = {
              justificacion: result.valor.justificacion,
              tipoAutorizacion: 'Acompañante extra'
            }
            break;
          case 3:
            this.tipoBoleto = 2;
            this.tipo_boleto = " con Medio boleto";
            this.autorizacionesEspeciales = {
              justificacion: result.valor,
              tipoAutorizacion: 'Medio boleto para niños menores de 5 años'
            }
            this.observacionesPac = " Autorizacion de medio boleto.";
            break;
          case 4:
            this.exepcion = undefined;
            this.borrarExepciones();
            break;
        }
      }
    });
  }

  borrarExepciones() {
    if (this.preguntarFechaNac) {
      this.validarFechaCitaFechaNac();
    }
    if (this.tipo_boleto !== "Medio Boleto INAPAN") {
      this.validarEdad(this.anioNacimiento);
    }
    this.acompanianteExtra = false;
    this.observacionesPac = "";
    this.recoleccionMed = 0;
    this.observacionesAcompa = "";
    this.autorizacionesEspeciales = {
      justificacion: '',
      tipoAutorizacion: '',
    }
  }

  imprimirSoloPlantilla() {
    this.imprimirPlantilla = !this.imprimirPlantilla;
    this.controlVoBo.reset();
    this.controlValidate.reset();
  }

  sendPrintTemplate() {
    this._solicitud
      .PrintTemplate(
        this.unidadMedica.unidadMedId,
        this.controlVoBo.value,
        this.controlValidate.value
      )
      .subscribe(
        (res) => {
          const downloadURL = URL.createObjectURL(res);
          window.open(downloadURL);
        },
        (err) => {
          if (err.error !== "" || err.error !== undefined) {
            Swal.fire({
              position: "center",
              title: "Error",
              text: "No se pudo descargar la plantilla",
              icon: "error",
              showConfirmButton: true,
            });
          }
        }
      );
  }
}

//Validaciones de este modulo de solicitudes
//1.-Validacion de disponibilidad de rutas
//2.-Validacion de fecha de la cita con la vigencia del contrato
//3.-Validacion de la fecha de la cita con el tipo de boleto (medio - entero)
//4.-Validacion del acompañanante
