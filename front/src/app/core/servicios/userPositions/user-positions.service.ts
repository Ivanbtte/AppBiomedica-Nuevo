import { Injectable } from "@angular/core";
import { HttpClient, HttpParams } from "@angular/common/http";
import { RespuestaPeticion } from "../../models/estructuras_generales";
import { environment } from "../../../../environments/environment";
import { map } from "rxjs/operators";

@Injectable({
  providedIn: "root",
})
export class UserPositionsService {
  constructor(private httpClient: HttpClient) {}

  consultUsersValid(status, valid, unitMedId, clvDele, departmentId: string) {
    let queryParams = new HttpParams();
    if (valid !== "") {
      queryParams = queryParams.set("valid", valid);
    }
    if (status !== "") {
      queryParams = queryParams.set("status", status);
    }
    if (unitMedId !== "") {
      queryParams = queryParams.set("unitMedId", unitMedId);
    }
    if (clvDele !== "") {
      queryParams = queryParams.set("clvDele", clvDele);
    }
    if (departmentId !== "") {
      queryParams = queryParams.set("departmentId", departmentId);
    }
    return this.httpClient
      .get<RespuestaPeticion>(`${environment.host}/usuarios/valid`, {
        params: queryParams,
      })
      .pipe(
        map((res: RespuestaPeticion) => {
          return res;
        })
      );
  }

  addNewUserPosition(obj: any, userType: string, update: string) {
    let queryParams = new HttpParams().set("userType", userType);
    if (update !== "") {
      queryParams = queryParams.set("update", update);
    }
    return this.httpClient
      .post<RespuestaPeticion>(`${environment.host}/usuarioPuesto`, obj, {
        params: queryParams,
      })
      .pipe(
        map((res: RespuestaPeticion) => {
          return res;
        })
      );
  }
}
