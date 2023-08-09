package DepartamentosModel
import "github.com/jinzhu/gorm"

func (od *OrganigramaDelegacion) AddOrganigramaDelegation(tx *gorm.DB) error {
	return tx.Create(od).Error
}