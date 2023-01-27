package vendors

import "github.com/giwiro/escrap/modules/provider/model"

func NewScrapProviderVendor(provider model.ScrapProvider) model.ScrapProviderVendor {
	switch provider {
	case model.Amazon:
		return &scrapProviderAmazonVendor{}
	}

	return nil
}
