package reaction

import "time"

type Reaction struct {
	ID        int
	UserID    int
	PostID    *int
	CommentID *int
	Type      int// "like" or "dislike"
	CreatedAt time.Time
}
