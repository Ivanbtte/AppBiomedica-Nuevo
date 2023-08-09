import {Component, OnInit, ViewChild} from '@angular/core';
import {MatTableDataSource} from '@angular/material/table';
import {MatPaginator} from '@angular/material/paginator';
import {ContratoService} from '../../core/servicios/ContratosServicios/contrato.service';
import {MatDialog} from '@angular/material/dialog';
import {AgregarEditarContratoComponent} from '../../shared/Dialogs/Contratos/agregar-editar-contrato/agregar-editar-contrato.component';
import {Router} from '@angular/router';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';
import {CargaMasivaContratosComponent} from '../../shared/Dialogs/Contratos/carga-masiva-contratos/carga-masiva-contratos.component';
import {EnviarArchivoConceptComponent} from '../../shared/Dialogs/Contratos/carga-masiva-conceptos/carga-masiva-conceptos.component';
import {CargaMasivaDistribucionComponent} from '../../shared/Dialogs/Contratos/carga-masiva-distribucion/carga-masiva-distribucion.component';
import {ProveedoresInterface} from '../../core/models/proveedores.interface';
import {DetalleProveedorComponent} from '../../shared/Dialogs/detalle-proveedor/detalle-proveedor.component';
import {FiltrosContratosComponent} from '../../shared/Dialogs/Contratos/filtros-contratos/filtros-contratos.component';
import {ConsultarContratoComponent} from '../../shared/Dialogs/Contratos/consultar-contrato/consultar-contrato.component';
import Swal from 'sweetalert2';
import {AgregarUnContratoComponent} from '../../shared/Dialogs/Contratos/agregar-un-contrato/agregar-un-contrato.component';
import {ContratoPuestoService} from '../../core/servicios/ContratosServicios/contrato-puesto.service';
import jwt_decode from 'jwt-decode';
import {ArchivosContratoComponent} from '../../shared/Dialogs/Contratos/archivos-contrato/archivos-contrato.component';


@Component({
  selector: 'app-administrar-contratos',
  templateUrl: './administrar-contratos.component.html',
  styleUrls: ['./administrar-contratos.component.css']
})
export class AdministrarContratosComponent implements OnInit {
  dataSource: any;
  elementos: any[] = [];
  uploadForm!: FormGroup;
  cargando = false;
  @ViewChild(MatPaginator, {static: true}) paginator!: MatPaginator;
  tamanioPagina = 15;
  tamanios = [15, 30, 45];
  totalRegistros = 0;
  paginaActual = 1;
  control_nombre = new FormControl();
  menu_distribucion!: boolean;
  /*variables para los filtros*/
  selectable = true;
  removable = true;
  proveedor: any[] = [];
  delegacion: any[] = [];
  unidad_med: any;
  nombre = '';
  n_prov_imss = '';
  delegaciones_filtro = '';
  tipo_contrato = '';

  constructor(
    private _contrato: ContratoService,
    public dialog: MatDialog,
    private router: Router,
    private formBuilder: FormBuilder,
    private contrato_puesto: ContratoPuestoService
  ) {
  }

  ngOnInit() {
    const decode: any = jwt_decode(<string> localStorage.getItem('token'));
    this.contrato_puesto.Consultarcontrato_puesto(decode.PuestoId).subscribe(res => {
      const ids = res.Data.map((it:any) => it.TipoContrato.Id);
      this.tipo_contrato = ids.join();
      console.log(this.tipo_contrato);
      for (let item of ids) {
        if (item === '5f323c89d7be0a421a2305ff') {
          this.menu_distribucion = true;
        }
      }
      this.consultarContratos();
    });
    this.uploadForm = this.formBuilder.group({
      profile: ['']
    });
    this.paginator._intl.itemsPerPageLabel = 'Resultados por pagina';
    this.paginator.page.subscribe(page => {
      if (page) {
        this.paginaActual = page['pageIndex'] + 1;
        this.tamanioPagina = page['pageSize'];
        this.consultarContratos();
      }
    });
  }

  consultarContratos() {
    this.n_prov_imss = '';
    this.delegaciones_filtro = '';
    const nombre = this.control_nombre.value !== null ? this.control_nombre.value : '';
    if (this.proveedor.length > 0) {
      const n_prov = this.proveedor.map(it => it.NProvImss);
      this.n_prov_imss = n_prov.join();
    }
    if (this.delegacion.length > 0) {
      const dele = this.delegacion.map(it => it.ClvDele);
      this.delegaciones_filtro = dele.join();
    }
    this._contrato.ConsultarContratos('', this.tipo_contrato, nombre, this.n_prov_imss, this.delegaciones_filtro, '', '', '', this.paginaActual.toString(),
      this.tamanioPagina.toString()).subscribe(resultado => {
      this.dataSource = new MatTableDataSource();
      this.dataSource.data = resultado.Data['registros'];
      this.elementos = this.dataSource.data;
      this.totalRegistros = resultado.Data['total'];
      this.cargando = true;
      if (this.elementos.length === 0) {
        Swal.fire({
          icon: 'error',
          title: 'Sin Resultados',
          text: 'No se encontraron resultados para tu busqueda',
        });
      }
    });
  }

  buscarPorNombre() {
    this.consultarContratos();
    this.paginator.firstPage();
  }

  openDialogAgregar(): void {
    const dialogRef = this.dialog.open(AgregarEditarContratoComponent, {
      width: '95%',
      disableClose: true,
      data: {
        editar: false,
        contrato: ''
      }
    });
    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.consultarContratos();
      }
    });
  }

  editando(contrato: any) {
    const dialogRef = this.dialog.open(AgregarUnContratoComponent, {
      height: 'auto',
      width: '90%',
      disableClose: true,
      data: {
        contrato: contrato,
        representante: '',
        proveedor: contrato.Proveedor,
        editar: true,
      }
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.consultarContratos();
      }
    });
  }

  /* ******  Funcion para abrir la vista de archivos de un contrato  ******   */
  AbrirDialogArchivos(contrato: any) {
    const dialogRef = this.dialog.open(ArchivosContratoComponent, {
      height: 'auto',
      width: '65%',
      disableClose: true,
      data: {
        contrato: contrato,
      }
    });

    dialogRef.afterClosed().subscribe(result => {
    });
  }
// ********************************************************************************************************
  /* *****  Funcion para abrir la ventana de Carga Masiva de Contratos  *****   */
  openDialogCargaMasiva(): void {
    const dialogRef = this.dialog.open(CargaMasivaContratosComponent, {
      height: 'auto',
      width: '90%',
      disableClose: true,

    });

    dialogRef.afterClosed().subscribe(result => {
      if (result === true) {
        this.consultarContratos();
      }
    });
  }

// ********************************************************************************************************
  /* *****  Funcion para abrir el componente para subir un archivo para carga masiva de conceptos  *****   */
  openDialogConceptosArchivo(): void {
    const dialogRef = this.dialog.open(EnviarArchivoConceptComponent, {
      height: 'auto',
      width: '90%',
      disableClose: true,

    });

    dialogRef.afterClosed().subscribe(result => {
      if (result === true) {
        this.consultarContratos();
      }
    });
  }

  openDialogDistribucionArchivo(conceptos: boolean): void {
    const dialogRef = this.dialog.open(CargaMasivaDistribucionComponent, {
      height: 'auto',
      width: '90%',
      disableClose: true,
      data: {conceptos: conceptos}
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result === true) {
        this.consultarContratos();
      }
    });
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

  // ********************************************************************************************************
  /* *****  Funcion para abrir la ventana de filtros  *****   */
  AbrirVentanaFiltros(): void {
    const dialogRef = this.dialog.open(FiltrosContratosComponent, {
        height: 'auto',
        width: '75%',
        disableClose: true,
        data: {
          proveedor: this.proveedor,
          delegacion: this.delegacion,
          unidad_med: '',
          nombreAutoriza: this.nombre
        }
      })
    ;
    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.proveedor = result.proveedor;
        this.nombre = result.nombre;
        this.delegacion = result.delegacion;
        this.consultarContratos();
      }
    });
  }

  // ********************************************************************************************************
  /* *****  Funcion para eliminar los filtros  *****   */
  EliminarFiltro(chip: string) {
    if (chip === 'proveedor') {
      this.proveedor = [];
      this.n_prov_imss = '';
      this.nombre = '';
    }
    if (chip === 'delegacion') {
      this.delegaciones_filtro = '';
      this.delegacion = [];
    }
    this.consultarContratos();
  }

  AbrirDetallecontrato(row: any) {
    const dialogRef = this.dialog.open(ConsultarContratoComponent, {
      height: 'auto',
      width: '50%',
      disableClose: true,
      data: {
        num_contrato: row,
      }
    });
    dialogRef.afterClosed().subscribe(result => {
    });
  }
}
