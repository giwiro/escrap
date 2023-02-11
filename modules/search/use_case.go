package search

import (
	"errors"
	"github.com/giwiro/escrap/modules/provider/dao"
	"github.com/giwiro/escrap/modules/provider/model"
	"github.com/giwiro/escrap/modules/provider/vendors"
)

type UseCase interface {
	Search(keyword string, page uint, provider model.ScrapProvider) (map[string]interface{}, error)
}
type useCase struct {
	scrapProviderDao dao.ScrapProviderDao
}

func NewUseCase(scrapProviderDao dao.ScrapProviderDao) UseCase {
	return &useCase{scrapProviderDao}
}

func (u useCase) Search(keyword string, page uint, provider model.ScrapProvider) (response map[string]interface{}, err error) {
	response = nil
	providerVendor, err := vendors.NewScrapProviderVendor(provider, &u.scrapProviderDao)

	if err != nil {
		return
	}

	if providerVendor == nil {
		err = errors.New("no provider vendor was found")
		return
	}

	response, err = providerVendor.Search(keyword, page)

	return
}
