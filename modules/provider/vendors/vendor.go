package vendors

import (
	"github.com/giwiro/escrap/modules/provider/dao"
	"github.com/giwiro/escrap/modules/provider/model"
)

func NewScrapProviderVendor(provider model.ScrapProvider, providerDao *dao.ScrapProviderDao) (model.ScrapProviderVendor, error) {
	switch provider {
	case model.Amazon:
		return NewScrapProviderAmazonVendor(providerDao)
	}

	return nil, nil
}
