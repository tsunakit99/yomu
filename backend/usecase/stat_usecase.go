package usecase

import (
	"github.com/tsunakit99/yomu/domain/model"
	"github.com/tsunakit99/yomu/domain/repository"
)

type StatUsecase interface {
	GetArticleStats(slug string) (*model.ArticleStat, error)
}

type statUsecase struct {
	repo repository.StatRepository
}

func NewStatUsecase(repo repository.StatRepository) StatUsecase {
	return &statUsecase{repo: repo}
}

func (u *statUsecase) GetArticleStats(slug string) (*model.ArticleStat, error) {
	return u.repo.GetStats(slug)
}
