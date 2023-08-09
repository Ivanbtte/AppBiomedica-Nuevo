import {Component, Inject, OnInit} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {ProveedoresInterface} from '../../../../core/models/proveedores.interface';
import {FormControl} from '@angular/forms';
import {ContratoService} from '../../../../core/servicios/ContratosServicios/contrato.service';
import {ProveedoresService} from '../../../../core/servicios/ProveedoresServicios/proveedores.service';
import {DelegacionesService} from '../../../../core/servicios/delegaciones.service';

export interface Filtros {
  proveedor: any;
  nombre: string;
  delegacion: any;
  unidad_med: any;
}

@Component({
  selector: 'app-filtros-contratos',
  templateUrl: './filtros-contratos.component.html',
  styleUrls: ['./filtros-contratos.component.css']
})
export class FiltrosContratosComponent implements OnInit {
  provedores: ProveedoresInterface[] = [];
  nombre_proveedor = new FormControl(this.data.nombre);
  resultado_proveedores = new FormControl('');
  delegaciones = new FormControl();
  sin_resultados = false;
  delegacionesArray: any[]=[];

  constructor(
    private _contrato: ContratoService,
    private  _proveedores: ProveedoresService,
    public dialogRef: MatDialogRef<FiltrosContratosComponent>,
    @Inject(MAT_DIALOG_DATA) public data: Filtros,
    private _delegaciones: DelegacionesService,
  ) {
    console.log(this.data);
    if (this.data.nombre !== null && this.data.nombre !== '' && this.data.nombre !== undefined) {
      this.nombre_proveedor.setValue(this.data.nombre);
      this.buscarProveedor();
      this.resultado_proveedores.setValue(this.data.proveedor);
    }
    if (this.data.delegacion !== null || this.data.delegacion !== '') {
      this.delegaciones.setValue(this.data.delegacion);
    }
  }

  ngOnInit() {
    this.nombre_proveedor.valueChanges.subscribe(res => {
      if (res === '') {
        this.provedores = [];
        this.sin_resultados = false;
        this.resultado_proveedores.setValue('');
      }
    });
    this._delegaciones.ConsultarDelegaciones().subscribe(resultado => {
      this.delegacionesArray = resultado.Data;
    });
  }

  limpiarFiltros() {
    this.nombre_proveedor.setValue('');
    this.provedores = [];
    this.sin_resultados = false;
    this.resultado_proveedores.setValue('');
    this.delegaciones.setValue('');
  }

  buscarProveedor() {
    console.log(this.nombre_proveedor.value);
    const filtro = this.nombre_proveedor.value !== null ? this.nombre_proveedor.value : '';
    this._proveedores.ConsultarProveedoresFiltro(filtro).subscribe(resultado => {
      this.provedores = resultado.Data;
      this.sin_resultados = false;
      console.log(this.provedores);
      if (this.provedores.length === 0) {
        this.sin_resultados = true;
      }
    });
  }

  filtrar() {
    console.log(this.resultado_proveedores.value, this.delegaciones.value);
    this.data.proveedor = this.resultado_proveedores.value;
    this.data.nombre = this.nombre_proveedor.value;
    this.data.delegacion = this.delegaciones.value;
    this.dialogRef.close(this.data);
  }

  onNoClick(): void {
    this.data.delegacion = '';
    this.data.proveedor = '';
    this.data.nombre = '';
    this.dialogRef.close(this.data);
  }

  compare(c1: { NombreEmpresa: string }, c2: { NombreEmpresa: string }) {
    return c1 && c2 && c1.NombreEmpresa === c2.NombreEmpresa;
  }

  compareDele(c1: { NombreDele: string }, c2: { NombreDele: string }) {
    return c1 && c2 && c1.NombreDele === c2.NombreDele;
  }
}
