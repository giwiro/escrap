package dao

import "github.com/giwiro/escrap/modules/provider/model"

type ScrapProviderDao interface {
	FindProviderByDomain(domain string) (model.ScrapProvider, error)
}
