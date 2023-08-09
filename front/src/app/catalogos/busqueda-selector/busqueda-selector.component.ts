import {Component, Inject, ViewChild} from '@angular/core';
import {MatDialogRef, MAT_DIALOG_DATA} from '@angular/material/dialog';
import {TemaGrupo} from '../../core/models/interfaceCatalogos/filtros.interface';
import {FormControl} from '@angular/forms';
import {Observable} from 'rxjs';
import {map, startWith} from 'rxjs/operators';
import {FiltrosService} from '../servicios/filtros.service';
import {MatSelectionList} from '@angular/material/list';

@Component({
  selector: 'app-busqueda-selector',
  templateUrl: './busqueda-selector.component.html',
  styleUrls: ['./busqueda-selector.component.css']
})
export class BusquedaSelectorComponent {
  filtroTema: TemaGrupo[];
  myControl = new FormControl();
  filteredOptions: Observable<TemaGrupo[]>;
  lastFilter = '';
  Filtro: TemaGrupo[] = [];
  radioTema = new FormControl();
  filtroPorTema = false;
  filtroGrupo = false;
  @ViewChild('shoes') shoes!: MatSelectionList;

  constructor(
    public dialogRef: MatDialogRef<BusquedaSelectorComponent>,
    @Inject(MAT_DIALOG_DATA) public data: any,
    private filtrosService: FiltrosService,
  ) {
    this.filtroTema = data.datos;
    console.log(this.filtroTema);
    this.filteredOptions = this.myControl.valueChanges.pipe(
      startWith<string | TemaGrupo[]>(''),
      map(value => typeof value === 'string' ? value : this.lastFilter),
      map(filter => this._filter(filter))
    );
    this.filtrosService.getGrupoTemas('true', '').subscribe(res => {
      this.Filtro = res.Data;
    }, err => {
      console.log(err);
    });
  }

  onNoClick(): void {
    if (this.radioTema.value === '1') {
      const arr = this.shoes.selectedOptions.selected;
      const seleccionados = arr.map(it => it.value);
      console.log(seleccionados);
      this.dialogRef.close(seleccionados);
    }
    if (this.radioTema.value === '2') {
      console.log(this.myControl.value);
      if (this.myControl.value !== null) {
        const seleccionados2 = [{
          Id: this.myControl.value.Grupo,
          Tema: this.myControl.value.Tema.Tema
        }];
        console.log('grupos', seleccionados2);
        this.dialogRef.close(seleccionados2);
      } else {
        this.dialogRef.close([]);

      }

    }
  }

  displayFn(tema: TemaGrupo) {
    return tema ? tema.Tema.Tema : '';
  }

  borrar() {
    console.log('abri');
    this.myControl.setValue('');
  }

  private _filter(filter: string): TemaGrupo[] {
    if (filter) {
      this.lastFilter = filter.toLowerCase();
      return this.Filtro.filter(option => option.Tema.Tema.toLowerCase().indexOf(this.lastFilter) === 0);
    } else {
      return this.Filtro;
    }
  }

  valorRadioButton() {
    if (this.radioTema.value === '1') {
      this.filtroPorTema = true;
      this.filtroGrupo = false;
      this.myControl.reset();
    }
    if (this.radioTema.value === '2') {
      this.filtroPorTema = false;
      this.filtroGrupo = true;
      if (this.shoes) {
        this.shoes.deselectAll();
      }
    }
  }
}
