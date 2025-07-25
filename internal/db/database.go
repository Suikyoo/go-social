package db

import (
	"database/sql"
	"time"
  _ "github.com/go-sql-driver/mysql"
)

type Options struct {
  maxOpenConns int
  maxIdleConns int
  connMaxLifetime time.Duration

}

//since it does not fall into any interface, I don't bother using addresses for its operation
func NewOptions() Options {
  options := Options{
    maxOpenConns: 10,
    maxIdleConns: 10,
    connMaxLifetime: time.Minute * 3,
  }

  return options

}

func New(dataSource string) (*sql.DB, error){
  options := NewOptions()
  db, err := NewWithOptions(dataSource, options)
  return db, err

}

func NewWithOptions(dataSource string, options Options) (*sql.DB, error){
  db, err := sql.Open("mysql", dataSource)

  if err != nil {
    return nil, err
  }

  db.SetMaxOpenConns(options.maxOpenConns)
  db.SetMaxIdleConns(options.maxIdleConns)
  db.SetConnMaxLifetime(options.connMaxLifetime)

  return db, nil

}
