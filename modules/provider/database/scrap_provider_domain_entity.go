package database

type ScrapProviderDomainEntity struct {
	ScrapProviderDomainId uint `gorm:"primaryKey"`
	ScrapProviderId       uint
	Domain                string
}

func (s ScrapProviderDomainEntity) TableName() string {
	return "escrap.scrap_provider_domain"
}
