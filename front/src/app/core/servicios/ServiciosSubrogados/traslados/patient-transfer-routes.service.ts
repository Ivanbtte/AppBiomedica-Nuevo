import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {RespuestaPeticion} from '../../../models/estructuras_generales';
import {environment} from '../../../../../environments/environment';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class PatientTransferRoutesService {

  constructor(private httpClient: HttpClient) {
  }

  /* ******  Servicio para consultar los destinos por unidad medica  ******   */
  ConsultarDestinos(numero_contrato: string, unidad_id: string) {
    const queryParams = new HttpParams().set('numero_contrato', numero_contrato).set('unidad_id', unidad_id);
    return this.httpClient.get<RespuestaPeticion>(`${environment.host}/traslados/destinos`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  /* ******  Servicio para consultar los origenes por unidad medica  ******   */
  ConsultarOrigenes(numero_contrato: string, unidad_id: string, destino: string) {
    const queryParams = new HttpParams().set('numero_contrato', numero_contrato).set('unidad_id', unidad_id).set('destino', destino);
    return this.httpClient.get<RespuestaPeticion>(`${environment.host}/traslados/origenes`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  ConsultBudget(contractNumber, unitId: string) {
    const queryParams = new HttpParams().set('contractNumber', contractNumber).set('unitId', unitId);
    return this.httpClient.get<RespuestaPeticion>(`${environment.host}/traslados/budget`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  ConsultEarn(state, contractNumber, unitId: string) {
    const queryParams = new HttpParams().set('state', state).set('contractNumber', contractNumber).set('unitId', unitId);
    return this.httpClient.get<RespuestaPeticion>(`${environment.host}/traslados/earn`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  ConsultTicketGoingAvailable(destination, typeTicket, contractNumber, state, unitId: string) {
    let queryParams = new HttpParams();
    if (destination !== '') {
      queryParams = queryParams.set('destination', destination);
    }
    if (typeTicket !== '') {
      queryParams = queryParams.set('typeTicket', typeTicket);
    }
    if (contractNumber !== '') {
      queryParams = queryParams.set('contractNumber', contractNumber);
    }
    if (state !== '') {
      queryParams = queryParams.set('state', state);
    }
    if (unitId !== '') {
      queryParams = queryParams.set('unitId', unitId);
    }
    return this.httpClient.get<RespuestaPeticion>(`${environment.host}/traslados/ticketGoing`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
