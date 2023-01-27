package web

import (
	"errors"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/giwiro/escrap/common"
	"github.com/giwiro/escrap/modules/provider"
	"github.com/giwiro/escrap/modules/scrap"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

type ScrapController interface {
	common.Controller
}
type scrapController struct {
	scrapUseCase    scrap.UseCase
	providerUseCase provider.UseCase
}

func NewScrapController(scrapUseCase scrap.UseCase, providerUseCase provider.UseCase) ScrapController {
	return &scrapController{scrapUseCase, providerUseCase}
}

func (sc *scrapController) RegisterRoutes(group *gin.RouterGroup) {
	scrapRouterGroup := group.Group("/scrap")

	scrapRouterGroup.POST("", func(c *gin.Context) {
		var request ScrapRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			spew.Dump(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		scrapProvider, scrapProviderErr := sc.providerUseCase.FindProviderByUrl(request.Url)

		if scrapProviderErr != nil {
			if !errors.Is(scrapProviderErr, gorm.ErrRecordNotFound) {
				log.Info(scrapProviderErr.Error())
			}

			c.JSON(http.StatusNotFound, gin.H{"error": "Could not find scrap provider"})
			return
		}

		result, product, err := sc.scrapUseCase.ScrapUrl(request.Url, scrapProvider)

		if err != nil {
			log.Error(err.Error())
			c.JSON(http.StatusNotFound, gin.H{"error": "Could scrap url"})
			return
		}

		spew.Dump(result)
		spew.Dump(product)

		c.String(200, "Gaaa")
	})
}
