import { Component, Input, OnInit, ViewChild } from '@angular/core';
import { Contrato } from '../../../core/models/contrato.interface';
import { ConceptosContratoService } from '../../../core/servicios/ContratosServicios/conceptos-contrato.service';
import { MatPaginator } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';

@Component({
  selector: 'app-estudios-laboratorio',
  templateUrl: './estudios-laboratorio.component.html',
  styleUrls: ['./estudios-laboratorio.component.css']
})
export class EstudiosLaboratorioComponent implements OnInit {
  @Input() data!: Contrato[];
  @Input() tipo!: string;
  @Input() subTipo!: string;
  @ViewChild(MatPaginator,{static:true}) paginator!: MatPaginator;
  MontoContratado = 0;
  dataSource: any;
  displayedColumns: string[] = ['unMed', 'especialidad', 'descripcion', 'unidadMedida', 'precioCon','canMax', 'impMax',
    'EditarPro', 'EliminarPro'];
  tamanioPagina = 15;
  tamanios = [15, 30, 45];
  totalRegistros = 0;
  paginaActual = 1;
  constructor(private _conceptos: ConceptosContratoService) { }

  ngOnInit(): void {
    console.log('2232323', this.data);
    this.consultarConceptos();
    this.paginator._intl.itemsPerPageLabel = 'Resultados por pagina';
    this.paginator.page.subscribe(page => {
      if (page) {
        this.paginaActual = page['pageIndex'] + 1;
        this.tamanioPagina = page['pageSize'];
        this.consultarConceptos();
      }
    });
  }

  consultarConceptos() {
    const id_contrato = this.data[0].NumeroContrato;
    const tipo = this.tipo !== null ? this.tipo : '';
    const sub_tipo = this.subTipo !== null ? this.subTipo : '';
    console.log('sdfsdf-*-***', tipo, sub_tipo);
    this._conceptos.ConsultarConceptos(id_contrato, tipo, sub_tipo, this.paginaActual.toString(), this.tamanioPagina.toString()).subscribe(resultado => {
      console.log(resultado.Data)
      this.dataSource = new MatTableDataSource();
      this.dataSource.data = resultado.Data['registros'].map((it:any) => it.estudiosLaboratorio)[0];
      this.MontoContratado = resultado.Data['total'];
      this.totalRegistros = resultado.Data['numReg'];
      console.log(this.dataSource);
    });
  }
}
