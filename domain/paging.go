package domain

type Paging[T any] struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`

	Total int `json:"total"`

	HasNext bool `json:"has_next"`
	HasPrev bool `json:"has_prev"`

	NextOffset int `json:"next_offset"`
	PrevOffset int `json:"prev_offset"`

	Data []T `json:"data"`

	PageNumber int `json:"page_number"`
	TotalPage  int `json:"total_page"`
}
