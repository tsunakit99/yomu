package infra

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/tsunakit99/yomu/domain/model"
	"github.com/tsunakit99/yomu/domain/repository"
	"github.com/yuin/goldmark"
)

type MarkdownArticleRepository struct {
	PostDir string
}

func NewMarkdownArticleRepository(postDir string) repository.ArticleRepository {
	return &MarkdownArticleRepository{PostDir: postDir}
}

type frontMatterData struct {
	Title string   `yaml:"title"`
	Date  string   `yaml:"date"`
	Tags  []string `yaml:"tags"`
}

func (r *MarkdownArticleRepository) GetAll() ([]*model.Article, error) {
	files, err := filepath.Glob(filepath.Join(r.PostDir, "*.md"))
	if err != nil {
		return nil, err
	}

	var articles []*model.Article

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			continue // skip faulty files
		}

		var fm frontMatterData
		_, err = frontmatter.Parse(strings.NewReader(string(content)), &fm)
		if err != nil {
			continue // skip if no frontmatter
		}

		slug := strings.TrimSuffix(filepath.Base(file), ".md")

		articles = append(articles, &model.Article{
			Slug:  slug,
			Title: fm.Title,
			Date:  fm.Date,
			Tags:  fm.Tags,
		})
	}

	return articles, nil
}

func (r *MarkdownArticleRepository) GetBySlug(slug string) (*model.ArticleDetail, error) {
	filePath := filepath.Join(r.PostDir, slug+".md")
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var fm frontMatterData
	body, err := frontmatter.Parse(strings.NewReader(string(content)), &fm)
	if err != nil {
		return nil, err
	}

	var buf strings.Builder
	if err := goldmark.Convert(body, &buf); err != nil {
		return nil, err
	}

	return &model.ArticleDetail{
		Slug:        slug,
		Title:       fm.Title,
		Date:        fm.Date,
		Tags:        fm.Tags,
		ContentHTML: buf.String(),
	}, nil
}
