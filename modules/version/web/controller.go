package web

import (
	"github.com/gin-gonic/gin"
	"github.com/giwiro/escrap/common"
	"github.com/giwiro/escrap/modules/version"
	"net/http"
)

type VersionController interface {
	common.Controller
}
type versionController struct {
}

func NewVersionController() VersionController {
	return &versionController{}
}

func (vc *versionController) RegisterRoutes(group *gin.RouterGroup) {
	group.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"version": version.Version})
	})
}
