import {Component, OnInit, ViewChild, ViewEncapsulation} from '@angular/core';
import {MatAccordion} from '@angular/material/expansion';

@Component({
  selector: 'app-listas',
  templateUrl: './listas.component.html',
  styleUrls: ['./listas.component.scss'],
  encapsulation: ViewEncapsulation.None
})
export class ListasComponent implements OnInit {
  @ViewChild(MatAccordion) accordion: MatAccordion;
  medicos = true;
  abasto = true;
  finanzas = true;
  catalogs = true;

  constructor() {
  }

  ngOnInit() {
  }

  toggleMedicos() {
    this.medicos = !this.medicos;
  }

  toggleAbastecimiento() {
    this.abasto = !this.abasto;
  }

  toggleFinanzas() {
    this.finanzas = !this.finanzas;
  }

  toggleCatalogs() {
    this.catalogs = !this.catalogs;
  }
}
