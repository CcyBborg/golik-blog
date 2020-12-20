package schema

import (
	"github.com/CcyBborg/golik-blog/internal/models"
)

func ConvertPosts(posts []models.Post) []ListPost {
	schemaPosts := make([]ListPost, len(posts))

	for i, post := range posts {
		schemaPosts[i] = ListPost{
			ID: post.ID,
			Author: User{
				ID:       post.Author.ID,
				Username: post.Author.Username,
			},
			Title:      post.Title,
			Summary:    post.Summary,
			UpdatedAt:  post.UpdatedAt,
			Categories: ConvertCategories(post.Categories),
		}

		if !post.PublishedAt.IsZero() {
			schemaPosts[i].PublishedAt = post.PublishedAt
		}
	}

	return schemaPosts
}

func ConvertPost(post models.Post) Post {
	schemaPost := Post{
		ID: post.ID,
		Author: User{
			ID:       post.Author.ID,
			Username: post.Author.Username,
		},
		Title:      post.Title,
		Summary:    post.Summary,
		UpdatedAt:  post.UpdatedAt,
		Categories: ConvertCategories(post.Categories),
		Content:    post.Content,
	}

	if !post.PublishedAt.IsZero() {
		schemaPost.PublishedAt = post.PublishedAt
	}

	return schemaPost
}

func ConvertCategories(categories []models.Category) []Category {
	schemaCategories := make([]Category, len(categories))

	for i, category := range categories {
		schemaCategories[i] = Category{
			ID:    category.ID,
			Title: category.Title,
		}
	}

	return schemaCategories
}

func ConvertComments(comments []models.Comment) []Comment {
	schemaComments := make([]Comment, len(comments))

	for i, comment := range comments {
		schemaComments[i] = Comment{
			ID:        comment.ID,
			CreatedAt: comment.CreatedAt,
			Author: User{
				ID:       comment.Author.ID,
				Username: comment.Author.Username,
			},
			Content: comment.Content,
		}
	}

	return schemaComments
}
