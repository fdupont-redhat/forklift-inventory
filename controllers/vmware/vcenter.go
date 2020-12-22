package vmware

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/konveyor/forklift-inventory/models"
	"github.com/konveyor/forklift-inventory/models/vmware"
)

// FindVCenters godoc
// @Summary Find all VMware vCenters
// @Description get array of hashes
func FindVCenters(c *gin.Context) {
	var vcenters []vmware.VCenter
	models.DB.Find(&vcenters)

	c.JSON(http.StatusOK, gin.H{"data": vcenters})
}
