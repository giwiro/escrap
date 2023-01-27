package provider

import (
	"github.com/giwiro/escrap/modules/provider/dao"
	"github.com/giwiro/escrap/modules/provider/model"
	log "github.com/sirupsen/logrus"
	"net/url"
)

type UseCase interface {
	FindProviderByUrl(_url string) (model.ScrapProvider, error)
}
type useCase struct {
	scrapProviderDao dao.ScrapProviderDao
}

func NewUseCase(scrapProviderDao dao.ScrapProviderDao) UseCase {
	return &useCase{scrapProviderDao}
}

func (u *useCase) FindProviderByUrl(_url string) (model.ScrapProvider, error) {
	parsedUrl, _ := url.Parse(_url)
	domain := parsedUrl.Hostname()

	log.Debugf("hostname => %s", parsedUrl.Hostname())

	provider, err := u.scrapProviderDao.FindProviderByDomain(domain)

	if err != nil {
		return model.Unset, err
	}

	return provider, nil
}
