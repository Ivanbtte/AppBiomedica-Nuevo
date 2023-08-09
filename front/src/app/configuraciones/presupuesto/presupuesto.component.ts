import {Component, OnInit} from '@angular/core';
import {PresupuestoService} from '../serviciosConfiguraciones/presupuesto.service';
import {FormControl} from '@angular/forms';
import {Presupuesto} from '../../core/models/presupuesto';
import Swal from 'sweetalert2';

@Component({
  selector: 'app-presupuesto',
  templateUrl: './presupuesto.component.html',
  styleUrls: ['./presupuesto.component.css']
})
export class PresupuestoComponent implements OnInit {
  botonActualizar = false;
  botonEditar = false;
  botonAgregar = true;
  habilitarInput = false;
  presupuesto = new FormControl();
  value = '';
  resultado!: Presupuesto;

  constructor(
    private _presupuesto: PresupuestoService,
  ) {
  }

  ngOnInit() {
    this.consultarPresupuesto();
  }

  consultarPresupuesto() {
    this._presupuesto.ConsultaPresupuesto().subscribe(res => {
      this.resultado = res.Data;
      console.log(res.Data);
      this.botonAgregar = false;
      this.botonEditar = true;
      this.botonActualizar = false;
      this.value = this.resultado.Presupuesto.toString();
      this.habilitarInput = true;
    });
  }

  habilitarEdicion() {
    this.botonActualizar = true;
    this.habilitarInput = false;
  }

  actualizarPresupuesto() {
    const pre = this.presupuesto.value;
    console.log(pre);
    if (pre !== null) {
      this._presupuesto.ActualizarPresupuesto({
        Id: this.resultado.Id,
        Presupuesto: pre,
        Estado: this.resultado.Estado
      }).subscribe(res => {
        console.log(res.Data);
        this.consultarPresupuesto();
        Swal.fire(
          'Actualizado',
          'Tu presupuesto ha sido actualizado',
          'success'
        );
      });
    }
  }

  agregarPresupuesto() {
    const pre = this.presupuesto.value;
    console.log(pre);
    if (pre !== null) {
      this._presupuesto.AgregarPresupuesto({Id: '', Presupuesto: pre, Estado: true}).subscribe(res => {
        console.log(res.Data);
        this.consultarPresupuesto();
        Swal.fire(
          'Agregado',
          'Tu presupuesto ha sido agregado',
          'success'
        );
      });
    }
  }
}
