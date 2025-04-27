package repository

import "github.com/tsunakit99/yomu/domain/model"

type ArticleRepository interface {
	GetAll() ([]*model.Article, error)
	GetBySlug(slug string) (*model.ArticleDetail, error)
	Create(slug string, input *model.ArticleInput) error
	Update(slug string, input *model.ArticleInput) error
	Delete(slug string) error
}
