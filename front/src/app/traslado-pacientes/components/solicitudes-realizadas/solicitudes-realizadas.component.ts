import { Component, OnInit, ViewChild } from "@angular/core";
import { MatPaginator } from "@angular/material/paginator";
import { MatTableDataSource } from "@angular/material/table";
import { BreakpointObserver } from "@angular/cdk/layout";
import { UsuarioService } from "../../../authentication/servicios/usuario.service";
import { UsuarioLogeado } from "../../../core/models/UsuarioLogin.interface";
import { SolicitudTrasladoService } from "../../../core/servicios/SolicitudTraslados/solicitud-traslado.service";
import Swal from "sweetalert2";
import { FormBuilder, FormControl, FormGroup } from "@angular/forms";
import { PerfectScrollbarConfigInterface } from "ngx-perfect-scrollbar";
import { MatSidenav } from "@angular/material/sidenav";
import jwt_decode from "jwt-decode";
import { DateAdapter } from "@angular/material/core";
import * as _moment from "moment";
import { UserUnitMedical } from "../../../core/models/userUnitMed.model";
import { Solicitudtraslado } from "../../../core/models/ServiciosSubrogados/trasladoPacientes/solicitud_traslado.interface";
import { RequestFiltersModel } from "../../../core/models/ServiciosSubrogados/trasladoPacientes/requestFilters.model";
import { ContratoService } from "../../../core/servicios/ContratosServicios/contrato.service";
import { Contrato } from "../../../core/models/contrato.interface";

const moment = _moment;

@Component({
  selector: "app-solicitudes-realizadas",
  templateUrl: "./solicitudes-realizadas.component.html",
  styleUrls: ["./solicitudes-realizadas.component.css"],
})
export class SolicitudesRealizadasComponent implements OnInit {
  unitMedic: any;
  public config: PerfectScrollbarConfigInterface = {};
  @ViewChild("sidenav") sidenav: MatSidenav;
  userToken: UsuarioLogeado;
  displayedColumns = [
    "folio",
    "elaboracion",
    "ftp",
    "origen",
    "destino",
    "nombre_pac",
    "fecha",
    "accion",
  ];
  pageSize = 20;
  pageSizeOptions = [20, 50, 100];
  length = 0;
  pageIndex = 1;
  loadingData = false;
  dataSource: MatTableDataSource<Solicitudtraslado>;
  @ViewChild(MatPaginator, { static: true }) paginator!: MatPaginator;
  searchInputControl = new FormControl();
  range: FormGroup;
  filters: RequestFiltersModel[] = [];
  controlCompanion = new FormControl();
  controlDestination = new FormControl();
  controlOrigin = new FormControl();
  controlNumberContract = new FormControl();
  origins: string[] = [];
  destinations: string[] = [];
  withCompanion = true;
  unaccompanied = true;
  unitMedicalId = "";
  contracts: Contrato[];
  filterNumberContract = "";
  minDate: Date;
  maxDate: Date;
  dateStart = "";
  dateEnd = "";

  constructor(
    private userService: UsuarioService,
    breakpointObserver: BreakpointObserver,
    private transferRequestService: SolicitudTrasladoService,
    fb: FormBuilder,
    private adapter: DateAdapter<any>,
    private contractService: ContratoService
  ) {
    this.adapter.setLocale("es-mx");
    breakpointObserver.observe(["(max-width: 600px)"]).subscribe((result) => {
      this.displayedColumns = result.matches
        ? [
            "folio",
            "elaboracion",
            "ftp",
            "origen",
            "destino",
            "nombre_pac",
            "fecha",
            "accion",
          ]
        : [
            "folio",
            "elaboracion",
            "ftp",
            "origen",
            "destino",
            "nombre_pac",
            "fecha",
            "accion",
          ];
    });
    this.userToken = jwt_decode(<string>localStorage.getItem("token"));
    this.range = fb.group({
      start: new FormControl(null),
      end: new FormControl(null),
    });
  }

  ngOnInit() {
    this.userService
      .unidadUsuario(this.userToken.UsuarioId)
      .subscribe((res) => {
        this.unitMedic = res.Data;
        this.unitMedicalId = this.unitMedic.unidadMed.Id.toString();
        this.consultCurrentContract();
      });
    this.paginator.page.subscribe((page) => {
      if (page) {
        this.pageIndex = page["pageIndex"] + 1;
        this.pageSize = page["pageSize"];
        this.consultRequest();
      }
    });
    this.paginator._intl.itemsPerPageLabel = "Resultados por pagina";
    this.paginator._intl.nextPageLabel = "Siguiente página";
    this.paginator._intl.lastPageLabel = "Última página";
    this.paginator._intl.previousPageLabel = "Pagina anterior";
    this.paginator._intl.firstPageLabel = "Primera página";
    this.searchInputControl.valueChanges.subscribe((res) => {
      if (!res) {
        this.consultRequest();
      }
    });
  }

  consultCurrentContract() {
    const now = moment().format("YYYY-MM-DDTHH:mm:ssZ");
    this.contractService
      .ConsultarContratoTraslado(this.unitMedicalId, now)
      .subscribe((result) => {
        if (result.Data != null) {
          this.contracts = result.Data;
          console.log(result.Data);
          this.determineFilters();
          this.determineRangeDates();
        } else {
          void Swal.fire({
            icon: "error",
            title: "Sin Resultados",
            text: "El ultimo contrato vigente esta vencido, puedes consultar solicitudes anteriores cambiando el rango de fechas para busqueda.",
          });
        }
      });
  }

  determineFilters() {
    const temporal = this.contracts.map((contract) => contract.NumeroContrato);
    this.filterNumberContract = temporal.join(",");
    if (this.contracts.length === 1) {
      this.controlNumberContract.setValue(this.filterNumberContract);
      this.controlNumberContract.disable();
    }
    this.consultRequest();
    this.consultOriginsFilter(
      this.unitMedicalId,
      "true",
      "",
      "",
      this.filterNumberContract,
      this.dateStart,
      this.dateEnd
    );
    this.consultDestinationsFilter(
      this.unitMedicalId,
      "true",
      "",
      "",
      this.filterNumberContract,
      this.dateStart,
      this.dateEnd
    );
    this.consultCompanionFilter(
      this.unitMedicalId,
      "true",
      "",
      "",
      this.filterNumberContract,
      this.dateStart,
      this.dateEnd
    );
  }

  determineRangeDates() {
    if (this.contracts.length > 1) {
      const dateMin = this.contracts
        .map((contract) => contract.InicioContrato)
        .reduce(function (valor1, valor2) {
          return new Date(valor1) < new Date(valor2) ? valor1 : valor2;
        });
      const dateMomentObjectMin = moment(dateMin, "DD/MM/YYYY");
      this.minDate = dateMomentObjectMin.toDate();
      const dateMax = this.contracts
        .map((contract) => contract.FinContrato)
        .reduce(function (valor1, valor2) {
          return new Date(valor1) > new Date(valor2) ? valor1 : valor2;
        });
      const dateMomentObject = moment(dateMax, "DD/MM/YYYY");
      this.maxDate = dateMomentObject.toDate();
    } else {
      const dateMomentObjectMin = moment(
        this.contracts[0].InicioContrato,
        "DD/MM/YYYY"
      );
      this.minDate = dateMomentObjectMin.toDate();
      const dateMomentObjectMax = moment(
        this.contracts[0].FinContrato,
        "DD/MM/YYYY"
      );
      this.maxDate = dateMomentObjectMax.toDate();
    }
  }

  consultCompanionFilter(
    unitId,
    state,
    origin,
    destination,
    contractsNumber,
    startDate,
    endDate: string
  ) {
    this.transferRequestService
      .GetCompanionFilters(
        unitId,
        state,
        origin,
        destination,
        contractsNumber,
        startDate,
        endDate
      )
      .subscribe((res) => {
        const temporal = res.Data.map((orig) => orig.acompañante.Bool);
        this.defineCompanionFilterStatus(temporal);
      });
  }

  defineCompanionFilterStatus(status: boolean[]) {
    this.withCompanion = true;
    this.unaccompanied = true;
    for (let entry of status) {
      if (entry) {
        this.withCompanion = false;
      } else {
        this.unaccompanied = false;
      }
    }
  }

  consultOriginsFilter(
    unitId,
    state,
    destination,
    companion,
    numberContracts,
    startDate,
    endDate: string
  ) {
    this.transferRequestService
      .GetOriginFilters(
        unitId,
        state,
        destination,
        companion,
        numberContracts,
        startDate,
        endDate
      )
      .subscribe((res) => {
        this.origins = res.Data.map((orig) => orig.origen);
      });
  }

  consultDestinationsFilter(
    unitId,
    state,
    origin,
    companion,
    numberContracts,
    startDate,
    endDate: string
  ) {
    this.transferRequestService
      .GetDestinationFilters(
        unitId,
        state,
        origin,
        companion,
        numberContracts,
        startDate,
        endDate
      )
      .subscribe((res) => {
        this.destinations = res.Data.map((orig) => orig.destino);
      });
  }

  // searchByDate() {
  // }
  removeChips(filter: RequestFiltersModel): void {
    const index = this.filters.indexOf(filter);
    if (index >= 0) {
      this.filters.splice(index, 1);
      this.resetValueFilters(filter.filter);
      this.consultRequest();
    }
  }

  resetValueFilters(filter: string) {
    switch (filter) {
      case "Acompañante":
        this.resetFilterCompanion();
        break;
      case "Destino":
        this.resetFilterDestination();
        break;
      case "Origen":
        this.resetFilterOrigin();
        break;
      case "Numero Contrato":
        this.resetFilterContract();
        break;
    }
  }

  resetFilterCompanion() {
    this.controlCompanion.reset();
    this.resetAllFilters();
  }

  resetFilterDestination() {
    this.controlDestination.reset();
    this.resetAllFilters();
  }

  resetFilterOrigin() {
    this.controlOrigin.reset();
    this.resetAllFilters();
  }

  resetFilterContract() {
    this.controlNumberContract.reset();
    this.consultCurrentContract();
    this.resetAllFilters();
  }

  resetFilterDates() {
    this.range.reset();
    this.range.markAsUntouched();
    this.range.markAsPristine();
    this.dateStart = "";
    this.dateEnd = "";
    this.resetAllFilters();
    this.consultRequest();
  }

  resetAllFilters() {
    const filterOrigin =
      this.controlOrigin.value !== null ? this.controlOrigin.value : "";
    const filterDestination =
      this.controlDestination.value !== null
        ? this.controlDestination.value
        : "";
    const filterCompanion =
      this.controlCompanion.value !== null ? this.controlCompanion.value : "";
    this.consultOriginsFilter(
      this.unitMedicalId,
      "true",
      filterDestination,
      filterCompanion,
      this.filterNumberContract,
      this.dateStart,
      this.dateEnd
    );
    this.consultDestinationsFilter(
      this.unitMedicalId,
      "true",
      filterOrigin,
      filterCompanion,
      this.filterNumberContract,
      this.dateStart,
      this.dateEnd
    );
    this.consultCompanionFilter(
      this.unitMedicalId,
      "true",
      filterOrigin,
      filterDestination,
      this.filterNumberContract,
      this.dateStart,
      this.dateEnd
    );
  }

  changeOrigin(chipNew: RequestFiltersModel) {
    this.checkRepeatNotChips(chipNew);
    this.resetAllFilters();
  }

  changeDestination(chipNew: RequestFiltersModel) {
    this.checkRepeatNotChips(chipNew);
    this.resetAllFilters();
  }

  changeCompanion(chipNew: RequestFiltersModel) {
    const filterCompanion =
      this.controlCompanion.value !== null ? this.controlCompanion.value : "";
    const filterOrigin =
      this.controlOrigin.value !== null ? this.controlOrigin.value : "";
    const filterDestination =
      this.controlDestination.value !== null
        ? this.controlDestination.value
        : "";
    this.consultDestinationsFilter(
      this.unitMedicalId,
      "true",
      filterOrigin,
      filterCompanion,
      this.filterNumberContract,
      this.dateStart,
      this.dateEnd
    );
    this.consultOriginsFilter(
      this.unitMedicalId,
      "true",
      filterDestination,
      filterCompanion,
      this.filterNumberContract,
      this.dateStart,
      this.dateEnd
    );
    this.checkRepeatNotChips(chipNew);
  }

  changeContractNumber(chipNew: RequestFiltersModel) {
    this.filterNumberContract = chipNew.value;
    this.checkRepeatNotChips(chipNew);
    this.resetAllFilters();
  }

  checkRepeatNotChips(chipNew: RequestFiltersModel) {
    if (!this.filters.find((chip) => chip.filter === chipNew.filter)) {
      this.addChips(chipNew);
    } else {
      this.updateValueChip(chipNew);
    }
  }

  addChips(chipNew: RequestFiltersModel) {
    this.filters.push(chipNew);
    this.consultRequest();
  }

  updateValueChip(chipUpdate: RequestFiltersModel) {
    const index = this.filters.findIndex(
      (chip) => chip.filter === chipUpdate.filter
    );
    this.filters[index].value = chipUpdate.value;
    this.consultRequest();
  }

  consultRequest() {
    const filter =
      this.searchInputControl.value !== null
        ? this.searchInputControl.value
        : "";
    const filterCompanion =
      this.controlCompanion.value !== null ? this.controlCompanion.value : "";
    const filterOrigin =
      this.controlOrigin.value !== null ? this.controlOrigin.value : "";
    const filterDestination =
      this.controlDestination.value !== null
        ? this.controlDestination.value
        : "";
    this.transferRequestService
      .ConsultSolicitudes(
        this.unitMedicalId,
        this.pageIndex.toString(),
        this.pageSize.toString(),
        filter,
        "true",
        filterCompanion,
        filterOrigin,
        filterDestination,
        this.filterNumberContract,
        this.dateStart,
        this.dateEnd
      )
      .subscribe((res) => {
        console.log(res.Data);
        this.dataSource = new MatTableDataSource<Solicitudtraslado>(
          res.Data["registros"]
        );
        this.length = res.Data["total"];
        this.loadingData = true;
        if (this.dataSource.data.length === 0) {
          void Swal.fire({
            icon: "error",
            title: "Sin Resultados",
            text: "No se encontraron resultados para tu busqueda",
          });
        }
      });
  }

  formatDate(valor: string): string {
    return moment(valor).format("DD-MM-YYYY");
  }

  formatExpeditionDate(valor: string): string {
    return moment(valor).format("DD-MM-YYYY");
  }

  getPDF(folio: string) {
    this.transferRequestService.printRequestComplete(folio, "true",'true','').subscribe(
      (res) => {
        const downloadURL = URL.createObjectURL(res);
        window.open(downloadURL, "Solituddetraslado");
      },
      (err) => {
        if (err.error !== "" || err.error !== undefined) {
          void Swal.fire({
            position: "center",
            title: "Error",
            text: "No se pudo descargar la plantilla",
            icon: "error",
            showConfirmButton: true,
          });
        }
      }
    );
  }
  imprimirSoloDatos(folio: string) {
    this.transferRequestService.printDataOnly(folio,'true','').subscribe(res => {
      const downloadURL = URL.createObjectURL(res);
      window.open(downloadURL);
    }, err => {
      if (err.error !== '' || err.error !== undefined) {
        Swal.fire({
          position: 'center',
          title: 'Error',
          text: 'No se pudo descargar la plantilla',
          icon: 'error',
          showConfirmButton: true,
        });
      }
    });
  }

  rangeDateChange() {
    if (this.range.get("end").value) {
      console.log("Filtrando por rango");
      this.dateStart = this.range.get("start").value.toISOString();
      this.dateEnd = this.range.get("end").value.toISOString();
      this.consultRequest();
      this.resetAllFilters();
    }
  }
}
