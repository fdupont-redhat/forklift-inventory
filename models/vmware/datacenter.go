package vmware

import (
	"github.com/jinzhu/gorm"
)

type Datacenter struct {
	gorm.Model
	Name		string	`json:"name" gorm:"not null"`
	FolderID	uint16
}
