package vmware

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/konveyor/forklift-inventory/models"
	"github.com/konveyor/forklift-inventory/models/vmware"
)

func FindDatacenters(c *gin.Context) {
	var datacenters []vmware.Datacenter
	models.DB.Find(&datacenters)

	c.JSON(http.StatusOK, gin.H{"data": datacenters})
}
