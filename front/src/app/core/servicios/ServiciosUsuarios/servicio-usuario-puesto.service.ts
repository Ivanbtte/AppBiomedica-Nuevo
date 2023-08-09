import { HttpClient, HttpParams } from "@angular/common/http";
import { not } from "@angular/compiler/src/output/output_ast";
import { Injectable } from "@angular/core";
import { map } from "rxjs/operators";
import { environment } from "src/environments/environment";
import {
  BitacoraUsuarioPuesto,
  UnsuscribeUser,
} from "../../models/bitacora_usuario_puesto.interface";
import { RespuestaPeticion } from "../../models/estructuras_generales";

@Injectable({
  providedIn: "root",
})
export class ServicioUsuarioPuestoService {
  constructor(private http: HttpClient) {}

  AddBitacoraUsuarioPuesto(obj: Partial<BitacoraUsuarioPuesto>, status: string, roleId: string) {
    const queryParams = new HttpParams().set("status", status).set('roleId',roleId);
    return this.http
      .post<RespuestaPeticion>(`${environment.host}/usuarioPuesto/bitacora`,obj,{ params: queryParams })
      .pipe(map((res: RespuestaPeticion) => {
          return res;
        })
      );
  }

  GetTipoModificaciones(notId: string) {
    const queryParams = new HttpParams().set("notId", notId);
    return this.http
      .get<RespuestaPeticion>(
        `${environment.host}/usuarios/tipo/modificacion`,
        { params: queryParams }
      )
      .pipe(
        map((res: RespuestaPeticion) => {
          return res;
        })
      );
  }

  UnsuscribeUsuarioPuesto(
    unsuscribe: Partial<BitacoraUsuarioPuesto>,
    idUser: string
  ) {
    const queryParams = new HttpParams().set("idUser", idUser);
    return this.http
      .put<RespuestaPeticion>(
        `${environment.host}/usuarios/unsuscribe`,
        unsuscribe,
        { params: queryParams }
      )
      .pipe(
        map((res: RespuestaPeticion) => {
          return res;
        })
      );
  }
}
