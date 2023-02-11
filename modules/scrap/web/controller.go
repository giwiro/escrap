package web

import (
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
			log.Debug(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		scrapProvider, scrapProviderErr := sc.providerUseCase.FindProviderByUrl(request.Url)

		if scrapProviderErr != nil || scrapProvider == model.Unset {
			if scrapProviderErr != nil {
				log.Error(scrapProviderErr)
			}

			c.JSON(http.StatusNotFound, gin.H{"error": "Could not find scrap provider"})
			return
		}

		result, product, err := sc.scrapUseCase.ScrapUrl(request.Url, scrapProvider)

		if err != nil {
			log.Error(err.Error())
			c.JSON(http.StatusNotFound, gin.H{"error": "Could not scrap url"})
			return
		}

		c.Header("X-Escrap-Result-Id", strconv.Itoa(int(result.Id)))

		c.JSON(200, map[string]interface{}{
			"product": product,
			"result":  result,
		})
	})
}
