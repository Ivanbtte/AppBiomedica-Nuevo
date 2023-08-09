import {Component, Inject, OnInit} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {FormControl} from '@angular/forms';
import {VisualizadorPDFComponent} from '../visualizador-pdf/visualizador-pdf.component';
import {MatDialog} from '@angular/material/dialog';
import {ArchivosContratosService} from '../../../../core/servicios/ContratosServicios/archivos-contratos.service';
import {MatTabChangeEvent} from '@angular/material/tabs';

export interface DialogData {
  contrato: any;
}

@Component({
  selector: 'app-archivos-contrato',
  templateUrl: './archivos-contrato.component.html',
  styleUrls: ['./archivos-contrato.component.css']
})
export class ArchivosContratoComponent implements OnInit {
  public Contrato: any;
  public Carpetas: any[]=[];
  public Archivos: any[]=[];
  public ArchivosSubcarpeta: any[]=[];
  tabs = [];
  selected = new FormControl(0);
  tab_seleccioando = '';

  constructor(
    public dialogRef: MatDialogRef<ArchivosContratoComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
    private _archivos: ArchivosContratosService,
    public dialog: MatDialog,
  ) {
  }

  ngOnInit() {
    this.Contrato = this.data.contrato;
    console.log(this.Contrato);
    this.consultarCarpetas('');
  }

  removeTab(index: number) {
    this.tabs.splice(index, 1);
  }

  consultarCarpetas(path: string) {
    this._archivos.ConsultarCarpetas(path).subscribe(res => {
      this.tabs = res.Data;
    });
  }

  consultarArchivos(path: string) {
    this._archivos.ConsultarArchivos(this.Contrato.NumeroContrato, path).subscribe(res => {
      this.Archivos = res.Data;
      console.log(this.Archivos);
    });
  }

  consultarArchivossubcarpeta(path: string) {
    this._archivos.ConsultarArchivos(this.Contrato.NumeroContrato, path).subscribe(res => {
      this.ArchivosSubcarpeta = res.Data;
      console.log(this.ArchivosSubcarpeta);
    });
  }

  tabSelecionado(event: MatTabChangeEvent) {
    this.tab_seleccioando = event.tab.textLabel;
    console.log('Seleccionaste este', event.tab.textLabel);
    this.selecionarUnaCarpeta(event.tab.textLabel);
    this.consultarArchivos(event.tab.textLabel);
  }

  selecionarUnaCarpeta(nombre: string) {
    this._archivos.ConsultarCarpetas(nombre).subscribe(res => {
      this.Carpetas = res.Data;
      console.log(this.Carpetas);
    });
  }

  onNoClick(): void {
    this.dialogRef.close();
  }

  /* ******  Funcion para abrir la visualizacion de un pdf ******   */
  AbrirPDF(item: any) {
    const dialogRef = this.dialog.open(VisualizadorPDFComponent, {
      height: 'auto',
      width: '80%',
      disableClose: true,
      data: {
        contrato: this.Contrato,
        path: item.Ruta + '/' + item.Nombre
      }
    });

    dialogRef.afterClosed().subscribe(result => {
    });
  }

  visualizar(item: any): boolean {
    return (item.Nombre.includes('.pdf'));
  }

  archivo_carpeta(archivo: any, carpeta: any): boolean {
    return (archivo.Ruta.includes(carpeta.Nombre));
  }

  descargar_archivo(item: any) {
    this._archivos.DescargarArchivo(item.Ruta + '/' + item.Nombre).subscribe(res => {
      console.log(res);
      if (File) {
        console.log(File);
      }
    });
  }
}
