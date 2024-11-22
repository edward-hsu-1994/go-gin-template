package accesses

import (
	"fmt"
	"go-fiber-template/domain"
	"strconv"
)

var mockPostsList = []*domain.Post{
	{
		ID:          1,
		Title:       "Title 1",
		Content:     "Content 1Content 1Content 1Content 1Content 1Content 1Content 1Content 1Content 1Content 1Content 1Content 1",
		CreatedAt:   "2021-01-01T00:00:00Z",
		UpdatedAt:   "2021-01-01T00:00:00Z",
		PublishedAt: nil,
		Status:      domain.PostStatusDraft.String(),
	},
	{
		ID:          2,
		Title:       "Title 2",
		Content:     "Content 2Content 2Content 2Content 2Content 2Content 2Content 2Content 2Content 2Content 2Content 2Content 2",
		CreatedAt:   "2021-01-01T00:00:00Z",
		UpdatedAt:   "2021-01-01T00:00:00Z",
		PublishedAt: nil,
		Status:      domain.PostStatusDraft.String(),
	},
}

type MockPostRepository struct {
}

func NewMockPostRepository() PostRepository {
	return &MockPostRepository{}
}

func (mr *MockPostRepository) ListPosts(
	offset int,
	limit int,
) ([]*domain.PostSummary, error) {
	queryResult := mockPostsList[offset : offset+limit]

	posts := make([]*domain.PostSummary, 0)

	for _, post := range queryResult {
		posts = append(posts, domain.NewPostSummary(post))
	}

	return posts, nil
}

func (mr *MockPostRepository) TotalPostsCount() (int, error) {
	queryResult := len(mockPostsList)

	return queryResult, nil
}

func (mr *MockPostRepository) GetPostById(id string) (*domain.Post, error) {
	for _, post := range mockPostsList {
		postId := strconv.Itoa(post.ID)
		if id == postId {
			return post, nil
		}
	}

	return nil, domain.ErrorPostNotFound.New(fmt.Sprintf("Post with ID %s not found", id))
}
