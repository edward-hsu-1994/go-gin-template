package domain

import "fmt"

type PostStatus int

const (
	PostStatusDeleted PostStatus = iota
	PostStatusDraft
	PostStatusPublished
)

var postStatusStrings = map[PostStatus]string{
	PostStatusDeleted:   "deleted",
	PostStatusDraft:     "draft",
	PostStatusPublished: "published",
}

var validStatuses = map[PostStatus]bool{
	PostStatusDeleted:   true,
	PostStatusDraft:     true,
	PostStatusPublished: true,
}

func isValidStatus(status PostStatus) bool {
	_, ok := validStatuses[status]
	return ok
}

func (status PostStatus) String() string {
	return postStatusStrings[status]
}

func ParsePostStatus(status string) (PostStatus, error) {
	for k, v := range postStatusStrings {
		if v == status {
			return k, nil
		}
	}

	return PostStatusDraft, ErrorInvalidPostStatus.New(fmt.Sprintf("Invalid post status: %s", status))
}
