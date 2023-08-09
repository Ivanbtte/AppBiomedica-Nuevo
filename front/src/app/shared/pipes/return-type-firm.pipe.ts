import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'returnTypeFirm'
})
export class ReturnTypeFirmPipe implements PipeTransform {

  transform(type: number): string {
    if (type === 1) {
      return 'Vo. Bo.';
    } else {
      return 'Autoriza';
    }
  }

}
