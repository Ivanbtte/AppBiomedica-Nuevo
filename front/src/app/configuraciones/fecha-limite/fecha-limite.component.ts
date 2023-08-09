import {Component, OnInit} from '@angular/core';
import {MAT_MOMENT_DATE_FORMATS, MomentDateAdapter} from '@angular/material-moment-adapter';
import {DateAdapter, MAT_DATE_FORMATS, MAT_DATE_LOCALE} from '@angular/material/core';
import {FormControl} from '@angular/forms';
import {FechaLimiteService} from '../serviciosConfiguraciones/fecha-limite.service';
import {Fecha} from '../../core/models/fechaLimite.interface';
import Swal from 'sweetalert2';
import {MatDatepickerInputEvent} from '@angular/material/datepicker';

@Component({
  selector: 'app-fecha-limite',
  templateUrl: './fecha-limite.component.html',
  styleUrls: ['./fecha-limite.component.css'],
  providers: [
    {provide: MAT_DATE_LOCALE, useValue: 'ja-JP'},
    {provide: DateAdapter, useClass: MomentDateAdapter, deps: [MAT_DATE_LOCALE]},
    {provide: MAT_DATE_FORMATS, useValue: MAT_MOMENT_DATE_FORMATS},
  ],
})
export class FechaLimiteComponent implements OnInit {
  date = new FormControl(new Date());
  fechaActual = new Date();
  m = this.fechaActual.getUTCMonth();
  d = this.fechaActual.getDate();
  minDate = new Date(2019, this.m, this.d);
  maxDate = new Date(2030, 0, 1);
  minDate2 = new Date(2019, 0, 1);
  maxDate2 = new Date(2030, 0, 1);
  fechaInico = '';
  fechaLimite = '';
  fechaEstado = false;
  botonGuardar = false;
  botonAgregar = true;
  botonActualizar = false;
  estadoDataPick = true;
  fechaBackEnd!: Fecha;
  botonEditar = false;

  constructor(
    private adapter: DateAdapter<any>,
    private _fecha: FechaLimiteService
  ) {
    this.adapter.setLocale('es-mx');
    console.log(this.m, this.d);
  }

  ngOnInit() {
    this.consultarFecha();
  }

  consultarFecha() {
    this._fecha.ConsultaFecha().subscribe(res => {
      this.fechaBackEnd = res.Data;
      this.fechaInico = this.fechaBackEnd.FechaInicio;
      this.fechaLimite = this.fechaBackEnd.FechaLimite;
      this.fechaEstado = true;
      this.botonGuardar = false;
      this.botonActualizar = false;
      this.botonAgregar = false;
      this.botonEditar = true;
      this.estadoDataPick = true;
      console.log(this.fechaBackEnd);
    }, err => {
      console.log(err);
    });
  }

  actualizarFechas() {
    Swal.fire({
      title: 'Actualizar Fechas',
      html: 'Fecha de inicio: ' + this.fechaInico + ' Fecha límite: ' + this.fechaLimite,
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      cancelButtonText: 'Cancelar',
      confirmButtonText: 'Si ,actualizar'
    }).then((result) => {
      if (result.value) {
        this._fecha.ActualizarFecha({
          Id: this.fechaBackEnd.Id,
          FechaInicio: this.fechaInico,
          FechaLimite: this.fechaLimite,
          Estado: this.fechaBackEnd.Estado
        }).subscribe(res => {
          console.log(res.Data);
          Swal.fire(
            'Actualizado',
            'Tus fechas han sido actualizadas',
            'success'
          );
          this.consultarFecha();
        });
      }
    });
  }

  guardarFechas() {
    if (this.fechaInico !== '' && this.fechaLimite !== '') {
      Swal.fire({
        title: 'Establecer Fechas',
        html: 'Fecha de inicio: ' + this.fechaInico + ' Fecha límite: ' + this.fechaLimite,
        icon: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#3085d6',
        cancelButtonColor: '#d33',
        cancelButtonText: 'Cancelar',
        confirmButtonText: 'Si ,enviar'
      }).then((result) => {
        if (result.value) {
          this._fecha.AgregarFecha({
            Id: '',
            FechaInicio: this.fechaInico,
            FechaLimite: this.fechaLimite,
            Estado: true
          }).subscribe(res => {
            console.log(res.Data);
            Swal.fire(
              'Enviada',
              'Tus fechas han sido agregadas',
              'success'
            );
            this.consultarFecha();
          });
        }
      });
    } else {
      Swal.fire({
        icon: 'error',
        title: 'Error',
        text: 'Debes elegir las fechas',
      });
    }
  }

  hacerVisibleFormulario() {
    this.fechaEstado = true;
    this.botonAgregar = false;
    this.estadoDataPick = false;
    this.botonGuardar = true;
  }

  editarFechas() {
    this.estadoDataPick = false;
    this.botonActualizar = true;
  }

/*  valorFecha(event: MatDatepickerInputEvent<Date>) {
    this.fechaInico = event.value.
    const dia = event.value['_i'].date;
    const mes = event.value['_i'].month;
    const anio = event.value['_i'].year;
    this.minDate2 = new Date(anio, mes, dia + 1);
    console.log(event);
  }

  valorFechaLimite(event: MatDatepickerInputEvent<Date>) {
    this.fechaLimite = event.targetElement['value'];
    const dia = event.value['_i'].date;
    const mes = event.value['_i'].month;
    const anio = event.value['_i'].year;
    this.maxDate = new Date(anio, mes, dia - 1);
  }*/
}
