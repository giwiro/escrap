package web

type SearchRequest struct {
	Keyword  string `json:"keyword" binding:"required"`
	Provider string `json:"provider" binding:"required,alpha"`
	Page     uint   `json:"page" binding:"required,numeric"`
}
