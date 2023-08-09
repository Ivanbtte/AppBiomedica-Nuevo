import {Injectable} from '@angular/core';
import {RequestFiltersModel} from '../../core/models/ServiciosSubrogados/trasladoPacientes/requestFilters.model';
import {Contrato} from '../../core/models/contrato.interface';
import {SolicitudTrasladoService} from '../../core/servicios/SolicitudTraslados/solicitud-traslado.service';

@Injectable({
  providedIn: 'root'
})
export class SidenavFiltersService {
  unitMedicalId = '';
  filterNumberContract = '';
  dateStart = '';
  dateEnd = '';
  withCompanion = true;
  unaccompanied = true;
  filters: RequestFiltersModel[] = [];
  origins: string[] = [];
  destinations: string[] = [];
  contracts: Contrato[];

  constructor(private transferRequestService: SolicitudTrasladoService,
  ) {
  }
}
