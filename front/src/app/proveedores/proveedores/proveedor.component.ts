import {Component, OnInit, ViewChild} from '@angular/core';
import {MatPaginator} from '@angular/material/paginator';
import {MatSort} from '@angular/material/sort';
import {ProveedoresService} from '../../core/servicios/ProveedoresServicios/proveedores.service';
import {MatTableDataSource} from '@angular/material/table';
import {ProveedoresInterface} from '../../core/models/proveedores.interface';
import Swal from 'sweetalert2';
import {MatDialog} from '@angular/material/dialog';
import {AgregarProveedorComponent} from '../../shared/Dialogs/agregar-proveedor/agregar-proveedor.component';
import {ContactosProveedoresComponent} from '../../shared/Dialogs/contactos-proveedores/contactos-proveedores.component';
import {DetalleProveedorComponent} from '../../shared/Dialogs/detalle-proveedor/detalle-proveedor.component';
import {UsuarioLogeado} from '../../core/models/UsuarioLogin.interface';
import jwt_decode from 'jwt-decode';
import {FormControl} from '@angular/forms';

@Component({
  selector: 'app-agregar-proveedor',
  templateUrl: './proveedor.component.html',
  styleUrls: ['./proveedor.component.css']
})
export class ProveedorComponent implements OnInit {
  // ********************************************************************************************************
  /* ***** Variables para la tabla  *****   */
  dataSource: any;
  @ViewChild(MatPaginator, {static:true}) paginator!: MatPaginator;
  @ViewChild(MatSort,{static:true}) sort!: MatSort;
  tamanioPagina = 50;
  tamanios = [20, 50, 100];
  totalRegistros = 0;
  paginaActual = 1;
  displayedColumns: string[] = ['DetallePro', 'Alias', 'Estado', 'CorreoPro', 'RFC',
    'Telefono', 'VerPro', 'EditarPro', 'EliminarPro'];
  cargando = false;
  // ********************************************************************************************************
  /* *****  Variables para el servicio de consultar proveedores  *****   */
  id_proveedor = '';
  control_nombre = new FormControl();
  public usuarioLog!: UsuarioLogeado;
  permisos = true;

  constructor(
    private _proveedores: ProveedoresService,
    public dialog: MatDialog,
  ) {
  }

  ngOnInit() {
    this.usuarioLog = jwt_decode(<string>localStorage.getItem('token'));
    this.llenarTabla();
    this.paginator._intl.itemsPerPageLabel = 'Resultados por pagina';
    this.paginator.page.subscribe(page => {
      if (page) {
        this.paginaActual = page['pageIndex'] + 1;
        this.tamanioPagina = page['pageSize'];
        this.llenarTabla();
      }
    });
  }

  llenarTabla() {
    const nombre = this.control_nombre.value !== null ? this.control_nombre.value : '';
    this._proveedores.ConsultarProveedores(this.id_proveedor, nombre, this.paginaActual.toString(),
      this.tamanioPagina.toString()).subscribe(resultado => {
      const arregloProv: ProveedoresInterface = resultado.Data['registros'];
      this.dataSource = new MatTableDataSource();
      this.dataSource.data = arregloProv;
      this.totalRegistros = resultado.Data['total'];
      this.cargando = true;
      console.log(arregloProv);
    });
  }

  openDialogAgregar(): void {
    const dialogRef = this.dialog.open(AgregarProveedorComponent, {
      height: 'auto',
      width: '80%',
      disableClose: true,
      data: {
        editar: false,
        proveedor: ''
      }
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.llenarTabla();
      }
    });
  }

  openDialogEditar(provee: ProveedoresInterface): void {
    const dialogRef = this.dialog.open(AgregarProveedorComponent, {
      height: 'auto',
      width: '80%',
      disableClose: true,
      data: {
        editar: true,
        proveedor: provee
      }
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.llenarTabla();
      }
    });
  }

  openDialogVer(id: ProveedoresInterface): void {
    const dialogRef = this.dialog.open(ContactosProveedoresComponent, {
      height: 'auto',
      width: '80%',
      disableClose: true,
      data: {
        proveedor: id,
        seleccionar: false
      }
    });

    dialogRef.afterClosed().subscribe(result => {
    });
  }

  openDialogDetalleProveedor(provee: ProveedoresInterface): void {
    const dialogRef = this.dialog.open(DetalleProveedorComponent, {
      height: 'auto',
      width: '35%',
      disableClose: false,
      data: {
        proveedor: provee,
      }
    });

    dialogRef.afterClosed().subscribe(result => {
    });
  }

  EliminarProveedor(id_proveedor: string) {
    Swal.fire({
      title: 'Estas seguro',
      text: 'Deseas eliminar este proveedor',
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      confirmButtonText: 'Si Eliminar'
    }).then((result) => {
      if (result.value) {
        this._proveedores.EliminarProveedor(id_proveedor).subscribe(resultado => {
          console.log(resultado);
          Swal.fire(
            'Eliminado',
            'El proveedor ha sido eliminado correctamente',
            'success'
          );
          this.llenarTabla();
        });
      }
    });
  }

  buscarPorNombre() {
    this.llenarTabla();
    this.paginator.firstPage();
  }
}
