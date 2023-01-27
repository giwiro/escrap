package model

type ScrapProvider uint

const (
	Unset  ScrapProvider = 0
	Amazon               = 1
	Ebay                 = 2
)

type ScrapProviderVendor interface {
	Scrap(url string) (*ScrapResult, *ScrapProduct, error)
}

func GetScrapProvider(id uint) ScrapProvider {
	switch id {
	case 1:
		return Amazon
	case 2:
		return Ebay
	}
	return 0
}
