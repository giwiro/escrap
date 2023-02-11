package database

import (
	"errors"
	"github.com/giwiro/escrap/modules/provider/dao"
	"github.com/giwiro/escrap/modules/provider/model"
	"github.com/giwiro/escrap/modules/provider/request"
	"gorm.io/gorm"
	"time"
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

func (spd *scrapProviderPgDao) FindProviderByName(name string) (provider model.ScrapProvider, err error) {
	provider = model.Unset

	switch name {
	case "amazon":
		provider = model.Amazon
		break
	case "ebay":
		provider = model.Ebay
		break
	}

	return
}

func (spd *scrapProviderPgDao) FindProviderByDomain(domain string) (provider model.ScrapProvider, err error) {
	var scrapProvider ScrapProviderDomainEntity
	provider = model.Unset

	query := spd.db.Where("domain = ?", domain).First(&scrapProvider)

	err = query.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Unset, nil
		}
		return
	}

	provider = model.ScrapProvider(scrapProvider.ScrapProviderId)

	return
}

func (spd *scrapProviderPgDao) FindProduct(vendorId string) (*model.ScrapProduct, error) {
	var scrapProduct ScrapProductEntity

	query := spd.db.Where("vendor_id = ?", vendorId).First(&scrapProduct)

	if err := query.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return scrapProduct.MapToModel(), nil
}

func (spd *scrapProviderPgDao) InsertProduct(request *request.InsertProductRequest) (*model.ScrapProduct, error) {
	product := &ScrapProductEntity{
		ScrapProviderId: uint(request.ScrapProvider),
		Name:            request.Name,
		Url:             request.Url,
		Price:           request.Price,
		Height:          request.Height,
		Length:          request.Length,
		Weight:          request.Weight,
		Width:           request.Width,
		VendorId:        request.VendorId,
		Description:     request.Description,
		ImageUrl:        request.ImageUrl,
		LastScrappedAt:  time.Now(),
	}
	spd.db.Create(&product)

	return product.MapToModel(), nil
}

func (spd *scrapProviderPgDao) FindLastResult(request *request.FindLastResultRequest) (*model.ScrapResult, error) {
	var resultEntity = ScrapResultEntity{}
	query := spd.db.Where("scrap_product_id = ?", request.ProductId).Order("created_at desc").First(&resultEntity)

	if err := query.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return resultEntity.MapToModel(), nil
}

func (spd *scrapProviderPgDao) InsertResult(request *request.InsertResultRequest) (*model.ScrapResult, error) {
	resultEntity := ScrapResultEntity{
		ScrapProductId:     request.ProductId,
		ScrapResultStateId: uint(request.StateId),
		ApiResult:          request.ApiResult,
		ApiResult2:         request.ApiResult2,
	}

	spd.db.Create(&resultEntity)

	return resultEntity.MapToModel(), nil
}

func (spd *scrapProviderPgDao) UpdateProduct(request *request.UpdateProductRequest) (*model.ScrapProduct, error) {
	var productEntity = ScrapProductEntity{}
	query := spd.db.Where("scrap_product_id = ?", request.ProductId).First(&productEntity)

	if err := query.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	spd.db.Model(&productEntity).Updates(ScrapProductEntity{
		Name:           request.Name,
		Url:            request.Url,
		Price:          request.Price,
		Height:         request.Height,
		Length:         request.Length,
		Weight:         request.Weight,
		Width:          request.Width,
		Description:    request.Description,
		ImageUrl:       request.ImageUrl,
		LastScrappedAt: request.LastScrappedAt,
	})

	return productEntity.MapToModel(), nil
}
