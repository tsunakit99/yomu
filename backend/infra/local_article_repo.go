package infra

import (
	"github.com/tsunakit99/yomu/domain/model"
	"github.com/tsunakit99/yomu/domain/repository"
)

type LocalArticleRepository struct{}

func NewLocalArticleRepository() repository.ArticleRepository {
	return &LocalArticleRepository{}
}

func (r *LocalArticleRepository) GetAll() ([]*model.Article, error) {
	// 今は仮でハードコード
	return []*model.Article{
		{Slug: "hello-yomu", Title: "Yomuへようこそ", Date: "2024-04-27", Tags: []string{"welcome", "intro"}},
	}, nil
}

func (r *LocalArticleRepository) GetBySlug(slug string) (*model.ArticleDetail, error) {
	// 今は仮でハードコード
	if slug == "hello-yomu" {
		return &model.ArticleDetail{
			Slug:        "hello-yomu",
			Title:       "Yomuへようこそ",
			Date:        "2024-04-27",
			Tags:        []string{"welcome", "intro"},
			ContentHTML: "<p>YomuはMarkdownで記事を書くためのツールです。</p>",
		}, nil
	}
	return nil, nil
}

func (r *LocalArticleRepository) Create(slug string, input *model.ArticleInput) error {
	// 今は何もしない
	return nil
}
func (r *LocalArticleRepository) Update(slug string, input *model.ArticleInput) error {
	// 今は何もしない
	return nil
}
func (r *LocalArticleRepository) Delete(slug string) error {
	// 今は何もしない
	return nil
}
