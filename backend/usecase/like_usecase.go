package usecase

import (
	"github.com/tsunakit99/yomu/domain/repository"
)

type LikeUsecase interface {
	AddLike(slug string) error
}

type likeUsecase struct {
	repo repository.LikeRepository
}

func NewLikeUsecase(repo repository.LikeRepository) LikeUsecase {
	return &likeUsecase{
		repo: repo,
	}
}

func (u *likeUsecase) AddLike(slug string) error {
	return u.repo.IncrementLike(slug)
}
