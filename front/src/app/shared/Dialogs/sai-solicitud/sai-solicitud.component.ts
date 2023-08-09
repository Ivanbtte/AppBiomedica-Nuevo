import {Component, Inject, OnInit} from '@angular/core';
import {MAT_DIALOG_DATA} from '@angular/material/dialog';

@Component({
  selector: 'app-sai-solicitud',
  templateUrl: './sai-solicitud.component.html',
  styleUrls: ['./sai-solicitud.component.css']
})
export class SaiSolicitudComponent implements OnInit {
  estado = false;

  constructor(
    @Inject(MAT_DIALOG_DATA) public data: any
  ) {
    this.estado = data;
    console.log(this.estado);
  }

  ngOnInit() {
  }

}
