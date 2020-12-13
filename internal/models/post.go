package models

import (
	"time"
)

type Post struct {
	ID          int64
	Author      User
	Title       string
	Summary     string
	Content     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	PublishedAt time.Time
	Categories  []Category
}
