import {Pipe, PipeTransform} from '@angular/core';

@Pipe({
  name: 'returnConditionFirm'
})
export class ReturnConditionFirmPipe implements PipeTransform {

  transform(state: boolean): string {
    if (!state) {
      return 'Activar';
    } else {
      return 'Inactivar';
    }
  }

}
