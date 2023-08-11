import { Component, OnInit } from "@angular/core";
import { DepartmentTreeComponent } from "src/app/authentication/department-tree/department-tree.component";
import { AdministradoresTrasladosService } from "../../core/administradores-traslados.service";
import { MatDialog } from "@angular/material/dialog";
export interface AdminTraslado {
  unidad: any;
  datos: any;
  usuario: any;
  estado: boolean;
}

@Component({
  selector: "app-administrar-usuarios",
  templateUrl: "./administrar-usuarios.component.html",
  styleUrls: ["./administrar-usuarios.component.scss"],
})

export class AdministrarUsuariosComponent implements OnInit {
  displayedColumns: string[] = [
    "accion",
    "unidad",
    "departamento",
    "puesto",
    "nombre"
  ];
  dataSource: any[];
  user: any;
  constructor(
    private administradores: AdministradoresTrasladosService,
    private dialog: MatDialog
  ) {
    this.user = JSON.parse(<string>localStorage.getItem("user"));
    this.consultarAdministradores();
  }

  ngOnInit(): void { }
  consultarAdministradores() {
    this.administradores.ConsultAdministradoresUnidad(this.user.PuestoOrganigrama.organigramaDelegacion.delegaciones.ClvDele, "614c9be1d7be0a7dcda60e33")
      .subscribe((result) => {
        const temp: AdminTraslado[] = result.Data.sort((a, b) => Number(b.estado) - Number(a.estado));
        console.log(result)
        const datosTabla = temp.map(item => {
          return {
            estado: item.estado,
            unidad: item.unidad.DenominacionUni.trim(),
            unidadId: item.unidad.Id,
            departamento: Object.keys(item.datos).length !== 0 ? item.datos.puestoOrganigrama.organigramaUnidad.departamento.Nombre : '',
            puesto: Object.keys(item.datos).length !== 0 ? item.datos.puestoOrganigrama.catalogoPuesto.Puesto : '',
            usuario: item.usuario.Id != '' ? item.usuario.Usuario.Persona.Nombre +
              " " + item.usuario.Usuario.Persona.ApellidoPat + " " + item.usuario.Usuario.Persona.ApellidoMat : '',
            idDepartamento: Object.keys(item.datos).length !== 0 ? item.datos.puestoOrganigrama.organigramaUnidad.departamento.Id : '',
            idPuesto: Object.keys(item.datos).length !== 0 ? item.datos.puestoOrganigrama.catalogoPuesto.Id : '',
          }
        })
        this.dataSource = datosTabla
      });
  }

  openDialogDepartment(item: any) {
    console.log(item)
    let clvDele = "";
    let Dele = "";
    let unitId = 0;
    let unit = "";
    unitId = item.unidadId;
    const dialogRef = this.dialog.open(DepartmentTreeComponent, {
      height: "auto",
      width: "60%",
      disableClose: true,
      data: {
        unitMedId: unitId,
        UnitMed: unit,
        clvDelegation: clvDele,
        Delegation: Dele,
        puesto: true,
        idDepartamento: item.idDepartamento,
        idPuesto: item.idPuesto
      },
    });
    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        this.consultarAdministradores();
      }
    });
  }
}
