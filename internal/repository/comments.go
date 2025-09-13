package repository

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type Comment struct {
  ID int64 `json:"id"`
  UserID int64 `json:"user_id"`
  PostID int64 `json:"post_id"`
  Username string `json:"username"`
  Content string `json:"content"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  Likes int64 `json:"likes"`
}

type CommentRepository interface {
  CreateRepository[Comment]
	GetFeedByPostID(context.Context, int64, int8) ([]*Comment, error)
}

type SqlCommentRepository struct {
  db *sql.DB
}

type TestCommentRepository struct {}

func (r *SqlCommentRepository) GetFeedByPostID(ctx context.Context, postID int64, amt int8) ([]*Comment, error) {
  query := `
  SELECT users.username, users.id, posts.id, comments.id, comments.content, comments.created_at, comments.updated_at
  FROM comments
  INNER JOIN users ON comments.user_id = users.id
  INNER JOIN posts ON comments.post_id = posts.id
	WHERE posts.id = $1
  LIMIT $2
  `
  ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
  defer cancel()

  rows, err := r.db.QueryContext(
    ctx,
    query,
		postID,
    amt,
  )
  if err != nil {
    return []*Comment{}, err
  }

  comments := make([]*Comment, 0)

  for rows.Next() {
    comment := Comment{}
    err = rows.Scan(
      &comment.Username,
      &comment.UserID,
      &comment.PostID,
      &comment.ID,
      &comment.Content,
      &comment.CreatedAt,
      &comment.UpdatedAt,
    )

    if err != nil {
      return []*Comment{}, err
    }
    comments = append(comments, &comment)
  }
  return comments, nil
  

}

func (r *SqlCommentRepository) Create(ctx context.Context, comment *Comment) error {
	query := `
	WITH inserted AS (
		INSERT INTO comments (post_id, user_id, content)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, created_at, updated_at, likes
	)
	SELECT inserted.id, username, created_at, updated_at, likes
	FROM inserted
	INNER JOIN users on inserted.user_id = users.id
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := r.db.QueryRowContext(
		ctx, 
		query,
		comment.PostID,
		comment.UserID,
		comment.Content,
	).Scan(
		&comment.ID,
		&comment.Username,
		&comment.CreatedAt,
		&comment.UpdatedAt,
		&comment.Likes,
	)

	if err != nil {
		return err
	}
	
	return nil
}

func (r *TestCommentRepository) Create(ctx context.Context, comment *Comment) error {
	log.Print("comment created")
	return nil
}

func (r *TestCommentRepository) GetFeedByPostID(ctx context.Context, _ int64, amt int8) ([]*Comment, error) {
	log.Print("comments fetched")
	return append(make([]*Comment, 0), &Comment{}), nil
}
