package repository

import "github.com/tsunakit99/yomu/domain/model"

type StatRepository interface {
	GetStats(slug string) (*model.ArticleStat, error)
}
