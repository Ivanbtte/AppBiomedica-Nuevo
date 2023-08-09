import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Params} from '@angular/router';
import {MatTableDataSource} from '@angular/material/table';
import {ModulesService} from '../../core/modules.service';
import {AddPermitModuleComponent} from '../dialogs/add-permit-module/add-permit-module.component';
import {MatDialog} from '@angular/material/dialog';
import {UsuarioService} from '../../../authentication/servicios/usuario.service';

@Component({
  selector: 'app-view-permit-by-user',
  templateUrl: './view-permit-by-user.component.html',
  styleUrls: ['./view-permit-by-user.component.scss']
})
export class ViewPermitByUserComponent implements OnInit {
  unitMedic = '';
  userId = '';
  displayedColumns: string[] = ['module', 'edit',];
  dataSource: MatTableDataSource<any>;
  user: any;

  constructor(
    private route: ActivatedRoute,
    private modulesService: ModulesService,
    public dialog: MatDialog,
    private usuarioService: UsuarioService
  ) {
  }

  ngOnInit(): void {
    this.route.params.subscribe((params: Params) => {
      this.unitMedic = params.unitMedic;
      this.userId = params.id;
    });
    this.consultModules();
    this.consultUser();
  }

  consultModules() {
    this.modulesService.ConsultModules('61366fb1d7be0a40b654873d').subscribe(result => {
      this.dataSource = new MatTableDataSource<any>(result.Data);
      console.log(result.Data);
    });
  }

  consultUser() {
    this.usuarioService.ConsultUSerByPosition(this.userId).subscribe(result => {
      this.user = result.Data;
      console.log(result.Data);
    });
  }

  addPermitModule(name, moduleId: string) {
    const dialogRef = this.dialog.open(AddPermitModuleComponent, {
      height: 'auto',
      width: '80%',
      disableClose: true,
      data: {
        user: this.user,
        nameModule: name,
        moduleId: moduleId
      }
    });
    dialogRef.afterClosed().subscribe(result => {
      console.log('The dialog was closed');
    });
  }
}
