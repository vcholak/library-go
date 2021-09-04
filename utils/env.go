package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// EnvVar returns the value of the key
func EnvVar(key string, testing bool) string {

  var err error
  // load .env file
  if testing {
    err = godotenv.Load("../.env")
  } else {
    err = godotenv.Load()
  }

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  return os.Getenv(key)
}
