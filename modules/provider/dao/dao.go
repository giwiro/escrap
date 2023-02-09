package dao

import (
	"github.com/giwiro/escrap/modules/provider/model"
	"github.com/giwiro/escrap/modules/provider/request"
)

type ScrapProviderDao interface {
	FindProviderByDomain(domain string) (model.ScrapProvider, error)
	FindProduct(vendorId string) (*model.ScrapProduct, error)
	InsertProduct(request *request.InsertProductRequest) (*model.ScrapProduct, error)
	FindLastResult(request *request.FindLastResultRequest) (*model.ScrapResult, error)
	InsertResult(request *request.InsertResultRequest) (*model.ScrapResult, error)
	UpdateProduct(request *request.UpdateProductRequest) (*model.ScrapProduct, error)
}
