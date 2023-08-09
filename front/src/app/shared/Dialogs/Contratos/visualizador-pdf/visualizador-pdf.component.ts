import {Component, Inject, OnInit, ViewChild} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {ArchivosContratosService} from '../../../../core/servicios/ContratosServicios/archivos-contratos.service';

export interface DialogData {
  contrato: any;
  path: string;
}

@Component({
  selector: 'app-visualizador-pdf',
  templateUrl: './visualizador-pdf.component.html',
  styleUrls: ['./visualizador-pdf.component.css']
})
export class VisualizadorPDFComponent implements OnInit {
  // @ViewChild('pdfViewerOnDemand') pdfViewerOnDemand!;
  // @ViewChild('pdfViewerAutoLoad') pdfViewerAutoLoad!;
  nombre_descarga = '';

  constructor(
    public dialogRef: MatDialogRef<VisualizadorPDFComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
    private _archivos: ArchivosContratosService
  ) {
  }

  ngOnInit() {
    this.descargar_archivo(this.data.path);
    this.nombre_descarga = 'Contrato_' + this.data.contrato.NumeroContrato + '.pdf';
  }

  descargar_archivo(path: string) {
    this._archivos.DescargarArchivo(path).subscribe(res => {
      console.log(res);
      // this.pdfViewerAutoLoad.pdfSrc = res; // pdfSrc can be Blob or Uint8Array
      // this.pdfViewerAutoLoad.refresh(); // Ask pdf viewer to load/refresh pdf
      if (File) {
        console.log(File);
      }
    });
  }

  onNoClick(): void {
    this.dialogRef.close();
  }
}
