import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {RespuestaPeticion} from '../../../models/estructuras_generales';
import {environment} from '../../../../../environments/environment';
import {map} from 'rxjs/operators';
import {FirmaSolicitud} from '../../../models/ServiciosSubrogados/trasladoPacientes/firma_solicitud.interface';

@Injectable({
  providedIn: 'root'
})
export class FirmService {

  constructor(private http: HttpClient) {
  }

  /* ******  Servicio para consultar firma por unidad medica y tipo de firma  ******   */
  GetAllTypeFirms(unidad_id: string, tipo: string, state: string) {
    const queryParams = new HttpParams().set('tipo', tipo).set('unidad_id', unidad_id).set('state', state);
    return this.http.get<RespuestaPeticion>(`${environment.host}/firma`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  /* ******  Servicio para actualizar  una firma de una solicitud  ******   */
  UpdateFirm(obj: Partial<FirmaSolicitud>, id: string, unitId: string, matric: string, userId: string) {
    const queryParams = new HttpParams().set('id', id).set('unitId', unitId).set('matric', matric).set('userId', userId);
    return this.http.put<RespuestaPeticion>(`${environment.host}/firma`, obj, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
          return res;
        }
      )
    );
  }

  /* ******  Servicio para agregar un nombre y un cargo de una firma de una solicitud  ******   */
  AgregarFirma(obj: Partial<FirmaSolicitud>) {
    return this.http.post<RespuestaPeticion>(`${environment.host}/firma`, obj).pipe(
      map((res: RespuestaPeticion) => {
          return res;
        }
      )
    );
  }

  /* ******  Servicio para consultar todas las firmas por unidad medica  ******   */

  GetFirmsByUnit(unit_id: string) {
    const queryParams = new HttpParams().set('unitId', unit_id);
    return this.http.get<RespuestaPeticion>(`${environment.host}/firma/all`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  /* ******  Servicio para habilitar o deshabilitar una firma  ******   */
  EnableFirm(id: string, state: string, unitId: string, userId: string) {
    let queryParams = new HttpParams().set('id', id);
    if (state !== '') {
      queryParams = queryParams.set('state', state);
    }
    if (unitId !== '') {
      queryParams = queryParams.set('unitId', unitId);
    }
    if (userId !== '') {
      queryParams = queryParams.set('userId', userId);
    }
    return this.http.put<RespuestaPeticion>(`${environment.host}/firma/enable`, {}, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
          return res;
        }
      )
    );
  }
}
