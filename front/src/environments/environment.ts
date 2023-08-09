// The file contents for the current environment will overwrite these during build.
// The build system defaults to the dev environment which uses `environment.ts`, but if you do
// `ng build --env=prod` then `environment.prod.ts` will be used instead.
// The list of which env maps to which file can be found in `.angular-cli.json`.
// ********************************************************************************************************
/* *****  funcion para modo produccion  *****   */
// export const environment = {
//   production: false,
//   host: 'http://172.25.251.100:80'
// };
// ********************************************************************************************************
/* *****  funcion para modo pruebas  *****   */
export const environment = {
  production: true,
  // host: 'http://192.168.1.121:8080'
   host: 'http://192.168.1.134:8080'
};

