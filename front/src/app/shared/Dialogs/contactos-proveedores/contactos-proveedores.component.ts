import {Component, OnInit, Inject, ViewChild} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialog, MatDialogRef} from '@angular/material/dialog';
import {ContactoProveedor, ProveedoresInterface} from '../../../core/models/proveedores.interface';
import {BreakpointObserver} from '@angular/cdk/layout';
import {MatTableDataSource} from '@angular/material/table';
import {DetalleContactoProveedorComponent} from '../detalle-contacto-proveedor/detalle-contacto-proveedor.component';
import Swal from 'sweetalert2';
import {AgregarContactoProveedorComponent} from '../agregar-contacto-proveedor/agregar-contacto-proveedor.component';
import {UsuarioLogeado} from '../../../core/models/UsuarioLogin.interface';
import jwt_decode from 'jwt-decode';
import {ContactoProveedoresService} from '../../../core/servicios/ProveedoresServicios/contacto-proveedores.service';


export interface DialogDataContacto {
  proveedor: ProveedoresInterface;
  seleccionar: boolean;
}

@Component({
  selector: 'app-contactos-proveedores',
  templateUrl: './contactos-proveedores.component.html',
  styleUrls: ['./contactos-proveedores.component.css']
})
export class ContactosProveedoresComponent implements OnInit {
  provee!: ProveedoresInterface;
  contactos: ContactoProveedor[]=[];
  displayedColumns: string[] = ['NombreCon', 'Departamento', 'CargoCon', 'CelularCon', 'CorreoCon', 'VerPro', 'EditarPro', 'EliminarPro'];
  dataSource: any;
  public usuarioLog!: UsuarioLogeado;
  permisos = true;
  titulo = 'Contactos del Proveedor';

  constructor(
    public dialogRef: MatDialogRef<ContactosProveedoresComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogDataContacto,
    private _contactos: ContactoProveedoresService,
    breakpointObserver: BreakpointObserver,
    public dialog: MatDialog,
  ) {
    breakpointObserver.observe(['(max-width: 600px)']).subscribe(result => {
      this.displayedColumns = result.matches ?
        ['NombreCon', 'Departamento', 'CargoCon', 'CelularCon', 'CorreoCon', 'VerPro', 'EditarPro', 'EliminarPro'] :
        ['NombreCon', 'Departamento', 'CargoCon', 'CelularCon', 'CorreoCon', 'VerPro', 'EditarPro', 'EliminarPro'];
    });
  }

  ngOnInit() {
    this.usuarioLog = jwt_decode(<string>localStorage.getItem('token'));
    this.provee = this.data.proveedor;
    this.llenarTabla();
    if (this.data.seleccionar) {
      this.displayedColumns.pop();
      this.displayedColumns.pop();
      this.displayedColumns.pop();
      this.displayedColumns.push('Seleccionar');
      this.titulo = 'Seleccionar Representante Legal';
    }
  }

  llenarTabla() {
    this._contactos.ConsultarContactoProveedor('', this.provee.NProvImss, '').subscribe(resultado => {
      this.contactos = resultado.Data;
      console.log(this.contactos);
      this.dataSource = new MatTableDataSource(this.contactos);
    });
  }

  applyFilter(filterValue: string) {
    this.dataSource.filter = filterValue.trim().toLowerCase();
  }

  cerrarDialog() {
    this.dialogRef.close();
  }

  openDialogDetalleContacto(contacto: ContactoProveedor): void {
    const dialogRef = this.dialog.open(DetalleContactoProveedorComponent, {
      height: 'auto',
      width: '35%',
      disableClose: false,
      data: {
        contacto: contacto,
      }
    });

    dialogRef.afterClosed().subscribe(result => {
    });
  }

  EliminarContacto(id_contacto: string) {
    Swal.fire({
      title: 'Estas seguro',
      text: 'Deseas eliminar este contacto',
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      confirmButtonText: 'Si Eliminar'
    }).then((result) => {
      if (result.value) {
        this._contactos.EliminarContacto(id_contacto).subscribe(resultado => {
          console.log(resultado);
          Swal.fire(
            'Eliminado',
            'El contacto ha sido eliminado correctamente',
            'success'
          );
          this.llenarTabla();
        });
      }
    });
  }

  Selecionar(contacto: ContactoProveedor) {
    Swal.fire({
      title: 'Seleccionar',
      text: 'Proveedor:' + this.provee.NombreEmpresa + 'Contacto:' + contacto.NombreCompleto,
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      confirmButtonText: 'Si Seleccionar'
    }).then((result) => {
      if (result.value) {
        this.dialogRef.close(contacto);
      }
    });
  }

  openDialogAgregar(): void {
    const dialogRef = this.dialog.open(AgregarContactoProveedorComponent, {
      height: 'auto',
      width: '80%',
      disableClose: true,
      data: {
        editar: false,
        contacto: '',
        proveedor_id: this.provee.NProvImss
      }
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.llenarTabla();
      }
    });
  }

  openDialogEditar(contac: ContactoProveedor): void {
    const dialogRef = this.dialog.open(AgregarContactoProveedorComponent, {
      height: 'auto',
      width: '80%',
      disableClose: true,
      data: {
        editar: true,
        contacto: contac,
        proveedor_id: this.provee.Id
      }
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.llenarTabla();
      }
    });
  }
}
