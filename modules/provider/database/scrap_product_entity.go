package database

import (
	"github.com/giwiro/escrap/modules/provider/model"
	"github.com/shopspring/decimal"
	"time"
)

type ScrapProductEntity struct {
	ScrapProductId  uint `gorm:"primaryKey"`
	ScrapProviderId uint
	Name            string
	Url             string
	Price           decimal.Decimal `json:"amount" sql:"type:decimal(12,2);"`
	VendorId        string
	Description     string
	ImageUrl        string
	LastScrappedAt  time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (s ScrapProductEntity) TableName() string {
	return "escrap.scrap_product"
}

func (s ScrapProductEntity) MapToModel() *model.ScrapProduct {
	return &model.ScrapProduct{
		Id:             s.ScrapProductId,
		ScrapProvider:  model.ScrapProvider(s.ScrapProviderId),
		Name:           s.Name,
		Url:            s.Url,
		Price:          s.Price.StringFixedBank(2),
		VendorId:       s.VendorId,
		Description:    s.Description,
		ImageUrl:       s.ImageUrl,
		LastScrappedAt: s.LastScrappedAt,
		CreatedAt:      s.CreatedAt,
		UpdatedAt:      s.UpdatedAt,
	}
}
