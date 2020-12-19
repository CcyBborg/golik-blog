package schema

import (
	"time"
)

type Category struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

type ListPost struct {
	ID            int64      `json:"id"`
	Author        User       `json:"author"`
	Title         string     `json:"title"`
	Summary       string     `json:"summary"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	PublishedAt   time.Time  `json:"publishedAt,omitempty"`
	Categories    []Category `json:"categoryList"`
	CommentsCount int64      `json:"commentsCount"`
}

type Post struct {
	ID            int64      `json:"id"`
	Author        User       `json:"author"`
	Title         string     `json:"title"`
	Summary       string     `json:"summary"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	PublishedAt   time.Time  `json:"publishedAt,omitempty"`
	Categories    []Category `json:"categoryList"`
	CommentsCount int64      `json:"commentsCount"`
	Content       string     `json:"content"`
}
