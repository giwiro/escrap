package request

import (
	"github.com/giwiro/escrap/modules/provider/model"
	"github.com/shopspring/decimal"
	"time"
)

type InsertProductRequest struct {
	Name          string
	Url           string
	Price         decimal.NullDecimal
	Height        decimal.NullDecimal
	Length        decimal.NullDecimal
	Weight        decimal.NullDecimal
	Width         decimal.NullDecimal
	Description   string
	ImageUrl      string
	ScrapProvider model.ScrapProvider
	VendorId      string
}

type UpdateProductRequest struct {
	ProductId      uint
	Name           string
	Url            string
	Price          decimal.NullDecimal
	Height         decimal.NullDecimal
	Length         decimal.NullDecimal
	Weight         decimal.NullDecimal
	Width          decimal.NullDecimal
	Description    string
	ImageUrl       string
	LastScrappedAt time.Time
}

type FindLastResultRequest struct {
	ProductId uint
}
type InsertResultRequest struct {
	ProductId  uint
	StateId    model.ScrapResultState
	ApiResult  map[string]interface{}
	ApiResult2 map[string]interface{}
}
