package repository

import (
	"context"
	"database/sql"

	"github.com/Suikyoo/go-social/internal/authutils"
)


type User struct {
	ID   int64  `json:"id"`
	Username string `json:"username"`
  Password authutils.Password `json:"-"`
}

type UserRepository interface {
	RegularRepository[User]
  GetByUsername(context.Context, string) (*User, error)
}

type SqlUserRepository struct {
	db *sql.DB
}

func (r *SqlUserRepository) Get(ctx context.Context, id int64) (*User, error) {
	query := `
  SELECT id, username, password
  FROM users
  WHERE id = $1
  `

  user := User{}

  ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
  defer cancel()

  err := r.db.QueryRowContext(
    ctx,
    query,
    id,
  ).Scan(
    &user.ID,
    &user.Username,
    &user.Password,

    //don't include the password on scan
    //although the password would be nil,
    //its json tag is "-" anyways so it won't get sent
    
    //update: nah, just include the password dawg
    //deal with the aftermaths later
  )

  if err != nil {
    return nil, err
  }

  return &user, nil

}

func (r *SqlUserRepository) GetByUsername (ctx context.Context, username string) (*User, error) {
  query := `
  SELECT id, username, password
  FROM users
  WHERE username = $1
  `
  user := User{}

  ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
  defer cancel()

  err := r.db.QueryRowContext(
    ctx, 
    query, 
    username,
  ).Scan(
    &user.ID,
    &user.Username,
    &user.Password,
  )

  if err != nil {
    return nil, err
  }

  return &user, nil

}

func (r *SqlUserRepository) Create(ctx context.Context, user *User) error {
  query := `
  INSERT INTO users (username, password)
  VALUES ($1, $2)
  RETURNING id
  `
  
  ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
  defer cancel()

  err := r.db.QueryRowContext(
    ctx, 
    query,
    user.Username,
    []byte(user.Password),
  ).Scan(
    &user.ID,
  )

  if err != nil {
    return err
  }

  return nil

}

func (r *SqlUserRepository) GetFeed(ctx context.Context, amt int8) ([]*User, error) {
	query := `
  SELECT id, username
  FROM users
  LIMIT $1
  `
	userFeed := make([]*User, 0)

  ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
  defer cancel()

	rows, err := r.db.QueryContext(ctx, query, amt)

	if err != nil {
		return []*User{}, err
	}

	for rows.Next() {
		user := User{}
		err = rows.Scan(
			&user.ID,
			&user.Username,
		)

		//comment on sibling, posts.go
		if err != nil {
			return []*User{}, err
		}
		userFeed = append(userFeed, &user)

	}
	return userFeed, nil

}
