package RoleModel

import (
	"github.com/jinzhu/gorm"
)

func (r *Role) AddRole(tx *gorm.DB) error {
	return tx.Create(r).Error
}
