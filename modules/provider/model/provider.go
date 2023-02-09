package model

type ScrapProvider uint

const (
	Unset  ScrapProvider = 0
	Amazon               = 1
	Ebay                 = 2
)

type ScrapProviderVendor interface {
	GetVendorId(url string) (string, error)
	Scrap(url string, product *ScrapProduct) (*ScrapResult, *ScrapProduct, error)
}
