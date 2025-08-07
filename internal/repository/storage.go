package repository

import (
	"context"
	"database/sql"
	"time"
)

//gets at most n [int8] amount of records 
type FeedRepository[T any] interface {
	GetFeed(context.Context, int8) ([]*T, error)
}

type CreateRepository[T any] interface {
	Create(context.Context, *T) error
}

//gets record using id field
type GetRepository[T any] interface {
  Get(context.Context, int64) (*T, error)
}

type RegularRepository[T any] interface {
  CreateRepository[T]
  GetRepository[T]
  FeedRepository[T]

}

type Storage struct {
	Posts PostRepository
  Users UserRepository
}

var (
  QueryTimeoutDuration = time.Second * 5
)

func NewStorage(db *sql.DB) *Storage {
	storage := Storage{
		Posts: &SqlPostRepository{db: db},
    Users: &SqlUserRepository{db: db},
	}
  return &storage

}

