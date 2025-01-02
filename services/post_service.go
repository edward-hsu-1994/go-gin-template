package services

import (
	"go-gin-template/accesses"
	"go-gin-template/domain"
	"math"
)

type PostService struct {
	postRepository accesses.PostRepository
}

func NewPostService(
	postRepository accesses.PostRepository,
) *PostService {
	return &PostService{
		postRepository: postRepository,
	}
}

func (ps *PostService) ListPosts(offset int, limit int) (*domain.Paging[*domain.PostSummary], error) {
	posts, err := ps.postRepository.ListPosts(offset, limit)

	if err != nil {
		return nil, err
	}

	totalCount, err := ps.postRepository.TotalPostsCount()

	if err != nil {
		return nil, err
	}

	nextOffset := offset + limit
	prevOffset := offset - limit
	if prevOffset < 0 {
		prevOffset = 0
	}

	return &domain.Paging[*domain.PostSummary]{
		Offset: offset,
		Limit:  limit,
		Total:  totalCount,
		Data:   posts,

		HasNext:    offset+limit < totalCount,
		HasPrev:    offset > 0,
		NextOffset: nextOffset,
		PrevOffset: prevOffset,

		PageNumber: int((float64(offset) / float64(limit)) + 1),
		TotalPage:  int(math.Floor(float64(totalCount)/float64(limit)) + 1),
	}, nil
}

func (ps *PostService) GetPostById(id string) (*domain.Post, error) {
	return ps.postRepository.GetPostById(id)
}
