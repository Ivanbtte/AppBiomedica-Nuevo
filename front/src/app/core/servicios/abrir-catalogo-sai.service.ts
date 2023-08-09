import {Injectable} from '@angular/core';
import {BehaviorSubject, Observable, Subject} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AbrirCatalogoSaiService {
  estado = false;
  btn_estado = false;
  cuadro = '';
  private enviarEstado = new Subject<boolean>();
  private regresar = new Subject<boolean>();
  private ruta = new Subject<string>();
  ruta_regresar = this.ruta.asObservable();
  private enviarCuadro = new Subject<string>();
  private activarBoton = new BehaviorSubject<boolean>(false);

  public getEstadoBtn() {
    return this.activarBoton.asObservable();
  }

  constructor() {
  }

  public setEstadoBtn(valor: boolean) {
    return this.activarBoton.next(valor);
  }

  public setRuta(valor: string) {
    return this.ruta.next(valor);
  }

  public getRuta(): Observable<string> {
    return this.ruta.asObservable();
  }

  enviarEstadoComoponente(estado: boolean) {
    this.estado = estado;
    this.enviarEstado.next(this.estado);
  }

  enviarCuadroComponente(cuadro: string) {
    this.cuadro = cuadro;
    this.enviarCuadro.next((this.cuadro));
  }

  obtenerCuadro() {
    return this.cuadro;
  }

  obtenerEstado() {
    return this.estado;
  }
}
