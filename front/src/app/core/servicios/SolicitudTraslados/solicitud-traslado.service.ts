import { Injectable } from '@angular/core';
import { RespuestaPeticion } from '../../models/estructuras_generales';
import { environment } from '../../../../environments/environment';
import { map } from 'rxjs/operators';
import { HttpClient, HttpParams } from '@angular/common/http';
import { SolicitudtrasladoBoletos } from '../../models/ServiciosSubrogados/trasladoPacientes/solicitud_traslado.interface';

@Injectable({
  providedIn: 'root'
})
export class SolicitudTrasladoService {

  constructor(
    private http: HttpClient) {
  }

  RealizarSolicitud(obj: SolicitudtrasladoBoletos, fecha: string, uiPrei: string, firmsId: string, folioSeparado: string) {
    let queryParams = new HttpParams();
    if (fecha !== '') {
      queryParams = queryParams.set('fecha', fecha);
    }
    if (uiPrei !== '') {
      queryParams = queryParams.set('uiPrei', uiPrei);
    }
    if (firmsId !== '') {
      queryParams = queryParams.set('firmsId', firmsId);
    }
    if (folioSeparado !== '') {
      queryParams = queryParams.set('folioSeparado', folioSeparado)
    }
    return this.http.post<RespuestaPeticion>(`${environment.host}/traslados/solicitud`, obj, { params: queryParams }).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  ConsultSolicitudes(unit, numberPage, pageSize, filter, state, companion, origin, destination, numberContract, dateStart, dateEnd: string) {
    let queryParams = new HttpParams().set('unitMedical', unit).set('numberPage', numberPage).set('pageSize', pageSize);
    if (filter !== '') {
      queryParams = queryParams.set('filter', filter);
    }
    if (state !== '') {
      queryParams = queryParams.set('state', state);
    }
    if (companion !== '') {
      queryParams = queryParams.set('companion', companion);
    }
    if (origin !== '') {
      queryParams = queryParams.set('origin', origin);
    }
    if (destination !== '') {
      queryParams = queryParams.set('destination', destination);
    }
    if (dateStart !== '') {
      queryParams = queryParams.set('dateStart', dateStart);
    }
    if (dateEnd !== '') {
      queryParams = queryParams.set('dateEnd', dateEnd);
    }
    if (numberPage !== '') {
      queryParams = queryParams.set('numberContract', numberContract);
    }

    return this.http.get<RespuestaPeticion>(`${environment.host}/traslados/solicitados`, { params: queryParams }).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  ConsultarFolioSolicitud(folio: string) {
    let queryParams = new HttpParams();
    if (folio !== '') {
      queryParams = queryParams.set('folio', folio);
    }
    return this.http.get<RespuestaPeticion>(`${environment.host}/traslados/folio`, { params: queryParams }).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  CancelarFolioSolicitud(obj: any, folio: string,boletosSeparados:string) {
    let queryParams = new HttpParams();
    if (folio !== '') {
      queryParams = queryParams.set('folio', folio);
    }
    if (boletosSeparados !== '') {
      queryParams = queryParams.set('boletoSeparados', boletosSeparados);
    }
    return this.http.put<RespuestaPeticion>(`${environment.host}/traslados/cancelar`, obj, { params: queryParams }).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  PrintTemplate(idUnit: string, idVoBo: string, idAuthorize: string) {
    let queryParams = new HttpParams();
    if (idUnit !== '') {
      queryParams = queryParams.set('id', idUnit);
    }
    if (idVoBo !== '') {
      queryParams = queryParams.set('idVoBo', idVoBo);
    }
    if (idAuthorize !== '') {
      queryParams = queryParams.set('idAuthorize', idAuthorize);
    }
    return this.http.get(`${environment.host}/traslados/pdf/plantilla`, { params: queryParams, responseType: 'blob' },);
  }

  printDataOnly(folio, reimpresion: string, folioSeparados: string) {
    let queryParams = new HttpParams();
    if (folio !== '') {
      queryParams = queryParams.set('folio', folio);
    }
    if (reimpresion !== '') {
      queryParams = queryParams.set('reimpresion', reimpresion);
    }
    if (folioSeparados !== '') {
      queryParams = queryParams.set('folioSeparados', folioSeparados);
    }
    return this.http.get(`${environment.host}/traslados/pdf/datos`, { params: queryParams, responseType: 'blob' },);
  }

  printRequestComplete(folio: string, estado: string, reimpresion: string, folioSeparados: string) {
    let queryParams = new HttpParams();
    if (folio !== '') {
      queryParams = queryParams.set('folio', folio);
    }
    if (estado !== '') {
      queryParams = queryParams.set('estado', estado);
    }
    if (reimpresion !== '') {
      queryParams = queryParams.set('reimpresion', reimpresion);
    }
    if (folioSeparados !== '') {
      queryParams = queryParams.set('folioSeparados', folioSeparados);
    }
    return this.http.get(`${environment.host}/traslados/pdf/completo`, { params: queryParams, responseType: 'blob' },);
  }

  GetOriginFilters(unitId, state, destination, companion, contractsNumber, startDate, endDate: string) {
    let queryParams = new HttpParams().set('unitId', unitId).set('state', state);
    if (destination !== '') {
      queryParams = queryParams.set('destination', destination);
    }
    if (companion !== '') {
      queryParams = queryParams.set('companion', companion);
    }
    if (contractsNumber !== '') {
      queryParams = queryParams.set('numberContract', contractsNumber);
    }
    if (startDate !== '' && endDate !== '') {
      queryParams = queryParams.set('startDate', startDate);
      queryParams = queryParams.set('endDate', endDate);
    }
    return this.http.get<RespuestaPeticion>(`${environment.host}/traslados/filter/origin`, { params: queryParams }).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  GetDestinationFilters(unitId, state, origin, companion, contractsNumber, startDate, endDate: string) {
    let queryParams = new HttpParams().set('unitId', unitId).set('state', state);
    if (origin !== '') {
      queryParams = queryParams.set('origin', origin);
    }
    if (companion !== '') {
      queryParams = queryParams.set('companion', companion);
    }
    if (contractsNumber !== '') {
      queryParams = queryParams.set('numberContract', contractsNumber);
    }
    if (startDate !== '' && endDate !== '') {
      queryParams = queryParams.set('startDate', startDate);
      queryParams = queryParams.set('endDate', endDate);
    }
    return this.http.get<RespuestaPeticion>(`${environment.host}/traslados/filter/destination`, { params: queryParams }).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  GetCompanionFilters(unitId, state, origin, destination, contractsNumber, startDate, endDate: string) {
    let queryParams = new HttpParams().set('unitId', unitId).set('state', state);
    if (origin !== '') {
      queryParams = queryParams.set('origin', origin);
    }
    if (destination !== '') {
      queryParams = queryParams.set('destination', destination);
    }
    if (contractsNumber !== '') {
      queryParams = queryParams.set('numberContract', contractsNumber);
    }
    if (startDate !== '' && endDate !== '') {
      queryParams = queryParams.set('startDate', startDate);
      queryParams = queryParams.set('endDate', endDate);
    }
    return this.http.get<RespuestaPeticion>(`${environment.host}/traslados/filter/companion`, { params: queryParams }).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
