import {Component, OnInit, ViewChild} from '@angular/core';
import {ProformaService} from '../servicios/proforma.service';
import {FormControl} from '@angular/forms';
import {ServicioProforma} from '../../core/models/interfaceCatalogos/servicioPro.interface';
import {Observable} from 'rxjs';
import {MatDialog} from '@angular/material/dialog';
import {MatSort} from '@angular/material/sort';
import {ServicioProformaService} from '../servicios/servicio-proforma.service';
import {map, startWith} from 'rxjs/operators';
import {DetalleProformaComponent} from '../detalle-proforma/detalle-proforma.component';
import {MatPaginator} from '@angular/material/paginator';
import {MatTableDataSource} from '@angular/material/table';

@Component({
  selector: 'app-proforma',
  templateUrl: './proforma.component.html',
  styleUrls: ['./proforma.component.css']
})
export class ProformaComponent implements OnInit {
  datosProforma: any;
  myControl = new FormControl();
  control_nombre = new FormControl();
  displayedColumns: string[] = ['IdArticulo', 'ClaveCuadroBasico', 'Descripcion', 'Vista'];
  cargando = false;
  dataSource: any;
  options: ServicioProforma[] = [];
  filteredOptions!: Observable<ServicioProforma[]>;
  lastFilter = '';
  @ViewChild(MatPaginator, {static: true}) paginator!: MatPaginator;
  @ViewChild(MatSort, {static: true}) sort!: MatSort;
  tamanioPagina = 20;
  tamanios = [20, 50, 100];
  totalRegistros = 0;
  paginaActual = 1;
  serviceProforma = '';

  constructor(
    private _proforma: ProformaService,
    private _serviciosPro: ServicioProformaService,
    public dialog: MatDialog
  ) {
  }

  /* *****  funcion para abrir el dialog de la clace a detallar, recibe el id de la clave  *****   */
  openDialog(id: string): void {
    const dialogRef = this.dialog.open(DetalleProformaComponent, {
      height: 'auto',
      width: 'auto',
      data: {id: id}
    });

    dialogRef.afterClosed().subscribe(result => {
    });
  }

  /* *****  metodo que que ejecuta al iniciar el componente  *****   */
  ngOnInit() {
    this.llenarTabla();
    this.llenarSelect();
    this.filteredOptions = this.myControl.valueChanges.pipe(
      startWith(''),
      map(value => typeof value === 'string' ? value : value.Descripcion),
      map(descripcion => descripcion ? this._filter(descripcion) : this.options.slice())
    );
    this.paginator.page.subscribe(page => {
      if (page) {
        this.paginaActual = page['pageIndex'] + 1;
        this.tamanioPagina = page['pageSize'];
        this.llenarTabla();
      }
    });
    this.myControl.valueChanges.subscribe(result => {
      console.log(result);
      if (result === '') {
        this.buscarPorGrupo();
      }
    });
  }

  /* *****  funcion que llena la tabla   *****   */
  llenarTabla() {
    const nombre = this.control_nombre.value !== null ? this.control_nombre.value : '';
    this._proforma.getProforma(this.serviceProforma, nombre, this.paginaActual.toString(), this.tamanioPagina.toString()).subscribe(res => {
      this.dataSource = new MatTableDataSource();
      this.dataSource.data = res.Data['registros'];
      this.totalRegistros = res.Data['total'];
      this.cargando = true;
      this.datosProforma = res.Data;
      console.log(this.datosProforma);
    }, err => {
      console.log(err);
    });
  }

  /* ***** funcion que llena el select de los filtros  *****   */
  llenarSelect() {
    this._serviciosPro.getServiciosProf().subscribe(res => {
      this.options = res.Data;
    }, err => {
      console.log(err);
    });
  }

  buscarPorNombre() {
    this.paginator.firstPage();
    this.llenarTabla();
  }

  buscarPorGrupo() {
    this.serviceProforma = this.myControl.value !== '' ? this.myControl.value.Id : '0';
    this.control_nombre.reset();
    this.paginator.firstPage();
    this.llenarTabla();
  }

  private _filter(filter: string): ServicioProforma[] {
    const filterValue = filter.toLowerCase();
    return this.options.filter(option => option.Descripcion.toLowerCase().includes(filterValue));
  }

  displayFn(servicio: ServicioProforma) {
    return servicio ? servicio.Descripcion : '';
  }

}
