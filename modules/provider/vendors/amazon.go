package vendors

import (
	"context"
	"errors"
	"github.com/davecgh/go-spew/spew"
	"github.com/giwiro/escrap/common"
	"github.com/giwiro/escrap/config"
	"github.com/giwiro/escrap/modules/provider/dao"
	"github.com/giwiro/escrap/modules/provider/model"
	"github.com/giwiro/escrap/modules/provider/request"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"github.com/utekaravinash/gopaapi5"
	"github.com/utekaravinash/gopaapi5/api"
	"regexp"
	"strings"
	"time"
)

type ScrapProviderAmazonVendor interface {
	model.ScrapProviderVendor
	ItemLookupByAsin(asin string) (*api.GetItemsResponse, error)
	GetAbsoluteParentAsin(item *api.Item) string
	GetVariations(asin string) ([]api.GetVariationsResponse, error)
	GetPrice(offers api.Offers) decimal.NullDecimal
	GetImage(images api.Images) string
	GetHeight(itemInfo api.ItemInfo) decimal.NullDecimal
	GetLength(itemInfo api.ItemInfo) decimal.NullDecimal
	GetWeight(itemInfo api.ItemInfo) decimal.NullDecimal
	GetWidth(itemInfo api.ItemInfo) decimal.NullDecimal
}

type scrapProviderAmazonVendor struct {
	client           *gopaapi5.Client
	scrapProviderDao *dao.ScrapProviderDao
}

var amazonRegex = regexp.MustCompile(`^(https?://)?[^.]+\.amazon\.com/.*/([A-Z0-9]{10})[/?]?`)

var commonItemResources = []api.Resource{
	// images
	//api.ImagesPrimarySmall,
	//api.ImagesPrimaryMedium,
	api.ImagesPrimaryLarge,
	//api.ImagesVariantsSmall,
	//api.ImagesVariantsMedium,
	api.ImagesVariantsLarge,
	// item info
	api.ItemInfoByLineInfo,
	api.ItemInfoContentInfo,
	api.ItemInfoContentRating,
	api.ItemInfoClassifications,
	api.ItemInfoExternalIds,
	api.ItemInfoFeatures,
	api.ItemInfoManufactureInfo,
	api.ItemInfoProductInfo,
	api.ItemInfoTechnicalInfo,
	api.ItemInfoTitle,
	api.ItemInfoTradeInInfo,
	// parent
	api.ParentASIN,
	// browse node info
	api.BrowseNodeInfoBrowseNodes,
	api.BrowseNodeInfoBrowseNodesAncestor,
	api.BrowseNodeInfoBrowseNodesSalesRank,
	api.BrowseNodeInfoWebsiteSalesRank,
	// offers
	api.OffersListingsAvailabilityMaxOrderQuantity,
	api.OffersListingsAvailabilityMessage,
	api.OffersListingsAvailabilityMinOrderQuantity,
	api.OffersListingsAvailabilityType,
	api.OffersListingsCondition,
	api.OffersListingsConditionSubCondition,
	api.OffersListingsDeliveryInfoIsAmazonFulfilled,
	api.OffersListingsDeliveryInfoIsFreeShippingEligible,
	api.OffersListingsDeliveryInfoIsPrimeEligible,
	api.OffersListingsDeliveryInfoShippingCharges,
	//api.OffersListingsIsBuyBoxWinner,
	//api.OffersListingsLoyaltyPointsPoints,
	api.OffersListingsMerchantInfo,
	api.OffersListingsPrice,
	api.OffersListingsProgramEligibilityIsPrimeExclusive,
	api.OffersListingsProgramEligibilityIsPrimePantry,
	api.OffersListingsPromotions,
	api.OffersListingsSavingBasis,
	api.OffersSummariesHighestPrice,
	api.OffersSummariesLowestPrice,
	api.OffersSummariesOfferCount,
}

func NewScrapProviderAmazonVendor(providerDao *dao.ScrapProviderDao) (ScrapProviderAmazonVendor, error) {
	client, err := gopaapi5.NewClient(
		config.Conf.Scrapper.Amazon.AccessKey,
		config.Conf.Scrapper.Amazon.SecretKey,
		config.Conf.Scrapper.Amazon.AssociateTag,
		api.UnitedStates,
	)

	if err != nil {
		return nil, err
	}

	return &scrapProviderAmazonVendor{client, providerDao}, nil
}
func (s *scrapProviderAmazonVendor) ItemLookupByAsin(asin string) (*api.GetItemsResponse, error) {
	log.Debugf("[Amazon] ItemLookupByAsin ASIN: %s", asin)
	params := api.GetItemsParams{
		ItemIds: []string{
			asin,
		},
		Resources: commonItemResources,
	}

	response, err := s.client.GetItems(context.Background(), &params)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *scrapProviderAmazonVendor) GetAbsoluteParentAsin(item *api.Item) string {
	if item == nil {
		return ""
	}

	if item.ParentASIN == "" {
		return item.ASIN
	}

	parentItem, parentItemErr := s.ItemLookupByAsin(item.ParentASIN)

	if parentItemErr != nil {
		log.Errorf("[Amazon] Could not obtain parent asin from: %s", item.ParentASIN)
		return ""
	}

	if len(parentItem.ItemsResult.Items) == 0 {
		return item.ASIN
	}

	return s.GetAbsoluteParentAsin(&parentItem.ItemsResult.Items[0])
}

func (s *scrapProviderAmazonVendor) GetVariations(asin string) (variations []api.GetVariationsResponse, err error) {
	var page = 1

	for {
		params := api.GetVariationsParams{
			ASIN:           asin,
			VariationPage:  page,
			VariationCount: 10,
			Resources:      commonItemResources,
		}

		log.Debugf("[Amazon] Get variation of ASIN=%s page=%d", asin, page)

		response, err := s.client.GetVariations(context.Background(), &params)

		if err != nil {
			log.Error(err)
			return nil, err
		}

		variations = append(variations, *response)

		log.Debugf("[Amazon] VariationSummary.PageCount=%d", response.VariationsResult.VariationSummary.PageCount)

		if response.VariationsResult.VariationSummary.PageCount == page ||
			response.VariationsResult.VariationSummary.PageCount == 0 {
			break
		}

		page += 1
	}

	return variations, err
}

func (s *scrapProviderAmazonVendor) GetPrice(offers api.Offers) decimal.NullDecimal {
	price := decimal.NullDecimal{
		Decimal: decimal.Zero,
		Valid:   false,
	}
	if len(offers.Listings) > 0 && offers.Listings[0].Price.Amount != 0 {
		return decimal.NewNullDecimal(decimal.NewFromFloat32(offers.Listings[0].Price.Amount))
	}
	return price
}

func (s *scrapProviderAmazonVendor) GetImage(images api.Images) string {
	if images.Primary.Large.URL != "" {
		return images.Primary.Large.URL
	}

	if len(images.Variants) > 0 && images.Variants[0].Large.URL != "" {
		return images.Variants[0].Large.URL
	}

	return ""
}

func (s *scrapProviderAmazonVendor) Scrap(url string, p *model.ScrapProduct) (*model.ScrapResult, *model.ScrapProduct, error) {
	var product = p

	vendorId, vendorIdError := s.GetVendorId(url)

	if vendorIdError != nil {
		return nil, nil, errors.New("wrong url format, could not get vendor id")
	}

	log.Infof("[Amazon] Scrapping vendorId: %s", vendorId)

	response, responseErr := s.ItemLookupByAsin(vendorId)

	if responseErr != nil {
		spew.Dump(responseErr)
		log.Errorf("Error looking by asin: %s", responseErr.Error())
		return nil, nil, responseErr
	}

	item := response.ItemsResult.Items[0]

	log.Debugf("[Amazon] itemASIN: %s", item.ASIN)

	parentASIN := s.GetAbsoluteParentAsin(&item)

	log.Debugf("[Amazon] parentASIN: %s", parentASIN)

	variations, variationsError := s.GetVariations(parentASIN)

	if variationsError != variationsError {
		log.Error(variationsError)
	}

	if p == nil {
		insertProduct, err := (*s.scrapProviderDao).InsertProduct(&request.InsertProductRequest{
			Name:          item.ItemInfo.Title.DisplayValue,
			Url:           item.DetailPageURL,
			Price:         s.GetPrice(item.Offers),
			Height:        s.GetHeight(item.ItemInfo),
			Length:        s.GetLength(item.ItemInfo),
			Weight:        s.GetWeight(item.ItemInfo),
			Width:         s.GetWidth(item.ItemInfo),
			Description:   strings.Join(item.ItemInfo.Features.DisplayValues, "\n"),
			ImageUrl:      s.GetImage(item.Images),
			ScrapProvider: model.Amazon,
			VendorId:      vendorId,
		})

		if err != nil {
			return nil, nil, err
		}

		product = insertProduct
	}

	insertRequest := request.InsertResultRequest{
		ProductId: product.Id,
		StateId:   model.Success,
	}

	if apiResult, err := common.StructToMap(response); err == nil {
		insertRequest.ApiResult = apiResult
	}

	if apiResult2, err := common.StructToMap(map[string][]api.GetVariationsResponse{
		"VariationsResponses": variations,
	}); err == nil {
		insertRequest.ApiResult2 = apiResult2
	}

	result, resultErr := (*s.scrapProviderDao).InsertResult(&insertRequest)

	if resultErr != nil {
		return nil, nil, resultErr
	}

	if p != nil {
		updatedProduct, updatedProductErr := (*s.scrapProviderDao).UpdateProduct(&request.UpdateProductRequest{
			ProductId:      p.Id,
			Name:           item.ItemInfo.Title.DisplayValue,
			Url:            item.DetailPageURL,
			Price:          s.GetPrice(item.Offers),
			Height:         s.GetHeight(item.ItemInfo),
			Length:         s.GetLength(item.ItemInfo),
			Weight:         s.GetWeight(item.ItemInfo),
			Width:          s.GetWidth(item.ItemInfo),
			Description:    strings.Join(item.ItemInfo.Features.DisplayValues, "\n"),
			ImageUrl:       s.GetImage(item.Images),
			LastScrappedAt: time.Now(),
		})

		if updatedProductErr != nil {
			return nil, nil, updatedProductErr
		}

		product = updatedProduct
	}

	return result, product, nil
}

func (s *scrapProviderAmazonVendor) Search(keyword string, page uint) (response map[string]interface{}, err error) {
	params := api.SearchItemsParams{
		Keywords:  keyword,
		ItemPage:  int(page),
		ItemCount: 10, // max value
		Resources: commonItemResources,
	}

	resp, err := s.client.SearchItems(context.Background(), &params)

	if err != nil {
		panic(err)
	}

	response, err = common.StructToMap(resp)

	return
}

func (s *scrapProviderAmazonVendor) GetVendorId(url string) (string, error) {
	matches := amazonRegex.FindStringSubmatch(url)

	if len(matches) != 3 {
		return "", errors.New("wrong url format, could not get vendor id")
	}

	return matches[2], nil
}

func (s *scrapProviderAmazonVendor) GetHeight(itemInfo api.ItemInfo) decimal.NullDecimal {
	var height = decimal.NullDecimal{
		Decimal: decimal.Zero,
		Valid:   false,
	}

	switch strings.ToLower(itemInfo.ProductInfo.ItemDimensions.Height.Unit) {
	case "inches":
		inchFactor := decimal.NewFromFloat32(2.54)
		inches := decimal.NewFromFloat32(itemInfo.ProductInfo.ItemDimensions.Height.DisplayValue)
		height = decimal.NewNullDecimal(inches.Mul(inchFactor))
		break
	}

	return height
}

func (s *scrapProviderAmazonVendor) GetLength(itemInfo api.ItemInfo) decimal.NullDecimal {
	var length = decimal.NullDecimal{
		Decimal: decimal.Zero,
		Valid:   false,
	}

	switch strings.ToLower(itemInfo.ProductInfo.ItemDimensions.Length.Unit) {
	case "inches":
		inchFactor := decimal.NewFromFloat32(2.54)
		inches := decimal.NewFromFloat32(itemInfo.ProductInfo.ItemDimensions.Length.DisplayValue)
		length = decimal.NewNullDecimal(inches.Mul(inchFactor))
		break
	}

	return length
}

func (s *scrapProviderAmazonVendor) GetWeight(itemInfo api.ItemInfo) decimal.NullDecimal {
	var weight = decimal.NullDecimal{
		Decimal: decimal.Zero,
		Valid:   false,
	}

	switch strings.ToLower(itemInfo.ProductInfo.ItemDimensions.Weight.Unit) {
	case "pounds":
		poundFactor := decimal.NewFromFloat32(0.453592)
		pounds := decimal.NewFromFloat32(itemInfo.ProductInfo.ItemDimensions.Weight.DisplayValue)
		weight = decimal.NewNullDecimal(pounds.Mul(poundFactor))
	}

	return weight
}

func (s *scrapProviderAmazonVendor) GetWidth(itemInfo api.ItemInfo) decimal.NullDecimal {
	var width = decimal.NullDecimal{
		Decimal: decimal.Zero,
		Valid:   false,
	}

	switch strings.ToLower(itemInfo.ProductInfo.ItemDimensions.Width.Unit) {
	case "inches":
		inchFactor := decimal.NewFromFloat32(2.54)
		inches := decimal.NewFromFloat32(itemInfo.ProductInfo.ItemDimensions.Width.DisplayValue)
		width = decimal.NewNullDecimal(inches.Mul(inchFactor))
		break
	}

	return width
}
