package model

import (
	"time"
)

type ScrapProduct struct {
	Id             uint          `json:"-"`
	ScrapProvider  ScrapProvider `json:"-"`
	Name           string        `json:"name"`
	Url            string        `json:"url"`
	Price          string        `json:"price"`
	VendorId       string        `json:"vendorId"`
	Description    string        `json:"description"`
	ImageUrl       string        `json:"imageUrl"`
	LastScrappedAt time.Time     `json:"lastScrappedAt"`
	CreatedAt      time.Time     `json:"createdAt"`
	UpdatedAt      time.Time     `json:"updatedAt"`
}
