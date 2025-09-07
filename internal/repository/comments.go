package repository

import (
	"context"
	"database/sql"
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
  FeedRepository[Comment]
  CreateRepository[Comment]
}

type SqlCommentRepository struct {
  db *sql.DB
}

func (r *SqlCommentRepository) GetFeed(ctx context.Context, amt int8) ([]Comment, error) {
  query := `
  SELECT users.username, users.id, posts.id, comments.id, comments.content, comments.created_at, comments.updated_at
  FROM comments
  INNER JOIN users ON comments.user_id = users.id
  INNER JOIN posts ON comments.post_id = posts.id
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
    return []Comment{}, err
  }

  comments := make([]Comment, 0)

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
      return []Comment{}, err
    }
    comments = append(comments, comment)
  }

  return comments, nil
  

}
