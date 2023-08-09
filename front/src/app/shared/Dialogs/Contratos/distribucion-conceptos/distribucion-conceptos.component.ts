import {Component, Inject, OnInit} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialog, MatDialogRef} from '@angular/material/dialog';
import { MatTableDataSource} from '@angular/material/table';
import Swal from 'sweetalert2';
import {ConceptoContrato} from '../../../../core/models/concepto_contrato.interface';
import {FormBuilder} from '@angular/forms';
import {UsuarioService} from '../../../../authentication/servicios/usuario.service';
import {AgregarEditarDistrComponent} from '../agregar-editar-distr/agregar-editar-distr.component';
import {RespuestaPeticion} from '../../../../core/models/estructuras_generales';
import {DistribucionConceptosService} from '../../../../core/servicios/ContratosServicios/distribucion-conceptos.service';
import {ConceptosContratoService} from '../../../../core/servicios/ContratosServicios/conceptos-contrato.service';
import {DelegacionesService} from '../../../../core/servicios/delegaciones.service';

export interface DialogData {
  concepto: string;
  cantidad: number;
  archivo: boolean;
}

@Component({
  selector: 'app-distribucion-conceptos',
  templateUrl: './distribucion-conceptos.component.html',
  styleUrls: ['./distribucion-conceptos.component.css']
})
export class DistribucionConceptosComponent implements OnInit {
  concepto: any;
  claveCuadro='';
  distribucion: any[]=[];
  sumatoriaCantidad!: number;
  displayedColumns: string[] = ['Delegacion', 'unidadM', 'CantidadDis', 'PrecioDis', 'EditarPro', 'EliminarPro'];
  dataSource: any;
  dist_dupl = false;
  delegacionesArray: any[]=[];
  por_distribuir!: number;
  estado_distribucion = false;

  constructor(
    public dialogRef: MatDialogRef<DistribucionConceptosComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
    private  _distribucion: DistribucionConceptosService,
    private _conceptos: ConceptosContratoService,
    private _delegaciones: DelegacionesService,
    private fb: FormBuilder,
    private _unidades: UsuarioService,
    public dialog: MatDialog,
  ) {
    this.concepto = data.concepto;
    this.claveCuadro = this.concepto.Prei.Grupo + '.' + this.concepto.Prei.Generico + '.' + this.concepto.Prei.Especifico + '.' +
      this.concepto.Prei.Diferenciador + '.' + this.concepto.Prei.Variable;
    console.log(this.concepto);
    console.log(this.data.cantidad);
    if (data.cantidad !== 0) {
      this.dist_dupl = true;
    }
  }

  applyFilter(filterValue: string) {
    this.dataSource.filter = filterValue.trim().toLowerCase();
  }

  ngOnInit() {
    this.consultarDistribucion();
    this._delegaciones.ConsultarDelegaciones().subscribe(resultado => {
      this.delegacionesArray = resultado.Data;
      console.log(this.delegacionesArray);
    });
  }

  sumatoriaCantidades(array: any[]) {
    const sumatoria = array.reduce((acc, item) => acc + item.Cantidad, 0);
    this.sumatoriaCantidad = sumatoria;
  }

  consultarDistribucion() {
    this._distribucion.ConsultarDistribucion(this.concepto.Id).subscribe(resultado => {
      console.log(resultado.Data);
      this.distribucion = resultado.Data;
      this.dataSource = new MatTableDataSource();
      this.dataSource.data = this.distribucion;
      this.sumatoriaCantidades(resultado.Data);
      this.por_distribuir = this.concepto.CantidadConcepto - this.sumatoriaCantidad;
      this.dataSource.filterPredicate = (data: any, filter: string) => {
        return data.UnidadMed.DenominacionUni.trim().toLocaleLowerCase().indexOf(filter) !== -1 ||
          data.UnidadMed.Delegacion.NombreDele.trim().toLocaleLowerCase().indexOf(filter) !== -1;
      };
    });
  }

  regresaDescripcion(valor: string): string {
    const descripcion = valor.split('.', 1);
    return descripcion.toString();
  }

  onNoClick(): void {
    this.dialogRef.close(this.estado_distribucion);
  }

  async cambiarCantidad(concepto: ConceptoContrato) {
    const {value: email} = await Swal.fire({
      title: 'Cantidad Contratada',
      text: concepto.CantidadConcepto.toString(),
      input: 'number',
      inputPlaceholder: 'Cantidad a editar',
      showCancelButton: true,
    });

    if (email) {
      concepto.CantidadConcepto = parseInt(email, 10);
      this._conceptos.ActualizarConceptoContrato(concepto.Id, concepto, concepto.ContratoNumeroContrato).subscribe(res => {
        console.log(res);
        this.consultarDistribucion();
        this.dialogRef.close(true);
      });
    }
  }

  agregardistribucion() {
    const dialogRef = this.dialog.open(AgregarEditarDistrComponent, {
      width: '50%',
      height: 'auto',
      disableClose: true,
      data: {
        distribucion: '',
        concepto: this.regresaDescripcion(this.concepto.Prei.Descripcion),
        editar: false,
        cantidadMax: this.por_distribuir,
        id_concepto: this.concepto
      }
    });
    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.consultarDistribucion();
      }
    });

  }

  EditarCantidad(distribucion: any) {
    console.log('Editando');
    const dialogRef = this.dialog.open(AgregarEditarDistrComponent, {
      width: '50%',
      height: 'auto',
      disableClose: true,
      data: {
        distribucion: distribucion,
        concepto: this.regresaDescripcion(this.concepto.Prei.Descripcion),
        editar: true,
        cantidadMax: this.por_distribuir,
        id_concepto: this.concepto
      }
    });
    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.consultarDistribucion();
      }
    });
  }

  EliminarCantidad(distribucion: any) {
    console.log(distribucion);
    Swal.fire({
      title: '¿Estás seguro que deseas eliminar esta distribución?',
      icon: 'warning',
      showCancelButton: true,
      cancelButtonColor: '#a2a2a2',
      confirmButtonColor: '#e62324',
      cancelButtonText: 'Cancelar',
      confirmButtonText: 'Si, Eliminar',
      reverseButtons: true
    }).then((result) => {
      if (result.value) {
        this._distribucion.EliminarDistribucion(distribucion.Id, this.concepto.ContratoNumeroContrato,
          distribucion.UnidadMed.DelegacionesClvDele).subscribe(res => {
          Swal.fire(
            'Correcto',
            res.Mensaje,
            'success'
          );
          this.consultarDistribucion();
        }, err => {
          if (err.error !== '' || err.error !== undefined) {
            const erores = <RespuestaPeticion>err.error;
            Swal.fire({
              position: 'center',
              title: erores.Mensaje,
              text: erores.Data,
              icon: 'error',
              showConfirmButton: true,
            });
          }
        });
      }
    });
  }

  // ********************************************************************************************************
  /* *****   *****   */
  /*  const index = this.chipsVer.indexOf(chips);
    if (index >= 0) {
    this.chipsVer.splice(index, 1);
    this.buscarPorTema();
  }*/
}
