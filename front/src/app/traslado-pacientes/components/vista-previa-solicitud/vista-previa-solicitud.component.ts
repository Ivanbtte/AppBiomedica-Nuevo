import { Component, Inject, OnInit } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import {
  Acompaniante,
  Boletos,
  Solicitudtraslado,
  SolicitudtrasladoBoletos
} from '../../../core/models/ServiciosSubrogados/trasladoPacientes/solicitud_traslado.interface';
import { SolicitudTrasladoService } from '../../../core/servicios/SolicitudTraslados/solicitud-traslado.service';
import { RespuestaPeticion } from '../../../core/models/estructuras_generales';
import Swal from 'sweetalert2';
import * as _moment from 'moment';
import { AutorizacionesEspeciales } from '../../../core/models/ServiciosSubrogados/trasladoPacientes/solicitud_traslado.interface';
import { flatMap } from 'rxjs/operators';

const moment = _moment;

export interface DialogData {
  solicitud: Solicitudtraslado;
  boletos: Boletos[];
  unidad: string;
  edad: number;
  uiPrei: string;
  autorizacion: Acompaniante;
  firms: string;
  autorizacionEsp: AutorizacionesEspeciales;
  foliosSeparados: boolean;
}

@Component({
  selector: 'app-vista-previa-solicitud',
  templateUrl: './vista-previa-solicitud.component.html',
  styleUrls: ['./vista-previa-solicitud.component.scss'],
})
export class VistaPreviaSolicitudComponent implements OnInit {
  acompaniante = '';
  displayedColumns: string[] = ['Tipo', 'Ida', 'Regreso', 'Boleto', 'Total', 'Observaciones'];
  datasource: any;
  solicitudRealizada = true;
  nota = '';
  solicitudValidada = false;
  datos!: SolicitudtrasladoBoletos;
  folioSolicitud = '';
  dateCita = '';

  constructor(
    public dialogRef: MatDialogRef<VistaPreviaSolicitudComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
    private _solicitud: SolicitudTrasladoService
  ) {
  }

  ngOnInit() {
    this.dateCita = moment(this.data.solicitud.fechaCita).format();
    //this.dateCita = moment(this.data.solicitud.fechaCita).format('DD-MM-YYYY');
    const pasajeros = this.data.boletos.map(ite => {
      return {
        tipo: this.regresaTipo(ite.Tipo),
        ida: this.regresaBoleano(ite.Ida),
        regreso: this.regresaBoleano(ite.Regreso),
        total: ite.Total,
        tipoBoleto: this.regresaBoleto(ite.TipoBoleto),
        observaciones: ite.Observaciones
      };
    });
    this.datasource = pasajeros;
    if (this.data.solicitud.acompañante === true) {
      this.acompaniante = 'Si';
    } else {
      this.acompaniante = 'No';
    }
    console.log(this.data.solicitud.agregadoMedico.substring(6, 8), this.data.solicitud.agregadoMedico)
    if (this.data.edad < 16 && this.data.boletos.length === 1 && this.data.solicitud.agregadoMedico.substring(6, 8) != 'ES') {
      this.nota = 'Un menor de edad no tiene permitido viajar sin acompañante';
      this.solicitudValidada = true;
    }
    if (this.data.boletos.length === 1 && this.data.boletos[0].Tipo === 3 && this.data.solicitud.agregadoMedico.substring(6, 8) != 'ES') {
      this.solicitudValidada = false;
    }
    if (this.data.boletos.length === 0) {
      this.solicitudValidada = true;
    }
  }
  regresaBoleano(valor: boolean): string {
    if (valor) {
      return 'Si';
    } else {
      return ' - ';
    }
  }
  regresaBoleto(valor: number) {
    switch (valor) {
      case 0:
        return 'Sin boleto';
      case 1:
        return 'Entero';
      case 2:
        return 'Medio';
    }
  }
  regresaTipo(valor: number) {
    switch (valor) {
      case 1:
        return this.validaTipo();
      case 2:
        return this.validaTipo();
      case 3:
        return 'Acompañante Adulto';
    }
  }
  validaTipo(): string {
    if (this.data.edad < 18) {
      return 'Paciente Niño'
    } else {
      return 'Paciente Adulto'
    }
  }
  realizarsolicitud() {
    const dateTemporal = this.data.solicitud.fechaCita;
    this.datos = {
      Solicitud: this.data.solicitud,
      Boletos: this.data.boletos,
      Acompaniante: this.data.autorizacion,
      AutorizacionesEspeciales: this.data.autorizacionEsp
    };
    let folioSepa = ''
    if (this.data.foliosSeparados) {
      folioSepa = 'True'
    }
    console.log('final', this.datos)
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
        delete this.datos.Solicitud.fechaCita;
        this._solicitud.RealizarSolicitud(this.datos, dateTemporal, this.data.uiPrei, this.data.firms, folioSepa).subscribe(res => {
          if (this.data.foliosSeparados) {
            this.folioSolicitud = res.Data.folio1.folio + ',' + res.Data.folio2.folio;
          } else {
            this.folioSolicitud = res.Data.folio;
          }
          console.log(this.folioSolicitud)
          this.solicitudRealizada = false;
          Swal.fire({
            position: 'center',
            title: res.Mensaje,
            icon: 'success',
            showConfirmButton: true,
          });
        }, err => {
          if (err.error !== '' || err.error !== undefined) {
            const erores = <RespuestaPeticion>err.error;
            Swal.fire({
              title: 'Error al registrar la solictud',
              text: erores.Mensaje,
              icon: 'error',
            });
          }
        });
      }
    });

  }

  onNoClick(): void {
    this.dialogRef.close();
  }

  Cerrar() {
    this.dialogRef.close(true);
  }

  imprimirSolicitud() {
    let folioSepa = ''
    if (this.data.foliosSeparados) {
      folioSepa = 'True'
    }
    this._solicitud.printDataOnly(this.folioSolicitud, '', folioSepa).subscribe(res => {
      const downloadURL = URL.createObjectURL(res);
      window.open(downloadURL);
    }, err => {
      if (err.error !== '' || err.error !== undefined) {
        Swal.fire({
          position: 'center',
          title: 'Error',
          text: 'No se pudo descargar la plantilla',
          icon: 'error',
          showConfirmButton: true,
        });
      }
    });
  }

  imprimirSolicitudCompleta() {
    let folioSepa = ''
    if (this.data.foliosSeparados) {
      folioSepa = 'True'
    }
    this._solicitud.printRequestComplete(this.folioSolicitud, 'true', '', folioSepa).subscribe(res => {
      const downloadURL = URL.createObjectURL(res);
      window.open(downloadURL);
    }, err => {
      if (err.error !== '' || err.error !== undefined) {
        Swal.fire({
          position: 'center',
          title: 'Error',
          text: 'No se pudo descargar la plantilla',
          icon: 'error',
          showConfirmButton: true,
        });
      }
    });
  }

}
