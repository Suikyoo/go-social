package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load()

  if err != nil {
    log.Fatal(err)

  }

  //metadata
  cfg := config{
    addr: os.Getenv("ADDR"),

  }

  //object components
  app := application{
    config: cfg,

  }

  mux := app.Mount()

  app.Run(mux)


}
