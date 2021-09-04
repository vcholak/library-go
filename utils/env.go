package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// EnvVar returns the value of the key
func EnvVar(key string) string {

  // load .env file
  err := godotenv.Load("../.env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  return os.Getenv(key)
}
