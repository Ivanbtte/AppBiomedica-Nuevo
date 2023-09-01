import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { SolicitudServLaboratorio } from '../../models/ServiciosSubrogados/trasladoPacientes/servicios_laboratorio.interface';

@Injectable({
  providedIn: 'root'
})
export class ServiciosSubrugadosLaboratorioService {

  constructor(private http: HttpClient) { }
  RealizarSolicitud(obj: SolicitudServLaboratorio){
    
  }
}
