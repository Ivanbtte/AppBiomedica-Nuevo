import {
  Component,
  Inject,
  OnInit,
  ViewChild,
  AfterViewInit,
} from "@angular/core";
import {
  MAT_DIALOG_DATA,
  MatDialog,
  MatDialogRef,
} from "@angular/material/dialog";
import { MatPaginator, PageEvent } from "@angular/material/paginator";
import { MatSort } from "@angular/material/sort";
import { ProveedoresInterface } from "../../../../core/models/proveedores.interface";
import { MatTableDataSource } from "@angular/material/table";
import { AgregarProveedorComponent } from "../../agregar-proveedor/agregar-proveedor.component";
import { DetalleProveedorComponent } from "../../detalle-proveedor/detalle-proveedor.component";
import Swal from "sweetalert2";
import { FormControl } from "@angular/forms";
import { ProveedoresService } from "../../../../core/servicios/ProveedoresServicios/proveedores.service";

export interface DialogData {
  proveedor: any;
}

@Component({
  selector: "app-seleccionar-proveedor-representante",
  templateUrl: "./seleccionar-proveedor-representante.component.html",
  styleUrls: ["./seleccionar-proveedor-representante.component.css"],
})
export class SeleccionarProveedorRepresentanteComponent
  implements OnInit, AfterViewInit
{
  // ********************************************************************************************************
  /* ***** Variables para la tabla  *****   */
  // dataSource: any;
  //dataSource: MatTableDataSource<ProveedoresInterface>;
  dataSource = new MatTableDataSource<ProveedoresInterface>();

  @ViewChild(MatPaginator) paginator: MatPaginator;
  tamanioPagina = 20;
  tamanios = [20, 50, 100];
  totalRegistros = 0;
  paginaActual = 1;
  displayedColumns: string[] = [
    "DetallePro",
    "Alias",
    "Estado",
    "Municipio",
    "CorreoPro",
    "RFC",
    "Telefono",
    "SeleccionarPro",
  ];
  cargando = false;
  // ********************************************************************************************************
  /* *****  Variables para el servicio de consultar proveedores  *****   */
  id_proveedor = "";
  nombre_proveedor = "";
  control_nombre = new FormControl();

  constructor(
    public dialogRef: MatDialogRef<SeleccionarProveedorRepresentanteComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
    private _proveedores: ProveedoresService,
    public dialog: MatDialog
  ) {}
  ngAfterViewInit() {
    this.dataSource.paginator = this.paginator;
    this.dataSource.paginator.page.subscribe((page) => {
      if (page) {
        this.paginaActual = page["pageIndex"] + 1;
        this.tamanioPagina = page["pageSize"];
        this.llenarTabla();
      }
    });
    this.paginator._intl.itemsPerPageLabel = "Resultados por pagina";
  }
  ngOnInit() {
    this.llenarTabla();
    this.control_nombre.valueChanges.subscribe((result) => {
      console.log(result);
      if (result === "" || result === null) {
        this.llenarTabla();
      }
    });
  }

  llenarTabla() {
    const nombre =
      this.control_nombre.value !== null ? this.control_nombre.value : "";
    this._proveedores
      .ConsultarProveedores(
        this.id_proveedor,
        nombre,
        this.paginaActual.toString(),
        this.tamanioPagina.toString()
      )
      .subscribe((resultado) => {
        console.log(resultado);
        const arregloProv: ProveedoresInterface[] = resultado.Data["registros"];
        this.dataSource = new MatTableDataSource();
        this.dataSource = new MatTableDataSource(arregloProv);
        this.totalRegistros = resultado.Data["total"];
        this.cargando = true;
      });
  }

  openDialogAgregar(): void {
    const dialogRef = this.dialog.open(AgregarProveedorComponent, {
      height: "auto",
      width: "80%",
      disableClose: true,
      data: {
        editar: false,
        proveedor: "",
      },
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        this.llenarTabla();
      }
    });
  }

  SelecionarProveedor(proveedor: ProveedoresInterface) {
    Swal.fire({
      title: "Has decidido seleccionar a este proveedor",
      text: proveedor.AliasEmpresa,
      icon: "warning",
      showCancelButton: true,
      confirmButtonColor: "#3085d6",
      cancelButtonColor: "#a2a2a2",
      confirmButtonText: "Si, Seleccionar!",
      cancelButtonText: "Cancelar",
      reverseButtons: true,
    }).then((result) => {
      if (result.value) {
        this.dialogRef.close(proveedor);
      }
    });
  }

  openDialogDetalleProveedor(provee: ProveedoresInterface): void {
    const dialogRef = this.dialog.open(DetalleProveedorComponent, {
      height: "auto",
      width: "40%",
      disableClose: false,
      data: {
        proveedor: provee,
      },
    });

    dialogRef.afterClosed().subscribe((result) => {});
  }

  buscarPorNombre() {
    this.llenarTabla();
    this.paginator.firstPage();
  }

  cerrarDialog() {
    this.dialogRef.close();
  }
}
