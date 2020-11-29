package vmware

import (
	"github.com/gin-gonic/gin"
	"github.com/konveyor/forklift-inventory/models"
)

func FindVCenters(c *gin.Context) {
	var vcenters []models.vmware.VCenter
	models.DB.Find(&vcenters)

	c.JSON(http.StatusOK, gin.H{"data": vcenters})
}
