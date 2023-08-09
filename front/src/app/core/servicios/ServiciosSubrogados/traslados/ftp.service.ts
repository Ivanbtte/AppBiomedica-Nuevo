import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {RespuestaPeticion} from '../../../models/estructuras_generales';
import {environment} from '../../../../../environments/environment';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class FtpService {

  constructor(private _http: HttpClient) {
  }

  /* ******  Servicio para consultar los folios del ftp  ******   */
  consultarFolioFtp(folio: string) {
    const queryParams = new HttpParams().set('folio', folio);
    return this._http.get<RespuestaPeticion>(`${environment.host}/ftp`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }


  CargarFolios(data: FormData) {
    return this._http.post<RespuestaPeticion>(`${environment.host}/ftp`, data).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
