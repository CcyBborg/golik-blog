package models

import (
	"time"
)

type Comment struct {
	ID        int64
	PostID    int64
	Author    User
	CreatedAt time.Time
	Content   string
}
