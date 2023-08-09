package TrasladoPacientesModel

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modulos/Estructuras"
	"github.com/jinzhu/gorm"
)

/* *********************   Funcion para agregar una ruta consumida de un contrato     ********************* */
func (rc *RutasConsumidas) AgregarRutaConsumida(tx *gorm.DB) error {
	return tx.Create(rc).Error
}

func ConsultEarn(state bool, unitId int, contractNumber string) (error, float64) {
	type Result struct {
		Total float64
	}
	var results Result
	db := Conexion.GetDB()
	if unitId != 0 {
		db = db.Where("st.unidad_med_id = ?", unitId)
	}
	err := db.Table("rutas_consumidas").Select("sum(cantidad * tp.precio_ofertado) as total").
		Joins("inner join traslado_pacientes tp on rutas_consumidas.traslado_pacientes_id = tp.id").
		Joins(" inner join boletos b on b.id = rutas_consumidas.boletos_id").
		Joins("inner join solicitud_traslado st on b.solicitud_traslado_folio = st.folio").
		Where("st.estado = ?", state).Where("st.contrato_numero_contrato = ?", contractNumber).Scan(&results).Error
	if err != nil {
		return err, 0
	}
	return nil, results.Total
}
func ConsultTicketGoingAvailable(destination, typeTicket, contractNumber string, state bool, unitId int) (error,
	int64) {
	type Result struct {
		Total int64
	}
	var results Result
	db := Conexion.GetDB()
	err := db.Table("rutas_consumidas").Select("sum(cantidad) as total").
		Joins("inner join traslado_pacientes tp on rutas_consumidas.traslado_pacientes_id = tp.id").
		Joins(" inner join boletos b on b.id = rutas_consumidas.boletos_id").
		Joins("inner join solicitud_traslado st on b.solicitud_traslado_folio = st.folio").
		Where("st.estado = ? AND st.unidad_med_id = ? AND st.contrato_numero_contrato =? ", state, unitId, contractNumber).
		Where("trim(lower(tp.destino)) like ?", "%"+Estructuras.ConvertLowercase(destination)+"%").
		Where("trim(lower(tp.unidad_medida)) like ?", "%"+Estructuras.ConvertLowercase(typeTicket)+"%").Scan(&results).
		Error
	if err != nil {
		return err, 0
	}
	return nil, results.Total
}

func BuscarBoletosTrasladoPrecio(folio string) *[]RutasConsumidas {
	var boleto []RutasConsumidas
	db := Conexion.GetDB()
	err := db.Joins("inner join boletos b on rutas_consumidas.boletos_id = b.id").
		Joins("inner join traslado_pacientes tp on rutas_consumidas.traslado_pacientes_id = tp.id ").
		Where("b.solicitud_traslado_folio = ?", folio).Preload("Boletos").Preload("TrasladoPacientes").Find(
		&boleto).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil
		}
		return nil
	}
	return &boleto
}

type DevengoPorMes struct {
	Mes     int     `json:"mes"`
	Devengo float32 `json:"devengo"`
}

func ConsultarDevengoPorMes(state bool, unitId int, contractNumber string) ([]DevengoPorMes, error) {
	var results []DevengoPorMes
	db := Conexion.GetDB()
	if unitId != 0 {
		db = db.Where("st.unidad_med_id = ?", unitId)
	}
	err := db.Table("rutas_consumidas").Select("to_char(st.fecha_expedicion, 'MM') as mes,(sum(cantidad * tp.precio_ofertado))*1.16 as devengo").
		Joins("inner join traslado_pacientes tp on rutas_consumidas.traslado_pacientes_id = tp.id").
		Joins(" inner join boletos b on b.id = rutas_consumidas.boletos_id").
		Joins("inner join solicitud_traslado st on b.solicitud_traslado_folio = st.folio").
		Where("st.estado = ?", state).Where("st.contrato_numero_contrato = ?", contractNumber).Group("to_char(st.fecha_expedicion, 'MM')").Scan(&results).Error
	if err != nil {
		return results, err
	}
	return results, nil
}

type DevengoPorUnidad struct {
	Unidad  int     `json:"unidad"`
	Devengo float32 `json:"devengo"`
}

func ConsultarDevengoMesUnidad(state bool, mes, numContrato string) ([]DevengoPorUnidad, error) {
	var results []DevengoPorUnidad
	db := Conexion.GetDB()
	err := db.Table("rutas_consumidas").Select("	um.denominacion_uni as nombre, (sum(cantidad * tp.precio_ofertado))* 1.16 as devengo").
		Joins("inner join traslado_pacientes tp on rutas_consumidas.traslado_pacientes_id = tp.id").
		Joins("inner join boletos b on b.id = rutas_consumidas.boletos_id").
		Joins("inner join solicitud_traslado st on b.solicitud_traslado_folio = st.folio").
		Joins("inner join unidad_med um on st.unidad_med_id = um .id").
		Where("st.estado =? and st.contrato_numero_contrato =? and to_char(st.fecha_expedicion, 'MM') =?", state, numContrato, mes).
		Group("nombre").Scan(&results).Error
	if err != nil {
		return results, err
	}
	return results, nil
}
