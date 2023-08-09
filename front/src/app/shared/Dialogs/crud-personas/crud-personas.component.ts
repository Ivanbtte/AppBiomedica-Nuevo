import {Component, Inject, OnInit, ViewChild} from '@angular/core';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {DialogData} from '../Contratos/agregar-editar-contrato/agregar-editar-contrato.component';
import {MatPaginator} from '@angular/material/paginator';
import {MatSort} from '@angular/material/sort';
import {PersonaService} from '../../../authentication/servicios/persona.service';
import {Personas} from '../../../authentication/interfaces/usuario.interface';
import {MatTableDataSource} from '@angular/material/table';
import Swal from 'sweetalert2';


@Component({
  selector: 'app-crud-personas',
  templateUrl: './crud-personas.component.html',
  styleUrls: ['./crud-personas.component.css']
})
export class CrudPersonasComponent implements OnInit {
  public formulario_persona!: FormGroup;
  estadoBoton = false;
  public mensajeErrorCorreo = 'example@imss.gob.mx';
  public botonVisible = true;
  cerrarPanel = false;
  // ********************************************************************************************************
  /* ***** Variables para la tabla  *****   */
  dataSource: any;
  @ViewChild(MatPaginator) paginator!: MatPaginator;
  @ViewChild(MatSort) sort!: MatSort;
  tamanioPagina = 50;
  tamanios = [20, 50, 100];
  totalRegistros = 0;
  paginaActual = 1;
  displayedColumns: string[] = ['Nombre', 'ApellidoP', 'ApellidoM', 'Correo', 'Telefono', 'Matricula', 'SeleccionarPer'];
  cargando = false;

  constructor(
    public dialogRef: MatDialogRef<CrudPersonasComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
    private fb: FormBuilder,
    private _personas: PersonaService
  ) {
  }

  ngOnInit() {
    this.crear_formulario_persona();
    this.llenarTabla();

  }

  llenarTabla() {
    this._personas.getAllPersona(this.paginaActual.toString(), this.tamanioPagina.toString()).subscribe(resultado => {
      const arregloResultado: Personas = resultado.Data['registros'];
      this.dataSource = new MatTableDataSource();
      this.dataSource.data = arregloResultado;
      this.totalRegistros = resultado.Data['total'];
      this.cargando = true;
      console.log(arregloResultado);
    });
  }

  // Workaround for angular component issue #13870
  disableAnimation = true;

  ngAfterViewInit(): void {
    // timeout required to avoid the dreaded 'ExpressionChangedAfterItHasBeenCheckedError'
    setTimeout(() => this.disableAnimation = false);
  }

  crear_formulario_persona() {
    this.formulario_persona = this.fb.group({
      Nombre: [null, Validators.compose([Validators.pattern('^[a-zA-ZáéíóúÁÉÍÓÚñÑ\\s]*$'),
        Validators.required])],
      ApellidoPat: [null, Validators.compose([Validators.pattern('^[a-zA-ZáéíóúÁÉÍÓÚñÑ\\s]*$'),
        Validators.required])],
      ApellidoMat: [null, Validators.compose([Validators.pattern('^[a-zA-ZáéíóúÁÉÍÓÚñÑ\\s]*$'),
        Validators.required])],
      FechaNac: [null, Validators.compose([Validators.required])],
      Genero: [null, Validators.compose([Validators.required])],
      Correo: [null, Validators.compose([Validators.required, Validators.pattern('^[_a-z0-9-]+(.[_a-z0-9-]+)*@imss.gob.mx$')])
      ],
      Telefono: [null, Validators.compose([Validators.pattern('^([0-9])*$')])],
      Extension: [null, Validators.compose([Validators.pattern('^([0-9])*$')])],
      Matricula: [null, Validators.compose([Validators.pattern('^([0-9])*$')])],
      Celular: [null, Validators.compose([Validators.pattern('^([0-9])*$')])]
    });
  }

  cambiarValidaCorreo() {
    this.formulario_persona.removeControl('Correo');
    this.formulario_persona.addControl('Correo', this.fb.control('', Validators.compose([Validators.required, Validators.email])));
    this.mensajeErrorCorreo = 'Correo invalido';
    this.botonVisible = false;
    this.estadoBoton = true;
  }

  onNoClick(): void {
    this.dialogRef.close();
  }

  applyFilter(filterValue: string) {
    this.dataSource.filter = filterValue.trim().toLowerCase();
    console.log(this.dataSource.filteredData);
  }

  agregar_persona() {
    const datos = <Personas> this.formulario_persona.getRawValue();
    console.log(datos);
    Swal.fire({
      title: 'Estas seguro',
      text: 'Deseas agregar esta persona',
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      confirmButtonText: 'Si Agregar'
    }).then((result) => {
      if (result.value) {
        this.cerrarPanel = true;
        this._personas.AgregarPersona(datos).subscribe(resultado => {
          console.log(resultado.Data);
          Swal.fire(
            'Agregado',
            'La persona ha sido agregada correctamente',
            'success'
          );
          this.dialogRef.close(resultado.Data);
          this.cerrarPanel = false;
        });
      }
    });
  }

  Selecionar(persona: Personas) {
    Swal.fire({
      title: 'Seleccionar',
      text: 'Persona:' + persona.Nombre + ' ' + persona.ApellidoPat + ' ' + persona.ApellidoMat,
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      confirmButtonText: 'Si Seleccionar'
    }).then((result) => {
      if (result.value) {
        this.dialogRef.close(persona);
      }
    });
  }
}
