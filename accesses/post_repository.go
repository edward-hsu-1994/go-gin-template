package accesses

import (
	"go-fiber-template/domain"
)

type PostRepository interface {
	ListPosts(offset int, limit int) ([]*domain.PostSummary, error)
	TotalPostsCount() (int, error)
	GetPostById(id string) (*domain.Post, error)
}
