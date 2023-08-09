import {Component, OnInit} from '@angular/core';
import {Location} from '@angular/common';
import {RespuestaPeticion} from '../../core/models/estructuras_generales';
import {FianzaService} from '../../core/servicios/ContratosServicios/fianza.service';
import {ActivatedRoute, Params} from '@angular/router';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import Swal from 'sweetalert2';

@Component({
  selector: 'app-fianza',
  templateUrl: './fianza.component.html',
  styleUrls: ['./fianza.component.css']
})
export class FianzaComponent implements OnInit {
  estado_fianza!: boolean;
  numero_contrato! :string;
  formulario_fianza!: FormGroup;

  constructor(
    private _ruta: Location,
    private _fianza: FianzaService,
    private _route: ActivatedRoute,
    private fb: FormBuilder,
  ) {
  }

  ngOnInit() {
    this._route.params.subscribe((params: Params) => {
      this.numero_contrato = params.numero_contrato;
      this.consultarFianza();
      this.crearFormulario();
      // this.product = this.productsService.getProduct(id);
    });
  }

  crearFormulario() {
    this.formulario_fianza = this.fb.group({
      Afianzadora: ['', Validators.compose([Validators.required])],
      NumPoliza: ['', Validators.compose([Validators.required])],
      Tipo: ['', Validators.compose([Validators.required])],
      MontoFianza: ['', Validators.compose([Validators.required])],
    });
  }

  /* ******  Funcion para consultar fianza  ******   */
  consultarFianza() {
    this._fianza.Consultarfianza(this.numero_contrato).subscribe(res => {
      this.estado_fianza = res.Data['estado'];
      console.log(res.Data['datos']);
      const fianza = res.Data['datos'];
      this.formulario_fianza.controls['Afianzadora'].setValue(fianza.Afianzadora);
      this.formulario_fianza.controls['NumPoliza'].setValue(fianza.NumPoliza);
      this.formulario_fianza.controls['Tipo'].setValue(fianza.Tipo);
      this.formulario_fianza.controls['MontoFianza'].setValue(fianza.MontoFianza);
      this.formulario_fianza.disable();
    }, err => {
      if (err.error !== '' || err.error !== undefined) {
        const erores = <RespuestaPeticion> err.error;
        console.log(erores);
        this.estado_fianza = erores.Data['estado'];
      }
    });
  }

  /* ******  Funcion para agregar una fianza  ******   */
  agregarFianza() {
    console.log(this.formulario_fianza.value);
    Swal.fire({
      title: 'Confirmación',
      text: '¿Estas seguro de agregar estos datos?',
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#a2a2a2',
      confirmButtonText: 'Si',
      cancelButtonText: 'Cancelar',
      reverseButtons: true
    }).then((result) => {
      if (result.value) {
        this._fianza.AgregarFianza(this.numero_contrato, this.formulario_fianza.value).subscribe(res => {
          Swal.fire(
            'Agregado',
            res.Mensaje,
            'success'
          );
          this._ruta.back();
        }, err => {
          if (err.error !== '' || err.error !== undefined) {
            const erores = <RespuestaPeticion> err.error;
            Swal.fire(
              'Error',
              erores.Mensaje,
              'error'
            );
          }
        });
      }
    });
  }

  regresar() {
    this._ruta.back();
  }
}
