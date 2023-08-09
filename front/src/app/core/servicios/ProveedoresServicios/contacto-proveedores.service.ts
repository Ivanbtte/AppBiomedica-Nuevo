import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {RespuestaPeticion} from '../../models/estructuras_generales';
import {environment} from '../../../../environments/environment';
import {map} from 'rxjs/operators';
import {ContactoProveedor} from '../../models/proveedores.interface';

@Injectable({
  providedIn: 'root'
})
export class ContactoProveedoresService {

  constructor(
    private http: HttpClient
  ) {
  }

  ConsultarContactoProveedor(id_contacto: string, id_proveedor: string, nombre: string) {
    let queryParams = new HttpParams().set('id_proveedor', id_proveedor);
    if (id_contacto !== '') {
      queryParams = queryParams.set('id_contacto', id_contacto);
    }
    if (nombre !== '') {
      queryParams = queryParams.set('nombre_proveedor', nombre);
    }
    return this.http.get<RespuestaPeticion>(`${environment.host}/proveedores/contacto`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  AgregarContactoProveedor(obj: ContactoProveedor) {
    return this.http.post<RespuestaPeticion>(`${environment.host}/proveedores/contacto`, obj).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  ActualizarContactoProveedor(obj: ContactoProveedor, id_contacto: string) {
    const queryParams = new HttpParams().set('id_contacto', id_contacto);
    return this.http.put<RespuestaPeticion>(`${environment.host}/proveedores/contacto`, obj, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }

  EliminarContacto(id_contacto: string) {
    const queryParams = new HttpParams().set('id_contacto', id_contacto);
    return this.http.delete<RespuestaPeticion>(`${environment.host}/proveedores/contacto`, {params: queryParams}).pipe(
      map((res: RespuestaPeticion) => {
        return res;
      })
    );
  }
}
