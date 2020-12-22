package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/konveyor/forklift-inventory/models/vmware"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect to database.")
	}

	database.AutoMigrate(&vmware.VCenter{})
	database.AutoMigrate(&vmware.Folder{})
	database.AutoMigrate(&vmware.Datacenter{})

	DB = database
}
