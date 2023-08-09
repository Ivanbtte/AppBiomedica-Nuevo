package BitacoraModel

import (
	"github.com/jinzhu/gorm"
)

//********************************************************************************************************
/* *********************    Funcion para agregar un registro a la bitacora     ********************* */
func (p *BitacoraClave) AgregaBitacoraClave(tx *gorm.DB) error {
	return tx.Create(p).Error
}
