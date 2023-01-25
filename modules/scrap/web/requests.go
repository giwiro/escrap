package web

type ScrapRequest struct {
	Url string `json:"url" binding:"required"`
}
