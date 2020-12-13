package postshandler

import (
	"github.com/CcyBborg/golik-blog/internal/api/schema"
	"github.com/CcyBborg/golik-blog/internal/models"
)

func convertPosts(posts []models.Post) []schema.ListPost {
	schemaPosts := make([]schema.ListPost, len(posts))

	for i, post := range posts {
		schemaPosts[i] = convertPost(post)
	}

	return schemaPosts
}

func convertPost(post models.Post) schema.ListPost {
	return schema.ListPost{
		ID: post.ID,
		Author: schema.User{
			ID:       post.Author.ID,
			Username: post.Author.Username,
		},
		Title:       post.Title,
		Summary:     post.Summary,
		PublishedAt: post.PublishedAt,
		UpdatedAt:   post.UpdatedAt,
		Categories:  convertCategories(post.Categories),
	}
}

func convertCategories(categories []models.Category) []schema.Category {
	schemaCategories := make([]schema.Category, len(categories))

	for i, category := range categories {
		schemaCategories[i] = schema.Category{
			ID:    category.ID,
			Title: category.Title,
		}
	}

	return schemaCategories
}
