import { Component, ComponentFactoryResolver, ComponentRef, OnDestroy, OnInit, ViewChild, ViewContainerRef } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { ContratoService } from '../../core/servicios/ContratosServicios/contrato.service';
import { Contrato } from '../../core/models/contrato.interface';
import { ConceptosContratoService } from '../../core/servicios/ContratosServicios/conceptos-contrato.service';
import { MatDialog } from '@angular/material/dialog';
import { AbrirCatalogoSaiService } from '../../core/servicios/abrir-catalogo-sai.service';
import { DistribucionConceptosComponent } from '../../shared/Dialogs/Contratos/distribucion-conceptos/distribucion-conceptos.component';
import { EditarConceptoContratoComponent } from '../../shared/Dialogs/Contratos/editar-concepto-contrato/editar-concepto-contrato.component';
import { RepresentanteLegalService } from '../../core/servicios/ContratosServicios/representante-legal.service';
import { ProveedoresInterface } from '../../core/models/proveedores.interface';
import { DetalleProveedorComponent } from '../../shared/Dialogs/detalle-proveedor/detalle-proveedor.component';
import { DetalleContacContrComponent } from '../../shared/Dialogs/Contratos/detalle-contac-contr/detalle-contac-contr.component';
import Swal from 'sweetalert2';
import { RespuestaPeticion } from '../../core/models/estructuras_generales';
import { PenasService } from '../../core/servicios/ContratosServicios/penas.service';
import { FianzaService } from '../../core/servicios/ContratosServicios/fianza.service';
import { AdministradorService } from '../../core/servicios/ContratosServicios/administrador.service';
import { TrasladoPacientesComponent } from '../ConceptosServicios/traslado-pacientes/traslado-pacientes.component';
import { SubTipoService } from '../../core/servicios/ContratosServicios/sub-tipo.service';
import { SubTipo, SubTipoContrato } from '../../core/models/sub_tipo.interface';
import { SubTipoContratoService } from '../../core/servicios/ContratosServicios/sub-tipo-contrato.service';
import { EquipoMedicoComponent } from '../ConceptosServicios/equipo-medico/equipo-medico.component';
import { EstudiosLaboratorioComponent } from '../ConceptosServicios/estudios-laboratorio/estudios-laboratorio.component';

@Component({
  selector: 'app-detalle-contrato',
  templateUrl: './detalle-contrato.component.html',
  styleUrls: ['./detalle-contrato.component.css'],
})
export class DetalleContratoComponent implements OnInit, OnDestroy {
  @ViewChild('conceptos', { read: ViewContainerRef }) container!: any;
  componentRef!: ComponentRef<any>;
  contrato_id: any;
  contrato: Contrato[] = [];
  btn_agregar = false;
  dataSource: any;
  displayedColumns: string[] = ['id_prei', 'clave_cuadro', 'descripcion_prei', 'marca', 'modelo', 'fecha_max', 'precio', 'distribucion',
    'EditarPro', 'EliminarPro'];
  representante: any;
  estado_deductivas!: boolean;
  estado_convencionales!: boolean;
  estado_administradores!: boolean;
  estado_fianza!: boolean;
  lista_sub_tipos: SubTipo[] = [];
  sub_tipo_contr!: SubTipoContrato;
  tipo_contrato = '';

  constructor(
    private _route: ActivatedRoute,
    private _contrato: ContratoService,
    private router: Router,
    private _conceptos: ConceptosContratoService,
    public dialog: MatDialog,
    private _representante: RepresentanteLegalService,
    private _activar_btn_cancelar: AbrirCatalogoSaiService,
    private _penas: PenasService,
    private _fianza: FianzaService,
    private _administrador: AdministradorService,
    private _sub_tipo: SubTipoService,
    private _sub_tipo_contrato: SubTipoContratoService,
    private resolver: ComponentFactoryResolver
  ) {

  }

  ngOnInit() {
    this.contrato_id = this._route.snapshot.paramMap.get('contrato_id');
    this.consultarContrato();
    this._activar_btn_cancelar.setEstadoBtn(true);
    this._activar_btn_cancelar.setRuta('/contratos/administrar');
    this.consultarRepresentante();
    this.consultarPenasDeductivas();
    this.consultarPenasConvencionales();
    this.consultarAdministradores();
    this.consultarFianza();
  }

  ngOnDestroy(): void {
    this._activar_btn_cancelar.setEstadoBtn(false);
    this._activar_btn_cancelar.setRuta('/inicio');
  }

  /* ******  Funcion para crear un componente dinamico dependiendo del tipo de contrato  ******   */
  crearComponente(id: any) {
    switch (id) {
      case '5f587535d7be0a74bf7db093':
        const factory = this.resolver.resolveComponentFactory(TrasladoPacientesComponent);
        this.componentRef = this.container.createComponent(factory);
        this.componentRef.instance.data = this.contrato;
        this.componentRef.instance.tipo = this.tipo_contrato;
        this.componentRef.instance.subTipo = this.sub_tipo_contr.SubTipoId;
        console.log('It is a Sunday.');
        break;
      case '5f323c89d7be0a421a2305ff':
        const factory1 = this.resolver.resolveComponentFactory(EquipoMedicoComponent);
        this.componentRef = this.container.createComponent(factory1);
        break;
      case '5f587730d7be0a790dbbbe60':
        const factory2 = this.resolver.resolveComponentFactory(EstudiosLaboratorioComponent)
        this.componentRef = this.container.createComponent(factory2);
        this.componentRef.instance.data = this.contrato;
        this.componentRef.instance.tipo = this.tipo_contrato;
        this.componentRef.instance.subTipo = this.sub_tipo_contr.SubTipoId;
        break;
      default:
        console.log('No such day exists!');
        break;
    }
  }

  /* ******  Funcion para consultar los subtipos de contratos  ******   */
  consultarSubTipo(tipo_contrato: any) {
    this._sub_tipo.ConsultaSubTipo(tipo_contrato).subscribe(res => {
      this.lista_sub_tipos = res.Data;
      if (this.lista_sub_tipos.length > 0) {
        console.log('Si hay registros');
        tipo_contrato = '';
        this.consultarSubTipoContrato(this.contrato[0].NumeroContrato);
      } else {
        this.crearComponente(tipo_contrato);
        // this.consultarConceptos();
      }
    }, err => {
      if (err.error !== '' || err.error !== undefined) {
        const erores = <RespuestaPeticion>err.error;
        console.log(erores);
      }
    });
  }

  /* ******  Funcion para consultar que subtipo es un contrato si es que tiene un subtipo  ******   */
  consultarSubTipoContrato(numero_contrato: any) {
    this._sub_tipo_contrato.ConsultaSubTipoContrato(numero_contrato).subscribe(res => {
      this.sub_tipo_contr = res.Data;
      console.log(this.sub_tipo_contr);
      this.crearComponente(this.sub_tipo_contr.SubTipoId);
      // this.consultarConceptos();
    }, err => {
      if (err.error !== '' || err.error !== undefined) {
        const erores = <RespuestaPeticion>err.error;
        console.log(erores);
      }
    });
  }

  /********    Funcion para consultar penas deductivas   ********/
  consultarPenasDeductivas() {
    this._penas.ConsultarPenaDeductiva(this.contrato_id).subscribe(resultado => {
      this.estado_deductivas = resultado.Data['estado'];
    });
  }

  /* ******  Funcion para consultar administradores o auxiliares  ******   */
  consultarAdministradores() {
    this._administrador.ConsultarAdministradores(this.contrato_id).subscribe(res => {
      this.estado_administradores = res.Data['estado'];
    });
  }

  /* ******  Funcion para consultar fianza  ******   */
  consultarFianza() {
    this._fianza.Consultarfianza(this.contrato_id).subscribe(res => {
      this.estado_fianza = res.Data['estado'];
    }, err => {
      if (err.error !== '' || err.error !== undefined) {
        const erores = <RespuestaPeticion>err.error;
        this.estado_fianza = erores.Data['estado'];
      }
    });
  }

  consultarPenasConvencionales() {
    this._penas.ConsultarPenaConvencional(this.contrato_id).subscribe(resultado => {
      this.estado_convencionales = resultado.Data['estado'];
    }, err => {
      if (err.error !== '' || err.error !== undefined) {
        const erores = <RespuestaPeticion>err.error;
        this.estado_convencionales = erores.Data['estado'];
      }
    });
  }

  consultarRepresentante() {
    this._representante.ConsultaRepresentante(this.contrato_id).subscribe(resul => {
      this.representante = resul.Data;
    });
  }


  consultarContrato() {
    this._contrato.ConsultarContratos(this.contrato_id, '', '', '', '', '', '', '',
      '', '').subscribe(resultado => {
        this.contrato = resultado.Data['registros'];
        console.log('Este es el contrato', this.contrato);
        this.tipo_contrato = this.contrato[0].TipoContratoId;
        this.consultarSubTipo(this.contrato[0].TipoContratoId);
      });
  }

  regresaDescripcion(valor: string): string {
    const descripcion = valor.split('.', 1);
    return descripcion.toString();
  }

  /*  consultarConceptos() {
      const id_contrato = this.contrato[0].NumeroContrato;
      const tipo = this.tipo_contrato !== null ? this.tipo_contrato : '';
      const sub_tipo = this.sub_tipo_contr.SubTipoId !== null ? this.sub_tipo_contr.SubTipoId : '';
      console.log('sdfsdf-*-***', tipo, sub_tipo);
      this._conceptos.ConsultarConceptos(id_contrato, tipo, sub_tipo).subscribe(resultado => {
        console.log(resultado.Data);
        this.dataSource = new MatTableDataSource();
        this.dataSource.data = resultado.Data['registros'];
        this.MontoContratado = resultado.Data['total'];
      });
    }*/

  estadoDistribucion(distribucion: any) {
    if (distribucion.CantidadDistribuida === 0) {
      return '#c42534';
    }
    if (distribucion.CantidadDistribuida === distribucion.CantidadConcepto) {
      return '#18903f';
    }
    if (distribucion.CantidadDistribuida < distribucion.CantidadConcepto) {
      return '#d4782c';
    }
  }

  openDialogDistribucion(concepto: any): void {
    const dialogRef = this.dialog.open(DistribucionConceptosComponent, {
      height: 'auto',
      width: '90%',
      disableClose: true,
      data: {
        concepto: concepto,
        cantidad: 0,
        archivo: false
      }
    });
    dialogRef.afterClosed().subscribe(result => {
      // this.consultarConceptos();
    });
  }

  openDialogActualizar(concepto: any): void {
    const dialogRef = this.dialog.open(EditarConceptoContratoComponent, {
      width: '90%',
      height: 'auto',
      disableClose: true,
      data: { conceptoEditar: concepto }
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result === true) {
        // this.consultarConceptos();
      }
    });
  }

  enviarRuta() {
    this.router.navigate(['/contratos/agregar', this.contrato_id]);
  }

  openDialogDetalleProveedor(provee: ProveedoresInterface): void {
    const dialogRef = this.dialog.open(DetalleProveedorComponent, {
      height: 'auto',
      width: '30%',
      disableClose: false,
      data: {
        proveedor: provee,
      }
    });

    dialogRef.afterClosed().subscribe(result => {
    });
  }

  openDialogContacto(): void {
    const dialogRef = this.dialog.open(DetalleContacContrComponent, {
      width: '65%',
      height: 'auto',
      disableClose: true,
      data: {
        n_contrato: this.contrato_id,
        representante: this.representante
      }
    });

    dialogRef.afterClosed().subscribe(result => {
    });
  }

  EliminarConcepto(id: string) {
    Swal.fire({
      title: 'Eliminar',
      text: 'Estas seguro que deseas eliminar este concepto',
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      confirmButtonText: 'Si, Eliminar',
      cancelButtonText: 'Cancelar'
    }).then((result) => {
      if (result.value) {
        this._conceptos.EliminarConcepto(id).subscribe(res => {
          Swal.fire(
            'Correcto',
            res.Mensaje,
            'success'
          );
          // this.consultarConceptos();
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

  /* ******  Funcion para direccionar a la vista de fianza  ******   */
  abrirFianza() {
    this.router.navigate(['/contratos/fianza', this.contrato_id]);
  }
}
