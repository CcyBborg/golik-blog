package store

// TODO: Contains lots of buisness logic/remove later to Repository
import (
	"database/sql"
	"fmt"

	"github.com/CcyBborg/golik-blog/internal/models"

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
	getUserByIDPattern             = `SELECT id, username FROM "user" WHERE id = %d;`
	insertPostQuery                = `INSERT INTO "post" (author_id, title, summary, content) VALUES (%d, '%s', '%s', '%s') RETURNING id;`
	insertPostCategory             = `INSERT INTO "post_category" (post_id, category_id) VALUES (%d, %d);`
	getPostQueryPattern            = `SELECT id, author_id, title, summary, content, created_at, updated_at, published_at FROM "post" WHERE id = $1;`
	updatePostQueryPattern         = `UPDATE post SET (title, summary, content) = ($1, $2, $3) WHERE id = $4;`
	publishPostQueryPattern        = `UPDATE post SET published_at = NOW() WHERE id = $1;`
	deletePostCategoryQueryPattern = `DELETE FROM post_category WHERE post_id = $1;`
	deletePostQueryPattern         = `DELETE FROM post WHERE id = $1;`
	getCommentsQueryPattern        = `SELECT id, author_id, created_at, content FROM post_comment WHERE post_id = $1 ORDER BY created_at;`
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

func (s *Store) GetPosts(opts Opts) ([]models.Post, error) {
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

func (s *Store) InsertPost(post models.Post) (postID int64, err error) {
	query := fmt.Sprintf(insertPostQuery, post.Author.ID, post.Title, post.Summary, post.Content)

	if err = s.db.QueryRow(query).Scan(&postID); err != nil {
		return postID, err
	}

	for _, category := range post.Categories {
		query := fmt.Sprintf(insertPostCategory, postID, category.ID)
		if _, err = s.db.Exec(query); err != nil {
			return postID, err
		}
	}

	return postID, err
}

func (s *Store) GetPost(postID int64) (post models.Post, err error) {
	var authorID int64
	var updatedAt, publishedAt sql.NullTime
	err = s.db.QueryRow(getPostQueryPattern, postID).Scan(
		&post.ID, &authorID, &post.Title, &post.Summary,
		&post.Content, &post.CreatedAt, &updatedAt, &publishedAt)
	if err != nil {
		return post, err
	}

	if updatedAt.Valid {
		post.UpdatedAt = updatedAt.Time
	}
	if publishedAt.Valid {
		post.PublishedAt = publishedAt.Time
	}
	if post.Categories, err = s.getPostCategories(post.ID); err != nil {
		return post, err
	}
	if post.Author, err = s.getUserByID(authorID); err != nil {
		return post, err
	}

	return post, nil
}

func (s *Store) UpdatePost(post models.Post, publish bool) error {
	if _, err := s.db.Exec(updatePostQueryPattern, post.Title, post.Summary, post.Content, post.ID); err != nil {
		return err
	}

	if publish {
		_, err := s.db.Exec(publishPostQueryPattern, post.ID)
		return err
	}

	return nil
}

func (s *Store) DeletePost(postID int64) error {
	_, err := s.db.Exec(deletePostCategoryQueryPattern, postID)
	if err != nil {
		return nil
	}

	_, err = s.db.Exec(deletePostQueryPattern, postID)

	return err
}

func (s *Store) GetComments(postID int64) ([]models.Comment, error) {
	rows, err := s.db.Query(getCommentsQueryPattern, postID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	comments := make([]models.Comment, 0)
	for rows.Next() {
		comment := models.Comment{}
		var authorID int64
		if err := rows.Scan(&comment.ID, &authorID, &comment.CreatedAt, &comment.Content); err != nil {
			return nil, err
		}
		if comment.Author, err = s.getUserByID(authorID); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
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
