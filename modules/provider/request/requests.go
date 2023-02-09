package request

import (
	"github.com/giwiro/escrap/modules/provider/model"
	"github.com/shopspring/decimal"
	"time"
)

type InsertProductRequest struct {
	Name          string
	Url           string
	Price         decimal.Decimal
	Description   string
	ImageUrl      string
	ScrapProvider model.ScrapProvider
	VendorId      string
}

type UpdateProductRequest struct {
	ProductId      uint
	Name           string
	Url            string
	Price          decimal.Decimal
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
	ApiResult  []byte
	ApiResult2 []byte
	ApiResult3 []byte
}
