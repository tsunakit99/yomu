package model

type ArticleDetail struct {
	Slug        string   `json:"slug"`
	Title       string   `json:"title"`
	Date        string   `json:"date"`
	Tags        []string `json:"tags"`
	ContentHTML string   `json:"contentHtml"`
}
