import {Component, Inject, OnInit} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {ModulesService} from '../../../core/modules.service';
import {MatTableDataSource} from '@angular/material/table';
import {ActionsPermit} from '../../../../core/models/ServiciosSubrogados/trasladoPacientes/module.interface';
import Swal from 'sweetalert2';

export interface DialogData {
  user: any;
  nameModule: string;
  moduleId: string;
}

@Component({
  selector: 'app-add-permit-module',
  templateUrl: './add-permit-module.component.html',
  styleUrls: ['./add-permit-module.component.scss']
})
export class AddPermitModuleComponent implements OnInit {
  displayedColumns: string[] = ['name', 'description', 'edit',];
  dataSource: MatTableDataSource<ActionsPermit>;

  constructor(
    public dialogRef: MatDialogRef<AddPermitModuleComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
    private modulesService: ModulesService
  ) {
  }

  ngOnInit(): void {
    this.consultModules();
  }

  consultModules() {
    this.modulesService.ConsultActionsByModule(this.data.moduleId, this.data.user.Id).subscribe(result => {
      this.dataSource = new MatTableDataSource<ActionsPermit>(result.Data);
      console.log(result.Data);
    });
  }

  onNoClick(): void {
    this.dialogRef.close();
  }

  changePermit(item: ActionsPermit) {
    console.log(item);
    const message = item.permit !== true ? 'Activar' : 'Desactivar';
    const text = item.permit !== true ? 'De activar el permiso de la funcion' : 'De desactivar el permiso de la funcion';
    Swal.fire({
      title: '¿Estas seguro?',
      text: text + ': ' + item.action.name,
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      confirmButtonText: 'Si, ' + message,
      reverseButtons: true,
      cancelButtonText: 'Cancelar'
    }).then((result) => {
      if (result.isConfirmed) {
        Swal.fire(
          'Deleted!',
          'Your file has been deleted.',
          'success'
        );
      }
    });
  }
}
