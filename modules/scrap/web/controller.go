package web

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/giwiro/escrap/common"
	"github.com/giwiro/escrap/modules/provider"
	"github.com/giwiro/escrap/modules/provider/model"
	"github.com/giwiro/escrap/modules/scrap"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
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

		if scrapProviderErr != nil || scrapProvider == model.Unset {
			if scrapProviderErr != nil {
				log.Info(scrapProviderErr.Error())
			}

			c.JSON(http.StatusNotFound, gin.H{"error": "Could not find scrap provider"})
			return
		}

		result, _, err := sc.scrapUseCase.ScrapUrl(request.Url, scrapProvider)

		if err != nil {
			log.Error(err.Error())
			c.JSON(http.StatusNotFound, gin.H{"error": "Could scrap url"})
			return
		}

		/*spew.Dump(result)
		spew.Dump(product)*/

		c.Header("X-Escrap-Result-Id", strconv.Itoa(int(result.Id)))

		c.JSON(200, result)
	})
}
