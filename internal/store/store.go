package store

import (
	"database/sql"
	"fmt"

	"github.com/CcyBborg/golik-blog/internal/models"
	"github.com/CcyBborg/golik-blog/internal/services/posts"

	_ "github.com/lib/pq"
)

const (
	getPostsQueryPattern = `SELECT id, author_id, title, summary, content, created_at, updated_at, published_at
						FROM "post"
						WHERE published_at IS NOT NULL
						ORDER BY updated_at %s
						LIMIT %d OFFSET %d;`
	getCategoriesForPostPattern = `SELECT category.id, category.title FROM category, post_category
						WHERE category.id = post_category.category_id AND post_category.post_id = %d;`
	getUserByIDPattern = `SELECT id, username FROM "user" WHERE id = %d;`
)

type Store struct {
	config *Config
	db     *sql.DB
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) GetPosts(opts posts.Opts) ([]models.Post, error) {
	query := fmt.Sprintf(getPostsQueryPattern, opts.Sort.UpdatedAt, opts.Pagination.Limit,
		opts.Pagination.Offset)

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	postList := make([]models.Post, 0)
	for rows.Next() {
		post := models.Post{}
		var authorID int64
		var updatedAt, publishedAt sql.NullTime
		if err := rows.Scan(&post.ID, &authorID, &post.Title, &post.Summary,
			&post.Content, &post.CreatedAt, &updatedAt, &publishedAt); err != nil {
			return nil, err
		}
		if updatedAt.Valid {
			post.UpdatedAt = updatedAt.Time
		}
		if publishedAt.Valid {
			post.PublishedAt = publishedAt.Time
		}
		if post.Categories, err = s.getPostCategories(post.ID); err != nil {
			return nil, err
		}
		if post.Author, err = s.getUserByID(authorID); err != nil {
			return nil, err
		}
		postList = append(postList, post)
	}

	return postList, nil
}

func (s *Store) getPostCategories(postID int64) ([]models.Category, error) {
	query := fmt.Sprintf(getCategoriesForPostPattern, postID)

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	categoryList := make([]models.Category, 0)
	for rows.Next() {
		category := models.Category{}
		if err := rows.Scan(&category.ID, &category.Title); err != nil {
			return nil, err
		}
		categoryList = append(categoryList, category)
	}

	return categoryList, nil
}

func (s *Store) getUserByID(userID int64) (user models.User, err error) {
	query := fmt.Sprintf(getUserByIDPattern, userID)
	err = s.db.QueryRow(query).Scan(&user.ID, &user.Username)

	return user, err
}
