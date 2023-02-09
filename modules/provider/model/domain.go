package model

type ScrapProviderDomain struct {
	Id            uint          `json:"-"`
	ScrapProvider ScrapProvider `json:"-"`
	Domain        string        `json:"domain"`
}
