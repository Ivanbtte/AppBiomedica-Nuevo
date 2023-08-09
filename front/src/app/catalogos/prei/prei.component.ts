import {Component, OnInit, ViewChild} from '@angular/core';
import {PreiService} from '../servicios/prei.service';
import {MatDialog} from '@angular/material/dialog';
import {MatTableDataSource} from '@angular/material/table';
import {FormControl} from '@angular/forms';
import {Observable} from 'rxjs';
import {CuadroBasico} from '../../core/models/interfaceCatalogos/prei.interface';
import {map, startWith} from 'rxjs/operators';
import {DetalleDialogComponent} from '../detalle-dialog/detalle-dialog.component';
import {MatPaginator} from '@angular/material/paginator';

@Component({
    selector: 'app-prei',
    templateUrl: './prei.component.html',
    styleUrls: ['./prei.component.scss']
})
export class PreiComponent implements OnInit {
    myControl = new FormControl();
    control_nombre = new FormControl();
    control_grupo = new FormControl();
    control_generico = new FormControl();
    control_especifico = new FormControl();
    displayedColumns: string[] = ['Clave', 'IdArticulo', 'Descripcion', 'Tipo', 'Precio', 'Vista'];
    cargando = false;
    cuadro_id = '';
    grupo_id = '';
    generico_id = '';
    especifico_id = '';
    dataSource: any;
    options: CuadroBasico [] = [];
    filteredOptions!: Observable<CuadroBasico[]>;
    lastFilter = '';
    @ViewChild(MatPaginator, {static: true}) paginator!: MatPaginator;
    tamanioPagina = 50;
    tamanios = [20, 50, 100];
    totalRegistros = 0;
    paginaActual = 1;
    arreglo_grupos: any[] = [];
    arreglo_genericos: any[] = [];
    arreglo_especificos: any[] = [];
    filteredGrupos!: Observable<string[]>;
    filteredGenerico!: Observable<string[]>;
    filteredEspecifico!: Observable<string[]>;
    lastFilterGrupos = '';
    lastFilterGenerico = '';
    lastFilterEspecifico = '';

    constructor(
        private _prei: PreiService,
        public dialog: MatDialog
    ) {
    }

    ngOnInit() {
        this.cargarSelect();
        this.llenarArreglo();
        this.consultarGrupo();

        this.control_grupo.valueChanges.subscribe(res => {
            if (res === '') {
                this.borrarGrupo();
            }
        });
        this.control_generico.valueChanges.subscribe(resultado => {
            if (resultado === '') {
                this.borrarGenerico();
            }
        });
        this.control_especifico.valueChanges.subscribe(res => {
            if (res === '') {
                this.borrarEspecifico();
            }
        });
        this.myControl.valueChanges.subscribe(result => {
            console.log(result);
            if (result === '') {
                this.buscarPorTema();
            }
        });
        this.paginator._intl.itemsPerPageLabel = 'Resultados por pagina';
        this.paginator.page.subscribe(page => {
            if (page) {
                this.paginaActual = page['pageIndex'] + 1;
                this.tamanioPagina = page['pageSize'];
                this.llenarArreglo();
            }
        });
        this.filteredOptions = this.myControl.valueChanges.pipe(
            startWith(''),
            map(value => typeof value === 'string' ? value : value.Grupo),
            map(grupo => grupo ? this._filter2(grupo) : this.options.slice())
        );
    }

    // ********************************************************************************************************
    /* *****  funcion para abrir el dialog del detalle de la clave  *****   */
    openDialog(id: string): void {
        const dialogRef = this.dialog.open(DetalleDialogComponent, {
            height: 'auto',
            width: 'auto',
            data: {id: id}
        });

        dialogRef.afterClosed().subscribe(result => {
        });
    }

    // ********************************************************************************************************
    /* *****  Funciones para cargar las opciones del autocomplete de los inputs  *****   */
    cargaAutoCompleteGenerico() {
        this.filteredGenerico = this.control_generico.valueChanges.pipe(
            startWith<string>(''),
            map(value => typeof value === 'string' ? value : this.lastFilterGenerico),
            map(filter => this._filter_generico(filter))
        );
    }

    cargaAutoCompleteEspecifico() {
        this.filteredEspecifico = this.control_especifico.valueChanges.pipe(
            startWith<string>(''),
            map(value => typeof value === 'string' ? value : this.lastFilterEspecifico),
            map(filter => this._filter_especifico(filter))
        );
    }

    cargaAutocompleteGrupo() {
        this.filteredGrupos = this.control_grupo.valueChanges.pipe(
            startWith<string>(''),
            map(value5 => typeof value5 === 'string' ? value5 : this.lastFilterGrupos),
            map(filter5 => this._filter_grupos(filter5))
        );
    }

    // ********************************************************************************************************
    /* *****  funcion para llenar el arreglo del select de filtro por temas  *****   */
    cargarSelect() {
        this._prei.getCuadroBasico().subscribe(res => {
            this.options = res.Data;
        }, err => {
            console.log(err);
        });
    }

    // ********************************************************************************************************
    /* *****  funciones para consultar los genericos o especificos de los autocomplete  *****   */
    consultarGrupo() {
        this._prei.consultarFiltrosPrei('', '', true).subscribe(resultado => {
            this.arreglo_grupos = resultado.Data.map((it: { Grupo: any; }) => it.Grupo);
            this.cargaAutocompleteGrupo();
            console.log(this.arreglo_grupos);
        });
    }

    consultarGenerico(grupo: string) {
        this._prei.consultarFiltrosPrei(grupo, '', false).subscribe(resultado => {
            this.arreglo_genericos = resultado.Data.map((it: { Generico: any; }) => it.Generico);
            console.log(this.arreglo_genericos);
            this.cargaAutoCompleteGenerico();
        });
    }

    consultaEspecifico(grupo: string, generico: string) {
        this._prei.consultarFiltrosPrei(grupo, generico, false).subscribe(resultado => {
            this.arreglo_especificos = resultado.Data.map((it: { Especifico: any; }) => it.Especifico);
            console.log(this.arreglo_especificos);
            this.cargaAutoCompleteEspecifico();
        });
    }

    // ********************************************************************************************************
    /* *****  funciones para realizar los diferentes filtros  *****   */
    buscarPorNombre() {
        this.paginator.firstPage();
        this.llenarArreglo();
    }

    buscarPorTema() {
/*        if (this.myControl.value !== '') {
            this.cuadro_id = this.myControl.value !== null ? this.myControl.value.Id : '0';
        }*/
        this.cuadro_id = this.myControl.value !== '' ? this.myControl.value.Id : '0';
        this.paginator.firstPage();
        this.control_nombre.reset();
        this.grupo_id = '';
        this.control_grupo.reset();
        this.generico_id = '';
        this.control_generico.reset();
        this.especifico_id = '';
        this.control_especifico.reset();
        this.llenarArreglo();
        this.consultarGrupo();
    }

    buscarPorGrupo() {
        this.grupo_id = this.control_grupo.value !== null ? this.control_grupo.value : 0;
        this.cuadro_id = '';
        this.myControl.reset();
        this.generico_id = '';
        this.control_generico.reset();
        this.especifico_id = '';
        this.control_nombre.reset();
        this.paginator.firstPage();
        this.llenarArreglo();
        this.consultarGenerico(this.grupo_id);
    }

    buscarPorGenerico() {
        if (this.control_generico.value !== null) {
            console.log('buscando por generico');
            this.generico_id = this.control_generico.value;
            this.control_nombre.reset();
            this.paginator.firstPage();
            this.llenarArreglo();
            this.consultaEspecifico(this.grupo_id, this.generico_id);
        }
    }

    buscarPorEspecifico() {
        if (this.control_especifico.value !== null) {
            this.especifico_id = this.control_especifico.value;
            this.control_nombre.reset();
            this.paginator.firstPage();
            this.llenarArreglo();
        }
    }

    // ********************************************************************************************************
    /* *****  funcion que carga los resultados en la tabla  *****   */
    llenarArreglo() {
        const nombre = this.control_nombre.value !== null ? this.control_nombre.value : '';
        this._prei.getPrei(this.cuadro_id, nombre, this.paginaActual.toString(), this.tamanioPagina.toString()
            , this.grupo_id, this.generico_id, this.especifico_id).subscribe(res => {
            this.dataSource = new MatTableDataSource();
            this.dataSource.data = res.Data['registros'];
            console.log(res.Data['registros']);
            this.totalRegistros = res.Data['total'];
            this.cargando = true;
        }, err => {
            console.log(err);
        });
    }

    // ********************************************************************************************************
    /* *****  funciones para hacer el filtro de los autocomplete  *****   */
    private _filter2(value: string): CuadroBasico[] {
        const filterValue = value.toLowerCase();
        return this.options.filter(option => option.Grupo.toLowerCase().includes(filterValue));
    }

    private _filter_grupos(value: string): string[] {
        if (value) {
            this.lastFilterGrupos = value.toLowerCase();
            return this.arreglo_grupos.filter(option => option.toLowerCase().indexOf(this.lastFilterGrupos) === 0);
        } else {
            return this.arreglo_grupos;
        }
    }

    private _filter_generico(value: string): string[] {
        if (value) {
            this.lastFilterGenerico = value.toLowerCase();
            return this.arreglo_genericos.filter(option => option.toLowerCase().indexOf(this.lastFilterGenerico) === 0);
        } else {
            return this.arreglo_genericos;
        }
    }

    private _filter_especifico(value: string): string[] {
        if (value) {
            this.lastFilterEspecifico = value.toLowerCase();
            return this.arreglo_especificos.filter(option => option.toLowerCase().indexOf(this.lastFilterEspecifico) === 0);
        } else {
            return this.arreglo_especificos;
        }
    }

    // ********************************************************************************************************
    /* *****  funcion para mostrar el valor seleccionado en el autocomplete de temas  *****   */
    displayFn(cuadro: CuadroBasico): string {
        return cuadro && cuadro.Grupo ? cuadro.Grupo : '';
    }

    // ********************************************************************************************************
    /* *****  funciones para eliminar y resetear los inputs de grupo, generioo, especifico  *****   */
    borrarGrupo() {
        this.grupo_id = '';
        this.generico_id = '';
        this.especifico_id = '';
        this.control_grupo.reset();
        this.control_generico.reset();
        this.control_especifico.reset();
        this.control_nombre.reset();
        this.llenarArreglo();
        this.consultarGrupo();
    }

    borrarGenerico() {
        this.generico_id = '';
        this.especifico_id = '';
        this.control_generico.reset();
        this.control_especifico.reset();
        this.control_nombre.reset();
        this.llenarArreglo();
    }

    borrarEspecifico() {
        this.especifico_id = '';
        this.control_especifico.reset();
        this.control_nombre.reset();
        this.llenarArreglo();
    }

    regresaDescripcion(valor: string): string {
        const descripcion = valor.split('.', 1);
        return descripcion.toString();
    }
}
