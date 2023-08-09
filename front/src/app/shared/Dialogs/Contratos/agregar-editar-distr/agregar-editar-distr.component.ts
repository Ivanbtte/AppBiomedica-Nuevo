import {Component, Inject, OnInit} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {FormBuilder, FormControl, FormGroup, Validators} from '@angular/forms';
import {UsuarioService} from '../../../../authentication/servicios/usuario.service';
import {UnidadesM} from '../../../../authentication/interfaces/usuario.interface';
import {DistribucionUnidadMed} from '../../../../core/models/concepto_contrato.interface';
import Swal from 'sweetalert2';
import {RespuestaPeticion} from '../../../../core/models/estructuras_generales';
import {DistribucionConceptosService} from '../../../../core/servicios/ContratosServicios/distribucion-conceptos.service';
import {DelegacionesService} from '../../../../core/servicios/delegaciones.service';

export interface DistribucionData {
  distribucion: any;
  concepto: any;
  editar: boolean;
  cantidadMax: number;
  id_concepto: any;
}

@Component({
  selector: 'app-agregar-editar-distr',
  templateUrl: './agregar-editar-distr.component.html',
  styleUrls: ['./agregar-editar-distr.component.css']
})
export class AgregarEditarDistrComponent implements OnInit {
  formularioDistribucion!: FormGroup;
  por_distribuir!: number;
  UM: UnidadesM []=[];
  distribucion: any[]=[];
  delegacionesArray: any[]=[];
  cambioDelegacion = false;

  constructor(
    public dialogRef: MatDialogRef<AgregarEditarDistrComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DistribucionData,
    private fb: FormBuilder,
    private _unidades: UsuarioService,
    private  _distribucion: DistribucionConceptosService,
    private _delegaciones: DelegacionesService,
  ) {
  }

  ngOnInit() {
    this.por_distribuir = this.data.cantidadMax;
    this.consultarDistribucion();
    this._delegaciones.ConsultarDelegaciones().subscribe(resultado => {
      this.delegacionesArray = resultado.Data;
      this.CrearFormularioReactivo();
    });
  }

  onNoClick(): void {
    this.dialogRef.close();
  }

  consultarDistribucion() {
    this._distribucion.ConsultarDistribucion(this.data.id_concepto.Id).subscribe(resultado => {
      this.distribucion = resultado.Data;
    });
  }

  CrearFormularioReactivo() {
    this.formularioDistribucion = this.fb.group({
      Cantidad: [null, Validators.compose([Validators.pattern('^([0-9])*$'),
        Validators.required, Validators.min(1), Validators.max(this.por_distribuir)])],
      UnidadMedClvPresupuestal: new FormControl({value: null, disabled: true}, Validators.required),
      Delegacion: [null, Validators.required],
    });
    if (this.data.editar) {
      this.formularioDistribucion.controls['Cantidad'].setValue(this.data.distribucion.Cantidad);
      this.formularioDistribucion.controls['Cantidad'].setValidators([Validators.pattern('^([0-9])*$'),
        Validators.required, Validators.min(1), Validators.max(this.por_distribuir + this.data.distribucion.Cantidad)]);
      this.formularioDistribucion.controls['Delegacion'].setValue(this.data.distribucion.UnidadMed.Delegacion);
      this.formularioDistribucion.controls['UnidadMedClvPresupuestal'].setValue(this.data.distribucion.UnidadMed);
      this.cambioDelegacion = true;
      this.valorSelecDeleg(this.data.distribucion.UnidadMed.Delegacion);
    }
  }

  compare(c1: { NombreDele: string }, c2: { NombreDele: string }) {
    return c1 && c2 && c1.NombreDele === c2.NombreDele;
  }

  UnidadSeleccionada(um: { DenominacionUni: string }, um2: { DenominacionUni: string }) {
    return um && um2 && um.DenominacionUni === um2.DenominacionUni;
  }

  valorSelecDeleg(value: any) {
    if (value !== null || value !== '') {
      this._unidades.unidades(value.ClvDele).subscribe(res => {
        this.UM = res.Data;
        this.formularioDistribucion.controls['UnidadMedClvPresupuestal'].reset();
        this.formularioDistribucion.controls['UnidadMedClvPresupuestal'].enable();
        this.FiltrarUnidadesMed();
      }, err => {
        console.log(err);
      });
    }
  }

  // ********************************************************************************************************
  /* *****  Funcion para eliminar unidades medicas del select que ya estan registradas   *****   */
  FiltrarUnidadesMed() {
    const unidades: Array<any> = this.distribucion.map(it => it.UnidadMed);
    for (let unidad of unidades) {
      const index = this.UM.map(function (e) {
        return e.Id;
      }).indexOf(unidad.Id);
      if (index >= 0) {
        this.UM.splice(index, 1);
      }
    }
    if (this.data.editar && this.cambioDelegacion) {
      this.formularioDistribucion.controls['UnidadMedClvPresupuestal'].setValue(this.data.distribucion.UnidadMed);
      this.UM.push(this.data.distribucion.UnidadMed);
      this.cambioDelegacion = false;
    }
    console.log('Eliminados', this.UM);
  }

  agregardistribucion() {
    const datos = this.formularioDistribucion.getRawValue();
    const clvDele = this.formularioDistribucion.controls['Delegacion'].value;
    console.log(datos);
    const distribucion: DistribucionUnidadMed = {
      Id: '',
      UnidadMedClvPresupuestal: datos.UnidadMedClvPresupuestal.ClvPresupuestal,
      ConceptoContratoId: this.data.id_concepto.Id,
      Cantidad: parseInt(datos.Cantidad, 10)
    };
    console.log(distribucion, this.data.id_concepto.ContratoNumeroContrato, clvDele.ClvDele, this.data.id_concepto.CantidadConcepto);
    Swal.fire({
      title: 'Confirmacion',
      text: 'Estas seguro de agregar esta distribucion',
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#a2a2a2',
      confirmButtonText: 'Si, Agregar',
      cancelButtonText: 'Cancelar',
      reverseButtons: true

    }).then((result) => {
      if (result.value) {
        this._distribucion.AgregaDistribucion(distribucion, this.data.id_concepto.CantidadConcepto, this.data.id_concepto.ContratoNumeroContrato,
          clvDele.ClvDele).subscribe(res => {
          Swal.fire(
            'Correcto',
            res.Mensaje,
            'success'
          );
          this.dialogRef.close(true);
        }, err => {
          if (err.error !== '' || err.error !== undefined) {
            const erores = <RespuestaPeticion>err.error;
            Swal.fire({
              position: 'center',
              title: erores.Mensaje,
              text: erores.Data,
              icon: 'error',
              showConfirmButton: true,
            });
          }
          console.log(err);
        });
      }
    });
  }

  EditarDistribucion() {
    const editar = this.formularioDistribucion.getRawValue();
    const clvDele = this.formularioDistribucion.controls['Delegacion'].value;
    console.log(editar);
    const distribucion: DistribucionUnidadMed = {
      Id: this.data.distribucion.Id,
      UnidadMedClvPresupuestal: editar.UnidadMedClvPresupuestal.ClvPresupuestal,
      ConceptoContratoId: this.data.id_concepto.Id,
      Cantidad: parseInt(editar.Cantidad, 10)
    };
    Swal.fire({
      title: 'Confirmacion',
      text: 'Estas seguro de editar esta distribucion',
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#a2a2a2',
      confirmButtonText: 'Si, Editar',
      cancelButtonText: 'Cancelar',
      reverseButtons: true

    }).then((result) => {
      if (result.value) {
        this._distribucion.ActualizarDistribucion(distribucion, this.data.id_concepto.ContratoNumeroContrato, clvDele.ClvDele).subscribe(res => {
          Swal.fire(
            'Correcto',
            res.Mensaje,
            'success'
          );
          this.dialogRef.close(true);
        }, err => {
          if (err.error !== '' || err.error !== undefined) {
            const erores = <RespuestaPeticion>err.error;
            Swal.fire({
              position: 'center',
              title: erores.Mensaje,
              text: erores.Data,
              icon: 'error',
              showConfirmButton: true,
            });
          }
        });
      }
    });
  }
}
