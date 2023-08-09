package RoleModel

import (
	"github.com/jinzhu/gorm"
)

func (rma *RoleModuleAction) AddRoleModuleAction(tx *gorm.DB) error {
	return tx.Create(rma).Error
}
