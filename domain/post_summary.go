package domain

type PostSummary struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Summary string `json:"summary"`

	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	PublishedAt *string `json:"published_at"`

	Status string `json:"status"`
}

func NewPostSummary(post *Post) *PostSummary {
	if post == nil {
		return nil
	}

	return &PostSummary{
		ID:          post.ID,
		Title:       post.Title,
		Summary:     post.Content[:200],
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		PublishedAt: post.PublishedAt,
		Status:      post.Status,
	}
}
