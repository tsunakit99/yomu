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
