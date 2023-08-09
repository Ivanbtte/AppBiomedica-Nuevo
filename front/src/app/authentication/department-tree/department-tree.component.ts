import { Component, Inject, OnInit } from "@angular/core";
import { MAT_DIALOG_DATA, MatDialogRef } from "@angular/material/dialog";
import { FlatTreeControl } from "@angular/cdk/tree";
import {
  MatTreeFlatDataSource,
  MatTreeFlattener,
} from "@angular/material/tree";
import { DepartamentosService } from "../../core/servicios/DepartamentosSubdireccion/departamentos.service";
import Swal from "sweetalert2";
import { RespuestaPeticion } from "../../core/models/estructuras_generales";
import { PuestosService } from '../servicios/puestos.service';
import { CompanyPosition } from "../interfaces/company.position.interface";
import { FormControl } from '@angular/forms';
import { PuestoOrganigramaRole } from '../interfaces/puesto.role.interface';
import { Organigrama } from '../interfaces/departament.interface';
import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';

/* ******  interface para el dialog  ******   */
export interface DialogData {
  unitMedId: number;
  UnitMed: any;
  clvDelegation: string;
  Delegation: any;
  puesto: boolean;
  idDepartamento: string;
  idPuesto: string;
}

/* ******  elementos para el componente tree  ******   */

interface OrganizationUnit {
  id: string;
  nivelJerarquico: number;
  departamentoId: string;
  departamento: {
    Id: string;
    Nombre: string;
    DepartamentoTipoId: string;
    DepartamentoTipo: any;
    UsuarioTipoId: string;
    UsuarioTipo: any;
    Estado: boolean;
  };
  areaId: string;
  area: any;
  organigramaParent: string;
  organigramas: OrganizationUnit[];
  estado: boolean;
}

/** Flat node with expandable and level information */
interface ExampleFlatNode {
  expandable: boolean;
  name: string;
  areaId: string;
  id: string;
  level: number;
  selected: boolean;
  departamentoId: string;
}

@Component({
  selector: "app-department-tree",
  templateUrl: "./department-tree.component.html",
  styleUrls: ["./department-tree.component.css"],
})
export class DepartmentTreeComponent implements OnInit {
  mnsError = "";
  private _transformer = (node: OrganizationUnit, level: number) => {
    return {
      expandable: !!node.organigramas && node.organigramas.length > 0,
      name: node.departamento.Nombre,
      areaId: node.areaId,
      id: node.id,
      level: level,
      selected: false,
      departamentoId: node.departamentoId,
    };
  };
  treeControl = new FlatTreeControl<ExampleFlatNode>(
    (node) => node.level,
    (node) => node.expandable
  );
  treeFlattener = new MatTreeFlattener(
    this._transformer,
    (node) => node.level,
    (node) => node.expandable,
    (node) => node.organigramas
  );
  dataSource = new MatTreeFlatDataSource(this.treeControl, this.treeFlattener);
  positions: CompanyPosition[] = [];
  puestosBuscar = false;
  puesto = new FormControl()
  puestoRole: Partial<PuestoOrganigramaRole>
  idDepartamento = ''
  constructor(
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
    private department: DepartamentosService,
    public dialogRef: MatDialogRef<DepartmentTreeComponent>,
    private _puestos: PuestosService,

  ) {
    console.log(data)
    if (data.unitMedId != 0) {
      this.department.ConsultOrganizationUnitMed(data.unitMedId.toString(), "").subscribe(result => {
        this.dataSource.data = result.Data;
        console.log(this.dataSource.data)
        if (data.idDepartamento) {
          this.treeControl.collapseAll();
          this.expandir(this.dataSource.data,data.idDepartamento)
        } else {
          this.treeControl.expand(this.treeControl.dataNodes[0]);
        }
      }, (err) => {
        if (err.error !== "" || err.error !== undefined) {
          const errores = <RespuestaPeticion>err.error;
          this.mnsError = errores.Mensaje;
        }
      }
      );
    } else if (data.clvDelegation != "") {
      this.department.ConsultOrganizationUnitMed("0", data.clvDelegation).subscribe(result => {
        this.dataSource.data = result.Data;
        this.treeControl.expand(this.treeControl.dataNodes[0]);
      }, (err) => {
        if (err.error !== "" || err.error !== undefined) {
          const errores = <RespuestaPeticion>err.error;
          this.mnsError = errores.Mensaje;
        }
      }
      );
    } else {
      this.mnsError = "No se han registrado departamentos para este tipo de usuario: Nivel Central";
    }
  }
   
hasChild = (_: number, node: ExampleFlatNode) => node.expandable;

ngOnInit() {
  this.idDepartamento = this.data.idDepartamento
  console.log(this.idDepartamento)
}
expandir(data: OrganizationUnit[], idDepartamento: string): any {
  data.forEach(node => {
    console.log(node)
  },this);
}
selectedDepartment(node: ExampleFlatNode) {
  console.log(node)
  let mns = 'Mi departamento de adscipcion es:';
  if (this.data.puesto) {
    mns = 'El departamento es:'
  }
  Swal.fire({
    title: mns,
    text: node.name,
    icon: "warning",
    showCancelButton: true,
    confirmButtonColor: "#3085d6",
    cancelButtonColor: "#d33",
    cancelButtonText: "Seleccionar otro",
    confirmButtonText: "Aceptar",
  }).then((result) => {
    if (result.isConfirmed) {
      this.idDepartamento = node.departamentoId
      this.positions = [];
      this.puesto.reset();
      if (!this.data.puesto) {
        this.dialogRef.close(node);
      } else {
        this.buscarPuestos(node.id);
      }
    }
  });
}
buscarPuestos(id: any) {
  this._puestos.ConsultPositionOrganizations(id, "").subscribe((result) => {
    this.puestosBuscar = true;
    this.positions = result.Data;
  });
}
agregarAdministrador() {
  this.puestoRole = {
    puestoOrganigramaId: this.puesto.value.id,
    roleId: '614c9be1d7be0a7dcda60e33'
  }
  Swal.fire({
    title: 'Confirmación',
    text: "¿Estas seguro de asignar este puesto como administrador del modulo?",
    icon: 'warning',
    showCancelButton: true,
    confirmButtonColor: '#3085d6',
    cancelButtonColor: '#d33',
    confirmButtonText: 'Si, agregar',
    cancelButtonText: 'Cancelar',
    reverseButtons: true
  }).then((result) => {
    if (result.isConfirmed) {
      this._puestos.AgregaPuestoRole(this.puestoRole).subscribe(result => {
        Swal.fire(
          'Agregado',
          result.Mensaje,
          'success'
        )
        this.dialogRef.close(true);
      }, err => {
        if (err.error !== '' || err.error !== undefined) {
          const erores = <RespuestaPeticion>err.error;
          Swal.fire({
            position: 'center',
            title: erores.Mensaje,
            icon: 'error',
            showConfirmButton: false,
          });
        }
      });
    }
  })
}
eliminarAdministrador() {
  Swal.fire({
    title: 'Eliminar administrador',
    text: "Estas seguro de eliminar el administrador del modulo",
    icon: 'warning',
    showCancelButton: true,
    confirmButtonColor: '#3085d6',
    cancelButtonColor: '#d33',
    confirmButtonText: 'Si eliminar',
    cancelButtonText: 'Cancelar',
    reverseButtons: true
  }).then((result) => {
    if (result.isConfirmed) {
      Swal.fire(
        'Deleted!',
        'Your file has been deleted.',
        'success'
      )
    }
  })
}
}
