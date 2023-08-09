import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {AbrirCatalogoSaiService} from './servicios/abrir-catalogo-sai.service';
import {DelegacionesService} from './servicios/delegaciones.service';
import {EntidadFederativaService} from './servicios/entidad-federativa.service';
import {EstatusService} from './servicios/estatus.service';
import {NombramientoService} from './servicios/nombramiento.service';
import {ServiciosUnidadesService} from './servicios/servicios-unidades.service';
import {UnidadNombramientoService} from './servicios/unidad-nombramiento.service';
import {UnidadesMedicasService} from './servicios/unidades-medicas.service';
import {UsuarioTipoService} from './servicios/usuario-tipo.service';
import {AreaServicioService} from './servicios/AreasServicios/area-servicio.service';
import {BitacoraSolicitudService} from './servicios/bitacora/bitacora-solicitud.service';
import {AdministradorService} from './servicios/ContratosServicios/administrador.service';
import {ArchivosContratosService} from './servicios/ContratosServicios/archivos-contratos.service';
import {ConceptosContratoService} from './servicios/ContratosServicios/conceptos-contrato.service';
import {ContactosContratoService} from './servicios/ContratosServicios/contactos-contrato.service';
import {ContratoService} from './servicios/ContratosServicios/contrato.service';
import {ContratoDelegacionService} from './servicios/ContratosServicios/contrato-delegacion.service';
import {ContratoPuestoService} from './servicios/ContratosServicios/contrato-puesto.service';
import {DistribucionConceptosService} from './servicios/ContratosServicios/distribucion-conceptos.service';
import {FianzaService} from './servicios/ContratosServicios/fianza.service';
import {PenasService} from './servicios/ContratosServicios/penas.service';
import {RepresentanteLegalService} from './servicios/ContratosServicios/representante-legal.service';
import {SubTipoService} from './servicios/ContratosServicios/sub-tipo.service';
import {SubTipoContratoService} from './servicios/ContratosServicios/sub-tipo-contrato.service';
import {DepartamentosService} from './servicios/DepartamentosSubdireccion/departamentos.service';
import {ContactoProveedoresService} from './servicios/ProveedoresServicios/contacto-proveedores.service';
import {ProveedoresService} from './servicios/ProveedoresServicios/proveedores.service';
import {FtpService} from './servicios/ServiciosSubrogados/traslados/ftp.service';
import {PatientTransferRoutesService} from './servicios/ServiciosSubrogados/traslados/patient-transfer-routes.service';
import {ServiciosUsuariosService} from './servicios/ServiciosUsuarios/servicios-usuarios.service';
import {ClavesService} from './servicios/solicitud/claves.service';
import {SolicitudService} from './servicios/solicitud/solicitud.service';
import {SolicitudTrasladoService} from './servicios/SolicitudTraslados/solicitud-traslado.service';
import {FirmService} from './servicios/ServiciosSubrogados/traslados/firm.service';
import { DisableFormControlDirective } from './directive/disable-form-control.directive';

@NgModule({
    declarations: [
        DisableFormControlDirective
    ],
    imports: [
        CommonModule
    ],
    exports: [
        DisableFormControlDirective
    ],
    providers: [
        AbrirCatalogoSaiService,
        DelegacionesService,
        EntidadFederativaService,
        EstatusService,
        NombramientoService,
        ServiciosUnidadesService,
        UnidadNombramientoService,
        UnidadesMedicasService,
        UsuarioTipoService,
        AreaServicioService,
        BitacoraSolicitudService,
        AdministradorService,
        ArchivosContratosService,
        ConceptosContratoService,
        ContactosContratoService,
        ContratoService,
        ContratoDelegacionService,
        ContratoPuestoService,
        DistribucionConceptosService,
        FianzaService,
        PenasService,
        RepresentanteLegalService,
        SubTipoService,
        SubTipoContratoService,
        DepartamentosService,
        ContactoProveedoresService,
        ProveedoresService,
        FtpService,
        PatientTransferRoutesService,
        ServiciosUsuariosService,
        ClavesService,
        SolicitudService,
        SolicitudTrasladoService,
        FirmService
    ]
})
export class CoreModule {
}
