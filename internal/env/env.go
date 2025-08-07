package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Load() {

  if os.Getenv("ENV") == "" {
    err := godotenv.Load()
    if err != nil {
      log.Fatalf("Local environment variable error on setup. error: %v", err.Error())
    }
  }

}

func GetString(key string, callback string) string {
  value :=  os.Getenv(key)
  if value == "" {
    return callback
  }
  return value
}

func GetInt(key string, callback int) int {
  value := os.Getenv(key)
  num, err := strconv.Atoi(value) 
  if err != nil {
    return callback
  }

  if value  == "" {
    return callback
  }

  return num
}

func GetBytes(key string, callback []byte) []byte {
  value := os.Getenv(key)

  if value == "" {
    return callback
  }

  return []byte(value)
}
