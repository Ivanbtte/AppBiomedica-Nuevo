import { Component, OnInit, ViewEncapsulation } from "@angular/core";
import { FtpService } from "../../../core/servicios/ServiciosSubrogados/traslados/ftp.service";
import Swal from "sweetalert2";
import { RespuestaPeticion } from "../../../core/models/estructuras_generales";

@Component({
  selector: "app-cargar-ftp",
  templateUrl: "./cargar-ftp.component.html",
  styleUrls: ["./cargar-ftp.component.css"],
  encapsulation: ViewEncapsulation.None,
})
export class CargarFtpComponent implements OnInit {
  files: File[] = [];
  public validado = false;

  constructor(private ftpService: FtpService) {}

  ngOnInit(): void {}
  onSelect(event: any) {
    if (this.files.length < 1) {
      this.files.push(...event.addedFiles);
    }
  }
  onRemove(event: any) {
    this.files.splice(this.files.indexOf(event), 1);
    this.validado = false;
  }
  enviarArchivo() {
    this.validado = true;
    const formData = new FormData();
    formData.append("file", this.files[0]);
    this.ftpService.CargarFolios(formData).subscribe(
      (result) => {
        Swal.fire({
          position: "center",
          title: result.Mensaje,
          text: `${result.Data} folios`,
          icon: "success",
          showConfirmButton: true,
        });
        this.validado = false;
        this.files = [];
      },
      (err) => {
        if (err.error !== "" || err.error !== undefined) {
          const erores = <RespuestaPeticion>err.error;
          Swal.fire({
            position: "center",
            title: erores.Mensaje,
            icon: "error",
            showConfirmButton: true,
          });
        }
      }
    );
  }
}
