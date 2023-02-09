package scrap

import (
	"errors"
	"github.com/giwiro/escrap/config"
	"github.com/giwiro/escrap/modules/provider/dao"
	"github.com/giwiro/escrap/modules/provider/model"
	"github.com/giwiro/escrap/modules/provider/request"
	"github.com/giwiro/escrap/modules/provider/vendors"
	log "github.com/sirupsen/logrus"
	"time"
)

type UseCase interface {
	ScrapUrl(url string, provider model.ScrapProvider) (*model.ScrapResult, *model.ScrapProduct, error)
}
type useCase struct {
	scrapProviderDao dao.ScrapProviderDao
}

func NewUseCase(scrapProviderDao dao.ScrapProviderDao) UseCase {
	return &useCase{scrapProviderDao}
}

func (u *useCase) ScrapUrl(url string, provider model.ScrapProvider) (*model.ScrapResult, *model.ScrapProduct, error) {
	providerVendor, err := vendors.NewScrapProviderVendor(provider, &u.scrapProviderDao)

	if err != nil {
		return nil, nil, err
	}

	if providerVendor == nil {
		return nil, nil, errors.New("no provider vendor was found")
	}

	vendorId, vendorIdError := providerVendor.GetVendorId(url)

	if vendorIdError != nil {
		return nil, nil, errors.New("wrong url format, could not get vendor id")
	}

	product, productErr := u.scrapProviderDao.FindProduct(vendorId)

	if productErr != nil {
		return nil, nil, errors.New("could not retrieve product")
	}

	if product != nil {
		if time.Now().Sub(product.LastScrappedAt).Milliseconds() < config.Conf.Server.Ttl {
			result, resultErr := u.scrapProviderDao.FindLastResult(&request.FindLastResultRequest{
				ProductId: product.Id,
			})

			if resultErr != nil {
				return nil, nil, errors.New("could not retrieve result")
			}

			if result != nil {
				log.Debugf("[Amazon] Fetch from cache: %s", vendorId)
				return result, product, nil
			}
		}
	}

	return providerVendor.Scrap(url, product)
}
