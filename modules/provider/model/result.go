package model

type ScrapResult struct {
	Id               uint             `json:"-"`
	ScrapProductId   uint             `json:"-"`
	ScrapResultState ScrapResultState `json:"-"`
	ApiResult        []byte           `json:"apiResult"`
	ApiResult2       []byte           `json:"apiResult2,omitempty"`
}
