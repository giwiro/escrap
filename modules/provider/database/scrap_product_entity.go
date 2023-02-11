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
	Price           decimal.NullDecimal `sql:"type:decimal(12,2);"`
	Height          decimal.NullDecimal `sql:"type:decimal(12,2);"`
	Length          decimal.NullDecimal `sql:"type:decimal(12,2);"`
	Weight          decimal.NullDecimal `sql:"type:decimal(12,2);"`
	Width           decimal.NullDecimal `sql:"type:decimal(12,2);"`
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
	product := model.ScrapProduct{
		Id:             s.ScrapProductId,
		ScrapProvider:  model.ScrapProvider(s.ScrapProviderId),
		Name:           s.Name,
		Url:            s.Url,
		Price:          s.Price,
		Height:         s.Height,
		Length:         s.Length,
		Weight:         s.Weight,
		Width:          s.Width,
		VendorId:       s.VendorId,
		Description:    s.Description,
		ImageUrl:       s.ImageUrl,
		LastScrappedAt: s.LastScrappedAt,
		CreatedAt:      s.CreatedAt,
		UpdatedAt:      s.UpdatedAt,
	}

	return &product
}
