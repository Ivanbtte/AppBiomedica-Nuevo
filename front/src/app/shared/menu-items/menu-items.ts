import {Injectable} from '@angular/core';


export interface BadgeItem {
  type: string;
  value: string;
}

export interface Separator {
  name: string;
  type?: string;
}

export interface SubChildren {
  state: string;
  name: string;
  type?: string;
  permit?: string;
}

export interface ChildrenItems {
  state: string;
  name: string;
  type?: string;
  child?: SubChildren[];
}

export interface Menu {
  state: string;
  id:string;
  moduloId:string;
  orden:number;
  name: string;
  type: string;
  icon: string;
  status:boolean;
  permit?: string;
  badge?: BadgeItem[];
  separator?: Separator[];
  children?: ChildrenItems[];
}

// const MENUITEMS = [
//   {
//     state: 'inicio',
//     name: 'Inicio',
//     type: 'link',
//     icon: 'apps',
//     permiso: 'Consulta',
//   },
//   {
//     state: 'catalogos',
//     name: 'Catálogos',
//     type: 'sub',
//     icon: 'assignment',
//     permiso: 'Consulta',
//     children: [
//       {state: 'sai', name: 'SAI', type: 'link'},
//       {state: 'prei', name: 'PREI', type: 'link'},
//       {state: 'proforma', name: 'PROFORMA', type: 'link'},
//       /*{state: 'manuales', name: 'MANUALES', type: 'link'},
//       {state: 'otros', name: 'OTROS CATALOGOS', type: 'link'}*/
//
//     ]
//   },
//   {
//     state: 'listas',
//     name: 'Listas de Interés',
//     type: 'link',
//     icon: 'link',
//     permiso: 'Consulta'
//   },
//   {
//     state: 'area',
//     name: 'Areas Servicios',
//     type: 'sub',
//     icon: 'business',
//     permiso: 'Enfermeras',
//     children: [
//       {state: 'nueva', name: 'Agregar', type: 'link'},
//     ]
//   },
//   {
//     state: 'planeacion',
//     name: 'Fondo Fijo',
//     type: 'sub',
//     icon: 'today',
//     permiso: 'Enfermeras',
//
//     children: [
//       {
//         state: 'enfermeria', name: 'Enfermeria', type: 'subchild',
//         subchildren: [
//           {
//             state: 'cuadrosi',
//             name: 'En Cuadro Basico',
//             type: 'link'
//           },
//           {
//             state: 'cuadrono',
//             name: 'Fuera de Cuadro Basico',
//             type: 'link'
//           },
//         ]
//       },
//     ]
//   },
//   {
//     state: 'solicitudes',
//     name: 'Solicitudes',
//     type: 'sub',
//     icon: 'description',
//     permiso: 'Solicitud',
//
//     children: [
//       {
//         state: 'nueva',
//         name: 'Nueva',
//         type: 'subchild',
//         subchildren: [
//           {
//             state: 'consumo',
//             name: 'Bienes de Consumo',
//             type: 'link'
//           },
//           /*          {
//                       state: 'cocina',
//                       name: 'Equipo de Cocina',
//                       type: 'link'
//                     },*/
//           /*          {
//                       state: 'medico',
//                       name: 'Equipo Medico',
//                       type: 'link'
//                     },*/
//           /*          {
//                       state: 'instrumental',
//                       name: 'Instrumental',
//                       type: 'link'
//                     },*/
//         ]
//       },
//       {state: 'seguimiento', name: 'Seguimiento', type: 'link'},
//       {state: 'editar', name: 'Editar', type: 'link'}
//     ]
//   },
//   {
//     state: 'configuraciones',
//     name: 'Configuraciones',
//     type: 'sub',
//     icon: 'build',
//     permiso: 'Jefe Biomedico',
//
//     children: [
//       {state: 'fecha', name: 'Fecha Límite', type: 'link'},
//       {state: 'presupuesto', name: 'Presupuesto', type: 'link'}
//
//     ]
//   },
//   {
//     state: 'usuarios',
//     name: 'Usuarios',
//     type: 'sub',
//     icon: 'supervised_user_circle',
//     permiso: 'Jefe Biomedico',
//
//     children: [
//       {state: 'validar', name: 'Validar Usuarios', type: 'link'},
//       {state: 'bloquear', name: 'Bloquear', type: 'link'}
//
//     ]
//   },
//   {
//     state: 'requerimientos',
//     name: 'Validar Solicitudes',
//     type: 'sub',
//     icon: 'assignment_turned_in',
//     permiso: 'Jefe Biomedico',
//
//     children: [
//       {state: 'revision', name: 'Validar Solicitud', type: 'link'},
//       {state: 'validado', name: 'Solicitudes Aprobadas', type: 'link'},
//       {state: 'final', name: 'Solicitud Final', type: 'link'}
//
//     ]
//   },
//   {
//     state: 'proveedores',
//     name: 'Proveedores',
//     type: 'sub',
//     icon: 'apartment',
//     permiso: 'Consulta',
//
//     children: [
//       {state: 'lista', name: 'Lista de Proveedores', type: 'link'},
//     ]
//   },
//   {
//     state: 'contratos',
//     name: 'Contratos',
//     type: 'sub',
//     icon: 'receipt',
//     permiso: 'Administrador de contratos',
//
//     children: [
//       {state: 'administrar', name: 'Administrar', type: 'link'},
//       {state: 'seguimiento', name: 'Seguimiento', type: 'link'},
//     ]
//   },
//   {
//     state: 'traslado',
//     name: 'Traslados de pacientes',
//     type: 'sub',
//     icon: 'receipt',
//     permiso: 'Solicitud traslado',
//
//     children: [
//       {state: 'solicitud', name: 'Nueva Solicitud', type: 'link'},
//       {state: 'realizadas', name: 'Solicitudes Realizadas', type: 'link'},
//       {state: 'cancelar', name: 'Cancelar Solicitud', type: 'link'},
//     ]
//   },
// ];

@Injectable()
export class MenuItems {
  // public men: any[];
  // public men1: any[];
  // public men2: any[];
  // public men3: any[];
  // public men4: any[];
  // public permisos: any[];
  public MENUITEM = [];

  constructor() {
  }


  getMenuitems(): Menu[] {
    return this.MENUITEM;
  }

  // getMenuitem(): Menu[] {
  //   /* ******  funcion que concatena los menus por default quue tienen todos los usuarios  ******   */
  //   this.men = MENUITEMS.filter(menu => {
  //     return menu.permiso === 'Consulta';
  //   });
  //   for (let i of this.permisos) {
  //     if (i === 'Administrador de contratos') {
  //       this.men2 = MENUITEMS.filter(menu => {
  //         return menu.permiso === 'Administrador de contratos';
  //       });
  //       this.men = this.men.concat(this.men2);
  //     }
  //     if (i === 'Jefe Biomedico') {
  //       this.men1 = MENUITEMS.filter(menu => {
  //         return menu.permiso === i;
  //       });
  //       this.men = this.men.concat(this.men1);
  //     }
  //     if (i === 'Solicitud traslado') {
  //       this.men4 = MENUITEMS.filter(menu => {
  //         return menu.permiso === i;
  //       });
  //       this.men = this.men.concat(this.men4);
  //     }
  //     if (i === 'Solicitud') {
  //       this.men3 = MENUITEMS.filter(menu => {
  //         return menu.permiso === i;
  //       });
  //       this.men = this.men.concat(this.men3);
  //     }
  //   }
  //   /* ******  regreso el menu  ******   */
  //   const tem = new Set(this.men);
  //   this.men = Array.from(tem);
  //   return this.men;
  // }
}
