package usecase

import (
	"strings"

	"github.com/tsunakit99/yomu/domain/model"
	"github.com/tsunakit99/yomu/domain/repository"
)

type ArticleUsecase interface {
	GetAllArticles() ([]*model.Article, error)
	GetArticleBySlug(slug string) (*model.ArticleDetail, error)
	CreateArticle(input *model.ArticleInput) error
	UpdateArticle(slug string, input *model.ArticleInput) error
	DeleteArticle(slug string) error
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
	return u.repo.GetAll()
}

func (u *articleUsecase) GetArticleBySlug(slug string) (*model.ArticleDetail, error) {
	return u.repo.GetBySlug(slug)
}

func (u *articleUsecase) CreateArticle(input *model.ArticleInput) error {
	slug := generateSlug(input.Title)

	return u.repo.Create(slug, input)
}

func (u *articleUsecase) UpdateArticle(slug string, input *model.ArticleInput) error {
	return u.repo.Update(slug, input)
}

func (u *articleUsecase) DeleteArticle(slug string) error {
	return u.repo.Delete(slug)
}

// 🔥 slug自動生成（超シンプル版）
func generateSlug(title string) string {
	return strings.ReplaceAll(strings.ToLower(title), " ", "-")
}
