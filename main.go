package main

import (
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/konveyor/forklift-inventory/docs"

	"github.com/konveyor/forklift-inventory/models"
	"github.com/konveyor/forklift-inventory/controllers/vmware"
)

// @title Forklift Inventory API
// @version 1.0
// @description This is the inventory componenat of the Forklift project

// @contact.name Konveyor - Forklift development team
// @contact.url https://forklift.konveyor.io
// @contact.email https://forklift-dev@konveyor.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api/v1
func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/vcenters", vmware.FindVCenters)
	r.GET("/datacenters", vmware.FindDatacenters)
	r.GET("/folders", vmware.FindFolders)

	swaggerURL := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerURL))

	r.Run()
}
