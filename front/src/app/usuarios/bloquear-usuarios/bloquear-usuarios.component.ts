import {Component, OnInit} from '@angular/core';
import {UsuarioService} from '../../authentication/servicios/usuario.service';
import {MatTableDataSource} from '@angular/material/table';
import {MatDialog} from '@angular/material/dialog';
import {DetalleUsuarioComponent} from '../detalle-usuario/detalle-usuario.component';
import Swal from 'sweetalert2';

@Component({
  selector: 'app-bloquear-usuarios',
  templateUrl: './bloquear-usuarios.component.html',
  styleUrls: ['./bloquear-usuarios.component.css']
})
export class BloquearUsuariosComponent implements OnInit {
  displayedColumns: string[] = ['IdUser', 'Nombre', 'Apellidos', 'Correo', 'Matricula', 'Detalle', 'Validar'];
  dataSource: any;
  resultado: any[]=[];

  constructor(
    private _usuarios: UsuarioService,
    public dialog: MatDialog,
  ) {
  }

  ngOnInit() {
    this.consultarUsuarios();
  }

  openDialogClave(id: string): void {
    const dialogRef = this.dialog.open(DetalleUsuarioComponent, {
      height: 'auto',
      width: 'auto',
      data: {
        id: id,
      }
    });

    dialogRef.afterClosed().subscribe(result => {
    });
  }

  consultarUsuarios() {
    this._usuarios.ConsultarUsuarios('true').subscribe(res => {
      this.resultado = res.Data;
      this.dataSource = new MatTableDataSource();
      this.dataSource.data = this.resultado;
      console.log(this.resultado);
    });
  }

  bloquearUsuario(Id: string) {
    Swal.fire({
      title: 'Bloqueo Temporal',
      text: 'Estás seguro de bloquear este usuario',
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      cancelButtonText: 'Cancelar',
      confirmButtonText: 'Si, bloquear'
    }).then((result) => {
      if (result.value) {
        this._usuarios.ValidarUsuario(Id, 'false', {}).subscribe(res => {
          Swal.fire(
            'Bloqueado',
            'Este usuario ha sido bloqueado',
            'success'
          );
          this.consultarUsuarios();
        });
      }
    });
  }
}
