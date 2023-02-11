package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type ScrapProduct struct {
	Id             uint                `json:"-"`
	ScrapProvider  ScrapProvider       `json:"-"`
	Name           string              `json:"name"`
	Url            string              `json:"url"`
	Price          decimal.NullDecimal `json:"price"`  // Units: dollars
	Height         decimal.NullDecimal `json:"height"` // Units: cm
	Length         decimal.NullDecimal `json:"length"` // Units: cm
	Weight         decimal.NullDecimal `json:"weight"` // Units: kg
	Width          decimal.NullDecimal `json:"width"`  // Units: cm
	VendorId       string              `json:"vendorId"`
	Description    string              `json:"description"`
	ImageUrl       string              `json:"imageUrl"`
	LastScrappedAt time.Time           `json:"lastScrappedAt"`
	CreatedAt      time.Time           `json:"createdAt"`
	UpdatedAt      time.Time           `json:"updatedAt"`
}
