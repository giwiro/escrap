package web

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/giwiro/escrap/common"
	"net/http"
)

type ScrapController interface {
	common.Controller
}
type scrapController struct {
}

func NewScrapController() ScrapController {
	return &scrapController{}
}

func (scrapController scrapController) RegisterRoutes(group *gin.RouterGroup) {
	scrapRouterGroup := group.Group("/scrap")

	scrapRouterGroup.POST("", func(c *gin.Context) {
		var request ScrapRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		spew.Dump(request)

		c.String(200, "NO trailing")
	})
}
