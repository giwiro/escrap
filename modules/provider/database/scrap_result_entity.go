package database

import (
	"github.com/giwiro/escrap/modules/provider/model"
	"time"
)
import "gorm.io/datatypes"

type ScrapResultEntity struct {
	ScrapResultId      uint `gorm:"primaryKey"`
	ScrapProductId     uint
	ScrapResultStateId uint
	ApiResult          datatypes.JSON `gorm:"column:api_result" sql:"type:jsonb"`
	ApiResult2         datatypes.JSON `gorm:"column:api_result_2" sql:"type:jsonb"`
	ApiResult3         datatypes.JSON `gorm:"column:api_result_3" sql:"type:jsonb"`
	CreatedAt          time.Time
}

func (s ScrapResultEntity) TableName() string {
	return "escrap.scrap_result"
}

func (s ScrapResultEntity) MapToModel() *model.ScrapResult {
	return &model.ScrapResult{
		Id:               s.ScrapResultId,
		ScrapProductId:   s.ScrapProductId,
		ScrapResultState: model.ScrapResultState(s.ScrapResultStateId),
		ApiResult:        s.ApiResult,
		ApiResult2:       s.ApiResult2,
	}
}
