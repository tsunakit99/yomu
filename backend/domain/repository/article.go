package repository

import "github.com/tsunakit99/yomu/domain/model"

type ArticleRepository interface {
	GetAll() ([]*model.Article, error)
}
