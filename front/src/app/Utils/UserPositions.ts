export class UserPositions {

  static returnDepartment(row: any): any {
    if (row.PuestoOrganigrama.organigramaUnidadId) {
      return {
        department: row.PuestoOrganigrama.organigramaUnidad.departamento.Nombre,
        delegation: row.PuestoOrganigrama.organigramaUnidad.unidadMed.Delegacion.ClvDele + '  ' + row.PuestoOrganigrama.organigramaUnidad.unidadMed.Delegacion.NombreDele,
        unidad: row.PuestoOrganigrama.organigramaUnidad.unidadMed.DenominacionUni
      };
    } else if (row.PuestoOrganigrama.organigramaDelegacionId) {
      return {
        department: row.PuestoOrganigrama.organigramaDelegacion.departamento.Nombre,
        delegation: row.PuestoOrganigrama.organigramaDelegacion.delegaciones.ClvDele + '  ' + row.PuestoOrganigrama.organigramaDelegacion.delegaciones.NombreDele,
        unidad: ''
      };
    }
  }
}
