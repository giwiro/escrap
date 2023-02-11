package web

import (
	"github.com/gin-gonic/gin"
	"github.com/giwiro/escrap/common"
	"github.com/giwiro/escrap/modules/provider"
	"github.com/giwiro/escrap/modules/provider/model"
	"github.com/giwiro/escrap/modules/search"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type SearchController interface {
	common.Controller
}
type searchController struct {
	searchUseCase   search.UseCase
	providerUseCase provider.UseCase
}

func NewSearchController(searchUseCase search.UseCase, providerUseCase provider.UseCase) SearchController {
	return &searchController{searchUseCase, providerUseCase}
}

func (s searchController) RegisterRoutes(group *gin.RouterGroup) {
	searchRouterGroup := group.Group("/search")

	searchRouterGroup.POST("", func(c *gin.Context) {
		var request SearchRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			log.Debug(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		scrapProvider, scrapProviderErr := s.providerUseCase.FindProviderByName(request.Provider)

		if scrapProviderErr != nil || scrapProvider == model.Unset {
			if scrapProviderErr != nil {
				log.Error(scrapProviderErr)
			}

			c.JSON(http.StatusNotFound, gin.H{"error": "Could not find scrap provider"})
			return
		}

		response, err := s.searchUseCase.Search(request.Keyword, request.Page, scrapProvider)

		if err != nil {
			log.Error(err.Error())
			c.JSON(http.StatusNotFound, gin.H{"error": "Could not search"})
			return
		}

		c.JSON(200, response)
	})
}
