import {Component, OnInit} from '@angular/core';
import {UsuarioService} from '../../authentication/servicios/usuario.service';
import {MatTableDataSource} from '@angular/material/table';
import {MatDialog} from '@angular/material/dialog';
import {DetalleUsuarioComponent} from '../detalle-usuario/detalle-usuario.component';
import Swal from 'sweetalert2';
import {RespuestaPeticion} from '../../core/models/estructuras_generales';
import {UserPositionsService} from '../../core/servicios/userPositions/user-positions.service';
import {UserPositions} from '../../Utils/UserPositions';

@Component({
  selector: 'app-validar-usuarios',
  templateUrl: './validar-usuarios.component.html',
  styleUrls: ['./validar-usuarios.component.scss']
})
export class ValidarUsuariosComponent implements OnInit {
  displayedColumns: string[] = ['nameUser', 'departmentUser', 'cargoUser', 'detalleUser', 'validarUser', 'rechazarUser'];
  dataSource: any;
  resultado: any[] = [];
  userValidator: any;

  constructor(
    private _usuarios: UsuarioService,
    public dialog: MatDialog,
    private userPositionsService: UserPositionsService
  ) {
    this.userValidator = JSON.parse(<string> localStorage.getItem('user'));
    console.log(this.userValidator);
  }

  ngOnInit() {
    if (this.userValidator.PuestoOrganigrama.organigramaUnidadId) {
      this.consultUsersValid(this.userValidator.PuestoOrganigrama.organigramaUnidad.unidadMedId, '', this.userValidator.PuestoOrganigrama.organigramaUnidad.id);
    }
    if (this.userValidator.PuestoOrganigrama.organigramaDelegacionId) {
      this.consultUsersValid('', this.userValidator.PuestoOrganigrama.organigramaDelegacion.delegacionesClvDele, this.userValidator.PuestoOrganigrama.organigramaDelegacion.id);
    }
  }

  returnDepartment(row: any): any {
    return UserPositions.returnDepartment(row);
  }

  openDialogClave(user: any): void {
    const dialogRef = this.dialog.open(DetalleUsuarioComponent, {
      height: 'auto',
      width: '45%',
      data: {
        user: user,
      }
    });
    dialogRef.afterClosed().subscribe(result => {
    });
  }

  consultarUsuarios() {
    this._usuarios.ConsultarUsuarios('false').subscribe(res => {
      this.resultado = res.Data;
      this.dataSource = new MatTableDataSource();
      this.dataSource.data = this.resultado;
      console.log(this.resultado);
    });
  }

  consultUsersValid(unitId, clvDele, departmentId: string) {
    this.userPositionsService.consultUsersValid('true', 'false', unitId, clvDele, departmentId).subscribe(result => {
      this.resultado = result.Data;
      this.dataSource = new MatTableDataSource();
      this.dataSource.data = this.resultado;
      console.log(result.Data);
    });
  }

  validarUsuario(Id: string) {
    Swal.fire({
      title: 'Validación',
      text: 'Estás seguro de validar este usuario',
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      cancelButtonText: 'Cancelar',
      confirmButtonText: 'Si, validar'
    }).then((result) => {
      if (result.value) {
        this._usuarios.ValidarUsuario(Id, 'true', {}).subscribe(res => {
          Swal.fire(
            'Validado',
            'Este usuario ha sido validado',
            'success'
          );
          this.consultarUsuarios();
        });
      }
    });
  }

  async rechazarUsuario(clave: any) {
    console.log(clave);
    const {value: comentario} = await Swal.fire({
      icon: 'warning',
      title: '¿Estás seguro?',
      text: '¡No podrás revertir esto!',
      input: 'text',
      inputPlaceholder: 'Escribe el motivo',
      showCancelButton: true,
      cancelButtonText: 'Cancelar',
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      confirmButtonText: 'Si, rechazar',
      preConfirm: (value) => {
        if (!value) {
          return 'Necesitas escribir un comentario';
        }
      }
    });
    if (comentario) {
      this._usuarios.rechazarUsuario(clave.Id, clave.PersonaId, comentario).subscribe(res => {
        console.log(res.Data);
        Swal.fire({
          icon: 'success',
          title: 'Correcto',
          text: res.Mensaje,
          showConfirmButton: true,

        });
        this.consultarUsuarios();

      }, err => {
        if (err.error !== '' || err.error !== undefined) {
          const erores = <RespuestaPeticion> err.error;
          Swal.fire({
            icon: 'error',
            title: 'Error',
            text: erores.Mensaje,
            showConfirmButton: true,

          });
        }
      });
    }

  }
}
