package vmware

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/konveyor/forklift-inventory/models"
	"github.com/konveyor/forklift-inventory/models/vmware"
)

func FindFolders(c *gin.Context) {
	var folders []vmware.Folder
	models.DB.Find(&folders)

	c.JSON(http.StatusOK, gin.H{"data": folders})
}
