import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'returnStateFirm'
})
export class ReturnStateFirmPipe implements PipeTransform {

  transform(state: boolean): any {
    if (state) {
      return {text: 'Activo', color: '#09b107'};
    } else {
      return {text: 'Inactivo', color: '#d60017'};
    }
  }

}
