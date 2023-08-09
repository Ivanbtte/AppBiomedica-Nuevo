import {Component, OnInit, ViewChild} from '@angular/core';
import {MatDialog} from '@angular/material/dialog';
import {MatPaginator} from '@angular/material/paginator';
import {MatTableDataSource} from '@angular/material/table';
import {DialogNewFirmComponent} from '../dialogs/dialog-new-firm/dialog-new-firm.component';
import {DialogEditFirmComponent} from '../dialogs/dialog-edit-firm/dialog-edit-firm.component';
import {FirmService} from '../../../core/servicios/ServiciosSubrogados/traslados/firm.service';
import {FirmaSolicitud} from '../../../core/models/ServiciosSubrogados/trasladoPacientes/firma_solicitud.interface';
import Swal from 'sweetalert2';
import {RespuestaPeticion} from '../../../core/models/estructuras_generales';
import {ReturnConditionFirmPipe} from '../../../shared/pipes/return-condition-firm.pipe';
import {ReturnTypeFirmPipe} from '../../../shared/pipes/return-type-firm.pipe';

@Component({
  selector: 'app-firmas',
  templateUrl: './firmas.component.html',
  styleUrls: ['./firmas.component.css']
})
export class FirmasComponent implements OnInit {
  displayedColumns = ['name', 'position', 'matric', 'type', 'state', 'action'];
  @ViewChild(MatPaginator, {static: true}) paginator!: MatPaginator;
  pageSize = 20;
  size = [20, 50, 100];
  totalRecords = 0;
  actualPage = 1;
  dataSource: any;
  unidadMedica: any;

  constructor(
    public dialog: MatDialog,
    private firmService: FirmService
  ) {
    this.dataSource = new MatTableDataSource();
  }

  ngOnInit() {
    this.unidadMedica = JSON.parse(<string>localStorage.getItem("unit"));
    console.log(this.unidadMedica)
    this.paginator._intl.itemsPerPageLabel = 'Resultados por pagina';
    this.getAllFirm();
  }

  getAllFirm() {
    this.firmService.GetFirmsByUnit(this.unidadMedica.unidadMedId).subscribe(result => {
      console.log(result);
      this.dataSource.data = <FirmaSolicitud> result.Data;
      this.dataSource.paginator = this.paginator;
      this.totalRecords = this.dataSource.data.length;
    });
  }

  openDialogFirm() {
    const dialogRef = this.dialog.open(DialogNewFirmComponent, {
      height: 'auto',
      width: '55%',
      disableClose: true,
    });
    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.getAllFirm();
      }
    });

  }

  openDialogFirmEdit(row: FirmaSolicitud) {
    const dialogRef = this.dialog.open(DialogEditFirmComponent, {
      height: 'auto',
      width: '55%',
      disableClose: true,
      data: {firm: row},
    });
    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.getAllFirm();
      }
    });
  }

  enableFirm(row: FirmaSolicitud) {
    console.log(row);
    const returnConditionFirmPipe = new ReturnConditionFirmPipe();
    const stateFirm = returnConditionFirmPipe.transform(row.estado);
    const returnTypeFirmPipe = new ReturnTypeFirmPipe();
    const type = returnTypeFirmPipe.transform(row.tipo);
    const unitName: any = JSON.parse(<string>localStorage.getItem('unit'));
    console.log(unitName);
    Swal.fire({
      title: '¿Deseas ' + stateFirm + '?',
      html:
        'A <b>' + row.nombre + '</b>,<br>' +
        'Como: <b>' + type + '</b> <br>' +
        'Para el:  <b>' + unitName.unidadMed.DenominacionUni + '</b> ',
      showCancelButton: true,
      icon: 'warning',
      confirmButtonColor: '#00b162',
      cancelButtonColor: '#d33',
      confirmButtonText: stateFirm,
      cancelButtonText: 'Cancelar'
    }).then((result) => {
      if (result.value) {
        if (!row.estado) {
          row.estado = false;
        }
        const userId = <string>localStorage.getItem('idUser');
        this.firmService.EnableFirm(row.id, row.estado.toString(), row.unidadMedId.toString(), userId).subscribe(resp => {
          Swal.fire({
            position: 'center',
            title: resp.Mensaje,
            icon: 'success',
            showConfirmButton: true,
          });
          console.log(resp.Data);
          this.getAllFirm();
        }, err => {
          if (err.error !== '' || err.error !== undefined) {
            const erores = <RespuestaPeticion> err.error;
            Swal.fire({
              position: 'center',
              title: 'Error',
              text: erores.Mensaje,
              icon: 'error',
              showConfirmButton: true,
            });
          }
        });
      }
    });
  }
}
