package vendors

import (
	"github.com/giwiro/escrap/modules/provider/model"
	"regexp"
)

type ScrapProviderAmazonVendor interface {
	model.ScrapProviderVendor
}

type scrapProviderAmazonVendor struct {
}

var amazonRegex = regexp.MustCompile(`^(https?://)?[^.]+\.amazon\.com/.*/([A-Z0-9]{10})[/?]?`)

func (s *scrapProviderAmazonVendor) Scrap(url string) (*model.ScrapResult, *model.ScrapProduct, error) {
	matches := amazonRegex.FindStringSubmatch(url)

	vendorId := matches[2]

	return nil, nil, nil
}
