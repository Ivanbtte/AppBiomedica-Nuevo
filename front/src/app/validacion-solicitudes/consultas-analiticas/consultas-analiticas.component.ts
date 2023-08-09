import {Component, Inject} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialog, MatDialogRef} from '@angular/material/dialog';
import {ClavesService} from '../../core/servicios/solicitud/claves.service';
import {SolicitudService} from '../../core/servicios/solicitud/solicitud.service';
import {PresupuestoService} from '../../configuraciones/serviciosConfiguraciones/presupuesto.service';
import Swal from 'sweetalert2';
import {UsuarioLogeado} from '../../core/models/UsuarioLogin.interface';
import jwt_decode from 'jwt-decode';
import {CantidadDialogComponent} from '../../shared/Dialogs/cantidad-dialog/cantidad-dialog.component';

export interface DialogData {
  idSolicitud: string;
  clave: string;
  clavesSolicitud: any;
}

@Component({
  selector: 'app-consultas-analiticas',
  templateUrl: './consultas-analiticas.component.html',
  styleUrls: ['./consultas-analiticas.component.css']
})
export class ConsultasAnaliticasComponent {
  /* *****  la variable ""clave""alamcena todas los registros de claves que son iguales a la que se va a analizar  *****   */
  public clave: any[]=[];
  analizarclaveEstado = false;
  /* *****  la variable ""totalClave"" cuenta la cantidad total que se ha solicitado este año  *****   */
  totalClave: any;
  /* *****  la variable ""TotalPorServio"" cuenta la cantidad total pedida por el mismo servicio en  oaxaca  *****   */
  public TotalPorServio: any;
  /* *****  la variable ""objClave"" es el objeto completo de la clave seleccionada  *****   */
  objClave: any;
  /* *****  la variable ""totalSolicitud"" cuenta la cantidad total en esta solicitud  *****   */
  totalSolicitud: any;
  totalPorUnidades: any;
  totalporSercvicio: any;
  totalPorSolicitud: any;
  desgloceUnidad = false;
  desgloceServicio = false;
  desgloceSolicitud = false;
  presupuestoPorSolici: any;
  presupuestoGenera: any;
  presupuestoDelegacional: any;
  public usuarioLog: UsuarioLogeado;

  constructor(
    public dialogRef: MatDialogRef<ConsultasAnaliticasComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
    private _conceptosSolicitud: ClavesService,
    private _solicitudes: SolicitudService,
    private _presupuesto: PresupuestoService,
    public dialog: MatDialog,
  ) {
    this.usuarioLog = jwt_decode(<string>localStorage.getItem('token'));
    this.consultaPresupuestoBD();
    this.objClave = this.data.clave;
    console.log(this.data);
    console.log(this.objClave);
    this.consultarClave();
    this.consultarTodasSolicitudes();
  }

  consultarClave() {
    this._conceptosSolicitud.ConsultarClaves('', this.objClave.ClaveId, '', '', '').subscribe(res => {
      this.clave = res.Data;
      console.log(res.Data);
      this.totalClave = this.clave.map(it => it.CantidadSolicitada).reduce((acc, value) => acc + value);
      console.log('la misma clave', this.totalClave);
      this.cantidadTotalPorSErvicio(this.clave, this.objClave.ServiciosProformaId);
      this.cantidadTotalEstasolicitud(this.clave, this.objClave.SolicitudId);
      this.verUnidadClave(this.clave);
    });
  }

  consultarTodasSolicitudes() {
    this._solicitudes.ConsultarTodasSolicitudes('', 'true').subscribe(res => {
      console.log(res.Data);
      this.calcularPresupues(res.Data);
    });
  }

  consultaPresupuestoBD() {
    this._presupuesto.ConsultaPresupuesto().subscribe(res => {
      this.presupuestoDelegacional = res.Data;
      console.log(this.presupuestoDelegacional);
    });
  }

  calcularPresupues(arr: any[]) {
    this.presupuestoPorSolici = arr.map(item => {
        return {
          Unidad: item.UnidadMed.DenominacionUni,
          Total: item.Total
        };
      }
    );
    console.log(this.presupuestoPorSolici);
    this.presupuestoGenera = arr.map(ite => ite.Total).reduce((acc, value) => acc + value);
  }

  cantidadTotalEstasolicitud(clave: any[], idSolicitud: any) {
    const totalSolicituds = clave.filter(item => {
      return item.SolicitudId === idSolicitud;
    });
    console.log('filter', totalSolicituds);
    this.totalPorSolicitud = totalSolicituds;
    this.totalSolicitud = totalSolicituds.map(itw => itw.CantidadSolicitada).reduce((acc, value) => acc + value);
    console.log('total en esta solicitud', this.totalSolicitud);
  }

  verUnidadClave(total: any[]) {
    const unidades = total.map(valor => {
      return {
        unidad: valor.Solicitud.UnidadMed.DenominacionUni,
        cantidad: valor.CantidadSolicitada
      };
    });
    // let matriz = {};
    //
    // unidades.forEach(function(registro) {
    //   let pais = registro['unidad'];
    //   matriz[pais] = matriz[pais] ? (matriz[pais] + registro.cantidad) : registro.cantidad;
    // });
    // matriz = Object.keys(matriz).map(function(pais) {
    //   return {nombreAutoriza: pais, cant: matriz[pais]};
    // });
    // console.log(matriz);
    // this.totalPorUnidades = matriz;
    console.log(this.totalPorUnidades);
  }

  cantidadTotalPorSErvicio(total: any[], servicio: any) {
    const prefiltro = total.filter(it => {
        console.log(it.ServiciosProformaId, '  ', servicio);
        return it.ServiciosProformaId === servicio;
      }
    );
    this.totalporSercvicio = prefiltro;
    console.log(prefiltro);
    this.TotalPorServio = prefiltro.map(it => it.CantidadSolicitada).reduce((acc, value) => acc + value);
    console.log(this.TotalPorServio);
  }

  validarCantidad() {
    Swal.fire({
      title: 'Confirmación',
      text: 'Estás seguro de aprobar esta clave',
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      cancelButtonText: 'Cancelar',
      confirmButtonText: 'Si, aprobar'
    }).then((result) => {
      if (result.value) {
        this.objClave.Total = this.objClave.Precio * this.objClave.CantidadSolicitada;
        this._conceptosSolicitud.ActualizarClave(this.objClave.Id, '', '', '', 'true', {
          SolicitudId: this.objClave.SolicitudId,
          ClaveId: this.objClave.ClaveId,
          Descripcion: this.objClave.Descripcion,
          CantidadSolicitada: this.objClave.CantidadSolicitada,
          Precio: this.objClave.Precio,
          Total: this.objClave.Total,
          ServiciosProformaId: this.objClave.ServiciosProformaId,
          EstatusId: this.objClave.EstatusId,
          Fecha: this.objClave.Fecha
        }).subscribe(res => {
          console.log(res.Data);
          Swal.fire({
            icon: 'success',
            title: 'Correcto',
            text: res.Mensaje,
          });
          this.onNoClick();
        });
      }
    });
  }

  openDialogCantidad(clave: any): void {
    const dialogRef = this.dialog.open(CantidadDialogComponent, {
        height: 'auto',
        width: '65%',
        disableClose: true,
        data: {
          clave: clave,
          editar: false,
          agregar: false,
          aprobar: true
        }
      })
    ;

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        console.log('entreee');
        this.consultarLaClave(this.objClave.Id);
        this.consultarClave();
        this.consultarTodasSolicitudes();
      }
    });
  }

  consultarLaClave(id: string) {
    this._conceptosSolicitud.ConsultarUnaClaves(id).subscribe(res => {
      console.log(res.Data);
      this.objClave = res.Data;
    });
  }

  onNoClick() {
    this.dialogRef.close();
  }
}
