package domain

import "time"

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`

	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	PublishedAt *string `json:"published_at"`

	Status string `json:"status"`
}

func NewPost() *Post {
	return &Post{
		ID:          0,
		Title:       "",
		Content:     "",
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
		PublishedAt: nil,
		Status:      PostStatusDraft.String(),
	}
}
