package web

import (
	"github.com/gin-gonic/gin"
	"github.com/giwiro/escrap/common"
)

type ScrapController interface {
	common.Controller
}
type scrapController struct {
}

func (scrapController scrapController) RegisterRoutes(group *gin.RouterGroup) {
	// scrapRouterGroup := group.Group("/scrap")

	// scrapRouterGroup.GET()
}
