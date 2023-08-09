export interface ProveedoresInterface {
  Id: string;
  AliasEmpresa: string;
  NombreEmpresa: string;
  Correo: string;
  RFC: string;
  NProvImss: string;
  Telefono: string;
  Direccion: string;
  EstadosId: string;
  Estados: any;
  Municipio: string;
  CP: string;
  Estado: boolean;
}

export interface ContactoProveedor {
  Id: string;
  NombreCompleto: string;
  Departamento: string;
  Cargo: string;
  Telefono: string;
  Extension: string;
  Celular: string;
  Correo: string;
  Comentarios: string;
  Asunto: string;
  ProveedorNProvImss: string;
  Proveedor: any;
  Estado: boolean;
}
