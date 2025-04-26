package repository

type LikeRepository interface {
	IncrementLike(slug string) error
}
