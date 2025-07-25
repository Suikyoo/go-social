package repository

import (
  "context"
)

type Post struct {

}

type PostRepository interface{
  Create(context.Context, *Post) error
  GetById(context.Context, int64) (*Post, error)
  Get(context.Context, int64) ([]*Post, error)
}
  
