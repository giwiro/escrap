package scrap

import (
	"errors"
	"github.com/giwiro/escrap/modules/provider/model"
	"github.com/giwiro/escrap/modules/provider/vendors"
)

type UseCase interface {
	ScrapUrl(url string, provider model.ScrapProvider) (*model.ScrapResult, *model.ScrapProduct, error)
}
type useCase struct {
}

func NewUseCase() UseCase {
	return &useCase{}
}

func (u *useCase) ScrapUrl(url string, provider model.ScrapProvider) (*model.ScrapResult, *model.ScrapProduct, error) {
	providerVendor := vendors.NewScrapProviderVendor(provider)

	if providerVendor == nil {
		return nil, nil, errors.New("No provider vendor was found")
	}

	return providerVendor.Scrap(url)
}
