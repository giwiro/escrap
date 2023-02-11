package model

type ScrapResult struct {
	Id               uint                   `json:"-"`
	ScrapProductId   uint                   `json:"-"`
	ScrapResultState ScrapResultState       `json:"-"`
	ApiResult        map[string]interface{} `json:"apiResult"`
	ApiResult2       map[string]interface{} `json:"apiResult2,omitempty"`
}
