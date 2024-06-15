package dto

type ArticleDTO struct {
	Id         uint   `json:"articleId"`
	Title      string `json:"title"`
	Url        string `json:"url"`
	Tags       string `json:"tags"`
	CategoryId int    `json:"categoryId"`
	Status     bool   `json:"status"`
}
