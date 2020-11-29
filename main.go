package main

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/konveyor/forklift-inventory/models"
	"github.com/konveyor/forklift-inventory/controllers"
)

func main() {
	r := gin.Default()

	db := models.ConnectDatabase()

	r.GET("/vcenters", controllers.vmware.FindVCenters)

	r.Run()
}
