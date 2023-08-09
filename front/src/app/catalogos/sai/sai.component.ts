import { Component, OnInit, ViewChild, ViewEncapsulation } from "@angular/core";
import { SaiService } from "../servicios/sai.service";
import { MatPaginator } from "@angular/material/paginator";
import {
  Tema,
  TemaGrupo,
} from "../../core/models/interfaceCatalogos/filtros.interface";
import { FiltrosService } from "../servicios/filtros.service";
import { FormControl } from "@angular/forms";
import { Observable } from "rxjs";
import { map, startWith } from "rxjs/operators";
import { MatDialog } from "@angular/material/dialog";
import { BusquedaSelectorComponent } from "../busqueda-selector/busqueda-selector.component";
import { DetalleSaiComponent } from "../detalle-sai/detalle-sai.component";
import { AbrirCatalogoSaiService } from "../../core/servicios/abrir-catalogo-sai.service";
import Swal from "sweetalert2";
import { ClavesService } from "../../core/servicios/solicitud/claves.service";
import { RespuestaPeticion } from "../../core/models/estructuras_generales";
import { CantidadDialogComponent } from "../../shared/Dialogs/cantidad-dialog/cantidad-dialog.component";
import { FiltrosSaiService } from "../servicios/filtros-sai.service";
import { MatSort } from "@angular/material/sort";
import { MatTableDataSource } from "@angular/material/table";

@Component({
  selector: "app-sai",
  templateUrl: "./sai.component.html",
  styleUrls: ["./sai.component.scss"],
})
export class SaiComponent implements OnInit {
  myControl = new FormControl();
  control_generico = new FormControl();
  control_nombre = new FormControl();
  control_especifico = new FormControl();
  displayedColumns: string[] = [
    "Clave",
    "Descripcion",
    "UnidadPres",
    "Cantidad",
    "Tipo",
    "Vista",
    "Agregar",
  ];
  cargando = false;
  dataSource: any;
  grupo_id = "";
  especifico_id = "";
  temas_id = "";
  generico_id = "";
  chipsVer: Tema[] = [];
  selectable = true;
  removable = true;
  Filtro: TemaGrupo[] = [];
  filtroTema: TemaGrupo[] = [];
  arreglo_genericos: string[] = [];
  arreglo_especificos: string[] = [];
  filteredOptions!: Observable<TemaGrupo[]>;
  filteredGenericos!: Observable<string[]>;
  filteredEspecificos!: Observable<string[]>;
  BanderaPorGrupo = false;
  lastFilter = "";
  lastFilterGenerico = "";
  lastFilterEspecifico = "";
  @ViewChild(MatPaginator, { static: true }) paginator!: MatPaginator;
  tamanioPagina = 50;
  tamanios = [20, 50, 100];
  totalRegistros: number = 0;
  paginaActual = 1;
  /* *****  variables para el apartado de las solicitudes   *****   */
  estado: boolean;
  cuadro = "";
  public fecha: any;
  data: any[] = [];

  constructor(
    private _sai: SaiService,
    private _filtro: FiltrosService,
    public dialog: MatDialog,
    private _abrirSai: AbrirCatalogoSaiService,
    private _conceptos: ClavesService,
    private _sai_filtros: FiltrosSaiService
  ) {
    this.estado = this._abrirSai.obtenerEstado();
    this.cuadro = this._abrirSai.obtenerCuadro();
    const date = new Date();
    this.fecha =
      date.getUTCDate() +
      "/" +
      date.getUTCMonth() +
      "/" +
      date.getUTCFullYear();
  }

  openDialog(): void {
    const dialogRef = this.dialog.open(BusquedaSelectorComponent, {
      height: "auto",
      width: "95%",
      disableClose: true,
      data: { datos: this.filtroTema },
    });

    dialogRef.afterClosed().subscribe((result) => {
      console.log(result + "<-----resultado");
      if (
        (this.chipsVer.length === 0 && result) ||
        (this.chipsVer.length !== 0 && result)
      ) {
        this.grupo_id = "";
        this.generico_id = "";
        if (
          result.length !== 0 &&
          dialogRef.componentInstance.filtroGrupo === false
        ) {
          console.log("Entre a busqueda por temas ");
          this.chipsVer = result;
          this.buscarPorTema();
        }
        if (result.length !== 0 && dialogRef.componentInstance.filtroGrupo) {
          this.BanderaPorGrupo = true;
          console.log("Entre a busqueda por grupos ");
          this.chipsVer = result;
          console.log(this.chipsVer);
          this.grupo_id = this.chipsVer[0].Id.toString();
          this.consultarGenerico(this.grupo_id, false);
          this.buscarPorNombreGrupo();
        }
      }
    });
  }

  openDialogClave(id: string): void {
    const dialogRef = this.dialog.open(DetalleSaiComponent, {
      height: "auto",
      width: "auto",
      data: {
        id: id,
        estadoAgregar: this.estado,
        estadoCantidad: this.estado,
      },
    });

    dialogRef.afterClosed().subscribe((result) => {});
  }

  openDialogCantidad(clave: any): void {
    const idSolicitud = <string>localStorage.getItem("solicitudId");
    const idServicio = <string>localStorage.getItem("servicioId");
    const idUnidad = <string>localStorage.getItem("unidadId");
    const dialogRef = this.dialog.open(CantidadDialogComponent, {
      height: "85%",
      width: "70%",
      disableClose: true,
      data: {
        clave: clave,
        editar: false,
        agregar: true,
        aprobar: false,
      },
    });
    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        console.log("entreee", result);
        const total = clave.Precio * result;
        console.log("totoal:", total);
        this._conceptos
          .AgregarClave(
            {
              SolicitudId: idSolicitud,
              ClaveId:
                clave.Grupo +
                clave.Generico +
                clave.Especifico +
                clave.Diferenciador +
                clave.Variable,
              Descripcion: clave.Descripcion,
              CantidadSolicitada: parseInt(result, 10),
              Precio: clave.Precio,
              Total: total,
              ServiciosProformaId: idServicio,
              EstatusId: 1,
              Fecha: this.fecha,
            },
            idUnidad
          )
          .subscribe(
            (res) => {
              console.log("Agregado Correctamente", res.Mensaje);
              Swal.fire({
                icon: "success",
                title: "Correcto",
                text: res.Mensaje,
              });
            },
            (err) => {
              if (err.error !== "" || err.error !== undefined) {
                const erores = <RespuestaPeticion>err.error;
                Swal.fire({
                  icon: "error",
                  title: "Error",
                  text: erores.Mensaje,
                });
              }
              console.log("ocurrio un error", err);
            }
          );
      }
    });
  }

  ngOnInit() {
    this.llenarArreglo();
    this._filtro.getGrupoTemas("true", "").subscribe(
      (res) => {
        this.Filtro = res.Data;
      },
      (err) => {
        console.log(err);
      }
    );

    this._filtro.getGrupoTemas("false", "").subscribe(
      (res) => {
        this.filtroTema = res.Data;
      },
      (err) => {
        console.log(err);
      }
    );
    this.filteredOptions = this.myControl.valueChanges.pipe(
      startWith<string | TemaGrupo[]>(""),
      map((value) => (typeof value === "string" ? value : this.lastFilter)),
      map((filter) => this._filter(filter))
    );
    this.filteredGenericos = this.control_generico.valueChanges.pipe(
      startWith<string>(""),
      map((value2) =>
        typeof value2 === "string" ? value2 : this.lastFilterGenerico
      ),
      map((filter2) => this._filter_genericos(filter2))
    );
    this.filteredEspecificos = this.control_especifico.valueChanges.pipe(
      startWith<string>(""),
      map((value3) =>
        typeof value3 === "string" ? value3 : this.lastFilterEspecifico
      ),
      map((filter3) => this._filter_especificos(filter3))
    );
    this.control_generico.valueChanges.subscribe((res) => {
      if (res === "") {
        this.borrarFiltroGenerico();
      }
    });
    this.myControl.valueChanges.subscribe((res) => {
      if (res === "") {
        this.borrarGrupo();
      }
    });
    this.control_especifico.valueChanges.subscribe((res) => {
      if (res === "") {
        this.borrarFiltroEspecifico();
      }
    });
    this.paginator.page.subscribe((page) => {
      console.log("wewewewewewe", page);
      if (page) {
        this.paginaActual = page["pageIndex"] + 1;
        this.tamanioPagina = page["pageSize"];
        this.llenarArreglo();
      }
    });
    this.paginator._intl.itemsPerPageLabel = "Resultados por pagina";
    if (!this.estado) {
      this.displayedColumns.pop();
    }
  }

  cargarOpcionesAutocompleteGenerico() {
    this.filteredGenericos = this.control_generico.valueChanges.pipe(
      startWith<string>(""),
      map((value2) =>
        typeof value2 === "string" ? value2 : this.lastFilterGenerico
      ),
      map((filter2) => this._filter_genericos(filter2))
    );
  }

  cargarOpcionesAutocompleteGrupo() {
    this.filteredOptions = this.myControl.valueChanges.pipe(
      startWith<string | TemaGrupo[]>(""),
      map((value) => (typeof value === "string" ? value : this.lastFilter)),
      map((filter) => this._filter(filter))
    );
  }

  cargarOpcionesAutocompleteEspecifico() {
    this.filteredEspecificos = this.control_especifico.valueChanges.pipe(
      startWith<string>(""),
      map((value3) =>
        typeof value3 === "string" ? value3 : this.lastFilterEspecifico
      ),
      map((filter3) => this._filter_especificos(filter3))
    );
  }

  buscarPorTema() {
    const ids = this.chipsVer.map((it) => it.Id);
    this.temas_id = ids.join();
    this.grupo_id = "";
    this.generico_id = "";
    this.control_generico.reset();
    this.myControl.reset();
    this.control_nombre.reset();
    this.paginator.firstPage();
    this.BanderaPorGrupo = false;
    this._filtro.getGrupoTemas("true", this.temas_id).subscribe((res) => {
      this.Filtro = res.Data;
      console.log(this.Filtro);
      this.cargarOpcionesAutocompleteGrupo();
    });
    this.llenarArreglo();
  }

  buscarPorNombre() {
    this.llenarArreglo();
    this.paginator.firstPage();
  }

  /* *****  busqueda solo por el campo de generico  *****   */
  buscarPorGenericoSolamente() {
    if (this.grupo_id === "" && this.control_generico.value !== null) {
      console.log("buscaras solo por generico");
      this.generico_id = this.control_generico.value;
      this.llenarArreglo();
    }
  }

  borrarFiltroGenerico() {
    if (this.BanderaPorGrupo === false) {
      if (this.grupo_id !== "") {
        this.grupo_id = this.grupo_id !== null ? this.myControl.value.Grupo : 0;
      }
    }
    this.control_nombre.reset();
    this.control_generico.reset();
    this.generico_id = "";
    this.paginator.firstPage();
    this.llenarArreglo();
    this.cargarOpcionesAutocompleteGenerico();
    // this.consultarGenerico(this.grupo_id, false);
  }

  borrarFiltroEspecifico() {
    this.control_especifico.reset();
    this.especifico_id = "";
    this.llenarArreglo();
  }
  keytab(event) {
    console.log("hasasas");
    let element = event.srcElement.nextElementSibling; // get the sibling element

    if (element == null)
      // check if its null
      return;
    else element.focus(); // focus if not null
  }
  buscarPorGrupo() {
    this.grupo_id = this.grupo_id !== null ? this.myControl.value.Grupo : 0;
    this.control_nombre.reset();
    this.control_generico.reset();
    this.generico_id = "";
    this.control_especifico.reset();
    this.especifico_id = "";
    this.paginator.firstPage();
    this.llenarArreglo();
    this.consultarGenerico(this.grupo_id, false);
  }

  buscarPorEspecifico() {
    if (this.control_especifico.value !== null) {
      console.log("buscando por especifico");
      this.especifico_id = this.control_especifico.value;
      this.llenarArreglo();
    }
  }

  buscarPorGenerico() {
    this.generico_id =
      this.generico_id !== null ? this.control_generico.value : 0;
    this.control_nombre.reset();
    this.paginator.firstPage();
    this.consultarEspecificos(this.generico_id, false);
    this.llenarArreglo();
  }

  consultarGenerico(grupo: string, todos_genericos: boolean) {
    let resultado = [];
    this._sai_filtros
      .consultarFiltrosGrupos(grupo, "", todos_genericos, false)
      .subscribe((res) => {
        resultado = res.Data;
        this.arreglo_genericos = resultado.map(
          (it: { Generico: any }) => it.Generico
        );
        this.cargarOpcionesAutocompleteGenerico();
      });
  }

  consultarEspecificos(generico: string, todos_especificos: boolean) {
    let resultado2 = [];
    this._sai_filtros
      .consultarFiltrosGrupos(this.grupo_id, generico, false, todos_especificos)
      .subscribe((res) => {
        resultado2 = res.Data;
        this.arreglo_especificos = resultado2.map(
          (it: { Especifico: any }) => it.Especifico
        );
        console.log(this.arreglo_especificos);
        this.cargarOpcionesAutocompleteEspecifico();
      });
  }

  buscarPorNombreGrupo() {
    this.myControl.reset();
    this.control_generico.reset();
    this.temas_id = "";
    this.control_nombre.reset();
    this.paginator.firstPage();
    this.llenarArreglo();
  }

  llenarArreglo() {
    const nombre =
      this.control_nombre.value !== null ? this.control_nombre.value : "";
    this._sai
      .getSai(
        this.grupo_id,
        this.temas_id,
        nombre,
        this.paginaActual.toString(),
        this.tamanioPagina.toString(),
        this.cuadro,
        this.generico_id,
        this.especifico_id
      )
      .subscribe(
        (res) => {
          this.data = res.Data["registros"];
          this.dataSource = new MatTableDataSource();
          this.dataSource.data = this.data;
          this.totalRegistros = res.Data["total"];
          this.cargando = true;
          console.log(res.Data["registros"]);
          if (this.totalRegistros === 0) {
            Swal.fire({
              icon: "error",
              title: "Sin Resultados",
              text: "Lo siento no encontramos ningun resultado con tu busqueda",
            });
          }
        },
        (err) => {
          console.log(err);
        }
      );
    this.cargando = false;
  }

  private _filter(filter: string): TemaGrupo[] {
    if (filter) {
      this.lastFilter = filter.toLowerCase();
      return this.Filtro.filter(
        (option) => option.Grupo.toLowerCase().indexOf(this.lastFilter) === 0
      );
    } else {
      return this.Filtro;
    }
  }

  private _filter_genericos(value: string): string[] {
    if (value) {
      this.lastFilterGenerico = value.toLowerCase();
      return this.arreglo_genericos.filter(
        (option) => option.toLowerCase().indexOf(this.lastFilterGenerico) === 0
      );
    } else {
      return this.arreglo_genericos;
    }
  }

  private _filter_especificos(value: string): string[] {
    if (value) {
      this.lastFilterEspecifico = value.toLowerCase();
      return this.arreglo_especificos.filter(
        (option) =>
          option.toLowerCase().indexOf(this.lastFilterEspecifico) === 0
      );
    } else {
      return this.arreglo_especificos;
    }
  }

  borrarTodosChips(): void {
    this.chipsVer = [];
  }

  remove(chips: Tema): void {
    const index = this.chipsVer.indexOf(chips);
    if (index >= 0) {
      this.chipsVer.splice(index, 1);
      this.buscarPorTema();
    }
  }

  displayFn(tema: TemaGrupo) {
    return tema ? tema.Grupo : "";
  }

  borrarGrupo() {
    this.grupo_id = "";
    this.myControl.reset();
    this.generico_id = "";
    this.control_generico.reset();
    this.control_especifico.reset();
    this.especifico_id = "";
    this.arreglo_genericos = [];
    this.llenarArreglo();
    this.cargarOpcionesAutocompleteGrupo();
  }
}
