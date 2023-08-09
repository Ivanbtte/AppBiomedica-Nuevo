package ContratoModelo

import (
	"appbiomedica/Conexion"
	"appbiomedica/Modulos/Query"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

//********************************************************************************************************
/* *********************    Funcion para agregar un contrato       ********************* */
func (p *Contrato) AgregarContrato(tx *gorm.DB) error {
	return tx.Create(p).Error
}

//********************************************************************************************************
/* *********************    Funcion para verificar que no exita el numero de contrato    ********************* */
func CompruebaContrato(n_contrato string) bool {
	var contrato Contrato
	db := Conexion.GetDB()
	err := db.Where("trim(lower(numero_contrato)) = ?", n_contrato).First(&contrato).Error
	// if err != nil {
	// 	return false
	// }
	// repetido error
	//return true
	return err == nil
}

//********************************************************************************************************
/* *********************    Funcion para regresar un contrato contrato    ********************* */
func RegresaContrato(n_contrato string) *Contrato {
	var contrato Contrato
	db := Conexion.GetDB()
	err := db.Where("numero_contrato = ?", n_contrato).Preload("Proveedor").First(&contrato).Error
	if err != nil {
		return nil
	}
	return &contrato
}

/* *********************    Funcion para regresar un contrato de traslados de pacientes    ********************* */
func RegresaContratoTraslado(unidad_med string) (*[]Contrato, error) {
	var contrato []Contrato
	db := Conexion.GetDB()
	err := db.Select("distinct on (numero_contrato)*").
		Joins("inner join conceptos c on contrato.numero_contrato = c.contrato_numero_contrato").
		Joins("inner join traslado_pacientes tp on c.id = tp.conceptos_id").
		Where("unidad_med_id = ?", unidad_med).
		Find(&contrato).Error
	if err != nil {
		return nil, err
	}
	return &contrato, nil
}

//********************************************************************************************************
/* *********************    Funcion para actualizar un contrato    ********************* */
func (p *Contrato) ActualizarContrato(id_contrato string, tx *gorm.DB) error {
	var contrato Contrato
	err := tx.Where("id = ?", id_contrato).First(&contrato).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro el id")
		}
		return tx.Rollback().Error
	}
	return tx.Model(&p).Where("id=?", id_contrato).Update(Contrato{NumeroContrato: p.NumeroContrato,
		InicioContrato: p.InicioContrato, FinContrato: p.FinContrato, ProcedContratacion: p.ProcedContratacion,
		TipoProcedContratacion: p.TipoProcedContratacion, FechaFallo: p.FechaFallo, MontoTotal: p.MontoTotal,
		TipoContrato: p.TipoContrato, ProveedorNProvImss: p.ProveedorNProvImss}).Error
}

//********************************************************************************************************
/* *******************    Funcion para regresar la lista de contratos o un solo contrato     ******************** */
func ConsultarContrato(id_contrato string, tipo, filtro, proveedores, delegaciones []string, numPagina,
	numRegistros int) (error, *[]Contrato, int) {
	var resultado []Contrato
	var total int
	db := Conexion.GetDB()
	if id_contrato != "" {
		db = db.Where("numero_contrato = ?", id_contrato)
	}
	if len(tipo) != 0 {
		db = db.Where("tipo_contrato_id IN (?)", tipo)
	}
	if len(proveedores) != 0 {
		db = db.Where("proveedor_n_prov_imss IN (?)", proveedores).Group("id")
	}
	if len(delegaciones) != 0 {
		db = db.Joins("INNER JOIN contrato_delegaciones cd on contrato.numero_contrato = cd.contrato_numero_contrato").
			Where("cd.delegaciones_clv_dele IN (?)", delegaciones).Group("id")
	}
	if len(filtro) != 0 {
		db = Query.ContruirQueryContrato(db, filtro)
	}
	err := db.Find(&[]Contrato{}).Count(&total).Error
	if err != nil {
		return err, nil, total
	}
	err = db.Offset((numPagina - 1) * numRegistros).Limit(numRegistros).Preload("Proveedor").Preload("TipoContrato").
		Preload("Proveedor.Estados").Order("(to_date(fin_contrato,'DD/MM/YYYY')) desc").Find(&resultado).Error
	if err != nil {
		return err, nil, total
	}
	fmt.Println("Total de datos consultados: ", total)
	return nil, &resultado, total
}

//********************************************************************************************************
/* *********************    Funcion para eliminar un proveedor     ********************* */
func EliminarContrato(id string, tx *gorm.DB) (error, *Contrato) {
	var item Contrato
	err := tx.Where("id = ?", id).Find(&item).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("No se encontro el id"), nil
		}
		return err, nil
	}
	return tx.Model(&Contrato{}).Where("id=?", id).Delete(&item).Error, &item
}

//********************************************************************************************************
/* *********************   Funcion que agrega un arreglo de delgaciones por contrato     ********************* */
func (admin *Contrato_Delegaciones) AgregarDelegacionContrato(idContrato string, delegaciones []string, tx *gorm.DB) error {
	tabla1 := Contrato_Delegaciones{}
	tabla1.ContratoNumeroContrato = idContrato
	for _, dele := range delegaciones {
		tabla1.DelegacionesClvDele = dele
		err := tx.Create(tabla1).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Error
}

//********************************************************************************************************
/* *********************    Funcion para regresar una delegacion de un contrato    ********************* */
func RegresaDelegContrato(n_contrato, clvDele string) *Contrato_Delegaciones {
	var dele Contrato_Delegaciones
	db := Conexion.GetDB()
	err := db.Where("contrato_numero_contrato = ? AND delegaciones_clv_dele = ?", n_contrato, clvDele).First(&dele).Error
	if err != nil {
		return nil
	}
	return &dele
}

//********************************************************************************************************
/* *********************    Funcion que regresa proveedores de los contratos     ********************* */
func RegresaProveedoresContratos() (*[]Contrato, error) {
	var provee []Contrato
	db := Conexion.GetDB()
	err := db.Select("distinct proveedor_n_prov_imss").Preload("Proveedor").Find(&provee).Error
	if err != nil {
		return nil, err
	}
	return &provee, nil
}
