package repository

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type Post struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Username  string    `json:"username"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostRepository interface {
	RegularRepository[Post]
}

type SqlPostRepository struct {
	db *sql.DB
}

type TestPostRepository struct {
}

func (r *SqlPostRepository) Create(ctx context.Context, post *Post) error {
	query := `
  INSERT INTO posts (title, content, user_id)
  VALUES ($1, $2, $3) 
  RETURNING id, created_at, updated_at
  `
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := r.db.QueryRowContext(
		ctx,
		query,
		post.Title,
		post.Content,
		post.UserID,
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *SqlPostRepository) Get(ctx context.Context, id int64) (*Post, error) {
	query := `
  SELECT posts.id, posts.user_id, users.username, posts.title, posts.content, posts.created_at, posts.updated_at
  FROM posts
  INNER JOIN users
  ON posts.user_id = users.id
  WHERE posts.id = $1
  `
	post := Post{}

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := r.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
		&post.ID,
		&post.UserID,
		&post.Username,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &post, nil

}

func (r *SqlPostRepository) GetFeed(ctx context.Context, amt int8) ([]*Post, error) {
	query := `
  SELECT posts.id, posts.user_id, users.username, posts.title, posts.content, posts.created_at, posts.updated_at
  FROM posts
  INNER JOIN users
  ON posts.user_id = users.id
  LIMIT $1
  `
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := r.db.QueryContext(
		ctx,
		query,
		amt,
	)

	if err != nil {
		return []*Post{}, err
	}

	defer rows.Close()


	postFeed := make([]*Post, 0)

	for rows.Next() {
		post := Post{}
		err = rows.Scan(
			&post.ID,
			&post.UserID,
			&post.Username,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.UpdatedAt,
		)

		//im just gonna return an empty slice if even one of the scan iterations fail
		if err != nil {
			return []*Post{}, err
		}

		postFeed = append(postFeed, &post)

	}

	return postFeed, nil

}

func (r *TestPostRepository) Create(ctx context.Context, post *Post) error {
	log.Print("post created")
	return nil
}

func (r *TestPostRepository) Get(ctx context.Context, id int64) (*Post, error) {
	log.Print("post fetched")
	return &Post{}, nil

}

func (r *TestPostRepository) GetFeed(ctx context.Context, amt int8) ([]*Post, error) {
	log.Print("posts fetched")
	return append(make([]*Post, 0), &Post{}), nil
}
