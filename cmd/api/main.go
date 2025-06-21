package main

import (
	"log"
	"os"
	"example.com/social/internal/env"
)


func main() {
	conf := config{
		addr: env,
	}
	app := &application{
		config: conf,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
