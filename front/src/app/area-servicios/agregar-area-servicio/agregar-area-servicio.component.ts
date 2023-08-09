import {Component, OnInit} from '@angular/core';
import {AreaServicioService} from '../../core/servicios/AreasServicios/area-servicio.service';
import jwt_decode from 'jwt-decode';
import {UsuarioLogeado} from '../../core/models/UsuarioLogin.interface';
import {ServiciosUsuariosService} from '../../core/servicios/ServiciosUsuarios/servicios-usuarios.service';
import {NombramientoService} from '../../core/servicios/nombramiento.service';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {AreaServicioUnidades} from '../../core/models/areaServicios';
import Swal from 'sweetalert2';

@Component({
  selector: 'app-agregar-area-servicio',
  templateUrl: './agregar-area-servicio.component.html',
  styleUrls: ['./agregar-area-servicio.component.css']
})
export class AgregarAreaServicioComponent implements OnInit {
  formularioVisible = false;
  configuraciones = {
    columns: {
      unidad: {
        title: 'Unidad Medica',
        filter: true
      },
      servicio: {
        title: 'Servicio',
        filter: true
      },
      area: {
        title: 'Area',
        filter: true
      },
      cc: {
        title: 'Centro de Costo',
        filter: true
      }
    },
    actions: false,
  };
  public usuarioLog!: UsuarioLogeado;
  public datos!: any;
  public unidadesMedicas!: any[];
  public serviciosUsuarioResul!: any[];
  public id_unidad: any;
  public id_servicio: any;
  public formularioAgregar!: FormGroup;

  constructor(
    private _areaServicios: AreaServicioService,
    private _serviciosUsuarios: ServiciosUsuariosService,
    private _nombramiento: NombramientoService,
    private formBuilder: FormBuilder
  ) {
  }

  /* ******  Refactorizar este modulo----  ******   */
  ngOnInit() {
    this.usuarioLog = jwt_decode(<string> localStorage.getItem('token'));
    this.fetchAreaService();
    this.getUnitMedic();
    this.formularioAgregar = this.formBuilder.group({
      UnidadMedId: ['', Validators.required],
      ServiciosProformaId: ['', Validators.required],
      Area: ['', Validators.required],
      CentroCosto: [''],
    });
  }

  getUnitMedic() {
    this._nombramiento.nombramiento(this.usuarioLog.PersonaId).subscribe(res => {
      const datos = res.Data;
      this.unidadesMedicas = datos.UnidadMed;
      console.log(this.unidadesMedicas);
    });
  }

  fetchAreaService() {
    this._areaServicios.ConsultarAreaServicio('', '', this.usuarioLog.UsuarioId).subscribe(res => {
      console.log(res.Data);
      const service = res.Data.map((ite: {
        UnidadMed: { DenominacionUni: any; }; ServiciosProforma: { Descripcion: any; };
        Area: any; CentroCosto: any;
      }) => {
        return {
          unidad: ite.UnidadMed.DenominacionUni,
          servicio: ite.ServiciosProforma.Descripcion,
          area: ite.Area,
          cc: ite.CentroCosto
        };
      });
      console.log(service);
      this.datos = service;
    });
  }

  valueSelectUnit(value: string) {
    console.log(value);
    this.id_unidad = value;
    this.fetchServiceUser(value, this.usuarioLog.UsuarioId);
  }

  valueService(value: string) {
    this.id_servicio = value;
  }

  fetchServiceUser(idUnidad: string, idUsuario: string) {
    this._serviciosUsuarios.serviciosUsuarios(idUnidad, idUsuario).subscribe(res => {
      console.log(res.Data);
      this.serviciosUsuarioResul = res.Data;
    });
  }

  addArea() {
    const datos = <AreaServicioUnidades> this.formularioAgregar.getRawValue();
    const floa = Number(this.formularioAgregar.controls['CentroCosto']);
    datos.CentroCosto = floa;
    console.log(datos);
    this._areaServicios.AgregarAreaServicio(datos).subscribe(res => {
      console.log(res.Data);
      Swal.fire({
        position: 'center',
        title: res.Mensaje,
        icon: 'success',
      });
      this.fetchAreaService();
      this.formularioVisible = false;
      this.formularioAgregar.reset();
    });
  }
}
