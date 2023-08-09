import { Component, OnInit, ViewEncapsulation } from "@angular/core";
import { UsuarioService } from "../../../authentication/servicios/usuario.service";
import { MatTableDataSource } from "@angular/material/table";
import Swal from "sweetalert2";
import {
  BitacoraUsuarioPuesto,
  UnsuscribeUser,
} from "../../../core/models/bitacora_usuario_puesto.interface";
import { ServicioUsuarioPuestoService } from "../../../core/servicios/ServiciosUsuarios/servicio-usuario-puesto.service";
import { MatDialog } from "@angular/material/dialog";
import { DialogDeleteUserComponent } from "../dialogs/dialog-delete-user/dialog-delete-user.component";
import { B, I } from "@angular/cdk/keycodes";
import { Justificacion } from "../dialogs/dialog-autoriza-acomp/dialog-autoriza-acomp.component";
import { RespuestaPeticion } from "src/app/core/models/estructuras_generales";

@Component({
  selector: "app-transfer-manager",
  templateUrl: "./transfer-manager.component.html",
  styleUrls: ["./transfer-manager.component.scss"],
  encapsulation: ViewEncapsulation.None,
})
export class TransferManagerComponent implements OnInit {
  unitMedic: any;
  
  // displayedColumns: string[] = [
  //   "config",
  //   "nameTras",
  //   "positionTras",
  //   "departmentTras",
  //   "unitMedic",
  //   "action",
  // ];
  displayedColumns: string[] = [
    "nameTras",
    "positionTras",
    "departmentTras",
    "unitMedic",
    "action",
  ];
  dataSource: MatTableDataSource<any>;
  userId: any;
  constructor(
    private usuarioService: UsuarioService,
    private usaurioP: ServicioUsuarioPuestoService,
    public dialog: MatDialog
  ) {
    this.unitMedic = JSON.parse(<string>localStorage.getItem("unit"));
    console.log(this.unitMedic);
    this.userId = JSON.parse(localStorage.getItem("user"));
    this.consultUsers(this.userId.Usuario.Id);
    console.log(this.userId.Usuario.Id);
  }

  consultUsers(userId: string) {
    this.usuarioService
      .ConsultUsersByDepartment(this.unitMedic.id, userId)
      .subscribe((result) => {
        this.dataSource = new MatTableDataSource<any>(result.Data);
        console.log(result.Data);
      });
  }

  ngOnInit(): void {}

  async ValidateUser(data: any) {
    const bitacora: Partial<BitacoraUsuarioPuesto> = {
      usuarioModificadoId: data.user.Id,
      usuarioModificanteId: this.userId.Id,
      tipoModificacionId: "61f1a754d7be0a2b052e887e",
    };
    Swal.fire({
      title: "Validación de usuario",
      html: `Estas validando que: <b>${data.user.Usuario.Persona.Nombre} ${data.user.Usuario.Persona.ApellidoPat}.</b><br>
       Pertenece al departamento: <b>${data.user.PuestoOrganigrama.organigramaUnidad.departamento.Nombre}.</b><br>
       Con el puesto: <b>${data.user.PuestoOrganigrama.catalogoPuesto.Puesto}.</b>`,
      icon: "warning",
      showCancelButton: true,
      confirmButtonColor: "#3085d6",
      cancelButtonColor: "#d33",
      cancelButtonText: "Cancelar",
      reverseButtons: true,
      confirmButtonText: "Si, validar",
    }).then((result) => {
      if (result.isConfirmed) {
        this.example(bitacora);
      }
    });
  }

  async example(bitacora:Partial<BitacoraUsuarioPuesto>){
    console.log('nuevo swal');
    const { value: fruit } = await Swal.fire({
      title: 'Selecciona un rol',
      input: 'select',
      allowOutsideClick: false,
      inputOptions: {
        '614ca03ed7be0a0b8e27fca4': 'Capturista de vales de traslado'
      },
      inputPlaceholder: 'Selecciona',
      showCancelButton: true,
      inputValidator: (value) => {
        return new Promise((resolve) => {
          if (value !== '') {
            console.log('seleccionado', value)
            this.usaurioP.AddBitacoraUsuarioPuesto(bitacora, "true",value).subscribe(
              (result) => {
               this.consultUsers(this.userId.Usuario.Id);
                Swal.fire("Validado", "Este usuario ha sido validado.", "success");
              },
              (err) => {
                if (err.error !== "" || err.error !== undefined) {
                  Swal.fire({
                    position: "center",
                    title: "Error",
                    text: "No se pudo descargar la plantilla",
                    icon: "error",
                    showConfirmButton: true,
                  });
                }
              }
            );
            resolve('')
          } else {
            resolve('Necesitas elegi un puesto')
          }
        })
      }
    })
  }
  openDialog(value: any) {
    console.log(value);
    console.log(value.user.Usuario.Id, 'weeeee');
    const dialogRef = this.dialog.open(DialogDeleteUserComponent, {
      height: "auto",
      width: "55%",
      disableClose: true,
    });
    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        const bitacora: Partial<BitacoraUsuarioPuesto> = {
          usuarioModificadoId: value.user.Id,
          usuarioModificanteId: this.userId.Id,
          tipoModificacionId: result.tipoModificacion.id,
          comentario: result.justificacion,
        };
        this.usaurioP
          .UnsuscribeUsuarioPuesto(bitacora, value.user.Usuario.Id)
          .subscribe(
            (result) => {
              console.log(result.Data);
              Swal.fire(
                "Correcto",
                "Este usuario se ha dado de baja.",
                "success"
              );
              this.consultUsers(this.userId.Usuario.Id);
            },
            (err) => {
              if (err.error !== "" || err.error !== undefined) {
                const erores = <RespuestaPeticion>err.error;
                Swal.fire({
                  position: "center",
                  title: erores.Mensaje,
                  icon: "error",
                  showConfirmButton: true,
                });
              }
            }
          );
      }
    });
  }
}
