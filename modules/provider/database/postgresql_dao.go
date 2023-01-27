package database

import (
	"github.com/giwiro/escrap/modules/provider/dao"
	"github.com/giwiro/escrap/modules/provider/model"
	"gorm.io/gorm"
)

type ScrapProviderPgDao interface {
	dao.ScrapProviderDao
}

type scrapProviderPgDao struct {
	db *gorm.DB
}

func NewScrapProviderPgDao(db *gorm.DB) ScrapProviderPgDao {
	return &scrapProviderPgDao{db}
}

func (spd *scrapProviderPgDao) FindProviderByDomain(domain string) (model.ScrapProvider, error) {
	var scrapProvider ScrapProviderDomainEntity

	query := spd.db.Where("domain = ?", domain).First(&scrapProvider)

	if err := query.Error; err != nil {
		return model.Unset, err
	}

	return model.GetScrapProvider(scrapProvider.ScrapProviderId), nil
}
