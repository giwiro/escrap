package provider

import (
	"github.com/giwiro/escrap/modules/provider/dao"
	"github.com/giwiro/escrap/modules/provider/model"
	log "github.com/sirupsen/logrus"
	"net/url"
)

type UseCase interface {
	FindProviderByName(name string) (model.ScrapProvider, error)
	FindProviderByUrl(url string) (model.ScrapProvider, error)
}
type useCase struct {
	scrapProviderDao dao.ScrapProviderDao
}

func NewUseCase(scrapProviderDao dao.ScrapProviderDao) UseCase {
	return &useCase{scrapProviderDao}
}

func (u *useCase) FindProviderByName(name string) (provider model.ScrapProvider, err error) {
	provider = model.Unset

	provider, err = u.scrapProviderDao.FindProviderByName(name)

	log.Debugf("name => %s", name)

	if err != nil {
		return
	}

	return
}

func (u *useCase) FindProviderByUrl(_url string) (provider model.ScrapProvider, err error) {
	provider = model.Unset
	parsedUrl, err := url.Parse(_url)
	domain := parsedUrl.Hostname()

	log.Debugf("hostname => %s", parsedUrl.Hostname())

	provider, err = u.scrapProviderDao.FindProviderByDomain(domain)

	if err != nil {
		return
	}

	return
}
