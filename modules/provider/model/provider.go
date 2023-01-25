package model

type ScrapProvider int

const (
	Amazon ScrapProvider = 1
	Ebay                 = 2
)

type ScrapProviderVendor interface {
	Scrap(url string) (*ScrapResult, *ScrapProduct, error)
}
