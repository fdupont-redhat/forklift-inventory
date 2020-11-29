package vmware

import (
	"github.com/jinzhu/gorm"
)

type Datacenter struct {
	gorm.Model
	ID		uint16	`json:"id" gorm:"primaryKey"`
	Name		string	`json:"name" gorm:"not null"`
	FolderID	uint16
}
