import { AfterViewInit, Component, ElementRef, Inject, OnInit, ViewChild, ViewEncapsulation } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MAT_MOMENT_DATE_FORMATS, MomentDateAdapter } from '@angular/material-moment-adapter';
import { DateAdapter, MAT_DATE_FORMATS, MAT_DATE_LOCALE } from '@angular/material/core';
import jwt_decode from "jwt-decode";
import { RespuestaPeticion } from 'src/app/core/models/estructuras_generales';
import { AccederService } from 'src/app/traslado-pacientes/core/acceder.service';
import { ModulesService } from 'src/app/traslado-pacientes/core/modules.service';
import Swal from 'sweetalert2';
import { SolicitudServLaboratorio } from '../../core/models/ServiciosSubrogados/trasladoPacientes/servicios_laboratorio.interface';
import * as moment from 'moment';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';
import { ServiciosSubrugadosLaboratorioService } from 'src/app/core/servicios/ServiciosSubrugadosLaboratorio/servicios-subrugados-laboratorio.service';
export interface DialogData {
  solicitud: SolicitudServLaboratorio;
}
@Component({
  selector: 'app-solicitud-subrogado',
  templateUrl: './solicitud-subrogado.component.html',
  styleUrls: ['./solicitud-subrogado.component.css'],
  providers: [
    { provide: MAT_DATE_LOCALE, useValue: "en-GB" },
    {
      provide: DateAdapter, useClass: MomentDateAdapter,
      deps: [MAT_DATE_LOCALE],
    },
    { provide: MAT_DATE_FORMATS, useValue: MAT_MOMENT_DATE_FORMATS },
  ],
  encapsulation: ViewEncapsulation.None
})
export class SolicitudSubrogadoComponent implements OnInit/*, AfterViewInit */ {
  @ViewChild("folionu", { static: false, read: ElementRef }) input1: ElementRef;
  unidad: any;
  unidadMedica: any;
  usuarioLog: any;
  today = new Date();
  folioSeparado = false;
  imprimirPlantilla = true; x
  controlVoBo = new FormControl("", [Validators.required]);
  controlValidate = new FormControl("", [Validators.required]);
  formSolicitud!: FormGroup;
  respuestaAcceder = false;
  dateCita = '';

  constructor(
    private modules: ModulesService,
    private fb: FormBuilder,
    private adapter: DateAdapter<any>,
    private accederService: AccederService,
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
    private _solicitud: ServiciosSubrugadosLaboratorioService,
  ) {
    this.adapter.setLocale("es-mx");
    this.today.setHours(0, 0, 0, 0);
  }

  ngAfterViewInit() {
    this.input1.nativeElement.focus();
  }

  ngOnInit(): void {
    this.dateCita = moment(this.data.solicitud.fecha).format();
    this.usuarioLog = jwt_decode(<string>localStorage.getItem("token"));
    this.unidad = JSON.parse(<string>localStorage.getItem("unit"));
    this.formularioSolicitud();
    this.consultarPermisoFolioSeparado(this.unidadMedica.unidadMed.Id, '64d51dd479b12822bc91a359ñ')
    this.formSolicitud.controls["unidadMedId"].setValue(this.unidadMedica.unidadMedId);
    console.log(this.unidad)
  }
  dataSourse: any[]

  //Método para revisar firmas
  fetchFirms() {
    throw new Error('Method not implemented.');
  }

  //Método que contiene todos los datos de la solicitud 
  formularioSolicitud() {
    this.formSolicitud = this.fb.group({
      ftp03: [
        null,
        Validators.compose([
          Validators.required,
          Validators.maxLength(11),
          Validators.minLength(11),
          Validators.pattern("^([0-9])*$"),
        ]),
      ],
      ftp04: [
        null,
        Validators.compose([
          Validators.required,
          Validators.maxLength(6),
          Validators.minLength(6),
          Validators.pattern("^([0-9])*$"),
        ]),
      ],
      nombreser: [
        { value: "", disabled: false },
        Validators.compose([Validators.required]),
      ],
      direccionser: [
        { value: "", disabled: false },
        Validators.compose([Validators.required]),
      ],
      NombrePac: [
        { value: "", disabled: true },
        Validators.compose([Validators.required]),
      ],
      ApellidoPac: [
        { value: "", disabled: true },
        Validators.compose([Validators.required]),
      ],
      CurpPac: [
        { value: "", disabled: true },
        Validators.compose([Validators.required]),
      ],
      VigenciaRamSeg: [
        { value: "", disabled: true },
        Validators.compose([Validators.required]),
      ],
      DiagnosticoClin: [
        { value: "", disabled: false },
        Validators.compose([Validators.required]),
      ],
      especificacion1: [
        { value: "", disabled: false },
        Validators.compose([Validators.required]),
      ],
      NomRaProve: [
        { value: "", disabled: false },
        Validators.compose([Validators.required]),
      ],
      RfcProve: [
        { value: "", disabled: false },
        Validators.compose([Validators.required]),
      ],
      DomiProve: [
        { value: "", disabled: false },
        Validators.compose([Validators.required]),
      ],
      TelProve: [
        { value: "", disabled: false },
        Validators.compose([Validators.required]),
      ],
      ContraProve: [
        { value: "", disabled: false },
        Validators.compose([Validators.required]),
      ],
      VegenProve1: [
        { value: "", disabled: false },
        Validators.compose([Validators.required]),
      ],
      VegenProve2: [
        { value: "", disabled: false },
        Validators.compose([Validators.required]),
      ],
      EspeProve: [
        { value: "", disabled: false },
        Validators.compose([Validators.required]),
      ],
      ElabFir: [
        { value: "", disabled: false },
        Validators.compose([Validators.required]),
      ],
      JefeSerFir: [
        { value: "", disabled: false },
        Validators.compose([Validators.required]),
      ],
      DirecFir: [
        { value: "", disabled: false },
        Validators.compose([Validators.required]),
      ],
      ElabFir2: [
        { value: "", disabled: false },
        Validators.compose([Validators.required]),
      ],
      JefeSerFir2: [
        { value: "", disabled: false },
        Validators.compose([Validators.required]),
      ],
      DirecFir2: [
        { value: "", disabled: false },
        Validators.compose([Validators.required]),
      ],
      ElabFir3: [
        { value: "", disabled: false },
        Validators.compose([Validators.required]),
      ],
      JefeSerFir3: [
        { value: "", disabled: false },
        Validators.compose([Validators.required]),
      ],
      DirecFir3: [
        { value: "", disabled: false },
        Validators.compose([Validators.required]),
      ],
    });

  }

  //Método que se encarga de imprimir la plantilla
  imprimirSoloPlantilla() {
    this.imprimirPlantilla = !this.imprimirPlantilla;
    this.controlVoBo.reset();
    this.controlValidate.reset();
  }

  //Consulta los permisos del folio
  consultarPermisoFolioSeparado(unit, modulo: string) {
    this.modules.ConsultarPermisosUnidad(unit, modulo).subscribe(result => {
      console.log(result.Data)
      this.folioSeparado = result.Data
    })
  }

  consultarNSS(nss: string) {
    this.respuestaAcceder = true
    this.accederService.ConsultarDerechoabiente(nss).subscribe(result => {
      if (result.Data.Stat && result != null) {
        this.respuestaAcceder = false;
        console.log(result)
      }
    },
      (err) => {
        if (err.error !== "" || err.error !== undefined) {
          const errors = <RespuestaPeticion>err.error;
          Swal.fire({
            position: "center",
            title: errors.Mensaje,
            allowOutsideClick: false,
            text: 'Ingresa los datos del paciente de forma manual',
            icon: "error",
            showConfirmButton: true,
          });
          this.formSolicitud.controls["NombrePac"].enable();
          this.formSolicitud.controls["ApellidoPac"].enable();
          this.formSolicitud.controls["CurpPac"].enable();
        }
      }
    );
  }

  guardarDatos() {
    const dateTemporal = this.data.solicitud.fecha;
    Swal.fire({
      title: 'Confirmar',
      text: 'Al enviar la solicitud, no podras editar nuevamente los datos. En caso de requerir cambios posteriores tendras que cancelar' +
        ' el folio',
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#36b729',
      cancelButtonColor: '#d33',
      confirmButtonText: 'Enviar',
      cancelButtonText: 'Cancelar'
    }).then((result) => {
      if (result.value) {
        this._solicitud.RealizarSolicitud(dateTemporal).subscribe(res => {


        }, err => {
          if (err.error !== '' || err.error !== undefined) {
            const erores = <RespuestaPeticion>err.error;
            Swal.fire({
              title: 'Error al registrar la solictud',
              text: erores.Mensaje,
              icon: 'error',
            });
          }
        })
      }
    });
  }
}


