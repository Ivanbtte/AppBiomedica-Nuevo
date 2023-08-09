import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-solicitud-subrogado',
  templateUrl: './solicitud-subrogado.component.html',
  styleUrls: ['./solicitud-subrogado.component.css']
})
export class SolicitudSubrogadoComponent implements OnInit {
unidad: any;
  constructor() { }

  ngOnInit(): void {
    this.unidad = JSON.parse(<string>localStorage.getItem("unit"));
    console.log(this.unidad)
  } 
  dataSource: any[];

}
