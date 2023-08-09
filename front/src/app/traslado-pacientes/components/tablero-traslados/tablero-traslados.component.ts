import { Component, OnInit } from "@angular/core";
import { ContratoService } from "../../../core/servicios/ContratosServicios/contrato.service";
import { Router } from "@angular/router";

@Component({
  selector: "app-tablero-traslados",
  templateUrl: "./tablero-traslados.component.html",
  styleUrls: ["./tablero-traslados.component.css"],
})
export class TableroTrasladosComponent implements OnInit {
  contratos: any[] = [];
  constructor(private contrato: ContratoService, private router: Router) {}

  ngOnInit(): void {
    this.consultarContratos();
  }

  consultarContratos() {
    this.contrato
      .ConsultarContratoSubTipo("5f587535d7be0a74bf7db093")
      .subscribe((result) => {
        this.contratos = result.Data;
        console.log(this.contratos);
      });
  }
  masDetalles(item: any) {
    // guardar en el localstorage los datos del contrato
    localStorage.setItem('devengoContrato', JSON.stringify(item));
    this.router.navigateByUrl('/traslado/detalle/contrato');
    // this.masDetallesBoton = !this.masDetallesBoton;
    // const ids = this.contratos.map((it) => it.contrato.NumeroContrato);
    // const numContratos = ids.join(",")
    // this.contrato.ConsultarDetalleContrato(numContratos).subscribe(result =>{
    //   console.log(result.Data)
    // })
  }
}
