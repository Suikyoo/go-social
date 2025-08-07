package main

import (
	"log"
	"strings"

	"github.com/Suikyoo/go-social/internal/db"
	"github.com/Suikyoo/go-social/internal/env"
	"github.com/Suikyoo/go-social/internal/repository"
)

func main() {
  env.Load()

	//config
	cfg := config{
    server: serverConfig{
      addr: strings.Join([]string{env.GetString("IP_ADDR", "127.0.0.1"), env.GetString("PORT", "8080")}, ":"),
    },
    auth: authConfig{
      jwtSecret: env.GetBytes("JWT_SECRET", []byte{}),
      tokenExpiry: 3600,

    },
	}


  //database
	db_conn_pool, err := db.New(env.GetString("DATABASE_SRC", ""))
	if err != nil {
		log.Fatal(err)
	}

  //repository
	store := repository.NewStorage(db_conn_pool)

  //service layer is almost nonexistent (already coupled with the api layer)
	app := application{
		config: cfg,
    store: store,
	}

	mux := app.Mount()

	app.Run(mux)

}
