package model

type ArticleInput struct {
	Title   string   `json:"title" validate:"required"`
	Date    string   `json:"date" validate:"required"`
	Tags    []string `json:"tags" validate:"required,dive,required"`
	Content string   `json:"content" validate:"required"`
}
