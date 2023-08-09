import {PersonaService} from '../authentication/servicios/persona.service';
import {AbstractControl} from '@angular/forms';
import {map} from 'rxjs/operators';
import {RespuestaPeticion} from '../core/models/estructuras_generales';

export class MyValidators {

  static validateEnrollment(service: PersonaService) {
    return (control: AbstractControl) => {
      const value = control.value;
      if (value) {
        return service.checkEnrollment(value)
          .pipe(
            map((response: RespuestaPeticion) => {
              return response.Data['status'] ? {not_available: true} : null;
            })
          );
      }
    };
  }

}
