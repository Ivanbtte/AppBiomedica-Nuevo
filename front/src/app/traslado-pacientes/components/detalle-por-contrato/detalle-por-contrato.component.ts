import { Component, OnInit, ViewChild, ViewEncapsulation } from '@angular/core';
import { Router } from "@angular/router";
import { ContratoService } from "../../../core/servicios/ContratosServicios/contrato.service";
import {
  ChartConfiguration,
  ChartEvent,
  ChartType,
} from "chart.js";
import { BaseChartDirective } from "ng2-charts";
import DataLabelsPlugin from "chartjs-plugin-datalabels";
export interface DevengoPorMes {
  mes: number;
  devengo: number;
}
export interface DetalleTablero {
  devengo: DevengoPorMes[];
  acompaniante: any[];
}
export interface DataDevengo {
  mes: string;
  devengo: number;
  promedio: number;
}
export interface Meses {
  numero: number;
  nombre: string;
}
export interface HorizontalData {
  unidad: string;
  total: number;
  acompaniante: number;
}
export interface PorcentajeAcompaniante {
  unidad: string;
  promedio: string;
}
@Component({
  selector: "app-detalle-por-contrato",
  templateUrl: "./detalle-por-contrato.component.html",
  styleUrls: ["./detalle-por-contrato.component.scss"],
  encapsulation: ViewEncapsulation.None
})
export class DetallePorContratoComponent implements OnInit {
  colorTab:'primary'
  datosHorizontal: HorizontalData[];
  datosPorcentaje: PorcentajeAcompaniante[];
  gastoPromedio: number;
  mesesNombre: Meses[] = [
    { numero: 1, nombre: "ENERO" },
    { numero: 2, nombre: "FEBRERO" },
    { numero: 3, nombre: "MARZO" },
    { numero: 4, nombre: "ABRIL" },
    { numero: 5, nombre: "MAYO" },
    { numero: 6, nombre: "JUNIO" },
    { numero: 7, nombre: "JULIO" },
    { numero: 8, nombre: "AGOSTO" },
    { numero: 9, nombre: "SEPTIEMBRE" },
    { numero: 10, nombre: "OCTUBRE" },
    { numero: 11, nombre: "NOVIEMBRE" },
    { numero: 12, nombre: "DICIEMBRE" },
  ];
  item = JSON.parse(<string>localStorage.getItem("devengoContrato"));
  @ViewChild(BaseChartDirective) chart: BaseChartDirective | undefined;
  datos: DetalleTablero;
  meses: DataDevengo[] = [];
  // Fucniones para la grafica

  public barChartOptions: ChartConfiguration["options"] = {
    responsive: true,
    color: "black",
    // We use these empty structures as placeholders for dynamic theming.
    scales: {
      x: {
        ticks: {
          color: "black",
        },
      },
      y: {
        ticks: {
          color: "black",
          // Include a dollar sign in the ticks
          callback: function (value, index, values) {
            if (value >= 1000) {
              return (
                "$" + value.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",")
              );
            } else {
              return "$" + value;
            }
          },
        },
      },
    },
    plugins: {
      tooltip: {
        callbacks: {
          label: function (context) {
            let label = context.dataset.label || "";
            return label;
          },
        },
      },
      legend: {
        display: true,
      },
      title: {
        display: true,
        text: "Devengo por mes",
        color: "black",
      },
      datalabels: {
        anchor: "end",
        align: "end",
        color: "#C02942",
        //produccion
        formatter: function (value, context) {
          if (context.dataset.type === "bar") {
            if (value >= 1000) {
              return (
                "$" + value.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",")
              );
            } else {
              return "$" + value;
            }
          } else {
            return "";
          }
        },
        font: {
          weight: "bold",
          size: 13,
        },
      },
    },
  };
  public barChartOptions2: ChartConfiguration["options"] = {
    responsive: true,
    indexAxis: "y",
    color: "black",
    // We use these empty structures as placeholders for dynamic theming.
    scales: {
      x: {
        ticks: {
          color: "black",
        },
      },
      y: {
        ticks: {
          color: "black",
        },
      },
    },
    plugins: {
      legend: {
        display: true,
        position: "right",
      },
      title: {
        display: true,
        text: "Solicitudes por unidad",
        color: "black",
      },
      datalabels: {
        anchor: "end",
        align: "end",
        color: "#C02942",
        font: {
          weight: "bold",
          size: 13,
        },
        formatter: function (value) {
          const item = " " + value
          return item;
        },
      },
    },
  };
  public barChartOptions3: ChartConfiguration["options"] = {
    responsive: true,
    indexAxis: "y",
    color: "black",
    // We use these empty structures as placeholders for dynamic theming.
    scales: {
      x: {
        ticks: {
          color: "black",
          callback: function (value, index, values) {
            const valueTmp = parseFloat(value.toString());
            const tmp = new Intl.NumberFormat("en-US", {
              style: "percent",
              minimumFractionDigits: 0,
              maximumFractionDigits: 0,
            }).format(valueTmp);
            return tmp;
          },
        },
      },
      y: {
        ticks: {
          color: "black",
        },
      },
    },
    plugins: {
      tooltip: {
        callbacks: {
          label: function (context) {
            let label = context.dataset.label || "";
            if (label) {
              label += ": ";
            }
            if (context.parsed.x !== null) {
              label += new Intl.NumberFormat("en-US", {
                style: "percent",
                minimumFractionDigits: 2,
                maximumFractionDigits: 2,
              }).format(context.parsed.x);
            }
            return label;
          },
        },
      },
      legend: {
        display: true,
        position: "top",
      },
      title: {
        display: true,
        text: "Porcentaje de solicitudes con acompañantes por unidad",
        color: "black",
      },
      datalabels: {
        anchor: "center",
        align: "center",
        color: "white",
        font: {
          weight: "bold",
          size: 20,
        },
        formatter: function (value, context) {
          const tmp = new Intl.NumberFormat("en-US", {
            style: "percent",
            minimumFractionDigits: 2,
            maximumFractionDigits: 2,
          }).format(value);
          return tmp;
        },
      },
    },
  };
  public barChartType: ChartType = "bar";
  public barChartPlugins = [DataLabelsPlugin];
  data2: any;
  data3: any;
  data4: any;
  constructor(private router: Router, private contrato: ContratoService) {}

  ngOnInit(): void {
    this.cargarDatos();
  }
  regresar() {
    localStorage.removeItem("devengoContrato");
    this.router.navigateByUrl("/traslado/tablero");
  }
  cargarDatos() {
    this.contrato
      .ConsultarDetalleContrato(this.item.contrato.NumeroContrato)
      .subscribe((result) => {
        this.datos = result.Data;
        this.crearArreglo();
        this.asignarValores();
      });
  }
  crearArreglo() {
    const extraerFechaInicio = this.item.contrato.InicioContrato.split("/");
    const extraerFechaFin = this.item.contrato.FinContrato.split("/");
    const inicio = new Date(
      extraerFechaInicio[2],
      extraerFechaInicio[1] - 1,
      extraerFechaInicio[0]
    ).getTime();
    const fin = new Date(
      extraerFechaFin[2],
      extraerFechaFin[1] - 1,
      extraerFechaFin[0]
    ).getTime();
    const diferenciaMeses = fin - inicio;
    const diff = diferenciaMeses / (1000 * 60 * 60 * 24);
    this.gastoPromedio = ((this.item.contrato.MontoTotal * 1.16) / diff) * 30;
    for (let item = 0; item < this.mesesNombre.length; item++) {
      const temp = this.datos.devengo.filter((x) => {
        if (this.mesesNombre[item].numero === x.mes) {
          return { mes: this.mesesNombre[item].nombre, devengo: x.devengo };
        }
      });
      if (temp.length > 0) {
        this.meses.push({
          mes: this.mesesNombre[item].nombre,
          devengo: temp[0].devengo,
          promedio: this.gastoPromedio,
        });
      } else {
        this.meses.push({
          mes: this.mesesNombre[item].nombre,
          devengo: 0,
          promedio: this.gastoPromedio,
        });
      }
    }
    const temp = this.datos.acompaniante.map((item) => {
      return {
        unidad: item.unidad.DenominacionUni,
        total: item.total,
        acompaniante: item.acompaniante,
      };
    });
    const tempPromedio = this.datos.acompaniante.map((item) => {
      return {
        unidad: item.unidad.DenominacionUni,
        promedio: item.promedio,
      };
    });
    this.datosHorizontal = temp;
    this.datosPorcentaje = tempPromedio;
  }
  asignarValores() {
    const prom = new Intl.NumberFormat("en", {
      style: "currency",
      currency: "MXN",
    }).format(this.gastoPromedio);
    this.datosPorcentaje.sort((a, b) => {
      if (b.promedio > a.promedio) return 1;
      if (b.promedio < a.promedio) return -1;
      return 0;
    });
    console.log(this.meses)
    this.data2 = {
      labels: this.meses.map((item) => item.mes),
      datasets: [
        {
          type: "bar",
          data: this.meses.map((item) => item.devengo),
          label: this.item.contrato.NumeroContrato,
          backgroundColor: "#032E46",
          borderColor: "#032E46",
          hoverBackgroundColor: "#032E46",
        },
        {
          type: "line",
          data: this.meses.map((item) => item.promedio),
          label: `Promedio mensual ${prom}`,
          order: 1,
          borderDash: [10, 5],
          backgroundColor: "#DDCAA1",
          borderColor: "#DDCAA1",
          pointBackgroundColor: "#BF9B51",
        },
      ],
    };
    this.data3 = {
      labels: this.datosHorizontal.map((item) => item.unidad),
      datasets: [
        {
          data: this.datosHorizontal.map((item) => item.total),
          label: "Solicitudes totales",
          backgroundColor: "#0C555C",
          borderColor: "#0C555C",
          hoverBackgroundColor: "#0C555C",
        },
        {
          data: this.datosHorizontal.map((item) => item.acompaniante),
          label: "Solicitudes con acompañantes",
          backgroundColor: "#D6523C",
          borderColor: "#D6523C",
          hoverBackgroundColor: "#D6523C",
        },
      ],
    };
    this.data4 = {
      labels: this.datosPorcentaje.map((item) => item.unidad),
      datasets: [
        {
          data: this.datosPorcentaje.map((item) => item.promedio),
          label: "Porcentaje",
          backgroundColor: "#0F3057",
          borderColor: "#0F3057",
          hoverBackgroundColor: "#0F3057",
        },
      ],
    };
  }
  // events
  public chartClicked(event:any): void {
   console.log('event:',event.event.chart);
  }
  descargarImagen(grafica: string, nombre: string) {
    const imageLink = document.createElement("a");
    const canvas = document.getElementById(grafica) as HTMLCanvasElement;
    imageLink.download = `${nombre}.png`;
    imageLink.href = canvas.toDataURL("image/png", 1);
    imageLink.click();
  }
  public chartHovered({
    event,
    active,
  }: {
    event?: ChartEvent;
    active?: {}[];
  }): void {}
}
