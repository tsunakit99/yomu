package usecase

import (
	"github.com/tsunakit99/yomu/domain/model"
	"github.com/tsunakit99/yomu/domain/repository"
)

type ArticleUsecase interface {
	GetAllArticles() ([]*model.Article, error)
}

type articleUsecase struct {
	repo repository.ArticleRepository
}

func NewArticleUsecase(repo repository.ArticleRepository) ArticleUsecase {
	return &articleUsecase{
		repo: repo,
	}
}

func (u *articleUsecase) GetAllArticles() ([]*model.Article, error) {
	articles, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return articles, nil
}
