package TrasladoPacientesModel

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modelos/ContratoModel/ConceptosServiciosModelo"
	"appbiomedica/Modelos/ModeloUnidadM"
	"appbiomedica/Modulos/Estructuras"
	"appbiomedica/Modulos/Query"
	"database/sql"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

//********************************************************************************************************
/* *********************   Funcion para agregar un concepto de un contrato     ********************* */
func (st *SolicitudTraslado) AgregarSolicitud(tx *gorm.DB) error {
	fmt.Println("Finalllll", st)
	return tx.Create(st).Error
}

//********************************************************************************************************
/* *********************    Funcion que busca que no se repita un folio ftp     ********************* */
func BuscarFolioFtp(folio, estado string) (bool, *SolicitudTraslado) {
	var solicitud SolicitudTraslado
	db := Conexion.GetDB()
	err := db.Where("ftp01 = ? AND estado =?", folio, estado).Preload("UnidadMed").First(&solicitud).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		return false, nil
	}
	// repetido error
	return true, &solicitud
}

/* *********************    Funcion que busca una solicitud de traslado por numero de folio     ********************* */
func BuscarFolioSolicitud(folio string) (bool, *SolicitudTraslado) {
	var solicitud SolicitudTraslado
	db := Conexion.GetDB()
	err := db.Where("folio = ?", folio).Preload("UnidadMed").Preload("Usuario").Preload("Usuario.Persona").First(&solicitud).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		return false, nil
	}
	return true, &solicitud
}

/* *********************    Funcion que busca una solicitud de traslado por numero de folio     ********************* */
func BuscarFolioSolicitudes(folio string) (bool, *[]SolicitudTraslado) {
	var solicitud []SolicitudTraslado
	db := Conexion.GetDB()
	err := db.Where("ftp01 = ? and estado =? ", folio, true).Find(&solicitud).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		return false, nil
	}
	return true, &solicitud
}

//********************************************************************************************************
/* *********************    Funcion que busca el concepto de traslado     ********************* */
func SearchTicket(unitId int, origin, destination, ticketType, contractNumber string) (bool,
	*ConceptosServiciosModelo.TrasladoPacientes) {
	var tickets ConceptosServiciosModelo.TrasladoPacientes
	db := Conexion.GetDB()
	err := db.Joins("inner join conceptos c on c.id = traslado_pacientes.conceptos_id").
		Where("unidad_med_id = ? "+
			"AND (trim(lower(origen))) = ? "+
			"AND (trim(lower(destino))) = ? "+
			"AND c.contrato_numero_contrato = ?"+
			"AND trim(lower("+"unidad_medida)) LIKE ?", unitId, Estructuras.ConvertLowercase(origin),
			Estructuras.ConvertLowercase(destination), contractNumber, "%"+strings.ToLower(ticketType)+"%").First(&tickets).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		return false, nil
	}
	return true, &tickets
}

//********************************************************************************************************
/* ***************    Funcion que busca el numero consecutivo que sigue por unidad y por año     ******************* */
func BuscarFolioConsecutivo(unidad, anio int) (error, int) {
	var total int
	db := Conexion.GetDB()
	err := db.Model(&SolicitudTraslado{}).Where(" date_part('year', "+
		"fecha_cita) = ? AND unidad_med_id = ?", anio, unidad).Count(&total).Error
	if err != nil {
		return err, 0
	}
	fmt.Println("Total de datos consultados", total)
	return nil, total
}

//********************************************************************************************************
/* ***************    Funcion que busca por unidad medica todas las solicitudes realizadas     ******************* */
func ConsultarSolicitudes(unidad, numPagina, numRegistro int, filtro, estado, companion, origin,
	destination string, contractsNumber []string, startDate, endDate string) (error, *[]SolicitudTraslado,
	int) {
	var datos []SolicitudTraslado
	var total int
	db := Conexion.GetDB()
	db = db.Where("estado =?", estado)
	db = db.Where("unidad_med_id =?", unidad)
	if len(filtro) != 0 {
		db = Query.ContruirQuerySolicitudTraslado(db, filtro)
	}
	if startDate != "" && endDate != "" {
		db = db.Where("fecha_expedicion BETWEEN ? AND ?", startDate, endDate)
	}
	if companion != "" {
		if companion == "1" {
			db = db.Where("acompañante =?", true)
		} else {
			db = db.Where("acompañante =?", false)
		}
	}
	if origin != "" {
		db = db.Where("origen =?", origin)
	}
	if destination != "" {
		db = db.Where("destino =?", destination)
	}
	if len(contractsNumber) != 0 {
		db = db.Where("contrato_numero_contrato IN (?)", contractsNumber)
	}
	err := db.Find(&[]SolicitudTraslado{}).Count(&total).Error
	if err != nil {
		return err, nil, total
	}
	err = db.Offset((numPagina - 1) * numRegistro).Limit(numRegistro).Order("fecha_expedicion desc").Find(
		&datos).Error
	if err != nil {
		return err, nil, total
	}
	return nil, &datos, total
}

/*****  Funcion para cancelar una solicitud de traslado  *****/
func (st SolicitudTraslado) CancelarSolicitud(id string, tx *gorm.DB) error {
	return tx.Model(&st).Where("folio=?", id).Where("estado = ?", true).Update("estado", false).Error
}

/*****  Funcion para cancelar una solicitud de traslado con folios separados  *****/
func (st SolicitudTraslado) CancelarSolicitudSeparados(ftp string, tx *gorm.DB) error {
	return tx.Model(&st).Where("ftp01=?", ftp).Where("estado = ?", true).Update("estado", false).Error
}

/* *********************   Funcion para agregar un registro a la bitacora de solicitudes canceladas
********************* */
func (sc *SolicitudesCanceladas) AgregarBitacoraCancelacion(tx *gorm.DB) error {
	return tx.Create(sc).Error
}

func SearchOriginFilter(unitId int, state bool, destination, companion string, contractsNumber []string, startDate,
	endDate string) (*[]SolicitudTraslado, error) {
	var origins []SolicitudTraslado
	db := Conexion.GetDB()
	if destination != "" {
		db = db.Where("destino =?", destination)
	}
	if companion != "" {
		if companion == "1" {
			db = db.Where("acompañante =?", true)
		} else {
			db = db.Where("acompañante =?", false)
		}
	}
	if len(contractsNumber) != 0 {
		db = db.Where("contrato_numero_contrato IN (?)", contractsNumber)
	}
	if startDate != "" && endDate != "" {
		db = db.Where("fecha_expedicion BETWEEN ? AND ?", startDate, endDate)
	}
	err := db.Select("distinct origen").Where("unidad_med_id = ?  AND estado = ?", unitId, state).Find(&origins).Error
	if err != nil {
		return nil, err
	}
	return &origins, nil
}
func SearchDestinationFilter(unitId int, state bool, origin, companion string, contractsNumber []string, startDate,
	endDate string) (*[]SolicitudTraslado, error) {
	var origins []SolicitudTraslado
	db := Conexion.GetDB()
	if origin != "" {
		db = db.Where("origen =?", origin)
	}
	if companion != "" {
		if companion == "1" {
			db = db.Where("acompañante =?", true)
		} else {
			db = db.Where("acompañante =?", false)
		}
	}
	if len(contractsNumber) != 0 {
		db = db.Where("contrato_numero_contrato IN (?)", contractsNumber)
	}
	if startDate != "" && endDate != "" {
		db = db.Where("fecha_expedicion BETWEEN ? AND ?", startDate, endDate)
	}
	err := db.Select("distinct destino").Where("unidad_med_id=? AND estado = ?", unitId, state).Find(&origins).Error
	if err != nil {
		return nil, err
	}
	return &origins, nil
}

type Results struct {
	Acompañante sql.NullBool `json:"acompañante,omitempty"`
}

func SearchCompanionFilter(unitId int, state bool, origin, destination string, contractsNumber []string, startDate,
	endDate string) ([]Results, error) {
	var result []Results
	db := Conexion.GetDB()
	if origin != "" {
		db = db.Where("origen =?", origin)
	}
	if destination != "" {
		db = db.Where("destino =?", destination)
	}
	if len(contractsNumber) != 0 {
		db = db.Where("contrato_numero_contrato IN (?)", contractsNumber)
	}
	if startDate != "" && endDate != "" {
		db = db.Where("fecha_expedicion BETWEEN ? AND ?", startDate, endDate)
	}
	err := db.Table("solicitud_traslado").Select("distinct acompañante").Where("unidad_med_id=? AND estado = ?",
		unitId, state).Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Funcion para consultar el total de solicitudes realizadas por contrato
func ContarSolicitudesRealizadas(numeroContrato string, estado bool) (int64, error) {
	var total int64
	db := Conexion.GetDB()
	err := db.Model(&SolicitudTraslado{}).Where("contrato_numero_contrato = ? and estado = ?", numeroContrato, estado).Count(&total).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return 0, nil
		}
		return 0, err
	}
	return total, nil
}

// Funcion para contar las solicitudes con o sin acompaniante
func ContarSolicitudesAcompaniante(numeroContrato string, acompaniante, estadoSolicitud bool) (int32, error) {
	var data int32
	db := Conexion.GetDB()
	err := db.Model(&SolicitudTraslado{}).Where("contrato_numero_contrato = ? and acompañante = ? and estado =?", numeroContrato, acompaniante, estadoSolicitud).Count(&data).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return 0, nil
		}
		return 0, err
	}
	return data, nil
}

type Promedio struct {
	UnidadMedId  int
	Total        int64
	Acompaniante int64
	Promedio     float64
}

// Funcion que regresa la unidad con el mayor porcentaje de solicitudes con acompañante
func PorcentajeMayorAcompaniante(numContrato string) ([]Promedio, error) {
	var datos []Promedio
	db := Conexion.GetDB()
	rows, err := db.Raw(`select a.unidad_med_id, a.total, b.acompaniante, b.acompaniante::decimal / a.total::decimal as promedio from(select unidad_med_id ,count(*) as total from solicitud_traslado where contrato_numero_contrato =? group by (unidad_med_id)) a inner join (select unidad_med_id ,count(*) as acompaniante from solicitud_traslado where contrato_numero_contrato =? AND acompañante =? group by(unidad_med_id))b on a.unidad_med_id = b.unidad_med_id order by promedio desc ;`, numContrato, numContrato, true).Rows()
	if err != nil {
		return datos, err
	}
	defer rows.Close()
	for rows.Next() {
		var e Promedio
		db.ScanRows(rows, &e)
		datos = append(datos, e)
	}
	return datos, nil
}

// Funcion para consultar que unidad medica tiene el mayor numero de solicitudes realizadas por contrato
type UnidadMayorSolicitud struct {
	Unidad int `json:"unidad"`
	Total  int `json:"total"`
}

func UnidadMayorSolicitudes(numContrato string, estado bool) (UnidadMayorSolicitud, error) {
	var datos UnidadMayorSolicitud
	db := Conexion.GetDB()
	err := db.Table("solicitud_traslado").Select("unidad_med_id as unidad ,count(*) as total").Where("contrato_numero_contrato = ? and estado = ?", numContrato, estado).Group("unidad_med_id").Order("total desc").Limit(1).Scan(&datos).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return datos, nil
		}
		return datos, err
	}
	return datos, nil
}

// Funcion para poder consultar cuantas solicitudes ha echo cada unidad medica con acompañante o sin acompañante ordenado de mayor a menor

type AcompañantePorUnidad struct {
	Numero       int                     `json:"numero"`
	Unidad       ModeloUnidadM.UnidadMed `json:"unidad"`
	Total        int                     `json:"total"`
	Acompaniante int                     `json:"acompaniante"`
	Promedio     float32                 `json:"promedio"`
}

func ContarSolicitudesAcompaniantePorUnidad(numeroContrato string, acompaniante, estadoSolicitud bool) ([]AcompañantePorUnidad, error) {
	var data []AcompañantePorUnidad
	db := Conexion.GetDB()
	rows, err := db.Raw(`select a.numero, a.total, b.acompaniante, b.acompaniante::decimal / a.total::decimal as promedio from (select unidad_med_id as numero, count(*) as total from solicitud_traslado st where (contrato_numero_contrato = ? and estado = ?) group by numero) a
    inner join (select unidad_med_id as numero, count(*) as acompaniante from solicitud_traslado st2 where(contrato_numero_contrato = ? and acompañante = ? and estado = ?) group by numero) b on a.numero = b.numero order by total desc;`, numeroContrato, estadoSolicitud, numeroContrato, acompaniante, estadoSolicitud).Rows()
	if err != nil {
		return data, err
	}
	defer rows.Close()
	for rows.Next() {
		var e AcompañantePorUnidad
		db.ScanRows(rows, &e)
		data = append(data, e)
	}
	return data, nil
}
